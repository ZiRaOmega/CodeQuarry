package app

import (
	"log"
	"net/http"
	"path"
	"text/template"
)


func parseTemplates(componentName string, parts ...string) *template.Template {
	// Construct the paths for common template parts.
	templatePath := "public/templates/"
	componentPath := "public/components/"
	var paths []string
	for _, part := range parts {
		paths = append(paths, path.Join(templatePath, part, part+".html"))
	}
	// Add the component template.
	paths = append(paths, path.Join(componentPath, componentName, componentName+".html"))
	
	// Use template.Must to panic if there's an error.
	return template.Must(template.ParseFiles(paths...))
}

var templates = map[string]*template.Template{}
func SendComponent(componentName string) http.HandlerFunc {
	templates["home"] = parseTemplates("home", "head", "header", "footer", "script")
	templates["auth"] = parseTemplates("auth", "head", "footer", "script")
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[SendComponent:%s] New Client with IP: %s\n", componentName, r.RemoteAddr)

		tmpl, ok := templates[componentName]
		if !ok {
			http.Error(w, "Component not found", http.StatusNotFound)
			return
		}

		if err := tmpl.ExecuteTemplate(w, componentName, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

// handle CSS files
func CssHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the styles.css file when the /styles.css route is accessed
	http.ServeFile(w, r, "public/styles/styles.css")
}

func HeaderCssHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the styles.css file when the /styles.css route is accessed
	http.ServeFile(w, r, "public/templates/header/header.css")
}

func CQcssHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the styles.css file when the /styles.css route is accessed
	http.ServeFile(w, r, "public/components/home/home.css")
}

func LogoHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the codeQuarry.html file as the default page.
	http.ServeFile(w, r, "public/images/CODEQUARRY.webp")
}

func AnimationsHandler(w http.ResponseWriter, r *http.Request) {
	// serve the animation js file
	http.ServeFile(w, r, "scripts/animation.js")
}

func ErrorsHandler(w http.ResponseWriter, r *http.Request) {
	// serve the animation js file
	http.ServeFile(w, r, "public/components/auth/auth_obfuscate.js")
}

func SubjectsHandlerJS(w http.ResponseWriter, r *http.Request) {
	// Serve the subjects.html file as the default page
	http.ServeFile(w, r, "scripts/subjects.js")
}

func HandleCodeQuarry(w http.ResponseWriter, r *http.Request) {
	// Serve the codeQuarry.html file as the default page
	http.ServeFile(w, r, "public/components/home/home.html")
}

func WebsocketFileHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the codeQuarry.html file as the default page
	http.ServeFile(w, r, "scripts/websocket.js")
}

func VoteHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the codeQuarry.html file as the default page
	http.ServeFile(w, r, "scripts/votes.js")
}
