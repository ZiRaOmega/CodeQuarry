
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
