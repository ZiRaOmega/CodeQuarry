document.addEventListener("DOMContentLoaded", function () {
  const listElement = document.getElementById("subjectsList");
  const questionsList = document.getElementById("questionsList"); // Now directly referencing the static HTML element
  const returnButton = document.getElementById("returnButton");
  returnButton.style.display = "none";

  returnButton.addEventListener("click", function () {
    returnButton.style.display = "none";
    localStorage.removeItem("subjectId");
    localStorage.removeItem("subjectTitle");
    listElement.style.display = "";
    questionsList.style.display = "none";
    questionsList.innerHTML = ""; // Clear previous questions
  });

  fetch("/api/subjects")
    .then((response) => response.json())
    .then((subjects) => {
      subjects.forEach((subject) => {
        const listItem = document.createElement("div");
        listItem.classList.add("category_cards");

        const title = document.createElement("h2");
        title.textContent = subject.title;
        listItem.appendChild(title);

        const description = document.createElement("p");
        description.textContent = subject.description;
        listItem.appendChild(description);

        listItem.addEventListener("click", function () {
          localStorage.setItem("subjectId", subject.id);
          localStorage.setItem("subjectTitle", subject.title);
          listElement.style.display = "none"; // Hide the list
          returnButton.style.display = ""; // Show return button
          fetchQuestions(subject.id);
        });

        listElement.appendChild(listItem);
      });
    });
});

function fetchQuestions(subjectId) {
  fetch(`/api/questions?subjectId=${subjectId}`)
    .then((response) => response.json())
    .then((questions) => {
      questionsList.innerHTML = ""; // Clear previous questions
      questions.forEach((question) => {
        const questionItem = document.createElement("p");
        questionItem.textContent = question;
        questionsList.appendChild(questionItem);
      });
      questionsList.style.display = ""; // Show the questions list
    });
}
