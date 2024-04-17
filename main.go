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
	cmd := exec.Command("python3", "obfuscator/obfuscate.py", inputPath, outputPath)
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
	http.HandleFunc("/", app.LoginhandlerPage)
	http.HandleFunc("/styles.css", app.CssHandler)
	http.HandleFunc("/scripts/animation.js", app.AnimationsHandler)
	http.HandleFunc("/scripts/errors_obfuscate.js", app.ErrorsHandler)

	http.HandleFunc("/register", app.RegisterHandler(db))
	http.HandleFunc("/login", app.LoginHandler(db))

	fmt.Println("Server is running on http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
