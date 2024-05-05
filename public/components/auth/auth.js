let registerLastName = document.getElementById("registerLastName");
let registerFirstName = document.getElementById("registerFirstName");
let registerUsername = document.getElementById("registerUsername");
let registerEmail = document.getElementById("registerEmail");
let registerPassword = document.getElementById("registerPassword");
let registerPasswordConfirmation = document.getElementById(
  "registerPasswordConfirmation"
);
let registerForm = document.getElementById("registerForm");
let registerSubmit = document.getElementById("registerSubmit");
let contentAlert = document.getElementById("contentAlert");
//registerSubmit.addEventListener("click", function (event) {
$(document).ready(function () {
  $("#registerForm").submit(function (event) {
    event.preventDefault();
    const fields = [
      { value: registerLastName.value, name: "LastName" },
      { value: registerFirstName.value, name: "FirstName" },
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
    if (registerPassword.value !== "") {
      var regex =
        /^(?=.*[0-9])(?=.*[^a-zA-Z0-9])[a-zA-Z0-9!@#$%^&*()_+=\-`~\[\]{};':"\\|,.<>\/?]{8,}$/;
      /* 
        ^ asserts position at the start of a line
        (?=.*[0-9]) asserts that the string contains at least one number
        (?=.*[^a-zA-Z0-9]) asserts that the string contains at least one special character
        [a-zA-Z0-9!@#$%^&*()_+=\-`~\[\]{};':"\\|,.<>\/?]{8,} matches any character in the set, including special characters, at least 8 times
        $ asserts position at the end of a line
      */
      if (!regex.test(registerPassword.value)) {
        errorMessage +=
          "Password must be at least 8 characters long, contain at least one number, and contain at least one special character.<br>";
      }
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
        .then((response) => response.json())
        .then((data) => {
          if (data.status === "success") {
            let registerBlock = document.getElementById("registerBlock");
            registerBlock.style.display = "none";
            Swal.fire({
              title: "Thank You!",
              text: data.message,
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
        })
        .catch((error) => {
          console.error("Error:", error);
          contentAlert.innerHTML = error.message;
        });
    } else {
      contentAlert.innerHTML = errorMessage;
    }
  });
});
let login = document.getElementById("login");
let usernameOrEmailLogin = document.getElementById("usernameOrEmailLogin");
let passwordLogin = document.getElementById("loginPassword");
let contentAlertLogin = document.getElementById("contentAlertLogin");
$(document).ready(function () {
  $("#loginForm").submit(function (event) {
    event.preventDefault(); // Prevent the default form submission
    var formData = {
      usernameOrEmailLogin: $("#usernameOrEmailLogin").val(),
      passwordLogin: $("#loginPassword").val(),
    };
    $.ajax({
      type: "POST",
      url: "/login",
      data: $.param(formData), // Correctly encode the data as URL-encoded string
      contentType: "application/x-www-form-urlencoded", // Ensure the content type is set correctly
      success: function (response) {
        if (response.status === "success") {
          window.location.href = "/home"; // Redirect to the forum page
        } else {
          let loginBlock = document.getElementById("loginBlock");
          loginBlock.style.display = "none";
          Swal.fire({
            icon: "error",
            title: "Oops...",
            text: response.message || "Invalid login credentials!",
            confirmButtonText: "OK",
          }).then((result) => {
            if (result.value) {
              setTimeout(() => {
                loginBlock.style.display = "flex";
              }, 500);
            }
          });
        }
      },
      error: function () {
        let loginBlock = document.getElementById("loginBlock");
        loginBlock.style.display = "none";
        Swal.fire({
          icon: "error",
          title: "Oops...",
          text: "Invalid login credentials!",
        }).then((result) => {
          if (result.value) {
            setTimeout(() => {
              loginBlock.style.display = "flex";
            }, 300);
            loginBlock.style.animation = "fadeIn 0.3s ease-in-out";
          }
        });
      },
    });
  });
});
