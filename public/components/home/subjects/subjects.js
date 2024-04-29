let SubjectsList = [];
let QuestionsElementsList = [];
let ListElement;
const questionsList = document.getElementById("questionsList");
const returnButton = document.createElement("div");
document.addEventListener("DOMContentLoaded", function () {
  localStorage.removeItem("subjectId");
  localStorage.removeItem("subjectTitle");
  returnButton.id = "returnButton";
  const listElement = document.getElementById("subjectsList");
  const questionsList = document.getElementById("questionsList");

  // Create the "All" subjects item with title and description
  const allSubjectsItem = document.createElement("div");
  allSubjectsItem.classList.add("category_cards");

  // Create and append the title
  const allTitle = document.createElement("h2");
  allTitle.classList.add("category_title");
  allTitle.textContent = "All";
  allSubjectsItem.appendChild(allTitle);

  // Create and append the description
  const allDescription = document.createElement("p");
  allDescription.classList.add("category_description");
  allDescription.textContent =
    "Click here to view all questions across all subjects.";
  allSubjectsItem.appendChild(allDescription);

  // Append the "All" item to the list
  listElement.appendChild(allSubjectsItem);
  ListElement = listElement;
  // Event listener for the "All" subjects item
  allSubjectsItem.addEventListener("click", function () {
    localStorage.setItem("subjectId", "all");
    localStorage.setItem("subjectTitle", "All Subjects");
    listElement.style.display = "none"; // Hide the list
    returnButton.style.display = ""; // Show return button
    fetchQuestions("all"); // Fetch all questions
  });

  returnButton.addEventListener("click", function () {
    localStorage.removeItem("subjectId");
    localStorage.removeItem("subjectTitle");
    listElement.style.display = "";
    questionsList.style.display = "none";
    questionsList.innerHTML = ""; // Clear previous questions
  });

  fetch("/api/subjects")
    .then((response) => response.json())
    .then((subjects) => {
      const allQestionCountDiv = document.createElement("div");
      allQestionCountDiv.classList.add("question_count_all");
      let totalQuestions = 0; // For counting all questions
      SubjectsList = [];
      subjects.forEach((subject) => {
        SubjectsList.push(subject);
        totalQuestions += subject.questionCount; // Sum up all questions
        allQestionCountDiv.textContent = totalQuestions;
        allSubjectsItem.appendChild(allQestionCountDiv);
        const listItem = document.createElement("div");
        listItem.classList.add("category_cards");

        const questionCountDiv = document.createElement("div");
        questionCountDiv.classList.add("question_count");
        questionCountDiv.setAttribute("data-subject-id", subject.id);
        questionCountDiv.textContent = subject.questionCount; // Assuming 'subject.questionCount' is the number of questions
        listItem.appendChild(questionCountDiv);

        const title = document.createElement("h2");
        title.classList.add("category_title");
        title.textContent = subject.title;
        listItem.appendChild(title);

        const description = document.createElement("p");
        description.classList.add("category_description");
        description.textContent = subject.description;
        listItem.appendChild(description);

        listItem.addEventListener("click", function () {
          localStorage.setItem("subjectId", subject.id);
          localStorage.setItem("subjectTitle", subject.title);
          listElement.style.display = "none";
          fetchQuestions(subject.id);
        });

        listElement.appendChild(listItem);
      });
    });
});

window.fetchQuestions = function (subjectId) {
  fetch(`/api/questions?subjectId=${subjectId}`)
    .then((response) => response.json())
    .then((questions) => {
      createFilter(questions);
      create_questions(questions);
    });
};

