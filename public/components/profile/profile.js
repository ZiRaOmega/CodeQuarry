const profileInfos = document.getElementById("profileInformations");
const profileForm = document.getElementById("editProfileForm");
const button = document.getElementById("editButton");
const lines = document.getElementsByClassName("informations")
const links = document.getElementsByClassName("links")

const birthDate = document.getElementById("birthDate");
const schoolDate = document.getElementById("schoolYear")

if (birthDate.textContent == " 01/01/0001") {
	birthDate.style.display = "none";
}

if (schoolDate.textContent == " 01/01/0001") {
	schoolDate.style.display = "none";
}

for (let j of links) {
	console.log(j.href)
	if (j.href == "https://localhost/profile") {
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
});

// Define a function to set the birth date
function setDateInInput(theDate, theInput) {
    // Extract the relevant theDate components using a regular expression
    const [, year, month, day] = theDate.match(/\{(\d{4})-(\d{2})-(\d{2})/);

    // Format the theDate components into YYYY-MM-DD format
    const formattedDate = `${year}-${day}-${month}`;

    // Set the value of the input element
    let DOBInput = document.getElementById(theInput);
    DOBInput.value = formattedDate;
}
