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

## `WebsocketHandler` Function Documentation

The `WebsocketHandler` function in Go is designed to facilitate real-time communication between a client and a server over WebSockets. This documentation provides a detailed explanation of the function's operation, including its parameters, core functionality, and error handling.

### Function Signature

```go
func WebsocketHandler(db *sql.DB) http.HandlerFunc
```

#### Parameters

- `db *sql.DB`: A pointer to an SQL database instance. This parameter allows the function to interact with a database, although the specific usage within the function isn't demonstrated in the given snippet.

#### Return Type

- `http.HandlerFunc`: The function returns an `http.HandlerFunc`, which is a type that handles HTTP requests in the Go standard library's `net/http` package.

### Functionality Overview

The function initializes a WebSocket connection and continuously reads messages from the client, sending them back as echoes. Here's a breakdown of the process:

1. **HTTP to WebSocket Upgrade**: 
   - The function begins by upgrading an incoming HTTP connection to a WebSocket connection using an `upgrader` object. This is critical for initiating WebSocket communication.
   - If the upgrade fails, it logs the error and sends an HTTP 500 internal server error response back to the client.

2. **Message Handling Loop**:
   - After successfully establishing a WebSocket connection, the function enters an infinite loop, continually waiting to read messages from the client.
   - If reading a message fails, it logs the error and sends an HTTP 500 internal server error response, then exits the loop and closes the connection.

3. **Echo Back**:
   - For each message received, the function immediately sends it back to the client, effectively echoing the received message.
   - If sending the message fails, it logs the error and sends an HTTP 500 internal server error response, then exits the loop and closes the connection.

### Error Handling

The function robustly handles errors at each step of the WebSocket communication process:
- **Connection Upgrade Error**: Logs and notifies the client of an inability to upgrade the HTTP connection.
- **Read Message Error**: Logs and notifies the client if a message cannot be read from the WebSocket.
- **Write Message Error**: Logs and notifies the client if a message cannot be sent back through the WebSocket.

Each error results in closing the WebSocket connection to ensure that no faulty or half-open connections persist.

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
