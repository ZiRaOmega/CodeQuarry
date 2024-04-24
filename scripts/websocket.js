var socket;
$(document).ready(function () {
    //Websocket connection
    var Current_Hostname = document.location.hostname
    socket = new WebSocket(`wss://${Current_Hostname}/ws`);
    socket.onopen = function (e) {
        console.log("[open] Connection established");
        console.log("Sending to server");
        socket.send("Hey there from client");
    };

    socket.onmessage = function (event) {
        console.log(`[message] Data received from server: ${event.data}`);
    };

    socket.onclose = function (event) {
        if (event.wasClean) {
            console.log(`[close] Connection closed cleanly, code=${event.code} reason=${event.reason}`);
        } else {
            // e.g. server process killed or network down
            // event.code is usually 1006 in this case
            console.error('[close] Connection died');
        }
    };

    socket.onerror = function (error) {
        console.error(`[error] ${error.message}`);
    };
  });