let QuestionsElementsList = [];
const questionContainer = document.getElementById("questions_container");
const questionsList = document.getElementById("questions_list");

document.addEventListener("DOMContentLoaded", function () {
    subjectId = window.location.pathname.split("/")[2];

    (async () => {
        try {
            const response = await fetch(`/api/questions?subjectId=${subjectId}`);
            //.then((response) => response.json());
            const questions = await response.json();
            console.log(questions);
            createFilter(questions);
            createQuestions(questions);
        } catch (error) {
            const errorH1 = document.createElement("h1");
            errorH1.textContent = "An error occured while fetching the questions";
            errorH1.style.color = "red";
            questionsList.appendChild(errorH1);
            console.error("There was a problem with your fetch operation:", error);
        }
    })();
});

function createFilter(questions) {
    questionTrackerCount = document.getElementById("question_tracker_count");

    if (questions == null) {
        questionTrackerCount.textContent = "0 question(s)";
        return;
    } else {
        questionTrackerCount.textContent = `${questions.length} question(s)`;
    }

    document.getElementById("filter_nbr_of_comments").onclick = () =>
        sortByNumberOfComments(questions);
    document.getElementById("filter_old").onclick = () =>
        sortOldestToNewest(questions);
    document.getElementById("filter_new").onclick = () =>
        sortNewestToOldest(questions);
    document.getElementById("filter_popular").onclick = () =>
        sortByUpvotes(questions);
    document.getElementById("filter_unpopular").onclick = () =>
        sortByDownvotes(questions);
}

function sortByNumberOfComments(questions) {
    questions.forEach((q) => (q.responses = q.responses || []));
    questions.sort((a, b) => b.responses.length - a.responses.length);
    refreshQuestionView(questions);
}

function sortOldestToNewest(questions) {
    questions.sort(
        (a, b) => new Date(a.creation_date) - new Date(b.creation_date)
    );
    refreshQuestionView(questions);
}

function sortNewestToOldest(questions) {
    questions.sort(
        (a, b) => new Date(b.creation_date) - new Date(a.creation_date)
    );
    refreshQuestionView(questions);
}

function sortByUpvotes(questions) {
    questions.sort((a, b) => b.upvotes - a.upvotes);
    refreshQuestionView(questions);
}

function sortByDownvotes(questions) {
    questions.sort((a, b) => a.upvotes - b.upvotes);
    refreshQuestionView(questions);
}

function refreshQuestionView(questions) {
    questionsList.innerHTML = ""; // Clear previous questions
    createFilter(questions);
    createQuestions(questions);
}

function createQuestions(questions) {
    questions.forEach((question) => {
        const questionElement = createQuestionElement(question);
        questionsList.appendChild(questionElement);
    });
    checkHighlight();
}

