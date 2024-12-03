package app

import (
	"database/sql"
	"net/http"
	"path"

	"html/template" // Add this import statement

	"github.com/gorilla/csrf"
)

type AuthInfo struct {
	Rank int
}

/* ======================= GLOBAL ======================= */
//Couche de Présentation (Gestion des Requêtes HTTP)
func SendComponent(component_name string, db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		csrfToken := csrf.TemplateField(r)
		if component_name == "auth" || component_name == "cgu" || component_name == "rgpd" {
			handleStaticComponent(component_name, csrfToken, w)
			return
		}
		handleDynamicComponent(component_name, csrfToken, db, w, r)
	}
}

// Couche de Présentation (Gestion des Requêtes HTTP)
func handleStaticComponent(component_name string, csrfToken template.HTML, w http.ResponseWriter) {
	tmplData := struct {
		CSRFToken template.HTML
	}{
		CSRFToken: csrfToken,
	}
	err := ParseAndExecuteTemplate(component_name, tmplData, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Couche de Présentation (Gestion des Requêtes HTTP)
func handleDynamicComponent(component_name string, csrfToken template.HTML, db *sql.DB, w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	session_id := cookie.Value
	if isValidSession(session_id, db) {
		user, err := getUserInfo(session_id, db)
		if err != nil {
			http.Error(w, "Error getting user info", http.StatusInternalServerError)
			return
		}
		user.CSRFToken = csrfToken

		if component_name == "profile" || component_name == "classement" {
			user.Rank.String, err = SetRankByXp(user)
			if err != nil {
				http.Error(w, "Error getting user rank", http.StatusInternalServerError)
				return
			}
		}
		err = ParseAndExecuteTemplate(component_name, user, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// Couche de Logique Métier (Services)
func getUserInfo(session_id string, db *sql.DB) (User, error) {
	var user User
	//Couche d'Accès aux Données (Interactions avec la Base de Données)
	user, err := GetUser(session_id, db)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// Let's make a list of all the templates we need in each case
var templates = map[string]*template.Template{}

func init() {
	// Pre-parse all templates.
	templates["home"] = parseTemplates("home", "head", "header", "search_bar", "all_subjects", "footer", "script")
	templates["auth"] = parseTemplates("auth", "head", "script")
	templates["subject"] = parseTemplates("subject", "head", "header", "search_bar", "footer", "script", "all_subjects")
	templates["profile"] = parseTemplates("profile", "head", "header", "search_bar", "footer", "script")
	templates["question_viewer"] = parseTemplates("question_viewer", "head", "header", "search_bar", "footer", "script")
	templates["classement"] = parseTemplates("classement", "head", "header", "search_bar", "footer", "script")
	templates["panel"] = parseTemplates("panel", "script", "head", "header", "search_bar")
	templates["cgu"] = parseTemplates("cgu", "head", "header", "search_bar", "footer", "script")
	templates["rgpd"] = parseTemplates("rgpd", "head", "footer")
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

func CheckLogoHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the codeQuarry.html file as the default page.
	http.ServeFile(w, r, "public/images/checked.png")
}

func AnimationsHandler(w http.ResponseWriter, r *http.Request) {
	// serve the animation js file
	http.ServeFile(w, r, "public/components/auth/animation_obfuscate.js")
}

func PanelCssHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/components/panel/panel.css")
}

func FaviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/images/CODEQUARRY.ico")
}

/* ======================= TEMPLATES ======================= */

/* --------------- HEADER ---------------- */
// CSS
func HeaderHandlerCss(w http.ResponseWriter, r *http.Request) {
	// Serve the styles.css file when the /styles.css route is accessed
	http.ServeFile(w, r, "public/templates/header/header.css")
}

/* --------------- FOOTER ---------------- */
func FooterHandlerCss(w http.ResponseWriter, r *http.Request) {
	// Serve the styles.css file when the /styles.css route is accessed
	http.ServeFile(w, r, "public/templates/footer/footer.css")
}

/* --------------- ALL SUBJECTS ---------------- */
// JS
func AllSubjectsHandlerJS(w http.ResponseWriter, r *http.Request) {
	// Serve the subjects.html file as the default page
	http.ServeFile(w, r, "public/templates/all_subjects/all_subjects_obfuscate.js")
}

// CSS
func AllSubjectsHandlerCSS(w http.ResponseWriter, r *http.Request) {
	// Serve the styles.css file when the /styles.css route is accessed
	http.ServeFile(w, r, "public/templates/all_subjects/all_subjects.css")
}

/* --------------- SEARCH_BAR ---------------- */
// JS
func SearchBarJS(w http.ResponseWriter, r *http.Request) {
	// Serve the codeQuarry.html file as the default page
	http.ServeFile(w, r, "public/templates/search_bar/input_obfuscate.js")
}

// CSS
func SearchBarCSS(w http.ResponseWriter, r *http.Request) {
	// Serve the styles.css file when the /styles.css route is accessed
	http.ServeFile(w, r, "public/templates/search_bar/search_bar.css")
}

/* ======================= COMPONENTS ======================= */

/* --------------- AUTH ---------------- */
// JS
func AuthHandler(w http.ResponseWriter, r *http.Request) {
	// serve the auth js file
	http.ServeFile(w, r, "public/components/auth/auth.js")
}

// CSS
func AuthCssHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the styles.css file when the /styles.css route is accessed
	http.ServeFile(w, r, "public/components/auth/auth.css")
}

/* --------------- HOME ---------------- */
// HTML

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

/*Panel*/
func PanelJSHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/components/panel/panel_obfuscate.js")
}

/* --------------- PROFILE ---------------- */

func ProfileCSSHandler(w http.ResponseWriter, r *http.Request) {
	// Servce profile.css
	http.ServeFile(w, r, "public/components/profile/profile.css")
}

func PostCSSHandler(w http.ResponseWriter, r *http.Request) {
	// Servce Post.css
	http.ServeFile(w, r, "public/components/home/posts/posts.css")
}

func ClassementCSSHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/components/classement/classement.css")
}

// JS
func SubjectHandlerJS(w http.ResponseWriter, r *http.Request) {
	// Serve the subjects.html file as the default page
	http.ServeFile(w, r, "public/components/subject/subject_obfuscate.js")
}

// CSS
func SubjectCSSHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the styles.css file when the /styles.css route is accessed
	http.ServeFile(w, r, "public/components/subject/subject.css")
}

