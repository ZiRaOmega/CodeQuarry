package main

import (
	"CQ/app"
	"fmt"
	"net/http"
)

func main() {
	db := app.InitDB("users.db")
	defer db.Close()

	app.SetupDB(db)

	http.HandleFunc("/", app.LoginhandlerPage)
	http.HandleFunc("/styles.css", app.CssHandler)
	http.HandleFunc("/scripts/animation.js", app.AnimationsHandler)
	http.HandleFunc("/scripts/errors.js", app.ErrorsHandler)
	http.HandleFunc("/register", app.RegisterHandler(db))
	http.HandleFunc("/login", app.LoginHandler(db))

	fmt.Println("Server is running on http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
