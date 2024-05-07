package app

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"
)

type Panel struct {
	Questions []Question
	Users     []User
	Subjects  []Subject
	Rank      int
}

func PanelAdminHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session_id, err := r.Cookie("session")
		if err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		user_id, err := getUserIDUsingSessionID(session_id.Value, db)
		if err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		fmt.Println(user_id)
		data := Panel{}
		rank := FetchRankByUserID(db, user_id)
		data.Rank = rank
		data.Subjects = FetchSubjects(db)
		data.Questions = FetchQuestions(db)
		data.Users, err = FetchUsers(db)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(rank, "RANK")
		fmt.Println(data.Questions)
		if rank == 2 || rank == 1 {
			err := ParseAndExecuteTemplate("panel", data, w)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			//Admin/Moderator
		} else if rank == 0 {
			//User
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}

func FetchRankByUserID(db *sql.DB, user_id int) int {
	var rank int
	stmt, err := db.Prepare("SELECT rang_rank_ FROM users WHERE id_student = $1")
	if err != nil {
		return -1
	}
	defer stmt.Close()
	if err := stmt.QueryRow(user_id).Scan(&rank); err != nil {
		return -1
	}
	return rank
}

func FetchSubjects(db *sql.DB) []Subject {
	/*Subject(
	id_subject SERIAL NOT NULL,
	title VARCHAR(50) NOT NULL,
	description VARCHAR(500) NOT NULL,
	creation_date DATE NOT NULL,
	update_date DATE NOT NULL,
	PRIMARY KEY(id_subject),
	UNIQUE(title)*/
	var subjects []Subject
	rows, err := db.Query("SELECT id_subject, title, description, creation_date, update_date FROM Subject")
	if err != nil {
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var title, description string
		var creation_date, update_date time.Time
		if err := rows.Scan(&id, &title, &description, &creation_date, &update_date); err != nil {
			continue
		}
		subjects = append(subjects, Subject{
			Id:           id,
			Title:        title,
			Description:  description,
			CreationDate: creation_date,
			UpdateDate:   update_date,
		})
	}
	return subjects
}

func FetchQuestions(db *sql.DB) []Question {
	/*Question(
	id_question SERIAL NOT NULL,
	id_subject INT NOT NULL,
	id_user INT NOT NULL,
	title VARCHAR(50) NOT NULL,
	description VARCHAR(500) NOT NULL,
	creation_date DATE NOT NULL,
	update_date DATE NOT NULL,
	PRIMARY KEY(id_question),
	FOREIGN KEY(id_subject) REFERENCES Subject(id_subject),
	FOREIGN KEY(id_user) REFERENCES users(id_student)*/
	var questions []Question
	rows, err := db.Query("SELECT id_question, id_subject, id_student, title, description, creation_date, update_date FROM Question")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var id, id_subject, id_user int
		var title, description string
		var creation_date, update_date time.Time
		if err := rows.Scan(&id, &id_subject, &id_user, &title, &description, &creation_date, &update_date); err != nil {
			fmt.Println(err.Error())

		}

		resp, err := FetchResponseByQuestion(db, id, 0)
		if err != nil {
			fmt.Println(err.Error())
		}
		questions = append(questions, Question{
			Id:           id,
			SubjectID:    id_subject,
			User_Id:      id_user,
			Title:        title,
			Description:  description,
			CreationDate: creation_date,
			UpdateDate:   update_date,
			Responses:    resp,
		})
	}
	return questions
}

func FetchUsers(db *sql.DB) ([]User, error) {
	var users []User
	query := `SELECT id_student, lastname, firstname, username, email, password, avatar, birth_date, bio, website, github, xp, rang_rank_, school_year, creation_date, update_date, deleting_date FROM users`
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query execution error: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		if err := rows.Scan(
			&user.ID,
			&user.LastName,
			&user.FirstName,
			&user.Username,
			&user.Email,
			&user.Password,
			&user.Avatar,
			&user.BirthDate,
			&user.Bio,
			&user.Website,
			&user.GitHub,
			&user.XP,
			&user.Rank_Panel,
			&user.SchoolYear,
			&user.CreationDate,
			&user.UpdateDate,
			&user.DeletingDate,
		); err != nil {
			return nil, fmt.Errorf("scan error: %v", err)
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows final error: %v", err)
	}
	return users, nil
}
