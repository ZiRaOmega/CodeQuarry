package main_test

import (
	_ "github.com/lib/pq"
)

/*
var testDB *sql.DB

// initTestDB initializes a test database
func initTestDBMain(m *testing.M) *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		"codequarry",
		"CQ1234",
		"localhost",
		"5432",
		"codequarrytest",
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func TestMain(m *testing.M) {
	var err error
	testDB = initTestDBMain(m)

	err = populateTestData(testDB)
	if err != nil {
		log.Fatalf("Failed to populate test data: %v", err)
	}

	code := m.Run()

	os.Exit(code)
}

func populateTestData(db *sql.DB) error {
	_, err := db.Exec(`TRUNCATE TABLE question, users, subject RESTART IDENTITY CASCADE`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`INSERT INTO users (id_student, lastname, firstname, username, email, password, avatar, birth_date, bio, website, github, xp, rang_rank_) VALUES (1, 'Doe', 'John', 'johndoe', 'm@m.m', 'password', 'avatar.jpg', '2000-01-01', 'I am John Doe', 'https://example.com', 'https://github.com', 0, 0)`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`INSERT INTO subject (id_subject, title, description, creation_date, update_date) VALUES (1, 'Test Subject', 'Desc', '2002-06-04', '2002-06-04')`)
	if err != nil {
		return err
	}

	questions := []app.Question{
		{
			Title:       "Test Question 1",
			Description: "Description for test question 1",
			Content:     "Content for test question 1",
		},
		{
			Title:       "Test Question 2",
			Description: "Description for test question 2",
			Content:     "Content for test question 2",
		},
	}

	for _, question := range questions {
		err := app.CreateQuestion(db, question, 1, 1)
		if err != nil {
			return err
		}
	}

	return nil
}

func TestFetchQuestionsBySubject(t *testing.T) {
	subjectID := "all"
	userID := 1

	questions, err := app.FetchQuestionsBySubject(testDB, subjectID, userID)
	assert.NoError(t, err)
	assert.NotEmpty(t, questions)
}

func TestFetchQuestionByQuestionID(t *testing.T) {
	questionID := 1
	userID := 1

	question, err := app.FetchQuestionByQuestionID(testDB, questionID, userID)
	assert.NoError(t, err)
	assert.Equal(t, questionID, question.Id)
}

func TestQuestionsHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/questions?subjectId=all", nil)
	assert.NoError(t, err)

	req.AddCookie(&http.Cookie{Name: "session", Value: "validsessionid"})

	rr := httptest.NewRecorder()
	handler := app.QuestionsHandler(testDB)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	var questions []app.Question
	err = json.NewDecoder(rr.Body).Decode(&questions)
	assert.NoError(t, err)
	assert.NotEmpty(t, questions)
}

func TestQuestionViewerHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/question_viewer?question_id=1", nil)
	assert.NoError(t, err)

	req.AddCookie(&http.Cookie{Name: "session", Value: "validsessionid"})

	rr := httptest.NewRecorder()
	handler := app.QuestionViewerHandler(testDB)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "question_viewer")
}

func TestCreateQuestion(t *testing.T) {
	question := app.Question{
		Title:       "Test Title",
		Description: "Test Description",
		Content:     "Test Content",
	}

	err := app.CreateQuestion(testDB, question, 1, 1)
	assert.NoError(t, err)
}

func TestUserDeleteQuestion(t *testing.T) {
	err := app.UserDeleteQuestion(testDB, 1, 1)
	assert.NoError(t, err)
} */
