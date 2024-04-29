const profileInfos = document.getElementById("profileInformations");
const profileForm = document.getElementById("editProfileForm");
const button = document.getElementById("editButton");

const birthDate = document.getElementById("birthDate")

if (birthDate.textContent == " 01/01/0001") {
	birthDate.style.display = "none"; 
}

button.addEventListener("click", function() {
	profileInfos.style.display = "none";
	profileForm.style.display = "flex";
})