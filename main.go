package main

import (
	"CQ/app"
	"fmt"
	"log"
	"net/http"
	"os/exec"
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

	db := app.InitDB("users.db")
	defer db.Close()

	app.SetupDB(db)
	inputPath := "scripts/errors.js"
	outputPath := "scripts/errors_obfuscate.js"
	obfuscateJavaScript(inputPath, outputPath)
	http.HandleFunc("/", app.AddSecurityHeaders(app.LoginhandlerPage))
	http.HandleFunc("/styles.css", app.CssHandler)
	http.HandleFunc("/scripts/animation.js", app.AnimationsHandler)
	http.HandleFunc("/scripts/errors_obfuscate.js", app.ErrorsHandler)
	http.HandleFunc("/codeQuarry", app.AddSecurityHeaders(app.HandleCodeQuarry))
	http.HandleFunc("/register", app.AddSecurityHeaders(app.RegisterHandler(db)))
	http.HandleFunc("/login", app.AddSecurityHeaders(app.LoginHandler(db)))

	fmt.Println("Server is running on https://localhost:443/")
	err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal("[DEBUG] ListenAndServe: ", err)
	}
}
