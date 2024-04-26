package app

import (
	"database/sql"
	"fmt"
)

// HandleUpvote updates the upvote count for a question
func HandleUpvote(db *sql.DB, questionID float64, sessionID string) error {
	user_Id, err := getUserIDUsingSessionID(sessionID, db)
	if err != nil {
		return err
	}
	if !CheckIfUserAlreadyVotedForThisQuestion(db, questionID, float64(user_Id)) {
		fmt.Println("Inserting user in vote_question")
		InsertUserInVoteQuestionDb(db, questionID, float64(user_Id), false, false)
	}
	// Increment upvote and potentially decrement downvote if already voted
	if !CheckIfUserUpvoted(db, questionID, float64(user_Id)) && !CheckIfUserDownvoted(db, questionID, float64(user_Id)) {
		fmt.Println("User upvoted")
		UpdateUserInVoteQuestionDb(db, questionID, float64(user_Id), true, false)
		query := `UPDATE question SET upvotes = upvotes + 1 WHERE id_question = $1`
		_, err = db.Exec(query, questionID)
		if err != nil {
			return err
		}
		return nil
	} else if CheckIfUserUpvoted(db, questionID, float64(user_Id)) && !CheckIfUserDownvoted(db, questionID, float64(user_Id)) {
		fmt.Println("User did not upvote")
		UpdateUserInVoteQuestionDb(db, questionID, float64(user_Id), false, false)
		query := `UPDATE question SET upvotes = upvotes - 1 WHERE id_question = $1`
		_, err = db.Exec(query, questionID)
		if err != nil {
			return err
		}
		return nil
	} else if !CheckIfUserUpvoted(db, questionID, float64(user_Id)) && CheckIfUserDownvoted(db, questionID, float64(user_Id)) {
		fmt.Println("User upvoted")
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
		fmt.Println("User did not upvote")
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
	user_Id, err := getUserIDUsingSessionID(sessionID, db)
	if err != nil {
		return err
	}
	if !CheckIfUserAlreadyVotedForThisQuestion(db, questionID, float64(user_Id)) {
		fmt.Println("Inserting user in vote_question")
		InsertUserInVoteQuestionDb(db, questionID, float64(user_Id), false, false)
	}
	// Increment downvote and potentially decrement upvote if already voted
	if !CheckIfUserDownvoted(db, questionID, float64(user_Id)) && !CheckIfUserUpvoted(db, questionID, float64(user_Id)) {
		fmt.Println("User downvoted")
		UpdateUserInVoteQuestionDb(db, questionID, float64(user_Id), false, true)
		query := `UPDATE question SET downvotes = downvotes + 1 WHERE id_question = $1`
		_, err = db.Exec(query, questionID)
		if err != nil {
			return err
		}
		return nil
	} else if CheckIfUserDownvoted(db, questionID, float64(user_Id)) && !CheckIfUserUpvoted(db, questionID, float64(user_Id)) {
		fmt.Println("User did not downvote")
		UpdateUserInVoteQuestionDb(db, questionID, float64(user_Id), false, false)
		query := `UPDATE question SET downvotes = downvotes - 1 WHERE id_question = $1`
		_, err = db.Exec(query, questionID)
		if err != nil {
			return err
		}
		return nil
	} else if !CheckIfUserDownvoted(db, questionID, float64(user_Id)) && CheckIfUserUpvoted(db, questionID, float64(user_Id)) {
		fmt.Println("User downvoted")
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
		fmt.Println("User did not downvote")
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
