package main

import (
	"fmt"
	"net/http"
)

// handle CSS files
func cssHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the styles.css file when the /styles.css route is accessed
	http.ServeFile(w, r, "styles.css")
}

func animationsHandler(w http.ResponseWriter, r *http.Request) {
	//serve the animation js file
	http.ServeFile(w, r, "animation.js")
}

// handle the main page
func loginhandler(w http.ResponseWriter, r *http.Request) {
	// Serve the login.html file as the default page
	http.ServeFile(w, r, "login.html")
}

func main() {
	// Set the handler function for the default route
	http.HandleFunc("/", loginhandler)
	// Set the handler function for the /styles.css route
	http.HandleFunc("/styles.css", cssHandler)
	// Set the handler function for the /animation.js route
	http.HandleFunc("/animation.js", animationsHandler)

	// Start the server on port 8080
	fmt.Println("Server is running on http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
