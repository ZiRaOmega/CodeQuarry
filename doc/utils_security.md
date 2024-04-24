#### 1. `ContainsXSS`

**Function Signature**:
```go
func ContainsXSS(input string) bool
```

**Purpose**:
Checks if the input string contains common patterns that may indicate an XSS attack.

**Parameters**:
- `input string`: The string to be checked for XSS patterns.

**Return Value**:
- `bool`: Returns `true` if any XSS patterns are detected, otherwise returns `false`.

**Common XSS Contexts Checked**:
- `<script>` tags
- 'javascript:' pseudo-protocol
- HTML event handlers (e.g., `onclick`)
- `<iframe>` tags
- `<img>` tags with `src='data:'`

#### 2. `SanitizeXSS`

**Function Signature**:
```go
func SanitizeXSS(input string) string
```

**Purpose**:
Removes common XSS patterns from the input string, effectively sanitizing the string.

**Parameters**:
- `input string`: The string to be sanitized.

**Return Value**:
- `string`: The sanitized string with all XSS patterns replaced with an empty string.

#### 3. `ContainsSQLi`

**Function Signature**:
```go
func ContainsSQLi(input string) bool
```

**Purpose**:
Determines if the input string contains patterns that might indicate an SQL injection attack.

**Parameters**:
- `input string`: The string to check for SQL injection patterns.

**Return Value**:
- `bool`: Returns `true` if SQL injection patterns are detected, otherwise returns `false`.

**Common SQLi Contexts Checked**:
- SQL keywords (e.g., SELECT, INSERT, UPDATE, DELETE)
- Comments (e.g., `--`, `/* */`)
- SQL operators like `UNION`, `OR`, `AND`

#### 4. `SanitizeSQLi`

**Function Signature**:
```go
func SanitizeSQLi(input string) string
```

**Purpose**:
Removes common SQL injection patterns from the input string, providing basic sanitization against SQL injection attacks.

**Parameters**:
- `input string`: The string to be sanitized.

**Return Value**:
- `string`: The sanitized string with all detected SQL injection patterns replaced with an empty string.

### Security Considerations

While these functions provide basic protection against some common web security threats, they are not foolproof:
- **Regex Limitations**: The reliance on regular expressions for detection and sanitization may not cover all possible variations of XSS and SQLi attacks, especially sophisticated or obfuscated attacks.
- **Contextual Security**: Security measures must be context-specific. For example, SQL query sanitization is better handled by using prepared statements rather than simple string sanitization.

### Best Practices

- **Combine Approaches**: Use these functions in conjunction with other security practices, such as validating and encoding user inputs based on the context in which they are used (HTML context, JavaScript context, SQL queries, etc.).
- **Update Regularly**: Regularly update the patterns to adapt to new vulnerabilities and attack vectors as they emerge.
