loginRedirect = document.getElementById("loginRedirect");
registerRedirect = document.getElementById("registerRedirect");
login_button = document.getElementById("login-button");
register_button = document.getElementById("register-button");
registerBlock = document.getElementById("registerBlock");
loginBlock = document.getElementById("loginBlock");
mainBlock = document.getElementById("auth_block");

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
