package app

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type Response struct {
	ResponseID   int       `json:"response_id"`
	Description  string    `json:"description"`
	Content      string    `json:"content"`
	UpVotes      int       `json:"upvotes"`
	DownVotes    int       `json:"downvotes"`
	BestAnswer   bool      `json:"best_answer"`
	CreationDate time.Time `json:"creation_date"`
	UpdateDate   time.Time `json:"update_date"`
	QuestionID   int       `json:"question_id"`
	StudentID    int       `json:"student_id"`
	StudentName  string    `json:"student_name"`
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
			var response Response
			var receive_data interface{}
			err := json.NewDecoder(r.Body).Decode(&receive_data)
			fmt.Println(receive_data)
			fmt.Println(response)
			session_id := receive_data.(map[string]interface{})["session_id"].(string)
			question_id := receive_data.(map[string]interface{})["response"].(map[string]interface{})["question_id"].(string)
			description := receive_data.(map[string]interface{})["response"].(map[string]interface{})["description"].(string)
			content := receive_data.(map[string]interface{})["response"].(map[string]interface{})["content"].(string)
			creation_date := time.Now()
			update_date := time.Now()
			userid, err := getUserIDUsingSessionID(session_id, db)
			if userid == 0 || err != nil {
				http.Error(w, "Invalid session", http.StatusBadRequest)
				return
			}
			question_id_convert, err := strconv.Atoi(question_id)
			response = Response{Description: description, Content: content, UpVotes: 0, DownVotes: 0, BestAnswer: false, CreationDate: creation_date, UpdateDate: update_date, QuestionID: question_id_convert, StudentID: userid}
			if err != nil {
				http.Error(w, "Invalid question_id", http.StatusBadRequest)
				return
			}
			// Insert response into database
			err = InsertResponse(db, response)
			if err != nil {
				http.Error(w, "Error inserting response", http.StatusInternalServerError)
				return
			}
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

		}
	}
}
func InsertResponse(db *sql.DB, response Response) error {
	query := `INSERT INTO response (description, content, upvotes, downvotes, best_answer, creation_date, update_date, id_question, id_student) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err := db.Exec(query, response.Description, response.Content, 0, 0, false, response.CreationDate, response.UpdateDate, response.QuestionID, response.StudentID)
	if err != nil {
		return err
	}
	return nil
}
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
