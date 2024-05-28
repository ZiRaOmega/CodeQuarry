package app

import (
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/gorilla/sessions"
)

var store *sessions.CookieStore

func InitStoreCSRFToken() func(http.Handler) http.Handler {
	// Replace with your own authentication key (32 bytes)
	csrfKey := []byte("32-byte-long-auth-key")

	// Initialize the store with secure options
	store = sessions.NewCookieStore(
		[]byte("Z9H6byZx3dvnA5H+GZNJsaWshO1iamGbhTsM4C4eFxI="),
	)
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600, // 1 hour
		HttpOnly: true, // Make sure cookie is not accessible via JavaScript
		Secure:   true, // Send cookie only over HTTPS
		SameSite: http.SameSiteStrictMode,
	}

	// Initialize CSRF protection
	CSRF := csrf.Protect(csrfKey, csrf.Secure(false))
	return CSRF

}

const (
	csrfTokenLength = 32
	csrfSessionKey  = "csrf_token"
)
