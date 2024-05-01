
let SubjectsList = [];
let QuestionsElementsList = [];
let ListElement;
const questionsList = document.getElementById("questionsList");
const returnButton = createReturnButton();

document.addEventListener("DOMContentLoaded", function () {
    initializeLocalStorage();
    ListElement = document.getElementById("subjectsList");

    const allSubjectsItem = createAllSubjectsItem();
    ListElement.appendChild(allSubjectsItem);
    addAllSubjectsClickListener(allSubjectsItem, ListElement, returnButton);

    returnButton.addEventListener("click", function () {
        resetSubjectSelection(ListElement, questionsList);
    });

    fetchSubjects(allSubjectsItem, ListElement);
});

function initializeLocalStorage() {
    localStorage.removeItem("subjectId");
    localStorage.removeItem("subjectTitle");
}

function createReturnButton() {
    const button = document.createElement("div");
    button.id = "returnButton";
    return button;
}

function createAllSubjectsItem() {
    const item = document.createElement("div");
    item.classList.add("category_cards");

    const title = document.createElement("h2");
    title.classList.add("category_title");
    title.textContent = "All";
    item.appendChild(title);

    const description = document.createElement("p");
    description.classList.add("category_description");
    description.textContent = "Click here to view all questions across all subjects.";
    item.appendChild(description);

    return item;
}

function addAllSubjectsClickListener(allSubjectsItem, listElement, returnButton) {
    allSubjectsItem.addEventListener("click", function () {
        localStorage.setItem("subjectId", "all");
        localStorage.setItem("subjectTitle", "All Subjects");
        listElement.style.display = "none";
        returnButton.style.display = "";
        fetchQuestions("all");
    });
}

function resetSubjectSelection(listElement, questionsList) {
    localStorage.removeItem("subjectId");
    localStorage.removeItem("subjectTitle");
    listElement.style.display = "";
    questionsList.style.display = "none";
    questionsList.innerHTML = ""; // Clear previous questions
}

function fetchSubjects(allSubjectsItem, listElement) {
    fetch("/api/subjects")
        .then((response) => response.json())
        .then((subjects) => {
            const allQestionCountDiv = document.createElement("div");
            allQestionCountDiv.classList.add("question_count_all");
            let totalQuestions = 0;

            SubjectsList = [];
            subjects.forEach((subject) => {
                SubjectsList.push(subject);
                totalQuestions += subject.questionCount;
                allQestionCountDiv.textContent = totalQuestions;
                allSubjectsItem.appendChild(allQestionCountDiv);

                const listItem = createSubjectItem(subject);
                addSubjectClickListener(listItem, subject, listElement);
                listElement.appendChild(listItem);
            });
        });
}

function createSubjectItem(subject) {
    const listItem = document.createElement("div");
    listItem.classList.add("category_cards");

    const questionCountDiv = document.createElement("div");
    questionCountDiv.classList.add("question_count");
    questionCountDiv.setAttribute("data-subject-id", subject.id);
    questionCountDiv.textContent = subject.questionCount;
    listItem.appendChild(questionCountDiv);

    const title = document.createElement("h2");
    title.classList.add("category_title");
    title.textContent = subject.title;
    listItem.appendChild(title);

    const description = document.createElement("p");
    description.classList.add("category_description");
    description.textContent = subject.description;
    listItem.appendChild(description);

    return listItem;
}

function addSubjectClickListener(listItem, subject, listElement) {
    listItem.addEventListener("click", function () {
        localStorage.setItem("subjectId", subject.id);
        localStorage.setItem("subjectTitle", subject.title);
        listElement.style.display = "none";
        fetchQuestions(subject.id);
    });
}

window.fetchQuestions = function (subjectId) {
    fetch(`/api/questions?subjectId=${subjectId}`)
        .then((response) => response.json())
        .then((questions) => {
            createFilter(questions);
            createQuestions(questions);
        });
};

function createFilter(questions) {
    questionsList.innerHTML = ""; // Clear previous questions
    const filterContainer = document.createElement("div");
    filterContainer.classList.add("filter_container");

    const questionFilter = createQuestionFilter(questions);
    returnButton.textContent = "⬅";
    filterContainer.appendChild(returnButton);
    filterContainer.appendChild(questionFilter);
    questionsList.appendChild(filterContainer);
}

function createQuestionFilter(questions) {
    const questionFilter = document.createElement("div");
    questionFilter.classList.add("question_filter");

    const questionTrackerCount = document.createElement("div");
    questionTrackerCount.classList.add("question_tracker_count");

    const filterQuestions = document.createElement("div");
    filterQuestions.classList.add("filter_questions");

    const filters = createFilterElements();
    filters.forEach(filter => filterQuestions.appendChild(filter));

    questionFilter.appendChild(questionTrackerCount);
    questionFilter.appendChild(filterQuestions);

    updateQuestionTrackerCount(questions, questionTrackerCount);

    filters[0].onclick = () => sortByNumberOfComments(questions);
    filters[1].onclick = () => sortOldestToNewest(questions);
    filters[2].onclick = () => sortByBestAnswer(questions);
    filters[3].onclick = () => sortNewestToOldest(questions);
    filters[4].onclick = () => sortByUpvotes(questions);
    filters[5].onclick = () => sortByDownvotes(questions);
  console.log(filters)
    return questionFilter;
}

