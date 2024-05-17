let potential_results = [];

document.addEventListener("DOMContentLoaded", function () {
  // Fetch all questions

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

        var result_text = document.createElement("span");
        result_text.className = "result_text";

        if (item.title.toLowerCase().indexOf(input) > -1) {
          prefix.textContent = "Title: ";
          result_text.innerHTML = slice_rslt(item.title, input);
        } else if (item.description.toLowerCase().indexOf(input) > -1) {
          prefix.textContent = "Description: ";
          result_text.innerHTML = slice_rslt(item.description, input);
        } else if (item.content.toLowerCase().indexOf(input) > -1) {
          prefix.textContent = "Content: ";
          result_text.innerHTML = slice_rslt(item.content, input);
        } else if (item.subject_title.toLowerCase().indexOf(input) > -1) {
          prefix.textContent = "Subject: ";
          result_text.innerHTML = slice_rslt(item.subject_title, input);
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
            <pre><code>${""}</code></pre>
        `;
        details.querySelector("code").textContent=item.content
        //console.log(item)

        //var result_text = document.createElement("span");
        //result_text.className = "result_text";
        //result_text.textContent = item.title;

        preview.appendChild(prefix);
        preview.appendChild(result_text);

        rslt_element.appendChild(preview);
        rslt_element.appendChild(details);


        rslt_element.onclick = function () {
          // Click on item to redirect to question_viewer page
          window.location.href = `/question_viewer?question_id=${item.id}`;
        };
        results.appendChild(rslt_element);
        checkHighlight();
      });
    } else {
      results.style.display = "none"; // Hide dropdown if input is empty
    }
  });


function slice_rslt(text, input) {
  // range 30 characters before and after the input and highlight the input
  var index = text.toLowerCase().indexOf(input);
  var start = index - 40;
  var end = index + input.length + 40;
  if (start < 0) {
    start = 0;
  }
  if (end > text.length) {
    end = text.length;
  }
  //console.log("start :", index - start,"\nend :", end - (index + input.length) );

  var sliced = text.slice(start, end);

  if (sliced.length > 80) {
    // cut the start and the end of the output if it's too long (more than 80 characters)
    sliced = sliced.slice(index - start, index - start + 80);

  }

  // replace the input with highlighted input and make sure that the input is text only and will not be treated as html
  var highlighted = sliced.replace(
    new RegExp(input, "gi"),
    (match) => `<span class="highlight">${match}</span>`
  );
  // slice the output if it's too long (more than 80 characters)

  
  return highlighted;

}
/*
    <div class="subject_tag">${question.subject_title}</div>
    <div class="clickable_container">
        <h3 class="question_title">${question.title}</h3>
        <p class="question_description">${question.description}</p>
        <pre><code>${question.content}</code></pre>
    </div>
*/