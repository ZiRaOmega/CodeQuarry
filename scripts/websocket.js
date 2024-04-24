// Declare a variable 'socket' to store the WebSocket connection object
var socket;

// Use jQuery's document ready function to ensure the DOM is fully loaded
$(document).ready(function () {
    // Create a WebSocket connection to the current hostname with the '/ws' endpoint
    socket = new WebSocket(`wss://${document.location.hostname}/ws`);

    // Define the onopen function to be called when the WebSocket connection is established
    socket.onopen = function (e) {
        // Log a message indicating the connection is established
        console.log("[open] Connection established");

        // Log a message indicating a message is being sent to the server
        console.log("Sending to server");

        // Send a message to the server
        socket.send("Hey there from client");
    };

    // Define the onmessage function to be called when a message is received from the server
    socket.onmessage = function (event) {
        // Log the message received from the server
        console.log(`[message] Data received from server: ${event.data}`);
    };

    // Define the onclose function to be called when the WebSocket connection is closed
    socket.onclose = function (event) {
        // Check if the connection was closed cleanly
        if (event.wasClean) {
            // Log a message indicating the connection was closed cleanly
            console.log(`[close] Connection closed cleanly, code=${event.code} reason=${event.reason}`);
        } else {
            // Log an error message indicating the connection was not closed cleanly
            console.error('[close] Connection died');
        }
    };

    // Define the onerror function to be called when an error occurs with the WebSocket connection
    socket.onerror = function (error) {
        // Log an error message indicating the error that occurred
        console.error(`[error] ${error.message}`);
    };
});