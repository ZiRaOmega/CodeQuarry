$(document).ready(function () {
  let terms = false;

  // Track terms acceptance for registration form
  document.getElementById("acceptTerms").addEventListener("change", function () {
      terms = this.checked; // Update terms based on checkbox status
  });

  // Submit handler for register form
  $("#registerForm").submit(function (event) {
      event.preventDefault(); // Prevent the default form submission

      // Check if terms are accepted
      if (!terms) {
          Swal.fire({
              icon: "error",
              title: "Oops...",
              text: "Please accept the terms and conditions",
              confirmButtonText: "OK",
          });
          return; // Stop execution if terms are not accepted
      }

      // Validate registration form fields
      let errorMessage = validateRegisterFields();
      if (errorMessage) {
          contentAlert.innerHTML = errorMessage;
          return; // Stop execution if there are validation errors
      }

      // Execute reCAPTCHA for register form if all checks pass
      grecaptcha.execute('6LfndH0qAAAAAI5c26D0fwZ0dqAbrLhk9nhWS8mj', { action: 'register' })
          .then(function (token) {
              onSubmitRegister(token); // Call onSubmitRegister with the reCAPTCHA token
          });
  });

  // Submit handler for login form
  $("#loginForm").submit(function (event) {
      event.preventDefault(); // Prevent the default form submission

      // Execute reCAPTCHA for login form
      grecaptcha.execute('6LfndH0qAAAAAI5c26D0fwZ0dqAbrLhk9nhWS8mj', { action: 'login' })
          .then(function (token) {
              onSubmitLogin(token); // Call onSubmitLogin with the reCAPTCHA token
          });
  });
});

// onSubmit function for registration form
function onSubmitRegister(token) {
  let form = new FormData(registerForm);
  form.append('g-recaptcha-response', token); // Append reCAPTCHA token to form data

  fetch("/register", {
      method: "POST",
      body: new URLSearchParams(form),
  })
  .then(response => response.json())
  .then(data => {
      handleRegisterResponse(data); // Call response handler for register
  })
  .catch(error => {
      contentAlert.innerHTML = error.message;
  });
}

// onSubmit function for login form
function onSubmitLogin(token) {
  let form = new FormData(loginForm);
  form.append('g-recaptcha-response', token); // Append reCAPTCHA token to form data

  fetch("/login", {
      method: "POST",
      body: new URLSearchParams(form),
  })
  .then(response => response.json())
  .then(data => {
      handleLoginResponse(data); // Call response handler for login
  })
  .catch(error => {
      displayLoginError();
  });
}

// Validation for registration fields
function validateRegisterFields() {
  const fields = [
      { value: registerLastName.value, name: "LastName" },
      { value: registerFirstName.value, name: "FirstName" },
      { value: registerUsername.value, name: "Username" },
      { value: registerEmail.value, name: "Email" },
      { value: registerPassword.value, name: "Password" },
      { value: registerPasswordConfirmation.value, name: "Password Confirmation" },
  ];

  let errorMessage = "";
  fields.forEach((field) => {
      if (field.value == "") {
          errorMessage += `${field.name} is required.<br>`;
      }
  });

  if (
      registerPassword.value !== registerPasswordConfirmation.value &&
      registerPassword.value !== "" &&
      registerPasswordConfirmation.value !== ""
  ) {
      errorMessage += "Passwords do not match.<br>";
  }

  let regex = /^(?=.*[0-9])(?=.*[^a-zA-Z0-9])[a-zA-Z0-9!@#$%^&*()_+=\-`~\[\]{};':"\\|,.<>\/?]{8,}$/;
  if (registerPassword.value !== "" && !regex.test(registerPassword.value)) {
      errorMessage += "Password must be at least 8 characters long, contain at least one number, and one special character.<br>";
  }

  if (registerEmail.value !== "" && !/^[^@]+@[^@]+\.[^@]+$/.test(registerEmail.value)) {
      errorMessage += "Email must be a valid address.<br>";
  }

  return errorMessage;
}

// Handle registration response
function handleRegisterResponse(data) {
  if (data.status === "success") {
      Swal.fire({
          title: "Thank You!",
          text: data.message + " don't forget to verify your email",
          icon: "success",
          confirmButtonText: "OK",
      }).then((result) => {
          if (result.value) {
              window.location.href = "/home";
          }
      });
  } else {
      throw new Error(data.message || "Registration failed");
  }
}

// Handle login response
function handleLoginResponse(data) {
  if (data.status === "success") {
      window.location.href = "/home";
  } else {
      displayLoginError(data.message || "Invalid login credentials or email not verified");
  }
}

// Display login error
function displayLoginError(message) {
  Swal.fire({
      icon: "error",
      title: "Oops...",
      text: message,
      confirmButtonText: "OK",
  }).then((result) => {
      if (result.value) {
          setTimeout(() => {
              document.location.reload();
          }, 500);
      }
  });
}