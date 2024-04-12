package app

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type PageData struct {
	Content string
}

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

		fullname := r.FormValue("fullname")
		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Error hashing password", http.StatusInternalServerError)
			return
		}

		stmt, err := db.Prepare("INSERT INTO users(fullname, username, email, password) VALUES(?, ?, ?, ?)")
		if err != nil {
			http.Error(w, "Error preparing the SQL statement", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		if _, err := stmt.Exec(fullname, username, email, string(hashedPassword)); err != nil {
			http.Error(w, "Error inserting the data into the database", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "Registration successful"})
	}
}

func LoginHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Error parsing form", http.StatusBadRequest)
				return
			}

			username := r.FormValue("usernameOrEmailLogin")
			password := r.FormValue("passwordLogin")

			// Fetch user from database
			var storedPassword string
			err = db.QueryRow("SELECT password FROM users WHERE username = ? OR email = ?", username, username).Scan(&storedPassword)
			if err != nil {
				if err == sql.ErrNoRows {
					fmt.Println(err, "No user found")
				} else {
					fmt.Println(err, "Error fetching user from database")
				}
				return
			}

			// Compare the stored hashed password with the password that was submitted
			err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
			if err != nil {
				// If the hash does not match the password provided
				fmt.Println(err, "Invalid login credentials")
				return
			}

			// Login successful
			fmt.Fprintf(w, "Login successful")
		} else {
			fmt.Println("Invalid request method")
		}
	}
}
