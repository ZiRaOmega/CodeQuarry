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

func FetchAllSubjects(db *sql.DB) ([]map[string]string, error) {
	var subjects []map[string]string
	query := "SELECT id_subject, title, description FROM Subject ORDER BY title ASC"
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error querying subjects: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id, title, description string
		if err := rows.Scan(&id, &title, &description); err != nil {
			log.Printf("Error scanning subject: %v", err)
			continue
		}
		subjects = append(subjects, map[string]string{"id": id, "title": title, "description": description})
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

// FetchQuestionsBySubject fetches questions for a specific subject from the database
func FetchQuestionsBySubject(db *sql.DB, subjectID string) ([]string, error) {
	var questions []string
	query := "SELECT title FROM question WHERE id_subject = $1"
	rows, err := db.Query(query, subjectID)
	if err != nil {
		log.Printf("Error querying questions: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var question string
		if err := rows.Scan(&question); err != nil {
			log.Printf("Error scanning question: %v", err)
			continue
		}
		questions = append(questions, question)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error reading question rows: %v", err)
		return nil, err
	}
	return questions, nil
}

// QuestionsHandler handles the API endpoint for fetching questions based on subject ID
func QuestionsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		subjectID := r.URL.Query().Get("subjectId")
		questions, err := FetchQuestionsBySubject(db, subjectID)
		if err != nil {
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(questions)
	}
}
