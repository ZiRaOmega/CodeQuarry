let fetch_results = [];

document.addEventListener("DOMContentLoaded", function () {
  // Fetch all questions
  (async () => {
    try {
      const response = await fetch(`/api/questions?subjectId=all`);
      const questions = await response.json();
      fetch_results = questions;

      /* //create the options
      const options = fetch_results.map((question) => {
        return {
          id: question.id,
          title: question.title,
          description: question.description,
          content: question.content,
          subject_title: question.subject_title,
        };
      }
      );
      datalist = document.getElementById("search_results");
      options.forEach((option) => {
        ['title','description','content','subject_title'].forEach((key)=>{
          let optionElement = document.createElement("option");
          optionElement.value = `${key}: ${option[key]}`;
          optionElement.innerText = option[key];
          optionElement.onclick = function () {console.log(option.id)};
          optionElement.setAttribute("data-id", option.id);
          optionElement.setAttribute("data-key", key);
          datalist.appendChild(optionElement);
        })

      }); */
    } catch (error) {
      console.error("There was a problem with your fetch operation:", error);
    }
  })();
});

document
  .getElementById("search_bar_input")
  .addEventListener("input", function () {
    var input = this.value.toLowerCase();
    var results_html = document.getElementById("search_results");
    results_html.innerHTML = "";

    if (input) {
      results_html.style.display = "block"; // Show dropdown

      let rslt_array = [];
      let rslt_array_word = [];

      let potential_results = [...fetch_results];

      // Process full input matches first
      potential_results.forEach(function (item) {
        let match = false;
        let rslt_object = {
          Prefix: "",
          Suffix: "",
          question: item,
        };

        if (item.title.toLowerCase().includes(input)) {
          rslt_object.Prefix = "Title: ";
          rslt_object.Suffix = item.title;
          match = true;
        } else if (item.description.toLowerCase().includes(input)) {
          rslt_object.Prefix = "Description: ";
          rslt_object.Suffix = item.description;
          match = true;
        } else if (item.content.toLowerCase().includes(input)) {
          rslt_object.Prefix = "Content: ";
          rslt_object.Suffix = item.content;
          match = true;
        } else if (item.subject_title.toLowerCase().includes(input)) {
          rslt_object.Prefix = "Subject: ";
          rslt_object.Suffix = item.subject_title;
          match = true;
        }

        if (match) {
          rslt_array.push(rslt_object);
        }
      });

      rslt_array = removeDuplicates(rslt_array);
      let sorted_array = sort_results(rslt_array, input);
      // Process individual word matches next
      let cut_input = input.split(" ");
      cut_input.forEach(function (word) {
        if (!word) return;
        else if (word.length == 0) return;
        fetch_results.forEach(function (item) {
          let match = false;
          let rslt_object = {
            Prefix: "",
            Suffix: "",
            question: item,
          };

          if (item.title.toLowerCase().includes(word)) {
            rslt_object.Prefix = "Title: ";
            rslt_object.Suffix = item.title;
            match = true;
          } else if (item.description.toLowerCase().includes(word)) {
            rslt_object.Prefix = "Description: ";
            rslt_object.Suffix = item.description;
            match = true;
          } else if (item.content.toLowerCase().includes(word)) {
            rslt_object.Prefix = "Content: ";
            rslt_object.Suffix = item.content;
            match = true;
          } else if (item.subject_title.toLowerCase().includes(word)) {
            rslt_object.Prefix = "Subject: ";
            rslt_object.Suffix = item.subject_title;
            match = true;
          }

          if (match) {
            rslt_array_word.push(rslt_object);
          }
        });
      });

      rslt_array_word = removeDuplicates(rslt_array_word);
      let sorted_array_word = sort_results(rslt_array_word, input);

      // Combine arrays, ensuring no duplicates
      let combined_results = [...sorted_array];

      sorted_array_word.forEach(function (item) {
        if (
          !combined_results.some(
            (result) => result.question.id === item.question.id
          )
        ) {
          combined_results.push(item);
        }
      });

      create_results(combined_results, input);
    } else {
      results_html.style.display = "none"; // Hide dropdown if input is empty
    }
  });

function removeDuplicates(results) {
  const uniqueResults = [];
  const ids = new Set();

  results.forEach((result) => {
    if (!ids.has(result.question.id)) {
      ids.add(result.question.id);
      uniqueResults.push(result);
    }
  });

  return uniqueResults;
}

function slice_rslt(text, input) {
  var index = text.toLowerCase().indexOf(input);
  var rslt_len = 80;
  var range = rslt_len - input.length;
  var start = index - range / 2;
  var end = index + input.length + range / 2;

  if (start < 0) {
    start = 0;
    end = rslt_len;
  }
  if (end > text.length) {
    end = text.length;
    start = end - rslt_len;
  }
  var sliced = text.slice(start, end);
  var highlighted = sliced.replace(
    new RegExp(input, "gi"),
    (match) => `<span class="highlight">${match}</span>`
  );
  return highlighted;
}

function sort_results(array, input) {
  array.sort((a, b) => {
    let a_dist = levenshtein(input, a.Suffix.toLowerCase());
    let b_dist = levenshtein(input, b.Suffix.toLowerCase());
    return a_dist - b_dist;
  });
  return array;
}

function levenshtein(a, b) {
  var dp = [];
  for (var i = 0; i <= a.length; i++) {
    dp[i] = [];
    dp[i][0] = i;
  }
  for (var j = 0; j <= b.length; j++) {
    dp[0][j] = j;
  }

  for (var i = 1; i <= a.length; i++) {
    for (var j = 1; j <= b.length; j++) {
      if (a[i - 1] === b[j - 1]) {
        dp[i][j] = dp[i - 1][j - 1];
      } else {
        dp[i][j] = 1 + Math.min(dp[i - 1][j - 1], dp[i - 1][j], dp[i][j - 1]);
      }
    }
  }

  return dp[a.length][b.length];
}

function create_results(array, input) {
  var results_html = document.getElementById("search_results");
  array.forEach(function (item) {
    var rslt_element = document.createElement("div");
    rslt_element.className = "result_element";

    var preview = document.createElement("div");
    preview.className = "result_element_preview";

    var det_title = slice_rslt(item.question.title, input);
    var det_desc = slice_rslt(item.question.description, input);

    var details = document.createElement("div");
    details.className = "result_element_details";
    details.innerHTML = `
      <div class="subject_tag">${item.question.subject_title}</div>
      <h3 class="question_title">${det_title}</h3>
      <p class="question_description">${det_desc}</p>
      <pre><code>${""}</code></pre>
    `;
    //console.log(item.question.content)
    details.querySelector("code").textContent = item.question.content;
    //console.log(details.querySelector("code"))
    var prev_suffix = slice_rslt(item.Suffix, input);
    preview.innerHTML = `<span class="result_prefix">${item.Prefix}</span><span class="result_text">${prev_suffix}</span>`;

    rslt_element.appendChild(preview);
    rslt_element.appendChild(details);

    rslt_element.onclick = function () {
      window.location.href = `/question_viewer?question_id=${item.question.id}`;
    };
    results_html.appendChild(rslt_element);
  });
  checkHighlight();
}

function checkHighlight() {
  // Ensure both hljs and Prism are defined
  if (typeof hljs === "undefined") {
    console.error("Highlight.js not found!");
    return;
  }

  document.querySelectorAll("pre code").forEach((block) => {
    // Apply Highlight.js
    hljs.highlightElement(block);
  });
}