function trololo(questions) {
    // Clear previous questions
    if (questions != null)
        questions.forEach((question) => {
            const questionContainer = document.createElement("div");
            questionContainer.classList.add("question");
            const clickable_container = document.createElement("div");
            clickable_container.classList.add("clickable_container");
            // Add subject title tag
            const subjectTag = document.createElement("div");
            subjectTag.classList.add("subject_tag");
            subjectTag.textContent = question.subject_title;
            questionContainer.appendChild(subjectTag);

            const questionTitle = document.createElement("h3");
            questionTitle.classList.add("question_title");
            questionTitle.textContent = question.title;
            clickable_container.appendChild(questionTitle);

            const questionDescription = document.createElement("p");
            questionDescription.classList.add("question_description");
            questionDescription.textContent = question.description;
            clickable_container.appendChild(questionDescription);

            const questionContent = document.createElement("p");
            questionContent.classList.add("question_content");
            questionContent.textContent = question.content;
            const preDiv = document.createElement("pre");
            const code = document.createElement("code");
            preDiv.appendChild(code);
            code.innerHTML = question.content;
            clickable_container.appendChild(preDiv);

            const ContainerCreatorAndDate = document.createElement("div");
            ContainerCreatorAndDate.classList.add("creator_and_date_container");

            const questionDate = document.createElement("p");
            questionDate.classList.add("question_creation_date");
            questionDate.textContent = `Publié le: ${new Date(
                question.creation_date
            ).toLocaleDateString()}`;
            ContainerCreatorAndDate.appendChild(questionDate);

            const questionCreator = document.createElement("p");
            questionCreator.classList.add("question_creator");
            questionCreator.textContent = "Publié par";
            const creatorName = document.createElement("span");
            creatorName.textContent = question.creator;
            creatorName.classList.add("creator_name");
            questionCreator.appendChild(creatorName);
            const responsesCounter = document.createElement("p");
            responsesCounter.classList.add("responses_counter");
            if (Array.isArray(question.responses)) {
                // Set text content to the length of the responses array
                responsesCounter.textContent = `${question.responses.length} reponse(s)`;
            } else {
                // Handle cases where 'responses' might not be defined or not an array
                responsesCounter.textContent = "0 reponse(s)";
            }
            ContainerCreatorAndDate.appendChild(responsesCounter);
            ContainerCreatorAndDate.appendChild(questionCreator);
            clickable_container.appendChild(ContainerCreatorAndDate);
            questionContainer.appendChild(clickable_container);
            QuestionsElementsList.push(questionContainer);
            const voteContainer = document.createElement("div");
            voteContainer.classList.add("vote_container");
            const upvoteContainer = document.createElement("div");
            upvoteContainer.classList.add("upvote_container");
            const upvoteText = document.createElement("div");
            upvoteText.classList.add("upvote_text");
            upvoteText.textContent = "+";
            const upvoteCount = document.createElement("p");
            upvoteCount.classList.add("upvote_count");
            upvoteCount.setAttribute("data-question-id", question.id);
            upvoteCount.textContent = question.upvotes;
            upvoteContainer.appendChild(upvoteText);
            upvoteContainer.appendChild(upvoteCount);
            voteContainer.appendChild(upvoteContainer);
            const downvoteContainer = document.createElement("div");
            downvoteContainer.classList.add("downvote_container");
            if (question.user_vote == "upvoted") {
                upvoteContainer.style.backgroundColor = "green";
            } else if (question.user_vote == "downvoted") {
                downvoteContainer.style.backgroundColor = "red";
            }
            const downvoteText = document.createElement("div");
            downvoteText.classList.add("downvote_text");
            downvoteText.textContent = "-";
            const downvoteCount = document.createElement("p");
            downvoteCount.classList.add("downvote_count");
            downvoteCount.setAttribute("data-question-id", question.id);
            downvoteCount.textContent = question.downvotes;
            downvoteContainer.appendChild(downvoteText);
            downvoteContainer.appendChild(downvoteCount);
            voteContainer.appendChild(downvoteContainer);
            questionContainer.appendChild(voteContainer);
            questionsList.appendChild(questionContainer);
            if (question.user_vote == "upvoted") {
                upvoteContainer.style.backgroundColor = "rgb(104, 195, 163)";
            } else if (question.user_vote == "downvoted") {
                downvoteContainer.style.backgroundColor = "rgb(196, 77, 86)";
            }
            let responseContainer = document.createElement("div");
            responseContainer.classList.add("response_container");
            clickable_container.addEventListener("click", () => {
                window.location.href = `https://${window.location.hostname}/question_viewer?question_id=${question.id}`;
            });
            upvoteContainer.onclick = function () {
                //if upvoteContainer backgroundColor is green then remove the color
                if (upvoteContainer.style.backgroundColor == "rgb(104, 195, 163)") {
                    upvoteContainer.style.backgroundColor = "";
                } else {
                    upvoteContainer.style.backgroundColor = "rgb(104, 195, 163)";
                    if (downvoteContainer.style.backgroundColor == "rgb(196, 77, 86)") {
                        downvoteContainer.style.backgroundColor = "";
                    }
                }
                socket.send(
                    JSON.stringify({
                        type: "upvote",
                        content: question.id,
                        session_id: getCookie("session"),
                    })
                );
            };

            downvoteContainer.onclick = function () {
                //if downvoteContainer backgroundColor is red then remove the color
                if (downvoteContainer.style.backgroundColor == "rgb(196, 77, 86)") {
                    downvoteContainer.style.backgroundColor = "";
                } else {
                    downvoteContainer.style.backgroundColor = "rgb(196, 77, 86)";
                    if (upvoteContainer.style.backgroundColor == "rgb(104, 195, 163)") {
                        upvoteContainer.style.backgroundColor = "";
                    }
                }
                socket.send(
                    JSON.stringify({
                        type: "downvote",
                        content: question.id,
                        session_id: getCookie("session"),
                    })
                );
            };
        });
    questionsList.style.display = ""; // Show the questions list
    checkHighlight();
    if (questions == null) {
        questionTrackerCount.textContent = "0 question(s)";
    }
}


