const questionsList = document.getElementById("questionsList"); 

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

