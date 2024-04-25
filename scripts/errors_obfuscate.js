const a0_0x4f81f2 = a0_0x5462;
(function (_0x4c2c74, _0x547410) {
  const _0x437db5 = a0_0x5462,
    _0x521150 = _0x4c2c74();
  while (!![]) {
    try {
      const _0x400a8a =
        parseInt(_0x437db5(0x1f7)) / 0x1 +
        parseInt(_0x437db5(0x21b)) / 0x2 +
        parseInt(_0x437db5(0x228)) / 0x3 +
        (-parseInt(_0x437db5(0x1f3)) / 0x4) *
          (-parseInt(_0x437db5(0x21f)) / 0x5) +
        (-parseInt(_0x437db5(0x205)) / 0x6) *
          (-parseInt(_0x437db5(0x20f)) / 0x7) +
        parseInt(_0x437db5(0x202)) / 0x8 +
        (parseInt(_0x437db5(0x218)) / 0x9) *
          (-parseInt(_0x437db5(0x229)) / 0xa);
      if (_0x400a8a === _0x547410) break;
      else _0x521150["push"](_0x521150["shift"]());
    } catch (_0x1780a8) {
      _0x521150["push"](_0x521150["shift"]());
    }
  }
})(a0_0x2f90, 0xbfe43);
let registerLastName = document[a0_0x4f81f2(0x22e)]("registerLastName"),
  registerFirstName = document[a0_0x4f81f2(0x22e)](a0_0x4f81f2(0x208)),
  registerUsername = document["getElementById"](a0_0x4f81f2(0x209)),
  registerEmail = document[a0_0x4f81f2(0x22e)](a0_0x4f81f2(0x226)),
  registerPassword = document[a0_0x4f81f2(0x22e)](a0_0x4f81f2(0x21e)),
  registerPasswordConfirmation = document["getElementById"](a0_0x4f81f2(0x20c)),
  registerForm = document[a0_0x4f81f2(0x22e)](a0_0x4f81f2(0x220)),
  registerSubmit = document["getElementById"]("registerSubmit"),
  contentAlert = document[a0_0x4f81f2(0x22e)]("contentAlert");
