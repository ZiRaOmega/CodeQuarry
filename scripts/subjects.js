let SubjectsList = [];
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
          returnButton.style.display = "";
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
      console.log(questions);
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
        const preDiv = document.createElement("pre");
        const code = document.createElement("code");
        preDiv.appendChild(code);
        code.innerHTML = question.content;
        questionContainer.appendChild(preDiv);

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
        let responseContainer= document.createElement("div");
        responseContainer.classList.add("response_container");
        if (question.responses!=undefined){

          question.responses.forEach((response) => {
            const responseContainer = document.createElement("div");
            responseContainer.classList.add("response");
            const responseContent = document.createElement("p");
            responseContent.classList.add("response_content");
            responseContent.textContent = response.Content;
            responseContainer.appendChild(responseContent);
            const responseCreator = document.createElement("p");
            responseCreator.classList.add("response_creator");
            responseCreator.textContent = `Réponse de ${response.StudentName}`;
            responseContainer.appendChild(responseCreator);
            const responseDate = document.createElement("p");
            responseDate.classList.add("response_creation_date");
            responseDate.textContent = `Publié le: ${new Date(response.CreationDate).toLocaleDateString()}`;
            responseContainer.appendChild(responseDate);
            
            responseContainer.appendChild(responseDate);
            questionContainer.appendChild(responseContainer);
          });
        }
        upvoteContainer.onclick = function () {
          socket.send(
            JSON.stringify({
              type: "upvote",
              content: question.id,
              session_id: getCookie("session"),
            })
          );
        };

        downvoteContainer.onclick = function () {
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
    });
};