function createFilterElements() {
    const filterPopular = document.createElement("div");
    filterPopular.classList.add("filter_popular");
    filterPopular.textContent = "Upvotes ↗";

    const filterUnpopular = document.createElement("div");
    filterUnpopular.classList.add("filter_unpopular");
    filterUnpopular.textContent = "Upvotes ↘";

    const filterNewest = document.createElement("div");
    filterNewest.classList.add("filter_newest");
    filterNewest.textContent = "Newest";

    const filterOldest = document.createElement("div");
    filterOldest.classList.add("filter_oldest");
    filterOldest.textContent = "Oldest";

    const filterBestAnswer = document.createElement("div");
    filterBestAnswer.classList.add("filter_best_answer");
    filterBestAnswer.textContent = "Answered ✔";

    const filterNumberOfComments = document.createElement("div");
    filterNumberOfComments.classList.add("filter_number_of_comments");
    filterNumberOfComments.textContent = "↑ Comments";

    return [filterNumberOfComments, filterOldest, filterBestAnswer, filterNewest, filterPopular, filterUnpopular];
}

function updateQuestionTrackerCount(questions, tracker) {
    tracker.textContent = questions ? `${questions.length} question(s)` : "0 question(s)";
}

function sortByNumberOfComments(questions) {
    questions.forEach(q => q.responses = q.responses || []);
    questions.sort((a, b) => b.responses.length - a.responses.length);
    refreshQuestionView(questions);
}

function sortOldestToNewest(questions) {
    questions.sort((a, b) => new Date(a.creation_date) - new Date(b.creation_date));
    refreshQuestionView(questions);
}

function sortNewestToOldest(questions) {
    questions.sort((a, b) => new Date(b.creation_date) - new Date(a.creation_date));
    refreshQuestionView(questions);
}

function sortByBestAnswer(questions) {
  questions.sort((a, b) => {
      a.responses = a.responses || [];
      b.responses = b.responses || [];
      return b.responses.filter(r => r.best_answer==true).length - a.responses.filter(r => r.best_answer==true).length;
  });
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
  // Clear previous questions
  if (questions != null)
    questions.forEach((question) => {
      const questionContainer = document.createElement("div");
      questionContainer.classList.add("question");
      const questionChecked = document.createElement("div");
      questionChecked.classList.add("question_checked");
      questionChecked.setAttribute("data-question-id", question.id);
      questionContainer.appendChild(questionChecked);
      if (question.responses != null) {
        if (question.responses.some((r) => r.best_answer == true)) {
          questionChecked.style.display = "block";
        } else {
          questionChecked.style.display = "none";
        }
      }
      questionContainer.setAttribute("data-question-id", question.id);
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
      code.textContent = question.content;
      document.querySelectorAll("pre code").forEach((block) => {
        hljs.highlightElement(block);
      });
      clickable_container.appendChild(preDiv);
      const ContainerCreatorAndDate = document.createElement("div");
      ContainerCreatorAndDate.classList.add("creator_and_date_container");

      const questionDate = document.createElement("p");
      questionDate.classList.add("question_creation_date");
      questionDate.textContent = `Posted the: ${new Date(
        question.creation_date
      ).toLocaleDateString()}`;
      ContainerCreatorAndDate.appendChild(questionDate);
      const questionCreator = document.createElement("p");
      questionCreator.classList.add("question_creator");
      questionCreator.textContent = "Created by";
      const creatorName = document.createElement("span");
      creatorName.textContent = question.creator;
      creatorName.classList.add("creator_name");
      questionCreator.appendChild(creatorName);
      const responsesCounter = document.createElement("p");
      responsesCounter.classList.add("responses_counter");
      if (Array.isArray(question.responses)) {
        // Set text content to the length of the responses array
        responsesCounter.textContent = `${question.responses.length} response(s)`;
      } else {
        // Handle cases where 'responses' might not be defined or not an array
        responsesCounter.textContent = "0 response(s)";
      }
      ContainerCreatorAndDate.appendChild(responsesCounter);
      ContainerCreatorAndDate.appendChild(questionCreator);
      clickable_container.appendChild(ContainerCreatorAndDate);
      questionContainer.appendChild(clickable_container);
      QuestionsElementsList.push(questionContainer);
      const voteContainer = document.createElement("div");
      voteContainer.classList.add("vote_container");
      const addFavoriElement = document.createElement("div");
      addFavoriElement.classList.add("favori");
      addFavoriElement.setAttribute("data-question-id", question.id);
      addFavoriElement.textContent = "☆";
      fetch("/api/favoris")
        .then((response) => response.json())
        .then((favoris) => {
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
      voteContainer.appendChild(addFavoriElement);
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
      console.log(question);
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
      if (question.user_vote == "upvoted") {
        upvoteContainer.style.backgroundColor = "rgb(104, 195, 163)";
      } else if (question.user_vote == "downvoted") {
        downvoteContainer.style.backgroundColor = "rgb(196, 77, 86)";
      }
      let responseContainer = document.createElement("div");
      responseContainer.classList.add("response_container");
      clickable_container.addEventListener("click", () => {
        window.location.href = `https://${window.location.hostname}/question_viewer?question_id=${question.id}`;
      });
      upvoteContainer.onclick = function () {
        //if upvoteContainer backgroundColor is green then remove the color
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
      };

      downvoteContainer.onclick = function () {
        //if downvoteContainer backgroundColor is red then remove the color
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
      };
    });
  questionsList.style.display = ""; // Show the questions list
  checkHighlight();
  if (questions == null) {
    let questionTrackerCount = document.querySelector(
      ".question_tracker_count"
    );
    questionTrackerCount.textContent = "0 question(s)";
  }
}
