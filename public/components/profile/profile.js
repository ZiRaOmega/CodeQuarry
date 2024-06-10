const profileInfos = document.getElementById("profileInformations");
const profileForm = document.getElementById("editProfileForm");
const deleteForm = document.getElementById("deleteProfileForm")
const button = document.getElementById("editButton");
const lines = document.getElementsByClassName("informations")
const links = document.getElementsByClassName("links")

const allposts = document.getElementById("all_posts")

const birthDate = document.getElementById("birthDate");

if (birthDate.textContent == " 01/01/0001") {
  birthDate.style.display = "none";
}

button.addEventListener("click", function () {
  profileInfos.style.display = "none";
  profileForm.forEach(e=>{

    e.style.display = "flex";
  })
});

let deleteButtons = document.querySelectorAll(".delete_button");
deleteButtons.forEach((button) => {
  button.addEventListener("click", function () {
    let id = button.getAttribute("question-id");
    socket.send(
      JSON.stringify({
        type: "deletePost",
        content: id,
		
      })
    );
  });
});

function getFile() {
  document.getElementById("photo_changer").click();
}

function sub(obj) {
  var file = obj.value;
  var fileName = file.split("\\");
  document.getElementById("yourBtn").textContent = fileName[fileName.length - 1] + " ✔";
  event.preventDefault();
}

let deleteButtonsfavori = document.querySelectorAll(".delete_button_favori");
deleteButtonsfavori.forEach((button) => {
  button.addEventListener("click", function () {
    let id = button.getAttribute("question-id");
    socket.send(
      JSON.stringify({
        type: "deleteFavori",
        content: id,
    
      })
    );
    //Delete parent
    button.parentNode.remove();
  });
});
const schoolDate = document.getElementById("schoolYear")

if (birthDate.textContent == " 01/01/0001") {
	birthDate.style.display = "none";
}

if (schoolDate.textContent == " 01/01/0001") {
	schoolDate.style.display = "none";
}

for (let j of links) {
	console.log(j.href)
	if (j.href == "https://codequarry.dev/profile") {
		j.style.display = "none";
	}
}

for (let i of lines){
	console.log(i.textContent)
	if (i.textContent == "") {
		i.style.display = "none"
	}
} 

button.addEventListener("click", function () {
	profileInfos.style.display = "none";
	profileForm.style.display = "flex";
  deleteForm.style.display = "flex";
  allposts.style.display = "none"
});

// Define a function to set the birth date
function setDateInInput(theDate, theInput) {
    // Extract the relevant theDate components using a regular expression
    const [, year, month, day] = theDate.match(/\{(\d{4})-(\d{2})-(\d{2})/);

    // Format the theDate components into YYYY-MM-DD format
    const formattedDate = `${year}-${month}-${day}`;

    // Set the value of the input element
    let DOBInput = document.getElementById(theInput);
    DOBInput.value = formattedDate;
}



document.getElementById("resend_email").onclick = ()=>{
  socket.send(
    JSON.stringify({
      type: "resendEmail",
      content: {
        email: document.getElementById("email").value,
      },
    })
  );
}