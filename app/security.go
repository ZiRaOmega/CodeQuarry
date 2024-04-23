package app

import (
	"log"
	"net/http"
	"os"
)

type LogLevel int

const (
	DebugLevel LogLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
)

// https://bruinsslot.jp/post/go-secure-webserver/index.html
func AddSecurityHeaders(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubdomains")
		w.Header().Add("Content-Security-Policy", "default-src 'self'")
		w.Header().Add("X-XSS-Protection", "1; mode=block")
		w.Header().Add("X-Frame-Options", "DENY")
		w.Header().Add("Referrer-Policy", "strict-origin-when-cross-origin")
		w.Header().Add("X-Content-Type-Options", "nosniff")
		//w.Header().Add("Content-Type", "text/plain")

		next.ServeHTTP(w, r)
	}
}

// func LogAudit(message string) {
// 	file, err := os.OpenFile("audit.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 	if err != nil {
// 		log.Fatalf("Failed to open audit.log: %s", err)
// 	}
// 	defer file.Close()

// 	logger := log.New(file, "LOG: ", log.LstdFlags)
// 	logger.Println(message)
// }

func Log(level LogLevel, message string) {
	file, err := os.OpenFile("audit.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open audit.log: %s", err)
	}
	defer file.Close()

	var prefix string
	switch level {
	case DebugLevel:
		prefix = "DEBUG: "
	case InfoLevel:
		prefix = "INFO: "
	case WarnLevel:
		prefix = "WARN: "
	case ErrorLevel:
		prefix = "ERROR: "
	}

	logger := log.New(file, prefix, log.LstdFlags)
	logger.Println(message)
}
