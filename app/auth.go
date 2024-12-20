package app

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/mail"
	"net/url"
	"os"
	"time"

	UUID "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"

	"codequarry/app/utils"
)

type PageData struct {
	Content string
}

const Cookie_Expiration = 5 * time.Hour

func RegisterHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
			return
		}
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}

		// Capture reCAPTCHA response
		recaptchaResponse := r.FormValue("g-recaptcha-response")
		if recaptchaResponse == "" {
			json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "reCAPTCHA response missing"})
			return
		}

		// Verify reCAPTCHA
		if !verifyRecaptcha(recaptchaResponse) {
			json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "reCAPTCHA verification failed"})
			return
		}

		lastname := r.FormValue("lastname")
		firstname := r.FormValue("firstname")
		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")
		if !isValidEmail(email) {
			Log(ErrorLevel, "Invalid email")
			// http.Error(w, "Invalid email", http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "Invalid email"})
			return
		}
		if len(lastname) < 1 || len(firstname) < 1 || len(username) < 2 || len(email) < 4 || len(password) < 5 {
			Log(ErrorLevel, "Invalid registration data")
			json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "Invalid registration data"})
			return
		}
		if utils.ContainsSQLi(lastname) || utils.ContainsSQLi(firstname) || utils.ContainsSQLi(username) || utils.ContainsSQLi(email) || utils.ContainsSQLi(password) {
			Log(ErrorLevel, "SQL injection detected")
			// http.Error(w, "SQL injection detected", http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "SQL injection detected"})
			return

		} else if utils.ContainsXSS(lastname) || utils.ContainsXSS(firstname) || utils.ContainsXSS(username) || utils.ContainsXSS(email) || utils.ContainsXSS(password) {
			Log(ErrorLevel, "XSS detected")
			// http.Error(w, "XSS detected", http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "XSS detected"})
			return
		}
		// !!! TODO : smth better than bcrypt?
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			Log(ErrorLevel, "Error hashing password")
			// http.Error(w, "Error hashing password", http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "Error hashing password"})
			return
		}
		today := time.Now()
		// In postgres, the placeholders are $1, $2, $3, etc. In MySQL, the placeholders are ?, ?, ?, etc.
		stmt, err := db.Prepare("INSERT INTO users(lastname, firstname, username, email, password, avatar ,xp,rang_rank_,creation_date) VALUES($1, $2, $3, $4, $5, $6, 0,0, $7)")
		if err != nil {

			Log(ErrorLevel, "Error preparing the SQL statement")
			// http.Error(w, "Error preparing the SQL statement", http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "Error preparing the SQL statement"})
			return
		}
		defer stmt.Close()
		if _, err := stmt.Exec(lastname, firstname, username, email, string(hashedPassword), "/img/defaultUser.png",today); err != nil {
			if err.Error() == "pq: duplicate key value violates unique constraint \"users_username_key\"" {
				Log(ErrorLevel, "Username already exists")
				// http.Error(w, "Username already exists", http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "Username already exists"})
				return
			} else if err.Error() == "pq: duplicate key value violates unique constraint \"users_email_key\"" {
				Log(ErrorLevel, "Email already exists")
				// http.Error(w, "Email already exists", http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "Email already exists"})
				return
			}
			Log(ErrorLevel, "Error inserting the data into the database"+err.Error())
			// http.Error(w, "Error inserting the data into the database", http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "Error inserting the data into the database"})
			return
		}
		token := GenerateTokenVerificationEmail()
		SendVerificationEmail(db, email, token)
		err = InsertVerifToken(db, email, token)
		if err != nil {
			Log(ErrorLevel, err.Error())
			return
		}
		err = CreateSession(username, db, w)
		if err != nil {

			Log(ErrorLevel, "Error creating session")
			json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "Error creating session"})
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "Registration successful"})
		Log(DebugLevel, "Registration successful: "+username+" at "+r.URL.Path)
	}
}

