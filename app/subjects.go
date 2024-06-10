package app

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Subject struct {
	Id            int        `json:"id"`
	Title         string     `json:"title"`
	Description   string     `json:"description"`
	CreationDate  time.Time  `json:"creation_date"`
	UpdateDate    time.Time  `json:"update_date"`
	QuestionCount int        `json:"questionCount"`
	Questions     []Question `json:"questions"`
}

// insertInSubject inserts a new subject into the database with additional fields
func InsertInSubject(db *sql.DB, title, description string) error {
	currentTime := time.Now().Format("2006-01-02")

	// Check if the subject already exists
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM Subject WHERE title = $1)", title).Scan(&exists)
	if err != nil {
		return err
	}

	if exists {
		//return new error
		return fmt.Errorf("Subject with title %s already exists", title)
	}

	stmt, err := db.Prepare("INSERT INTO Subject(title, description, creation_date, update_date) VALUES($1, $2, $3, $4)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(title, description, currentTime, currentTime); err != nil {
		return err
	}
	return err
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
		{"C#", "A general-purpose, multi-paradigm programming language"},
		{"Swift", "A powerful and intuitive programming language for macOS, iOS, watchOS, and tvOS"},
	}
	//
	for _, subject := range subjects {
		err := InsertInSubject(db, subject.Title, subject.Description)
		if err != nil {
			log.Printf("Error inserting subject %s: %v", subject.Title, err)
		}
	}
}

func FetchAllSubjects(db *sql.DB, user_id int) ([]Subject, error) {
	var subjects []Subject
	query := `
    SELECT s.id_subject, s.title, s.description, COUNT(q.id_question) as question_count
    FROM subject s
    LEFT JOIN question q ON s.id_subject = q.id_subject
    GROUP BY s.id_subject
    ORDER BY s.title ASC`
	// this query can be explained as follows:
	// 1. Select the subject id, title, description, and count of questions
	// 2. From the subject table
	// 3. Left join the question table on the subject id.
	// 4. Group the results by the subject id.
	// 5. Order the results by the subject title in ascending order

	// this query can be optimized by using a subquery to get the count of questions

	rows, err := db.Prepare(query)
	if err != nil {
		log.Printf("Error querying subjects: %v", err)
		return nil, err
	}
	// defer rows.Close()
	stmt, err := rows.Query()
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	for stmt.Next() {
		var id, title, description string
		var questionCount int
		if err := stmt.Scan(&id, &title, &description, &questionCount); err != nil {
			log.Printf("Error scanning subject: %v", err)
			continue
		}
		idint, err := strconv.Atoi(id)
		if err != nil {
			log.Printf("Error converting id to int: %v", err)
			continue
		}
		questions, err := FetchQuestionsBySubject(db, id, user_id)
		if err != nil {
			log.Printf("Error fetching questions for subject %s: %v", title, err)
			continue
		}
		subjects = append(subjects, Subject{Id: idint, Title: title, Description: description, QuestionCount: questionCount, Questions: questions})
	}
	if err := stmt.Err(); err != nil {
		log.Printf("Error reading subject rows: %v", err)
		return nil, err
	}
	return subjects, nil
}

func SubjectsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Get cookie
		cookie, err := r.Cookie("session")
		if err != nil {
			log.Printf("Error getting session cookie: %v", err)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		session_id := cookie.Value
		if !isValidSession(session_id, db) {
			log.Println("Invalid session")
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		//Get user_id
		user_id, err := GetUserIDUsingSessionID(session_id, db)
		if err != nil {
			log.Printf("Error getting user id: %v", err)
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}
		subjects, err := FetchAllSubjects(db, user_id)
		if err != nil {
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(subjects)
	}
}
