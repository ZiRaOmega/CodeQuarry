package app

//Import websocket
import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
var ConnectionList = []*websocket.Conn{} // List of all connections to the server

type WSMessage struct {
	Type      string      `json:"type"`
	Content   interface{} `json:"content"`
	SessionID string      `json:"session_id"`
}

type Vote struct {
	QuestionID int `json:"question_id"`
	Upvote     int `json:"upvote"`
	Downvote   int `json:"downvote"`
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
		// Add the new WebSocket connection
		ConnectionList = append(ConnectionList, conn)
		defer conn.Close()
		for {
			wsmessage := WSMessage{}
			// Read the message from the client
			err := conn.ReadJSON(&wsmessage)
			if err != nil {
				Log(ErrorLevel, "Error reading the message from the client")
				RemoveConnFromList(conn)
				conn.Close()
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
			case "upvote":
				err := HandleUpvote(db, wsmessage.Content.(float64), wsmessage.SessionID)
				upvote, downvote := SendNewVoteCount(db, wsmessage.Content.(float64))
				vote := Vote{QuestionID: int(wsmessage.Content.(float64)), Upvote: upvote, Downvote: downvote}
				if err != nil {
					conn.WriteJSON(WSMessage{Type: "error", Content: "Failed to upvote"})
				} else {
					conn.WriteJSON(WSMessage{Type: "voteUpdate", Content: vote, SessionID: wsmessage.SessionID})
					BroadcastMessage(WSMessage{Type: "voteUpdate", Content: vote, SessionID: ""}, conn)
				}
			case "downvote":
				err := HandleDownvote(db, wsmessage.Content.(float64), wsmessage.SessionID)
				upvote, downvote := SendNewVoteCount(db, wsmessage.Content.(float64))
				vote := Vote{QuestionID: int(wsmessage.Content.(float64)), Upvote: upvote, Downvote: downvote}
				if err != nil {
					conn.WriteJSON(WSMessage{Type: "error", Content: "Failed to downvote"})
				} else {
					conn.WriteJSON(WSMessage{Type: "voteUpdate", Content: vote, SessionID: wsmessage.SessionID})
					BroadcastMessage(WSMessage{Type: "voteUpdate", Content: vote, SessionID: ""}, conn)
				}
			case "createPost":
				// Assuming the content has all necessary information
				content := wsmessage.Content.(map[string]interface{})
				fmt.Println(content)
				user_id, err := getUserIDUsingSessionID(wsmessage.SessionID, db)
				if err != nil {
					conn.WriteJSON(WSMessage{Type: "error", Content: "Failed to identify user"})
					break
				}
				quest := Question{
					Title:       content["title"].(string),
					Description: content["description"].(string),
					Content:     content["content"].(string),
				}
				subject_id, _ := strconv.Atoi(content["subject_id"].(string)) // handle error properly in production
				err = CreateQuestion(db, quest, user_id, subject_id)
				if err != nil {
					conn.WriteJSON(WSMessage{Type: "error", Content: "Failed to create post"})
				} else {
					// On successful question creation, send an update message
					updatedSubject, _ := FetchSubjectWithQuestionCount(db, subject_id) // Implement this method
					conn.WriteJSON(WSMessage{Type: "postCreated", Content: updatedSubject})
					BroadcastMessage(WSMessage{Type: "postCreated", Content: updatedSubject, SessionID: ""}, conn)
				}
			}

		}
	}
}
func RemoveConnFromList(conn *websocket.Conn) {
	for i, c := range ConnectionList {
		if c == conn {
			ConnectionList = append(ConnectionList[:i], ConnectionList[i+1:]...)
			break
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
		DeleteSession(session_id, db)
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

func BroadcastMessage(message WSMessage, currentConn *websocket.Conn) {
	for _, conn := range ConnectionList {
		if conn != currentConn {
			conn.WriteJSON(message)
		}
	}
}
