package app

import (
	"database/sql"
	"fmt"
	"net/http"
)

func FilterHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":

		case "GET":
			r.ParseForm()
			// Get the form data
			author := r.FormValue("author")
			/* newer := r.FormValue("newer")
			older := r.FormValue("older") */
			// Call the filter function
			questions := FilterQuestions(db, author, true, false)
			err := ParseAndExecuteTemplate("home", questions, w)
			if err != nil {
				fmt.Println(err)
			}
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func FilterQuestions(db *sql.DB, author string, newer, older bool) []Question {
	// Query the database based on the filter criteria
	query := "SELECT id_question, title, description, content, upvotes,downvotes, creation_date, update_date FROM questions WHERE creator = ?"
	if newer {
		query += " AND creationdate > ?"
	}
	if older {
		query += " AND updateddate < ?"
	}
	// Execute the query
	rows, err := db.Query(query, author)
	if err != nil {
		// Handle the error
	}
	defer rows.Close()

	// Iterate over the rows and populate the result
	var questions []Question
	for rows.Next() {
		var question Question

		// Scan the row into the question struct
		err := rows.Scan(&question.Id, &question.Title, &question.Description, &question.Content, &question.Upvotes, &question.Downvotes, &question.CreationDate, &question.Creator, &question.SubjectID)
		if err != nil {
			// Handle the error
		}
		//Append Responses to the question
		responses, err := FetchResponseByQuestion(db, question.Id)
		if err != nil {
			//Handle error
		}
		question.SubjectID, err = FetchSubjectIDByQuestionID(db, question.Id)
		if err != nil {
			//handle error
		}
		question.SubjectTitle, err = FetchSubjectTitleByQuestionID(db, question.Id)
		if err != nil {
			//handle error
		}
		question.Responses = responses
		questions = append(questions, question)
	}

	// Check for any errors during iteration
	err = rows.Err()
	if err != nil {
		// Handle the error
	}

	// Return the filtered questions
	return questions
}
