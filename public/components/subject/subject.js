let QuestionsElementsList = [];
const questionContainer = document.getElementById("questions_container");
const questionsList = document.getElementById("questions_list");
const return_btn = document.getElementById("return_button");

return_btn.onclick = () => {
  window.location.href = "/home";
};

document.addEventListener("DOMContentLoaded", function () {
  let subjectId = window.location.pathname.split("/")[2];

  (async () => {
    try {
      console.log(SubjectsList)
      if (SubjectsList.length==0){
        const response = await fetch("/api/subjects");
        const subjects = await response.json();
        SubjectsList = subjects;
      }
      //get question from SubjectsList by subjectId
      const subject = SubjectsList.find((subject) => subject.id == subjectId);
      console.log(subject)
      const questions = subject.questions;

      createFilter(questions);
      createQuestions(questions);
    } catch (error) {
      const errorH1 = document.createElement("h1");
      errorH1.textContent = "An error occured while fetching the questions";
      errorH1.style.color = "red";
      questionsList.appendChild(errorH1);
      console.error("There was a problem with your fetch operation:", error);
    }
  })();
});

function createFilter(questions) {
  questionTrackerCount = document.getElementById("question_tracker_count");

  if (questions == null) {
    questionTrackerCount.textContent = "0 question(s)";
    return;
  } else {
    questionTrackerCount.textContent = `${questions.length} question(s)`;
  }

  document.getElementById("filter_nbr_of_comments").onclick = () =>
    sortByNumberOfComments(questions);
  document.getElementById("filter_old").onclick = () =>
    sortOldestToNewest(questions);
  document.getElementById("filter_new").onclick = () =>
    sortNewestToOldest(questions);
  document.getElementById("filter_popular").onclick = () =>
    sortByUpvotes(questions);
  document.getElementById("filter_unpopular").onclick = () =>
    sortByDownvotes(questions);
}

function sortByNumberOfComments(questions) {
  questions.forEach((q) => (q.responses = q.responses || []));
  questions.sort((a, b) => b.responses.length - a.responses.length);
  refreshQuestionView(questions);
}

function sortOldestToNewest(questions) {
  questions.sort(
    (a, b) => new Date(a.creation_date) - new Date(b.creation_date)
  );
  refreshQuestionView(questions);
}

function sortNewestToOldest(questions) {
  questions.sort(
    (a, b) => new Date(b.creation_date) - new Date(a.creation_date)
  );
  refreshQuestionView(questions);
}

function sortByUpvotes(questions) {
  questions.sort((a, b) => b.upvotes - a.upvotes);
  refreshQuestionView(questions);
}

function sortByDownvotes(questions) {
  questions.sort((a, b) => a.upvotes - b.upvotes);
  refreshQuestionView(questions);
}

function refreshQuestionView(questions) {
  questionsList.innerHTML = ""; // Clear previous questions
  createFilter(questions);
  createQuestions(questions);
}

function createQuestions(questions) {
  questions.forEach((question) => {
    const questionElement = createQuestionElement(question);
    questionsList.appendChild(questionElement);
  });
  checkHighlight();
}

function createQuestionElement(question) {
  const questionElement = document.createElement("div");
  questionElement.classList.add("question");
  questionElement.setAttribute("data-question-id", question.id);
  const questionChecked = document.createElement("div");
  questionChecked.classList.add("question_checked");
  questionChecked.setAttribute("data-question-id", question.id);
  if (question.responses != null) {
    if (question.responses.some((r) => r.best_answer == true)) {
      questionChecked.style.display = "block";
    } else {
      questionChecked.style.display = "none";
    }
  }
  if (question.responses == null) {
    question.responses = [];
  }

  questionElement.innerHTML = htmlQuestionConstructor(question);
  questionElement.appendChild(questionChecked);

  // Highlight the code block
  // !!! TODO : use checkHighlight() function from detect_lang.js
  // make sure to check global script
  /* code = questionElement.querySelector("code");
    code = hljs.highlightElement(code); */

  // Upvote and downvote event listeners
  upvoteContainer = questionElement.querySelector(".upvote_container");
  downvoteContainer = questionElement.querySelector(".downvote_container");
  switchVoteColor(question, upvoteContainer, downvoteContainer);

  // Handle favori
  favori = questionElement.querySelector(".favori");
  manageFavorite(favori, question.id);

  // Clickable container redirect to question viewer
  clickable_container = questionElement
    .querySelector(".clickable_container")
    .addEventListener("click", () => {
      window.location.href = `https://${window.location.hostname}/question_viewer?question_id=${question.id}`;
    });

  return questionElement;
}

