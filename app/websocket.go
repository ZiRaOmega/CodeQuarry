package app

//Import websocket
import (
	"database/sql"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// WebsocketHandler is a handler function that upgrades the HTTP connection to a WebSocket connection
// and handles the communication between the client and the server using WebSocket protocol.
// It takes a *sql.DB parameter to perform any necessary database operations.
// The function returns an http.HandlerFunc that can be used as a handler for WebSocket requests.
//
// Example usage:
//
//	http.HandleFunc("/websocket", WebsocketHandler(db))
//
// Parameters:
//   - db: A pointer to a sql.DB object representing the database connection.
//
// Returns:
//
//	An http.HandlerFunc that handles WebSocket requests.
func WebsocketHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Upgrade the HTTP connection to a WebSocket connection
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			Log(ErrorLevel, "Error upgrading the HTTP connection to a WebSocket connection")
			http.Error(w, "Error upgrading the HTTP connection to a WebSocket connection", http.StatusInternalServerError)
			return
		}
		defer conn.Close()
		for {

			// Read the message from the client
			_, msg, err := conn.ReadMessage()
			if err != nil {
				Log(ErrorLevel, "Error reading the message from the client")
				http.Error(w, "Error reading the message from the client", http.StatusInternalServerError)
				return
			}

			// Send the message back to the client
			if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				Log(ErrorLevel, "Error sending the message back to the client")
				http.Error(w, "Error sending the message back to the client", http.StatusInternalServerError)
				return
			}
		}
	}
}
