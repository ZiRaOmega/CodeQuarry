var socket;
const getCookie = (name) => {
  const value = `; ${document.cookie}`;
  const parts = value.split(`; ${name}=`);
  if (parts.length === 2) return parts.pop().split(";").shift();
};
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

    // Send a message to the server to check the session
    let message = {
      type: "session",
      content: getCookie("session"),
    };
    socket.send(JSON.stringify(message));
    //socket.send("Hey there from client");
  };

  // Define the onmessage function to be called when a message is received from the server
  socket.onmessage = function (event) {
    let msg = JSON.parse(event.data);
    switch (msg.type) {
      case "session":
        if (msg.content == "expired" || msg.content == "empty") {
          console.log(`Session ${msg.content}, redirecting to login page`);
          if (window.location.pathname != "/") window.location.href = "/";
        } else if (msg.content == "valid") {
          // !!! TODO show the website if the session is still valid
          if (window.location.pathname == "/")
            window.location.pathname = "/home";
          console.log("Session still valid");
        }
      case "voteUpdate":
        handleVoteUpdate(msg.content);
        break;
      case "postCreated":
        updateQuestionCount(msg.content); // Implement this function to update the UI
        if (localStorage.getItem("subjectId") == msg.content.id) {
          fetchQuestions(msg.content.id); // Implement this function to fetch and display questions
        }
        break;
      case "response":
        console.log("Response received from server:", msg.content);
        const answer_description = document.createElement("span");
        const creator = document.createElement("div");
        creator.classList.add("creator_name");
        answer_description.classList.add(
          "question-viewer__answers__answer__description"
        );
        answer_description.textContent = msg.content.description;
        const answerContainer = document.createElement("div");
        const creator_and_date_container = document.createElement("div");
        creator_and_date_container.classList.add(
          "creator_and_date_container_answers"
        );
        answerContainer.classList.add("question-viewer__answers__answer");

        // Content wrapper
        const answerContent = document.createElement("div");
        answerContent.classList.add(
          "question-viewer__answers__answer__content"
        );
        const pre = document.createElement("pre");
        const code = document.createElement("code");
        code.textContent = msg.content.content;
        pre.appendChild(code);
        answerContent.appendChild(pre);

        // Author info
        const answerAuthor = document.createElement("div");
        answerAuthor.classList.add("question-viewer__answers__answer__author");
        answerAuthor.textContent = "Réponse de";
        creator.textContent = msg.content.student_name;
        answerAuthor.appendChild(creator);

        // Date of the response
        const answerDate = document.createElement("div");
        answerDate.classList.add("question-viewer__answers__answer__date");
        answerDate.textContent = `Publié le: ${new Date(
          msg.content.creation_date
        ).toLocaleDateString()}`;

        // Best answer toggle
        const bestAnswerContainer = document.createElement("div");
        bestAnswerContainer.classList.add("best_answer_container");
        const bestAnswer = document.createElement("div");
        bestAnswer.classList.add("best_answer");
        bestAnswer.textContent = "Best answer ✔";
        bestAnswer.setAttribute("data-answer-id", msg.content.responseId);

        // Appending everything to the main container
        answerContainer.appendChild(answer_description);
        answerContainer.appendChild(answerContent);
        creator_and_date_container.appendChild(answerDate);
        creator_and_date_container.appendChild(answerAuthor);
        answerContainer.appendChild(creator_and_date_container);

        // Check if best answer
        if (msg.content.isBestAnswer) {
          bestAnswer.style.backgroundColor = "rgb(104, 195, 163)";
        }

        bestAnswer.onclick = function () {
          // Send WebSocket message to toggle best answer status
          socket.send(
            JSON.stringify({
              type: "bestAnswer",
              content: {
                answer_id: bestAnswer.getAttribute("data-answer-id"),
                question_id: msg.content.questionId,
              },
              session_id: getCookie("session"),
            })
          );
        };

        bestAnswerContainer.appendChild(bestAnswer);
        creator_and_date_container.appendChild(bestAnswerContainer);
        const answers = document.querySelector(".question-viewer__answers");
        answers.appendChild(answerContainer);

        checkHighlight();
        break;

      case "postDeleted":
        const question = document.querySelector(
          `.question[data-question-id="${msg.content}"]`
        );
        if (question) {
          question.remove();
          let subjectid = localStorage.getItem("subjectId");
          fetchQuestions(subjectid);
        }
        break;
      case "questionCompareUser":
        ItsMyQuestion(msg.content);
        break;
      case "bestAnswer":
        const questionID = msg.content.question_id;
        console.log("Best answer received from server:", msg.content);

        let answersBtn = document.querySelectorAll(".best_answer");
        let best_answer_check = document.querySelectorAll(".best_answer_check");
        let question_checked = document.querySelectorAll(".question_checked");
        question_checked.forEach((check) => {
          console.log(check.getAttribute("data-question-id"));
          console.log(questionID);
          if (check.getAttribute("data-question-id") == questionID.toString()) {
            if (check.style.display == "block") {
              check.style.display = "none";
            } else {
              check.style.display = "block";
            }
          }
        });

        best_answer_check.forEach((check) => {
          if (msg.content.question_best_answer == -1) {
            check.style.display = "none";
          } else if (
            check.getAttribute("data-answer-id") ==
            msg.content.question_best_answer.toString()
          ) {
            check.style.display = "flex";
          }
        });
        answersBtn.forEach((element) => {
          if (msg.content.question_best_answer == -1) {
            element.style.display = "flex";
            element.style.backgroundColor = "";
          } else if (
            element.dataset.answerId ==
            msg.content.question_best_answer.toString()
          ) {
            element.style.display = "flex";
            element.style.backgroundColor = "rgb(104, 195, 163)";
            element.classList.add("best_answer_container");
          } else {
            element.style.display = "none";
          }
        });
        break;
      case "XP":
        let xp = document.querySelector(".xp");
        xp.textContent = msg.content+" XP";
        break;
    }
    // Log the message received from the server
    console.log(`[message] Data received from server: ${event.data}`);
  };

  // Define the onclose function to be called when the WebSocket connection is closed
  socket.onclose = function (event) {
    // Check if the connection was closed cleanly
    if (event.wasClean) {
      // Log a message indicating the connection was closed cleanly
      console.log(
        `[close] Connection closed cleanly, code=${event.code} reason=${event.reason}`
      );
    } else {
      // Log an error message indicating the connection was not closed cleanly
      console.error("[close] Connection died");
    }
  };

  // Define the onerror function to be called when an error occurs with the WebSocket connection
  socket.onerror = function (error) {
    // Log an error message indicating the error that occurred
    console.error(`[error] ${error.message}`);
  };
});

