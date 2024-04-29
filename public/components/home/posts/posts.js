function createPostInput() {
  let title = document.createElement("input");
  let description = document.createElement("textarea");
  let content = document.createElement("textarea");
  let subjectLists = document.createElement("select");
  SubjectsList.forEach((subject) => {
    let option = document.createElement("option");
    option.value = subject.id;
    option.textContent = subject.title;
    subjectLists.appendChild(option);
  });

  title.setAttribute("type", "text");
  title.setAttribute("placeholder", "Title");
  description.setAttribute("placeholder", "Description");
  content.setAttribute("placeholder", "Content");
  let create_post_description = document.createElement("p");
  create_post_description.textContent = "Qu'est-ce qui vous creuse l'esprit ?";
  let container = document.createElement("div");
  container.id = "create_post_container";
  let black_background = document.createElement("div");
  black_background.id = "black_background";
  let submit = document.createElement("button");
  submit.textContent = "Publier";
  submit.onclick = function () {
    let form = {
      title: title.value,
      description: description.value,
      content: content.value,
      subject_id: subjectLists.value,
    };
    socket.send(
      JSON.stringify({
        type: "createPost",
        content: form,
        session_id: getCookie("session"),
      })
    );
    black_background.remove();
  };
  let cross = document.createElement("div");
  cross.textContent = "Annuler";
  cross.id = "cross";
  container.appendChild(create_post_description);
  container.appendChild(cross);
  container.appendChild(title);
  container.appendChild(description);
  container.appendChild(content);
  container.appendChild(subjectLists);
  container.appendChild(submit);
  black_background.appendChild(container);
  document.body.appendChild(black_background);

  cross.onclick = function () {
    black_background.remove();
  };
  black_background.onclick = function (e) {
    if (e.target.id == "black_background") black_background.remove();
  };
}

document
  .getElementById("create_post_button")
  .addEventListener("click", createPostInput);
