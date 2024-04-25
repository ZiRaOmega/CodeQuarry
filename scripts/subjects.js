document.addEventListener("DOMContentLoaded", function () {
  const listElement = document.getElementById("subjectsList");
  const questionsList = document.getElementById("questionsList");
  const returnButton = document.getElementById("returnButton");
  returnButton.style.display = "none";

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

  // Event listener for the "All" subjects item
  allSubjectsItem.addEventListener("click", function () {
    localStorage.setItem("subjectId", "all");
    localStorage.setItem("subjectTitle", "All Subjects");
    listElement.style.display = "none"; // Hide the list
    returnButton.style.display = ""; // Show return button
    fetchQuestions("all"); // Fetch all questions
  });

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
          returnButton.style.display = "";
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
      const questionsList = document.getElementById("questionsList");
      questionsList.innerHTML = ""; // Clear previous questions
      questions.forEach((question) => {
        const questionContainer = document.createElement("div");
        questionContainer.classList.add("question");

        // Add subject title tag
        const subjectTag = document.createElement("div");
        subjectTag.classList.add("subject_tag");
        subjectTag.textContent = question.subject_title;
        questionContainer.appendChild(subjectTag);

        const questionTitle = document.createElement("h3");
        questionTitle.classList.add("question_title");
        questionTitle.textContent = question.title;
        questionContainer.appendChild(questionTitle);

        const questionContent = document.createElement("p");
        questionContent.classList.add("question_content");
        questionContent.textContent = question.content;
        questionContainer.appendChild(questionContent);

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
        questionCreator.textContent = "Crée par ";
        const creatorName = document.createElement("span");
        creatorName.textContent = question.creator;
        creatorName.classList.add("creator_name");
        questionCreator.appendChild(creatorName);
        ContainerCreatorAndDate.appendChild(questionCreator);

        questionContainer.appendChild(ContainerCreatorAndDate);
        questionsList.appendChild(questionContainer);
      });
      questionsList.style.display = ""; // Show the questions list
    });
}
