function a0_0x1bfe(_0x4f85ac, _0x526963) {
  const _0x567d29 = a0_0x567d();
  return (
    (a0_0x1bfe = function (_0x1bfe71, _0x49ebaf) {
      _0x1bfe71 = _0x1bfe71 - 0x1b9;
      let _0xb92a55 = _0x567d29[_0x1bfe71];
      return _0xb92a55;
    }),
    a0_0x1bfe(_0x4f85ac, _0x526963)
  );
}
const a0_0x5b642d = a0_0x1bfe;
(function (_0x115653, _0xc24411) {
  const _0x11b25a = a0_0x1bfe,
    _0x1e9c6a = _0x115653();
  while (!![]) {
    try {
      const _0xaa81d5 =
        -parseInt(_0x11b25a(0x1d1)) / 0x1 +
        parseInt(_0x11b25a(0x1ed)) / 0x2 +
        -parseInt(_0x11b25a(0x1b9)) / 0x3 +
        (parseInt(_0x11b25a(0x1bc)) / 0x4) *
          (-parseInt(_0x11b25a(0x1da)) / 0x5) +
        parseInt(_0x11b25a(0x1e4)) / 0x6 +
        (-parseInt(_0x11b25a(0x1e3)) / 0x7) *
          (-parseInt(_0x11b25a(0x1be)) / 0x8) +
        (-parseInt(_0x11b25a(0x1e8)) / 0x9) *
          (-parseInt(_0x11b25a(0x1db)) / 0xa);
      if (_0xaa81d5 === _0xc24411) break;
      else _0x1e9c6a["push"](_0x1e9c6a["shift"]());
    } catch (_0x1fb048) {
      _0x1e9c6a["push"](_0x1e9c6a["shift"]());
    }
  }
})(a0_0x567d, 0x8550e);
function a0_0x567d() {
  const _0xce604e = [
    "ready",
    "#loginForm",
    "registerBlock",
    "registerFirstName",
    "Password\x20must\x20be\x20at\x20least\x208\x20characters\x20long,\x20contain\x20at\x20least\x20one\x20number,\x20and\x20contain\x20at\x20least\x20one\x20special\x20character.<br>",
    "fire",
    "val",
    "test",
    "Email\x20must\x20be\x20a\x20valid\x20address.<br>",
    "style",
    "386520MrIail",
    "#loginPassword",
    "innerHTML",
    "href",
    "registerSubmit",
    "usernameOrEmailLogin",
    "submit",
    "catch",
    "flex",
    "5CCtUNy",
    "5028290oRTlrl",
    "Invalid\x20login\x20credentials!",
    "none",
    "contentAlert",
    "registerPasswordConfirmation",
    "#registerForm",
    "/register",
    "display",
    "245swgsPo",
    "2554542tbWftQ",
    "POST",
    "#usernameOrEmailLogin",
    "Password",
    "27sPiJHI",
    "Email",
    "login",
    "location",
    "loginBlock",
    "198066kFFjSq",
    "value",
    "json",
    "error",
    "Passwords\x20do\x20not\x20match.<br>",
    "success",
    "forEach",
    "2451471cjJKqw",
    "getElementById",
    "then",
    "2066972kcMbhp",
    "preventDefault",
    "53304xVhJGM",
    "registerForm",
    "message",
    "status",
    "name",
    "ajax",
    "Thank\x20You!",
    "Oops...",
    "application/x-www-form-urlencoded",
  ];
  a0_0x567d = function () {
    return _0xce604e;
  };
  return a0_0x567d();
}
let registerLastName = document["getElementById"]("registerLastName"),
  registerFirstName = document["getElementById"](a0_0x5b642d(0x1ca)),
  registerUsername = document[a0_0x5b642d(0x1ba)]("registerUsername"),
  registerEmail = document[a0_0x5b642d(0x1ba)]("registerEmail"),
  registerPassword = document[a0_0x5b642d(0x1ba)]("registerPassword"),
  registerPasswordConfirmation = document[a0_0x5b642d(0x1ba)](
    a0_0x5b642d(0x1df)
  ),
  registerForm = document[a0_0x5b642d(0x1ba)](a0_0x5b642d(0x1bf)),
  registerSubmit = document[a0_0x5b642d(0x1ba)](a0_0x5b642d(0x1d5)),
  contentAlert = document[a0_0x5b642d(0x1ba)](a0_0x5b642d(0x1de));
