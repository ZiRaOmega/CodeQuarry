## WebSockets Overview

WebSockets provide a way to open a bi-directional, persistent communication channel between a client (typically a web browser) and a server. This allows for real-time data exchange without the need to repeatedly send HTTP requests.

### How WebSockets Work

WebSockets operate over a single long-lived connection, which makes them suitable for applications that require real-time updates such as gaming, trading platforms, or social feeds.

#### Opening a Connection

A WebSocket connection is initiated by the client through a WebSocket handshake. This is essentially an HTTP Upgrade request that tells the server that the client wishes to establish a WebSocket connection. If the server supports WebSockets, it will respond with an HTTP 101 status code, switching protocols from HTTP to WebSockets.

Example Go code using the Gorilla WebSocket library:

```go
var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func handler(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Printf("error upgrading HTTP to WebSocket: %v", err)
        return
    }
    defer conn.Close()
    // WebSocket connection successfully established
}
```

#### Message Handling

Once a WebSocket connection is established, data can be sent back and forth in the form of messages. Each message is either text or binary. WebSockets ensure that messages are sent and received in full before the next message is handled.

```go
for {
    messageType, message, err := conn.ReadMessage()
    if err != nil {
        log.Printf("error reading message: %v", err)
        break
    }
    err = conn.WriteMessage(messageType, message)
    if err != nil {
        log.Printf("error sending message: %v", err)
        break
    }
}
```

#### Closing a Connection

Either the client or server can close the WebSocket connection. The closing process involves sending a control frame with a status code indicating why the connection is closing.

### Error Handling

Error handling in WebSockets is crucial to maintain the reliability and robustness of the communication process. Typical errors include:

- **Connection Errors**: Failures during the handshake or due to network issues.
- **Read/Write Errors**: Issues when receiving or sending data, often due to connection drops or protocol violations.
- **Protocol Compliance**: Handling scenarios where either the client or server does not comply with the WebSocket protocol standards.

In Go, using the Gorilla WebSocket package, each of these errors needs to be appropriately logged and the connection should be closed cleanly to free up resources.

### Conclusion

WebSockets offer a powerful method for real-time communication in modern web applications. Proper implementation and error handling are essential to harness their full potential, especially in environments where real-time data and high reliability are crucial.

**Language > Specialist**: Go > WebSocket Communication and Session Management Expert
**Includes**: `net/http`, `database/sql`, `github.com/gorilla/websocket`
**Requirements**: Detailed documentation for a WebSocket handler function in Go that manages real-time communication and session validation, including database interactions for session management.

## Package `websocket` Documentation

This documentation covers the `WebsocketHandler` function and its associated helper functions, `isValidSession` and `DeleteSession`, which collectively manage WebSocket communications and session validations.

### 1. `WebsocketHandler`

**Function Signature**:
```go
func WebsocketHandler(db *sql.DB) http.HandlerFunc
```

**Purpose**:
`WebsocketHandler` creates an HTTP handler function that upgrades HTTP connections to WebSocket connections and handles WebSocket messages, particularly focusing on session management messages.

**Parameters**:
- `db *sql.DB`: A pointer to a `sql.DB` object representing the database connection, used for session validation and other database operations.

**Operations**:
1. **Connection Upgrade**: The handler upgrades the incoming HTTP connection to a WebSocket connection using the `upgrader.Upgrade` method.
2. **Message Handling**: The function enters a continuous loop to read and respond to WebSocket messages, handling session-related messages by validating session IDs stored in the database.
3. **Session Validation**: Upon receiving a message of type "session", it validates the session ID against the database.

**Error Handling**:
- Logs and returns an error if the connection upgrade fails.
- Ends the WebSocket connection if reading from the connection fails.

**Return Value**:
- `http.HandlerFunc`: A handler function that manages WebSocket communications.

### 2. `isValidSession`

