function editQuestion(id) {
  const questionElement = document.querySelector(`[data-question-id="${id}"]`);
  const inputs = questionElement.querySelectorAll(".input-field");
  const textareas = questionElement.querySelectorAll(".textarea-field");

  const data = {
    type: "editQuestionPanel",
    content: {
      id: id,
      title: inputs[0].value,
      description: textareas[0].value,
      content: textareas[1].value,
      creationDate: inputs[1].value,
      updateDate: inputs[2].value,
      upvotes: inputs[3].value,
      downvotes: inputs[4].value,
    },
    session_id: getCookie("session"),
  };
  socket.send(JSON.stringify(data));
}

function editResponse(responseId, question_id) {
  // Select the response element by its data attribute
  const responseElement = document.querySelector(
    `.response[data-response-id="${responseId}"]`
  );

  // Retrieve content from textarea and input fields
  const inputs = responseElement.querySelectorAll(".input-field");
  const textareas = responseElement.querySelectorAll(".textarea-field");
  //Create the data object
  const data = {
    type: "editResponsePanel",
    content: {
      id: responseId,
      question_id: question_id,
      content: textareas[0].value,
      description: inputs[0].value,
      creationDate: inputs[1].value,
      updateDate: inputs[2].value,
      upvotes: inputs[3].value,
      downvotes: inputs[4].value,
    },
    session_id: getCookie("session"),
  };

  // Send the JSON stringified data through the WebSocket
  socket.send(JSON.stringify(data));
}

function editSubject(id) {
  const subjectElement = document.querySelector(`[data-subject-id="${id}"]`);
  const inputs = subjectElement.querySelectorAll(".input-field");
  const textareas = subjectElement.querySelectorAll(".textarea-field");

  const data = {
    type: "editSubjectPanel",
    content: {
      id: id,
      title: inputs[0].value,
      description: textareas[0].value,
      creationDate: inputs[1].value,
      updateDate: inputs[2].value,
    },
    session_id: getCookie("session"),
  };
  socket.send(JSON.stringify(data));
}

function addSubject() {
  const newSubjectElement = document.querySelector("#add_subject");
  const inputs = newSubjectElement.querySelector(".input-field");
  const textareas = newSubjectElement.querySelector(".textarea-field");
  console.log(newSubjectElement);
  console.log(inputs);
  console.log(textareas);
  if (textareas.value == "" || inputs.value == "") {
    return;
  }

  const data = {
    type: "addSubjectPanel",
    content: {
      title: inputs.value,
      description: textareas.value,
    },
    session_id: getCookie("session"),
  };
  socket.send(JSON.stringify(data));
  //Reset fields
  inputs.value = "";
  textareas.value = "";
}

function deleteSubject(id) {
  const data = {
    type: "deleteSubjectPanel",
    content: {
      id: id,
    },
    session_id: getCookie("session"),
  };
  socket.send(JSON.stringify(data));
}

function deleteQuestion(id) {
  const data = {
    type: "deleteQuestionPanel",
    content: {
      id: id,
    },
    session_id: getCookie("session"),
  };
  socket.send(JSON.stringify(data));
}
function deleteResponse(id, question_id) {
  const data = {
    type: "deleteResponsePanel",
    content: {
      id: id,
      question_id: question_id,
    },
    session_id: getCookie("session"),
  };
  socket.send(JSON.stringify(data));
}

function editUser(id) {
  /** <div class="user_details" data-user-id="{{.ID}}">
                <input class="input-field" type="text" value="{{.FirstName}} " /><br>
                <input class="input-field" type="text" value="{{.LastName}}" /><br>
                <input class="input-field" type="text" value="@{{.Username}}" /><br>
                <input class="input-field" type="email" value="{{.Email}}" /><br>
                <textarea class="textarea-field">{{if .Bio.Valid}}{{.Bio.String}}{{else}}Not provided{{end}}</textarea><br>
                <input class="input-field" type="url" value="{{if .Website.Valid}}{{.Website.String}}{{end}}" /><br>
                <input class="input-field" type="url" value="{{if .GitHub.Valid}}{{.GitHub.String}}{{end}}" /><br>
                <input class="input-field" type="number" value="{{if .XP.Valid}}{{.XP.Int64}}{{else}}0{{end}}" /><br>
                <input class="input-field" type="text" value="{{if .Rank.Valid}}{{.Rank.String}}{{else}}Unranked{{end}}" /><br>
                <input class="input-field" type="date" value="{{if .SchoolYear.Valid}}{{.SchoolYear.Time}}{{end}}" /><br>
                <button >Edit</button>
                <button >Delete</button>
            </div>*/
  const userElement = document.querySelector(`[data-user-id="${id}"]`);
  const inputs = userElement.querySelectorAll(".input-field");
  const textareas = userElement.querySelectorAll(".textarea-field");
  //Create data object
  const data = {
    type: "editUserPanel",
    content: {
      id: id,
      firstname: inputs[0].value,
      lastname: inputs[1].value,
      username: inputs[2].value.replace("@", ""),
      email: inputs[3].value,
      bio: textareas[0].value,
      website: inputs[4].value,
      github: inputs[5].value,
      xp: inputs[6].value,
      rank: inputs[7].value,
      schoolyear: inputs[8].value,
    },
    session_id: getCookie("session"),
  };
  //Send data through WebSocket
  socket.send(JSON.stringify(data));
}
function deleteAvatar(user_id) {
  const data = {
    type: "deleteAvatarPanel",
    content: {
      user_id: user_id,
    },
    session_id: getCookie("session"),
  };
  socket.send(JSON.stringify(data));
}
function deleteUser(id) {
  const data = {
    type: "deleteUserPanel",
    content: {
      id: id,
    },
    session_id: getCookie("session"),
  };
  socket.send(JSON.stringify(data));
}

