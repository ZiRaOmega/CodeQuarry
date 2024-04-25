## Go Templates Overview

Templates in Go provide a powerful way for generating text outputs based on predefined templates and dynamic data. Go supports templates through two primary packages: `text/template` and `html/template`. These packages allow developers to define templates that can then dynamically integrate data at runtime to produce outputs like HTML, emails, or other text formats.

### Packages Overview

- **`text/template`**: This package is designed for generating textual output from a template, where the output does not require any escaping for safety (e.g., generating configuration files, source code).
- **`html/template`**: This package is an extension of `text/template` and provides the same interface and functionality but is specifically designed for generating HTML output. It automatically secures the output by escaping HTML content to prevent Cross-Site Scripting (XSS) vulnerabilities.

### Template Parsing and Execution

#### 1. Define a Template

A template is a string or a file containing one or more definitions that can be combined with data to produce a final output. Templates use placeholders, often referred to as "actions," enclosed in `{{` and `}}`.

Example of a simple template:
```go
const tmpl = `Hello, {{.Name}}! Welcome to our service.`
```

#### 2. Parse the Template

Before a template can be used, it must be parsed. Parsing a template involves analyzing the template's syntax and preparing it for execution.

```go
t, err := template.New("welcome").Parse(tmpl)
if err != nil {
    log.Fatalf("parsing: %s", err)
}
```

#### 3. Execute the Template

Executing a template involves applying it to a specified data structure. The template engine replaces actions with the corresponding data from the provided structure.

```go
type User struct {
    Name string
}

data := User{Name: "John Doe"}

if err := t.Execute(os.Stdout, data); err != nil {
    log.Fatalf("execution: %s", err)
}
```

### Safety Features

The `html/template` package automatically escapes all inputs when rendering HTML, which ensures that the output is safe to display in user browsers. It prevents common web vulnerabilities such as XSS by ensuring that inputs are not interpreted as executable HTML or JavaScript.

### Best Practices

- **Contextual Escaping**: `html/template` understands HTML context and will escape data based on where it appears within your HTML (e.g., within a script tag or as a URL).
- **Prefer `html/template` for Web Content**: Always use `html/template` for generating HTML content to leverage its auto-escaping capabilities for better security.
- **Reuse Templates**: Parse templates once and execute them multiple times with different data. Parsing is a CPU-intensive operation.
- **Error Handling**: Always handle errors from parsing and executing templates, especially when dealing with user-generated content.

### Conclusion

Templates in Go are a flexible and secure way to generate dynamic text and HTML content. Understanding the distinction between `text/template` and `html/template`, as well as following best practices for their use, is crucial for creating secure and efficient applications.


## Function `SendTemplate` Documentation

The `SendTemplate` function in Go is designed to serve a specific HTML page constructed from multiple templates (header, footer, and a main template). This function is typically used in web applications to generate dynamic web pages with a consistent layout.

### Function Signature

```go
func SendTemplate(template_name string) http.HandlerFunc
```

#### Parameters

- `template_name string`: The base name of the main template file (excluding "public/" prefix and ".html" suffix) that will be combined with the header and footer to create a complete web page.

#### Return Type

- `http.HandlerFunc`: Returns an HTTP handler function that can be used to handle HTTP requests.

### Detailed Breakdown

```go
return func(w http.ResponseWriter, r *http.Request) {
```
- **Return a Closure**: This line returns an anonymous function that fits the `http.HandlerFunc` type, suitable for handling HTTP requests.

```go
log.Printf("[SendIndex:%s] New Client with IP: %s\n", r.URL.Path, r.RemoteAddr)
```
- **Logging Access**: Logs the accessed URL path and the IP address of the client making the request, which helps in monitoring and debugging client interactions.

```go
tmpl, err := template.ParseFiles("public/header.html", "public/footer.html", "public/"+template_name+".html")
```
- **Parse HTML Templates**: Attempts to parse the header, footer, and a specific template from the `public` directory. These files are assembled to form the full HTML document.
  
- **Dynamic Template Naming**: Constructs the path to the main template using the `template_name` parameter, allowing flexibility in serving different pages using the same structure.

```go
if err != nil {
    panic(err)
    // Parsing the HTML templates for the header, footer, and index. If there is an error, the program will panic and stop execution.
}
```
- **Error Handling**: Checks if there was an error in parsing the templates. If an error occurs, it panics, effectively stopping the server with an error message. This is a harsh way to handle errors and is not recommended for production environments.

```go
err = tmpl.ExecuteTemplate(w, template_name, nil)
```
- **Execute and Send Template**: Executes the specified template and sends the output directly to the client (`w`). Uses `template_name` as the identifier of the main template part to be executed.
- **Data Passing**: The `nil` argument implies that no data is passed to the templates, which means they should be designed to operate without external data inputs.

```go
if err != nil {
    panic(err)
    // Executing the "index" template and sending it to the client. If there is an error, the program will panic and stop execution.
}
```
- **Execution Error Handling**: Similar to the parsing error handling, it panics if there is an issue in executing the template, which terminates the server. In practice, it's better to handle such errors by logging them and sending an HTTP error response to the client.

## Comprehensive Documentation on HTML Templates in Go

The provided templates are structured using Go's `html/template` package, which is designed for generating HTML output that is safe against code injection. This is particularly important for web applications where user input is directly incorporated into the page. The templates defined (`header`, `footer`, `login`) are typical for a web application requiring user authentication.

### Template Definitions

#### `header` Template

```go
{{define "header"}}
<head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Code Quarry</title>
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.5.1/jquery.min.js"></script>
    <script src="https://cdn.jsdelivr.net/npm/sweetalert2@11"></script>
    <link rel="stylesheet" href="styles.css" />
</head>
{{end}}
```
- **Purpose**: Sets up the HTML head section for all pages. Includes metadata, script references (jQuery, SweetAlert2), and a CSS link.
- **Usage**: Included in each HTML page that needs a common header. Provides a consistent title and ensures that styles and scripts are uniformly available across pages.

#### `footer` Template

```go
{{define "footer"}}
<script src="../scripts/animation.js"></script>
<script src="../scripts/errors_obfuscate.js"></script>
<script src="../scripts/websocket.js"></script>
{{end}}
```
- **Purpose**: Defines the closing scripts that should be loaded after the main content of the webpage. This typically includes animations, error handling, and WebSocket interactions.
- **Usage**: Appended to the end of the HTML body to ensure scripts execute after the HTML content is fully loaded.

#### `login` Template

```go
{{define "login"}}
<!DOCTYPE html>
<html lang="en">
    {{template "header"}}
    <body>
        <!-- Structure for login and registration forms -->
        ...
        {{template "footer"}}
    </body>
</html>
{{end}}
```
- **Structure**: Combines the `header` and `footer` templates around specific content for a user login and registration page.
- **Content**:
  - **Login and Registration Forms**: Provides interactive forms for user authentication and account creation, using POST methods to secure endpoints `/login` and `/register`.
  - **Dynamic Content**: Contains toggles between login and registration views without needing to load new pages, presumably handled by JavaScript not included in the snippet.

### Rendering and Usage

These templates would be parsed and executed by a Go server as follows:

1. **Parsing Templates**: All templates are parsed using `template.ParseFiles()` or `template.ParseGlob()` if they are in separate files or a single file with multiple definitions respectively.
2. **Executing Templates**: Specific templates are executed based on the user's navigation, typically by calling `tmpl.ExecuteTemplate(w, "login", data)` where `data` contains dynamic content passed to templates.
3. **Error Handling**: Proper error checks should be implemented after parsing and executing to ensure that any issues with template rendering are caught and handled gracefully.

