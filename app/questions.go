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
	User_Id      int        `json:"user_id"`
	SubjectTitle string     `json:"subject_title"`
	SubjectID    int        `json:"subject_id"`
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	Content      string     `json:"content"`
	CreationDate time.Time  `json:"creation_date"`
	UpdateDate   time.Time  `json:"update_date"`
	Creator      string     `json:"creator"`
	Upvotes      int        `json:"upvotes"`
	Downvotes    int        `json:"downvotes"`
	Responses    []Response `json:"responses"`
	UserVote     string     `json:"user_vote"`
}
type QuestionViewer struct {
	Question   Question
	Rank_Panel sql.NullInt64
}

// FetchQuestionsBySubject retrieves a list of questions based on the subject ID.
// If the subject ID is "all", it fetches all questions from the database.
// Otherwise, it fetches questions only for the specified subject ID.
// It returns a slice of Question structs and an error if any.
func FetchQuestionsBySubject(db *sql.DB, subjectID string, user_id int) ([]Question, error) {
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
	voted_question, err := FetchVotedQuestions(db, user_id)
	if err != nil {
		log.Printf("Error fetching voted questions: %v", err)
		return nil, err
	}
	for rows.Next() {
		var q Question
		if err := rows.Scan(&q.Id, &q.SubjectTitle, &q.Title, &q.Description, &q.Content, &q.CreationDate, &q.Creator, &q.Upvotes, &q.Downvotes); err != nil {
			log.Printf("Error scanning question: %v", err)
			continue
		}
		for _, voted := range voted_question {
			if voted.Q == q.Id && voted.Upvote {
				q.UserVote = "upvoted"
			} else if voted.Q == q.Id && voted.Downvote {
				q.UserVote = "downvoted"
			}
		}
		q.Responses, err = FetchResponseByQuestion(db, q.Id, user_id)
		if err != nil {
			log.Printf("Error fetching responses: %v", err)
			continue
		}

		questions = append(questions, q)
	}
	if len(questions) == 0 {
		return []Question{}, nil

	}
	if err := rows.Err(); err != nil {
		log.Printf("Error reading question rows: %v", err)
		return nil, err
	}
	return questions, nil
}

