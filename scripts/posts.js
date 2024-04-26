function createPostInput() {
  if (document.getElementById("create_post_container") != null) {
    document.getElementById("create_post_container").remove();
  } else {
    let title = document.createElement("input");
    let content = document.createElement("textarea");
    let submit = document.createElement("button");
    title.setAttribute("type", "text");
    title.setAttribute("placeholder", "Title");
    content.setAttribute("placeholder", "Content");
    submit.textContent = "Submit";
    let container = document.createElement("div");
    container.id = "create_post_container";
    container.style.position = "absolute"
    container.style.backgroundColor = "beige";
    container.style.top = "50%";
    submit.onclick = function () {
      let form = {
        title: title.value,
        content: content.value,
        subject_id: localStorage.getItem("subjectId")
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
    container.appendChild(submit);

    document.body.appendChild(container);
  }
}

document
  .getElementById("create_post_button")
  .addEventListener("click", createPostInput);
