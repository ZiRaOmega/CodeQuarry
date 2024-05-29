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

// TestCssHandler tests the CssHandler function from the app package.
// It sends a GET request to the /global_style/global.css endpoint
// and checks if the response status code is 200 (OK) and if the
// Content-Type header starts with "text/css".
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

// TestFaviconHandler tests the FaviconHandler function from the app package.
// It sends a GET request to the /favicon.ico endpoint and checks if the
// response status code is 200 (OK) and if the Content-Type header starts with "image/".
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

// TestSendComponentAuth tests the SendComponent function from the app package
// for the "auth" component. It initializes the database connection using
// environment variables, sends a GET request to the root endpoint ("/"),
// and checks if the response status code is 200 (OK) and if the Content-Type
// header starts with "text/html".
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

// TestLoginHandler tests the LoginHandler function from the app package.
// It initializes the database connection using environment variables,
// sends a POST request to the /login endpoint with form data, and checks if
// the response status code is 200 (OK).
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

// TestRegisterHandler tests the RegisterHandler function from the app package.
// It initializes the database connection using environment variables,
// sends a POST request to the /register endpoint with form data, and checks if
// the response status code is 200 (OK).
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

// TestLogoutHandler tests the LogoutHandler function from the app package.
// It initializes the database connection using environment variables,
// sends a POST request to the /logout endpoint with a session cookie,
// and checks if the response status code is 303 (See Other).
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

// TestWebsocketHandler tests the WebsocketHandler function from the app package.
// It initializes the database connection using environment variables,
// starts a new HTTPS server with the WebSocket handler, connects to the WebSocket
// server, sends a test message, and checks if the response message is as expected.
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
