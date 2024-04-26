function setupVoteButtons(questionId) {
  const upvoteButton = document.getElementsByClassName("upvote_container");
  upvoteButton.array.forEach((element) => {
    element.onclick = function () {
      console.log("Upvote button clicked for question:", questionId);
      socket.send(JSON.stringify({ type: "upvote", content: questionId }));
    };
  });

  const downvoteButton = document.getElementsByClassName("downvote_container");
  downvoteButton.onclick = function () {
    socket.send(JSON.stringify({ type: "downvote", content: questionId }));
  };

  return { upvoteButton, downvoteButton };
}

// Handling WebSocket messages for votes
socket.onmessage = function (event) {
  let msg = JSON.parse(event.data);
  if (msg.type === "voteUpdate") {
    console.log("Vote updated for question:", msg.content.questionId);
    // Update UI accordingly
  }
};
