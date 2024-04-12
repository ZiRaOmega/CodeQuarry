registerName = document.getElementById("registerName");
registerUsername = document.getElementById("registerUsername");
registerEmail = document.getElementById("registerEmail");
registerPassword = document.getElementById("registerPassword");
registerPasswordConfirmation = document.getElementById(
  "registerPasswordConfirmation"
);
registerForm = document.getElementById("registerForm");
registerSubmit = document.getElementById("registerSubmit");
contentAlert = document.getElementById("contentAlert");

registerSubmit.addEventListener("click", function (event) {
  const fields = [
    { value: registerName.value, name: "Name" },
    { value: registerUsername.value, name: "Username" },
    { value: registerEmail.value, name: "Email" },
    { value: registerPassword.value, name: "Password" },
    {
      value: registerPasswordConfirmation.value,
      name: "Password Confirmation",
    },
  ];

  let errorMessage = "";
  fields.forEach((field) => {
    if (field.value == "") {
      errorMessage += `${field.name} is required.<br>`;
    }
  });

  // Check if both password fields are not empty but do not match
  if (
    registerPassword.value !== registerPasswordConfirmation.value &&
    registerPassword.value !== "" &&
    registerPasswordConfirmation.value !== ""
  ) {
    errorMessage += "Passwords do not match.<br>";
  }

  // Check for password length and numeric requirements
  if (
    registerPassword.value !== "" &&
    !/^(?=.*\d)[a-zA-Z0-9]{8,}$/.test(registerPassword.value)
  ) {
    errorMessage +=
      "Password must be at least 8 characters long and contain at least one number.<br>";
  }

  // Check if the email is not empty and valid
  if (
    registerEmail.value !== "" &&
    !/^[^@]+@[^@]+\.[^@]+$/.test(registerEmail.value)
  ) {
    errorMessage += "Email must be a valid address.<br>";
  }

  event.preventDefault();

  if (!errorMessage) {
    fetch("/register", {
      method: "POST",
      headers: {
        "Content-Type": "application/x-www-form-urlencoded",
      },
      body: new URLSearchParams(new FormData(registerForm)),
    })
      .then((response) => {
        if (response.ok) {
          return response.json();
        }
        throw new Error("Registration failed");
      })
      .then((data) => {
        let registerBlock = document.getElementById("registerBlock");
        registerBlock.style.display = "none";
        Swal.fire({
          title: "Thank You!",
          text: "You have successfully registered.",
          icon: "success",
          confirmButtonText: "OK",
        }).then((result) => {
          if (result.value) {
            window.location.reload(); // Reload the page
          }
        });
      })
      .catch((error) => {
        console.error("Error:", error);
        contentAlert.innerHTML = "An error occurred during registration.";
      });
  } else {
    contentAlert.innerHTML = errorMessage;
  }
});
