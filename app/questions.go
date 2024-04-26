package app

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Question represents the data structure for a question
type Question struct {
	Id           int       `json:"id"`
	SubjectTitle string    `json:"subject_title"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	CreationDate time.Time `json:"creation_date"`
	Creator      string    `json:"creator"`
	Upvotes      int       `json:"upvotes"`
	Downvotes    int       `json:"downvotes"`
}

func FetchQuestionsBySubject(db *sql.DB, subjectID string) ([]Question, error) {
	var questions []Question
	var rows *sql.Rows
	var err error

	if subjectID == "all" {
		query := `SELECT s.title AS subject_title, q.title, q.content, q.creation_date, u.username, q.upvotes, q.downvotes
                  FROM question q
                  JOIN users u ON q.id_student = u.id_student
                  JOIN subject s ON q.id_subject = s.id_subject`
		rows, err = db.Query(query)
	} else {
		query := `SELECT s.title AS subject_title, q.title, q.content, q.creation_date, u.username, q.upvotes, q.downvotes
                  FROM question q
                  JOIN users u ON q.id_student = u.id_student
                  JOIN subject s ON q.id_subject = s.id_subject
                  WHERE q.id_subject = $1`
		rows, err = db.Query(query, subjectID)
	}

	if err != nil {
		log.Printf("Error querying questions: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var q Question
		if err := rows.Scan(&q.SubjectTitle, &q.Title, &q.Content, &q.CreationDate, &q.Creator, &q.Upvotes, &q.Downvotes); err != nil {
			log.Printf("Error scanning question: %v", err)
			continue
		}
		questions = append(questions, q)
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