// Verify reCAPTCHA response with Google
func verifyRecaptcha(token string) bool {
	secret := os.Getenv("RECAPTCHA_SERVER_KEY")
	endpoint := "https://www.google.com/recaptcha/api/siteverify"

	resp, err := http.PostForm(endpoint, url.Values{
		"secret":   {secret},
		"response": {token},
	})
	if err != nil {
		log.Println("Error verifying reCAPTCHA:", err)
		return false
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Println("Error parsing reCAPTCHA response:", err)
		return false
	}

	return result["success"].(bool)
}
func GetEmailWithUsername(db *sql.DB, username string) string {
	//fmt.Println(username)
	stmt, err := db.Prepare("SELECT email FROM users WHERE username = $1")
	if err != nil {
		Log(ErrorLevel, err.Error())
	}
	defer stmt.Close()
	var email string
	err = stmt.QueryRow(username).Scan(&email)
	if err != nil {
		Log(ErrorLevel, err.Error())
	}
	return email
}
func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
func LoginHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
			return
		}
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}
		// Capture reCAPTCHA response
		recaptchaResponse := r.FormValue("g-recaptcha-response")
		if recaptchaResponse == "" {
			json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "reCAPTCHA response missing"})
			return
		}

		// Verify reCAPTCHA
		if !verifyRecaptcha(recaptchaResponse) {
			json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "reCAPTCHA verification failed"})
			return
		}

		username := r.FormValue("usernameOrEmailLogin")
		password := r.FormValue("passwordLogin")
		if utils.ContainsSQLi(username) || utils.ContainsSQLi(password) {
			Log(ErrorLevel, "SQL injection detected")
			json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "SQL injection detected"})
			return
		} else if utils.ContainsXSS(username) || utils.ContainsXSS(password) {
			Log(ErrorLevel, "XSS detected")
			json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "XSS detected"})
			return
		}
		email := GetEmailWithUsername(db, username)
		if !isEmailVerified(db, email) {
			Log(ErrorLevel, "Email not verified"+email)
			json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "User Not registered or has not Verify email"})
			return
		}
		// Fetch user from database
		var storedPassword string
		var deletingDate sql.NullTime
		query := "SELECT password, deleting_date FROM users WHERE username = $1 OR email = $2"
		err := db.QueryRow(query, username, username).Scan(&storedPassword, &deletingDate)
		if err != nil {
			Log(ErrorLevel, err.Error())
			json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "Error while fetching in db"})
			return
		}

		if deletingDate.Valid {
			Log(ErrorLevel, "Account is in deleting state")
			json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "User account is in deleting state contact admin to retrieve this account"})
			return
		}
		if err != nil {
			if err == sql.ErrNoRows {
				Log(ErrorLevel, "No user found with the provided credentials"+username+" at "+r.URL.Path)

				json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "No user found with the provided credentials"})
				return
			} else {
				Log(ErrorLevel, "Error fetching user from database"+username+" at "+r.URL.Path)

				json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "Error fetching user from database"})
				return
			}
		}
		// Compare the stored hashed password with the password that was submitted
		//Function to check the password hash
		err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
		if err != nil {
			Log(DebugLevel, "Invalid login credentials"+username+" at "+r.URL.Path)

			json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "Invalid login credentials"})
			return
		}
		err = CreateSession(username, db, w)
		if err != nil {
			Log(ErrorLevel, "Error creating session")
			json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "Error creating session"})
			return
		}
		// Login successful
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "Login successful"})
		Log(DebugLevel, "Login successful: "+username+" at "+r.URL.Path)

	}
}

// LogoutHandler
func LogoutHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		if err != nil {
			Log(ErrorLevel, "Error getting session cookie")
			http.Error(w, "Error getting session cookie", http.StatusInternalServerError)
			return
		}
		session_id := cookie.Value
		err = DeleteSession(session_id, db)
		if err != nil {
			Log(ErrorLevel, "Error deleting session")
			http.Error(w, "Error deleting session", http.StatusInternalServerError)
			return
		}
		// Remove cookie
		cookie.Expires = time.Now().AddDate(0, 0, -1)
		http.SetCookie(w, cookie)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

// CreateSession creates a new session for the given username.
// It retrieves the user ID from the database, generates a UUID for the session,
// inserts the session into the database, and sets a session cookie in the response.
// The session cookie is set to expire after a certain duration.
// Parameters:
//   - username: The username for which to create the session.
//   - db: The database connection.
//   - w: The HTTP response writer.
//
// Returns:
//   - error: An error if any occurred during the session creation process.
func CreateSession(username string, db *sql.DB, w http.ResponseWriter) error {
	user_id, err := GetUserIDFromDB(username, db)
	if err != nil {
		return err
	}
	user_uuid := UUID.NewV4().String()
	createdAt := time.Now()
	expireAt := createdAt.Add(Cookie_Expiration)
	err = InsertSessionToDB(db, user_id, user_uuid, createdAt, expireAt)
	if err != nil {
		return err
	}
	cookie := http.Cookie{
		Name:     "session",
		Value:    user_uuid,
		Expires:  expireAt,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Secure:   true,
	}
	http.SetCookie(w, &cookie)
	return err
}
