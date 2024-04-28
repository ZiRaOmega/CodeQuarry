package app

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"codequarry/app/utils"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func ProfileHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get cookie
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
			fmt.Println(err.Error())
			http.Error(w, "Error getting user info", http.StatusInternalServerError)
			return
		}
		ParseAndExecuteTemplate("profile", user, w)
	}
}

func (U *User) FormatBirthDate() string {
	return U.BirthDate.Time.Format("2006-01-02")
}

func (U *User) FormatSchoolYear() string {
	return U.SchoolYear.Time.Format("2006-01-02")
}

// Define User structure based on your database schema
type User struct {
	ID                 int
	LastName           string
	FirstName          string
	Username           string
	Email              string
	Password           string
	Avatar             sql.NullString
	BirthDate          sql.NullTime // Adjusted for possible NULL values
	Birth_Date_Format  string
	Bio                sql.NullString
	Website            sql.NullString
	GitHub             sql.NullString
	XP                 sql.NullInt64
	Rank               sql.NullString
	SchoolYear         sql.NullTime // Adjusted for possible NULL values
	School_Year_Format string
	CreationDate       sql.NullTime // Adjusted for possible NULL values
	UpdateDate         sql.NullTime // Adjusted for possible NULL values
	DeletingDate       sql.NullTime // Adjusted for possible NULL values
	My_Post            []Question
	VotedQuestion      []Question
}

// GetUser fetches user details from the database based on the session ID
func GetUser(session_id string, db *sql.DB) (User, error) {
	var user User
	if db == nil {
		return user, errors.New("database connection is nil")
	}

	// Prepare the SQL statement for fetching user_id from session_id
	stmt, err := db.Prepare("SELECT user_id FROM Sessions WHERE uuid = $1")
	if err != nil {
		return user, err // Return immediately if there's an error preparing the statement
	}
	defer stmt.Close() // Ensure statement is always closed after use

	// Execute the statement and scan the result
	var user_id int
	if err = stmt.QueryRow(session_id).Scan(&user_id); err != nil {
		return user, err
	}

	// Prepare the SQL statement for fetching user details from user_id
	stmt, err = db.Prepare("SELECT id_student, lastname, firstname, username, email, password, avatar, birth_date, bio, website, github, xp, rang_rank_, school_year, creation_date, update_date, deleting_date FROM users WHERE id_student = $1")
	if err != nil {
		return user, err // Handle preparation errors
	}
	defer stmt.Close() // Ensure statement is always closed after use

	// Execute the statement and scan the result into the User struct
	if err = stmt.QueryRow(user_id).Scan(&user.ID, &user.LastName, &user.FirstName, &user.Username, &user.Email, &user.Password, &user.Avatar, &user.BirthDate, &user.Bio, &user.Website, &user.GitHub, &user.XP, &user.Rank, &user.SchoolYear, &user.CreationDate, &user.UpdateDate, &user.DeletingDate); err != nil {
		return user, err
	}
	// FormatDates formats CreationDate and
	// UpdateDate to a human-readable format dd/mm/yyyy
	user.Birth_Date_Format = user.FormatBirthDate()
	user.School_Year_Format = user.FormatSchoolYear()
	Posts, err := FetchQuestionsByUserID(db, user.ID)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(Posts)
	user.My_Post = Posts
	//user.VotedQuestion, err = FetchVotedQuestions(db, user.ID)
	return user, nil // Return the populated user object
}

type QuestionVote struct {
	Upvote   bool
	Downvote bool
	Q        Question
}

