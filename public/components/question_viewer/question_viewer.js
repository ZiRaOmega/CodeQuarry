let response_submit = document.getElementById("response_submit");
const response_description = document.getElementById("response_description");
const response_content = document.getElementById("response_content");
const error_message_response = document.getElementById(
  "error_message_response"
);

function sendResponse() {
  if (response_description.value == "" && response_content.value == "") {
    error_message_response.style.display = "block";
    error_message_response.innerText = "Veuillez remplir tous les champs.";
  } else if (response_description.value == "") {
    error_message_response.style.display = "block";
    error_message_response.innerText = "Veuillez remplir la description.";
  } else if (response_content.value == "") {
    error_message_response.style.display = "block";
    error_message_response.innerText = "Veuillez remplir le contenu.";
  } else {
    const question_id = document
      .getElementById("question_id")
      .getAttribute("question-id");
    const response_description = document.getElementById(
      "response_description"
    ).value;
    const response_content = document.getElementById("response_content").value;
    const response = {
      question_id: getUrlArgument("question_id"),
      description: response_description,
      content: response_content,
    };
    fetch("/api/responses", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        response: response,
        session_id: getCookie("session"),
      }),
    })
      .then((response) => response.json())
      .then((data) => {
        if (data.status === "success") {
          window.location.reload();
        }
      })
      .catch((error) => {
        console.error("Error:", error);
      });
  }
}

response_submit.addEventListener("click", function () {
  sendResponse();
});

function handleKeyPress(event) {
  if (event.keyCode === 13) {
    // Check if the key pressed is Enter
    event.preventDefault(); // Prevent the default action to stop from submitting the form
    sendResponse();
  }
}

response_description.addEventListener("keypress", handleKeyPress);
response_content.addEventListener("keypress", handleKeyPress);

function getUrlArgument(name) {
  const response_description = document.getElementById("response_description");
  const response_content = document.getElementById("response_content");
  response_description.value = "";
  response_content.value = "";
  const url = new URL(window.location.href);
  return url.searchParams.get(name);
}

const upvoteContainer = document.querySelectorAll(".upvote_container");
const downvoteContainer = document.querySelectorAll(".downvote_container");
const upvoteCount = document.querySelectorAll(".upvote_count");
const downvoteCount = document.querySelectorAll(".downvote_count");
const question_viewer__question__title = document.querySelector(
  ".question-viewer__question__title"
);

const question_viewer__question__description = document.querySelector(
  ".question-viewer__question__description"
);

const question_viewer__question__content = document.querySelector(
  ".question-viewer__question__content"
);

const question_viewer__answers = document.querySelector(
  ".question-viewer__answers"
);

const preDiv = document.createElement("pre");
const code = document.createElement("code");
preDiv.appendChild(code);