**Function Signature**:
```go
func isValidSession(session_id string, db *sql.DB) bool
```

**Purpose**:
Checks if a given session ID is valid based on its expiration time stored in the database.

**Parameters**:
- `session_id string`: The session ID to validate.
- `db *sql.DB`: Database connection to query the session's expiration.

**Operations**:
1. **Query Database**: Fetches the session's expiration time from the database.
2. **Validate Session**: Compares the current time with the expiration time to determine if the session is still valid.

**Error Handling**:
- Logs errors related to database operations or if the session has expired.

**Return Value**:
- `bool`: Returns `true` if the session is valid, otherwise `false`.

### 3. `DeleteSession`

**Function Signature**:
```go
func DeleteSession(session_id string, db *sql.DB) error
```

**Purpose**:
Deletes a session from the database using the session ID.

**Parameters**:
- `session_id string`: The session ID to delete.
- `db *sql.DB`: Database connection to execute the delete operation.

**Operations**:
1. **Prepare SQL Statement**: Prepares a SQL statement for deleting the session.
2. **Execute SQL Statement**: Executes the deletion.

**Error Handling**:
- Handles and returns errors related to preparing or executing the SQL statement.

**Return Value**:
- `error`: Returns an error object if there is a failure in deleting the session.

### Example Usage

```go
http.HandleFunc("/websocket", WebsocketHandler(db))
```

### Best Practices

- **WebSocket Security**: Ensure WebSocket connections are only accepted from authenticated and authorized users.
- **Database Connection Management**: Use connection pooling and handle database connections carefully to avoid leaks.
- **Error Logging**: Implement comprehensive logging for debugging and monitoring the application's behavior.


The `WebsocketHandler` function in Go is designed to upgrade HTTP connections to WebSocket connections, facilitating real-time communication between the client and the server. It also performs database operations to validate session information received over the WebSocket.

### Function Signature

```go
func WebsocketHandler(db *sql.DB) http.HandlerFunc
```

#### Parameters

- `db *sql.DB`: A pointer to a `sql.DB` object, representing the active database connection. This connection is used to perform queries related to session validation.

#### Returns

- `http.HandlerFunc`: A function that conforms to the `http.HandlerFunc` type, which can be used as a handler for WebSocket requests.

### Detailed Breakdown

**Function Behavior**:

```go
return func(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        Log(ErrorLevel, "Error upgrading the HTTP connection to a WebSocket connection")
        http.Error(w, "Error upgrading the HTTP connection to a WebSocket connection", http.StatusInternalServerError)
        return
    }
    defer conn.Close()

    for {
        wsmessage := WSMessage{}
        err := conn.ReadJSON(&wsmessage)
        if err != nil {
            Log(ErrorLevel, "Error reading the message from the client")
            return
        }
        
        // Message handling based on type
        switch wsmessage.Type {
        case "session":
            handleSessionMessage(wsmessage, conn, db)
        }
    }
}
```

**Operations**:

1. **WebSocket Upgrade**: The HTTP connection is upgraded to a WebSocket connection using an upgrader object. Errors during this process are logged and an HTTP 500 error is returned to the client.

2. **Message Loop**: The handler enters a continuous loop where it reads JSON-formatted messages from the client.

3. **Session Message Handling**: If the message type is "session", it checks the session ID against the database:
   - **Validate Session**: Uses the `isValidSession` function to check if the session ID from the message is still valid.
   - **Response**: Depending on the validity of the session, a response is sent back to the client using `WriteJSON`. The response indicates whether the session is "valid", "expired", or "empty".

**Auxiliary Functions**:

- `isValidSession(session_id string, db *sql.DB) bool`: Checks the database to see if the provided session ID corresponds to an active session.

- `handleSessionMessage(wsmessage WSMessage, conn *websocket.Conn, db *sql.DB)`: Handles session-related messages by performing database checks and sending responses.

### Example Usage