function createFilter(questions) {
  questionsList.innerHTML = ""; // Clear previous questions
  const filterContainer = document.createElement("div");
  filterContainer.classList.add("filter_container");
  const questionFilter = document.createElement("div");
  questionFilter.classList.add("question_filter");
  const questionTrackerCount = document.createElement("div");
  questionTrackerCount.classList.add("question_tracker_count");
  const filterQuestions = document.createElement("div");
  filterQuestions.classList.add("filter_questions");
  const filterPopular = document.createElement("div");
  filterPopular.classList.add("filter_popular");
  filterPopular.textContent = "Croissant ↗";
  const filterUnpopular = document.createElement("div");
  filterUnpopular.classList.add("filter_unpopular");
  filterUnpopular.textContent = "Décroissant ↘";
  const filterNewest = document.createElement("div");
  const filterOldest = document.createElement("div");
  filterOldest.classList.add("filter_oldest");
  const filterNumberOfComments = document.createElement("div");
  filterNumberOfComments.classList.add("filter_number_of_comments");
  filterNewest.classList.add("filter_newest");
  filterNumberOfComments.textContent = "↑ Commentaires";
  filterNewest.textContent = "Recent";
  filterOldest.textContent = "Ancien";
  filterQuestions.appendChild(filterNumberOfComments);
  filterQuestions.appendChild(filterOldest);
  filterQuestions.appendChild(filterNewest);
  filterQuestions.appendChild(filterPopular);
  filterQuestions.appendChild(filterUnpopular);
  questionFilter.appendChild(questionTrackerCount);
  questionFilter.appendChild(filterQuestions);
  returnButton.textContent = "⬅";
  filterContainer.appendChild(returnButton);
  filterContainer.appendChild(questionFilter);
  questionsList.appendChild(filterContainer);
  if (questions == null) {
    questionTrackerCount.textContent = "0 question(s)";
    return;
  } else {
    questionTrackerCount.textContent = `${questions.length} question(s)`;
  }

  filterNumberOfComments.onclick = function () {
    console.log("sorting by number of comments", questions);
    //check if responses is null if yes set it to 0
    questions.forEach((question) => {
      if (question.responses == null) {
        question.responses = [];
      }
    });
    questions.sort((a, b) => b.responses.length - a.responses.length);
    questionsList.innerHTML = ""; // Clear previous questions
    createFilter(questions);
    create_questions(questions);
  };

  filterOldest.onclick = function () {
    console.log("sorting by date", questions);
    questions.sort(
      (a, b) => new Date(a.creation_date) - new Date(b.creation_date)
    );
    questionsList.innerHTML = ""; // Clear previous questions
    createFilter(questions);
    create_questions(questions);
  };

  filterNewest.onclick = function () {
    console.log("sorting by date", questions);
    questions.sort(
      (a, b) => new Date(b.creation_date) - new Date(a.creation_date)
    );
    questionsList.innerHTML = ""; // Clear previous questions
    createFilter(questions);
    create_questions(questions);
  };

  filterPopular.onclick = function () {
    console.log("sorting by upvotes", questions);
    questions.sort((a, b) => b.upvotes - a.upvotes);
    questionsList.innerHTML = ""; // Clear previous questions
    createFilter(questions);
    create_questions(questions);
  };

  filterUnpopular.onclick = function () {
    questions.sort((a, b) => a.upvotes - b.upvotes);
    questionsList.innerHTML = ""; // Clear previous questions
    createFilter(questions);
    create_questions(questions);
  };
}

