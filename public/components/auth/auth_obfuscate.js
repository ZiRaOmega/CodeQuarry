const a0_0x2882dd=a0_0x46fa;(function(_0x1e7da8,_0x295f77){const _0x82fd8d=a0_0x46fa,_0x580a6b=_0x1e7da8();while(!![]){try{const _0x4e9278=parseInt(_0x82fd8d(0x7b))/0x1*(parseInt(_0x82fd8d(0x89))/0x2)+parseInt(_0x82fd8d(0x82))/0x3+parseInt(_0x82fd8d(0xa2))/0x4*(parseInt(_0x82fd8d(0xb9))/0x5)+-parseInt(_0x82fd8d(0xc0))/0x6*(parseInt(_0x82fd8d(0x88))/0x7)+-parseInt(_0x82fd8d(0xa3))/0x8*(-parseInt(_0x82fd8d(0xb5))/0x9)+-parseInt(_0x82fd8d(0x9d))/0xa*(-parseInt(_0x82fd8d(0x81))/0xb)+-parseInt(_0x82fd8d(0xad))/0xc;if(_0x4e9278===_0x295f77)break;else _0x580a6b['push'](_0x580a6b['shift']());}catch(_0x5ba088){_0x580a6b['push'](_0x580a6b['shift']());}}}(a0_0x69a8,0x35645));let registerLastName=document[a0_0x2882dd(0x8f)](a0_0x2882dd(0xa0)),registerFirstName=document[a0_0x2882dd(0x8f)](a0_0x2882dd(0x8b)),registerUsername=document['getElementById'](a0_0x2882dd(0xa4)),registerEmail=document[a0_0x2882dd(0x8f)](a0_0x2882dd(0xa9)),registerPassword=document['getElementById'](a0_0x2882dd(0xb7)),registerPasswordConfirmation=document[a0_0x2882dd(0x8f)](a0_0x2882dd(0xb8)),registerForm=document[a0_0x2882dd(0x8f)](a0_0x2882dd(0x97)),registerSubmit=document[a0_0x2882dd(0x8f)]('registerSubmit'),contentAlert=document[a0_0x2882dd(0x8f)](a0_0x2882dd(0xa6));$(document)[a0_0x2882dd(0x8e)](function(){$('#registerForm')['submit'](function(_0x5cba90){const _0x4f77c8=a0_0x46fa;_0x5cba90[_0x4f77c8(0x94)]();const _0x502a3c=[{'value':registerLastName[_0x4f77c8(0x9a)],'name':_0x4f77c8(0x7d)},{'value':registerFirstName[_0x4f77c8(0x9a)],'name':_0x4f77c8(0xbd)},{'value':registerUsername[_0x4f77c8(0x9a)],'name':_0x4f77c8(0xb3)},{'value':registerEmail[_0x4f77c8(0x9a)],'name':'Email'},{'value':registerPassword[_0x4f77c8(0x9a)],'name':_0x4f77c8(0xbc)},{'value':registerPasswordConfirmation[_0x4f77c8(0x9a)],'name':'Password\x20Confirmation'}];let _0x22e6ee='';_0x502a3c[_0x4f77c8(0xac)](_0x144c82=>{const _0x25b6ff=_0x4f77c8;_0x144c82[_0x25b6ff(0x9a)]==''&&(_0x22e6ee+=_0x144c82[_0x25b6ff(0x84)]+_0x25b6ff(0x9f));});registerPassword[_0x4f77c8(0x9a)]!==registerPasswordConfirmation['value']&&registerPassword['value']!==''&&registerPasswordConfirmation[_0x4f77c8(0x9a)]!==''&&(_0x22e6ee+=_0x4f77c8(0xbe));if(registerPassword[_0x4f77c8(0x9a)]!==''){var _0x3fab48=/^(?=.*[0-9])(?=.*[^a-zA-Z0-9])[a-zA-Z0-9!@#$%^&*()_+=\-`~\[\]{};':"\\|,.<>\/?]{8,}$/;!_0x3fab48['test'](registerPassword['value'])&&(_0x22e6ee+=_0x4f77c8(0x96));}registerEmail[_0x4f77c8(0x9a)]!==''&&!/^[^@]+@[^@]+\.[^@]+$/[_0x4f77c8(0x87)](registerEmail[_0x4f77c8(0x9a)])&&(_0x22e6ee+=_0x4f77c8(0x93)),_0x5cba90[_0x4f77c8(0x94)](),!_0x22e6ee?fetch(_0x4f77c8(0x85),{'method':_0x4f77c8(0xb2),'headers':{'Content-Type':'application/x-www-form-urlencoded'},'body':new URLSearchParams(new FormData(registerForm))})['then'](_0x264929=>_0x264929['json']())[_0x4f77c8(0xba)](_0x1e4df5=>{const _0x55472a=_0x4f77c8;if(_0x1e4df5[_0x55472a(0xbb)]==='success'){let _0xd551d6=document['getElementById'](_0x55472a(0x7c));_0xd551d6[_0x55472a(0x92)][_0x55472a(0x95)]=_0x55472a(0xb0),Swal[_0x55472a(0x7e)]({'title':_0x55472a(0x8d),'text':_0x1e4df5[_0x55472a(0x91)],'icon':_0x55472a(0x9c),'confirmButtonText':'OK'})[_0x55472a(0xba)](_0x522ebc=>{const _0x5c0a50=_0x55472a;_0x522ebc[_0x5c0a50(0x9a)]&&(window[_0x5c0a50(0x98)][_0x5c0a50(0xb1)]=_0x5c0a50(0x9e));});}else throw new Error(_0x1e4df5[_0x55472a(0x91)]||_0x55472a(0x8a));})[_0x4f77c8(0x80)](_0x2fd03f=>{const _0x5c16af=_0x4f77c8;console[_0x5c16af(0xab)](_0x5c16af(0x99),_0x2fd03f),contentAlert[_0x5c16af(0x90)]=_0x2fd03f[_0x5c16af(0x91)];}):contentAlert['innerHTML']=_0x22e6ee;});});let login=document[a0_0x2882dd(0x8f)](a0_0x2882dd(0xb4)),usernameOrEmailLogin=document['getElementById'](a0_0x2882dd(0xbf)),passwordLogin=document[a0_0x2882dd(0x8f)](a0_0x2882dd(0xa7)),contentAlertLogin=document['getElementById'](a0_0x2882dd(0xae));function a0_0x46fa(_0x153ae2,_0x1470e7){const _0x69a8c4=a0_0x69a8();return a0_0x46fa=function(_0x46faed,_0x15e6b2){_0x46faed=_0x46faed-0x7b;let _0x9a19cb=_0x69a8c4[_0x46faed];return _0x9a19cb;},a0_0x46fa(_0x153ae2,_0x1470e7);}$(document)['ready'](function(){const _0x557b94=a0_0x2882dd;$('#loginForm')[_0x557b94(0x83)](function(_0x1061c1){const _0x48e436=_0x557b94;_0x1061c1[_0x48e436(0x94)]();var _0x228234={'usernameOrEmailLogin':$(_0x48e436(0xaa))['val'](),'passwordLogin':$('#loginPassword')[_0x48e436(0xa5)]()};$[_0x48e436(0xaf)]({'type':_0x48e436(0xb2),'url':_0x48e436(0xb6),'data':$[_0x48e436(0xa8)](_0x228234),'contentType':_0x48e436(0x86),'success':function(_0x35985a){const _0x3b345b=_0x48e436;if(_0x35985a[_0x3b345b(0xbb)]===_0x3b345b(0x9c))window[_0x3b345b(0x98)]['href']='/home';else{let _0x4d35f8=document['getElementById'](_0x3b345b(0x7f));_0x4d35f8[_0x3b345b(0x92)][_0x3b345b(0x95)]=_0x3b345b(0xb0),Swal[_0x3b345b(0x7e)]({'icon':'error','title':_0x3b345b(0xa1),'text':_0x35985a['message']||_0x3b345b(0x9b),'confirmButtonText':'OK'})['then'](_0x43f018=>{const _0x3f8950=_0x3b345b;_0x43f018[_0x3f8950(0x9a)]&&setTimeout(()=>{const _0x1e77eb=_0x3f8950;_0x4d35f8[_0x1e77eb(0x92)][_0x1e77eb(0x95)]=_0x1e77eb(0x8c);},0x1f4);});}},'error':function(){const _0x152560=_0x48e436;let _0x3aa4cf=document[_0x152560(0x8f)](_0x152560(0x7f));_0x3aa4cf[_0x152560(0x92)][_0x152560(0x95)]='none',Swal['fire']({'icon':_0x152560(0xab),'title':_0x152560(0xa1),'text':'Invalid\x20login\x20credentials!'})['then'](_0x1a7272=>{const _0x1a71e3=_0x152560;_0x1a7272[_0x1a71e3(0x9a)]&&(setTimeout(()=>{const _0x3f54b8=_0x1a71e3;_0x3aa4cf[_0x3f54b8(0x92)]['display']=_0x3f54b8(0x8c);},0x12c),_0x3aa4cf[_0x1a71e3(0x92)]['animation']='fadeIn\x200.3s\x20ease-in-out');});}});});});function a0_0x69a8(){const _0x2a3d4a=['32195JbDpQP','then','status','Password','FirstName','Passwords\x20do\x20not\x20match.<br>','usernameOrEmailLogin','78KPOUug','34CtKZcq','registerBlock','LastName','fire','loginBlock','catch','35101IyBQyv','530283FTcEJK','submit','name','/register','application/x-www-form-urlencoded','test','65695kLISlt','7946weSCSo','Registration\x20failed','registerFirstName','flex','Thank\x20You!','ready','getElementById','innerHTML','message','style','Email\x20must\x20be\x20a\x20valid\x20address.<br>','preventDefault','display','Password\x20must\x20be\x20at\x20least\x208\x20characters\x20long,\x20contain\x20at\x20least\x20one\x20number,\x20and\x20contain\x20at\x20least\x20one\x20special\x20character.<br>','registerForm','location','Error:','value','Invalid\x20login\x20credentials!','success','940iaLtrj','home','\x20is\x20required.<br>','registerLastName','Oops...','12MmlhtA','16KgiAtS','registerUsername','val','contentAlert','loginPassword','param','registerEmail','#usernameOrEmailLogin','error','forEach','5624904rTNHVv','contentAlertLogin','ajax','none','href','POST','Username','login','802467aSqlPx','/login','registerPassword','registerPasswordConfirmation'];a0_0x69a8=function(){return _0x2a3d4a;};return a0_0x69a8();}