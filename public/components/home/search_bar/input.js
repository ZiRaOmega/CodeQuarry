// {subject: "Math", marks:[
let input = document.getElementById("search-bar-input");
input.addEventListener("keyup", function () {
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
});