```go
http.HandleFunc("/websocket", WebsocketHandler(db))
```

- This line of code sets up the `WebsocketHandler` as the handler for routes directed to "/websocket", enabling WebSocket communication on this endpoint.

### Best Practices

- **Error Handling**: Proper logging and error responses ensure that the server can gracefully handle issues during WebSocket communication.
- **Security Considerations**: Validate all data received through WebSockets to avoid potential security risks. Ensure that database operations do not expose sensitive information.
- **Resource Management**: Use `defer` to ensure resources like WebSocket connections are properly closed after use to prevent resource leaks.

---

**History**: Documented the `WebsocketHandler` function, which provides a robust method for handling WebSocket communications in Go, including session validation via database interaction.

**Source Tree**:
- (💾=saved) WebSocket Handler
  - ✅ Core WebSocket communication handling
  - ✅ Database interaction for session validation

**Next Task**: FINISHED. Consider enhancements such as support for more message types and integrating more complex session management features, like session renewal and detailed logging for monitoring and diagnostics.
## WebSocket JavaScript Client Documentation

This documentation provides an overview and detailed explanation of a JavaScript script designed to establish and handle a WebSocket connection. This script utilizes the WebSockets API to communicate with a server from a web client.

### Script Overview

The script initializes a WebSocket connection when the document is ready, which ensures that all HTML is loaded before the script runs. It is encapsulated within jQuery's `$(document).ready()` to guarantee this setup.

### Detailed Explanation

#### Initialization

```javascript
var socket;
$(document).ready(function () {
    var Current_Hostname = document.location.hostname
    socket = new WebSocket(`wss://${Current_Hostname}/ws`);
```

- **Variable Declaration**: `socket` is declared without initialization at the global scope to be accessible throughout the script.
- **Document Ready**: Ensures that the full HTML document is loaded before executing the script.
- **Current Hostname Retrieval**: Fetches the hostname of the current document location, which is used to dynamically set the WebSocket server's URL.
- **WebSocket Initialization**: Initializes a new WebSocket connection using a secure WebSocket protocol (`wss://`) to the current hostname appended with `/ws`. This URL pattern typically points to a WebSocket endpoint on the server.

#### WebSocket Event Handlers

```javascript
socket.onopen = function (e) {
    console.log("[open] Connection established");
    console.log("Sending to server");
    socket.send("Hey there from client");
};
```

- **Open Event**: Triggered when the WebSocket connection is successfully established.
- **Sending Message**: Logs the connection status and sends a greeting message to the server.

```javascript
socket.onmessage = function (event) {
    console.log(`[message] Data received from server: ${event.data}`);
};
```

- **Message Event**: Triggered when a message is received from the server. It logs the message data received.

```javascript
socket.onclose = function (event) {
    if (event.wasClean) {
        console.log(`[close] Connection closed cleanly, code=${event.code} reason=${event.reason}`);
    } else {
        console.error('[close] Connection died');
    }
};
```

- **Close Event**: Handles the closing of the WebSocket connection. It checks if the closure was clean (no errors or issues), logging appropriate messages. A clean close includes the server sending a close frame or the client successfully closing the connection.

```javascript
socket.onerror = function (error) {
    console.error(`[error] ${error.message}`);
};
```

- **Error Event**: Triggered on any error during the WebSocket communication. Logs the error message to the console.

### Best Practices

1. **Secure Connections**: Always use `wss://` (WebSocket Secure) over `ws://` to encrypt transmitted data, protecting against eavesdropping and tampering.
2. **Error Handling**: Robust error handling in WebSocket applications is crucial for reliability. Ensure that all potential events (open, message, error, close) have corresponding handlers.
3. **Reconnection Strategy**: Implement automatic reconnection in case of unexpected server disconnects or network issues.
4. **Data Handling**: Validate and sanitize all data sent to and received from the server to prevent security vulnerabilities, such as XSS (Cross-Site Scripting).
