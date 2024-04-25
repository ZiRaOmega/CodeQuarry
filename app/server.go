package app

import (
	"log"
	"net/http"
	"text/template"
)

func SendTemplate(template_name string, data interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[SendIndex:%s] New Client with IP: %s\n", r.URL.Path, r.RemoteAddr)
		// This function is used to handle requests to send the index page. It logs the IP address of the client making the request.
		tmpl, err := template.ParseFiles("public/header.html", "public/footer.html", "public/script.html", "public/head.html", "public/"+template_name+".html")
		if err != nil {
			panic(err)
			// Parsing the HTML templates for the header, footer, and index. If there is an error, the program will panic and stop execution.
		}
		err = tmpl.ExecuteTemplate(w, template_name, data)
		if err != nil {
			panic(err)
			// Executing the "index" template and sending it to the client. If there is an error, the program will panic and stop execution.
		}
	}
}

// handle CSS files
func CssHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the styles.css file when the /styles.css route is accessed
	http.ServeFile(w, r, "public/styles.css")
}

func CQcssHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the styles.css file when the /styles.css route is accessed
	http.ServeFile(w, r, "public/codeQuarry.css")
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
	http.ServeFile(w, r, "scripts/auth_obfuscate.js")
}

func SubjectsHandlerJS(w http.ResponseWriter, r *http.Request) {
	// Serve the subjects.html file as the default page
	http.ServeFile(w, r, "scripts/subjects.js")
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
