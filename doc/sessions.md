## Function `CreateSession` Documentation

The `CreateSession` function is designed to authenticate a user by their username, create a new session record in the database, and set an HTTP cookie with the session identifier. This function is a key component in managing user sessions securely and efficiently in a web application.

### Function Signature

```go
func CreateSession(username string, db *sql.DB, w http.ResponseWriter) error
```

#### Parameters

- `username string`: The username of the user for whom the session is being created.
- `db *sql.DB`: A pointer to the database connection pool used for SQL operations.
- `w http.ResponseWriter`: The HTTP response writer used to send back the session cookie to the client's browser.

#### Return Value

- `error`: Returns an error object that will be non-nil if an error occurs during the session creation process.

### Detailed Breakdown

```go
user_id, err := getUserID(username, db)
```

- **Retrieve User ID**: Fetches the user ID associated with the given username from the database. This ID is necessary to link the session to the correct user account.

```go
if err != nil {
    return err
}
```

- **Error Handling**: Immediately returns the error if fetching the user ID fails, stopping further execution of the function.

```go
user_uuid := UUID.NewV4().String()
createdAt := time.Now()
expireAt := createdAt.Add(Cookie_Expiration)
```

- **Session Identifier**: Generates a new UUID v4, which is used as a unique identifier for the session.
- **Session Timing**:
  - `createdAt`: Records the current time as the session's start time.
  - `expireAt`: Calculates the expiration time of the session based on a predefined duration (`Cookie_Expiration`).

```go
_, err = db.Exec("INSERT INTO Sessions(user_id,uuid,created_at,expire_at) VALUES($1,$2,$3,$4)", user_id, user_uuid, createdAt, expireAt)
```

- **Database Insertion**: Inserts the new session into the `Sessions` table with the user ID, UUID, creation time, and expiration time.
- **SQL Injection Prevention**: Uses parameterized SQL queries to safeguard against SQL injection.

```go
if err != nil {
    Log(ErrorLevel, "Error creating session "+err.Error())
}
```

- **Log Error**: If there's an error during the database operation, logs the error with an "ErrorLevel" indicating the severity.

```go
cookie := http.Cookie{
    Name:     "session",
    Value:    user_uuid,
    Expires:  expireAt,
    HttpOnly: true,
}
http.SetCookie(w, &cookie)
```

- **Cookie Configuration**:
  - `Name`: The name of the cookie, set to "session".
  - `Value`: The UUID of the newly created session.
  - `Expires`: The expiration time of the cookie, aligning with the session's expiration.
  - `HttpOnly`: Set to true to prevent access to the cookie via client-side scripts, enhancing security against XSS attacks.
- **Set Cookie**: Attaches the configured cookie to the HTTP response, effectively sending it to the client's browser.

```go
return err
```

- **Return Error**: Returns any error that occurred during session creation, allowing the calling function to handle it appropriately.

### Best Practices

- **Secure Cookie Handling**: Setting `HttpOnly` to true is crucial for preventing access to the cookie value via client-side JavaScript, which is a common vector for XSS attacks.
- **Error Logging and Handling**: Robust error handling and logging are implemented to facilitate troubleshooting and ensure that errors do not go unnoticed.
- **Parameterized SQL Queries**: Using parameterized queries is a best practice to prevent SQL injection, a prevalent security vulnerability in web applications.
