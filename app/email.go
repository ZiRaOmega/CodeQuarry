package app

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/smtp"

	"golang.org/x/crypto/bcrypt"
)

const (
	// SMTP server configuration
	SmtpHost = "smtp.gmail.com"
	SmtpPort = "587"
	SmtpUser = "psvforum01@gmail.com"
	SmtpPass = "ufzuqzwwerfsyscc"
)

func SendVerificationEmail(db *sql.DB, email, token string) {

	auth := smtp.PlainAuth("", SmtpUser, SmtpPass, SmtpHost)

	subject := "Subject: Verify Your Email\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := "<html><body><p>Please verify your email by clicking on the link below:</p>" +
		"<a href=\"https://localhost/verify?token=" + token + "\">Verify Email</a></body></html>"

	msg := []byte("From: " + SmtpUser + "\n" +
		"To: " + email + "\n" +
		subject + mime + body)
	err := InsertVerifToken(db, email, token)
	if err != nil {
		panic(err)
	}
	err = smtp.SendMail(SmtpHost+":"+SmtpPort, auth, SmtpUser, []string{email}, msg)
	if err != nil {
		panic(err) // handle the error appropriately in production code
	}
}

func GenerateTokenVerificationEmail() string {
	token := make([]byte, 32) // 256 bits are usually sufficient
	_, err := rand.Read(token)
	if err != nil {
		return "" // handle the error appropriately in production code
	}
	return base64.URLEncoding.EncodeToString(token)
}

func ResetPassword(db *sql.DB, email string) string {
	tempPassword := make([]byte, 12) // Generate a temporary 12-byte password
	_, err := rand.Read(tempPassword)
	if err != nil {
		return "" // handle the error appropriately in production code
	}
	tempPasswordStr := base64.URLEncoding.EncodeToString(tempPassword)[:12] // Shorten to 12 characters

	// Assuming you have a function to update the password in your user database
	err = UpdateUserPassword(db, email, tempPasswordStr)
	if err != nil {
		return "" // handle the error appropriately in production code
	}

	// Here, you should also send the new password to the user's email, which we will not demonstrate here
	return tempPasswordStr
}

func SendResetPasswordEmail(email, password string) {

	auth := smtp.PlainAuth("", SmtpUser, SmtpPass, SmtpHost)

	subject := "Subject: Reset Password\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := fmt.Sprintf(`<html><body><p>Your new password is: %s</p></body></html>`, password)

	msg := []byte("From: " + SmtpUser + "\n" +
		"To: " + email + "\n" +
		subject + mime + body)

	err := smtp.SendMail(SmtpHost+":"+SmtpPort, auth, SmtpUser, []string{email}, msg)
	if err != nil {
		panic(err) // handle the error appropriately in production code
	}
}

// UpdateUserPassword hashes a new password and updates it in the database using a prepared statement.
func UpdateUserPassword(db *sql.DB, email, password string) error {
	// Generate a hash of the password to store in the database.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Prepare the SQL statement for execution.
	stmt, err := db.Prepare("UPDATE users SET password = $1 WHERE email = $2")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the prepared statement with the hashed password and email as parameters.
	_, err = stmt.Exec(hashedPassword, email)
	return err
}

// InsertVerifToken inserts a verification token into the database using a prepared statement.
func InsertVerifToken(db *sql.DB, email, token string) error {
	// Prepare the SQL statement for execution.
	stmt, err := db.Prepare("INSERT INTO VerifyEMail(email, token) VALUES($1, $2)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the prepared statement with the email and token as parameters.
	_, err = stmt.Exec(email, token)
	return err
}

// isValidToken checks the validity of a token using a prepared statement.
func isValidToken(db *sql.DB, token string) bool {
	// Prepare the SQL query to check the token's validity.
	stmt, err := db.Prepare("SELECT COUNT(*) FROM VerifyEmail WHERE token = $1 AND validated = FALSE")
	if err != nil {
		return false
	}
	defer stmt.Close()

	var count int
	// Execute the query with the token as a parameter and scan the result.
	err = stmt.QueryRow(token).Scan(&count)
	if err != nil {
		return false
	}

	if count == 1 {
		// Update the validation status if the token is valid.
		return UpdateValidated(db, token) == nil
	}
	return false
}

// UpdateValidated updates the validated status of a token using a prepared statement.
func UpdateValidated(db *sql.DB, token string) error {
	stmt, err := db.Prepare("UPDATE VerifyEmail SET validated = TRUE WHERE token = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(token)
	return err
}

// isEmailVerified checks if an email is verified using a prepared statement.
func isEmailVerified(db *sql.DB, email string) bool {
	stmt, err := db.Prepare("SELECT COUNT(*) FROM VerifyEmail WHERE email = $1 AND validated = TRUE")
	if err != nil {
		return false
	}
	defer stmt.Close()

	var count int
	err = stmt.QueryRow(email).Scan(&count)
	if err != nil {
		return false
	}
	return count == 1
}

func ForgotPasswordHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		email := r.FormValue("email")
		if email == "" {
			http.Error(w, "Email is required", http.StatusBadRequest)
			return
		}
		if !isEmailVerified(db, email) {
			http.Error(w, "Email is not verified", http.StatusBadRequest)
			return
		}
		password := ResetPassword(db, email)
		SendResetPasswordEmail(email, password)
		//http.Error(w, "Password reset successful. Check your email for the new password", http.StatusOK)
		http.Redirect(w, r, "/", http.StatusOK)
	}
}
