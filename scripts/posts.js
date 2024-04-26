function createPostInput() {
  if (document.getElementById("create_post_container") != null) {
    document.getElementById("create_post_container").remove();
  } else {
    let title = document.createElement("input");
    let content = document.createElement("textarea");
    let subjectLists = document.createElement("select");
    SubjectsList.forEach((subject) => {
        let option = document.createElement("option");
        option.value = subject.id;
        option.textContent = subject.title;
        subjectLists.appendChild(option);
        });
    let submit = document.createElement("button");
    title.setAttribute("type", "text");
    title.setAttribute("placeholder", "Title");
    content.setAttribute("placeholder", "Content");
    submit.textContent = "Submit";
    let container = document.createElement("div");
    container.id = "create_post_container";
 
    submit.onclick = function () {
      let form = {
        title: title.value,
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
      container.remove();
    };
    container.appendChild(title);
    container.appendChild(content);
    container.appendChild(subjectLists);
    container.appendChild(submit);

    document.body.appendChild(container);
  }
}

document
  .getElementById("create_post_button")
  .addEventListener("click", createPostInput);
