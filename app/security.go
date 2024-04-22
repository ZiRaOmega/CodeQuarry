package app

import "net/http"

// https://bruinsslot.jp/post/go-secure-webserver/index.html
func AddSecurityHeaders(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubdomains")
        w.Header().Add("Content-Security-Policy", "default-src 'self'")
        w.Header().Add("X-XSS-Protection", "1; mode=block")
        w.Header().Add("X-Frame-Options", "DENY")
        w.Header().Add("Referrer-Policy", "strict-origin-when-cross-origin")
        w.Header().Add("X-Content-Type-Options", "nosniff")
        w.Header().Add("Content-Type", "text/plain")

		next.ServeHTTP(w, r)
	}
}
