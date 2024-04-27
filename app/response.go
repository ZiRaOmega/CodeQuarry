package app

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type Response struct {
	ResponseID   int
	Description  string
	Content      string
	UpVotes      int
	DownVotes    int
	BestAnswer   bool
	CreationDate time.Time
	UpdateDate   time.Time
	QuestionID   int
	StudentID    int
	StudentName  string
}

func ResponsesHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
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
			responses, err := FetchResponseByQuestion(db, idint)
			if err != nil {
				http.Error(w, "Error fetching responses", http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(responses)
		case http.MethodPost:
			// Handle post request
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
func FetchResponseByQuestion(db *sql.DB, questionID int) ([]Response, error) {
	var responses []Response
	query := `SELECT * FROM Response WHERE id_question = $1`
	rows, err := db.Query(query, questionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var r Response
		if err := rows.Scan(&r.ResponseID, &r.Description, &r.Content, &r.UpVotes, &r.DownVotes, &r.BestAnswer, &r.CreationDate, &r.UpdateDate, &r.QuestionID, &r.StudentID); err != nil {
			return nil, err
		}
		r.StudentName = GetUsernameWithUserID(db, r.StudentID)
		responses = append(responses, r)
	}
	return responses, nil
}