function addOneUpVoteResponse(Responseid) {
  // Select the response element by its data attribute
  const responseElement = document.querySelector(
    `.response[data-response-id="${Responseid}"]`
  );

  // Retrieve content from input field
  const inputs = responseElement.querySelectorAll(".input-field")[3];

  //update the input content
  inputs.value = parseInt(inputs.value) + 1;
}

function removeOneUpVoteResponse(Responseid) {
  // Select the response element by its data attribute
  const responseElement = document.querySelector(
    `.response[data-response-id="${Responseid}"]`
  );

  // Retrieve content from input field
  const inputs = responseElement.querySelectorAll(".input-field")[3];

  //update the input content
  inputs.value = parseInt(inputs.value) - 1;

  if (parseInt(inputs.value) < 0) {
    inputs.value = 0;
  }
}

function addOneUpVoteQuestion(questionId) {
  // Select the response element by its data attribute
  const questionElement = document.querySelector(
    `[data-question-id="${questionId}"]`
  );

  // Retrieve content from input field
  const inputs = questionElement.querySelectorAll(".input-field")[3];

  //update the input content
  inputs.value = parseInt(inputs.value) + 1;
}

function removeOneUpVoteQuestion(questionId) {
  // Select the response element by its data attribute
  const questionElement = document.querySelector(
    `[data-question-id="${questionId}"]`
  );

  // Retrieve content from input field
  const inputs = questionElement.querySelectorAll(".input-field")[3];

  //update the input content
  inputs.value = parseInt(inputs.value) - 1;

  if (parseInt(inputs.value) < 0) {
    inputs.value = 0;
  }
}

function addOneDownVoteResponse(Responseid) {
  // Select the response element by its data attribute
  const responseElement = document.querySelector(
    `.response[data-response-id="${Responseid}"]`
  );

  // Retrieve content from input field
  const inputs = responseElement.querySelectorAll(".input-field")[4];

  //update the input content
  inputs.value = parseInt(inputs.value) + 1;
}

function removeOneDownVoteResponse(Responseid) {
  // Select the response element by its data attribute
  const responseElement = document.querySelector(
    `.response[data-response-id="${Responseid}"]`
  );

  // Retrieve content from input field
  const inputs = responseElement.querySelectorAll(".input-field")[4];

  //update the input content
  inputs.value = parseInt(inputs.value) - 1;

  if (parseInt(inputs.value) < 0) {
    inputs.value = 0;
  }
}

function addOneDownVoteQuestion(questionId) {
  // Select the response element by its data attribute
  const questionElement = document.querySelector(
    `[data-question-id="${questionId}"]`
  );

  // Retrieve content from input field
  const inputs = questionElement.querySelectorAll(".input-field")[4];

  //update the input content
  inputs.value = parseInt(inputs.value) + 1;
}

function removeOneDownVoteQuestion(questionId) {
  // Select the response element by its data attribute
  const questionElement = document.querySelector(
    `[data-question-id="${questionId}"]`
  );

  // Retrieve content from input field
  const inputs = questionElement.querySelectorAll(".input-field")[4];

  //update the input content
  inputs.value = parseInt(inputs.value) - 1;

  if (parseInt(inputs.value) < 0) {
    inputs.value = 0;
  }
}

document.getElementById("show_all_users").onclick = () => {
  const all_users = document.getElementById("all_users");
  toggleVisibility(all_users);
};

document.getElementById("show_all_subjects").onclick = () => {
  const all_subjects = document.getElementById("all_subjects");
  toggleVisibility(all_subjects);
};

document.getElementById("show_all_questions").onclick = () => {
  const all_questions = document.getElementById("all_questions");
  toggleVisibility(all_questions);
};
function toggleVisibility(element) {
  // Check if the element is visible
  if (element.classList.contains("visible")) {
    element.classList.remove("visible");
    setTimeout(() => {
      element.style.height = "0"; // Collapse the element after the fade out
      element.style.display = "hidden";
    }, 500); // This should match the transition time
  } else {
    element.style.height = "auto"; // Expand the element before fading in

    element.classList.add("visible");
  }
}
function showResponses(id) {
  const responses = document
    .querySelector(`[data-question-id="${id}"]`)
    .querySelector(".all_responses");
  toggleVisibility(responses);
}

function ResendEmail(email) {
  const data = {
    type: "resendEmail",
    content: {
      email: email,
    },
    session_id: getCookie("session"),
  };
  socket.send(JSON.stringify(data));
}
function searchUsers() {
  const search = document.getElementById("search_bar_users").value;
  document.querySelectorAll(".user").forEach((u) => {
    for (let p = 0; p < u.querySelectorAll("input").length; p++) {
      let i = u.querySelectorAll("input")[p];
      console.log(i);
      if (i.value.toLowerCase().includes(search.toLowerCase())) {
        u.style.display = "block";
        console.log(i);
        break;
      } else {
        u.style.display = "none";
      }
    }
  });
}

document.getElementById("search_bar_users").addEventListener("keyup", (k) => {
  k.preventDefault();
  if (document.getElementById("search_bar_users").value == "") {
    document.querySelectorAll(".user").forEach((u) => {
      u.style.display = "block";
    });
  } else {
    searchUsers();
  }
});
