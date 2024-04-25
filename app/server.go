package app

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"
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

func ProfileHandler(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		//Get cookie
		cookie, err := r.Cookie("session")
		if err != nil {
			http.Error(w, "Error getting session cookie", http.StatusInternalServerError)
			return
		}
		session_id := cookie.Value
		//Get user info from user_id
		user, err := GetUser(session_id, db)
		if err != nil {
			fmt.Println(err.Error())
			http.Error(w, "Error getting user info", http.StatusInternalServerError)
			return
		}
		//Send user info to the profile page
		log.Printf("[SendIndex:%s] New Client with IP: %s\n", r.URL.Path, r.RemoteAddr)
		// This function is used to handle requests to send the index page. It logs the IP address of the client making the request.
		tmpl, err := template.ParseFiles("public/header.html", "public/footer.html", "public/script.html", "public/head.html", "public/"+"profile"+".html")
		if err != nil {
			panic(err)
			// Parsing the HTML templates for the header, footer, and index. If there is an error, the program will panic and stop execution.
		}
		err = tmpl.ExecuteTemplate(w, "profile", user)
		if err != nil {
			panic(err)
			// Executing the "index" template and sending it to the client. If there is an error, the program will panic and stop execution.
		}
	}
}

func (U *User) FormatBirthDate() string {
	return U.BirthDate.Time.Format("2006-02-01")
}

func (U *User) FormatSchoolYear() string {
	return U.SchoolYear.Time.Format("2006-02-01")
}

func (U *User) FormatXP() int64 {
	return U.XP.Int64
}

// Define User structure based on your database schema
type User struct {
	ID        int
	LastName  string
	FirstName string
	Username  string
	Email     string
	Password  string
	Avatar    sql.NullString

	BirthDate         sql.NullTime // Adjusted for possible NULL values
	Birth_Date_Format string
	Bio               sql.NullString

	Website sql.NullString

	GitHub sql.NullString

	XP sql.NullInt64

	Rank sql.NullString

	SchoolYear         sql.NullTime // Adjusted for possible NULL values
	School_Year_Format string
	CreationDate       sql.NullTime // Adjusted for possible NULL values
	UpdateDate         sql.NullTime // Adjusted for possible NULL values
	DeletingDate       sql.NullTime // Adjusted for possible NULL values
}

// GetUser fetches user details from the database based on the session ID
func GetUser(session_id string, db *sql.DB) (User, error) {
	var user User
	if db == nil {
		return user, errors.New("database connection is nil")
	}

	// Prepare the SQL statement for fetching user_id from session_id
	stmt, err := db.Prepare("SELECT user_id FROM Sessions WHERE uuid = $1")
	if err != nil {
		return user, err // Return immediately if there's an error preparing the statement
	}
	defer stmt.Close() // Ensure statement is always closed after use

	// Execute the statement and scan the result
	var user_id int
	if err = stmt.QueryRow(session_id).Scan(&user_id); err != nil {
		return user, err
	}

	// Prepare the SQL statement for fetching user details from user_id
	stmt, err = db.Prepare("SELECT id_student, lastname, firstname, username, email, password, avatar, birth_date, bio, website, github, xp, rang_rank_, school_year, creation_date, update_date, deleting_date FROM users WHERE id_student = $1")
	if err != nil {
		return user, err // Handle preparation errors
	}
	defer stmt.Close() // Ensure statement is always closed after use

	// Execute the statement and scan the result into the User struct
	if err = stmt.QueryRow(user_id).Scan(&user.ID, &user.LastName, &user.FirstName, &user.Username, &user.Email, &user.Password, &user.Avatar, &user.BirthDate, &user.Bio, &user.Website, &user.GitHub, &user.XP, &user.Rank, &user.SchoolYear, &user.CreationDate, &user.UpdateDate, &user.DeletingDate); err != nil {
		return user, err
	}
	// FormatDates formats CreationDate and
	// UpdateDate to a human-readable format dd/mm/yyyy
	user.Birth_Date_Format = user.FormatBirthDate()
	user.School_Year_Format = user.FormatSchoolYear()

	return user, nil // Return the populated user object
}

func UpdateProfileHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		//Check if the session is valid
		cookie, err := r.Cookie("session")
		if err != nil {
			http.Error(w, "Error getting session cookie", http.StatusInternalServerError)
			return
		}
		session_id := cookie.Value
		if !isValidSession(session_id, db) {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		//Get form
		user := User{}

		user.ID, err = getUserIDUsingSessionID(session_id, db)

		if err != nil {
			fmt.Println(err.Error())
		}
		if strconv.Itoa(user.ID) != r.PostFormValue("id_student") {
			http.Error(w, "Invalid user", http.StatusForbidden)
			return
		}
		user.LastName = r.PostFormValue("lastname")
		user.FirstName = r.PostFormValue("firstname")
		user.Username = r.PostFormValue("username")
		user.Email = r.PostFormValue("email")
		user.Password = r.PostFormValue("password")

		birthDateStr := r.PostFormValue("birth_date")
		birthDate, err := time.Parse("2006-01-02", birthDateStr)
		if err != nil {
			fmt.Println(err.Error())
		}
		user.BirthDate = sql.NullTime{Time: birthDate, Valid: true}
		user.Bio = sql.NullString{String: r.PostFormValue("bio"), Valid: true}
		user.Website = sql.NullString{String: r.PostFormValue("website"), Valid: true}
		user.GitHub = sql.NullString{String: r.PostFormValue("github"), Valid: true}
		user.SchoolYear = sql.NullTime{Time: time.Now(), Valid: true}
		//if Password is empty, don't update it
		if user.Password == "" {
			//Prepare
			stmt, err := db.Prepare("UPDATE users SET lastname = $1, firstname = $2, username = $3, email = $4, birth_date = $5, bio = $6, website = $7, github = $8, school_year = $9 WHERE id_student = $10")
			if err != nil {
				fmt.Println(err.Error())
			}
			defer stmt.Close()
			//Execute
			_, err = stmt.Exec(user.LastName, user.FirstName, user.Username, user.Email, user.BirthDate, user.Bio, user.Website, user.GitHub, user.SchoolYear, user.ID)
			if err != nil {
				fmt.Println(err.Error())
			}
		} else {
			//Prepare
			stmt, err := db.Prepare("UPDATE users SET lastname = $1, firstname = $2, username = $3, email = $4, password = $5, birth_date = $6, bio = $7, website = $8, github = $9, school_year = $10 WHERE id_student = $11")
			if err != nil {
				fmt.Println(err.Error())
			}
			defer stmt.Close()
			//Execute
			_, err = stmt.Exec(user.LastName, user.FirstName, user.Username, user.Email, user.Password, user.BirthDate, user.Bio, user.Website, user.GitHub, user.SchoolYear, user.ID)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
		http.Redirect(w, r, "/profile", http.StatusSeeOther)

	}
}