// FetchQuestionByQuestionID retrieves a question from the database based on the given question ID.
// It returns the retrieved question and an error, if any.
func FetchQuestionByQuestionID(db *sql.DB, questionID int, user_id int) (Question, error) {
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
	voted_question, err := FetchVotedQuestions(db, user_id)
	if err != nil {
		log.Printf("Error fetching voted questions: %v", err)
		return Question{}, err
	}
	for _, voted := range voted_question {
		if voted.Q == questionID && voted.Upvote {
			q.UserVote = "upvoted"
		} else if voted.Q == questionID && voted.Downvote {
			q.UserVote = "downvoted"
		}
	}
	q.Responses, err = FetchResponseByQuestion(db, q.Id, user_id)
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
		question_id := r.URL.Query().Get("question_id")
		session_id, err := r.Cookie("session")
		if err != nil {
			http.Error(w, "Session not found", http.StatusUnauthorized)
			return
		}

		user_id, err := getUserIDUsingSessionID(session_id.Value, db)
		if err != nil {
			http.Error(w, "Session not found", http.StatusUnauthorized)
			//Redirect to auth
			http.Redirect(w, r, "/auth", http.StatusSeeOther)
			return
		}

		if question_id == "" && subjectID != "" {

			questions, err := FetchQuestionsBySubject(db, subjectID, user_id)
			if err != nil {

				http.Error(w, "Error while fetching", http.StatusInternalServerError)
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(questions)
		} else if question_id != "" && subjectID == "" {
			question_id_int, err := strconv.Atoi(question_id)
			if err != nil {

				http.Error(w, "Error while fetching", http.StatusInternalServerError)
			}
			questions, err := FetchQuestionByQuestionID(db, question_id_int, user_id)
			if err != nil {

				http.Error(w, "Error while fetching", http.StatusInternalServerError)
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode([]Question{questions})
		}

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
		//get cookie session
		session_id, err := r.Cookie("session")

		if err != nil {
			//http.Error(w, "Session not found", http.StatusUnauthorized)
			http.Redirect(w, r, "/auth", http.StatusSeeOther)
			return
		}
		user_id, err := getUserIDUsingSessionID(session_id.Value, db)
		if err != nil {
			//http.Error(w, "Session not found", http.StatusUnauthorized)
			http.Redirect(w, r, "/auth", http.StatusSeeOther)
			return
		}
		questions, err := FetchQuestionByQuestionID(db, idint, user_id)
		if err != nil {
			http.Error(w, "Error fetching responses", http.StatusInternalServerError)
			return
		}
		rank := FetchRankByUserID(db, user_id)
		question_viewer := QuestionViewer{Question: questions, Rank_Panel: sql.NullInt64{Int64: int64(rank), Valid: true}}
		err = ParseAndExecuteTemplate("question_viewer", question_viewer, w)
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
	// Insert xp for user after creating a question
	err = InsertXP(db, user_id, 1000)
	if err != nil {
		log.Printf("Error updating XP: %v", err)
		return err
	}
	return nil
}

func InsertXP(db *sql.DB, user_id int, xp int) error {
	query := `UPDATE users SET xp = xp + $1 WHERE id_student = $2`
	_, err := db.Exec(query, xp, user_id)
	if err != nil {
		log.Printf("Error updating XP: %v", err)
		return err
	}
	return nil
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

type QuestionVote struct {
	Upvote   bool
	Downvote bool
	Q        int
}

func FetchVotedQuestions(db *sql.DB, userID int) ([]QuestionVote, error) {
	var questions []QuestionVote
	rows, err := db.Query(`SELECT q.id_question, v.upvote_q, v.downvote_q
						   FROM question q
						   JOIN users u ON q.id_student = u.id_student
						   JOIN subject s ON q.id_subject = s.id_subject
						   JOIN Vote_question v ON q.id_question = v.id_question
						   WHERE v.id_student = $1`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var q int
		var qq QuestionVote
		if err := rows.Scan(&q, &qq.Upvote, &qq.Downvote); err != nil {
			continue
		}
		qq.Q = q
		questions = append(questions, qq)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return questions, nil
}

func FetchStudentIDUsingQuestionID(db *sql.DB, questionID int) (int, error) {
	var studentID int
	err := db.QueryRow(`SELECT id_student FROM question WHERE id_question = $1`, questionID).Scan(&studentID)
	if err != nil {
		return 0, err
	}
	return studentID, nil
}

func UserDeleteQuestion(db *sql.DB, questionID int, userID int) error {
	studentID, err := FetchStudentIDUsingQuestionID(db, questionID)
	if err != nil {
		return err
	}
	if studentID != userID {

		return nil
	}
	//RemoveXP for question author and for all users who answer the question
	RemoveXP(db, studentID, 1000)
	//Remove xp for all users who answered the question
	answers, err := FetchResponseByQuestion(db, questionID, userID)
	if err != nil {
		return err
	}
	for _, answer := range answers {
		RemoveXP(db, answer.StudentID, 100)
		RemoveXP(db, studentID, 100)
	}
	//Delete on cascade
	_, err = db.Exec(`DELETE FROM question WHERE id_question = $1`, questionID)
	if err != nil {
		return err
	}
	return nil
}

func RemoveXP(db *sql.DB, user_id int, xp int) error {
	query := `UPDATE users SET xp = xp - $1 WHERE id_student = $2`
	_, err := db.Exec(query, xp, user_id)
	if err != nil {
		log.Printf("Error updating XP: %v", err)
		return err
	}
	return nil
}
func CheckIfQuestionIsMine(db *sql.DB, questionID int, userID float64) bool {
	var id float64
	err := db.QueryRow(`SELECT id_student FROM question WHERE id_question = $1`, questionID).Scan(&id)
	if err != nil {
		return false
	}
	if id == userID {
		return true
	}
	return false
}
func FetchXP(db *sql.DB, user_id int) (int, error) {
	var xp int
	err := db.QueryRow(`SELECT xp FROM users WHERE id_student = $1`, user_id).Scan(&xp)
	if err != nil {
		return 0, err
	}
	return xp, nil
}
func getQuestionIDFromResponseID(db *sql.DB, responseID int) int {
	var questionID int
	err := db.QueryRow(`SELECT id_question FROM response WHERE id_response = $1`, responseID).Scan(&questionID)
	if err != nil {
		return 0
	}
	return questionID
}

func ModifyQuestion(db *sql.DB, questionID int, title string, description string, content string, user_id int) error {
	//Fetch question
	question, err := FetchQuestionByQuestionID(db, questionID, user_id)
	if err != nil {
		return err
	}
	//if title description or content is empty, keep the old value
	if title == "" {
		title = question.Title
	}
	if description == "" {
		description = question.Description
	}
	if content == "" {
		content = question.Content
	}
	_, err = db.Exec(`UPDATE question SET title = $1, description = $2, content = $3 WHERE id_question = $4`, title, description, content, questionID)
	if err != nil {
		return err
	}
	return nil
}
