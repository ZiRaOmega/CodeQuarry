loginRedirect = document.getElementById("loginRedirect");
registerRedirect = document.getElementById("registerRedirect");
login_button = document.getElementById("login-button");
register_button = document.getElementById("register-button");
registerBlock = document.getElementById("registerBlock");
loginBlock = document.getElementById("loginBlock");
mainBlock = document.getElementById("mainBlock");
registerForm = document.getElementById("registerForm");
registerSubmit = document.getElementById("registerSubmit");

login_button.addEventListener("click", function () {
  loginBlock.style.display = "flex";
  registerBlock.style.display = "none";
  mainBlock.style.display = "none";
});

register_button.addEventListener("click", function () {
  registerBlock.style.display = "flex";
  loginBlock.style.display = "none";
  mainBlock.style.display = "none";
});

loginRedirect.addEventListener("click", function () {
  loginBlock.style.display = "flex";
  registerBlock.style.display = "none";
  mainBlock.style.display = "none";
});

registerRedirect.addEventListener("click", function () {
  registerBlock.style.display = "flex";
  loginBlock.style.display = "none";
  mainBlock.style.display = "none";
});

registerName = document.getElementById("registerName");
registerUsername = document.getElementById("registerUsername");
registerEmail = document.getElementById("registerEmail");
registerPassword = document.getElementById("registerPassword");
registerPasswordConfirmation = document.getElementById(
  "registerPasswordConfirmation"
);
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

  // Check for password length, numeric, and special character requirements
  if (
    registerPassword.value !== "" &&
    !/^(?=.*\d.*\d)(?=.*[!@#$%^&*])[0-9a-zA-Z!@#$%^&*]{8,}$/.test(
      registerPassword.value
    )
  ) {
    errorMessage +=
      "Password must be at least 8 characters long, contain 2 numbers, and 1 special character.<br>";
  }

  // Check if the email is not empty and valid
  if (
    registerEmail.value !== "" &&
    !/^[^@]+@[^@]+\.(com|fr)$/.test(registerEmail.value)
  ) {
    errorMessage += "Email must be a valid .com or .fr address.<br>";
  }

  if (errorMessage) {
    contentAlert.innerHTML = errorMessage;
    event.preventDefault();
  }
});
