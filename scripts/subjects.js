document.addEventListener("DOMContentLoaded", function () {
  fetch("/api/subjects")
    .then((response) => response.json())
    .then((subjects) => {
      const listElement = document.getElementById("subjectsList");
      subjects.forEach((subject) => {
        const listItem = document.createElement("div");
        listItem.classList.add("category_cards");

        const title = document.createElement("h2");
        title.textContent = subject.title;
        listItem.appendChild(title);

        const description = document.createElement("p");
        description.textContent = subject.description;
        listItem.appendChild(description);

        listElement.appendChild(listItem);
      });
    });
});
