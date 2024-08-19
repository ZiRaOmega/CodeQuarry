package main

import (
	"codequarry/app"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

func obfuscateJavaScript(inputPath, outputPath string, wg *sync.WaitGroup, errChan chan<- error, index int, total int) {
	defer wg.Done()

	// Ensure the path to the Python executable and the script is correct
	cmd := exec.Command("npx", "javascript-obfuscator", inputPath, "-o", outputPath)
	err := cmd.Run()
	if err != nil {
		errChan <- err
		return
	}
	log.Printf("Obfuscation successful %d/%d: %s -> %s", index, total, inputPath, outputPath)
}
func main() {
	err := godotenv.Load()
	if err != nil {
		app.Log(app.ErrorLevel, "Failed to load .env file")
		return
	}
	fmt.Println("Starting server...")
	// DSN for PostgreSQL
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbType := os.Getenv("DB_TYPE")
	URL := os.Getenv("URL")
	dsn := dbType + "://" + dbUser + ":" + dbPassword + "@" + dbHost + ":" + dbPort + "/" + dbName + "?sslmode=require"
	db := app.InitDB(dsn)
	CSRF := app.InitStoreCSRFToken()
	defer db.Close()
	app.SetupDB(db)
	// Obfuscate JavaScript files
	jsFiles := []struct {
		input  string
		output string
	}{
		{"public/components/auth/auth.js", "public/components/auth/auth_obfuscate.js"},
		{"public/components/auth/animation.js", "public/components/auth/animation_obfuscate.js"},
		{"public/components/classement/classement.js", "public/components/classement/classement_obfuscate.js"},
		{"public/components/home/posts/posts.js", "public/components/home/posts/posts_obfuscate.js"},
		{"public/components/panel/panel.js", "public/components/panel/panel_obfuscate.js"},
		{"public/components/profile/profile.js", "public/components/profile/profile_obfuscate.js"},
		{"public/components/question_viewer/question_viewer.js", "public/components/question_viewer/question_viewer_obfuscate.js"},
		{"public/components/subject/subject.js", "public/components/subject/subject_obfuscate.js"},
		{"public/templates/all_subjects/all_subjects.js", "public/templates/all_subjects/all_subjects_obfuscate.js"},
		{"public/templates/header/header.js", "public/templates/header/header_obfuscate.js"},
		{"public/templates/search_bar/input.js", "public/templates/search_bar/input_obfuscate.js"},
		{"scripts/websocket.js", "scripts/websocket_obfuscate.js"},
	}

	var wg sync.WaitGroup
	errChan := make(chan error, len(jsFiles))

	for index, file := range jsFiles {
		wg.Add(1)
		go obfuscateJavaScript(file.input, file.output, &wg, errChan, index, len(jsFiles))
	}

	// Wait for all goroutines to finish
	wg.Wait()
	close(errChan)

	// Check for errors
	for err := range errChan {
		if err != nil {
			log.Fatalf("Obfuscation failed: %s", err)
		}
	}

	RegisterRateLimiter := app.NewRateLimiter(5, time.Hour.Abs())
	GlobalrateLimiter := app.NewRateLimiter(10, time.Minute)
	// When adding secure headers on the root of the webserver, all pages going to have the same headers, so no need to add to all
	http.HandleFunc("/global_style/global.css", app.CssHandler)
	http.HandleFunc("/favicon.ico", app.FaviconHandler)
	http.Handle("/", app.CorsMiddleware(app.AddSecurityHeaders(CSRF(http.HandlerFunc(app.SendComponent("auth", db))).ServeHTTP)))

	http.HandleFunc("/components/auth/auth.css", app.AuthCssHandler)
	http.HandleFunc("/components/auth/auth_obfuscate.js", app.AuthHandler)
	http.HandleFunc("/components/auth/animation.js", app.AnimationsHandler)
	// Serve public/img folder
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("public/img"))))
	http.Handle("/register", CSRF(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		RegisterRateLimiter.Handle(app.RegisterHandler(db)).ServeHTTP(w, r)
	})))
	http.Handle("/login", CSRF(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		GlobalrateLimiter.Handle(app.LoginHandler(db)).ServeHTTP(w, r)
	})))
	http.HandleFunc("/images/logo.png", app.LogoHandler)
	http.HandleFunc("/checked", app.CheckLogoHandler)
	http.HandleFunc("/logo", app.LogoHandler)
	http.HandleFunc("/create_post", app.AddSecurityHeaders(app.CreatePostHandler))
	http.HandleFunc("/scripts/posts.js", app.PostsHandler)
	http.HandleFunc("/home", app.SendComponent("home", db))
	http.HandleFunc("/components/home/home.css", app.CQcssHandler)
	http.HandleFunc("/templates/header/header.css", app.HeaderHandlerCss)
	http.HandleFunc("/templates/footer/footer.css", app.FooterHandlerCss)
	http.HandleFunc("/templates/header/header.js", app.HandleHeaderJS)
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
	http.Handle("/api/responses", CSRF(app.ResponsesHandler(db)))
	http.HandleFunc("/api/favoris", app.FavoriHandler(db))
	http.HandleFunc("/api/classement", app.SendUsersInfoJson(db))
	http.HandleFunc("/detect_lang", app.DetectLanguageHandler)
	http.HandleFunc("/question_viewer", func(w http.ResponseWriter, r *http.Request) {
		CSRF(app.QuestionViewerHandler(db)).ServeHTTP(w, r)
	})
	http.HandleFunc("/scripts/question_viewer.js", app.QuestionViewerJSHandler)
	http.HandleFunc("/components/question_viewer/question_viewer.css", app.QuestionViewerCSSHandler)
	http.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		CSRF(app.SendComponent("profile", db)).ServeHTTP(w, r)
	})
	http.HandleFunc("/update-profile", func(w http.ResponseWriter, r *http.Request) {
		CSRF(http.HandlerFunc(app.UpdateProfileHandler(db))).ServeHTTP(w, r)
	})
	http.HandleFunc("/search_bar/input.js", app.SearchBarJS)
	http.HandleFunc("/templates/search_bar/search_bar.css", app.SearchBarCSS)
	http.HandleFunc("/components/profile/profile.js", app.ProfileJs)
	http.HandleFunc("/posts.css", app.PostCSSHandler)
	http.HandleFunc("/classement", app.SendComponent("classement", db))
	http.HandleFunc("/classement.css", app.ClassementCSSHandler)
	http.HandleFunc("/scripts/classement.js", app.ClassementJSHandler)
	http.HandleFunc("/panel", app.PanelAdminHandler(db))
	http.HandleFunc("/scripts/panel.js", app.PanelJSHandler)
	http.HandleFunc("/components/panel/panel.css", app.PanelCssHandler)
	http.HandleFunc("/verify", app.VerifEmailHandler(db))
	http.HandleFunc("/forgot-password", app.ForgotPasswordHandler(db))
	http.HandleFunc("/cgu", app.SendComponent("cgu", db))
	http.HandleFunc("/rgpd", app.SendComponent("rgpd", db))
	http.HandleFunc("/components/rgpd/rgpd.css", app.HandleCSSRGPD)
	http.Handle("/delete-profile", CSRF(http.HandlerFunc(app.DeleteProfileHandler(db))))
	http.HandleFunc("/robots.txt", app.HandleRobots)
	http.HandleFunc("/sitemap.xml", app.HandleSitemap)

	//go startHTTPServer()
	fmt.Println("Server is running on https://" + URL + ":443/")
	err = http.ListenAndServeTLS(":443", "./cert/fullchain1.pem", "./cert/privkey1.pem", nil)
	//For the server
	//err = http.ListenAndServeTLS(":443", "./cert/codequarry.dev/fullchain1.pem", "./cert/codequarry.dev/privkey1.pem", nil)
	if err != nil {
		app.Log(app.ErrorLevel, "Error starting the server")
		log.Fatal("[DEBUG] ListenAndServe: ", err)
	}

}

// Redirects HTTP requests to HTTPS
func redirectHTTPToHTTPS(w http.ResponseWriter, r *http.Request) {
	// Note: Use http.StatusMovedPermanently for permanent redirects
	http.Redirect(w, r, "https://"+r.Host+r.URL.String(), http.StatusTemporaryRedirect)
}
func startHTTPServer() {
	// Create a new ServeMux
	mux := http.NewServeMux()
	// Register the redirect function specifically
	mux.HandleFunc("/", redirectHTTPToHTTPS)
	// Listen on HTTP port 80 with the new ServeMux
	log.Fatal(http.ListenAndServe(":80", mux))
}
