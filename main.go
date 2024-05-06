package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"codequarry/app"

	"github.com/joho/godotenv"
)

func obfuscateJavaScript(inputPath, outputPath string) {
	// Ensure the path to the Python executable and the script is correct
	cmd := exec.Command("npx", "javascript-obfuscator", inputPath, "-o", outputPath)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Obfuscation failed: %s", err)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		// You may choose to handle the error differently based on your requirements
		return
	}

	// Update the DSN for PostgreSQL
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbType := os.Getenv("DB_TYPE")
	dsn := dbType + "://" + dbUser + ":" + dbPassword + "@" + dbHost + ":" + dbPort + "/" + dbName + "?sslmode=disable"

	db := app.InitDB(dsn)
	defer db.Close()

	app.SetupDB(db)
	inputPath := "public/components/auth/auth.js"
	outputPath := "public/components/auth/auth_obfuscate.js"
	obfuscateJavaScript(inputPath, outputPath)

	// When adding secure headers on the root of the webserver, all pages going to have the same headers, so no need to add to all

	http.HandleFunc("/global_style/global.css", app.CssHandler)

	http.HandleFunc("/", app.AddSecurityHeaders(app.SendComponent("auth", db)))
	http.HandleFunc("/components/auth/auth.css", app.AuthCssHandler)
	// http.HandleFunc("/scripts/auth_obfuscate.js", app.ErrorsHandler)
	http.HandleFunc("/components/auth/auth_obfuscate.js", app.AuthHandler)
	http.HandleFunc("/components/auth/animation.js", app.AnimationsHandler)
	// Serve public/img folder
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("public/img"))))
	http.HandleFunc("/register", app.RegisterHandler(db))
	http.HandleFunc("/login", app.LoginHandler(db))

	http.HandleFunc("/images/logo.png", app.LogoHandler)
	http.HandleFunc("/checked", app.CheckLogoHandler)
	http.HandleFunc("/logo", app.LogoHandler)

	http.HandleFunc("/create_post", app.CreatePostHandler)
	http.HandleFunc("/scripts/posts.js", app.PostsHandler)

	// http.HandleFunc("/codeQuarry", app.SendTemplate("codeQuarry"))
	http.HandleFunc("/home", app.SendComponent("home", db))
	// http.HandleFunc("/styles/codeQuarry.css", app.CQcssHandler)
	http.HandleFunc("/components/home/home.css", app.CQcssHandler)
	// http.HandleFunc("/styles/header.css", app.HeaderCssHandler)
	http.HandleFunc("/templates/header/header.css", app.HeaderHandlerCss)
	http.HandleFunc("/templates/footer/footer.css", app.FooterHandlerCss)

	http.HandleFunc("/components/profile/profile.css", app.ProfileCSSHandler)
	http.HandleFunc("/logout", app.LogoutHandler(db))

	http.HandleFunc("/ws", app.WebsocketHandler(db))
	http.HandleFunc("/scripts/websocket.js", app.WebsocketFileHandler)

	http.HandleFunc("/votes", app.VoteHandler)

	// ONE SUBJECT
	http.HandleFunc("/subject/", app.SendComponent("subject", db))
	http.HandleFunc("/components/subject/subject.css", app.SubjectCSSHandler)
	http.HandleFunc("/scripts/subject.js", app.SubjectHandlerJS)

	http.HandleFunc("/scripts/all_subjects.js", app.AllSubjectsHandlerJS)
	http.HandleFunc("/templates/all_subjects/all_subjects.css", app.AllSubjectsHandlerCSS)
	app.InsertMultipleSubjects(db)
	http.HandleFunc("/api/subjects", app.SubjectsHandler(db))
	http.HandleFunc("/api/questions", app.QuestionsHandler(db))
	http.HandleFunc("/api/responses", app.ResponsesHandler(db))
	http.HandleFunc("/api/favoris", app.FavoriHandler(db))
	http.HandleFunc("/api/classement", app.SendUsersInfoJson(db))
	http.HandleFunc("/detect_lang", app.DetectLanguageHandler)
	http.HandleFunc("/question_viewer", app.QuestionViewerHandler(db))
	http.HandleFunc("/scripts/question_viewer.js", app.QuestionViewerJSHandler)
	http.HandleFunc("/components/question_viewer/question_viewer.css", app.QuestionViewerCSSHandler)
	http.HandleFunc("/profile", app.SendComponent("profile", db))
	http.HandleFunc("/update-profile", app.UpdateProfileHandler(db))
	http.HandleFunc("/search_bar/input.js", app.SearchBarJS)
	http.HandleFunc("/components/profile/profile.js", app.ProfileJs)
	http.HandleFunc("/posts.css", app.PostCSSHandler)
	//http.HandleFunc("/classement", app.ClassementHandler(db))
	http.HandleFunc("/classement", app.SendComponent("classement",db))
	http.HandleFunc("/classement.css", app.ClassementCSSHandler)
	http.HandleFunc("/scripts/classement.js", app.ClassementJSHandler)

	fmt.Println("Server is running on https://localhost:443/")
	err = http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
	if err != nil {
		app.Log(app.ErrorLevel, "Error starting the server")
		log.Fatal("[DEBUG] ListenAndServe: ", err)
	}
}
