
#### `Question`
Represents the data structure for a question within the system.

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
- `Responses`: List of responses to the question.

### Functions

#### `FetchQuestionsBySubject`
Fetches questions filtered by the subject from the database.

**Parameters**:
- `db`: Database connection.
- `subjectID`: Identifier of the subject or "all" to fetch questions across all subjects.

**Returns**:
- `[]Question`: Slice of questions.
- `error`: Error object in case of failure.

#### `QuestionsHandler`
HTTP handler function that retrieves questions based on the subject ID provided in the query parameters and sends them back as JSON.

**Usage**:
```go
http.HandleFunc("/questions", QuestionsHandler(db))
```

#### `QuestionViewerHandler`
HTTP handler for displaying individual questions based on a provided question ID.

**Usage**:
```go
http.HandleFunc("/view-question", QuestionViewerHandler(db))
```

#### `CreateQuestion`
Inserts a new question into the database.

**Parameters**:
- `db`: Database connection.
- `question`: Question object containing the details to be inserted.
- `user_id`: User ID of the question's creator.
- `subject_id`: Subject ID associated with the question.

**Returns**:
- `error`: Error object if the insertion fails.

#### `FetchSubjectWithQuestionCount`
Retrieves a subject along with the count of questions linked to it.

**Parameters**:
- `db`: Database connection.
- `subjectId`: Identifier for the subject.

**Returns**:
- `Subject`: Subject object with the question count.
- `error`: Error object in case of failure.

#### `FetchQuestionsByUserID`
Fetches questions created by a specific user.

**Parameters**:
- `db`: Database connection.
- `userID`: User ID whose questions are to be fetched.

**Returns**:
- `[]Question`: Slice of questions.
- `error`: Error object in case of failure.