function createQuestionElement(question) {
    const questionElement = document.createElement("div");
    questionElement.classList.add("question");
    if (question.responses == null) {
        question.responses = [];
    }

    questionElement.innerHTML = htmlQuestionConstructor(question);

    // Highlight the code block
    // !!! TODO : use checkHighlight() function from detect_lang.js
    // make sure to check global script
    /* code = questionElement.querySelector("code");
    code = hljs.highlightElement(code); */

    // Upvote and downvote event listeners
    upvoteContainer = questionElement.querySelector(".upvote_container");
    downvoteContainer = questionElement.querySelector(".downvote_container");
    switchVoteColor(question, upvoteContainer, downvoteContainer);

    // Handle favori
    favori = questionElement.querySelector(".favori");
    manageFavorite(favori, question.id);

    // Clickable container redirect to question viewer
    clickable_container = questionElement.querySelector(".clickable_container")
        .addEventListener("click", () => {
            window.location.href = `https://${window.location.hostname}/question_viewer?question_id=${question.id}`;
        });

    return questionElement;
}

function manageFavorite(favori, questionId) {
    fetch("/api/favoris")
        .then((response) => response.json())
        .then((favoris) => {
            if (Array.isArray(favoris)) {
                if (favoris.some((f) => f == questionId)) {
                    favori.classList.add("favori_active");
                    favori.textContent = "★";
                    favori.style.backgroundColor = "gold";
                } else {
                    favori.classList.remove("favori_active");
                    favori.textContent = "☆";
                    favori.style.backgroundColor = "";
                }
            } else {
                favori.classList.remove("favori_active");
                favori.textContent = "☆";
                favori.style.backgroundColor = "";
            }
        })
        .catch((error) => {
            console.error("Problem with favori fetch:", error);
        });


    favori.onclick = function () {
        //AddFavori(questionId);
        if (favori.classList.contains("favori_active")) {
            favori.classList.remove("favori_active");
            favori.textContent = "☆";
            favori.style.backgroundColor = "";
        } else {
            favori.classList.add("favori_active");
            favori.textContent = "★";
            favori.style.backgroundColor = "gold";
        }
    };
}


function htmlQuestionConstructor(question) {
    return `
<div class="subject_tag">${question.subject_title}</div>
<div class="clickable_container">
    <h3 class="question_title">${question.title}</h3>
    <p class="question_description">${question.description}</p>
    <pre><code>${question.content}</code></pre>
    <div class="creator_and_date_container">
        <p class="question_creation_date">Publié le: ${new Date(question.creation_date).toLocaleDateString()}</p>
        <p class="responses_counter">${question.responses.length} reponse(s)</p>
        <p class="question_creator">Publié par <span class="creator_name">${question.creator}</span></p>
    </div>
</div>
<div class="vote_container">
    <div class="upvote_container">
        <div class="upvote_text">+</div>
        <p class="upvote_count" data-question-id="${question.id}">${question.upvotes}</p>
    </div>
    <div class="downvote_container">
        <div class="downvote_text">-</div>
        <p class="downvote_count" data-question-id="${question.id}">${question.downvotes}</p>
    </div>
    <div class="favori" data-question-id="${question.id}">☆</div>
</div>
`;
}

function switchVoteColor(question, upvoteContainer, downvoteContainer) {
    if (question.user_vote == "upvoted") {
        upvoteContainer.style.backgroundColor = "rgb(104, 195, 163)";
    } else if (question.user_vote == "downvoted") {
        downvoteContainer.style.backgroundColor = "rgb(196, 77, 86)";
    }

    upvoteContainer.addEventListener("click", () => {
        if (upvoteContainer.style.backgroundColor == "rgb(104, 195, 163)") {
            upvoteContainer.style.backgroundColor = "";
        } else {
            upvoteContainer.style.backgroundColor = "rgb(104, 195, 163)";
            if (downvoteContainer.style.backgroundColor == "rgb(196, 77, 86)") {
                downvoteContainer.style.backgroundColor = "";
            }
        }

        socket.send(
            JSON.stringify({
                type: "upvote",
                content: question.id,
                session_id: getCookie("session"),
            })
        );
    });

    downvoteContainer.addEventListener("click", () => {
        if (downvoteContainer.style.backgroundColor == "rgb(196, 77, 86)") {
            downvoteContainer.style.backgroundColor = "";
        } else {
            downvoteContainer.style.backgroundColor = "rgb(196, 77, 86)";
            if (upvoteContainer.style.backgroundColor == "rgb(104, 195, 163)") {
                upvoteContainer.style.backgroundColor = "";
            }
        }

        socket.send(
            JSON.stringify({
                type: "downvote",
                content: question.id,
                session_id: getCookie("session"),
            })
        );
    });
}