let potential_results = [];

document.addEventListener("DOMContentLoaded", function () {
  (async () => {
    try {
      const response = await fetch(`/api/questions?subjectId=all`);
      const questions = await response.json();
      console.log(questions);
      potential_results = questions;
    } catch (error) {
      console.error(
        "There was a problem with your fetch operation:",
        error
      );
    }
  })();
});

document
  .getElementById("search_bar_input")
  .addEventListener("input", function () {
    var input = this.value.toLowerCase();
    var results = document.getElementById("search_results");
    results.innerHTML = "";

    if (input) {
      results.style.display = "block"; // Show dropdown
      potential_results.forEach(function (item) {

        var prefix = document.createElement("span");
        prefix.className = "result_prefix";

        if (item.title.toLowerCase().indexOf(input) > -1) {
          prefix.textContent = "Title: ";
        } else if (item.content.toLowerCase().indexOf(input) > -1) {
          prefix.textContent = "Content: ";
        } else if (item.description.toLowerCase().indexOf(input) > -1) {
          prefix.textContent = "Description: ";
        } else if (item.subject_title.toLowerCase().indexOf(input) > -1) {
          prefix.textContent = "Subject: ";
        } else {
          return;
        }
        var rslt_element = document.createElement("div");
        rslt_element.className = "result_element";

        var preview = document.createElement("div");
        preview.className = "result_element_preview";

        var details = document.createElement("div");
        details.className = "result_element_details";
        details.innerHTML = `<div class="subject_tag">${item.subject_title}</div>
            <h3 class="question_title">${item.title}</h3>
            <p class="question_description">${item.description}</p>
            <pre><code>${item.content}</code></pre>
        `;


        var result_text = document.createElement("span");
        result_text.className = "result_text";
        result_text.textContent = item.title;

        preview.appendChild(prefix);
        preview.appendChild(result_text);

        rslt_element.appendChild(preview);
        rslt_element.appendChild(details);


        rslt_element.onclick = function () {
          // Click on item to redirect to question_viewer page
          window.location.href = `/question_viewer?question_id=7`;
        };
        results.appendChild(rslt_element);
        checkHighlight();
      });
    } else {
      results.style.display = "none"; // Hide dropdown if input is empty
    }
  });

/*
    <div class="subject_tag">${question.subject_title}</div>
    <div class="clickable_container">
        <h3 class="question_title">${question.title}</h3>
        <p class="question_description">${question.description}</p>
        <pre><code>${question.content}</code></pre>
    </div>
*/