// Handling WebSocket messages for votes
socket.onmessage = function (event) {
  let msg = JSON.parse(event.data);
  if (msg.type === "voteUpdate") {
    console.log("Vote updated for question:", msg.content.questionId);
    // Update UI accordingly
  }
};