function handleVoteUpdate(data) {
  // Assuming 'data' contains { questionId: 123, upvotes: 10, downvotes: 5 }
  const upvoteCountElement = document.querySelector(
    `.upvote_count[data-question-id="${data.question_id}"]`
  );
  const downvoteCountElement = document.querySelector(
    `.downvote_count[data-question-id="${data.question_id}"]`
  );

  if (upvoteCountElement) {
    upvoteCountElement.textContent = data.upvote;
  }
  if (downvoteCountElement) {
    downvoteCountElement.textContent = data.downvote;
  }
}

function updateQuestionCount(subject) {
  // Find the element displaying the question count and update it
  const questionCountDiv = document.querySelector(
    `.question_count[data-subject-id="${subject.id}"]`
  );
  if (questionCountDiv) {
    questionCountDiv.textContent = subject.questionCount;
  }

  // Update the total questions display if needed
  const totalQuestionsDiv = document.querySelector(".question_count_all");
  if (totalQuestionsDiv) {
    let totalQuestions = parseInt(totalQuestionsDiv.textContent, 10) || 0;
    totalQuestions++; // Increment since a new question was added
    totalQuestionsDiv.textContent = totalQuestions;
  }
}

function ItsMyQuestion(bool) {
  const best_answer = document.querySelectorAll(".best_answer_container");
  best_answer.forEach((element) => {
    element.style.display = bool ? "flex" : "none";
  });
}
