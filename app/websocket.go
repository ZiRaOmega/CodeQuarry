package app

//Import websocket
import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type WSMessage struct {
	Type    string      `json:"type"`
	Content interface{} `json:"content"`
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
			wsmessage := WSMessage{}
			// Read the message from the client
			err := conn.ReadJSON(&wsmessage)
			if err != nil {
				Log(ErrorLevel, "Error reading the message from the client")
				return
			}
			switch wsmessage.Type {
			case "session":
				if wsmessage.Content == nil {
					conn.WriteJSON(WSMessage{Type: "session", Content: "empty"})
					break
				}
				session_id := wsmessage.Content.(string)
				// Check if the session ID is valid
				if !isValidSession(session_id, db) {
					conn.WriteJSON(WSMessage{Type: "session", Content: "expired"})

				} else {
					conn.WriteJSON(WSMessage{Type: "session", Content: "valid"})
				}
			}
		}
	}
}

func isValidSession(session_id string, db *sql.DB) bool {
	var expireAt time.Time
	err := db.QueryRow("SELECT expire_at FROM Sessions WHERE uuid = $1", session_id).Scan(&expireAt)
	if err != nil {
		Log(ErrorLevel, "Error fetching session from database"+err.Error())
		return false
	}
	if time.Now().After(expireAt) {
		Log(ErrorLevel, "Session expired")
		//DeleteSession(session_id, db)
		return false
	} else {
		return true
	}
}

func DeleteSession(session_id string, db *sql.DB) error {
	stmt, err := db.Prepare("DELETE FROM Sessions WHERE uuid = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.Exec(session_id); err != nil {
		return err
	}
	return nil
}
