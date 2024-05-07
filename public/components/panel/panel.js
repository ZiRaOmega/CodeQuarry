
function editQuestion(id) {
    const questionElement = document.querySelector(`[data-question-id="${id}"]`);
    const inputs = questionElement.querySelectorAll('.input-field');
    const textareas = questionElement.querySelectorAll('.textarea-field');

    const data = {
        type: "editQuestionPanel",
        content: {
            id: id,
            title: inputs[0].value,
            description: textareas[0].value,
            content: textareas[1].value,
            creationDate: inputs[1].value,
            updateDate: inputs[2].value,
            upvotes: inputs[3].value,
            downvotes: inputs[4].value
        },
        session_id: getCookie("session")
    };
    socket.send(JSON.stringify(data));
}

function editResponse(responseId,question_id) {
    // Select the response element by its data attribute
    const responseElement = document.querySelector(`.response[data-response-id="${responseId}"]`);

    // Retrieve content from textarea and input fields
    const inputs = responseElement.querySelectorAll('.input-field');
    const textareas = responseElement.querySelectorAll('.textarea-field');
    //Create the data object 
    const data = {
        type: "editResponsePanel",
        content: {
            id: responseId,
            question_id: question_id,
            content: textareas[0].value,
            description: inputs[0].value,
            creationDate: inputs[1].value,
            updateDate: inputs[2].value,
            upvotes: inputs[3].value,
            downvotes: inputs[4].value
        },
        session_id: getCookie("session")
    };
    
    // Send the JSON stringified data through the WebSocket
    socket.send(JSON.stringify(data));
}

function editSubject(id){
    const subjectElement = document.querySelector(`[data-subject-id="${id}"]`);
    const inputs = subjectElement.querySelectorAll('.input-field');
    const textareas = subjectElement.querySelectorAll('.textarea-field');

    const data = {
        type: "editSubjectPanel",
        content: {
            id: id,
            title: inputs[0].value,
            description: textareas[0].value,
            creationDate: inputs[1].value,
            updateDate: inputs[2].value,
        },
        session_id: getCookie("session")
    };
    socket.send(JSON.stringify(data));
}

function addSubject(){
    const newSubjectElement = document.querySelector("#add_subject")
    const inputs = newSubjectElement.querySelector('.input-field');
    const textareas = newSubjectElement.querySelector('.textarea-field');
    console.log(newSubjectElement)
    console.log(inputs)
    console.log(textareas)
    if (textareas.value == "" || inputs.value == ""){return}

    const data = {
        type: "addSubjectPanel",
        content: {
            title: inputs.value,
            description: textareas.value,
        },
        session_id: getCookie("session")
    };
    socket.send(JSON.stringify(data));
    //Reset fields
    inputs.value = "";
    textareas.value = "";

}

function deleteSubject(id){
    const data = {
        type: "deleteSubjectPanel",
        content: {
            id: id,
        },
        session_id: getCookie("session")
    };
    socket.send(JSON.stringify(data));
}

function deleteQuestion(id){
    const data = {
        type: "deleteQuestionPanel",
        content: {
            id: id,
        },
        session_id: getCookie("session")
    };
    socket.send(JSON.stringify(data));
}
function deleteResponse(id,question_id){
    const data = {
        type: "deleteResponsePanel",
        content: {
            id: id,
            question_id: question_id
        },
        session_id: getCookie("session")
    };
    socket.send(JSON.stringify(data));
}

function editUser(id){
    /** <div class="user_details" data-user-id="{{.ID}}">
                <input class="input-field" type="text" value="{{.FirstName}} " /><br>
                <input class="input-field" type="text" value="{{.LastName}}" /><br>
                <input class="input-field" type="text" value="@{{.Username}}" /><br>
                <input class="input-field" type="email" value="{{.Email}}" /><br>
                <textarea class="textarea-field">{{if .Bio.Valid}}{{.Bio.String}}{{else}}Not provided{{end}}</textarea><br>
                <input class="input-field" type="url" value="{{if .Website.Valid}}{{.Website.String}}{{end}}" /><br>
                <input class="input-field" type="url" value="{{if .GitHub.Valid}}{{.GitHub.String}}{{end}}" /><br>
                <input class="input-field" type="number" value="{{if .XP.Valid}}{{.XP.Int64}}{{else}}0{{end}}" /><br>
                <input class="input-field" type="text" value="{{if .Rank.Valid}}{{.Rank.String}}{{else}}Unranked{{end}}" /><br>
                <input class="input-field" type="date" value="{{if .SchoolYear.Valid}}{{.SchoolYear.Time}}{{end}}" /><br>
                <button >Edit</button>
                <button >Delete</button>
            </div>*/
    const userElement = document.querySelector(`[data-user-id="${id}"]`);
    const inputs = userElement.querySelectorAll('.input-field');
    const textareas = userElement.querySelectorAll('.textarea-field');
    //Create data object
    const data = {
        type: "editUserPanel",
        content: {
            id: id,
            firstname : inputs[0].value,
            lastname : inputs[1].value,
            username : inputs[2].value.replace("@",""),
            email : inputs[3].value,
            bio : textareas[0].value,
            website : inputs[4].value,
            github : inputs[5].value,
            xp : inputs[6].value,
            rank : inputs[7].value,
            schoolyear : inputs[8].value
        },
        session_id: getCookie("session")
    };
    //Send data through WebSocket
    socket.send(JSON.stringify(data));
}

function deleteUser(id){
    const data = {
        type: "deleteUserPanel",
        content: {
            id: id,
        },
        session_id: getCookie("session")
    };
    socket.send(JSON.stringify(data));
}