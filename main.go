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
	inputPath := "scripts/auth.js"
	outputPath := "scripts/auth_obfuscate.js"
	obfuscateJavaScript(inputPath, outputPath)

	// When adding secure headers on the root of the webserver, all pages going to have the same headers, so no need to add to all

	http.HandleFunc("/", app.AddSecurityHeaders(app.SendTemplate("login", nil)))
	http.HandleFunc("/styles.css", app.CssHandler)
	http.HandleFunc("/codeQuarry.css", app.CQcssHandler)
	http.HandleFunc("/scripts/animation.js", app.AnimationsHandler)
	http.HandleFunc("/scripts/auth_obfuscate.js", app.ErrorsHandler)
	http.HandleFunc("/scripts/websocket.js", app.WebsocketFileHandler)
	http.HandleFunc("/scripts/subjects.js", app.SubjectsHandlerJS)
	//Serve public/img folder
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("public/img"))))
	http.HandleFunc("/codeQuarry", app.SendTemplate("codeQuarry", nil))
	http.HandleFunc("/register", app.RegisterHandler(db))
	http.HandleFunc("/login", app.LoginHandler(db))
	http.HandleFunc("/logo", app.LogoHandler)
	http.HandleFunc("/logout", app.LogoutHandler(db))
	http.HandleFunc("/ws", app.WebsocketHandler(db))
	http.HandleFunc("/profile", app.ProfileHandler(db))
	app.InsertMultipleSubjects(db)
	http.HandleFunc("/api/subjects", app.SubjectsHandler(db))
	http.HandleFunc("/update-profile", app.UpdateProfileHandler(db))
	fmt.Println("Server is running on https://localhost:443/")
	err = http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
	if err != nil {
		app.Log(app.ErrorLevel, "Error starting the server")
		log.Fatal("[DEBUG] ListenAndServe: ", err)
	}
}
