package main_test

import (
	"crypto/tls"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
	"strings"
	"testing"

	"codequarry/app"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

func TestCssHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/global_style/global.css", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.CssHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "text/css"
	contentType := rr.Header().Get("Content-Type")
	if !strings.HasPrefix(contentType, expected) {
		t.Errorf("handler returned wrong content type: got %v want %v", contentType, expected)
	}
}

func TestFaviconHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/favicon.ico", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.FaviconHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "image/"
	contentType := rr.Header().Get("Content-Type")
	if !strings.HasPrefix(contentType, expected) {
		t.Errorf("handler returned wrong content type: got %v want %v", contentType, expected)
	}
}

func TestSendComponentAuth(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		t.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbType := os.Getenv("DB_TYPE")

	dsn := dbType + "://" + dbUser + ":" + dbPassword + "@" + dbHost + ":" + dbPort + "/" + dbName + "?sslmode=disable"
	db := app.InitDB(dsn)
	defer db.Close()

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := app.SendComponent("auth", db)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "text/html"
	contentType := rr.Header().Get("Content-Type")
	if !strings.HasPrefix(contentType, expected) {
		t.Errorf("handler returned wrong content type: got %v want %v", contentType, expected)
	}
}

func TestLoginHandler(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		t.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbType := os.Getenv("DB_TYPE")

	dsn := dbType + "://" + dbUser + ":" + dbPassword + "@" + dbHost + ":" + dbPort + "/" + dbName + "?sslmode=disable"
	db := app.InitDB(dsn)
	defer db.Close()

	req, err := http.NewRequest("POST", "/login", strings.NewReader("username=test&password=test"))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.LoginHandler(db))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestRegisterHandler(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		t.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbType := os.Getenv("DB_TYPE")

	dsn := dbType + "://" + dbUser + ":" + dbPassword + "@" + dbHost + ":" + dbPort + "/" + dbName + "?sslmode=disable"
	db := app.InitDB(dsn)
	defer db.Close()

	req, err := http.NewRequest("POST", "/register", strings.NewReader("username=test&password=test&email=test@example.com"))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.RegisterHandler(db))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestLogoutHandler(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		t.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbType := os.Getenv("DB_TYPE")

	dsn := dbType + "://" + dbUser + ":" + dbPassword + "@" + dbHost + ":" + dbPort + "/" + dbName + "?sslmode=disable"
	db := app.InitDB(dsn)
	defer db.Close()

	req, err := http.NewRequest("POST", "/logout", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Set a session cookie
	cookie := &http.Cookie{
		Name:  "session",
		Value: "testsessionid",
	}
	req.AddCookie(cookie)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.LogoutHandler(db))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusSeeOther {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusSeeOther)
	}
}

func TestWebsocketHandler(t *testing.T) {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		t.Fatal("Error loading .env file")
	}

	// Retrieve database connection details from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbType := os.Getenv("DB_TYPE")

	// Form the Data Source Name (DSN) string
	dsn := dbType + "://" + dbUser + ":" + dbPassword + "@" + dbHost + ":" + dbPort + "/" + dbName + "?sslmode=disable"
	// Initialize the database connection
	db := app.InitDB(dsn)
	defer db.Close()

	// Start a new HTTPS server with the WebSocket handler
	server := httptest.NewTLSServer(http.HandlerFunc(app.WebsocketHandler(db)))
	defer server.Close()

	// Add the server URL to the allowed origins
	app.AllowedOrigins = append(app.AllowedOrigins, server.URL)

	// Construct the WebSocket URL with wss (WebSocket Secure)
	u := url.URL{Scheme: "wss", Host: server.Listener.Addr().String(), Path: "/ws"}
	headers := http.Header{"Origin": {server.URL}}

	// Dial the WebSocket server with TLS configuration to skip certificate verification
	dialer := websocket.Dialer{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	ws, _, err := dialer.Dial(u.String(), headers)
	if err != nil {
		t.Fatalf("failed to dial WebSocket: %v", err)
	}
	defer ws.Close()

	// Send a test message to the WebSocket server
	err = ws.WriteJSON(app.WSMessage{Type: "session"})
	if err != nil {
		t.Fatalf("failed to write message: %v", err)
	}

	// Read the response message from the WebSocket server
	wsmessage := app.WSMessage{}
	err = ws.ReadJSON(&wsmessage)
	if err != nil {
		t.Fatalf("failed to read message: %v", err)
	}

	// Define the expected response
	expected := app.WSMessage{Type: "session", Content: "expired"} // Adjust this according to your handler's response
	if !reflect.DeepEqual(expected, wsmessage) {
		t.Errorf("unexpected response: got %v want %v", wsmessage, expected)
	}
}
