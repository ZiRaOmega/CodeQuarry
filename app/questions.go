package app

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Question represents the data structure for a question
type Question struct {
	Id           int        `json:"id"`
	SubjectTitle string     `json:"subject_title"`
	SubjectID    int        `json:"subject_id"`
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	Content      string     `json:"content"`
	CreationDate time.Time  `json:"creation_date"`
	Creator      string     `json:"creator"`
	Upvotes      int        `json:"upvotes"`
	Downvotes    int        `json:"downvotes"`
	Responses    []Response `json:"responses"`
}

// FetchQuestionsBySubject retrieves a list of questions based on the subject ID.
// If the subject ID is "all", it fetches all questions from the database.
// Otherwise, it fetches questions only for the specified subject ID.
// It returns a slice of Question structs and an error if any.
func FetchQuestionsBySubject(db *sql.DB, subjectID string) ([]Question, error) {
	var questions []Question
	var rows *sql.Rows
	var err error

	if subjectID == "all" {
		query := `SELECT q.id_question, s.title AS subject_title, q.title, q.description, q.content, q.creation_date, u.username, q.upvotes, q.downvotes
                  FROM question q
                  JOIN users u ON q.id_student = u.id_student
                  JOIN subject s ON q.id_subject = s.id_subject`
		rows, err = db.Query(query)
	} else {
		query := `SELECT q.id_question, s.title AS subject_title, q.title, q.description, q.content, q.creation_date, u.username, q.upvotes, q.downvotes
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
		if err := rows.Scan(&q.Id, &q.SubjectTitle, &q.Title, &q.Description, &q.Content, &q.CreationDate, &q.Creator, &q.Upvotes, &q.Downvotes); err != nil {
			log.Printf("Error scanning question: %v", err)
			continue
		}
		q.Responses, err = FetchResponseByQuestion(db, q.Id)
		if err != nil {
			log.Printf("Error fetching responses: %v", err)
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

// FetchQuestionByQuestionID retrieves a question from the database based on the given question ID.
// It returns the retrieved question and an error, if any.
func FetchQuestionByQuestionID(db *sql.DB, questionID int) (Question, error) {
	var q Question
	query := `SELECT q.id_question, s.title AS subject_title, q.title, q.description, q.content, q.creation_date, u.username, q.upvotes, q.downvotes
				  FROM question q
				  JOIN users u ON q.id_student = u.id_student	
				  JOIN subject s ON q.id_subject = s.id_subject
				  WHERE q.id_question = $1`
	err := db.QueryRow(query, questionID).Scan(&q.Id, &q.SubjectTitle, &q.Title, &q.Description, &q.Content, &q.CreationDate, &q.Creator, &q.Upvotes, &q.Downvotes)
	if err != nil {
		return q, err
	}
	q.Responses, err = FetchResponseByQuestion(db, q.Id)
	if err != nil {
		return q, err
	}
	return q, nil
}

// GetUsernameWithUserID retrieves the username associated with the given userID from the database.
// It takes a *sql.DB pointer and an integer userID as parameters.
// It returns the username as a string. If an error occurs, it returns an empty string.
func GetUsernameWithUserID(db *sql.DB, userID int) string {
	var username string
	err := db.QueryRow(`SELECT username FROM users WHERE id_student = $1`, userID).Scan(&username)
	if err != nil {
		return ""
	}
	return username

}

// QuestionsHandler handles the HTTP request for retrieving questions by subject ID.
// It takes a database connection as input and returns an http.HandlerFunc.
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

// QuestionViewerHandler handles the HTTP request for viewing a question.
// It takes a database connection as input and returns an http.HandlerFunc.
// The returned handler function fetches the question details from the database
// based on the provided question_id parameter and renders the question viewer template.
func QuestionViewerHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		questionID := r.URL.Query().Get("question_id")
		if questionID == "" {
			http.Error(w, "Missing question_id parameter", http.StatusBadRequest)
			return
		}
		idint, err := strconv.Atoi(questionID)
		if err != nil {
			http.Error(w, "Invalid question_id parameter", http.StatusBadRequest)
			return
		}
		questions, err := FetchQuestionByQuestionID(db, idint)
		if err != nil {
			http.Error(w, "Error fetching responses", http.StatusInternalServerError)
			return
		}
		err = ParseAndExecuteTemplate("question_viewer", questions, w)
		if err != nil {
			http.Error(w, "Error parsing template", http.StatusInternalServerError)
			return

		}
	}

}

// CreateQuestion inserts a new question into the database.
// It takes a database connection, a question object, a user ID, and a subject ID as parameters.
// It returns an error if the insertion fails.
func CreateQuestion(db *sql.DB, question Question, user_id int, subject_id int) error {
	query := `INSERT INTO question (title, description, content, creation_date, update_date, id_student, id_subject, upvotes, downvotes)
			  VALUES ($1, $2, $3, $4, $5, $6, $7, 0, 0)`
	_, err := db.Exec(query, question.Title, question.Description, question.Content, time.Now(), time.Now(), user_id, subject_id)
	if err != nil {
		log.Printf("Error inserting question: %v", err)
		// Additional debugging information
		log.Printf("Attempted to insert: title='%s', user_id=%d, subject_id=%d", question.Title, user_id, subject_id)
		return err
	}
	return nil
}

type Subject struct {
	Id            int    `json:"id"`
	Title         string `json:"title"`
	QuestionCount int    `json:"questionCount"`
}

// FetchSubjectWithQuestionCount fetches a subject from the database along with the count of its associated questions.
// It takes a database connection (db) and a subject ID (subjectId) as parameters.
// It returns the fetched subject and an error, if any.
func FetchSubjectWithQuestionCount(db *sql.DB, subjectId int) (Subject, error) {
	var subject Subject
	query := `SELECT s.id_subject, s.title, COUNT(q.id_question) as question_count
			  FROM subject s
			  LEFT JOIN question q ON s.id_subject = q.id_subject
			  WHERE s.id_subject = $1
			  GROUP BY s.id_subject`
	err := db.QueryRow(query, subjectId).Scan(&subject.Id, &subject.Title, &subject.QuestionCount)
	if err != nil {
		log.Printf("Error fetching subject: %v", err)
		return Subject{}, err
	}
	return subject, nil
}

func FetchQuestionsByUserID(db *sql.DB, userID int) ([]Question, error) {
	var questions []Question
	rows, err := db.Query(`SELECT q.id_question, s.title AS subject_title, q.title, q.content, q.creation_date, u.username, q.upvotes, q.downvotes
						   FROM question q
						   JOIN users u ON q.id_student = u.id_student
						   JOIN subject s ON q.id_subject = s.id_subject
						   WHERE u.id_student = $1`, userID)
	if err != nil {
		log.Printf("Error querying questions: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var q Question
		if err := rows.Scan(&q.Id, &q.SubjectTitle, &q.Title, &q.Content, &q.CreationDate, &q.Creator, &q.Upvotes, &q.Downvotes); err != nil {
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
