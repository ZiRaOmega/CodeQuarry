let response_submit = document.getElementById("response_submit");
response_submit.addEventListener("click", function () {
  const question_id = document
    .getElementById("question_id")
    .getAttribute("question-id");
  const response_description = document.getElementById(
    "response_description"
  ).value;
  const response_content = document.getElementById("response_content").value;
  const response = {
    question_id: getUrlArgument("question_id"),
    description: response_description,
    content: response_content,
  };
  fetch("/api/responses", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      response: response,
      session_id: getCookie("session"),
    }),
  })
    .then((response) => response.json())
    .then((data) => {
      console.log(data);
      if (data.status === "success") {
        window.location.reload();
      }
    })
    .catch((error) => {
      console.error("Error:", error);
    });
});

function getUrlArgument(name) {
  const url = new URL(window.location.href);
  return url.searchParams.get(name);
}