let potential_results = [];

document.addEventListener("DOMContentLoaded", function () {
  // Fetch all questions

  (async () => {
    try {
      const response = await fetch(`/api/questions?subjectId=all`);
      const questions = await response.json();
      //console.log(questions);
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
    var results_html = document.getElementById("search_results");
    results_html.innerHTML = "";

    // map with the results prefix and suffix
    var rslt_array = [];

    

    if (input) {
      results_html.style.display = "block"; // Show dropdown
      potential_results.forEach(function (item) {

        /* var prefix = document.createElement("span");
        prefix.className = "result_prefix"; */

        /* var result_text = document.createElement("span");
        result_text.className = "result_text"; */

        let rslt_object = {
          Prefix: "",
          Suffix: ""
        };

        var prefix = "";
        var suffix = "";

        if (item.title.toLowerCase().indexOf(input) > -1) {
          prefix = "Title: ";
          suffix = item.title;
          //result_text.innerHTML = slice_rslt(item.title, input);
        } else if (item.description.toLowerCase().indexOf(input) > -1) {
          prefix = "Description: ";
          suffix = item.description;
          //result_text.innerHTML = slice_rslt(item.description, input);
        } else if (item.content.toLowerCase().indexOf(input) > -1) {
          prefix = "Content: ";
          suffix = item.content;
          //result_text.innerHTML = slice_rslt(item.content, input);
        } else if (item.subject_title.toLowerCase().indexOf(input) > -1) {
          prefix = "Subject: ";
          suffix = item.subject_title;
          //result_text.innerHTML = slice_rslt(item.subject_title, input);
        } else {
          return;
        }
        rslt_object.Prefix = prefix;
        rslt_object.Suffix = suffix;
        rslt_array.push(rslt_object);

        /* var rslt_element = document.createElement("div");
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
        
        preview.appendChild(prefix);
        preview.appendChild(result_text);
        
        rslt_element.appendChild(preview);
        rslt_element.appendChild(details);
        
        
        rslt_element.onclick = function () {
          // Click on item to redirect to question_viewer page
          window.location.href = `/question_viewer?question_id=${item.id}`;
        }; */

        //results_html.appendChild(rslt_element);
        //TODO : resolve the issue of the highlight
        //checkHighlight();
      });
      console.log(rslt_array);

      //results = sort_results(results, input);
      // sort the map based on the suffix selon les fonctions sort_results et levenshtein
      //var sorted_rslt_map = sort_results(rslt_map, input);
      //console.log(sorted_rslt_map);
      //results_html.innerHTML = sorted_rslt_map;

      var sorted_array = sort_results(rslt_array, input);
      console.log(sorted_array);
      create_results(sorted_array);
    } else {
      results_html.style.display = "none"; // Hide dropdown if input is empty
    }
  });


function slice_rslt(text, input) {
  var index = text.toLowerCase().indexOf(input);
  // modify here if we want to change the length of the output
  var rslt_len = 80;
  var range = rslt_len - input.length;
  var start = index - (range / 2);
  var end = index + input.length + (range / 2);

  // make sure that the start and end are within the text
  // if the result is too short, we display all of it
  // if the result is too long, the slice will always be 80 characters
  if (start < 0) {
    start = 0;
    // move the end to the right if the start is less than 0 to keep the length of the output = 80
    end = rslt_len;
  }
  if (end > text.length) {
    end = text.length;
    // move the start to the left if the end is more than the length of the text to keep the length of the output = 80
    start = end - rslt_len;
  }
  var sliced = text.slice(start, end);

  // replace the input with highlighted input and make sure that the input is text only and will not be treated as html
  var highlighted = sliced.replace(
    new RegExp(input, "gi"),
    (match) => `<span class="highlight">${match}</span>`
  );

  return highlighted;

}           

function sort_results(array, input) {
  // sort the results () based on how close the match is to the input
  // for example, if the input is "abc", the result "abcde" should be higher than "abde"
  // So, we will sort the results based on the pourcentage of the match
  // the higher the pourcentage, the higher the result
  // we will use the levenstein distance to calculate the pourcentage
  // let's begin : 

  // first, we will calculate the distance between the input and the result
  /* var distances = [];
  results.forEach(function (result) {
    var distance = levenshtein(input, result);
    distances.push(distance);
  });

  // second, we will calculate the pourcentage of the match
  var pourcentage = [];
  distances.forEach(function (distance) {
    var p = 1 - distance / input.length;
    pourcentage.push(p);
  });

  // third, we will sort the results based on the pourcentage
  var sorted_results = [];
  pourcentage.forEach(function (p, i) {
    sorted_results.push([results[i], p]);
  });
  sorted_results.sort(function (a, b) {
    return b[1] - a[1];
  }); */

  // same as above but with the suffix of each result
  // so the output will be the array of the results sorted based on the pourcentage of the match

  var distances = [];
  array.forEach(function (result) {
    var distance = levenshtein(input, result.Suffix);
    distances.push(distance);
  });

  var pourcentage = [];
  distances.forEach(function (distance) {
    var p = 1 - distance / input.length;
    pourcentage.push(p);
  });

  var sorted_results = [];
  pourcentage.forEach(function (p, i) {
    sorted_results.push([array[i], p]);
  });
  sorted_results.sort(function (a, b) {
    return b[1] - a[1];
  });

  return sorted_results;
}
/* 
function levenshtein(a, b) {
  var tmp;
  if (a.length === 0) {
    return b.length;
  }
  if (b.length === 0) {
    return a.length;
  }
  if (a.length > b.length) {
    tmp = a;
    a = b;
    b = tmp;
  }

  var i, j, res, alen = a.length, blen = b.length, row = Array(alen);
  for (i = 0; i <= alen; i++) {
    row[i] = i;
  }

  for (i = 1; i <= blen; i++) {
    res = i;
    for (j = 1; j <= alen; j++) {
      tmp = row[j - 1];
      row[j - 1] = res;
      res = b[i - 1] === a[j - 1] ? tmp : Math.min(tmp + 1, Math.min(res + 1, row[j] + 1));
    }
  }
  return res;
} */

function levenshtein(a, b) {
  // calculate the distance between two strings
  // the distance is the number of changes needed to transform a to b
  // the changes can be insertion, deletion, or substitution
  // we will use the dynamic programming to calculate the distance
  // we will use a 2D array to store the distances between the substrings of a and b
  // the distance between a and b will be stored in the last cell of the array
  // let's begin :

  // first, we will create the 2D array and fill the first row and the first column
  var dp = [];
  for (var i = 0; i <= a.length; i++) {
    dp[i] = [];
    dp[i][0] = i;
  }
  for (var j = 0; j <= b.length; j++) {
    dp[0][j] = j;
  }

  // second, we will fill the array based on the distance between the substrings
  for (var i = 1; i <= a.length; i++) {
    for (var j = 1; j <= b.length; j++) {
      if (a[i - 1] === b[j - 1]) {
        dp[i][j] = dp[i - 1][j - 1];
      } else {
        dp[i][j] = 1 + Math.min(dp[i - 1][j - 1], dp[i - 1][j], dp[i][j - 1]);
      }
    }
  }

  // finally, we will return the distance between a and b
  return dp[a.length][b.length];
}

function create_results(array) {
  var results_html = document.getElementById("search_results");
  array.forEach(function (item) {
    var rslt_element = document.createElement("div");
    rslt_element.className = "result_element";
    
    var preview = document.createElement("div");
    preview.className = "result_element_preview";
    
    var details = document.createElement("div");
    details.className = "result_element_details";
    details.innerHTML = `<div class="subject_tag">${item[0].Suffix}</div>
    <h3 class="question_title">${item[0].Suffix}</h3>
    <p class="question_description">${item[0].Suffix}</p>
    <pre><code>${""}</code></pre>
    `;
    details.querySelector("code").textContent=item[0].Suffix
    
    preview.innerHTML = `<span class="result_prefix">${item[0].Prefix}</span><span class="result_text">${item[0].Suffix}</span>`;
    
    rslt_element.appendChild(preview);
    rslt_element.appendChild(details);
    
    
    rslt_element.onclick = function () {
      // Click on item to redirect to question_viewer page
      window.location.href = `/question_viewer?question_id=${item[0].id}`;
    };
    results_html.appendChild(rslt_element);
  });
}


/* 
// some of example to test the sort_result
console.log("for the input abc");
console.log(sort_results(["abcde", "abde", "abce", "abde"], "abc"));
console.log("for the input abde");
console.log(sort_results(["abcde", "abde", "abce", "abde"], "abde"));
console.log("for the input ab");
console.log(sort_results(["abcde", "abde", "abce", "abde"], "ab"));
console.log("for the input lorem ipsum");
console.log(sort_results(["lorem ipsum", "lorem", "ipsum", "lorem ipsum dolor"], "lorem ipsum"));
 */


// !!! TODO : maybe replace include with some()
// to make this example work : 
//console.log("for the input lorem ipsum");
//console.log(sort_results(["lorem ipsum", "lorem", "ipsum", "lorem ipsum dolor"], "lorem ipsum"));
//output eventualy expected: [["lorem ipsum", 1], ["lorem ipsum dolor", 0.6666666666666666], ["lorem", 0.5], ["ipsum", 0.5]]
// for now, we have:
// output : [["lorem ipsum", 0.9090909090909091], ["lorem ipsum dolor", 0.36363636363636365],
