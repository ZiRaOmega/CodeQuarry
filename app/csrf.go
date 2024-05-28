package app

import (
	"net/http"
	"os"

	"github.com/gorilla/csrf"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

var store *sessions.CookieStore

// InitStoreCSRFToken initializes and returns a middleware function that adds CSRF protection to the HTTP handler.
// It loads the CSRF key and store key from environment variables, initializes the session store with secure options,
// and configures the CSRF protection. The returned middleware function can be used to wrap the HTTP handler.
func InitStoreCSRFToken() func(http.Handler) http.Handler {
	godotenv.Load()
	csrfKey := os.Getenv("CSRF_KEY")
	storeKey := os.Getenv("STORE_KEY")

	// Initialize the store with secure options
	store = sessions.NewCookieStore(
		[]byte(storeKey),
	)
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600, // 1 hour
		HttpOnly: true, // Make sure cookie is not accessible via JavaScript
		Secure:   true, // Send cookie only over HTTPS
		SameSite: http.SameSiteStrictMode,
	}

	// Initialize CSRF protection
	CSRF := csrf.Protect([]byte(csrfKey), csrf.Secure(false))
	return CSRF

}
