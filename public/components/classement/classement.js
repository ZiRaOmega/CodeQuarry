fetch("/api/classement", {
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
    if (data.status === "success") {
      window.location.reload();
    }
  });
