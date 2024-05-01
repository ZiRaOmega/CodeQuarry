const profileInfos = document.getElementById("profileInformations");
const profileForm = document.getElementById("editProfileForm");
const button = document.getElementById("editButton");

const birthDate = document.getElementById("birthDate");

if (birthDate.textContent == " 01/01/0001") {
  birthDate.style.display = "none";
}

button.addEventListener("click", function () {
  profileInfos.style.display = "none";
  profileForm.style.display = "flex";
});

let deleteButtons = document.querySelectorAll(".delete_button");
deleteButtons.forEach((button) => {
  button.addEventListener("click", function () {
    let id = button.getAttribute("question-id");
    socket.send(
      JSON.stringify({
        type: "deletePost",
        content: id,
		session_id: getCookie("session") 
      })
    );
  });
});

let deleteButtonsfavori = document.querySelectorAll(".delete_button_favori");
deleteButtonsfavori.forEach((button) => {
  button.addEventListener("click", function () {
    let id = button.getAttribute("question-id");
    socket.send(
      JSON.stringify({
        type: "deleteFavori",
        content: id,
    session_id: getCookie("session") 
      })
    );
    //Delete parent
    button.parentNode.remove();
  });
});