$(document)[a0_0x5b642d(0x1c7)](function () {
  const _0x109d4c = a0_0x5b642d;
  $(_0x109d4c(0x1e0))[_0x109d4c(0x1d7)](function (_0x155d11) {
    const _0x4cff4e = _0x109d4c;
    _0x155d11[_0x4cff4e(0x1bd)]();
    const _0x28f81d = [
      { value: registerLastName[_0x4cff4e(0x1ee)], name: "LastName" },
      { value: registerFirstName[_0x4cff4e(0x1ee)], name: "FirstName" },
      { value: registerUsername[_0x4cff4e(0x1ee)], name: "Username" },
      { value: registerEmail["value"], name: _0x4cff4e(0x1e9) },
      { value: registerPassword[_0x4cff4e(0x1ee)], name: _0x4cff4e(0x1e7) },
      {
        value: registerPasswordConfirmation[_0x4cff4e(0x1ee)],
        name: "Password\x20Confirmation",
      },
    ];
    let _0x6301f9 = "";
    _0x28f81d[_0x4cff4e(0x1f3)]((_0x1768e4) => {
      const _0x3fc6dd = _0x4cff4e;
      _0x1768e4[_0x3fc6dd(0x1ee)] == "" &&
        (_0x6301f9 += _0x1768e4[_0x3fc6dd(0x1c2)] + "\x20is\x20required.<br>");
    });
    registerPassword[_0x4cff4e(0x1ee)] !==
      registerPasswordConfirmation[_0x4cff4e(0x1ee)] &&
      registerPassword[_0x4cff4e(0x1ee)] !== "" &&
      registerPasswordConfirmation[_0x4cff4e(0x1ee)] !== "" &&
      (_0x6301f9 += _0x4cff4e(0x1f1));
    if (registerPassword["value"] !== "") {
      var _0x4fc4a4 =
        /^(?=.*[0-9])(?=.*[^a-zA-Z0-9])[a-zA-Z0-9!@#$%^&*()_+=\-`~\[\]{};':"\\|,.<>\/?]{8,}$/;
      !_0x4fc4a4[_0x4cff4e(0x1ce)](registerPassword[_0x4cff4e(0x1ee)]) &&
        (_0x6301f9 += _0x4cff4e(0x1cb));
    }
    registerEmail[_0x4cff4e(0x1ee)] !== "" &&
      !/^[^@]+@[^@]+\.[^@]+$/["test"](registerEmail[_0x4cff4e(0x1ee)]) &&
      (_0x6301f9 += _0x4cff4e(0x1cf)),
      _0x155d11["preventDefault"](),
      !_0x6301f9
        ? fetch(_0x4cff4e(0x1e1), {
            method: _0x4cff4e(0x1e5),
            headers: { "Content-Type": _0x4cff4e(0x1c6) },
            body: new URLSearchParams(new FormData(registerForm)),
          })
            ["then"]((_0x8a96b0) => _0x8a96b0[_0x4cff4e(0x1ef)]())
            [_0x4cff4e(0x1bb)]((_0x58216e) => {
              const _0x17e6b8 = _0x4cff4e;
              if (_0x58216e[_0x17e6b8(0x1c1)] === _0x17e6b8(0x1f2)) {
                let _0x508e8d = document[_0x17e6b8(0x1ba)](_0x17e6b8(0x1c9));
                (_0x508e8d["style"][_0x17e6b8(0x1e2)] = "none"),
                  Swal[_0x17e6b8(0x1cc)]({
                    title: _0x17e6b8(0x1c4),
                    text: _0x58216e[_0x17e6b8(0x1c0)],
                    icon: _0x17e6b8(0x1f2),
                    confirmButtonText: "OK",
                  })[_0x17e6b8(0x1bb)]((_0x33b0a8) => {
                    const _0xb1191d = _0x17e6b8;
                    _0x33b0a8[_0xb1191d(0x1ee)] &&
                      (window[_0xb1191d(0x1eb)][_0xb1191d(0x1d4)] =
                        "login.html");
                  });
              } else
                throw new Error(
                  _0x58216e[_0x17e6b8(0x1c0)] || "Registration\x20failed"
                );
            })
            [_0x4cff4e(0x1d8)]((_0x2b7660) => {
              const _0x32dc83 = _0x4cff4e;
              console[_0x32dc83(0x1f0)]("Error:", _0x2b7660),
                (contentAlert[_0x32dc83(0x1d3)] = _0x2b7660[_0x32dc83(0x1c0)]);
            })
        : (contentAlert["innerHTML"] = _0x6301f9);
  });
});
let login = document["getElementById"](a0_0x5b642d(0x1ea)),
  usernameOrEmailLogin = document["getElementById"](a0_0x5b642d(0x1d6)),
  passwordLogin = document["getElementById"]("loginPassword"),
  contentAlertLogin = document["getElementById"]("contentAlertLogin");
$(document)[a0_0x5b642d(0x1c7)](function () {
  const _0x47ec38 = a0_0x5b642d;
  $(_0x47ec38(0x1c8))["submit"](function (_0x5d09a7) {
    const _0x40d270 = _0x47ec38;
    _0x5d09a7[_0x40d270(0x1bd)]();
    var _0x5d0745 = {
      usernameOrEmailLogin: $(_0x40d270(0x1e6))[_0x40d270(0x1cd)](),
      passwordLogin: $(_0x40d270(0x1d2))["val"](),
    };
    $[_0x40d270(0x1c3)]({
      type: "POST",
      url: "/login",
      data: $["param"](_0x5d0745),
      contentType: _0x40d270(0x1c6),
      success: function (_0x5f1830) {
        const _0x557abb = _0x40d270;
        if (_0x5f1830["status"] === _0x557abb(0x1f2))
          window["location"][_0x557abb(0x1d4)] = "/codeQuarry";
        else {
          let _0x25629a = document["getElementById"](_0x557abb(0x1ec));
          (_0x25629a[_0x557abb(0x1d0)][_0x557abb(0x1e2)] = "none"),
            Swal[_0x557abb(0x1cc)]({
              icon: _0x557abb(0x1f0),
              title: _0x557abb(0x1c5),
              text: _0x5f1830[_0x557abb(0x1c0)] || _0x557abb(0x1dc),
              confirmButtonText: "OK",
            })["then"]((_0x504973) => {
              const _0x456259 = _0x557abb;
              _0x504973[_0x456259(0x1ee)] &&
                setTimeout(() => {
                  const _0xd7edee = _0x456259;
                  _0x25629a[_0xd7edee(0x1d0)]["display"] = "flex";
                }, 0x1f4);
            });
        }
      },
      error: function () {
        const _0x1ce5fd = _0x40d270;
        let _0x51938c = document[_0x1ce5fd(0x1ba)](_0x1ce5fd(0x1ec));
        (_0x51938c["style"][_0x1ce5fd(0x1e2)] = _0x1ce5fd(0x1dd)),
          Swal["fire"]({
            icon: _0x1ce5fd(0x1f0),
            title: _0x1ce5fd(0x1c5),
            text: _0x1ce5fd(0x1dc),
          })[_0x1ce5fd(0x1bb)]((_0x3fd7f3) => {
            const _0x4678e8 = _0x1ce5fd;
            _0x3fd7f3["value"] &&
              (setTimeout(() => {
                const _0x43464f = a0_0x1bfe;
                _0x51938c[_0x43464f(0x1d0)][_0x43464f(0x1e2)] =
                  _0x43464f(0x1d9);
              }, 0x12c),
              (_0x51938c[_0x4678e8(0x1d0)]["animation"] =
                "fadeIn\x200.3s\x20ease-in-out"));
          });
      },
    });
  });
});
