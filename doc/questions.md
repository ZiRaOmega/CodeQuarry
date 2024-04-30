
### Data Structures

#### `Question`
Represents a question within the application.

**Fields**:
- `Id`: Unique identifier for the question.
- `SubjectTitle`: Title of the subject associated with the question.
- `SubjectID`: Identifier for the subject.
- `Title`: Title of the question.
- `Description`: A brief description of the question.
- `Content`: Detailed content of the question.
- `CreationDate`: Timestamp of when the question was created.
- `Creator`: Username of the user who created the question.
- `Upvotes`: Number of upvotes the question has received.
- `Downvotes`: Number of downvotes the question has received.
- `Responses`: A slice of `Response` objects representing the answers to the question.
- `UserVote`: The current user's voting status on the question ("upvoted" or "downvoted").

### Functions

#### `FetchQuestionsBySubject`
Retrieves questions based on a subject ID from the database, optionally filtering to include all subjects.

**Parameters**:
- `db`: The database connection.
- `subjectID`: The identifier for the subject, or "all" to retrieve questions across all subjects.
- `user_id`: The identifier of the current user to check voting status.

**Returns**:
- `([]Question, error)`: A slice of `Question` structs and an error if the operation fails.

#### `FetchQuestionByQuestionID`
Retrieves a detailed view of a question based on its ID.

**Parameters**:
- `db`: The database connection.
- `questionID`: The identifier of the question.

**Returns**:
- `(Question, error)`: A `Question` struct filled with detailed information and an error if the fetch is unsuccessful.

#### `QuestionsHandler`
Provides an HTTP handler function for retrieving questions based on the subject ID provided in the HTTP request.

**Usage**:
```go
http.HandleFunc("/questions", QuestionsHandler(db))
```

#### `QuestionViewerHandler`
Serves as an HTTP handler for displaying detailed information about a single question.

**Usage**:
```go
http.HandleFunc("/question/view", QuestionViewerHandler(db))
```

#### `CreateQuestion`
Inserts a new question into the database.

**Parameters**:
- `db`: The database connection.
- `question`: The `Question` struct containing question data.
- `user_id`: The user ID of the creator.
- `subject_id`: The subject ID associated with the question.

**Returns**:
- `error`: An error object if the insertion fails.

#### `FetchSubjectWithQuestionCount`
Fetches a subject along with a count of how many questions are associated with it.

**Parameters**:
- `db`: The database connection.
- `subjectId`: The subject ID.

**Returns**:
- `(Subject, error)`: A `Subject` struct including a count of related questions, and an error if the fetch fails.

#### `FetchQuestionsByUserID`
Retrieves all questions created by a specific user.

**Parameters**:
- `db`: The database connection.
- `userID`: The user ID whose questions are to be fetched.

**Returns**:
- `([]Question, error)`: A slice of `Question` structs and an error if the operation fails.

#### `FetchVotedQuestions`
Retrieves the voting status of questions by a specific user.

**Parameters**:
- `db`: The database connection.
- `userID`: The user ID to fetch voting data for.

**Returns**:
- `([]QuestionVote, error)`: A slice of `QuestionVote` structs indicating the user's votes on specific questions, and an error if the operation fails.

### Best Practices

- **Error Handling**: All functions include comprehensive error handling to ensure robustness and reliability.
- **Database Interactions**: SQL statements are meticulously crafted to prevent SQL injection and ensure performance.
- **JSON Responses**: API endpoints return data in JSON format, adhering to RESTful principles and ensuring compatibility with modern web clients.

### Example Usage

Deploying an HTTP server with handlers for fetching and viewing questions:

```go
func main() {
    db := app.InitDB("your-dsn")
    http.HandleFunc("/questions", app.QuestionsHandler(db))
    http.HandleFunc("/question/view", app.QuestionViewerHandler(db))
    http.ListenAndServe(":8080", nil)
}
```