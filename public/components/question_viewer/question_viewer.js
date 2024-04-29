let response_submit = document.getElementById("response_submit");
response_submit.addEventListener("click", function () {
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
      console.log(data);
      if (data.status === "success") {
        window.location.reload();
      }
    })
    .catch((error) => {
      console.error("Error:", error);
    });
});

function getUrlArgument(name) {
  const url = new URL(window.location.href);
  return url.searchParams.get(name);
}

const upvoteContainer = document.querySelectorAll(".upvote_container");
const downvoteContainer = document.querySelectorAll(".downvote_container");
const upvoteCount = document.querySelectorAll(".upvote_count");
const downvoteCount = document.querySelectorAll(".downvote_count");

fetch("/api/questions?subjectId=all")
  .then((response) => response.json())
  .then((data) => {
    data.forEach((question) => {
      upvoteCount[0].setAttribute("data-question-id", question.id);
      downvoteCount[0].setAttribute("data-question-id", question.id);
      upvoteCount[0].textContent = question.upvotes;
      downvoteCount[0].textContent = question.downvotes;
      //get the question where the id is the same as the one in the url
      if (question.id == getUrlArgument("question_id")) {
        if (question.user_vote == "upvoted") {
          upvoteContainer[0].style.backgroundColor = "rgb(104, 195, 163)";
        } else if (question.user_vote == "downvoted") {
          downvoteContainer[0].style.backgroundColor = "rgb(196, 77, 86)";
        }

        upvoteContainer[0].onclick = function () {
         
          console.log(question)
          //if upvoteContainer backgroundColor is green then remove the color
          if (
            upvoteContainer[0].style.backgroundColor == "rgb(104, 195, 163)"
          ) {
            upvoteContainer[0].style.backgroundColor = "";
            upvoteCount[0].textContent = parseInt(question.upvotes)-1;
          } else {
            upvoteContainer[0].style.backgroundColor = "rgb(104, 195, 163)";
            upvoteCount[0].textContent = parseInt(question.upvotes)+1;
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
          console.log("upvoted");
        };

        downvoteContainer[0].onclick = function () {
          if (downvoteContainer[0].style.backgroundColor == "rgb(196, 77, 86)") {
            }else{
              
           }   
          
          //if downvoteContainer backgroundColor is red then remove the color
          if (
            downvoteContainer[0].style.backgroundColor == "rgb(196, 77, 86)"
          ) {
            downvoteContainer[0].style.backgroundColor = "";
            downvoteCount[0].textContent = parseInt(question.downvotes)-1;
           
          } else {
            
           downvoteCount[0].textContent = parseInt(question.downvotes)+1;
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
      }
    });
  })
  .catch((error) => {
    console.error("Error:", error);
  });
