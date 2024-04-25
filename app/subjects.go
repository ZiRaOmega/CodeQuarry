package app

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// insertInSubject inserts a new subject into the database with additional fields
func InsertInSubject(db *sql.DB, title, description string) {
	currentTime := time.Now().Format("2006-01-02")

	// Check if the subject already exists
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM Subject WHERE title = $1)", title).Scan(&exists)
	if err != nil {
		log.Fatal("Error checking if subject exists: ", err)
	}

	if exists {
		return // Exit the function if the subject already exists
	}

	stmt, err := db.Prepare("INSERT INTO Subject(title, description, creation_date, update_date) VALUES($1, $2, $3, $4)")
	if err != nil {
		log.Fatal("Error preparing statement: ", err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(title, description, currentTime, currentTime); err != nil {
		log.Fatal("Error executing statement: ", err)
	} else {
		log.Printf("Inserted subject: %s", title)
	}
}

// Function to insert multiple subjects
func InsertMultipleSubjects(db *sql.DB) {
	subjects := []struct {
		Title       string
		Description string
	}{
		{"JavaScript", "Language for web development"},
		{"Golang", "Open source programming language that makes it easy to build simple, reliable, and efficient software"},
		{"C++", "General-purpose programming language created by Bjarne Stroustrup as an extension of the C programming language"},
		{"Ruby", "A dynamic, open source programming language with a focus on simplicity and productivity"},
		{"Rust", "A language empowering everyone to build reliable and efficient software"},
		{"Python", "An interpreted, high-level, general-purpose programming language"},
		{"Java", "A high-level, class-based, object-oriented programming language"},
	}
	fmt.Println("Inserting subjects...")
	for _, subject := range subjects {
		InsertInSubject(db, subject.Title, subject.Description)
	}
}

func FetchAllSubjects(db *sql.DB) ([]string, error) {
	var subjects []string
	query := "SELECT title FROM Subject ORDER BY title ASC"
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error querying subjects: %v", err)
		return nil, err
	}
	defer rows.Close()

	var title string
	for rows.Next() {
		if err := rows.Scan(&title); err != nil {
			log.Printf("Error scanning subject: %v", err)
			continue
		}
		subjects = append(subjects, title)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error reading subject rows: %v", err)
		return nil, err
	}
	return subjects, nil
}

func SubjectsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		subjects, err := FetchAllSubjects(db)
		if err != nil {
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(subjects)
	}
}