func FetchVotedQuestions(db *sql.DB, userID int) ([]QuestionVote, error) {
	var questions []QuestionVote
	/*Vote_question(
	id_student INT,
	id_question INT,
	upvote_q BOOLEAN NOT NULL,
	downvote_q BOOLEAN NOT NULL,
	PRIMARY KEY(id_student, id_question),
	FOREIGN KEY(id_student) REFERENCES users(id_student),
	FOREIGN KEY(id_question) REFERENCES Question(id_question)*/
	//Fetch using Vote_question
	rows, err := db.Query(`SELECT q.id_question, s.title AS subject_title, q.title, q.content, q.creation_date, u.username, q.upvotes, q.downvotes, v.upvote_q, v.downvote_q
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
		var q Question
		var qq QuestionVote
		if err := rows.Scan(&q.Id, &q.SubjectTitle, &q.Title, &q.Content, &q.CreationDate, &q.Creator, &q.Upvotes, &q.Downvotes, &qq.Upvote, &qq.Downvote); err != nil {
			continue
		}
		qq.Q = q
		questions = append(questions, qq)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	fmt.Println(questions)
	return questions, nil
}

// UpdateProfileHandler is an HTTP request handler that updates the user profile.
// It takes a database connection as input and returns an http.HandlerFunc.
// The handler checks if the request method is POST and if the session is valid.
// It then retrieves the user ID from the session and compares it with the ID in the request form.
// If the IDs match, it updates the user profile with the provided information.
// If the password is empty, it updates the profile without changing the password.
// If the password is not empty, it hashes the password before updating the profile.
// Finally, it redirects the user to the profile page.
func UpdateProfileHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		// Check if the session is valid
		cookie, err := r.Cookie("session")
		if err != nil {
			http.Error(w, "Error getting session cookie", http.StatusInternalServerError)
			return
		}
		session_id := cookie.Value
		if !isValidSession(session_id, db) {
			http.Redirect(w, r, "/auth", http.StatusSeeOther)
			return
		}

		user := User{}

		user.ID, err = getUserIDUsingSessionID(session_id, db)
		if err != nil {
			fmt.Println(err.Error())
		}
		if strconv.Itoa(user.ID) != r.PostFormValue("id_student") {
			http.Error(w, "Invalid user", http.StatusForbidden)
			return
		}
		user.LastName = r.PostFormValue("lastname")
		user.FirstName = r.PostFormValue("firstname")
		user.Username = r.PostFormValue("username")
		user.Email = r.PostFormValue("email")
		user.Password = r.PostFormValue("password")
		user.Avatar = sql.NullString{String: r.PostFormValue("avatar"), Valid: true}
		filename := ""

		filename, err = FileUpload(r)
		if err != nil {
			fmt.Println(err.Error())
		}
		if filename == "" {
			filename, err = getAvatar(db, user.ID)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
		fmt.Println(filename)
		user.Avatar = sql.NullString{String: filename, Valid: true}
		birthDateStr := r.PostFormValue("birth_date")
		birthDate, err := time.Parse("2006-01-02", birthDateStr)
		if err != nil {
			fmt.Println(err.Error())
		}
		user.BirthDate = sql.NullTime{Time: birthDate, Valid: true}
		user.Bio = sql.NullString{String: r.PostFormValue("bio"), Valid: true}
		user.Website = sql.NullString{String: r.PostFormValue("website"), Valid: true}
		user.GitHub = sql.NullString{String: r.PostFormValue("github"), Valid: true}
		user.SchoolYear = sql.NullTime{Time: time.Now(), Valid: true}
		if utils.ContainsSQLi(user.LastName) || utils.ContainsSQLi(user.FirstName) || utils.ContainsSQLi(user.Username) || utils.ContainsSQLi(user.Email) || utils.ContainsSQLi(user.Password) || utils.ContainsSQLi(user.Bio.String) || utils.ContainsSQLi(user.Website.String) || utils.ContainsSQLi(user.GitHub.String) {
			http.Error(w, "Invalid characters", http.StatusForbidden)
			return
		} else if utils.ContainsXSS(user.LastName) || utils.ContainsXSS(user.FirstName) || utils.ContainsXSS(user.Username) || utils.ContainsXSS(user.Email) || utils.ContainsXSS(user.Password) || utils.ContainsXSS(user.Bio.String) || utils.ContainsXSS(user.Website.String) || utils.ContainsXSS(user.GitHub.String) {
			http.Error(w, "Invalid characters", http.StatusForbidden)
			return
		}
		// if Password is empty, don't update it
		if user.Password == "" {
			// Prepare
			stmt, err := db.Prepare("UPDATE users SET lastname = $1, firstname = $2, username = $3, email = $4, birth_date = $5, avatar = $11, bio = $6, website = $7, github = $8, school_year = $9 WHERE id_student = $10")
			if err != nil {
				fmt.Println(err.Error())
			}
			defer stmt.Close()
			// Execute
			_, err = stmt.Exec(user.LastName, user.FirstName, user.Username, user.Email, user.BirthDate, user.Bio, user.Website, user.GitHub, user.SchoolYear, user.ID, user.Avatar.String)
			if err != nil {
				fmt.Println(err.Error())
			}
		} else {
			// Hash password
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
			if err != nil {
				fmt.Println(err.Error())
			}
			// Prepare
			stmt, err := db.Prepare("UPDATE users SET lastname = $1, firstname = $2, username = $3, email = $4, password = $5, birth_date = $6, avatar = $12, bio = $7, website = $8, github = $9, school_year = $10 WHERE id_student = $11")
			if err != nil {
				fmt.Println(err.Error())
			}
			defer stmt.Close()
			// Execute
			_, err = stmt.Exec(user.LastName, user.FirstName, user.Username, user.Email, hashedPassword, user.BirthDate, user.Bio, user.Website, user.GitHub, user.SchoolYear, user.ID, user.Avatar.String)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
		http.Redirect(w, r, "/profile", http.StatusSeeOther)
	}
}

// FileUpload handles the uploading of a file from a HTTP request.
// It expects a multipart form with a file field named "avatar".
// The function saves the uploaded file to the "./public/img/" directory with a generated name,
// and returns the path to the saved file on success.
// If there is an error during the file upload process, it returns an error.
func FileUpload(r *http.Request) (string, error) {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("avatar")
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	//if size is null
	if handler.Size == 0 {

		return "", nil
	}
	filename := strings.Split(handler.Filename, ".")
	extention := filename[len(filename)-1]
	uuid := uuid.NewV4()
	genname := uuid.String()[0:5] + "." + extention
	if handler.Filename[len(handler.Filename)-4:] != ".jpg" && handler.Filename[len(handler.Filename)-5:] != ".jpeg" && handler.Filename[len(handler.Filename)-4:] != ".png" && handler.Filename[len(handler.Filename)-4:] != ".gif" {
		return "", errors.New("wrong file type")
	}

	defer file.Close()
	_, err = os.Create("./public/img/" + genname)
	if err != nil {
		return "", err
	}
	f, err := os.OpenFile("./public/img/"+genname, os.O_WRONLY|os.O_CREATE, 0o666)
	if err != nil {
		return "", err
	}
	io.Copy(f, file)
	return "/img/" + genname, nil
}

// getAvatar
func getAvatar(db *sql.DB, user_id int) (string, error) {
	var avatar string
	stmt, err := db.Prepare("SELECT avatar FROM users WHERE id_student = $1")
	if err != nil {
		return "", err
	}
	defer stmt.Close()
	err = stmt.QueryRow(user_id).Scan(&avatar)
	if err != nil {
		return "", err
	}
	return avatar, nil
}
