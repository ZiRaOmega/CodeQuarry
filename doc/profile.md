
## Package `app` Documentation

The `app` package contains a collection of functions and types designed to manage user profiles in a web application environment. This includes handling HTTP requests for profile viewing and updating, securely fetching and storing user data, and managing file uploads.

### Functions Overview

#### 1. `ProfileHandler`

**Function Signature**:
```go
func ProfileHandler(db *sql.DB) http.HandlerFunc
```

**Purpose**:
Creates an HTTP handler function that retrieves user profile data based on a session ID stored in a cookie and renders it to a profile page using HTML templates.

**Workflow**:
- **Cookie Retrieval**: Attempts to retrieve a session cookie from the incoming HTTP request. If unsuccessful, it returns an HTTP error.
- **User Data Retrieval**: Uses the session ID from the cookie to fetch user details from the database. Handles errors by logging them and returning an HTTP error.
- **Template Parsing and Execution**: Parses multiple HTML template files and executes them to render the user's profile data. Handles parsing or execution errors by panicking, which is not recommended for production environments.

**Security Concerns**:
- Redirects to the error page or logs an error if the session cookie is missing or user data cannot be fetched, ensuring that unauthorized access is prevented.
- Uses panic for handling template errors, which should be replaced with more graceful error handling.

#### 2. `GetUser`

**Function Signature**:
```go
func GetUser(session_id string, db *sql.DB) (User, error)
```

**Purpose**:
Fetches user details from the database based on a session ID.

**Workflow**:
- **Session Verification**: Verifies the session by mapping the session ID to a user ID in the database.
- **Data Retrieval**: Fetches comprehensive user details based on the user ID and populates them into a `User` struct.
- **Error Handling**: Returns errors if database connections are not set or queries fail.

**Security Measures**:
- Uses prepared statements to prevent SQL injection.
- Properly handles database connections and statements with deferred closure.

#### 3. `UpdateProfileHandler`

**Function Signature**:
```go
func UpdateProfileHandler(db *sql.DB) http.HandlerFunc
```

**Purpose**:
Handles HTTP POST requests to update user profile data, ensuring that the session is valid and that the data provided does not contain harmful SQL or XSS content.

**Workflow**:
- **Method Check**: Ensures that only POST requests are processed.
- **Session Validation**: Checks if the session is valid; redirects to the login page if not.
- **Form Data Retrieval and Validation**: Retrieves data from the POST form, validates it for SQL injections and XSS attacks, and updates the user's data in the database.
- **Password Handling**: If the password is provided, hashes it before storing it in the database.

**Security Concerns**:
- Checks for SQL injections and XSS in user input before processing.
- Uses bcrypt for password hashing, which is a secure method of storing passwords.

#### 4. `FileUpload`

**Function Signature**:
```go
func FileUpload(r *http.Request) (string, error)
```

**Purpose**:
Handles file uploads for user avatars, ensuring that only files of certain types are uploaded and stored correctly.

**Workflow**:
- **Multipart Form Parsing**: Parses the multipart form data from the request.
- **File Type Validation**: Checks if the file type is one of the allowed types (.jpg, .jpeg, .png, .gif).
- **File Storage**: Saves the file to a designated directory and returns the path.

**Security Measures**:
- Validates file types to prevent malicious file uploads.
- Generates a new file name using a UUID to avoid conflicts and potential security issues related to user input file names.

### Conclusion

The `app` package provides robust functionalities for user profile management within a web application, implementing several security measures to protect against common web vulnerabilities. However, certain areas, such as error handling and file upload security, could be further enhanced to align with best security practices.
### Documentation for the "profile" HTML Template in Go

The "profile" template is part of a web application developed in Go, designed for rendering a user profile page that allows users to view and update their personal information. This template utilizes Go's `text/template` package to inject dynamic data into HTML, ensuring that the user's current information is displayed and can be updated through a form submission.

#### Template Structure

```html
{{define "profile"}}
<!DOCTYPE html>
<html lang="en">
{{template "head" .}}
<body>
    {{template "header" .}}
    <div class="profile">
        <form action="/update-profile" method="POST" enctype="multipart/form-data">
            ...
        </form>
    </div>
    {{template "script" .}}
</body>
</html>
{{end}}
```

Each component of the HTML structure is detailed below:

#### Head and Header

- **Head and Header Sections**: These sections are injected via `{{template "head" .}}` and `{{template "header" .}}` respectively. These calls include external template parts that contain HTML and possibly JavaScript or CSS that are common across pages, such as meta tags, stylesheets, and navigational elements.

#### Profile Form

- **Form Attributes**:
  - `action="/update-profile"`: Specifies the endpoint where the form data will be sent upon submission.
  - `method="POST"`: Sets the HTTP method used when submitting the form, ensuring data is sent as part of the request body.
  - `enctype="multipart/form-data"`: Allows the form to include file uploads, necessary for updating the avatar image.

- **Form Inputs**:
  Each input field in the form is pre-populated with the user's existing data:
  - **Hidden ID**: A hidden input that carries the user's ID to ensure the server knows which user profile to update.
  - **Text Inputs**: Fields for last name, first name, username, email, etc., are typical text inputs.
  - **Password Input**: For updating the password, shown as a blank field for security reasons.
  - **Avatar**: Includes an image tag displaying the current avatar and a file input for uploading a new avatar.
  - **Bio**: A textarea for a more extensive user biography.
  - **Experience and Rank**: Numeric inputs for experience points and rank, which might be used in gamification aspects of the application.
  - **Date Inputs**: For birth date and school year, formatted for ease of use in HTML date pickers.
  - **Submit Button**: Allows submission of the form to update the user's profile.

#### Script

- **Script Injection**: `{{template "script" .}}` includes additional JavaScript necessary for client-side functionality, which could be scripts for validation, interactivity, or integration with other systems.

### Best Practices and Security Considerations

1. **Data Validation and Sanitization**:
   - Server-side: Ensure that all incoming data from the form submission is properly validated and sanitized to prevent SQL injection and XSS attacks.
   - Client-side: Use JavaScript to provide immediate feedback for validation before submission to enhance user experience.

2. **Use of Prepared Statements**:
   - When updating information in the database, always use prepared statements to prevent SQL injection, as shown in the server-side handling of the update functionality.

3. **Secure File Handling**:
   - When dealing with file uploads (e.g., avatars), ensure the server properly checks the file type and size, and handles the storage securely to prevent execution of malicious files.

4. **Password Handling**:
   - Ensure passwords are never displayed in forms or logs.
   - Use strong hashing algorithms (e.g., bcrypt) when storing updated passwords.

This template is a crucial component of the user management system, providing an interface for users to manage their personal data. The integration of this template within a larger application requires careful attention to both functionality and security to ensure a safe and user-friendly experience.