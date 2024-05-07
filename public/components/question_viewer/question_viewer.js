(function () {
  var originalConsoleWarn = console.warn; // Save the original console.warn function

  console.warn = function () {
    var args = Array.prototype.slice.call(arguments);
    // Check if the first argument (the warning message) contains 'highlight.js'
    if (
      args.length > 0 &&
      typeof args[0] === "string" &&
      args[0].includes("highlight")
    ) {
      return; // Do not log Highlight.js warnings
    }
  };
})();

let response_submit = document.getElementById("response_submit");
const response_description = document.getElementById("response_description");
const response_content = document.getElementById("response_content");
const areYouSure = document.getElementsByClassName("areYouSure");
const areYouSureTitle = document.getElementsByClassName("areYouSureTitle");
const areYouSureText = document.getElementsByClassName("areYouSureText");
const Yes = document.getElementById("Yes");
const No = document.getElementById("No");
const best_answer_check = document.getElementsByClassName("best_answer_check");
//if best_answer_check stype display is flex, display none response_

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
const question_viewer = document.querySelector(".question-viewer__question")
const preDiv = document.createElement("pre");
const code = document.createElement("code");
preDiv.appendChild(code);

fetch("/api/questions?subjectId=all")
  .then((response) => response.json())
  .then((data) => {
    data.forEach((question) => {
      //get the question where the id is the same as the one in the url
      if (question.id == getUrlArgument("question_id")) {
        question_viewer.setAttribute("data-question-id",question.id)
        console.log(question);
        //when all is loaded
        let counter = 0;
        {
          let intervalId = setInterval(function () {
            if (counter >= 20) {
              clearInterval(intervalId); // Stop the interval if the counter is 10 or more
            } else {
              counter++;
              socket.send(
                JSON.stringify({
                  type: "questionCompareUser",
                  content: question.id,
                  session_id: getCookie("session"),
                })
              );
            }
          }, 150);
        }

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
            upvoteCount[0].textContent = parseInt(question.upvotes) - 1;
          } else {
            upvoteContainer[0].style.backgroundColor = "rgb(104, 195, 163)";
            upvoteCount[0].textContent = parseInt(question.upvotes) + 1;
            if (
              downvoteContainer[0].style.backgroundColor == "rgb(196, 77, 86)"
            ) {
              downvoteContainer[0].style.backgroundColor = "";
              downvoteCount[0].textContent = parseInt(question.downvotes);
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
            downvoteCount[0].textContent = parseInt(question.downvotes) - 1;
          } else {
            downvoteCount[0].textContent = parseInt(question.downvotes) + 1;
            downvoteContainer[0].style.backgroundColor = "rgb(196, 77, 86)";
            if (
              upvoteContainer[0].style.backgroundColor == "rgb(104, 195, 163)"
            ) {
              upvoteContainer[0].style.backgroundColor = "";
              upvoteCount[0].textContent = parseInt(question.upvotes);
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
        const addFavoriElement = document.createElement("div");
        addFavoriElement.classList.add("favori");
        addFavoriElement.setAttribute("data-question-id", question.id);
        addFavoriElement.textContent = "☆";
        fetch("/api/favoris")
          .then((response) => response.json())
          .then((favoris) => {
            console.log(Array.isArray(favoris));
            if (Array.isArray(favoris)) {
              if (favoris.some((f) => f == question.id)) {
                addFavoriElement.classList.add("favori_active");
                addFavoriElement.textContent = "★";
              } else {
                addFavoriElement.classList.remove("favori_active");
                addFavoriElement.textContent = "☆";
              }
            } else {
              addFavoriElement.classList.remove("favori_active");
              addFavoriElement.textContent = "☆";
            }
          });
        addFavoriElement.onclick = function () {
          AddFavori(question.id);
          if (addFavoriElement.classList.contains("favori_active")) {
            addFavoriElement.classList.remove("favori_active");
            addFavoriElement.textContent = "☆";
          } else {
            addFavoriElement.classList.add("favori_active");
            addFavoriElement.textContent = "★";
          }
        };
        const modifyButton = document.createElement("button");
        modifyButton.classList.add("modify_button");
        modifyButton.textContent = "Modify";
        modifyButton.onclick = function () {
          //if already exist just remove
          if (document.querySelector(".modify_container")){
            document.querySelector(".modify_container").remove()
            return
          }
          //create input with default value
          const question_title_input = document.createElement("input");
          question_title_input.setAttribute("type", "text");
          question_title_input.setAttribute("value", question.title);
          question_title_input.setAttribute("id", "question_title");
          question_title_input.classList.add("question_title_input");
          const question_description_input = document.createElement("textarea");
          question_description_input.innerText = question.description;
          question_description_input.setAttribute("id", "question_description");
          question_description_input.classList.add(
            "question_description_input"
          );
          const question_content_input = document.createElement("textarea");
          question_content_input.innerText = question.content;
          question_content_input.setAttribute("id", "question_content");
          question_content_input.classList.add("question_content_input");
          const modify_question = document.createElement("button");
          modify_question.classList.add("modify_question");
          modify_question.textContent = "Modify";
          const modifyContainer = document.createElement("div");
          modify_question.onclick = function () {
            ModifyQuestion();
            modifyContainer.remove()  
          };
          const cancel_button = document.createElement("button")
          cancel_button.innerText = "X"
          cancel_button.onclick = ()=>{
            modifyContainer.remove()
          }
          modifyContainer.appendChild(cancel_button)
          modifyContainer.classList.add("modify_container");
          modifyContainer.appendChild(question_title_input);
          modifyContainer.appendChild(question_description_input);
          modifyContainer.appendChild(question_content_input);
          modifyContainer.appendChild(modify_question);
          document
            .querySelector(".question-viewer__question")
            .appendChild(modifyContainer);
        };

        let voteContainer = document.querySelector(".vote_container");
        voteContainer.appendChild(addFavoriElement);
        voteContainer.appendChild(modifyButton);

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
            const bestAnswerCheck = document.createElement("div");
            bestAnswerCheck.classList.add("best_answer_check");
            bestAnswerCheck.setAttribute("data-answer-id", answer.response_id);
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
              "creator_and_date_container_answers"
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
            question_viewer__answers__answer__author.innerText = "Reponse from";
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
            question_viewer__answers__answer.setAttribute("data-answer-id",answer.response_id)
            question_viewer__answers__answer__description.classList.add(
              "question-viewer__answers__answer__description"
            );

            question_viewer__answers__answer__date.innerText = `Posted the: ${new Date(
              question.creation_date
            ).toLocaleDateString()}`;

            question_viewer__answers__answer__description.innerText =
              answer.description;

            code.textContent = answer.content;
            console.log(answer.is_author)
            
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
            question_viewer__answers__answer.appendChild(bestAnswerCheck);

            bestAnswerContainer.appendChild(bestAnswer);

            creator_and_date_container.appendChild(bestAnswerContainer);
            //upvote and downvote for response
            const vote_responseContainer = document.createElement("div");
            vote_responseContainer.classList.add("vote_response_container");
            const upvote_responseContainer = document.createElement("div");
            const downvote_responseContainer = document.createElement("div");
            const upvote_responseCount = document.createElement("div");
            const downvote_responseCount = document.createElement("div");
            upvote_responseContainer.classList.add("upvote_response_container");
            upvote_responseContainer.classList.add("upvote_container");
            downvote_responseContainer.classList.add(
              "downvote_response_container"
            );
            downvote_responseContainer.classList.add("downvote_container");
            upvote_responseCount.classList.add("upvote_response_count");
            downvote_responseCount.classList.add("downvote_response_count");
            upvote_responseCount.setAttribute(
              "data-answer-id",
              answer.response_id
            );
            downvote_responseCount.setAttribute(
              "data-answer-id",
              answer.response_id
            );
            console.log(answer);
            upvote_responseCount.textContent = "+ " + answer.upvotes;
            downvote_responseCount.textContent = "- " + answer.downvotes;
            upvote_responseContainer.appendChild(upvote_responseCount);
            downvote_responseContainer.appendChild(downvote_responseCount);
            vote_responseContainer.appendChild(upvote_responseContainer);
            vote_responseContainer.appendChild(downvote_responseContainer);
            question_viewer__answers__answer.appendChild(
              vote_responseContainer
            );
            if (answer.is_author){
              console.log("ghjklmù")
              const modify_button = document.createElement("button")
              modify_button.innerText = "Modify"
              vote_responseContainer.appendChild(modify_button)
              modify_button.addEventListener('click',()=>{
                if (document.querySelector(".modify_response_container")){
                  document.querySelector(".modify_response_container").remove()
                  return
                }
                console.log(modifyButton)
                const response_description_input = document.createElement("textarea");
                response_description_input.innerText = answer.description;
                response_description_input.setAttribute("id", "response_description");
                response_description_input.classList.add("response_description_input");
                const response_content_input = document.createElement("textarea");
                response_content_input.innerText = answer.content;
                response_content_input.setAttribute("id", "response_content");
                response_content_input.classList.add("response_content_input");
                const modify_response = document.createElement("button");
                modify_response.classList.add("modify_response");
                modify_response.textContent = "Modify";
                const modify_response_container = document.createElement("div");
                modify_response_container.classList.add("modify_response_container");
                modify_response.onclick = function () {
                  ModifyResponse(answer.response_id,response_content_input.value,response_description_input.value,answer.question_id);
                  modify_response_container.remove()
                };
                const cancel_button = document.createElement("button")
                cancel_button.innerText = "X"
                cancel_button.onclick = ()=>{
                  modify_response_container.remove()
                }
                modify_response_container.appendChild(cancel_button)
                modify_response_container.classList.add("modify_response_container");
                modify_response_container.appendChild(response_description_input);
                modify_response_container.appendChild(response_content_input);
                modify_response_container.appendChild(modify_response);
                document
                  .querySelector(".question-viewer__answers__answer")
                  .appendChild(modify_response_container);
              })

            }
            if (answer.user_vote == "upvoted") {
              upvote_responseContainer.style.backgroundColor =
                "rgb(104, 195, 163)";
              //upvote_responseCount.style.color = "white";
            } else if (answer.user_vote == "downvoted") {
              downvote_responseContainer.style.backgroundColor =
                "rgb(196, 77, 86)";
              //downvote_responseCount.style.color = "white";
            }

            upvote_responseContainer.onclick = function () {
              //if upvoteContainer backgroundColor is green then remove the color
              if (
                upvote_responseContainer.style.backgroundColor ==
                "rgb(104, 195, 163)"
              ) {
                upvote_responseContainer.style.backgroundColor = "";
                //upvote_responseCount.textContent = parseInt(answer.upvotes) - 1;
              } else {
                //upvote_responseCount.textContent = parseInt(answer.upvotes) + 1;
                upvote_responseContainer.style.backgroundColor =
                  "rgb(104, 195, 163)";
                if (
                  downvote_responseContainer.style.backgroundColor ==
                  "rgb(196, 77, 86)"
                ) {
                  downvote_responseContainer.style.backgroundColor = "";
                  /* downvote_responseCount.textContent = parseInt(
                    answer.downvotes
                  ); */
                }
              }
              socket.send(
                JSON.stringify({
                  type: "upvote_response",
                  content: answer.response_id,
                  session_id: getCookie("session"),
                })
              );
            };
            downvote_responseContainer.onclick = function () {
              //if downvote_responseContainer backgroundColor is red then remove the color
              if (
                downvote_responseContainer.style.backgroundColor ==
                "rgb(196, 77, 86)"
              ) {
                downvote_responseContainer.style.backgroundColor = "";
                downvoteCount.textContent = parseInt(answer.downvotes) - 1;
              } else {
                downvoteCount.textContent = parseInt(answer.downvotes) + 1;
                downvote_responseContainer.style.backgroundColor =
                  "rgb(196, 77, 86)";
                if (
                  upvote_responseContainer.style.backgroundColor ==
                  "rgb(104, 195, 163)"
                ) {
                  upvote_responseContainer.style.backgroundColor = "";
                  upvote_responseCount.textContent = parseInt(answer.upvotes);
                }
              }
              socket.send(
                JSON.stringify({
                  type: "downvote_response",
                  content: answer.response_id,
                  session_id: getCookie("session"),
                })
              );
            };
            let question_closed = document.getElementById("question_closed");
            if (containBestAnswer) {
              question_closed.style.display = "block";
              if (answer.best_answer) {
                bestAnswerCheck.style.display = "flex";
                bestAnswer.className = "best_answer best_answer_container";
                bestAnswer.style.display = "flex";
                bestAnswer.style.backgroundColor = "rgb(104, 195, 163)";
                document.getElementById("response_input").style.display =
                  "none";
              } else {
                bestAnswerCheck.style.display = "none";
                bestAnswer.style.display = "none";
              }
            } else {
              bestAnswer.style.display = "flex";
            }

            creator_and_date_container.appendChild(
              question_viewer__answers__answer__author
            );

            bestAnswer.onclick = function () {
              if (bestAnswer.className != "best_answer best_answer_container") {
                // Check both RGB and HEX
                areYouSure[0].style.display = "flex";

                Yes.onclick = function () {
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
                  areYouSure[0].style.display = "none";
                  document.getElementById("response_input").style.display =
                    "none";
                };

                No.onclick = function () {
                  areYouSure[0].style.display = "none";
                };
              } else {
                bestAnswer.className = "best_answer";
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
                document.getElementById("response_input").style.display =
                  "flex";
                question_closed.style.display = "none";
              }
            };

            document.querySelectorAll("pre code").forEach((block) => {
              hljs.highlightElement(block);
            });
          });
        }
      }
    });
  });

function ModifyQuestion() {
  const question_id = getUrlArgument("question_id");
  const question_title = document.getElementById("question_title").value;
  const question_description = document.getElementById(
    "question_description"
  ).value;
  const question_content = document.getElementById("question_content").value;

  socket.send(
    JSON.stringify({
      type: "modify_question",
      content: {
        question_id: question_id,
        title: question_title,
        description: question_description,
        content: question_content,
      },
      session_id: getCookie("session"),
    })
  );
}
