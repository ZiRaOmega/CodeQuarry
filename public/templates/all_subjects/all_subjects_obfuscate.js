const a0_0x152832=a0_0x3a57;(function(_0x5ea2e1,_0x5b2e7a){const _0x4e6abc=a0_0x3a57,_0x4e2248=_0x5ea2e1();while(!![]){try{const _0x116a15=-parseInt(_0x4e6abc(0x19a))/0x1*(parseInt(_0x4e6abc(0x14a))/0x2)+-parseInt(_0x4e6abc(0x164))/0x3+parseInt(_0x4e6abc(0x196))/0x4*(-parseInt(_0x4e6abc(0x169))/0x5)+parseInt(_0x4e6abc(0x16c))/0x6+parseInt(_0x4e6abc(0x133))/0x7+-parseInt(_0x4e6abc(0x14f))/0x8*(parseInt(_0x4e6abc(0x188))/0x9)+-parseInt(_0x4e6abc(0x170))/0xa*(-parseInt(_0x4e6abc(0x13d))/0xb);if(_0x116a15===_0x5b2e7a)break;else _0x4e2248['push'](_0x4e2248['shift']());}catch(_0x449306){_0x4e2248['push'](_0x4e2248['shift']());}}}(a0_0x361d,0xd4e60));function a0_0x3a57(_0x10ffe1,_0x46c063){const _0x361dfe=a0_0x361d();return a0_0x3a57=function(_0x3a57aa,_0x22ad80){_0x3a57aa=_0x3a57aa-0x12a;let _0x5af10c=_0x361dfe[_0x3a57aa];return _0x5af10c;},a0_0x3a57(_0x10ffe1,_0x46c063);}let SubjectsList=[],ListElement;function initializeLocalStorage(){const _0x46b6fe=a0_0x3a57;localStorage[_0x46b6fe(0x192)](_0x46b6fe(0x141)),localStorage[_0x46b6fe(0x192)]('subjectTitle');}document[a0_0x152832(0x191)]('DOMContentLoaded',function(){const _0x412b0b=a0_0x152832;initializeLocalStorage(),ListElement=document[_0x412b0b(0x131)](_0x412b0b(0x145))||document[_0x412b0b(0x190)](_0x412b0b(0x166)),((async()=>{const _0x315735=_0x412b0b;try{const _0xb70750=await fetch(_0x315735(0x143)),_0x584544=await _0xb70750[_0x315735(0x198)](),_0x2bf906=createAllSubjectsItem();ListElement['appendChild'](_0x2bf906),addAllSubjectsClickListener(_0x2bf906,ListElement),createSubjectItems(_0x2bf906,ListElement,_0x584544);}catch(_0x46aa93){const _0x5a4b87=document[_0x315735(0x190)]('h1');_0x5a4b87[_0x315735(0x181)]='An\x20error\x20occured\x20while\x20fetching\x20the\x20subjects',_0x5a4b87['style'][_0x315735(0x165)]=_0x315735(0x179),ListElement['appendChild'](_0x5a4b87),console['error'](_0x315735(0x149),_0x46aa93);}})());});function a0_0x361d(){const _0x233527=['length','question_filter','send','subjectId','all','/api/subjects','click','all_subjects_list','toLowerCase','upvotes','Upvotes\x20↗','There\x20was\x20a\x20problem\x20with\x20your\x20fetch\x20operation:','107102WRlsEr','Posted\x20the:\x20','title','setItem','filter_container','13482744jFJsos','filter_oldest','question_tracker_count','querySelectorAll','href','question_checked','All','Click\x20here\x20to\x20view\x20all\x20questions\x20across\x20all\x20subjects.','location','creator','session','some','question_creator','add','style','span','favori','question','data-subject-id','upvoted','display','4219983skfhNf','color','div','pre\x20code','toLocaleDateString','190715kUQsAE','responses_counter','/api/questions?subjectId=','5698260qEKTGe','/home','backgroundColor','/question_viewer?question_id=','40ErJpVw','green','downvote','/subject/','Upvotes\x20↘','innerHTML','questionCount','https://','none','red','querySelector','fetchQuestions','filter_questions','question_count','category_cards','downvoted','isArray','textContent','/subject/all','appendChild','filter_popular','Oldest','downvotes','favori_active','9wrcBez','downvote_text','push','rgb(196,\x2077,\x2086)','content','vote_container','Answered\x20✔','question_content','createElement','addEventListener','removeItem','user_vote','.question_tracker_count','rgb(104,\x20195,\x20163)','16DqQgJJ','question_description','json','forEach','5EYIUcr','stringify','0\x20question(s)','Newest','/api/favoris','then','setAttribute','subject_tag','classList','onclick','upvote','subjectTitle','responses','downvote_count','description','filter_best_answer','category_title','getElementById','↑\x20Comments','1742895VcRVrS','question_creation_date','response_container','log','remove','upvote_container','contains','hostname','category_description','data-question-id','8760587dzHGRU'];a0_0x361d=function(){return _0x233527;};return a0_0x361d();}const returnButton=document[a0_0x152832(0x190)](a0_0x152832(0x166));returnButton[a0_0x152832(0x1a3)]=()=>{const _0x46ae04=a0_0x152832;window[_0x46ae04(0x157)][_0x46ae04(0x153)]=_0x46ae04(0x16d);};function createAllSubjectsItem(){const _0x8f15aa=a0_0x152832,_0x2623be=document[_0x8f15aa(0x190)](_0x8f15aa(0x166));_0x2623be[_0x8f15aa(0x1a2)][_0x8f15aa(0x15c)](_0x8f15aa(0x17e));const _0x14f663=document[_0x8f15aa(0x190)]('h2');_0x14f663['classList'][_0x8f15aa(0x15c)](_0x8f15aa(0x130)),_0x14f663[_0x8f15aa(0x181)]=_0x8f15aa(0x155),_0x2623be[_0x8f15aa(0x183)](_0x14f663);const _0x4a5c47=document['createElement']('p');return _0x4a5c47[_0x8f15aa(0x1a2)][_0x8f15aa(0x15c)](_0x8f15aa(0x13b)),_0x4a5c47[_0x8f15aa(0x181)]=_0x8f15aa(0x156),_0x2623be[_0x8f15aa(0x183)](_0x4a5c47),_0x2623be;}function addAllSubjectsClickListener(_0x2559fd,_0x1d852a){const _0x3f48a7=a0_0x152832;_0x2559fd[_0x3f48a7(0x191)]('click',function(){const _0x50cd89=_0x3f48a7;localStorage['setItem']('subjectId',_0x50cd89(0x142)),localStorage['setItem'](_0x50cd89(0x12b),'All\x20Subjects'),window['location']['href']=_0x50cd89(0x182);});}function createSubjectItems(_0x1e6b81,_0x2d1335,_0x3dc281){const _0x47e081=a0_0x152832,_0x2256cd=document[_0x47e081(0x190)]('div');_0x2256cd[_0x47e081(0x1a2)][_0x47e081(0x15c)]('question_count_all');let _0x9fe9aa=0x0;SubjectsList=[],_0x3dc281[_0x47e081(0x199)](_0x127f9b=>{const _0x570047=_0x47e081;SubjectsList[_0x570047(0x18a)](_0x127f9b),_0x9fe9aa+=_0x127f9b[_0x570047(0x176)],_0x2256cd[_0x570047(0x181)]=_0x9fe9aa,_0x1e6b81[_0x570047(0x183)](_0x2256cd);const _0x48b20a=createItem(_0x127f9b);addSubjectClickListener(_0x48b20a,_0x127f9b,_0x2d1335),_0x2d1335[_0x570047(0x183)](_0x48b20a);});}function createItem(_0x21e254){const _0x3856dd=a0_0x152832,_0x1639e4=document['createElement'](_0x3856dd(0x166));_0x1639e4['classList']['add'](_0x3856dd(0x17e));const _0x36edb7=document[_0x3856dd(0x190)](_0x3856dd(0x166));_0x36edb7[_0x3856dd(0x1a2)][_0x3856dd(0x15c)](_0x3856dd(0x17d)),_0x36edb7[_0x3856dd(0x1a0)](_0x3856dd(0x161),_0x21e254['id']),_0x36edb7[_0x3856dd(0x181)]=_0x21e254[_0x3856dd(0x176)],_0x1639e4[_0x3856dd(0x183)](_0x36edb7);const _0x561146=document[_0x3856dd(0x190)]('h2');_0x561146[_0x3856dd(0x1a2)][_0x3856dd(0x15c)](_0x3856dd(0x130)),_0x561146['textContent']=_0x21e254[_0x3856dd(0x14c)],_0x1639e4['appendChild'](_0x561146);const _0x155264=document[_0x3856dd(0x190)]('p');return _0x155264[_0x3856dd(0x1a2)][_0x3856dd(0x15c)](_0x3856dd(0x13b)),_0x155264['textContent']=_0x21e254[_0x3856dd(0x12e)],_0x1639e4[_0x3856dd(0x183)](_0x155264),_0x1639e4;}function addSubjectClickListener(_0x146fb1,_0x489cc6,_0x5dd53b){const _0x53fcec=a0_0x152832;_0x146fb1['addEventListener'](_0x53fcec(0x144),function(){const _0x14394f=_0x53fcec;localStorage['setItem']('subjectId',_0x489cc6['id']),localStorage[_0x14394f(0x14d)](_0x14394f(0x12b),_0x489cc6[_0x14394f(0x14c)]),_0x489cc6[_0x14394f(0x14c)]=_0x489cc6['title'][_0x14394f(0x146)](),window['location'][_0x14394f(0x153)]=_0x14394f(0x173)+_0x489cc6['id']+'/'+_0x489cc6[_0x14394f(0x14c)];});}function fetchQuestions(_0x157c1c){const _0x19f60c=a0_0x152832;console[_0x19f60c(0x136)](_0x19f60c(0x17b)),fetch(_0x19f60c(0x16b)+_0x157c1c)[_0x19f60c(0x19f)](_0x2fe2fb=>_0x2fe2fb[_0x19f60c(0x198)]())[_0x19f60c(0x19f)](_0x4a9315=>{createFilter(_0x4a9315),createQuestions(_0x4a9315);});};function createFilter(_0x4eb346){const _0x3ed187=a0_0x152832;questionsList[_0x3ed187(0x175)]='';const _0x58f816=document[_0x3ed187(0x190)]('div');_0x58f816[_0x3ed187(0x1a2)][_0x3ed187(0x15c)](_0x3ed187(0x14e));const _0x3516e1=createQuestionFilter(_0x4eb346);returnButton[_0x3ed187(0x181)]='⬅',returnButton['id']='return_button',_0x58f816[_0x3ed187(0x183)](returnButton),_0x58f816[_0x3ed187(0x183)](_0x3516e1),questionsList[_0x3ed187(0x183)](_0x58f816);}function createQuestionFilter(_0x1a511c){const _0x4ab78f=a0_0x152832,_0x437399=document[_0x4ab78f(0x190)](_0x4ab78f(0x166));_0x437399['classList'][_0x4ab78f(0x15c)](_0x4ab78f(0x13f));const _0x279b57=document[_0x4ab78f(0x190)](_0x4ab78f(0x166));_0x279b57[_0x4ab78f(0x1a2)]['add'](_0x4ab78f(0x151));const _0x3869b4=document[_0x4ab78f(0x190)](_0x4ab78f(0x166));_0x3869b4['classList'][_0x4ab78f(0x15c)](_0x4ab78f(0x17c));const _0x1e4217=createFilterElements();return _0x1e4217[_0x4ab78f(0x199)](_0x355aef=>_0x3869b4['appendChild'](_0x355aef)),_0x437399[_0x4ab78f(0x183)](_0x279b57),_0x437399[_0x4ab78f(0x183)](_0x3869b4),updateQuestionTrackerCount(_0x1a511c,_0x279b57),_0x1e4217[0x0][_0x4ab78f(0x1a3)]=()=>sortByNumberOfComments(_0x1a511c),_0x1e4217[0x1][_0x4ab78f(0x1a3)]=()=>sortOldestToNewest(_0x1a511c),_0x1e4217[0x2][_0x4ab78f(0x1a3)]=()=>sortByBestAnswer(_0x1a511c),_0x1e4217[0x3][_0x4ab78f(0x1a3)]=()=>sortNewestToOldest(_0x1a511c),_0x1e4217[0x4][_0x4ab78f(0x1a3)]=()=>sortByUpvotes(_0x1a511c),_0x1e4217[0x5][_0x4ab78f(0x1a3)]=()=>sortByDownvotes(_0x1a511c),console[_0x4ab78f(0x136)](_0x1e4217),_0x437399;}function createFilterElements(){const _0x37d403=a0_0x152832,_0x315f6b=document[_0x37d403(0x190)](_0x37d403(0x166));_0x315f6b['classList']['add'](_0x37d403(0x184)),_0x315f6b[_0x37d403(0x181)]=_0x37d403(0x148);const _0x545f14=document[_0x37d403(0x190)]('div');_0x545f14['classList'][_0x37d403(0x15c)]('filter_unpopular'),_0x545f14[_0x37d403(0x181)]=_0x37d403(0x174);const _0x5876a1=document[_0x37d403(0x190)]('div');_0x5876a1[_0x37d403(0x1a2)][_0x37d403(0x15c)]('filter_newest'),_0x5876a1[_0x37d403(0x181)]=_0x37d403(0x19d);const _0x54d726=document['createElement'](_0x37d403(0x166));_0x54d726[_0x37d403(0x1a2)][_0x37d403(0x15c)](_0x37d403(0x150)),_0x54d726[_0x37d403(0x181)]=_0x37d403(0x185);const _0x14368d=document[_0x37d403(0x190)](_0x37d403(0x166));_0x14368d['classList']['add'](_0x37d403(0x12f)),_0x14368d[_0x37d403(0x181)]=_0x37d403(0x18e);const _0x19b95d=document[_0x37d403(0x190)](_0x37d403(0x166));return _0x19b95d[_0x37d403(0x1a2)][_0x37d403(0x15c)]('filter_number_of_comments'),_0x19b95d[_0x37d403(0x181)]=_0x37d403(0x132),[_0x19b95d,_0x54d726,_0x14368d,_0x5876a1,_0x315f6b,_0x545f14];}function updateQuestionTrackerCount(_0x3bb62a,_0x46ae1d){const _0x29fa4d=a0_0x152832;_0x46ae1d[_0x29fa4d(0x181)]=_0x3bb62a?_0x3bb62a[_0x29fa4d(0x13e)]+'\x20question(s)':_0x29fa4d(0x19c);}function createQuestions(_0x182163){const _0xe06558=a0_0x152832;if(_0x182163!=null)_0x182163[_0xe06558(0x199)](_0x77d1fd=>{const _0x1ca67e=_0xe06558,_0x387784=document[_0x1ca67e(0x190)](_0x1ca67e(0x166));_0x387784[_0x1ca67e(0x1a2)]['add'](_0x1ca67e(0x160));const _0x2072c2=document[_0x1ca67e(0x190)](_0x1ca67e(0x166));_0x2072c2[_0x1ca67e(0x1a2)][_0x1ca67e(0x15c)](_0x1ca67e(0x154)),_0x2072c2['setAttribute'](_0x1ca67e(0x13c),_0x77d1fd['id']),_0x387784[_0x1ca67e(0x183)](_0x2072c2);_0x77d1fd[_0x1ca67e(0x12c)]!=null&&(_0x77d1fd[_0x1ca67e(0x12c)][_0x1ca67e(0x15a)](_0x9c15de=>_0x9c15de['best_answer']==!![])?_0x2072c2[_0x1ca67e(0x15d)][_0x1ca67e(0x163)]='block':_0x2072c2['style'][_0x1ca67e(0x163)]=_0x1ca67e(0x178));_0x387784[_0x1ca67e(0x1a0)]('data-question-id',_0x77d1fd['id']);const _0x1ced84=document[_0x1ca67e(0x190)](_0x1ca67e(0x166));_0x1ced84[_0x1ca67e(0x1a2)][_0x1ca67e(0x15c)]('clickable_container');const _0x5eeb89=document['createElement']('div');_0x5eeb89[_0x1ca67e(0x1a2)][_0x1ca67e(0x15c)](_0x1ca67e(0x1a1)),_0x5eeb89[_0x1ca67e(0x181)]=_0x77d1fd['subject_title'],_0x387784[_0x1ca67e(0x183)](_0x5eeb89);const _0x21a0af=document[_0x1ca67e(0x190)]('h3');_0x21a0af[_0x1ca67e(0x1a2)][_0x1ca67e(0x15c)]('question_title'),_0x21a0af[_0x1ca67e(0x181)]=_0x77d1fd[_0x1ca67e(0x14c)],_0x1ced84['appendChild'](_0x21a0af);const _0x4e257a=document[_0x1ca67e(0x190)]('p');_0x4e257a[_0x1ca67e(0x1a2)][_0x1ca67e(0x15c)](_0x1ca67e(0x197)),_0x4e257a[_0x1ca67e(0x181)]=_0x77d1fd[_0x1ca67e(0x12e)],_0x1ced84[_0x1ca67e(0x183)](_0x4e257a);const _0xc19b=document[_0x1ca67e(0x190)]('p');_0xc19b[_0x1ca67e(0x1a2)]['add'](_0x1ca67e(0x18f)),_0xc19b[_0x1ca67e(0x181)]=_0x77d1fd[_0x1ca67e(0x18c)];const _0x345eda=document['createElement']('pre'),_0x15b30c=document[_0x1ca67e(0x190)]('code');_0x345eda[_0x1ca67e(0x183)](_0x15b30c),_0x15b30c['textContent']=_0x77d1fd['content'],document[_0x1ca67e(0x152)](_0x1ca67e(0x167))[_0x1ca67e(0x199)](_0x11c159=>{hljs['highlightElement'](_0x11c159);}),_0x1ced84['appendChild'](_0x345eda);const _0x59b9c4=document[_0x1ca67e(0x190)]('div');_0x59b9c4[_0x1ca67e(0x1a2)][_0x1ca67e(0x15c)]('creator_and_date_container');const _0x4d5ac5=document[_0x1ca67e(0x190)]('p');_0x4d5ac5[_0x1ca67e(0x1a2)][_0x1ca67e(0x15c)](_0x1ca67e(0x134)),_0x4d5ac5[_0x1ca67e(0x181)]=_0x1ca67e(0x14b)+new Date(_0x77d1fd['creation_date'])[_0x1ca67e(0x168)](),_0x59b9c4[_0x1ca67e(0x183)](_0x4d5ac5);const _0x9862c9=document['createElement']('p');_0x9862c9[_0x1ca67e(0x1a2)]['add'](_0x1ca67e(0x15b)),_0x9862c9['textContent']='Created\x20by';const _0x4eea8e=document[_0x1ca67e(0x190)](_0x1ca67e(0x15e));_0x4eea8e['textContent']=_0x77d1fd[_0x1ca67e(0x158)],_0x4eea8e['classList']['add']('creator_name'),_0x9862c9[_0x1ca67e(0x183)](_0x4eea8e);const _0x4b3542=document[_0x1ca67e(0x190)]('p');_0x4b3542[_0x1ca67e(0x1a2)][_0x1ca67e(0x15c)](_0x1ca67e(0x16a));Array[_0x1ca67e(0x180)](_0x77d1fd[_0x1ca67e(0x12c)])?_0x4b3542[_0x1ca67e(0x181)]=_0x77d1fd[_0x1ca67e(0x12c)]['length']+'\x20response(s)':_0x4b3542[_0x1ca67e(0x181)]='0\x20response(s)';_0x59b9c4[_0x1ca67e(0x183)](_0x4b3542),_0x59b9c4[_0x1ca67e(0x183)](_0x9862c9),_0x1ced84['appendChild'](_0x59b9c4),_0x387784[_0x1ca67e(0x183)](_0x1ced84),QuestionsElementsList[_0x1ca67e(0x18a)](_0x387784);const _0x130332=document['createElement'](_0x1ca67e(0x166));_0x130332['classList'][_0x1ca67e(0x15c)](_0x1ca67e(0x18d));const _0x3587c0=document[_0x1ca67e(0x190)](_0x1ca67e(0x166));_0x3587c0['classList'][_0x1ca67e(0x15c)](_0x1ca67e(0x15f)),_0x3587c0[_0x1ca67e(0x1a0)](_0x1ca67e(0x13c),_0x77d1fd['id']),_0x3587c0['textContent']='☆',fetch(_0x1ca67e(0x19e))[_0x1ca67e(0x19f)](_0xa9390a=>_0xa9390a[_0x1ca67e(0x198)]())[_0x1ca67e(0x19f)](_0x57d429=>{const _0x4fddfb=_0x1ca67e;Array[_0x4fddfb(0x180)](_0x57d429)?_0x57d429[_0x4fddfb(0x15a)](_0x4a6482=>_0x4a6482==_0x77d1fd['id'])?(_0x3587c0[_0x4fddfb(0x1a2)][_0x4fddfb(0x15c)](_0x4fddfb(0x187)),_0x3587c0[_0x4fddfb(0x181)]='★'):(_0x3587c0[_0x4fddfb(0x1a2)]['remove']('favori_active'),_0x3587c0[_0x4fddfb(0x181)]='☆'):(_0x3587c0[_0x4fddfb(0x1a2)][_0x4fddfb(0x137)](_0x4fddfb(0x187)),_0x3587c0['textContent']='☆');}),_0x3587c0[_0x1ca67e(0x1a3)]=function(){const _0x29196d=_0x1ca67e;AddFavori(_0x77d1fd['id']),_0x3587c0['classList'][_0x29196d(0x139)](_0x29196d(0x187))?(_0x3587c0['classList']['remove'](_0x29196d(0x187)),_0x3587c0['textContent']='☆'):(_0x3587c0[_0x29196d(0x1a2)]['add']('favori_active'),_0x3587c0[_0x29196d(0x181)]='★');},_0x130332[_0x1ca67e(0x183)](_0x3587c0);const _0x249026=document[_0x1ca67e(0x190)](_0x1ca67e(0x166));_0x249026['classList'][_0x1ca67e(0x15c)](_0x1ca67e(0x138));const _0xac0460=document[_0x1ca67e(0x190)](_0x1ca67e(0x166));_0xac0460[_0x1ca67e(0x1a2)][_0x1ca67e(0x15c)]('upvote_text'),_0xac0460[_0x1ca67e(0x181)]='+';const _0x1fa0c3=document[_0x1ca67e(0x190)]('p');_0x1fa0c3[_0x1ca67e(0x1a2)][_0x1ca67e(0x15c)]('upvote_count'),_0x1fa0c3[_0x1ca67e(0x1a0)](_0x1ca67e(0x13c),_0x77d1fd['id']),_0x1fa0c3[_0x1ca67e(0x181)]=_0x77d1fd[_0x1ca67e(0x147)],_0x249026['appendChild'](_0xac0460),_0x249026[_0x1ca67e(0x183)](_0x1fa0c3),_0x130332['appendChild'](_0x249026);const _0xf3c339=document[_0x1ca67e(0x190)](_0x1ca67e(0x166));_0xf3c339[_0x1ca67e(0x1a2)][_0x1ca67e(0x15c)]('downvote_container'),console[_0x1ca67e(0x136)](_0x77d1fd);if(_0x77d1fd['user_vote']=='upvoted')_0x249026[_0x1ca67e(0x15d)][_0x1ca67e(0x16e)]=_0x1ca67e(0x171);else _0x77d1fd['user_vote']==_0x1ca67e(0x17f)&&(_0xf3c339[_0x1ca67e(0x15d)][_0x1ca67e(0x16e)]=_0x1ca67e(0x179));const _0xeff26d=document[_0x1ca67e(0x190)](_0x1ca67e(0x166));_0xeff26d[_0x1ca67e(0x1a2)][_0x1ca67e(0x15c)](_0x1ca67e(0x189)),_0xeff26d[_0x1ca67e(0x181)]='-';const _0x32a839=document['createElement']('p');_0x32a839['classList']['add'](_0x1ca67e(0x12d)),_0x32a839['setAttribute'](_0x1ca67e(0x13c),_0x77d1fd['id']),_0x32a839[_0x1ca67e(0x181)]=_0x77d1fd[_0x1ca67e(0x186)],_0xf3c339[_0x1ca67e(0x183)](_0xeff26d),_0xf3c339[_0x1ca67e(0x183)](_0x32a839),_0x130332[_0x1ca67e(0x183)](_0xf3c339),_0x387784['appendChild'](_0x130332),questionsList[_0x1ca67e(0x183)](_0x387784);if(_0x77d1fd[_0x1ca67e(0x193)]==_0x1ca67e(0x162))_0x249026['style'][_0x1ca67e(0x16e)]=_0x1ca67e(0x195);else _0x77d1fd[_0x1ca67e(0x193)]==_0x1ca67e(0x17f)&&(_0xf3c339[_0x1ca67e(0x15d)][_0x1ca67e(0x16e)]='rgb(196,\x2077,\x2086)');let _0x36199e=document[_0x1ca67e(0x190)](_0x1ca67e(0x166));_0x36199e['classList'][_0x1ca67e(0x15c)](_0x1ca67e(0x135)),_0x1ced84[_0x1ca67e(0x191)](_0x1ca67e(0x144),()=>{const _0x3fa5a4=_0x1ca67e;window['location']['href']=_0x3fa5a4(0x177)+window[_0x3fa5a4(0x157)][_0x3fa5a4(0x13a)]+_0x3fa5a4(0x16f)+_0x77d1fd['id'];}),_0x249026['onclick']=function(){const _0x5de25d=_0x1ca67e;_0x249026[_0x5de25d(0x15d)][_0x5de25d(0x16e)]=='rgb(104,\x20195,\x20163)'?_0x249026['style'][_0x5de25d(0x16e)]='':(_0x249026['style'][_0x5de25d(0x16e)]=_0x5de25d(0x195),_0xf3c339[_0x5de25d(0x15d)][_0x5de25d(0x16e)]==_0x5de25d(0x18b)&&(_0xf3c339[_0x5de25d(0x15d)][_0x5de25d(0x16e)]='')),socket[_0x5de25d(0x140)](JSON[_0x5de25d(0x19b)]({'type':_0x5de25d(0x12a),'content':_0x77d1fd['id'],'session_id':getCookie('session')}));},_0xf3c339[_0x1ca67e(0x1a3)]=function(){const _0x36cdf9=_0x1ca67e;_0xf3c339[_0x36cdf9(0x15d)]['backgroundColor']=='rgb(196,\x2077,\x2086)'?_0xf3c339['style'][_0x36cdf9(0x16e)]='':(_0xf3c339[_0x36cdf9(0x15d)]['backgroundColor']=_0x36cdf9(0x18b),_0x249026['style'][_0x36cdf9(0x16e)]=='rgb(104,\x20195,\x20163)'&&(_0x249026['style'][_0x36cdf9(0x16e)]='')),socket[_0x36cdf9(0x140)](JSON[_0x36cdf9(0x19b)]({'type':_0x36cdf9(0x172),'content':_0x77d1fd['id'],'session_id':getCookie(_0x36cdf9(0x159))}));};});questionsList[_0xe06558(0x15d)]['display']='',checkHighlight();if(_0x182163==null){let _0x107e19=document[_0xe06558(0x17a)](_0xe06558(0x194));_0x107e19[_0xe06558(0x181)]='0\x20question(s)';}}