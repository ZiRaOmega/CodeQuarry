package app

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"codequarry/app/utils"
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
	UserVote     string    `json:"user_vote"`
	IsAuthor     bool      `json:"is_author"`
}

type ResponseVote struct {
	R        int  `json:"response"`
	UpVote   bool `json:"up_vote"`
	DownVote bool `json:"down_vote"`
}

func ResponsesHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			// Handle post request
			var response Response
			var receive_data interface{}
			err := json.NewDecoder(r.Body).Decode(&receive_data)
			if err != nil {
				json.NewEncoder(w).Encode(map[string]string{"error": "Error decoding request body"})
				http.Error(w, "Error decoding request body", http.StatusBadRequest)
				return
			}
			session_id := receive_data.(map[string]interface{})["session_id"].(string)
			question_id := receive_data.(map[string]interface{})["response"].(map[string]interface{})["question_id"].(string)
			description := receive_data.(map[string]interface{})["response"].(map[string]interface{})["description"].(string)
			content := receive_data.(map[string]interface{})["response"].(map[string]interface{})["content"].(string)
			//CheckXSS and SQLi
			if utils.ContainsSQLi(description) || utils.ContainsSQLi(content) {
				json.NewEncoder(w).Encode(map[string]string{"error": "SQL injection detected"})
				http.Error(w, "SQL injection detected", http.StatusBadRequest)
				return
			}
			if utils.ContainsXSS(description) || utils.ContainsXSS(content) {
				json.NewEncoder(w).Encode(map[string]string{"error": "XSS detected"})
				http.Error(w, "XSS detected", http.StatusBadRequest)
				return
			}
			creation_date := time.Now()

			userid, err := getUserIDUsingSessionID(session_id, db)
			if userid == 0 || err != nil {
				json.NewEncoder(w).Encode(map[string]string{"error": "Invalid session"})
				http.Error(w, "Invalid session", http.StatusBadRequest)
				return
			}
			question_id_convert, err := strconv.Atoi(question_id)
			//Check if question as best answer
			question_best_answer := GetBestAnswerFromQuestion(db, question_id_convert)
			if question_best_answer != -1 {
				json.NewEncoder(w).Encode(map[string]string{"error": "Question already has a best answer"})
				http.Error(w, "Question already has a best answer", http.StatusBadRequest)
				return
			}
			response = Response{Description: description, Content: content, UpVotes: 0, DownVotes: 0, BestAnswer: false, CreationDate: creation_date, UpdateDate: creation_date, QuestionID: question_id_convert, StudentID: userid, StudentName: GetUsernameWithUserID(db, userid)}
			if err != nil {
				json.NewEncoder(w).Encode(map[string]string{"error": "Invalid question_id"})
				http.Error(w, "Invalid question_id", http.StatusBadRequest)
				return
			}
			// Insert response into database
			err = InsertResponse(db, response)
			if err != nil {
				json.NewEncoder(w).Encode(map[string]string{"error": "Error inserting response"})
				http.Error(w, "Error inserting response", http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "Response successful"})
			wsmessage := WSMessage{Type: "response", Content: response}
			BroadcastMessage(wsmessage, nil)
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
	// Add XP to the student who answered the question
	err = InsertXP(db, response.StudentID, 100)
	if err != nil {
		return err
	}
	Question_Student_ID, err := FetchStudentIDUsingQuestionID(db, response.QuestionID)
	if err != nil {
		return err
	}
	// Add XP to the student who asked the question
	err = InsertXP(db, Question_Student_ID, 100)
	if err != nil {
		return err
	}
	return nil
}

func FetchResponseByQuestion(db *sql.DB, questionID int, user_id int) ([]Response, error) {
	var responses []Response
	query := `SELECT * FROM Response WHERE id_question = $1`
	rows, err := db.Query(query, questionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	voted_responses, err := FetchVotedResponses(db, user_id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var r Response
		if err := rows.Scan(&r.ResponseID, &r.Description, &r.Content, &r.UpVotes, &r.DownVotes, &r.BestAnswer, &r.CreationDate, &r.UpdateDate, &r.QuestionID, &r.StudentID); err != nil {
			return nil, err
		}
		for _, voted_response := range voted_responses {
			if voted_response.R == r.ResponseID && voted_response.UpVote {
				r.UserVote = "upvoted"
			} else if voted_response.R == r.ResponseID && voted_response.DownVote {
				r.UserVote = "downvoted"
			}
		}
		if user_id == r.StudentID {
			r.IsAuthor = true
		}

		r.StudentName = GetUsernameWithUserID(db, r.StudentID)
		responses = append(responses, r)
	}
	return responses, nil
}

// FetchVotedResponses retrieves responses voted by a specific user
func FetchVotedResponses(db *sql.DB, userID int) ([]ResponseVote, error) {
	query := `
        SELECT r.id_response, v.upvote_r, v.downvote_r
        FROM Response r
        JOIN users u ON r.id_student = u.id_student
        JOIN vote_response v ON r.id_response = v.id_response
        WHERE v.id_student = $1
    `

	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var responseVotes []ResponseVote
	for rows.Next() {
		var r int
		var rv ResponseVote
		var upVote, downVote bool

		if err := rows.Scan(&r, &upVote, &downVote); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		rv.R = r
		rv.UpVote = upVote
		rv.DownVote = downVote

		responseVotes = append(responseVotes, rv)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row error: %w", err)
	}

	return responseVotes, nil
}

// InsertBestAnswer toggles the best_answer status for a given response.
func InsertBestAnswer(db *sql.DB, responseID int) error {
	// First, check the current status of best_answer.
	currentStatus := CheckIfAleadyBestAnswer(db, responseID)
	// Toggle the status: true becomes false, false becomes true.
	newStatus := !currentStatus

	query := `UPDATE response SET best_answer = $1 WHERE id_response = $2`
	_, err := db.Exec(query, newStatus, responseID)
	if err != nil {
		return err
	}
	return nil
}

// CheckIfAleadyBestAnswer checks if the given response is already marked as the best answer.
func CheckIfAleadyBestAnswer(db *sql.DB, responseID int) bool {
	query := `SELECT best_answer FROM response WHERE id_response = $1`
	var bestAnswer bool
	err := db.QueryRow(query, responseID).Scan(&bestAnswer)
	if err != nil {
		// Handle error according to your logging or error handling strategy.
		// For simplicity, return false on error.
		return false
	}
	return bestAnswer
}

func GetBestAnswerFromQuestion(db *sql.DB, questionID int) int {
	//find the best answer when its true
	query := `SELECT id_response FROM response WHERE id_question = $1 AND best_answer = true`
	var bestAnswer int
	err := db.QueryRow(query, questionID).Scan(&bestAnswer)
	if err != nil {
		// Handle error according to your logging or error handling strategy.
		// For simplicity, return false on error.
		return -1
	}
	return bestAnswer
}
