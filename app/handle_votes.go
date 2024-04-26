package app

import "database/sql"

// HandleUpvote updates the upvote count for a question
func HandleUpvote(db *sql.DB, questionID float64) error {
	// Increment upvote and potentially decrement downvote if already voted
	query := `UPDATE question SET upvotes = upvotes + 1 WHERE id_question = $1`
	_, err := db.Exec(query, questionID)
	if err != nil {
		return err
	}
	return nil
}

// HandleDownvote updates the downvote count for a question
func HandleDownvote(db *sql.DB, questionID float64) error {
	// Increment downvote and potentially decrement upvote if already voted
	query := `UPDATE question SET downvotes = downvotes + 1 WHERE id_question = $1`
	_, err := db.Exec(query, questionID)
	if err != nil {
		return err
	}
	return nil
}

// ToggleVote removes a previous vote (up or down) when clicked again
func ToggleVote(db *sql.DB, questionID float64, voteType string) error {
	if voteType == "upvote" {
		query := `UPDATE question SET upvotes = upvotes - 1 WHERE id_question = $1 AND upvotes > 0`
		_, err := db.Exec(query, questionID)
		if err != nil {
			return err
		}
		return nil
	} else {
		query := `UPDATE question SET downvotes = downvotes - 1 WHERE id_question = $1 AND downvotes > 0`
		_, err := db.Exec(query, questionID)
		if err != nil {
			return err
		}
		return nil
	}

}
