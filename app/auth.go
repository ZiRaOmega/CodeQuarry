package app

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	UUID "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
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

		lastname := r.FormValue("lastname")
		firstname := r.FormValue("firstname")
		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")

		// !!! TODO : smth better than bcrypt?
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			Log(ErrorLevel, "Error hashing password")
			//http.Error(w, "Error hashing password", http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "Error hashing password"})
			return
		}

		// In postgres, the placeholders are $1, $2, $3, etc. In MySQL, the placeholders are ?, ?, ?, etc.
		stmt, err := db.Prepare("INSERT INTO users(lastname, firstname, username, email, password) VALUES($1, $2, $3, $4, $5)")
		if err != nil {
			fmt.Println(err)
			Log(ErrorLevel, "Error preparing the SQL statement")
			//http.Error(w, "Error preparing the SQL statement", http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "Error preparing the SQL statement"})

			return
		}
		defer stmt.Close()

		if _, err := stmt.Exec(lastname, firstname, username, email, string(hashedPassword)); err != nil {
			Log(ErrorLevel, "Error inserting the data into the database"+err.Error())
			//http.Error(w, "Error inserting the data into the database", http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "Error inserting the data into the database"})

			return
		}
		user_id, err := getUserID(username, db)
		if err != nil {
			return
		}
		user_uuid, cookie_exp, err := CreateSession(user_id, db)
		if err != nil {
			Log(ErrorLevel, "Error creating session")
			//http.Error(w, "Error creating session", http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "Error creating session"})

			return
		}
		cookie := http.Cookie{
			Name:     "session",
			Value:    user_uuid,
			Expires:  cookie_exp,
			SameSite: http.SameSiteStrictMode,
		}
		http.SetCookie(w, &cookie)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "Registration successful"})
		Log(DebugLevel, "Registration successful: "+username+" at "+r.URL.Path)
	}
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

		username := r.FormValue("usernameOrEmailLogin")
		password := r.FormValue("passwordLogin")

		// Fetch user from database
		var storedPassword string
		err := db.QueryRow("SELECT password FROM users WHERE username = $1 OR email = $2", username, username).Scan(&storedPassword)
		if err != nil {
			if err == sql.ErrNoRows {
				Log(ErrorLevel, "No user found with the provided credentials"+username+" at "+r.URL.Path)
				fmt.Println("No user found with the provided credentials")
				http.Error(w, "No user found with the provided credentials", http.StatusUnauthorized)
				return
			} else {
				Log(ErrorLevel, "Error fetching user from database"+username+" at "+r.URL.Path)
				fmt.Println("Error fetching user from database")
				http.Error(w, "Error fetching user from database", http.StatusInternalServerError)
				return
			}
		}

		// Compare the stored hashed password with the password that was submitted
		err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
		if err != nil {
			Log(DebugLevel, "Invalid login credentials"+username+" at "+r.URL.Path)
			fmt.Println("Invalid login credentials")
			http.Error(w, "Invalid login credentials", http.StatusUnauthorized)
			return
		}

		// Login successful
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "Login successful"})
		Log(DebugLevel, "Login successful: "+username+" at "+r.URL.Path)
		fmt.Println("Login successful")

	}
}

func CreateSession(user_id int, db *sql.DB) (string, time.Time, error) {
	user_uuid := UUID.NewV4().String()
	createdAt := time.Now()
	expireAt := createdAt.Add(Cookie_Expiration)
	_, err := db.Exec("INSERT INTO Sessions(user_id,uuid,created_at,expire_at) VALUES($1,$2,$3,$4)", user_id, user_uuid, createdAt, expireAt)
	if err != nil {
		Log(ErrorLevel, "Error creating session "+err.Error())
	}
	return user_uuid, expireAt, err
}

func getUserID(username string, db *sql.DB) (int, error) {
	var id int
	err := db.QueryRow("SELECT id_student FROM users WHERE username = $1", username).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
