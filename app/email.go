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
func UpdateUserPassword(db *sql.DB, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	// Update the password in the database
	_, err = db.Exec("UPDATE users SET password = $1 WHERE email = $2", hashedPassword, email)
	return err
}

func InsertVerifToken(db *sql.DB, email, token string) error {
	_, err := db.Exec("INSERT INTO VerifyEMail(email,token) VALUES($1,$2)", email, token)
	return err
}

func isValidToken(db *sql.DB, token string) bool {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM VerifyEmail WHERE token = $1 AND validated = FALSE", token).Scan(&count)

	if err != nil {
		return false
	}
	if count == 1 {
		err := UpdateValidated(db, token)
		if err != nil {
			return false
		}
		return true
	}
	return false
}
func UpdateValidated(db *sql.DB, token string) error {
	_, err := db.Exec("UPDATE VerifyEmail SET validated=true WHERE token=$1", token)
	return err
}

func isEmailVerified(db *sql.DB, email string) bool {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM VerifyEmail WHERE email = $1 AND validated = TRUE", email).Scan(&count)
	if err != nil {
		return false
	}
	if count == 1 {
		return true
	}
	return false
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
		http.Error(w, "Password reset successful. Check your email for the new password", http.StatusOK)
	}
}
