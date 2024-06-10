package app

import (
	"database/sql"
)

// HandleUpvote updates the upvote count for a question
func HandleUpvote(db *sql.DB, questionID float64, sessionID string) error {
	user_Id, err := GetUserIDUsingSessionID(sessionID, db)
	if err != nil {
		return err
	}
	if !CheckIfUserAlreadyVotedForThisQuestion(db, questionID, float64(user_Id)) {

		InsertUserInVoteQuestionDb(db, questionID, float64(user_Id), false, false)
	}
	// Increment upvote and potentially decrement downvote if already voted
	if !CheckIfUserUpvoted(db, questionID, float64(user_Id)) && !CheckIfUserDownvoted(db, questionID, float64(user_Id)) {

		UpdateUserInVoteQuestionDb(db, questionID, float64(user_Id), true, false)
		query := `UPDATE question SET upvotes = upvotes + 1 WHERE id_question = $1`
		_, err = db.Exec(query, questionID)
		if err != nil {
			return err
		}
		return nil
	} else if CheckIfUserUpvoted(db, questionID, float64(user_Id)) && !CheckIfUserDownvoted(db, questionID, float64(user_Id)) {

		UpdateUserInVoteQuestionDb(db, questionID, float64(user_Id), false, false)
		query := `UPDATE question SET upvotes = upvotes - 1 WHERE id_question = $1`
		_, err = db.Exec(query, questionID)
		if err != nil {
			return err
		}
		return nil
	} else if !CheckIfUserUpvoted(db, questionID, float64(user_Id)) && CheckIfUserDownvoted(db, questionID, float64(user_Id)) {

		UpdateUserInVoteQuestionDb(db, questionID, float64(user_Id), true, false)
		query := `UPDATE question SET upvotes = upvotes + 1 WHERE id_question = $1`
		_, err = db.Exec(query, questionID)
		if err != nil {
			return err
		}
		query = `UPDATE question SET downvotes = downvotes - 1 WHERE id_question = $1`
		_, err = db.Exec(query, questionID)
		if err != nil {
			return err
		}
		return nil
	} else {

		UpdateUserInVoteQuestionDb(db, questionID, float64(user_Id), false, false)
		query := `UPDATE question SET upvotes = upvotes - 1 WHERE id_question = $1`
		_, err = db.Exec(query, questionID)
		if err != nil {
			return err
		}
		return nil
	}

}

func SendNewVoteCount(db *sql.DB, questionID float64) (int, int) {
	query := `SELECT upvotes, downvotes FROM question WHERE id_question = $1`
	var upvotes int
	var downvotes int
	err := db.QueryRow(query, questionID).Scan(&upvotes, &downvotes)
	if err != nil {
		return 0, 0
	}
	return upvotes, downvotes
}

// HandleDownvote updates the downvote count for a question
func HandleDownvote(db *sql.DB, questionID float64, sessionID string) error {
	user_Id, err := GetUserIDUsingSessionID(sessionID, db)
	if err != nil {
		return err
	}
	if !CheckIfUserAlreadyVotedForThisQuestion(db, questionID, float64(user_Id)) {

		InsertUserInVoteQuestionDb(db, questionID, float64(user_Id), false, false)
	}
	// Increment downvote and potentially decrement upvote if already voted
	if !CheckIfUserDownvoted(db, questionID, float64(user_Id)) && !CheckIfUserUpvoted(db, questionID, float64(user_Id)) {

		UpdateUserInVoteQuestionDb(db, questionID, float64(user_Id), false, true)
		query := `UPDATE question SET downvotes = downvotes + 1 WHERE id_question = $1`
		_, err = db.Exec(query, questionID)
		if err != nil {
			return err
		}
		return nil
	} else if CheckIfUserDownvoted(db, questionID, float64(user_Id)) && !CheckIfUserUpvoted(db, questionID, float64(user_Id)) {

		UpdateUserInVoteQuestionDb(db, questionID, float64(user_Id), false, false)
		query := `UPDATE question SET downvotes = downvotes - 1 WHERE id_question = $1`
		_, err = db.Exec(query, questionID)
		if err != nil {
			return err
		}
		return nil
	} else if !CheckIfUserDownvoted(db, questionID, float64(user_Id)) && CheckIfUserUpvoted(db, questionID, float64(user_Id)) {

		UpdateUserInVoteQuestionDb(db, questionID, float64(user_Id), false, true)
		query := `UPDATE question SET downvotes = downvotes + 1 WHERE id_question = $1`
		_, err = db.Exec(query, questionID)
		if err != nil {
			return err
		}
		query = `UPDATE question SET upvotes = upvotes - 1 WHERE id_question = $1`
		_, err = db.Exec(query, questionID)
		if err != nil {
			return err
		}
		return nil
	} else {

		UpdateUserInVoteQuestionDb(db, questionID, float64(user_Id), false, false)
		query := `UPDATE question SET downvotes = downvotes - 1 WHERE id_question = $1`
		_, err = db.Exec(query, questionID)
		if err != nil {
			return err
		}
		return nil

	}
}

