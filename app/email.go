package app

/*
import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fmt"
	"net/smtp"
)

func SendVerificationEmail(email, token string) {
	from := "no-reply@yourdomain.com"
	pass := "your-email-password"
	smtpHost := "smtp.yourdomain.com"
	smtpPort := "587"
	auth := smtp.PlainAuth("", from, pass, smtpHost)

	subject := "Subject: Verify Your Email\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := "<html><body><p>Please verify your email by clicking on the link below:</p>" +
		"<a href=\"https://yourdomain.com/verify?token=" + token + "\">Verify Email</a></body></html>"

	msg := []byte("From: " + from + "\n" +
		"To: " + email + "\n" +
		subject + mime + body)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{email}, msg)
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

func ResetPassword(db *sql.DB,email string) string {
	tempPassword := make([]byte, 12) // Generate a temporary 12-byte password
	_, err := rand.Read(tempPassword)
	if err != nil {
		return "" // handle the error appropriately in production code
	}
	tempPasswordStr := base64.URLEncoding.EncodeToString(tempPassword)[:12] // Shorten to 12 characters

	// Assuming you have a function to update the password in your user database
	// UpdateUserPassword(email, tempPasswordStr)

	// Here, you should also send the new password to the user's email, which we will not demonstrate here
	return tempPasswordStr
}

func SendResetPasswordEmail(email, password string) {
	from := "no-reply@example.com"
	pass := "your-email-password"
	smtpHost := "smtp.example.com"
	smtpPort := "587"
	auth := smtp.PlainAuth("", from, pass, smtpHost)

	subject := "Subject: Reset Password\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := fmt.Sprintf(`<html><body><p>Your new password is: %s</p></body></html>`, password)

	msg := []byte("From: " + from + "\n" +
		"To: " + email + "\n" +
		subject + mime + body)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{email}, msg)
	if err != nil {
		panic(err) // handle the error appropriately in production code
	}
}
func UpdateUserPassword(db *sql.DB, email, password string) error {
	// Update the password in the database
	_, err := db.Exec("UPDATE users SET password = ? WHERE email = ?", password, email)
	return err
}
*/
