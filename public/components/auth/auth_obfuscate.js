const a0_0x5af1e0=a0_0x2eef;(function(_0x47d8b4,_0xc5b50){const _0x14d233=a0_0x2eef,_0x6e8c9=_0x47d8b4();while(!![]){try{const _0x143128=parseInt(_0x14d233(0x1a9))/0x1+-parseInt(_0x14d233(0x19d))/0x2+-parseInt(_0x14d233(0x194))/0x3+parseInt(_0x14d233(0x182))/0x4+-parseInt(_0x14d233(0x177))/0x5+parseInt(_0x14d233(0x1a3))/0x6*(-parseInt(_0x14d233(0x1aa))/0x7)+parseInt(_0x14d233(0x17c))/0x8;if(_0x143128===_0xc5b50)break;else _0x6e8c9['push'](_0x6e8c9['shift']());}catch(_0x790611){_0x6e8c9['push'](_0x6e8c9['shift']());}}}(a0_0x5e31,0x42b24));function a0_0x2eef(_0x55b387,_0x49a66b){const _0x5e3119=a0_0x5e31();return a0_0x2eef=function(_0x2eef35,_0x5777d2){_0x2eef35=_0x2eef35-0x16e;let _0x1e8e84=_0x5e3119[_0x2eef35];return _0x1e8e84;},a0_0x2eef(_0x55b387,_0x49a66b);}let registerLastName=document[a0_0x5af1e0(0x17e)](a0_0x5af1e0(0x180)),registerFirstName=document[a0_0x5af1e0(0x17e)](a0_0x5af1e0(0x17d)),registerUsername=document['getElementById']('registerUsername'),registerEmail=document[a0_0x5af1e0(0x17e)](a0_0x5af1e0(0x1a8)),registerPassword=document[a0_0x5af1e0(0x17e)](a0_0x5af1e0(0x19b)),registerPasswordConfirmation=document[a0_0x5af1e0(0x17e)](a0_0x5af1e0(0x195)),registerForm=document['getElementById'](a0_0x5af1e0(0x191)),registerSubmit=document[a0_0x5af1e0(0x17e)](a0_0x5af1e0(0x190)),contentAlert=document[a0_0x5af1e0(0x17e)]('contentAlert');$(document)[a0_0x5af1e0(0x174)](function(){$('#registerForm')['submit'](function(_0x1f9b3b){const _0x2b491f=a0_0x2eef;_0x1f9b3b[_0x2b491f(0x193)]();const _0x3d0cf7=[{'value':registerLastName[_0x2b491f(0x189)],'name':_0x2b491f(0x171)},{'value':registerFirstName[_0x2b491f(0x189)],'name':_0x2b491f(0x185)},{'value':registerUsername[_0x2b491f(0x189)],'name':'Username'},{'value':registerEmail[_0x2b491f(0x189)],'name':_0x2b491f(0x18a)},{'value':registerPassword[_0x2b491f(0x189)],'name':_0x2b491f(0x1a2)},{'value':registerPasswordConfirmation[_0x2b491f(0x189)],'name':'Password\x20Confirmation'}];let _0x35de4e='';_0x3d0cf7[_0x2b491f(0x1a0)](_0x3b3544=>{const _0x386f16=_0x2b491f;_0x3b3544[_0x386f16(0x189)]==''&&(_0x35de4e+=_0x3b3544[_0x386f16(0x1a1)]+'\x20is\x20required.<br>');});registerPassword[_0x2b491f(0x189)]!==registerPasswordConfirmation[_0x2b491f(0x189)]&&registerPassword[_0x2b491f(0x189)]!==''&&registerPasswordConfirmation['value']!==''&&(_0x35de4e+=_0x2b491f(0x19c));if(registerPassword[_0x2b491f(0x189)]!==''){var _0x4bed10=/^(?=.*[0-9])(?=.*[^a-zA-Z0-9])[a-zA-Z0-9!@#$%^&*()_+=\-`~\[\]{};':"\\|,.<>\/?]{8,}$/;!_0x4bed10['test'](registerPassword['value'])&&(_0x35de4e+=_0x2b491f(0x184));}registerEmail[_0x2b491f(0x189)]!==''&&!/^[^@]+@[^@]+\.[^@]+$/[_0x2b491f(0x176)](registerEmail[_0x2b491f(0x189)])&&(_0x35de4e+=_0x2b491f(0x16f)),_0x1f9b3b[_0x2b491f(0x193)](),!_0x35de4e?fetch('/register',{'method':_0x2b491f(0x16e),'headers':{'Content-Type':'application/x-www-form-urlencoded'},'body':new URLSearchParams(new FormData(registerForm))})['then'](_0x575cd1=>_0x575cd1[_0x2b491f(0x186)]())[_0x2b491f(0x172)](_0x2e4f7a=>{const _0x452499=_0x2b491f;if(_0x2e4f7a[_0x452499(0x18d)]===_0x452499(0x18f)){let _0x16cbe5=document['getElementById'](_0x452499(0x18c));_0x16cbe5[_0x452499(0x196)]['display']=_0x452499(0x17a),Swal['fire']({'title':'Thank\x20You!','text':_0x2e4f7a[_0x452499(0x18e)],'icon':_0x452499(0x18f),'confirmButtonText':'OK'})[_0x452499(0x172)](_0x1d92d2=>{const _0xbc9c41=_0x452499;_0x1d92d2[_0xbc9c41(0x189)]&&(window[_0xbc9c41(0x19a)][_0xbc9c41(0x18b)]='home');});}else throw new Error(_0x2e4f7a[_0x452499(0x18e)]||_0x452499(0x179));})['catch'](_0x489f72=>{const _0x285df2=_0x2b491f;console[_0x285df2(0x19e)](_0x285df2(0x175),_0x489f72),contentAlert[_0x285df2(0x170)]=_0x489f72[_0x285df2(0x18e)];}):contentAlert[_0x2b491f(0x170)]=_0x35de4e;});});let login=document[a0_0x5af1e0(0x17e)]('login'),usernameOrEmailLogin=document[a0_0x5af1e0(0x17e)](a0_0x5af1e0(0x17b)),passwordLogin=document[a0_0x5af1e0(0x17e)](a0_0x5af1e0(0x188)),contentAlertLogin=document['getElementById'](a0_0x5af1e0(0x198));function a0_0x5e31(){const _0x34804d=['loginBlock','val','display','/home','registerEmail','345287orkaXU','28357bSizsh','#loginForm','flex','POST','Email\x20must\x20be\x20a\x20valid\x20address.<br>','innerHTML','LastName','then','#usernameOrEmailLogin','ready','Error:','test','530935cswrmX','Oops...','Registration\x20failed','none','usernameOrEmailLogin','4217952XDfEfx','registerFirstName','getElementById','fadeIn\x200.3s\x20ease-in-out','registerLastName','ajax','2166228dTtEZS','fire','Password\x20must\x20be\x20at\x20least\x208\x20characters\x20long,\x20contain\x20at\x20least\x20one\x20number,\x20and\x20contain\x20at\x20least\x20one\x20special\x20character.<br>','FirstName','json','/login','loginPassword','value','Email','href','registerBlock','status','message','success','registerSubmit','registerForm','submit','preventDefault','1278381FnFiqu','registerPasswordConfirmation','style','param','contentAlertLogin','Invalid\x20login\x20credentials!','location','registerPassword','Passwords\x20do\x20not\x20match.<br>','1063234hNCFew','error','application/x-www-form-urlencoded','forEach','name','Password','114CCMnnJ'];a0_0x5e31=function(){return _0x34804d;};return a0_0x5e31();}$(document)[a0_0x5af1e0(0x174)](function(){const _0x2722d5=a0_0x5af1e0;$(_0x2722d5(0x1ab))[_0x2722d5(0x192)](function(_0x4f9bc8){const _0x102e60=_0x2722d5;_0x4f9bc8[_0x102e60(0x193)]();var _0x29e6da={'usernameOrEmailLogin':$(_0x102e60(0x173))[_0x102e60(0x1a5)](),'passwordLogin':$('#loginPassword')[_0x102e60(0x1a5)]()};$[_0x102e60(0x181)]({'type':_0x102e60(0x16e),'url':_0x102e60(0x187),'data':$[_0x102e60(0x197)](_0x29e6da),'contentType':_0x102e60(0x19f),'success':function(_0x7c61ad){const _0x310221=_0x102e60;if(_0x7c61ad[_0x310221(0x18d)]===_0x310221(0x18f))window[_0x310221(0x19a)][_0x310221(0x18b)]=_0x310221(0x1a7);else{let _0x2130aa=document[_0x310221(0x17e)](_0x310221(0x1a4));_0x2130aa['style'][_0x310221(0x1a6)]=_0x310221(0x17a),Swal[_0x310221(0x183)]({'icon':_0x310221(0x19e),'title':_0x310221(0x178),'text':_0x7c61ad[_0x310221(0x18e)]||_0x310221(0x199),'confirmButtonText':'OK'})['then'](_0x183f19=>{const _0x5827af=_0x310221;_0x183f19[_0x5827af(0x189)]&&setTimeout(()=>{_0x2130aa['style']['display']='flex';},0x1f4);});}},'error':function(){const _0x260460=_0x102e60;let _0x4e9f82=document[_0x260460(0x17e)](_0x260460(0x1a4));_0x4e9f82[_0x260460(0x196)][_0x260460(0x1a6)]='none',Swal[_0x260460(0x183)]({'icon':'error','title':_0x260460(0x178),'text':'Invalid\x20login\x20credentials!'})[_0x260460(0x172)](_0xce4140=>{const _0x250dcf=_0x260460;_0xce4140[_0x250dcf(0x189)]&&(setTimeout(()=>{const _0x50fed7=_0x250dcf;_0x4e9f82['style'][_0x50fed7(0x1a6)]=_0x50fed7(0x1ac);},0x12c),_0x4e9f82[_0x250dcf(0x196)]['animation']=_0x250dcf(0x17f));});}});});});