func CheckIfUserDownvoted(db *sql.DB, questionID float64, studentID float64) bool {
	query := `SELECT downvote_q FROM Vote_question WHERE id_student = $1 AND id_question = $2`
	var downvote bool
	err := db.QueryRow(query, studentID, questionID).Scan(&downvote)
	if err != nil {
		return false
	}
	return downvote
}

func CheckIfUserUpvoted(db *sql.DB, questionID float64, studentID float64) bool {
	query := `SELECT upvote_q FROM Vote_question WHERE id_student = $1 AND id_question = $2`
	var upvote bool
	err := db.QueryRow(query, studentID, questionID).Scan(&upvote)
	if err != nil {
		return false
	}
	return upvote
}

func CheckIfUserAlreadyVotedForThisQuestion(db *sql.DB, questionID float64, studentID float64) bool {
	//count number of rows where id_student = studentID and id_question = questionID
	query := `SELECT EXISTS(SELECT 1 FROM vote_question WHERE id_student = $1 AND id_question = $2)`
	var exists bool
	err := db.QueryRow(query, studentID, questionID).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}

func InsertUserInVoteQuestionDb(db *sql.DB, questionID float64, studentID float64, upvote bool, downvote bool) error {
	//insert user in vote_question table if not already present
	query := `INSERT INTO vote_question (id_student, id_question, upvote_q, downvote_q) VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(query, studentID, questionID, upvote, downvote)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserInVoteQuestionDb(db *sql.DB, questionID float64, studentID float64, upvote bool, downvote bool) error {
	query := `UPDATE vote_question SET upvote_q = $3, downvote_q = $4 WHERE id_student = $1 AND id_question = $2`
	_, err := db.Exec(query, studentID, questionID, upvote, downvote)
	if err != nil {
		return err
	}
	return nil
}

//upvotes and downvotes for responses

// HandleUpvoteResponse updates the upvote count for a response
func HandleUpvoteResponse(db *sql.DB, responseID float64, sessionID string) error {
	user_Id, err := GetUserIDUsingSessionID(sessionID, db)
	if err != nil {
		return err
	}
	if !CheckIfUserAlreadyVotedForThisResponse(db, responseID, float64(user_Id)) {

		InsertUserInVoteResponseDb(db, responseID, float64(user_Id), false, false)
	}
	// Increment upvote and potentially decrement downvote if already voted
	if !CheckIfUserUpvotedResponse(db, responseID, float64(user_Id)) && !CheckIfUserDownvotedResponse(db, responseID, float64(user_Id)) {

		UpdateUserInVoteResponseDb(db, responseID, float64(user_Id), true, false)
		query := `UPDATE response SET upvotes = upvotes + 1 WHERE id_response = $1`
		_, err = db.Exec(query, responseID)
		if err != nil {
			return err
		}
		return nil
	} else if CheckIfUserUpvotedResponse(db, responseID, float64(user_Id)) && !CheckIfUserDownvotedResponse(db, responseID, float64(user_Id)) {

		UpdateUserInVoteResponseDb(db, responseID, float64(user_Id), false, false)
		query := `UPDATE response SET upvotes = upvotes - 1 WHERE id_response = $1`
		_, err = db.Exec(query, responseID)
		if err != nil {
			return err
		}
		return nil
	} else if !CheckIfUserUpvotedResponse(db, responseID, float64(user_Id)) && CheckIfUserDownvotedResponse(db, responseID, float64(user_Id)) {

		UpdateUserInVoteResponseDb(db, responseID, float64(user_Id), true, false)
		query := `UPDATE response SET upvotes = upvotes + 1 WHERE id_response = $1`
		_, err = db.Exec(query, responseID)
		if err != nil {
			return err
		}
		query = `UPDATE response SET downvotes = downvotes - 1 WHERE id_response = $1`
		res, _ := db.Query(query, responseID)
		if res != nil {
			return err
		}
		return nil
	} else {

		UpdateUserInVoteResponseDb(db, responseID, float64(user_Id), false, false)
		query := `UPDATE response SET upvotes = upvotes - 1 WHERE id_response = $1`
		_, err = db.Exec(query, responseID)
		if err != nil {
			return err
		}
		return nil
	}
}

func CheckIfUserDownvotedResponse(db *sql.DB, responseID float64, studentID float64) bool {
	query := `SELECT downvote_r FROM Vote_response WHERE id_student = $1 AND id_response = $2`
	var downvote bool
	err := db.QueryRow(query, studentID, responseID).Scan(&downvote)
	if err != nil {
		return false
	}
	return downvote
}

func CheckIfUserUpvotedResponse(db *sql.DB, responseID float64, studentID float64) bool {
	query := `SELECT upvote_r FROM Vote_response WHERE id_student = $1 AND id_response = $2`
	var upvote bool
	err := db.QueryRow(query, studentID, responseID).Scan(&upvote)
	if err != nil {
		return false
	}
	return upvote
}

func CheckIfUserAlreadyVotedForThisResponse(db *sql.DB, responseID float64, studentID float64) bool {
	//count number of rows where id_student = studentID and id_response = responseID
	query := `SELECT EXISTS(SELECT 1 FROM vote_response WHERE id_student = $1 AND id_response = $2)`
	var exists bool
	err := db.QueryRow(query, studentID, responseID).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}

func InsertUserInVoteResponseDb(db *sql.DB, responseID float64, studentID float64, upvote bool, downvote bool) error {
	//insert user in vote_response table if not already present
	query := `INSERT INTO vote_response (id_student, id_response, upvote_r, downvote_r) VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(query, studentID, responseID, upvote, downvote)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserInVoteResponseDb(db *sql.DB, responseID float64, studentID float64, upvote bool, downvote bool) error {
	query := `UPDATE vote_response SET upvote_r = $3, downvote_r = $4 WHERE id_student = $1 AND id_response = $2`
	_, err := db.Exec(query, studentID, responseID, upvote, downvote)
	if err != nil {
		return err
	}
	return nil
}

// HandleDownvoteResponse updates the downvote count for a response

func HandleDownvoteResponse(db *sql.DB, responseID float64, sessionID string) error {
	user_Id, err := GetUserIDUsingSessionID(sessionID, db)
	if err != nil {
		return err
	}
	if !CheckIfUserAlreadyVotedForThisResponse(db, responseID, float64(user_Id)) {

		InsertUserInVoteResponseDb(db, responseID, float64(user_Id), false, false)
	}
	// Increment downvote and potentially decrement upvote if already voted
	if !CheckIfUserDownvotedResponse(db, responseID, float64(user_Id)) && !CheckIfUserUpvotedResponse(db, responseID, float64(user_Id)) {

		UpdateUserInVoteResponseDb(db, responseID, float64(user_Id), false, true)
		query := `UPDATE response SET downvotes = downvotes + 1 WHERE id_response = $1`
		_, err = db.Exec(query, responseID)
		if err != nil {
			return err
		}
		return nil
	} else if CheckIfUserDownvotedResponse(db, responseID, float64(user_Id)) && !CheckIfUserUpvotedResponse(db, responseID, float64(user_Id)) {

		UpdateUserInVoteResponseDb(db, responseID, float64(user_Id), false, false)
		query := `UPDATE response SET downvotes = downvotes - 1 WHERE id_response = $1`
		_, err = db.Exec(query, responseID)
		if err != nil {
			return err
		}
		return nil
	} else if !CheckIfUserDownvotedResponse(db, responseID, float64(user_Id)) && CheckIfUserUpvotedResponse(db, responseID, float64(user_Id)) {

		UpdateUserInVoteResponseDb(db, responseID, float64(user_Id), false, true)
		query := `UPDATE response SET downvotes = downvotes + 1 WHERE id_response = $1`
		_, err = db.Exec(query, responseID)
		if err != nil {
			return err
		}
		query = `UPDATE response SET upvotes = upvotes - 1 WHERE id_response = $1`
		res, _ := db.Query(query, responseID)
		if res != nil {
			return err
		}
		return nil
	} else {

		UpdateUserInVoteResponseDb(db, responseID, float64(user_Id), false, false)
		query := `UPDATE response SET downvotes = downvotes - 1 WHERE id_response = $1`
		_, err = db.Exec(query, responseID)
		if err != nil {
			return err
		}
		return nil
	}
}

func SendNewVoteCountResponse(db *sql.DB, responseID float64) (int, int) {
	query := `SELECT upvotes, downvotes FROM response WHERE id_response = $1`
	var upvotes int
	var downvotes int
	err := db.QueryRow(query, responseID).Scan(&upvotes, &downvotes)
	if err != nil {
		return 0, 0
	}
	return upvotes, downvotes
}