function manageFavorite(favori, questionId) {
  fetch("/api/favoris")
    .then((response) => response.json())
    .then((favoris) => {
      if (Array.isArray(favoris)) {
        console.log(favori);
        if (favoris.some((f) => f == questionId)) {
          favori.classList.add("favori_active");
          favori.textContent = "★";
          favori.style.backgroundColor = "gold";
        } else {
          favori.classList.remove("favori_active");
          favori.textContent = "☆";
          favori.style.backgroundColor = "";
        }
      } else {
        favori.classList.remove("favori_active");
        favori.textContent = "☆";
        favori.style.backgroundColor = "";
      }
    })
    .catch((error) => {
      console.error("Problem with favori fetch:", error);
    });

  favori.onclick = function () {
    AddFavori(questionId);
    if (favori.classList.contains("favori_active")) {
      favori.classList.remove("favori_active");
      favori.textContent = "☆";
      favori.style.backgroundColor = "";
    } else {
      favori.classList.add("favori_active");
      favori.textContent = "★";
      favori.style.backgroundColor = "gold";
    }
  };
}

function htmlQuestionConstructor(question) {
  return `
<div class="subject_tag">${question.subject_title}</div>
<div class="clickable_container">
    <h3 class="question_title">${question.title}</h3>
    <p class="question_description">${question.description}</p>
    <pre><code>${question.content}</code></pre>
    <div class="creator_and_date_container">
        <p class="question_creation_date">Publié le: ${new Date(
          question.creation_date
        ).toLocaleDateString()}</p>
        <p class="responses_counter">${question.responses.length} reponse(s)</p>
        <p class="question_creator">Publié par <span class="creator_name">${
          question.creator
        }</span></p>
    </div>
</div>
<div class="vote_container">
    <div class="upvote_container">
        <div class="upvote_text">+</div>
        <p class="upvote_count" data-question-id="${question.id}">${
    question.upvotes
  }</p>
    </div>
    <div class="downvote_container">
        <div class="downvote_text">-</div>
        <p class="downvote_count" data-question-id="${question.id}">${
    question.downvotes
  }</p>
    </div>
    <div class="favori" data-question-id="${question.id}">☆</div>
</div>
`;
}

function switchVoteColor(question, upvoteContainer, downvoteContainer) {
  if (question.user_vote == "upvoted") {
    upvoteContainer.style.backgroundColor = "rgb(104, 195, 163)";
  } else if (question.user_vote == "downvoted") {
    downvoteContainer.style.backgroundColor = "rgb(196, 77, 86)";
  }

  upvoteContainer.addEventListener("click", () => {
    if (upvoteContainer.style.backgroundColor == "rgb(104, 195, 163)") {
      upvoteContainer.style.backgroundColor = "";
    } else {
      upvoteContainer.style.backgroundColor = "rgb(104, 195, 163)";
      if (downvoteContainer.style.backgroundColor == "rgb(196, 77, 86)") {
        downvoteContainer.style.backgroundColor = "";
      }
    }

    socket.send(
      JSON.stringify({
        type: "upvote",
        content: question.id,
        session_id: getCookie("session"),
      })
    );
  });

  downvoteContainer.addEventListener("click", () => {
    if (downvoteContainer.style.backgroundColor == "rgb(196, 77, 86)") {
      downvoteContainer.style.backgroundColor = "";
    } else {
      downvoteContainer.style.backgroundColor = "rgb(196, 77, 86)";
      if (upvoteContainer.style.backgroundColor == "rgb(104, 195, 163)") {
        upvoteContainer.style.backgroundColor = "";
      }
    }

    socket.send(
      JSON.stringify({
        type: "downvote",
        content: question.id,
        session_id: getCookie("session"),
      })
    );
  });
}

// TODO:
/* 
const filterBestAnswer = document.createElement("div");
    filterBestAnswer.classList.add("filter_best_answer");
    filterBestAnswer.textContent = "Answered ✔";
*/