fetch("/api/questions?subjectId=all")
  .then((response) => response.json())
  .then((data) => {
    data.forEach((question) => {
      //get the question where the id is the same as the one in the url
      if (question.id == getUrlArgument("question_id")) {
        console.log(question);
        //when all is loaded

        setTimeout(() => {
          socket.send(
            JSON.stringify({
              type: "questionCompareUser",
              content: question.id,
              session_id: getCookie("session"),
            })
          );
        }, 500);

        question_viewer__question__title.innerText = question.title;
        question_viewer__question__description.innerText = question.description;
        code.textContent = question.content;
        question_viewer__question__content.appendChild(preDiv);
        document.querySelectorAll("pre code").forEach((block) => {
          // Apply Highlight.js
          hljs.highlightElement(block);
        });
        upvoteCount[0].setAttribute("data-question-id", question.id);
        downvoteCount[0].setAttribute("data-question-id", question.id);
        upvoteCount[0].textContent = question.upvotes;
        downvoteCount[0].textContent = question.downvotes;
        if (question.user_vote == "upvoted") {
          upvoteContainer[0].style.backgroundColor = "rgb(104, 195, 163)";
        } else if (question.user_vote == "downvoted") {
          downvoteContainer[0].style.backgroundColor = "rgb(196, 77, 86)";
        }
        upvoteContainer[0].onclick = function () {
          //if upvoteContainer backgroundColor is green then remove the color
          if (
            upvoteContainer[0].style.backgroundColor == "rgb(104, 195, 163)"
          ) {
            upvoteContainer[0].style.backgroundColor = "";
          } else {
            upvoteContainer[0].style.backgroundColor = "rgb(104, 195, 163)";
            if (
              downvoteContainer[0].style.backgroundColor == "rgb(196, 77, 86)"
            ) {
              downvoteContainer[0].style.backgroundColor = "";
            }
          }
          socket.send(
            JSON.stringify({
              type: "upvote",
              content: question.id,
              session_id: getCookie("session"),
            })
          );
        };

        downvoteContainer[0].onclick = function () {
          //if downvoteContainer backgroundColor is red then remove the color
          if (
            downvoteContainer[0].style.backgroundColor == "rgb(196, 77, 86)"
          ) {
            downvoteContainer[0].style.backgroundColor = "";
          } else {
            downvoteContainer[0].style.backgroundColor = "rgb(196, 77, 86)";
            if (
              upvoteContainer[0].style.backgroundColor == "rgb(104, 195, 163)"
            ) {
              upvoteContainer[0].style.backgroundColor = "";
            }
          }
          socket.send(
            JSON.stringify({
              type: "downvote",
              content: question.id,
              session_id: getCookie("session"),
            })
          );
        };

        if (question.responses != null) {
          //explain how does the sorting works

          const containBestAnswer = question.responses.some(
            (r) => r.best_answer == true
          );
          question.responses.sort((a, b) => {
            return (b.best_answer === true) - (a.best_answer === true);
          });

          question.responses.forEach((answer) => {
            const bestAnswerContainer = document.createElement("div");
            bestAnswerContainer.classList.add("best_answer_container");
            const bestAnswer = document.createElement("div");
            bestAnswer.classList.add("best_answer");
            bestAnswer.setAttribute("data-answer-id", answer.response_id);
            bestAnswer.innerText = "Best answer ✔";
            const question_viewer__answers__answer =
              document.createElement("div");
            const question_viewer__answers__answer__description =
              document.createElement("div");
            const question_viewer__answers__answer__content =
              document.createElement("div");
            const question_viewer__answers__answer__date =
              document.createElement("div");
            const creator_name = document.createElement("span");
            const creator_and_date_container = document.createElement("div");
            const pre = document.createElement("pre");
            const code = document.createElement("code");
            creator_and_date_container.classList.add(
              "creator_and_date_container"
            );
            creator_name.classList.add("creator_name");
            question_viewer__answers__answer__date.classList.add(
              "question-viewer__answers__answer__date"
            );
            const question_viewer__answers__answer__author =
              document.createElement("div");
            question_viewer__answers__answer__author.classList.add(
              "question-viewer__answers__answer__author"
            );
            question_viewer__answers__answer__author.innerText = "Réponse de ";
            creator_name.innerText = answer.student_name;
            question_viewer__answers__answer__content.classList.add(
              "question-viewer__answers__answer__content"
            );
            question_viewer__answers__answer__content.classList.add(
              "question-viewer__answers__answer__content"
            );
            question_viewer__answers__answer.classList.add(
              "question-viewer__answers__answer"
            );
            question_viewer__answers__answer__description.classList.add(
              "question-viewer__answers__answer__description"
            );

            question_viewer__answers__answer__date.innerText = `Publié le: ${new Date(
              question.creation_date
            ).toLocaleDateString()}`;

            question_viewer__answers__answer__description.innerText =
              answer.description;

            code.textContent = answer.content;

            question_viewer__answers__answer__author.appendChild(creator_name);
            pre.appendChild(code);
            question_viewer__answers__answer__content.appendChild(pre);
            question_viewer__answers__answer.appendChild(
              question_viewer__answers__answer__description
            );
            question_viewer__answers__answer.appendChild(
              question_viewer__answers__answer__content
            );
            question_viewer__answers.appendChild(
              question_viewer__answers__answer
            );
            question_viewer__answers__answer.appendChild(
              question_viewer__answers__answer__author
            );
            question_viewer__answers__answer.appendChild(
              creator_and_date_container
            );
            creator_and_date_container.appendChild(
              question_viewer__answers__answer__date
            );
            bestAnswerContainer.appendChild(bestAnswer);
            creator_and_date_container.appendChild(bestAnswerContainer);
            if (containBestAnswer) {
              if (answer.best_answer) {
                bestAnswer.style.display = "flex";
                bestAnswer.style.backgroundColor = "rgb(104, 195, 163)";
              } else {
                bestAnswer.style.display = "none";
              }
            } else {
              bestAnswer.style.display = "flex";
            }

            creator_and_date_container.appendChild(
              question_viewer__answers__answer__author
            );

            bestAnswer.onclick = function () {
              socket.send(
                JSON.stringify({
                  type: "bestAnswer",
                  content: {
                    answer_id: bestAnswer.getAttribute("data-answer-id"),
                    question_id: question.id.toString(),
                  },
                  session_id: getCookie("session"),
                })
              );
            };
            document.querySelectorAll("pre code").forEach((block) => {
              hljs.highlightElement(block);
            });
          });
        }
      }
    });
  })
  .catch((error) => {
    console.error("Error:", error);
  });
