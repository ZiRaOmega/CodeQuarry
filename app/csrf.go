package app

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("Z9H6byZx3dvnA5H+GZNJsaWshO1iamGbhTsM4C4eFxI="))

func InitStoreCSRFToken() {
	// Configure the store with secure options
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600, // 1 hour
		HttpOnly: true, // Make sure cookie is not accessible via JavaScript
		Secure:   true, // Send cookie only over HTTPS
		SameSite: http.SameSiteStrictMode,
	}
}

const (
	csrfTokenLength = 32
	csrfSessionKey  = "csrf_token"
)

func generateCSRFToken() (string, error) {
	token := make([]byte, csrfTokenLength)
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(token), nil
}

func SetCSRFToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Generate CSRF token
		token, err := generateCSRFToken()
		if err != nil {
			http.Error(w, "Unable to generate CSRF token", http.StatusInternalServerError)
			return
		}

		// Store CSRF token in session
		session, _ := store.Get(r, "csrfToken")
		session.Values[csrfSessionKey] = token
		session.Save(r, w)

		next.ServeHTTP(w, r)
	})
}
func VerifyCSRFToken(w http.ResponseWriter, r *http.Request) bool {
	// Get the CSRF token from the form
	formToken := r.FormValue("csrfToken")

	// Retrieve the CSRF token from session
	session, _ := store.Get(r, "csrfToken")
	sessionToken, ok := session.Values[csrfSessionKey].(string)
	if !ok {
		fmt.Println("CSRF token not found in session")
		return false
	}

	// Compare the tokens
	fmt.Println("Form Token:", formToken)
	fmt.Println("Session Token:", sessionToken)

	//Delete the token from store
	delete(session.Values, csrfSessionKey)
	session.Save(r, w)
	return formToken == sessionToken
}
