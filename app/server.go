package app

import (
	"net/http"
)

// handle CSS files
func CssHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the styles.css file when the /styles.css route is accessed
	http.ServeFile(w, r, "public/styles.css")
}

func LogoHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the codeQuarry.html file as the default page.
	http.ServeFile(w, r, "public/CODEQUARRY.webp")
}

func AnimationsHandler(w http.ResponseWriter, r *http.Request) {
	//serve the animation js file
	http.ServeFile(w, r, "scripts/animation.js")
}

func ErrorsHandler(w http.ResponseWriter, r *http.Request) {
	//serve the animation js file
	http.ServeFile(w, r, "scripts/errors_obfuscate.js")
}

// handle the main page
func LoginhandlerPage(w http.ResponseWriter, r *http.Request) {
	// Serve the login.html file as the default page
	http.ServeFile(w, r, "public/login.html")
}

func HandleCodeQuarry(w http.ResponseWriter, r *http.Request) {
	// Serve the codeQuarry.html file as the default page
	http.ServeFile(w, r, "public/codeQuarry.html")
}

func WebsocketFileHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the codeQuarry.html file as the default page
	http.ServeFile(w, r, "scripts/websocket.js")
}