function a0_0x2f90() {
  const _0x3a604e = [
    "registerFirstName",
    "registerUsername",
    "Password",
    "animation",
    "registerPasswordConfirmation",
    "Error:",
    "catch",
    "25459ZPwVve",
    "Email",
    "#registerForm",
    "preventDefault",
    "Thank\x20You!",
    "status",
    "Registration\x20failed",
    "LastName",
    "style",
    "9ckwQxR",
    "forEach",
    "usernameOrEmailLogin",
    "2161350IDEAaT",
    "none",
    "val",
    "registerPassword",
    "5sRSNuH",
    "registerForm",
    "ajax",
    "Username",
    "login.html",
    "fire",
    "POST",
    "registerEmail",
    "FirstName",
    "1596627lSCXNJ",
    "38560830rcShBQ",
    "loginBlock",
    "Oops...",
    "Password\x20must\x20be\x20at\x20least\x208\x20characters\x20long,\x20contain\x20at\x20least\x20one\x20number,\x20and\x20contain\x20at\x20least\x20one\x20special\x20character.<br>",
    "name",
    "getElementById",
    "param",
    "\x20is\x20required.<br>",
    "display",
    "location",
    "value",
    "2515124ufYhXJ",
    "test",
    "/register",
    "fadeIn\x200.3s\x20ease-in-out",
    "1013513ZnAsVP",
    "Invalid\x20login\x20credentials!",
    "/login",
    "application/x-www-form-urlencoded",
    "success",
    "flex",
    "innerHTML",
    "loginPassword",
    "error",
    "href",
    "submit",
    "6847120HgjZtx",
    "json",
    "ready",
    "876SGgcuG",
    "message",
    "then",
  ];
  a0_0x2f90 = function () {
    return _0x3a604e;
  };
  return a0_0x2f90();
}
function a0_0x5462(_0x35303b, _0x2169d0) {
  const _0x2f90d5 = a0_0x2f90();
  return (
    (a0_0x5462 = function (_0x54620c, _0x4312fe) {
      _0x54620c = _0x54620c - 0x1f3;
      let _0x17d2d8 = _0x2f90d5[_0x54620c];
      return _0x17d2d8;
    }),
    a0_0x5462(_0x35303b, _0x2169d0)
  );
}
$(document)[a0_0x4f81f2(0x204)](function () {
  const _0x396abc = a0_0x4f81f2;
  $(_0x396abc(0x211))["submit"](function (_0x1e7805) {
    const _0x10c17d = _0x396abc;
    _0x1e7805["preventDefault"]();
    const _0x2dcf0b = [
      { value: registerLastName[_0x10c17d(0x233)], name: _0x10c17d(0x216) },
      { value: registerFirstName["value"], name: _0x10c17d(0x227) },
      { value: registerUsername[_0x10c17d(0x233)], name: _0x10c17d(0x222) },
      { value: registerEmail["value"], name: _0x10c17d(0x210) },
      { value: registerPassword[_0x10c17d(0x233)], name: _0x10c17d(0x20a) },
      {
        value: registerPasswordConfirmation[_0x10c17d(0x233)],
        name: "Password\x20Confirmation",
      },
    ];
    let _0x52529f = "";
    _0x2dcf0b[_0x10c17d(0x219)]((_0x593970) => {
      const _0x2f7694 = _0x10c17d;
      _0x593970[_0x2f7694(0x233)] == "" &&
        (_0x52529f += _0x593970[_0x2f7694(0x22d)] + _0x2f7694(0x230));
    });
    registerPassword[_0x10c17d(0x233)] !==
      registerPasswordConfirmation[_0x10c17d(0x233)] &&
      registerPassword[_0x10c17d(0x233)] !== "" &&
      registerPasswordConfirmation[_0x10c17d(0x233)] !== "" &&
      (_0x52529f += "Passwords\x20do\x20not\x20match.<br>");
    if (registerPassword["value"] !== "") {
      var _0x2c723c =
        /^(?=.*[0-9])(?=.*[^a-zA-Z0-9])[a-zA-Z0-9!@#$%^&*()_+=\-`~\[\]{};':"\\|,.<>\/?]{8,}$/;
      !_0x2c723c[_0x10c17d(0x1f4)](registerPassword[_0x10c17d(0x233)]) &&
        (_0x52529f += _0x10c17d(0x22c));
    }
    registerEmail["value"] !== "" &&
      !/^[^@]+@[^@]+\.[^@]+$/[_0x10c17d(0x1f4)](
        registerEmail[_0x10c17d(0x233)]
      ) &&
      (_0x52529f += "Email\x20must\x20be\x20a\x20valid\x20address.<br>"),
      _0x1e7805[_0x10c17d(0x212)](),
      !_0x52529f
        ? fetch(_0x10c17d(0x1f5), {
            method: _0x10c17d(0x225),
            headers: { "Content-Type": _0x10c17d(0x1fa) },
            body: new URLSearchParams(new FormData(registerForm)),
          })
            ["then"]((_0x9b4115) => _0x9b4115[_0x10c17d(0x203)]())
            [_0x10c17d(0x207)]((_0x2d9239) => {
              const _0x10a934 = _0x10c17d;
              if (_0x2d9239[_0x10a934(0x214)] === _0x10a934(0x1fb)) {
                let _0x5963a0 = document[_0x10a934(0x22e)]("registerBlock");
                (_0x5963a0[_0x10a934(0x217)][_0x10a934(0x231)] =
                  _0x10a934(0x21c)),
                  Swal[_0x10a934(0x224)]({
                    title: _0x10a934(0x213),
                    text: _0x2d9239[_0x10a934(0x206)],
                    icon: _0x10a934(0x1fb),
                    confirmButtonText: "OK",
                  })[_0x10a934(0x207)]((_0x365c10) => {
                    const _0x24e015 = _0x10a934;
                    _0x365c10[_0x24e015(0x233)] &&
                      (window[_0x24e015(0x232)][_0x24e015(0x200)] =
                        _0x24e015(0x223));
                  });
              } else throw new Error(_0x2d9239["message"] || _0x10a934(0x215));
            })
            [_0x10c17d(0x20e)]((_0x322163) => {
              const _0x310c5d = _0x10c17d;
              console[_0x310c5d(0x1ff)](_0x310c5d(0x20d), _0x322163),
                (contentAlert[_0x310c5d(0x1fd)] = _0x322163[_0x310c5d(0x206)]);
            })
        : (contentAlert["innerHTML"] = _0x52529f);
  });
});
let login = document[a0_0x4f81f2(0x22e)]("login"),
  usernameOrEmailLogin = document[a0_0x4f81f2(0x22e)](a0_0x4f81f2(0x21a)),
  passwordLogin = document[a0_0x4f81f2(0x22e)](a0_0x4f81f2(0x1fe)),
  contentAlertLogin = document[a0_0x4f81f2(0x22e)]("contentAlertLogin");
$(document)[a0_0x4f81f2(0x204)](function () {
  const _0x276746 = a0_0x4f81f2;
  $("#loginForm")[_0x276746(0x201)](function (_0x49c5fd) {
    const _0x3d930f = _0x276746;
    _0x49c5fd[_0x3d930f(0x212)]();
    var _0x146ef2 = {
      usernameOrEmailLogin: $("#usernameOrEmailLogin")["val"](),
      passwordLogin: $("#loginPassword")[_0x3d930f(0x21d)](),
    };
    $[_0x3d930f(0x221)]({
      type: _0x3d930f(0x225),
      url: _0x3d930f(0x1f9),
      data: $[_0x3d930f(0x22f)](_0x146ef2),
      contentType: _0x3d930f(0x1fa),
      success: function (_0x4ce6dc) {
        const _0x2a801c = _0x3d930f;
        if (_0x4ce6dc["status"] === _0x2a801c(0x1fb))
          window[_0x2a801c(0x232)]["href"] = "/codeQuarry";
        else {
          let _0x33b136 = document["getElementById"](_0x2a801c(0x22a));
          (_0x33b136[_0x2a801c(0x217)][_0x2a801c(0x231)] = _0x2a801c(0x21c)),
            Swal["fire"]({
              icon: _0x2a801c(0x1ff),
              title: _0x2a801c(0x22b),
              text: _0x4ce6dc["message"] || _0x2a801c(0x1f8),
              confirmButtonText: "OK",
            })["then"]((_0xda2678) => {
              const _0x3781c4 = _0x2a801c;
              _0xda2678[_0x3781c4(0x233)] &&
                setTimeout(() => {
                  const _0x5e2bdf = _0x3781c4;
                  _0x33b136[_0x5e2bdf(0x217)][_0x5e2bdf(0x231)] =
                    _0x5e2bdf(0x1fc);
                }, 0x1f4);
            });
        }
      },
      error: function () {
        const _0xe993d5 = _0x3d930f;
        let _0x475886 = document[_0xe993d5(0x22e)]("loginBlock");
        (_0x475886["style"]["display"] = "none"),
          Swal[_0xe993d5(0x224)]({
            icon: "error",
            title: _0xe993d5(0x22b),
            text: _0xe993d5(0x1f8),
          })[_0xe993d5(0x207)]((_0x1d3dd9) => {
            const _0x5c462f = _0xe993d5;
            _0x1d3dd9[_0x5c462f(0x233)] &&
              (setTimeout(() => {
                const _0x82ce4e = _0x5c462f;
                _0x475886[_0x82ce4e(0x217)][_0x82ce4e(0x231)] = "flex";
              }, 0x12c),
              (_0x475886["style"][_0x5c462f(0x20b)] = _0x5c462f(0x1f6)));
          });
      },
    });
  });
});