function create_questions(questions) {
  // Clear previous questions
  if (questions != null)
    questions.forEach((question) => {
      const questionContainer = document.createElement("div");
      questionContainer.classList.add("question");
      const clickable_container = document.createElement("div");
      clickable_container.classList.add("clickable_container");
      // Add subject title tag
      const subjectTag = document.createElement("div");
      subjectTag.classList.add("subject_tag");
      subjectTag.textContent = question.subject_title;
      questionContainer.appendChild(subjectTag);

      const questionTitle = document.createElement("h3");
      questionTitle.classList.add("question_title");
      questionTitle.textContent = question.title;
      clickable_container.appendChild(questionTitle);

      const questionDescription = document.createElement("p");
      questionDescription.classList.add("question_description");
      questionDescription.textContent = question.description;
      clickable_container.appendChild(questionDescription);

      const questionContent = document.createElement("p");
      questionContent.classList.add("question_content");
      questionContent.textContent = question.content;
      const preDiv = document.createElement("pre");
      const code = document.createElement("code");
      preDiv.appendChild(code);
      code.innerHTML = question.content;
      clickable_container.appendChild(preDiv);

      const ContainerCreatorAndDate = document.createElement("div");
      ContainerCreatorAndDate.classList.add("creator_and_date_container");

      const questionDate = document.createElement("p");
      questionDate.classList.add("question_creation_date");
      questionDate.textContent = `Publié le: ${new Date(
        question.creation_date
      ).toLocaleDateString()}`;
      ContainerCreatorAndDate.appendChild(questionDate);

      const questionCreator = document.createElement("p");
      questionCreator.classList.add("question_creator");
      questionCreator.textContent = "Publié par";
      const creatorName = document.createElement("span");
      creatorName.textContent = question.creator;
      creatorName.classList.add("creator_name");
      questionCreator.appendChild(creatorName);
      const responsesCounter = document.createElement("p");
      responsesCounter.classList.add("responses_counter");
      if (Array.isArray(question.responses)) {
        // Set text content to the length of the responses array
        responsesCounter.textContent = `${question.responses.length} reponse(s)`;
      } else {
        // Handle cases where 'responses' might not be defined or not an array
        responsesCounter.textContent = "0 reponse(s)";
      }
      ContainerCreatorAndDate.appendChild(responsesCounter);
      ContainerCreatorAndDate.appendChild(questionCreator);
      clickable_container.appendChild(ContainerCreatorAndDate);
      questionContainer.appendChild(clickable_container);
      QuestionsElementsList.push(questionContainer);
      const voteContainer = document.createElement("div");
      voteContainer.classList.add("vote_container");
      const upvoteContainer = document.createElement("div");
      upvoteContainer.classList.add("upvote_container");
      const upvoteText = document.createElement("div");
      upvoteText.classList.add("upvote_text");
      upvoteText.textContent = "+";
      const upvoteCount = document.createElement("p");
      upvoteCount.classList.add("upvote_count");
      upvoteCount.setAttribute("data-question-id", question.id);
      upvoteCount.textContent = question.upvotes;
      upvoteContainer.appendChild(upvoteText);
      upvoteContainer.appendChild(upvoteCount);
      voteContainer.appendChild(upvoteContainer);
      const downvoteContainer = document.createElement("div");
      downvoteContainer.classList.add("downvote_container");
      console.log(question)
      if (question.user_vote == "upvoted") {
        upvoteContainer.style.backgroundColor = "green";
      } else if (question.user_vote == "downvoted") {
        downvoteContainer.style.backgroundColor = "red";
      }
      const downvoteText = document.createElement("div");
      downvoteText.classList.add("downvote_text");
      downvoteText.textContent = "-";
      const downvoteCount = document.createElement("p");
      downvoteCount.classList.add("downvote_count");
      downvoteCount.setAttribute("data-question-id", question.id);
      downvoteCount.textContent = question.downvotes;
      downvoteContainer.appendChild(downvoteText);
      downvoteContainer.appendChild(downvoteCount);
      voteContainer.appendChild(downvoteContainer);
      questionContainer.appendChild(voteContainer);
      questionsList.appendChild(questionContainer);
      let responseContainer = document.createElement("div");
      responseContainer.classList.add("response_container");
      clickable_container.addEventListener("click", () => {
        window.location.href = `https://${window.location.hostname}/question_viewer?question_id=${question.id}`;
      });
      upvoteContainer.onclick = function () {
        //if upvoteContainer backgroundColor is green then remove the color
        if (upvoteContainer.style.backgroundColor == "green") {
          upvoteContainer.style.backgroundColor = "";

        } else {
          upvoteContainer.style.backgroundColor = "green";
          if (downvoteContainer.style.backgroundColor == "red") {
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
      };

      downvoteContainer.onclick = function () {
        //if downvoteContainer backgroundColor is red then remove the color
        if (downvoteContainer.style.backgroundColor == "red") {
          downvoteContainer.style.backgroundColor = "";

        } else {
          downvoteContainer.style.backgroundColor = "red";
          if (upvoteContainer.style.backgroundColor == "green") {
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
      };
    });
  questionsList.style.display = ""; // Show the questions list
  checkHighlight();
  if (questions == null) {
    questionTrackerCount.textContent = "0 question(s)";
  }
}
