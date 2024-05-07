package app

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"
)

type Panel struct {
	Questions  []Question
	Users      []User
	Subjects   []Subject
	Rank_Panel sql.NullInt64
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
		data.Rank_Panel = sql.NullInt64{Int64: int64(rank), Valid: true}
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
	title VARCHAR(50) NOT NULL,
	description VARCHAR(1000) NOT NULL,
	content VARCHAR(10000) NOT NULL,
	upvotes INT,
	downvotes INT,
	creation_date DATE NOT NULL,
	update_date DATE NOT NULL,
	id_student INT NOT NULL,
	id_subject INT NOT NULL,
	PRIMARY KEY(id_question),
	UNIQUE(title),
	FOREIGN KEY(id_student) REFERENCES users(id_student),
	FOREIGN KEY(id_subject) REFERENCES Subject(id_subject)*/
	var questions []Question
	rows, err := db.Query("SELECT id_question, title, description, content, upvotes, downvotes, creation_date, update_date, id_student, id_subject FROM Question")
	if err != nil {
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		var id, upvotes, downvotes, student_id, subject_id int
		var title, description, content string
		var creation_date, update_date time.Time
		if err := rows.Scan(&id, &title, &description, &content, &upvotes, &downvotes, &creation_date, &update_date, &student_id, &subject_id); err != nil {
			continue
		}
		resp, err := FetchResponseByQuestion(db, id, 0)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(resp, "resp")
		questions = append(questions, Question{
			Id:           id,
			Title:        title,
			Description:  description,
			Content:      content,
			Upvotes:      upvotes,
			Downvotes:    downvotes,
			CreationDate: creation_date,
			UpdateDate:   update_date,
			User_Id:      student_id,
			SubjectID:    subject_id,
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
