const a0_0xe4fae9=a0_0x51d2;(function(_0x176ed8,_0x297d63){const _0xcb1678=a0_0x51d2,_0x4a9b6f=_0x176ed8();while(!![]){try{const _0x6f78cd=parseInt(_0xcb1678(0x1bb))/0x1*(parseInt(_0xcb1678(0x196))/0x2)+parseInt(_0xcb1678(0x198))/0x3*(-parseInt(_0xcb1678(0x192))/0x4)+parseInt(_0xcb1678(0x18e))/0x5*(-parseInt(_0xcb1678(0x1a9))/0x6)+parseInt(_0xcb1678(0x193))/0x7*(parseInt(_0xcb1678(0x1a4))/0x8)+parseInt(_0xcb1678(0x1b7))/0x9*(-parseInt(_0xcb1678(0x191))/0xa)+-parseInt(_0xcb1678(0x1bd))/0xb*(parseInt(_0xcb1678(0x1a3))/0xc)+parseInt(_0xcb1678(0x1ac))/0xd;if(_0x6f78cd===_0x297d63)break;else _0x4a9b6f['push'](_0x4a9b6f['shift']());}catch(_0x457639){_0x4a9b6f['push'](_0x4a9b6f['shift']());}}}(a0_0x5276,0x22786));let registerLastName=document['getElementById']('registerLastName'),registerFirstName=document[a0_0xe4fae9(0x18a)]('registerFirstName'),registerUsername=document[a0_0xe4fae9(0x18a)](a0_0xe4fae9(0x1a2)),registerEmail=document[a0_0xe4fae9(0x18a)]('registerEmail'),registerPassword=document['getElementById'](a0_0xe4fae9(0x19f)),registerPasswordConfirmation=document['getElementById'](a0_0xe4fae9(0x18d)),registerForm=document['getElementById'](a0_0xe4fae9(0x1ce)),registerSubmit=document['getElementById'](a0_0xe4fae9(0x1c3)),contentAlert=document[a0_0xe4fae9(0x18a)](a0_0xe4fae9(0x1ab));$(document)[a0_0xe4fae9(0x1b2)](function(){const _0x3ba9df=a0_0xe4fae9;$('#registerForm')[_0x3ba9df(0x1cd)](function(_0x234cda){const _0x59fb53=_0x3ba9df;_0x234cda[_0x59fb53(0x1bc)]();const _0x416ed1=[{'value':registerLastName[_0x59fb53(0x1b6)],'name':_0x59fb53(0x1c0)},{'value':registerFirstName[_0x59fb53(0x1b6)],'name':_0x59fb53(0x1c6)},{'value':registerUsername[_0x59fb53(0x1b6)],'name':_0x59fb53(0x1b8)},{'value':registerEmail[_0x59fb53(0x1b6)],'name':_0x59fb53(0x19a)},{'value':registerPassword[_0x59fb53(0x1b6)],'name':_0x59fb53(0x1b0)},{'value':registerPasswordConfirmation[_0x59fb53(0x1b6)],'name':_0x59fb53(0x190)}];let _0xb92268='';_0x416ed1[_0x59fb53(0x197)](_0x3c6c64=>{const _0x34c55e=_0x59fb53;_0x3c6c64[_0x34c55e(0x1b6)]==''&&(_0xb92268+=_0x3c6c64[_0x34c55e(0x188)]+_0x34c55e(0x19e));});registerPassword[_0x59fb53(0x1b6)]!==registerPasswordConfirmation[_0x59fb53(0x1b6)]&&registerPassword[_0x59fb53(0x1b6)]!==''&&registerPasswordConfirmation['value']!==''&&(_0xb92268+=_0x59fb53(0x18f));if(registerPassword['value']!==''){var _0x284119=/^(?=.*[0-9])(?=.*[^a-zA-Z0-9])[a-zA-Z0-9!@#$%^&*()_+=\-`~\[\]{};':"\\|,.<>\/?]{8,}$/;!_0x284119[_0x59fb53(0x1c1)](registerPassword[_0x59fb53(0x1b6)])&&(_0xb92268+='Password\x20must\x20be\x20at\x20least\x208\x20characters\x20long,\x20contain\x20at\x20least\x20one\x20number,\x20and\x20contain\x20at\x20least\x20one\x20special\x20character.<br>');}registerEmail[_0x59fb53(0x1b6)]!==''&&!/^[^@]+@[^@]+\.[^@]+$/[_0x59fb53(0x1c1)](registerEmail[_0x59fb53(0x1b6)])&&(_0xb92268+=_0x59fb53(0x1cc)),_0x234cda['preventDefault'](),!_0xb92268?fetch(_0x59fb53(0x1a6),{'method':_0x59fb53(0x1a1),'headers':{'Content-Type':_0x59fb53(0x1b3)},'body':new URLSearchParams(new FormData(registerForm))})[_0x59fb53(0x1af)](_0x14d083=>_0x14d083[_0x59fb53(0x1bf)]())['then'](_0x4b3592=>{const _0x334b73=_0x59fb53;if(_0x4b3592[_0x334b73(0x1a8)]===_0x334b73(0x1be)){let _0x3c86c3=document[_0x334b73(0x18a)](_0x334b73(0x18c));_0x3c86c3[_0x334b73(0x1c8)][_0x334b73(0x199)]=_0x334b73(0x1c4),Swal[_0x334b73(0x195)]({'title':_0x334b73(0x1b9),'text':_0x4b3592[_0x334b73(0x1b1)],'icon':_0x334b73(0x1be),'confirmButtonText':'OK'})[_0x334b73(0x1af)](_0x5aa216=>{const _0xd6fb15=_0x334b73;_0x5aa216['value']&&(window[_0xd6fb15(0x1b4)][_0xd6fb15(0x1c2)]=_0xd6fb15(0x1c7));});}else throw new Error(_0x4b3592[_0x334b73(0x1b1)]||_0x334b73(0x1c5));})['catch'](_0x2766ba=>{const _0x2b361c=_0x59fb53;console[_0x2b361c(0x1ae)](_0x2b361c(0x1c9),_0x2766ba),contentAlert[_0x2b361c(0x1aa)]=_0x2766ba[_0x2b361c(0x1b1)];}):contentAlert['innerHTML']=_0xb92268;});});let login=document[a0_0xe4fae9(0x18a)](a0_0xe4fae9(0x1a0)),usernameOrEmailLogin=document[a0_0xe4fae9(0x18a)]('usernameOrEmailLogin'),passwordLogin=document['getElementById'](a0_0xe4fae9(0x1a7)),contentAlertLogin=document['getElementById'](a0_0xe4fae9(0x19b));function a0_0x5276(){const _0x107ac3=['flex','value','9RASqyK','Username','Thank\x20You!','animation','42691BVvPla','preventDefault','11hGxUun','success','json','LastName','test','href','registerSubmit','none','Registration\x20failed','FirstName','/codeQuarry','style','Error:','fadeIn\x200.3s\x20ease-in-out','val','Email\x20must\x20be\x20a\x20valid\x20address.<br>','submit','registerForm','name','loginBlock','getElementById','param','registerBlock','registerPasswordConfirmation','157015wONSYL','Passwords\x20do\x20not\x20match.<br>','Password\x20Confirmation','97910plnZNd','17288BLbiJN','32501EdCxeR','#usernameOrEmailLogin','fire','4zlNGhz','forEach','102djsqOv','display','Email','contentAlertLogin','/login','ajax','\x20is\x20required.<br>','registerPassword','login','POST','registerUsername','1809192iuDrVd','376OntTcx','Oops...','/register','loginPassword','status','24mXDZyK','innerHTML','contentAlert','3519152wEWxen','#loginForm','error','then','Password','message','ready','application/x-www-form-urlencoded','location'];a0_0x5276=function(){return _0x107ac3;};return a0_0x5276();}function a0_0x51d2(_0xdbe43e,_0x5a09af){const _0x5276e3=a0_0x5276();return a0_0x51d2=function(_0x51d2cc,_0x437825){_0x51d2cc=_0x51d2cc-0x188;let _0x471ee4=_0x5276e3[_0x51d2cc];return _0x471ee4;},a0_0x51d2(_0xdbe43e,_0x5a09af);}$(document)[a0_0xe4fae9(0x1b2)](function(){const _0x42144b=a0_0xe4fae9;$(_0x42144b(0x1ad))['submit'](function(_0xf73638){const _0x4fc99e=_0x42144b;_0xf73638[_0x4fc99e(0x1bc)]();var _0x54a921={'usernameOrEmailLogin':$(_0x4fc99e(0x194))[_0x4fc99e(0x1cb)](),'passwordLogin':$('#loginPassword')['val']()};$[_0x4fc99e(0x19d)]({'type':_0x4fc99e(0x1a1),'url':_0x4fc99e(0x19c),'data':$[_0x4fc99e(0x18b)](_0x54a921),'contentType':_0x4fc99e(0x1b3),'success':function(_0x41d81e){const _0x963f91=_0x4fc99e;if(_0x41d81e[_0x963f91(0x1a8)]===_0x963f91(0x1be))window[_0x963f91(0x1b4)]['href']=_0x963f91(0x1c7);else{let _0x3f5120=document[_0x963f91(0x18a)](_0x963f91(0x189));_0x3f5120[_0x963f91(0x1c8)]['display']=_0x963f91(0x1c4),Swal[_0x963f91(0x195)]({'icon':'error','title':_0x963f91(0x1a5),'text':_0x41d81e['message']||'Invalid\x20login\x20credentials!','confirmButtonText':'OK'})[_0x963f91(0x1af)](_0x1910dc=>{const _0x8719f7=_0x963f91;_0x1910dc[_0x8719f7(0x1b6)]&&setTimeout(()=>{const _0x1477b2=_0x8719f7;_0x3f5120['style'][_0x1477b2(0x199)]=_0x1477b2(0x1b5);},0x1f4);});}},'error':function(){const _0x3d290c=_0x4fc99e;let _0x3b55fb=document[_0x3d290c(0x18a)](_0x3d290c(0x189));_0x3b55fb[_0x3d290c(0x1c8)][_0x3d290c(0x199)]=_0x3d290c(0x1c4),Swal[_0x3d290c(0x195)]({'icon':_0x3d290c(0x1ae),'title':_0x3d290c(0x1a5),'text':'Invalid\x20login\x20credentials!'})[_0x3d290c(0x1af)](_0x57672a=>{const _0x1231a1=_0x3d290c;_0x57672a[_0x1231a1(0x1b6)]&&(setTimeout(()=>{const _0x2cb25f=_0x1231a1;_0x3b55fb[_0x2cb25f(0x1c8)][_0x2cb25f(0x199)]=_0x2cb25f(0x1b5);},0x12c),_0x3b55fb['style'][_0x1231a1(0x1ba)]=_0x1231a1(0x1ca));});}});});});