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
  };
  socket.send(JSON.stringify(data));
}

function addSubject() {
  const newSubjectElement = document.querySelector("#add_subject");
  const inputs = newSubjectElement.querySelector(".input-field");
  const textareas = newSubjectElement.querySelector(".textarea-field");
  if (textareas.value == "" || inputs.value == "") {
    return;
  }

  const data = {
    type: "addSubjectPanel",
    content: {
      title: inputs.value,
      description: textareas.value,
    },
  };
  socket.send(JSON.stringify(data));
  //Reset fields
  inputs.value = "";
  textareas.value = "";
}

function deleteSubject(id,element=new HTMLElement) {
  const data = {
    type: "deleteSubjectPanel",
    content: {
      id: id,
    },
  };
  socket.send(JSON.stringify(data));
  element.parentElement.remove()
}

function deleteQuestion(id,element) {
  const data = {
    type: "deleteQuestionPanel",
    content: {
      id: id,
    },
  };
  socket.send(JSON.stringify(data));
  element.parentElement.remove()
}
function deleteResponse(id, question_id, element) {
  const data = {
    type: "deleteResponsePanel",
    content: {
      id: id,
      question_id: question_id,
    },
  };
  socket.send(JSON.stringify(data));
  element.parentElement.remove()
}

function editUser(id) {
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
  };
  //Send data through WebSocket
  socket.send(JSON.stringify(data));
}
function deleteAvatar(user_id,element) {
  const data = {
    type: "deleteAvatarPanel",
    content: {
      user_id: user_id,
    },
  };
  socket.send(JSON.stringify(data));
  element.parentElement.remove()
}
function deleteUser(id,element=new HTMLElement) {
  const data = {
    type: "deleteUserPanel",
    content: {
      id: id,
    },
  };
  socket.send(JSON.stringify(data));
  element.parentElement.remove()
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
  setTimeout(() => {
    document.querySelectorAll(".user").forEach((u) => {
      toggleFlexNone(u);
    });
  }, 550);
};

document.getElementById("show_all_subjects").onclick = () => {
  const all_subjects = document.getElementById("all_subjects");
  toggleVisibility(all_subjects);
  setTimeout(() => {
    document.querySelectorAll(".subject").forEach((u) => {
      toggleFlexNone(u);
    });
  }, 550);
};

document.getElementById("show_all_questions").onclick = () => {
  const all_questions = document.getElementById("all_questions");
  toggleVisibility(all_questions);
  setTimeout(() => {
    document.querySelectorAll(".question").forEach((u) => {
      toggleFlexNone(u);
    });
  }, 550);
};
function toggleFlexNone(u) {
  //switch
  if (u.style.display == "" || u.style.display == "none") {
    u.style.display = "flex";
  } else {
    u.style.display = "none";
  }
}
function toggleVisibility(element) {
  // Check if the element is visible
  if (element.classList.contains("visible")) {
    element.classList.remove("visible");
    setTimeout(() => {
      element.style.height = "0"; // Collapse the element after the fade out
      element.style.display = "none";
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
  setTimeout(() => {
    document.querySelectorAll(".response").forEach((u) => {
      toggleFlexNone(u);
    });
  }, 550);
}

function ResendEmail(email) {
  const data = {
    type: "resendEmail",
    content: {
      email: email,
    },
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
        u.style.display = "flex";
        console.log(i);
        break;
      } else {
        u.style.display = "none";
      }
    }
  });
}

function searchSubjects() {
  const search = document.getElementById("search_bar_subjects").value;
  document.querySelectorAll(".subject").forEach((s) => {
    for (let p = 0; p < s.querySelectorAll("input").length; p++) {
      let i = s.querySelectorAll("input")[p];
      if (i.value.toLowerCase().includes(search.toLowerCase())) {
        s.style.display = "flex";
        break;
      } else {
        s.style.display = "none";
      }
    }
  });
}

function searchQuestions() {
  const search = document.getElementById("search_bar_questions").value;
  document.querySelectorAll(".question").forEach((q) => {
    for (let p = 0; p < q.querySelectorAll("input").length; p++) {
      let i = q.querySelectorAll("input")[p];
      if (i.value.toLowerCase().includes(search.toLowerCase())) {
        q.style.display = "flex";
        break;
      } else {
        q.style.display = "none";
      }
    }
  });
}

document.getElementById("search_bar_questions").addEventListener("keyup", (k) => {
  k.preventDefault();
  if (document.getElementById("search_bar_questions").value == "") {
    document.querySelectorAll(".question").forEach((u) => {
      u.style.display = "flex";
    });
  } else {
    searchQuestions();
  }
});

document.getElementById("search_bar_subjects").addEventListener("keyup", (k) => {
  k.preventDefault();
  if (document.getElementById("search_bar_subjects").value == "") {
    document.querySelectorAll(".subject").forEach((u) => {
      u.style.display = "flex";
    });
  } else {
    searchSubjects();
  }
});