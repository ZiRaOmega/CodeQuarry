const a0_0x7daf1a=a0_0x1400;function a0_0x1400(_0x5d1df7,_0x1deb0e){const _0x5e771d=a0_0x5e77();return a0_0x1400=function(_0x140018,_0x245fd9){_0x140018=_0x140018-0x135;let _0x419c1f=_0x5e771d[_0x140018];return _0x419c1f;},a0_0x1400(_0x5d1df7,_0x1deb0e);}(function(_0x3e3bd1,_0x5baa30){const _0xb5861b=a0_0x1400,_0x5c22de=_0x3e3bd1();while(!![]){try{const _0x449267=parseInt(_0xb5861b(0x148))/0x1*(-parseInt(_0xb5861b(0x13d))/0x2)+-parseInt(_0xb5861b(0x166))/0x3*(-parseInt(_0xb5861b(0x14a))/0x4)+parseInt(_0xb5861b(0x163))/0x5*(-parseInt(_0xb5861b(0x139))/0x6)+parseInt(_0xb5861b(0x15a))/0x7*(-parseInt(_0xb5861b(0x13a))/0x8)+parseInt(_0xb5861b(0x13e))/0x9*(parseInt(_0xb5861b(0x157))/0xa)+-parseInt(_0xb5861b(0x179))/0xb*(parseInt(_0xb5861b(0x145))/0xc)+parseInt(_0xb5861b(0x159))/0xd;if(_0x449267===_0x5baa30)break;else _0x5c22de['push'](_0x5c22de['shift']());}catch(_0x21bd65){_0x5c22de['push'](_0x5c22de['shift']());}}}(a0_0x5e77,0x67f9a));let registerLastName=document[a0_0x7daf1a(0x167)](a0_0x7daf1a(0x171)),registerFirstName=document['getElementById'](a0_0x7daf1a(0x17a)),registerUsername=document[a0_0x7daf1a(0x167)](a0_0x7daf1a(0x13b)),registerEmail=document[a0_0x7daf1a(0x167)]('registerEmail'),registerPassword=document[a0_0x7daf1a(0x167)](a0_0x7daf1a(0x140)),registerPasswordConfirmation=document[a0_0x7daf1a(0x167)](a0_0x7daf1a(0x162)),registerForm=document[a0_0x7daf1a(0x167)](a0_0x7daf1a(0x137)),registerSubmit=document[a0_0x7daf1a(0x167)](a0_0x7daf1a(0x16c)),contentAlert=document[a0_0x7daf1a(0x167)]('contentAlert');$(document)[a0_0x7daf1a(0x144)](function(){const _0x1430af=a0_0x7daf1a;$(_0x1430af(0x174))['submit'](function(_0x48fbdb){const _0x10729a=_0x1430af;_0x48fbdb[_0x10729a(0x169)]();const _0x4d4475=[{'value':registerLastName[_0x10729a(0x136)],'name':_0x10729a(0x15f)},{'value':registerFirstName[_0x10729a(0x136)],'name':_0x10729a(0x178)},{'value':registerUsername[_0x10729a(0x136)],'name':_0x10729a(0x16f)},{'value':registerEmail['value'],'name':_0x10729a(0x170)},{'value':registerPassword[_0x10729a(0x136)],'name':'Password'},{'value':registerPasswordConfirmation['value'],'name':_0x10729a(0x13c)}];let _0x30561f='';_0x4d4475[_0x10729a(0x158)](_0x2b5669=>{const _0x2b8ab1=_0x10729a;_0x2b5669[_0x2b8ab1(0x136)]==''&&(_0x30561f+=_0x2b5669[_0x2b8ab1(0x155)]+_0x2b8ab1(0x13f));});registerPassword['value']!==registerPasswordConfirmation['value']&&registerPassword['value']!==''&&registerPasswordConfirmation[_0x10729a(0x136)]!==''&&(_0x30561f+=_0x10729a(0x138));if(registerPassword[_0x10729a(0x136)]!==''){var _0xb5e40e=/^(?=.*[0-9])(?=.*[^a-zA-Z0-9])[a-zA-Z0-9!@#$%^&*()_+=\-`~\[\]{};':"\\|,.<>\/?]{8,}$/;!_0xb5e40e[_0x10729a(0x135)](registerPassword['value'])&&(_0x30561f+=_0x10729a(0x161));}registerEmail['value']!==''&&!/^[^@]+@[^@]+\.[^@]+$/['test'](registerEmail['value'])&&(_0x30561f+=_0x10729a(0x176)),_0x48fbdb[_0x10729a(0x169)](),!_0x30561f?fetch(_0x10729a(0x16a),{'method':_0x10729a(0x151),'headers':{'Content-Type':_0x10729a(0x153)},'body':new URLSearchParams(new FormData(registerForm))})[_0x10729a(0x160)](_0x5c27fb=>_0x5c27fb[_0x10729a(0x16e)]())['then'](_0x5d3184=>{const _0x123a98=_0x10729a;if(_0x5d3184['status']==='success'){let _0x418b6c=document[_0x123a98(0x167)](_0x123a98(0x168));_0x418b6c[_0x123a98(0x142)][_0x123a98(0x150)]=_0x123a98(0x147),Swal[_0x123a98(0x156)]({'title':_0x123a98(0x14c),'text':_0x5d3184[_0x123a98(0x15d)],'icon':_0x123a98(0x177),'confirmButtonText':'OK'})[_0x123a98(0x160)](_0x4f2e2b=>{const _0x3c2f15=_0x123a98;_0x4f2e2b['value']&&(window['location'][_0x3c2f15(0x149)]=_0x3c2f15(0x14b));});}else throw new Error(_0x5d3184[_0x123a98(0x15d)]||_0x123a98(0x15b));})[_0x10729a(0x154)](_0x27f4f3=>{const _0x30fbef=_0x10729a;console[_0x30fbef(0x15c)]('Error:',_0x27f4f3),contentAlert[_0x30fbef(0x164)]=_0x27f4f3[_0x30fbef(0x15d)];}):contentAlert[_0x10729a(0x164)]=_0x30561f;});});let login=document[a0_0x7daf1a(0x167)](a0_0x7daf1a(0x173)),usernameOrEmailLogin=document['getElementById'](a0_0x7daf1a(0x175)),passwordLogin=document[a0_0x7daf1a(0x167)]('loginPassword'),contentAlertLogin=document[a0_0x7daf1a(0x167)](a0_0x7daf1a(0x16d));function a0_0x5e77(){const _0x483d62=['usernameOrEmailLogin','Email\x20must\x20be\x20a\x20valid\x20address.<br>','success','FirstName','22pCwAOa','registerFirstName','test','value','registerForm','Passwords\x20do\x20not\x20match.<br>','6pMOnaA','2299032kUJfWt','registerUsername','Password\x20Confirmation','2TgVSsn','18ozAeWo','\x20is\x20required.<br>','registerPassword','loginBlock','style','#loginForm','ready','3475164vzeGUR','animation','none','10684iQpnRg','href','213904hIIDFu','/codeQuarry','Thank\x20You!','submit','/login','param','display','POST','Oops...','application/x-www-form-urlencoded','catch','name','fire','1774310wcfgQt','forEach','12123033Kbpcqi','14OnRdYY','Registration\x20failed','error','message','val','LastName','then','Password\x20must\x20be\x20at\x20least\x208\x20characters\x20long,\x20contain\x20at\x20least\x20one\x20number,\x20and\x20contain\x20at\x20least\x20one\x20special\x20character.<br>','registerPasswordConfirmation','623465BGPIOu','innerHTML','fadeIn\x200.3s\x20ease-in-out','24jlUcND','getElementById','registerBlock','preventDefault','/register','flex','registerSubmit','contentAlertLogin','json','Username','Email','registerLastName','Invalid\x20login\x20credentials!','login','#registerForm'];a0_0x5e77=function(){return _0x483d62;};return a0_0x5e77();}$(document)[a0_0x7daf1a(0x144)](function(){const _0x6aa1b5=a0_0x7daf1a;$(_0x6aa1b5(0x143))[_0x6aa1b5(0x14d)](function(_0x3c264e){const _0x4d5aaa=_0x6aa1b5;_0x3c264e[_0x4d5aaa(0x169)]();var _0x50f158={'usernameOrEmailLogin':$('#usernameOrEmailLogin')[_0x4d5aaa(0x15e)](),'passwordLogin':$('#loginPassword')[_0x4d5aaa(0x15e)]()};$['ajax']({'type':'POST','url':_0x4d5aaa(0x14e),'data':$[_0x4d5aaa(0x14f)](_0x50f158),'contentType':'application/x-www-form-urlencoded','success':function(_0x634802){const _0x35febc=_0x4d5aaa;if(_0x634802['status']==='success')window['location'][_0x35febc(0x149)]=_0x35febc(0x14b);else{let _0x50396f=document[_0x35febc(0x167)](_0x35febc(0x141));_0x50396f['style'][_0x35febc(0x150)]=_0x35febc(0x147),Swal[_0x35febc(0x156)]({'icon':_0x35febc(0x15c),'title':_0x35febc(0x152),'text':_0x634802[_0x35febc(0x15d)]||_0x35febc(0x172),'confirmButtonText':'OK'})['then'](_0x5de68d=>{_0x5de68d['value']&&setTimeout(()=>{const _0x2644f6=a0_0x1400;_0x50396f['style'][_0x2644f6(0x150)]=_0x2644f6(0x16b);},0x1f4);});}},'error':function(){const _0x2737f6=_0x4d5aaa;let _0x575cdf=document[_0x2737f6(0x167)](_0x2737f6(0x141));_0x575cdf['style'][_0x2737f6(0x150)]=_0x2737f6(0x147),Swal['fire']({'icon':_0x2737f6(0x15c),'title':_0x2737f6(0x152),'text':_0x2737f6(0x172)})[_0x2737f6(0x160)](_0x2959c7=>{const _0x3c75af=_0x2737f6;_0x2959c7['value']&&(setTimeout(()=>{const _0xfbad1a=a0_0x1400;_0x575cdf[_0xfbad1a(0x142)][_0xfbad1a(0x150)]='flex';},0x12c),_0x575cdf[_0x3c75af(0x142)][_0x3c75af(0x146)]=_0x3c75af(0x165));});}});});});