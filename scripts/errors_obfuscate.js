const a0_0x145fca=a0_0x21a8;(function(_0xbd5a8c,_0xc1503b){const _0x3fa68c=a0_0x21a8,_0x437358=_0xbd5a8c();while(!![]){try{const _0x125dca=-parseInt(_0x3fa68c(0x211))/0x1*(parseInt(_0x3fa68c(0x206))/0x2)+parseInt(_0x3fa68c(0x21d))/0x3+parseInt(_0x3fa68c(0x1f3))/0x4*(-parseInt(_0x3fa68c(0x219))/0x5)+parseInt(_0x3fa68c(0x209))/0x6*(parseInt(_0x3fa68c(0x20a))/0x7)+-parseInt(_0x3fa68c(0x229))/0x8+parseInt(_0x3fa68c(0x21a))/0x9*(-parseInt(_0x3fa68c(0x216))/0xa)+-parseInt(_0x3fa68c(0x22b))/0xb*(-parseInt(_0x3fa68c(0x213))/0xc);if(_0x125dca===_0xc1503b)break;else _0x437358['push'](_0x437358['shift']());}catch(_0x5cc9b8){_0x437358['push'](_0x437358['shift']());}}}(a0_0xd1f3,0x4bef7));function a0_0x21a8(_0x5a160e,_0x4ce66e){const _0xd1f337=a0_0xd1f3();return a0_0x21a8=function(_0x21a803,_0x23e75c){_0x21a803=_0x21a803-0x1ed;let _0x17ac70=_0xd1f337[_0x21a803];return _0x17ac70;},a0_0x21a8(_0x5a160e,_0x4ce66e);}let registerName=document['getElementById'](a0_0x145fca(0x20e)),registerUsername=document[a0_0x145fca(0x223)](a0_0x145fca(0x203)),registerEmail=document[a0_0x145fca(0x223)]('registerEmail'),registerPassword=document[a0_0x145fca(0x223)](a0_0x145fca(0x207)),registerPasswordConfirmation=document['getElementById'](a0_0x145fca(0x22d)),registerForm=document[a0_0x145fca(0x223)](a0_0x145fca(0x201)),registerSubmit=document['getElementById'](a0_0x145fca(0x20d)),contentAlert=document['getElementById']('contentAlert');$(document)[a0_0x145fca(0x1f2)](function(){$('#registerForm')['submit'](function(_0x4f7579){const _0x3aa79c=a0_0x21a8;_0x4f7579['preventDefault']();const _0x35a677=[{'value':registerName[_0x3aa79c(0x1f9)],'name':_0x3aa79c(0x1fe)},{'value':registerUsername['value'],'name':_0x3aa79c(0x204)},{'value':registerEmail[_0x3aa79c(0x1f9)],'name':_0x3aa79c(0x1ff)},{'value':registerPassword[_0x3aa79c(0x1f9)],'name':'Password'},{'value':registerPasswordConfirmation[_0x3aa79c(0x1f9)],'name':'Password\x20Confirmation'}];let _0x4fe4f7='';_0x35a677[_0x3aa79c(0x205)](_0x8429f2=>{const _0x282faa=_0x3aa79c;_0x8429f2[_0x282faa(0x1f9)]==''&&(_0x4fe4f7+=_0x8429f2[_0x282faa(0x1f8)]+_0x282faa(0x226));}),registerPassword[_0x3aa79c(0x1f9)]!==registerPasswordConfirmation[_0x3aa79c(0x1f9)]&&registerPassword[_0x3aa79c(0x1f9)]!==''&&registerPasswordConfirmation['value']!==''&&(_0x4fe4f7+='Passwords\x20do\x20not\x20match.<br>'),registerPassword[_0x3aa79c(0x1f9)]!==''&&!/^(?=.*\d)[a-zA-Z0-9]{8,}$/['test'](registerPassword['value'])&&(_0x4fe4f7+='Password\x20must\x20be\x20at\x20least\x208\x20characters\x20long\x20and\x20contain\x20at\x20least\x20one\x20number.<br>'),registerEmail[_0x3aa79c(0x1f9)]!==''&&!/^[^@]+@[^@]+\.[^@]+$/[_0x3aa79c(0x200)](registerEmail['value'])&&(_0x4fe4f7+='Email\x20must\x20be\x20a\x20valid\x20address.<br>'),_0x4f7579[_0x3aa79c(0x220)](),!_0x4fe4f7?fetch(_0x3aa79c(0x1fa),{'method':_0x3aa79c(0x217),'headers':{'Content-Type':_0x3aa79c(0x1ee)},'body':new URLSearchParams(new FormData(registerForm))})[_0x3aa79c(0x208)](_0x157da4=>_0x157da4['json']())['then'](_0x288812=>{const _0x28f3e1=_0x3aa79c;if(_0x288812[_0x28f3e1(0x1ed)]===_0x28f3e1(0x20f)){let _0x15fd4f=document['getElementById'](_0x28f3e1(0x21e));_0x15fd4f[_0x28f3e1(0x22c)][_0x28f3e1(0x1f7)]=_0x28f3e1(0x21b),Swal[_0x28f3e1(0x21f)]({'title':'Thank\x20You!','text':_0x288812[_0x28f3e1(0x1ef)],'icon':'success','confirmButtonText':'OK'})[_0x28f3e1(0x208)](_0x4dd3e8=>{const _0x1860db=_0x28f3e1;_0x4dd3e8[_0x1860db(0x1f9)]&&(window[_0x1860db(0x1f6)][_0x1860db(0x20b)]=_0x1860db(0x1fc));});}else throw new Error(_0x288812[_0x28f3e1(0x1ef)]||_0x28f3e1(0x1fb));})[_0x3aa79c(0x1f5)](_0x1738a1=>{const _0x53d350=_0x3aa79c;console['error'](_0x53d350(0x218),_0x1738a1),contentAlert[_0x53d350(0x1f4)]=_0x1738a1[_0x53d350(0x1ef)];}):contentAlert[_0x3aa79c(0x1f4)]=_0x4fe4f7;});});let login=document[a0_0x145fca(0x223)](a0_0x145fca(0x1f0)),usernameOrEmailLogin=document[a0_0x145fca(0x223)](a0_0x145fca(0x228)),passwordLogin=document['getElementById']('loginPassword'),contentAlertLogin=document[a0_0x145fca(0x223)]('contentAlertLogin');$(document)[a0_0x145fca(0x1f2)](function(){const _0x88f836=a0_0x145fca;$(_0x88f836(0x1f1))[_0x88f836(0x227)](function(_0x33f69e){const _0x5d5270=_0x88f836;_0x33f69e['preventDefault']();var _0x521aab={'usernameOrEmailLogin':$(_0x5d5270(0x212))[_0x5d5270(0x1fd)](),'passwordLogin':$(_0x5d5270(0x22e))[_0x5d5270(0x1fd)]()};$[_0x5d5270(0x20c)]({'type':_0x5d5270(0x217),'url':_0x5d5270(0x214),'data':$[_0x5d5270(0x225)](_0x521aab),'contentType':_0x5d5270(0x1ee),'success':function(_0xf1f4e){const _0x1fdac2=_0x5d5270;if(_0xf1f4e[_0x1fdac2(0x1ed)]==='success')window[_0x1fdac2(0x1f6)][_0x1fdac2(0x20b)]=_0x1fdac2(0x202);else{let _0x16610d=document['getElementById'](_0x1fdac2(0x222));_0x16610d[_0x1fdac2(0x22c)][_0x1fdac2(0x1f7)]=_0x1fdac2(0x21b),Swal[_0x1fdac2(0x21f)]({'icon':_0x1fdac2(0x22a),'title':_0x1fdac2(0x224),'text':_0xf1f4e[_0x1fdac2(0x1ef)]||_0x1fdac2(0x215),'confirmButtonText':'OK'})[_0x1fdac2(0x208)](_0x577bf9=>{const _0xfdb4ab=_0x1fdac2;_0x577bf9[_0xfdb4ab(0x1f9)]&&setTimeout(()=>{const _0x55d977=_0xfdb4ab;_0x16610d[_0x55d977(0x22c)][_0x55d977(0x1f7)]=_0x55d977(0x21c);},0x1f4);});}},'error':function(){const _0x5b6d24=_0x5d5270;let _0x929b71=document[_0x5b6d24(0x223)]('loginBlock');_0x929b71[_0x5b6d24(0x22c)][_0x5b6d24(0x1f7)]=_0x5b6d24(0x21b),Swal['fire']({'icon':'error','title':'Oops...','text':_0x5b6d24(0x215)})[_0x5b6d24(0x208)](_0x39d814=>{const _0x35d66f=_0x5b6d24;_0x39d814['value']&&(setTimeout(()=>{const _0x2a47e5=a0_0x21a8;_0x929b71[_0x2a47e5(0x22c)][_0x2a47e5(0x1f7)]=_0x2a47e5(0x21c);},0x12c),_0x929b71[_0x35d66f(0x22c)][_0x35d66f(0x210)]=_0x35d66f(0x221));});}});});});function a0_0xd1f3(){const _0x4b3ae6=['119364hcgZdZ','registerBlock','fire','preventDefault','fadeIn\x200.3s\x20ease-in-out','loginBlock','getElementById','Oops...','param','\x20is\x20required.<br>','submit','usernameOrEmailLogin','463744mCrasc','error','1124882CthENG','style','registerPasswordConfirmation','#loginPassword','status','application/x-www-form-urlencoded','message','login','#loginForm','ready','4948BIBUDW','innerHTML','catch','location','display','name','value','/register','Registration\x20failed','login.html','val','Name','Email','test','registerForm','/codeQuarry','registerUsername','Username','forEach','10GXEquX','registerPassword','then','332658RDQCyI','49lEqZEW','href','ajax','registerSubmit','registerName','success','animation','22303JsVTAt','#usernameOrEmailLogin','96esTTbK','/login','Invalid\x20login\x20credentials!','20qtXCYJ','POST','Error:','2305jJqErq','878463CaLYYl','none','flex'];a0_0xd1f3=function(){return _0x4b3ae6;};return a0_0xd1f3();}