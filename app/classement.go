package app

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

/*
func ClassementHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		if err != nil {
			http.Error(w, "Error getting session cookie", http.StatusInternalServerError)
			http.Redirect(w, r, "/auth", http.StatusSeeOther)
			return
		}
		session_id := cookie.Value
		// Get user info from user_id
		user, err := GetUser(session_id, db)
		if err != nil {

			http.Error(w, "Error getting user info", http.StatusInternalServerError)
			return
		}
		ParseAndExecuteTemplate("classement", user, w)
	}
} */

func GetUsersInfo(db *sql.DB) ([]User, error) {
	var users []User
	query := `
	SELECT id_student, username,lastname, firstname, xp
	FROM users
	ORDER BY xp DESC`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username, &user.LastName, &user.FirstName, &user.XP); err != nil {
			return nil, err
		}
		user.Rank.String, err = SetRankByXp(user)
		if err != nil {
			return nil, err
		}
		Posts, err := FetchQuestionsByUserID(db, user.ID)
		if err != nil {

		}

		user.My_Post = Posts
		users = append(users, user)
	}
	defer rows.Close()
	return users, nil
}

func SendUsersInfoJson(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := GetUsersInfo(db)
		if err != nil {

			http.Error(w, "Error getting users info", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}
