const a0_0x20e63a=a0_0x24c5;(function(_0x400901,_0x3f2513){const _0x5c208b=a0_0x24c5,_0x364daf=_0x400901();while(!![]){try{const _0xe8c186=-parseInt(_0x5c208b(0xca))/0x1+parseInt(_0x5c208b(0xa5))/0x2+-parseInt(_0x5c208b(0x9c))/0x3+parseInt(_0x5c208b(0xb4))/0x4*(parseInt(_0x5c208b(0xa6))/0x5)+parseInt(_0x5c208b(0xa2))/0x6*(parseInt(_0x5c208b(0x9a))/0x7)+parseInt(_0x5c208b(0xb5))/0x8+parseInt(_0x5c208b(0xc4))/0x9*(-parseInt(_0x5c208b(0xa4))/0xa);if(_0xe8c186===_0x3f2513)break;else _0x364daf['push'](_0x364daf['shift']());}catch(_0x564901){_0x364daf['push'](_0x364daf['shift']());}}}(a0_0x3f5f,0xe4e12));function a0_0x24c5(_0x2f57d7,_0xe3b89){const _0x3f5f5d=a0_0x3f5f();return a0_0x24c5=function(_0x24c599,_0x556d6b){_0x24c599=_0x24c599-0x8f;let _0x1e8d77=_0x3f5f5d[_0x24c599];return _0x1e8d77;},a0_0x24c5(_0x2f57d7,_0xe3b89);}let registerLastName=document[a0_0x20e63a(0x94)](a0_0x20e63a(0xb1)),registerFirstName=document[a0_0x20e63a(0x94)](a0_0x20e63a(0xac)),registerUsername=document[a0_0x20e63a(0x94)](a0_0x20e63a(0xa3)),registerEmail=document[a0_0x20e63a(0x94)]('registerEmail'),registerPassword=document[a0_0x20e63a(0x94)](a0_0x20e63a(0xb2)),registerPasswordConfirmation=document[a0_0x20e63a(0x94)](a0_0x20e63a(0xa1)),registerForm=document[a0_0x20e63a(0x94)](a0_0x20e63a(0xcd)),registerSubmit=document[a0_0x20e63a(0x94)](a0_0x20e63a(0x92)),contentAlert=document[a0_0x20e63a(0x94)](a0_0x20e63a(0x91));$(document)[a0_0x20e63a(0xcc)](function(){const _0x4fb5a5=a0_0x20e63a;$(_0x4fb5a5(0xc1))[_0x4fb5a5(0xbe)](function(_0x2dcbab){const _0x23996f=_0x4fb5a5;_0x2dcbab['preventDefault']();const _0x49e37a=[{'value':registerLastName[_0x23996f(0x95)],'name':_0x23996f(0xbb)},{'value':registerFirstName['value'],'name':_0x23996f(0xa7)},{'value':registerUsername['value'],'name':_0x23996f(0x99)},{'value':registerEmail[_0x23996f(0x95)],'name':_0x23996f(0x8f)},{'value':registerPassword[_0x23996f(0x95)],'name':'Password'},{'value':registerPasswordConfirmation[_0x23996f(0x95)],'name':_0x23996f(0xa0)}];let _0x14c959='';_0x49e37a[_0x23996f(0x9f)](_0x43d9f0=>{const _0x34e566=_0x23996f;_0x43d9f0[_0x34e566(0x95)]==''&&(_0x14c959+=_0x43d9f0[_0x34e566(0xc8)]+_0x34e566(0xbd));});registerPassword[_0x23996f(0x95)]!==registerPasswordConfirmation[_0x23996f(0x95)]&&registerPassword[_0x23996f(0x95)]!==''&&registerPasswordConfirmation[_0x23996f(0x95)]!==''&&(_0x14c959+=_0x23996f(0x9e));if(registerPassword[_0x23996f(0x95)]!==''){var _0x43d48e=/^(?=.*[0-9])(?=.*[^a-zA-Z0-9])[a-zA-Z0-9!@#$%^&*()_+=\-`~\[\]{};':"\\|,.<>\/?]{8,}$/;!_0x43d48e['test'](registerPassword['value'])&&(_0x14c959+=_0x23996f(0x98));}registerEmail['value']!==''&&!/^[^@]+@[^@]+\.[^@]+$/[_0x23996f(0xa9)](registerEmail['value'])&&(_0x14c959+='Email\x20must\x20be\x20a\x20valid\x20address.<br>'),_0x2dcbab[_0x23996f(0xbc)](),!_0x14c959?fetch(_0x23996f(0xaa),{'method':_0x23996f(0x96),'headers':{'Content-Type':_0x23996f(0xc5)},'body':new URLSearchParams(new FormData(registerForm))})[_0x23996f(0xba)](_0x2e1ec8=>_0x2e1ec8['json']())[_0x23996f(0xba)](_0x4cc500=>{const _0x495f4=_0x23996f;if(_0x4cc500[_0x495f4(0xc0)]===_0x495f4(0xcb)){let _0x247738=document[_0x495f4(0x94)]('registerBlock');_0x247738[_0x495f4(0x9b)]['display']=_0x495f4(0xc9),Swal[_0x495f4(0xb6)]({'title':_0x495f4(0xc3),'text':_0x4cc500['message'],'icon':_0x495f4(0xcb),'confirmButtonText':'OK'})['then'](_0x4c81d4=>{const _0x5141cb=_0x495f4;_0x4c81d4[_0x5141cb(0x95)]&&(window['location']['href']='/codeQuarry');});}else throw new Error(_0x4cc500['message']||'Registration\x20failed');})['catch'](_0x8eb17f=>{const _0x209eb4=_0x23996f;console[_0x209eb4(0xcf)](_0x209eb4(0x9d),_0x8eb17f),contentAlert['innerHTML']=_0x8eb17f['message'];}):contentAlert[_0x23996f(0xaf)]=_0x14c959;});});let login=document[a0_0x20e63a(0x94)](a0_0x20e63a(0xa8)),usernameOrEmailLogin=document[a0_0x20e63a(0x94)](a0_0x20e63a(0xb3)),passwordLogin=document[a0_0x20e63a(0x94)]('loginPassword'),contentAlertLogin=document['getElementById'](a0_0x20e63a(0xb0));function a0_0x3f5f(){const _0x5838c9=['7mkVpHn','style','895572eGRLYv','Error:','Passwords\x20do\x20not\x20match.<br>','forEach','Password\x20Confirmation','registerPasswordConfirmation','1021950DxZzyN','registerUsername','730pYPMRO','1113402zVCXRf','406070gVbvsl','FirstName','login','test','/register','flex','registerFirstName','val','location','innerHTML','contentAlertLogin','registerLastName','registerPassword','usernameOrEmailLogin','80COtNlD','1486128oynIka','fire','href','/login','animation','then','LastName','preventDefault','\x20is\x20required.<br>','submit','#usernameOrEmailLogin','status','#registerForm','display','Thank\x20You!','138069lFVAuZ','application/x-www-form-urlencoded','param','Invalid\x20login\x20credentials!','name','none','181165KbIHAD','success','ready','registerForm','fadeIn\x200.3s\x20ease-in-out','error','Email','#loginForm','contentAlert','registerSubmit','message','getElementById','value','POST','#loginPassword','Password\x20must\x20be\x20at\x20least\x208\x20characters\x20long,\x20contain\x20at\x20least\x20one\x20number,\x20and\x20contain\x20at\x20least\x20one\x20special\x20character.<br>','Username'];a0_0x3f5f=function(){return _0x5838c9;};return a0_0x3f5f();}$(document)[a0_0x20e63a(0xcc)](function(){const _0x5f4f75=a0_0x20e63a;$(_0x5f4f75(0x90))['submit'](function(_0x33f948){const _0x542b82=_0x5f4f75;_0x33f948['preventDefault']();var _0x47401e={'usernameOrEmailLogin':$(_0x542b82(0xbf))[_0x542b82(0xad)](),'passwordLogin':$(_0x542b82(0x97))['val']()};$['ajax']({'type':_0x542b82(0x96),'url':_0x542b82(0xb8),'data':$[_0x542b82(0xc6)](_0x47401e),'contentType':_0x542b82(0xc5),'success':function(_0xc13d64){const _0x42bbc1=_0x542b82;if(_0xc13d64[_0x42bbc1(0xc0)]==='success')window[_0x42bbc1(0xae)][_0x42bbc1(0xb7)]='/codeQuarry';else{let _0x1e3e65=document[_0x42bbc1(0x94)]('loginBlock');_0x1e3e65[_0x42bbc1(0x9b)]['display']=_0x42bbc1(0xc9),Swal[_0x42bbc1(0xb6)]({'icon':'error','title':'Oops...','text':_0xc13d64[_0x42bbc1(0x93)]||_0x42bbc1(0xc7),'confirmButtonText':'OK'})[_0x42bbc1(0xba)](_0x2b7bc5=>{const _0x50b4c8=_0x42bbc1;_0x2b7bc5[_0x50b4c8(0x95)]&&setTimeout(()=>{const _0x159dc4=_0x50b4c8;_0x1e3e65['style'][_0x159dc4(0xc2)]=_0x159dc4(0xab);},0x1f4);});}},'error':function(){const _0x4b34d2=_0x542b82;let _0x4cf948=document[_0x4b34d2(0x94)]('loginBlock');_0x4cf948[_0x4b34d2(0x9b)][_0x4b34d2(0xc2)]=_0x4b34d2(0xc9),Swal[_0x4b34d2(0xb6)]({'icon':_0x4b34d2(0xcf),'title':'Oops...','text':_0x4b34d2(0xc7)})['then'](_0x17ec9b=>{const _0xd6d413=_0x4b34d2;_0x17ec9b[_0xd6d413(0x95)]&&(setTimeout(()=>{const _0x4130a0=_0xd6d413;_0x4cf948[_0x4130a0(0x9b)][_0x4130a0(0xc2)]=_0x4130a0(0xab);},0x12c),_0x4cf948[_0xd6d413(0x9b)][_0xd6d413(0xb9)]=_0xd6d413(0xce));});}});});});