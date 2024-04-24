**Language > Specialist**: Go > Web Session Management Expert
**Includes**: `database/sql`, `net/http`, `time`, `github.com/satori/go.uuid`
**Requirements**: Detailed documentation for two Go functions responsible for session management, including inserting sessions into a database and creating a session with associated cookie handling.
## Plan
1. Document the `insertSessionToDB` function, detailing its purpose, parameters, operations, and error handling.
2. Document the `CreateSession` function, explaining its comprehensive role in session creation, from user ID retrieval to setting session cookies.

---

## Package `session` Documentation

This documentation covers two crucial functions in the session management package, which facilitate the creation and storage of user sessions in a web application environment.

### 1. `insertSessionToDB`

**Function Signature**:
```go
func insertSessionToDB(db *sql.DB, user_id int, user_uuid string, createdAt time.Time, expireAt time.Time) error
```

**Purpose**:
Inserts a new session record into the database. This function is typically used to persist session details after generating a new session for a user.

**Parameters**:
- `db *sql.DB`: The database connection used to execute SQL commands.
- `user_id int`: The identifier of the user for whom the session is created.
- `user_uuid string`: The UUID of the new session, serving as a unique identifier.
- `createdAt time.Time`: The timestamp marking the creation of the session.
- `expireAt time.Time`: The timestamp indicating when the session will expire.

**Operations**:
1. **Prepare SQL Statement**: Prepares an SQL statement for inserting session data into the `Sessions` table.
2. **Execute SQL Statement**: Executes the prepared statement with the provided parameters.
3. **Error Handling**: Returns any errors encountered during the SQL operation, ensuring that calling functions can respond appropriately.

**Return Value**:
- `error`: Returns an error object if there is a failure in inserting the session data into the database.

### 2. `CreateSession`

**Function Signature**:
```go
func CreateSession(username string, db *sql.DB, w http.ResponseWriter) error
```

**Purpose**:
Creates a new session for a user, retrieves the user ID, generates a session UUID, inserts session details into the database, and sets a session cookie on the client's browser.

**Parameters**:
- `username string`: The username of the user for whom the session is being created.
- `db *sql.DB`: The database connection for querying and inserting data.
- `w http.ResponseWriter`: The HTTP response writer used to set cookies on the client side.

**Operations**:
1. **Retrieve User ID**: Calls `getUserIDFromDB` to fetch the user ID based on the username.
2. **Generate Session UUID**: Generates a new UUID v4 for the session.
3. **Calculate Timestamps**: Determines the current time and expiration time for the session.
4. **Insert Session into Database**: Utilizes `insertSessionToDB` to store session details.
5. **Set HTTP Cookie**: Configures and sets an HTTP cookie with the session UUID, which includes HttpOnly flag to enhance security.

**Error Handling**:
- Checks and returns errors at each critical step, ensuring that the process halts if a problem occurs (e.g., user ID retrieval fails, database insertion errors).

**Return Value**:
- `error`: Returns an error object if any step in the session creation process encounters issues.

### Best Practices

- **Security Measures**: Use HttpOnly cookies to prevent client-side scripts from accessing the session token.
- **Error Handling**: Proper error handling in each function helps prevent security vulnerabilities and ensures the application remains stable and reliable.
- **UUID for Sessions**: Using UUIDs ensures that session identifiers are unique and unpredictable, reducing the risk of session hijacking.

