package app

import (
	"log"
	"net/http"
	"text/template"
)

/* ======================= GLOBAL ======================= */

func SendTemplate(template_name string, data interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[SendIndex:%s] New Client with IP: %s\n", r.URL.Path, r.RemoteAddr)
		err := ParseAndExecuteTemplate(template_name, data, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func ParseAndExecuteTemplate(template_name string, data interface{}, w http.ResponseWriter) error {
	tmpl, err := template.ParseFiles("public/templates/head/head.html", "public/templates/header/header.html", "public/templates/footer/footer.html", "public/templates/script/script.html", "public/components/"+template_name+"/"+template_name+".html")
	if err != nil {
		return err
	}
	err = tmpl.ExecuteTemplate(w, template_name, data)
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
	http.ServeFile(w, r, "scripts/animation.js")
}

/* ======================= TEMPLATES ======================= */

// HEADER CSS
func HeaderCssHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the styles.css file when the /styles.css route is accessed
	http.ServeFile(w, r, "public/templates/header/header.css")
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
func HandleCodeQuarry(w http.ResponseWriter, r *http.Request) {
	// Serve the codeQuarry.html file as the default page
	http.ServeFile(w, r, "public/components/home/home.html")
}

// CSS
func CQcssHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the styles.css file when the /styles.css route is accessed
	http.ServeFile(w, r, "public/components/home/home.css")
}

// JS
func SubjectsHandlerJS(w http.ResponseWriter, r *http.Request) {
	// Serve the subjects.html file as the default page
	http.ServeFile(w, r, "public/components/home/subjects/subjects.js")
}

func VoteHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the codeQuarry.html file as the default page
	http.ServeFile(w, r, "public/components/home/votes.js")
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/create_post.webp")
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
