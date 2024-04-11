package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type PageData struct {
	Content string
}

func registerHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Error parsing form", http.StatusBadRequest)
				return
			}

			fullname := r.FormValue("fullname")
			username := r.FormValue("username")
			email := r.FormValue("email")
			password := r.FormValue("password")

			// Hashing the password with a cost of 14
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			if err != nil {
				fmt.Println(err, "Error hashing password")
				return
			}

			stmt, err := db.Prepare("INSERT INTO users(fullname, username, email, password) VALUES(?, ?, ?, ?)")
			if err != nil {
				fmt.Println(err, "Error preparing the SQL statement")
				return
			}
			defer stmt.Close()

			_, err = stmt.Exec(fullname, username, email, string(hashedPassword))
			if err != nil {
				fmt.Println(err, "Error inserting the data into the database")
				return
			}

		}
	}
}
