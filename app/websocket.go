package app

// Import websocket
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
type Vote_response_count struct {
	ResponseID int `json:"response_id"`
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
			//fmt.Println(wsmessage)
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
				if !isValidSession(wsmessage.SessionID, db) {
					http.Redirect(w, r, "/", http.StatusSeeOther)
					return
				}
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
				if !isValidSession(wsmessage.SessionID, db) {
					http.Redirect(w, r, "/", http.StatusSeeOther)
					return
				}
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
				if !isValidSession(wsmessage.SessionID, db) {
					http.Redirect(w, r, "/", http.StatusSeeOther)
					return
				}
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
					fmt.Println(err.Error())
					conn.WriteJSON(WSMessage{Type: "error", Content: "Failed to create post"})
				} else {
					// On successful question creation, send an update message
					updatedSubject, _ := FetchSubjectWithQuestionCount(db, subject_id) // Implement this method
					conn.WriteJSON(WSMessage{Type: "postCreated", Content: updatedSubject})
					BroadcastMessage(WSMessage{Type: "postCreated", Content: updatedSubject, SessionID: ""}, conn)
				}
			case "deletePost":
				/*{
					content: id,
					type: "deletePost",
					session_id: getCookie("session")
				}*/
				question_id, err := strconv.Atoi(wsmessage.Content.(string))
				if err != nil {
					conn.WriteJSON(WSMessage{Type: "error", Content: "Invalid question ID"})
					break
				}
				user_id, err := getUserIDUsingSessionID(wsmessage.SessionID, db)
				if !isValidSession(wsmessage.SessionID, db) {
					http.Redirect(w, r, "/", http.StatusSeeOther)
					return
				}
				if err != nil {
					conn.WriteJSON(WSMessage{Type: "error", Content: "Failed to identify user"})
					break
				}
				err = UserDeleteQuestion(db, question_id, user_id)
				if err != nil {
					fmt.Println(err.Error())
					conn.WriteJSON(WSMessage{Type: "error", Content: "Failed to delete post"})
				} else {
					// On successful question deletion, send an update message
					//updatedSubject, _ := FetchSubjectWithQuestionCount(db, question_id) // Implement this method
					/* conn.WriteJSON(WSMessage{Type: "postDeleted", Content: question_id}) */
					XP, err := FetchXP(db, user_id)
					if err != nil {
						conn.WriteJSON(WSMessage{Type: "error", Content: "Failed to fetch XP"})
						break
					}
					conn.WriteJSON(WSMessage{Type: "XP", Content: XP, SessionID: wsmessage.SessionID})
					BroadcastMessage(WSMessage{Type: "postDeleted", Content: question_id, SessionID: ""}, nil)
				}

			case "questionCompareUser":
				content := wsmessage.Content.(float64)
				questionID := int(content)
				userID, err := getUserIDUsingSessionID(wsmessage.SessionID, db)
				if !isValidSession(wsmessage.SessionID, db) {
					http.Redirect(w, r, "/", http.StatusSeeOther)
					return
				}
				if err != nil {
					conn.WriteJSON(WSMessage{Type: "error", Content: "Failed to identify user"})
					break
				}
				if CheckIfQuestionIsMine(db, questionID, float64(userID)) {
					conn.WriteJSON(WSMessage{Type: "questionCompareUser", Content: true})
				} else {
					conn.WriteJSON(WSMessage{Type: "questionCompareUser", Content: false})
				}
			case "bestAnswer":
				contentMap, ok := wsmessage.Content.(map[string]interface{})
				if !isValidSession(wsmessage.SessionID, db) {
					http.Redirect(w, r, "/", http.StatusSeeOther)
					return
				}
				fmt.Println(contentMap)
				if !ok {
					fmt.Println("Invalid content type for bestAnswer")
					// Optionally send an error response back to the client
					continue
				}

				// Extracting the answer ID and converting it to an integer

				answerID, err := strconv.Atoi(contentMap["answer_id"].(string)) // JSON numbers are decoded as floats
				if err != nil {
					fmt.Println("Invalid or missing answer_id")
					// Optionally send an error response back to the client
					continue
				}

				questionID, err := strconv.Atoi(contentMap["question_id"].(string)) // JSON numbers are decoded as floats
				if err != nil {
					fmt.Println("Invalid or missing question_id")
					// Optionally send an error response back to the client
					continue
				}

				// Now attempt to insert the best answer

				err = InsertBestAnswer(db, answerID)

				if err != nil {
					fmt.Printf("Error inserting best answer: %v\n", err)
					// Optionally send an error response back to the client
				} else {
					fmt.Println("Successfully set best answer")
					question_best_answer := GetBestAnswerFromQuestion(db, questionID)
					question_id_from_answer_id := getQuestionIDFromResponseID(db, answerID)

					conn.WriteJSON(WSMessage{Type: "bestAnswer", Content: map[string]interface{}{"question_best_answer": question_best_answer, "answer_id": answerID, "question_id": question_id_from_answer_id}})
					BroadcastMessage(WSMessage{Type: "bestAnswer", Content: map[string]interface{}{"question_best_answer": question_best_answer, "answer_id": answerID, "question_id": question_id_from_answer_id}, SessionID: ""}, nil)
				}
			case "addFavori":
				session_id := wsmessage.SessionID
				contentMap := wsmessage.Content.(float64)
				if !isValidSession(wsmessage.SessionID, db) {
					http.Redirect(w, r, "/", http.StatusSeeOther)
					return
				}
				// Check session and get user_id
				user_id, err := getUserIDUsingSessionID(session_id, db)
				if err != nil {
					conn.WriteJSON(WSMessage{Type: "addFavori", Content: "error"})
				} else {
					question_id := int(contentMap)
					/* if err != nil {
						conn.WriteJSON(WSMessage{Type: "addFavori", Content: "error"})
					} */
					if isItInFavori(db, user_id, question_id) {
						err = DeleteFavori(db, user_id, question_id)
						if err == nil {
							conn.WriteJSON(WSMessage{Type: "addFavori", Content: GetQuestionIdOfFavorite(db, user_id)})
						} else {
							conn.WriteJSON(WSMessage{Type: "addFavori", Content: "error"})
						}
					} else {
						err = AddFavori(db, user_id, question_id)
						if err == nil {
							conn.WriteJSON(WSMessage{Type: "addFavori", Content: GetQuestionIdOfFavorite(db, user_id)})
						} else {
							conn.WriteJSON(WSMessage{Type: "addFavori", Content: "already In Favori"})
						}
					}
				}
			case "deleteFavori":
				session_id := wsmessage.SessionID
				contentMap := wsmessage.Content.(string)
				if !isValidSession(wsmessage.SessionID, db) {
					http.Redirect(w, r, "/", http.StatusSeeOther)
					return
				}
				// Check session and get user_id
				user_id, err := getUserIDUsingSessionID(session_id, db)
				if err != nil {
					conn.WriteJSON(WSMessage{Type: "deleteFavori", Content: "error"})
				} else {
					question_id, err := strconv.Atoi(contentMap)
					if err != nil {
						conn.WriteJSON(WSMessage{Type: "deleteFavori", Content: "error"})
					}
					err = DeleteFavori(db, user_id, question_id)
					if err == nil {
						conn.WriteJSON(WSMessage{Type: "deleteFavori", Content: "success"})
					} else {
						conn.WriteJSON(WSMessage{Type: "deleteFavori", Content: "error"})
					}
				}
			case "upvote_response":
				if !isValidSession(wsmessage.SessionID, db) {
					http.Redirect(w, r, "/", http.StatusSeeOther)
					return
				}
				err := HandleUpvoteResponse(db, wsmessage.Content.(float64), wsmessage.SessionID)
				if err != nil {
					conn.WriteJSON(WSMessage{Type: "error", Content: "Failed to upvote response"})
				} else {
					up, down := SendNewVoteCountResponse(db, wsmessage.Content.(float64))
					vote := Vote_response_count{ResponseID: int(wsmessage.Content.(float64)), Upvote: up, Downvote: down}
					conn.WriteJSON(WSMessage{Type: "responseVoteUpdate", Content: vote, SessionID: wsmessage.SessionID})
					BroadcastMessage(WSMessage{Type: "responseVoteUpdate", Content: vote, SessionID: ""}, conn)
				}
			case "downvote_response":
				if !isValidSession(wsmessage.SessionID, db) {
					http.Redirect(w, r, "/", http.StatusSeeOther)
					return
				}
				err := HandleDownvoteResponse(db, wsmessage.Content.(float64), wsmessage.SessionID)
				if err != nil {
					conn.WriteJSON(WSMessage{Type: "error", Content: "Failed to downvote response"})
				} else {
					up, down := SendNewVoteCountResponse(db, wsmessage.Content.(float64))
					vote := Vote_response_count{ResponseID: int(wsmessage.Content.(float64)), Upvote: up, Downvote: down}
					conn.WriteJSON(WSMessage{Type: "responseVoteUpdate", Content: vote, SessionID: wsmessage.SessionID})
					BroadcastMessage(WSMessage{Type: "responseVoteUpdate", Content: vote, SessionID: ""}, conn)
				}
			case "modify_question":
				contentMap, ok := wsmessage.Content.(map[string]interface{})
				if !ok {
					fmt.Println("Invalid content type for modify_question")
					// Optionally send an error response back to the client
					continue
				}
				user_id, err := getUserIDUsingSessionID(wsmessage.SessionID, db)
				if err != nil {
					conn.WriteJSON(WSMessage{Type: "error", Content: "Failed to identify user"})
					break
				}
				question_id, err := strconv.Atoi(contentMap["question_id"].(string))
				if err != nil {
					fmt.Println("Invalid or missing question_id")
					// Optionally send an error response back to the client
					continue
				}
				if isValidSession(wsmessage.SessionID, db) {
					ModifyQuestion(db, question_id, contentMap["title"].(string), contentMap["description"].(string), contentMap["content"].(string), user_id)
					updatedQuestion, err := FetchQuestionByQuestionID(db, question_id, user_id)
					if err != nil {
						conn.WriteJSON(WSMessage{Type: "error", Content: "Failed to fetch updated question"})
					} else {
						conn.WriteJSON(WSMessage{Type: "questionModified", Content: updatedQuestion})
						BroadcastMessage(WSMessage{Type: "questionModified", Content: updatedQuestion, SessionID: ""}, conn)
					}
				} else {
					conn.WriteJSON(WSMessage{Type: "error", Content: "Invalid session"})
					http.Redirect(w, r, "/", http.StatusSeeOther)
				}
			case "modify_response":
				contentMap, ok := wsmessage.Content.(map[string]interface{})
				if !ok {
					fmt.Println("Invalid content type for modify_response")
					// Optionally send an error response back to the client
					continue
				}
				user_id, err := getUserIDUsingSessionID(wsmessage.SessionID, db)
				if err != nil {
					conn.WriteJSON(WSMessage{Type: "error", Content: "Failed to identify user"})
					break
				}
				if !isValidSession(wsmessage.SessionID, db) {
					conn.WriteJSON(WSMessage{Type: "error", Content: "Invalid session"})
					http.Redirect(w, r, "/", http.StatusSeeOther)
					break
				}
				response_id := int(contentMap["response_id"].(float64))
				if err != nil {
					fmt.Println("Invalid or missing response_id")
					// Optionally send an error response back to the client
					continue
				}
				question_id := int(contentMap["question_id"].(float64))
				if err != nil {
					fmt.Println("Invalid or missing response_id")
					// Optionally send an error response back to the client
					continue
				}
				err = ModifyResponse(db, response_id, contentMap["content"].(string), contentMap["description"].(string), user_id)
				if err != nil {
					fmt.Println(err)
					conn.WriteJSON(WSMessage{Type: "error", Content: "Failed to modify response"})
				} else {
					updatedResponse, err := FetchResponseByQuestion(db, question_id, user_id)
					if err != nil {
						conn.WriteJSON(WSMessage{Type: "error", Content: "Failed to fetch updated response"})
					} else {
						conn.WriteJSON(WSMessage{Type: "responseModified", Content: updatedResponse})
						BroadcastMessage(WSMessage{Type: "responseModified", Content: updatedResponse, SessionID: ""}, conn)
					}
				}
			case "editQuestionPanel":
				/*const data = {
				    type: "editQuestionPanel",
				    content: {
				        id: id,
				        title: inputs[0].value,
				        description: textareas[0].value,
				        content: textareas[1].value,
				        creationDate: inputs[1].value,
				        updateDate: inputs[2].value,
				        upvotes: inputs[3].value,
				        downvotes: inputs[4].value
				    },
				    session_id: getCookie("session")
				};*/
				contentMap, ok := wsmessage.Content.(map[string]interface{})
				if !ok {
					fmt.Println("Invalid content type for editQuestionPanel")
					// Optionally send an error response back to the client
					continue
				}
				user_id, err := getUserIDUsingSessionID(wsmessage.SessionID, db)
				if err != nil {
					conn.WriteJSON(WSMessage{Type: "error", Content: "Failed to identify user"})
					break
				}
				question_id := int(contentMap["id"].(float64))
				question_title := contentMap["title"].(string)
				question_description := contentMap["description"].(string)
				question_content := contentMap["content"].(string)
				question_creation_date, err := time.Parse("2024-05-01", contentMap["creationDate"].(string))
				question_update_date, err := time.Parse("2024-05-0", contentMap["updateDate"].(string))
				question_upvotes, err := strconv.Atoi(contentMap["upvotes"].(string))
				question_downvotes, err := strconv.Atoi(contentMap["downvotes"].(string))
				if FetchRankByUserID(db, user_id) > 0 && isValidSession(wsmessage.SessionID, db) {
					err := ModifyQuestionPanel(db, question_id, question_title, question_description, question_content, question_creation_date, question_update_date, question_upvotes, question_downvotes)
					if err != nil {
						conn.WriteJSON(WSMessage{Type: "error", Content: "Failed to modify question"})
						break
					}
					updatedQuestion, err := FetchQuestionByQuestionID(db, question_id, user_id)
					if err != nil {
						conn.WriteJSON(WSMessage{Type: "error", Content: "Failed to fetch updated question"})
					} else {
						conn.WriteJSON(WSMessage{Type: "questionModified", Content: updatedQuestion})
						BroadcastMessage(WSMessage{Type: "questionModified", Content: updatedQuestion, SessionID: ""}, conn)
					}
				} else {
					conn.WriteJSON(WSMessage{Type: "error", Content: "Unauthorized"})
				}
			case "editResponsePanel":
				/*const data = {
				    type: "editResponsePanel",
				    content: {
				        id: responseId,
						question_id: question_id,
				        content: textareas[0].value,
				        description: inputs[0].value,
				        creationDate: inputs[1].value,
				        updateDate: inputs[2].value,
				        upvotes: inputs[3].value,
				        downvotes: inputs[4].value
				    },
				    session_id: getCookie("session")
				};*/
				contentMap, ok := wsmessage.Content.(map[string]interface{})
				if !ok {
					fmt.Println("Invalid content type for editResponsePanel")
					// Optionally send an error response back to the client
					continue
				}
				user_id, err := getUserIDUsingSessionID(wsmessage.SessionID, db)
				if err != nil {
					conn.WriteJSON(WSMessage{Type: "error", Content: "Failed to identify user"})
					break
				}
				response_id := int(contentMap["id"].(float64))
				question_id := int(contentMap["question_id"].(float64))
				response_content := contentMap["content"].(string)
				response_description := contentMap["description"].(string)
				response_creation_date, err := time.Parse("2024-05-01", contentMap["creationDate"].(string))
				response_update_date, err := time.Parse("2024-05-0", contentMap["updateDate"].(string))
				response_upvotes, err := strconv.Atoi(contentMap["upvotes"].(string))
				response_downvotes, err := strconv.Atoi(contentMap["downvotes"].(string))
				if FetchRankByUserID(db, user_id) > 0 && isValidSession(wsmessage.SessionID, db) {
					err := ModifyResponsePanel(db, response_id, response_content, response_description, response_creation_date, response_update_date, response_upvotes, response_downvotes)
					if err != nil {
						conn.WriteJSON(WSMessage{Type: "error", Content: "Failed to modify response"})
						break
					}
					updatedQuestion, err := FetchQuestionByQuestionID(db, question_id, user_id)
					if err != nil {
						conn.WriteJSON(WSMessage{Type: "error", Content: "Failed to fetch updated question"})
					} else {
						conn.WriteJSON(WSMessage{Type: "questionModified", Content: updatedQuestion})
						BroadcastMessage(WSMessage{Type: "questionModified", Content: updatedQuestion, SessionID: ""}, conn)
					}
				} else {
					conn.WriteJSON(WSMessage{Type: "error", Content: "Unauthorized"})
				}
			case "editSubjectPanel":
				/*const data = {
				    type: "editSubjectPanel",
				    content: {
				        id: id,
				        title: inputs[0].value,
				        description: textareas[0].value,
				        creationDate: inputs[1].value,
				        updateDate: inputs[2].value,
				    },
				    session_id: getCookie("session")
				};*/
				contentMap, ok := wsmessage.Content.(map[string]interface{})
				if !ok {
					fmt.Println("Invalid content type for editSubjectPanel")
					// Optionally send an error response back to the client
					continue
				}
				user_id, err := getUserIDUsingSessionID(wsmessage.SessionID, db)
				if err != nil {
					conn.WriteJSON(WSMessage{Type: "error", Content: "Failed to identify user"})
					break
				}
				subject_id := int(contentMap["id"].(float64))
				subject_title := contentMap["title"].(string)
				subject_description := contentMap["description"].(string)
				subject_creation_date, err := time.Parse("2024-05-01", contentMap["creationDate"].(string))
				subject_update_date, err := time.Parse("2024-05-0", contentMap["updateDate"].(string))

				if FetchRankByUserID(db, user_id) > 0 && isValidSession(wsmessage.SessionID, db) {
					err := ModifySubjectPanel(db, subject_id, subject_title, subject_description, subject_creation_date, subject_update_date)
					if err != nil {
						conn.WriteJSON(WSMessage{Type: "error", Content: "Failed to modify subject"})
						break
					}
					updatedSubject := FetchSubjects(db)
					if err != nil {
						conn.WriteJSON(WSMessage{Type: "error", Content: "Failed to fetch updated subject"})
					} else {
						conn.WriteJSON(WSMessage{Type: "subjectModified", Content: updatedSubject})
						BroadcastMessage(WSMessage{Type: "subjectModified", Content: updatedSubject, SessionID: ""}, conn)
					}
				}
			case "addSubjectPanel":
				/* const data = {
				    type: "addSubjectPanel",
				    content: {
				        title: inputs[0].value,
				        description: textareas[0].value,
				    },
				    session_id: getCookie("session")
				};*/
				contentMap, ok := wsmessage.Content.(map[string]interface{})
				if !ok || len(contentMap) == 0 {
					fmt.Println("Invalid content type for addSubjectPanel")
					// Optionally send an error response back to the client
					continue
				}
				user_id, err := getUserIDUsingSessionID(wsmessage.SessionID, db)
				if err != nil {
					conn.WriteJSON(WSMessage{Type: "error", Content: "Failed to identify user"})
					break
				}
				subject_title := contentMap["title"].(string)
				subject_description := contentMap["description"].(string)
				if FetchRankByUserID(db, user_id) > 0 && isValidSession(wsmessage.SessionID, db) {
					err := InsertInSubject(db, subject_title, subject_description)
					if err != nil {
						conn.WriteJSON(WSMessage{Type: "error", Content: "Failed to create subject"})
						break
					}
					updatedSubject := FetchSubjects(db)
					if err != nil {
						conn.WriteJSON(WSMessage{Type: "error", Content: "Failed to fetch updated subject"})
					} else {
						conn.WriteJSON(WSMessage{Type: "subjectModified", Content: updatedSubject})
						BroadcastMessage(WSMessage{Type: "subjectModified", Content: updatedSubject, SessionID: ""}, conn)
					}
				}
			case "deleteSubjectPanel":
				/*const data = {
				    type: "deleteSubjectPanel",
				    content: {
				        id: id,
				    },
				    session_id: getCookie("session")
				};*/
				contentMap, ok := wsmessage.Content.(map[string]interface{})
				if !ok {
					fmt.Println("Invalid content type for deleteSubjectPanel")
					// Optionally send an error response back to the client
					continue
				}
				user_id, err := getUserIDUsingSessionID(wsmessage.SessionID, db)
				if err != nil {
					conn.WriteJSON(WSMessage{Type: "error", Content: "Failed to identify user"})
					break
				}
				panel_id := int(contentMap["id"].(float64))
				if FetchRankByUserID(db, user_id) > 0 && isValidSession(wsmessage.SessionID, db) {
					err := DeleteSubjectPanel(db, panel_id)
					if err != nil {
						conn.WriteJSON(WSMessage{Type: "error", Content: "Failed to delete subject"})
						break
					}
					updatedSubject := FetchSubjects(db)
					if err != nil {
						conn.WriteJSON(WSMessage{Type: "error", Content: "Failed to fetch updated subject"})
					} else {
						conn.WriteJSON(WSMessage{Type: "subjectModified", Content: updatedSubject})
						BroadcastMessage(WSMessage{Type: "subjectModified", Content: updatedSubject, SessionID: ""}, conn)
					}
				}
			case "deleteQuestionPanel":
				/*const data = {
				    type: "deleteQuestionPanel",
				    content: {
				        id: id,
				    },
				    session_id: getCookie("session")
				};*/
				contentMap, ok := wsmessage.Content.(map[string]interface{})
				if !ok {
					fmt.Println("Invalid content type for deleteQuestionPanel")
					// Optionally send an error response back to the client
					continue
				}
				user_id, err := getUserIDUsingSessionID(wsmessage.SessionID, db)
				question_id := int(contentMap["id"].(float64))
				if err != nil {
					conn.WriteJSON(WSMessage{Type: "error", Content: "Failed to identify user"})
					break
				}
				if FetchRankByUserID(db, user_id) > 0 && isValidSession(wsmessage.SessionID, db) {
					err := DeleteQuestionPanel(db, question_id)
					if err != nil {
						conn.WriteJSON(WSMessage{Type: "error", Content: "Failed to delete question"})
						break
					}

					if err != nil {
						conn.WriteJSON(WSMessage{Type: "error", Content: "Failed to fetch updated subject"})
					} else {
						conn.WriteJSON(WSMessage{Type: "postDeleted", Content: question_id})
						BroadcastMessage(WSMessage{Type: "postDeleted", Content: question_id, SessionID: ""}, conn)
					}
				} else {
					conn.WriteJSON(WSMessage{Type: "error", Content: "Unauthorized"})
				}
			case "deleteResponsePanel":
				/* const data = {
				    type: "deleteResponsePanel",
				    content: {
				        id: id,
				        question_id: question_id
				    },
				    session_id: getCookie("session")
				};*/
				contentMap, ok := wsmessage.Content.(map[string]interface{})
				if !ok {
					fmt.Println("Invalid content type for deleteResponsePanel")
					// Optionally send an error response back to the client
					continue
				}
				user_id, err := getUserIDUsingSessionID(wsmessage.SessionID, db)
				if err != nil {
					conn.WriteJSON(WSMessage{Type: "error", Content: "Failed to identify user"})
					break
				}
				response_id := int(contentMap["id"].(float64))
				question_id := int(contentMap["question_id"].(float64))
				if FetchRankByUserID(db, user_id) > 0 && isValidSession(wsmessage.SessionID, db) {
					err := DeleteResponsePanel(db, response_id)
					if err != nil {
						conn.WriteJSON(WSMessage{Type: "error", Content: "Failed to delete response"})
						break
					}
					updatedQuestion, err := FetchQuestionByQuestionID(db, question_id, user_id)
					if err != nil {
						conn.WriteJSON(WSMessage{Type: "error", Content: "Failed to fetch updated question"})
					} else {
						conn.WriteJSON(WSMessage{Type: "questionModified", Content: updatedQuestion})
						BroadcastMessage(WSMessage{Type: "questionModified", Content: updatedQuestion, SessionID: ""}, conn)
					}
				} else {
					conn.WriteJSON(WSMessage{Type: "error", Content: "Unauthorized"})
				}
			case "editUserPanel":
				/*const data = {
				    type: "editUserPanel",
				    content: {
				        id: id,
				        firstname : inputs[0].value,
				        lastname : inputs[1].value,
				        username : inputs[2].value,
				        email : inputs[3].value,
				        bio : textareas[0].value,
				        website : inputs[4].value,
				        github : inputs[5].value,
				        xp : inputs[6].value,
				        rank : inputs[7].value,
				        schoolyear : inputs[8].value
				    },
				    session_id: getCookie("session")
				};*/
				contentMap, ok := wsmessage.Content.(map[string]interface{})
				if !ok {
					fmt.Println("Invalid content type for editUserPanel")
					// Optionally send an error response back to the client
					continue
				}
				user_id, err := getUserIDUsingSessionID(wsmessage.SessionID, db)
				if err != nil {
					conn.WriteJSON(WSMessage{Type: "error", Content: "Failed to identify user"})
					break
				}
				editor_rank := FetchRankByUserID(db, user_id)
				if editor_rank > 0 && isValidSession(wsmessage.SessionID, db) {
					user_id := int(contentMap["id"].(float64))
					firstname := contentMap["firstname"].(string)
					lastname := contentMap["lastname"].(string)
					username := contentMap["username"].(string)
					email := contentMap["email"].(string)
					bio := contentMap["bio"].(string)
					website := contentMap["website"].(string)
					github := contentMap["github"].(string)
					xp, err := strconv.Atoi(contentMap["xp"].(string))
					rank, err := strconv.Atoi(contentMap["rank"].(string))
					if editor_rank < rank {
						conn.WriteJSON(WSMessage{Type: "error", Content: "Unauthorized"})
						break
					}
					schoolyear, err := time.Parse("2024-04-30", contentMap["schoolyear"].(string))
					err = ModifyUserPanel(db, user_id, firstname, lastname, username, email, bio, website, github, xp, rank, schoolyear)
					if err != nil {
						fmt.Println(err.Error())
						conn.WriteJSON(WSMessage{Type: "error", Content: "Failed to modify user"})
						break
					}
				} else {
					conn.WriteJSON(WSMessage{Type: "error", Content: "Unauthorized"})
				}
			case "deleteUserPanel":
				/* const data = {
				    type: "deleteUserPanel",
				    content: {
				        id: id,
				    },
				    session_id: getCookie("session")
				};*/
				contentMap, ok := wsmessage.Content.(map[string]interface{})
				if !ok {
					fmt.Println("Invalid content type for deleteUserPanel")
					// Optionally send an error response back to the client
					continue
				}
				user_id, err := getUserIDUsingSessionID(wsmessage.SessionID, db)
				if err != nil {
					conn.WriteJSON(WSMessage{Type: "error", Content: "Failed to identify user"})
					break
				}
				if FetchRankByUserID(db, user_id) > 1 && isValidSession(wsmessage.SessionID, db) {
					user_id := int(contentMap["id"].(float64))
					err := DeleteUserPanel(db, user_id)
					if err != nil {
						conn.WriteJSON(WSMessage{Type: "error", Content: "Failed to delete user"})
						break
					}
				} else {
					conn.WriteJSON(WSMessage{Type: "error", Content: "Unauthorized"})
				}
			case "deleteAvatarPanel":
				/*const data = {
				    type: "deleteAvatarPanel",
				    content: {
				        user_id: user_id,
				    },
				    session_id: getCookie("session")
				};*/
				contentMap, ok := wsmessage.Content.(map[string]interface{})
				if !ok {
					fmt.Println("Invalid content type for deleteAvatarPanel")
					// Optionally send an error response back to the client
					continue
				}
				user_id, err := getUserIDUsingSessionID(wsmessage.SessionID, db)
				if err != nil {
					conn.WriteJSON(WSMessage{Type: "error", Content: "Failed to identify user"})
					break
				}
				if FetchRankByUserID(db, user_id) > 0 && isValidSession(wsmessage.SessionID, db) {
					user_id := int(contentMap["user_id"].(float64))
					err := DeleteAvatar(db, user_id)
					if err != nil {
						conn.WriteJSON(WSMessage{Type: "error", Content: "Failed to delete avatar"})
						break
					}
				} else {
					conn.WriteJSON(WSMessage{Type: "error", Content: "Unauthorized"})
				}
			}
		}
	}
}
func DeleteAvatar(db *sql.DB, user_id int) error {
	stmt, err := db.Prepare("UPDATE users SET avatar = NULL WHERE id_student = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.Exec(user_id); err != nil {
		return err
	}
	return nil
}
func DeleteUserPanel(db *sql.DB, user_id int) error {
	stmt, err := db.Prepare("DELETE FROM users WHERE id_student = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.Exec(user_id); err != nil {
		return err
	}
	return nil
}
func ModifyUserPanel(db *sql.DB, user_id int, firstname string, lastname string, username string, email string, bio string, website string, github string, xp int, rank int, schoolyear time.Time) error {
	stmt, err := db.Prepare("UPDATE users SET firstname = $1, lastname = $2, username = $3, email = $4, bio = $5, website = $6, github = $7, xp = $8, rang_rank_ = $9, school_year = $10 WHERE id_student = $11")
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.Exec(firstname, lastname, username, email, bio, website, github, xp, rank, schoolyear, user_id); err != nil {
		return err
	}
	return nil
}
func DeleteResponsePanel(db *sql.DB, response_id int) error {
	stmt, err := db.Prepare("DELETE FROM Response WHERE id_response = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.Exec(response_id); err != nil {
		return err
	}
	return nil
}
func DeleteQuestionPanel(db *sql.DB, question_id int) error {
	stmt, err := db.Prepare("DELETE FROM Question WHERE id_question = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.Exec(question_id); err != nil {
		return err
	}
	return nil
}
func DeleteSubjectPanel(db *sql.DB, subject_id int) error {
	stmt, err := db.Prepare("DELETE FROM Subject WHERE id_subject = $1")
	if err != nil {
		return err
	}
	_, err2 := stmt.Exec(subject_id)
	return err2
}
func ModifySubjectPanel(db *sql.DB, subject_id int, title string, description string, creation_date time.Time, update_date time.Time) error {
	stmt, err := db.Prepare("UPDATE Subject SET title = $1, description = $2, creation_date = $3, update_date = $4 WHERE id_subject = $5")
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.Exec(title, description, creation_date, update_date, subject_id); err != nil {
		return err
	}
	return nil
}
func ModifyResponsePanel(db *sql.DB, response_id int, content string, description string, creationDate time.Time, updateDate time.Time, upVotes int, downVotes int) error {
	stmt, err := db.Prepare("UPDATE Response SET content = $1, description = $2, creation_date = $3, update_date = $4, upvotes = $5, downvotes = $6 WHERE id_response = $7")
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.Exec(content, description, creationDate, updateDate, upVotes, downVotes, response_id); err != nil {
		return err
	}
	return nil
}
func ModifyQuestionPanel(db *sql.DB, question_id int, title string, description string, content string, creation_date time.Time, update_date time.Time, upvotes int, downvotes int) error {
	stmt, err := db.Prepare("UPDATE Question SET title = $1, description = $2, content = $3, creation_date = $4, update_date = $5, upvotes = $6, downvotes = $7 WHERE id_question = $8")
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.Exec(title, description, content, creation_date, update_date, upvotes, downvotes, question_id); err != nil {
		return err
	}
	return nil
}
func ModifyResponse(db *sql.DB, response_id int, content string, description string, user_id int) error {
	fmt.Println(content)
	stmt, err := db.Prepare("UPDATE Response SET content = $1, description = $2 WHERE id_response = $3 AND id_student = $4")
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.Exec(content, description, response_id, user_id); err != nil {
		return err
	}
	return nil
}

func DeleteFavori(db *sql.DB, id_student int, question_id int) error {
	stmt, err := db.Prepare("DELETE FROM Favori WHERE id_student = $1 AND id_question = $2")
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.Exec(id_student, question_id); err != nil {
		return err
	}
	return nil
}

func AddFavori(db *sql.DB, id_student int, question_id int) error {
	stmt, err := db.Prepare("INSERT INTO Favori(id_student, id_question) VALUES($1, $2)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.Exec(id_student, question_id); err != nil {
		return err
	}
	return nil
}

func isItInFavori(db *sql.DB, id_student int, question_id int) bool {
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM Favori WHERE id_student = $1 AND id_question = $2)", id_student, question_id).Scan(&exists)
	if err != nil {
		Log(ErrorLevel, "Error checking if question is in Favori: "+err.Error())
		return false
	}
	return exists
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
