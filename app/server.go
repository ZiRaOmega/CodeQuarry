package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"path"
	"text/template"
)

/* ======================= GLOBAL ======================= */

func SendComponent(component_name string, db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(component_name)
		if component_name == "auth" {
			log.Printf("[SendIndex:%s] New Client with IP: %s\n", r.URL.Path, r.RemoteAddr)
			err := ParseAndExecuteTemplate(component_name, nil, w)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}
		cookie, err := r.Cookie("session")
		if err != nil {
			log.Printf("Error getting session cookie: %v", err)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		session_id := cookie.Value
		if isValidSession(session_id, db) {
			// Get user info from user_id
			var user User
			if component_name == "profile" {
				user, err = GetUser(session_id, db)
				if err != nil {
					fmt.Println(err.Error())
					http.Error(w, "Error getting user info", http.StatusInternalServerError)
					return
				}
			}
			err = ParseAndExecuteTemplate(component_name, user, w)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			log.Printf("[SendIndex:%s] New Client with IP: %s\n", r.URL.Path, r.RemoteAddr)
		}
	}
}

// Let's make a list of all the templates we need in each case
var templates = map[string]*template.Template{}

func init() {
	// Pre-parse all templates.
	templates["home"] = parseTemplates("home", "head", "header", "all_subjects", "footer", "script")
	templates["auth"] = parseTemplates("auth", "head", "script")
	templates["subject"] = parseTemplates("subject", "head", "header", "footer", "script")
}

func parseTemplates(component_name string, parts ...string) *template.Template {
	
	// Construct the paths for common template parts.
	templatePath := "public/templates/"
	componentPath := "public/components/"
	var paths []string
	for _, part := range parts {
		paths = append(paths, path.Join(templatePath, part, part+".html"))
	}
	// Add the component template.
	paths = append(paths, path.Join(componentPath, component_name, component_name+".html"))

	// Use template.Must to panic if there's an error.
	return template.Must(template.ParseFiles(paths...))
}

func ParseAndExecuteTemplate(component_name string, data interface{}, w http.ResponseWriter) error {
	// Execute the template with the given data.
	tmpl, ok := templates[component_name]
	if !ok {
		http.Error(w, "Component not found", http.StatusNotFound)
	}
	err := tmpl.ExecuteTemplate(w, component_name, data)
	if err != nil {
		return err
	}
	return nil
}

// CSS
func CssHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the styles.css file when the /styles.css route is accessed
	http.ServeFile(w, r, "public/global_style/global.css")
}

func LogoHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the codeQuarry.html file as the default page.
	http.ServeFile(w, r, "public/images/CODEQUARRY.webp")
}

func AnimationsHandler(w http.ResponseWriter, r *http.Request) {
	// serve the animation js file
	http.ServeFile(w, r, "public/components/auth/animation.js")
}

/* ======================= TEMPLATES ======================= */

/* --------------- HEADER ---------------- */
// CSS
func HeaderHandlerCss(w http.ResponseWriter, r *http.Request) {
	// Serve the styles.css file when the /styles.css route is accessed
	http.ServeFile(w, r, "public/templates/header/header.css")
}

/* --------------- ALL SUBJECTS ---------------- */
// JS
func AllSubjectsHandlerJS(w http.ResponseWriter, r *http.Request) {
	// Serve the subjects.html file as the default page
	http.ServeFile(w, r, "public/templates/all_subjects/all_subjects.js")
}
// CSS
func AllSubjectsHandlerCSS(w http.ResponseWriter, r *http.Request) {
	// Serve the styles.css file when the /styles.css route is accessed
	http.ServeFile(w, r, "public/templates/all_subjects/all_subjects.css")
}

/* ======================= COMPONENTS ======================= */

/* --------------- AUTH ---------------- */
// JS
func AuthHandler(w http.ResponseWriter, r *http.Request) {
	// serve the animation js file
	http.ServeFile(w, r, "public/components/auth/auth_obfuscate.js")
}

// CSS
func AuthCssHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the styles.css file when the /styles.css route is accessed
	http.ServeFile(w, r, "public/components/auth/auth.css")
}

/* --------------- HOME ---------------- */
// HTML
/* func HandleCodeQuarry(w http.ResponseWriter, r *http.Request) {
	// Serve the codeQuarry.html file as the default page
	http.ServeFile(w, r, "public/components/home/home.html")
} */

// CSS
func CQcssHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the styles.css file when the /styles.css route is accessed
	http.ServeFile(w, r, "public/components/home/home.css")
}
/* --------------- Question viouheur ---------------- */

func QuestionViewerCSSHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the styles.css file when the /styles.css route is accessed
	http.ServeFile(w, r, "public/components/question_viewer/question_viewer.css")
}
/* --------------- PROFILE ---------------- */

func ProfileCSSHandler(w http.ResponseWriter, r *http.Request) {
	// Servce profile.css
	http.ServeFile(w, r, "public/components/profile/profile.css")
}
/* --------------- SUBJECT ---------------- */
//JS
func SubjectJSHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the subjects.html file as the default page
	http.ServeFile(w, r, "public/components/subject/subject.js")
}
// CSS
func SubjectCSSHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the styles.css file when the /styles.css route is accessed
	http.ServeFile(w, r, "public/components/subject/subject.css")
}

func SearchBarJS(w http.ResponseWriter, r *http.Request) {
	// Serve the codeQuarry.html file as the default page
	http.ServeFile(w, r, "public/components/home/search_bar/input.js")
}

func VoteHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the codeQuarry.html file as the default page
	http.ServeFile(w, r, "public/components/home/votes.js")
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/images/create_post.webp")
}

func PostsHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the codeQuarry.html file as the default page
	http.ServeFile(w, r, "public/components/home/posts/posts.js")
}

func DetectLanguageHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the codeQuarry.html file as the default page
	http.ServeFile(w, r, "public/components/home/posts/detect_lang/detect_lang.js")
}

func QuestionViewerJSHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the codeQuarry.html file as the default page
	http.ServeFile(w, r, "public/components/question_viewer/question_viewer.js")
}

/* ======================= WEB_SOCKETS ======================= */

func WebsocketFileHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the codeQuarry.html file as the default page
	http.ServeFile(w, r, "scripts/websocket.js")
}

func HandleQuestionViewer(w http.ResponseWriter, r *http.Request) {
	// Serve the codeQuarry.html file as the default page
	http.ServeFile(w, r, "public/components/home/question_viewer/question_viewer.html")
}
