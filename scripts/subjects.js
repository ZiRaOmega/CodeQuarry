document.addEventListener("DOMContentLoaded", function () {
  fetch("/api/subjects")
    .then((response) => response.json())
    .then((subjects) => {
      const listElement = document.getElementById("subjectsList");
      subjects.forEach((subject) => {
        const listItem = document.createElement("div");
        listItem.classList.add("category_cards");
        listItem.textContent = subject;
        listElement.appendChild(listItem);
      });
    })
    .catch((error) => console.error("Error loading subjects:", error));
});
