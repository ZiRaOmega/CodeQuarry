/* // {subject: "Math", marks:[
let input = document.getElementById("search-bar-input");
input.addEventListener("keyup", function () {
  if (QuestionsElementsList.length==0){
    fetchQuestions("all")
  }
  console.log("searching...");
  //if there is no subjectId inside localstorage
  if (
    !localStorage.getItem("subjectId") ||
    localStorage.getItem("subjectId") === ""
  ) {
    localStorage.setItem("subjectId", "all");
    localStorage.setItem("subjectTitle", "All Subjects");
    ListElement.style.display = "none"; // Hide the list
    fetchQuestions("all"); // Fetch all questions
  }
  const search = input.value.toLowerCase();
  for (let i = 0; i < QuestionsElementsList.length; i++) {
    const questionElement = QuestionsElementsList[i];
    const question_title =
      questionElement.querySelector(".question_title").innerText;
    const question_content =
      questionElement.querySelector("pre code").innerText;
    const question_description = questionElement.querySelector(
      ".question_description"
    ).innerText;
    if (
      question_title.toLowerCase().includes(search) ||
      question_content.toLowerCase().includes(search) ||
      question_description.toLowerCase().includes(search)
    ) {
      questionElement.style.display = "block";
    } else {
      questionElement.style.display = "none";
    }
  }
});
document.getElementById("home").addEventListener("click", function () {
  localStorage.removeItem("subjectId");
  localStorage.removeItem("subjectTitle");
  ListElement.style.display = "";
}); */

// Same but display results in a scroll dynamic way
let input = document.getElementById("search-bar-input");
let search_results = document.getElementById("search_results");
//search_results
input.addEventListener("keyup", function () {
  const search = input.value.toLowerCase();
  search_results.innerHTML = "";
  for (let i = 0; i < potential_results.length; i++) {
    const question = potential_results[i];
    const question_title = question.title;
    const question_content = question.content;
    const question_description = question.description;
    const question_subject = question.subject_title;
    if (search === "") {
      search_results.innerHTML = "";
    } else if (
      question_title.toLowerCase().includes(search) ||
      question_content.toLowerCase().includes(search) ||
      question_description.toLowerCase().includes(search) ||
      question_subject.toLowerCase().includes(search)
    ) {
      search_results.innerHTML += `
      <div class="subject_tag">${question.subject_title}</div>
      <div class="clickable_container">
          <h3 class="question_title">${question.title}</h3>
          <p class="question_description">${question.description}</p>
          <pre><code>${question.content}</code></pre>
      </div>
      `;
    }
    // highlight search results
    checkHighlight();
  }
});

let potential_results = [];

document.addEventListener("DOMContentLoaded", function () {
  (async () => {
    try {
      const response = await fetch(`/api/questions?subjectId=all`);
      const questions = await response.json();
      console.log(questions);
      potential_results = questions;

    } catch (error) {
      console.error("There was a problem with your fetch operation:", error);
    }
  }
  )();
}
);

