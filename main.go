package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Initialize the database
	db := initDB("users.db")
	defer db.Close()

	// Setup database tables
	setupDB(db)

	// Set the HTTP handlers for serving files and handling requests
	http.HandleFunc("/", loginhandler)
	http.HandleFunc("/styles.css", cssHandler)
	http.HandleFunc("/animation.js", animationsHandler)
	http.HandleFunc("/register", registerHandler(db))

	// Start the server
	fmt.Println("Server is running on http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