func ProfileJs(w http.ResponseWriter, r *http.Request) {
	// Serve the codeQuarry.html file as the default page
	http.ServeFile(w, r, "public/components/profile/profile_obfuscate.js")
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
	http.ServeFile(w, r, "public/components/home/posts/posts_obfuscate.js")
}

func DetectLanguageHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the codeQuarry.html file as the default page
	http.ServeFile(w, r, "public/components/home/posts/detect_lang/detect_lang.js")
}

func QuestionViewerJSHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the codeQuarry.html file as the default page
	http.ServeFile(w, r, "public/components/question_viewer/question_viewer_obfuscate.js")
}

func ClassementJSHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/components/classement/classement_obfuscate.js")
}

/* ======================= WEB_SOCKETS ======================= */

func WebsocketFileHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the codeQuarry.html file as the default page
	http.ServeFile(w, r, "scripts/websocket_obfuscate.js")
}

func HandleQuestionViewer(w http.ResponseWriter, r *http.Request) {
	// Serve the codeQuarry.html file as the default page
	http.ServeFile(w, r, "public/components/home/question_viewer/question_viewer.html")
}

func HandleHeaderJS(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/templates/header/header_obfuscate.js")
}

func HandleCSSRGPD(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "public/components/rgpd/rgpd.css")
}

func HandleRobots(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/robots.txt")
}
func HandleSitemap(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/sitemap.xml")
}
func VerifEmailHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			token := r.URL.Query().Get("token")
			if token == "" {
				http.Error(w, "Token not found", http.StatusBadRequest)
				return
			}
			// Check if token is valid
			if !isValidToken(db, token) {
				http.Error(w, "Token not valid", http.StatusBadRequest)
				return
			} else {
				//Verify Email succes page

				http.Redirect(w, r, "/home", http.StatusSeeOther)
			}
		}
	}
}
