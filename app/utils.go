package app

import "regexp"

// ContainsXSS checks if the input string contains common XSS attack vectors
func ContainsXSS(input string) bool {
	// Patterns to check for common XSS contexts:
	// - <script> tags
	// - 'javascript:' pseudo-protocol
	// - HTML event handlers (e.g., onclick)
	// - <iframe> tags
	// - <img> tags with src='data:'
	patterns := []string{
		`(?i)<script.*?>.*?</script>`,
		`(?i)javascript:`,
		`(?i)on\w+="[^"]*"`,
		`(?i)<iframe.*?>.*?</iframe>`,
		`(?i)<img.*?src=['"]data:`,
	}

	// Compile and match each pattern
	for _, pattern := range patterns {
		re := regexp.MustCompile(pattern)
		if re.MatchString(input) {
			return true
		}
	}

	return false
}

// SanitizeXSS removes common XSS attack vectors from the input string
func SanitizeXSS(input string) string {
	// Patterns to check for common XSS contexts:
	// - <script> tags
	// - 'javascript:' pseudo-protocol
	// - HTML event handlers (e.g., onclick)
	// - <iframe> tags
	// - <img> tags with src='data:'
	patterns := []string{
		`(?i)<script.*?>.*?</script>`,
		`(?i)javascript:`,
		`(?i)on\w+="[^"]*"`,
		`(?i)<iframe.*?>.*?</iframe>`,
		`(?i)<img.*?src=['"]data:`,
	}

	// Replace each pattern with an empty string
	for _, pattern := range patterns {
		re := regexp.MustCompile(pattern)
		input = re.ReplaceAllString(input, "")
	}

	return input
}

// ContainsSQLi checks if the input string contains common SQL injection attack vectors
func ContainsSQLi(input string) bool {
	// Patterns to check for common SQL injection contexts:
	// - SQL keywords (e.g., SELECT, INSERT, UPDATE, DELETE)
	// - Comments (e.g., --, /* */)
	// - UNION operator
	// - OR operator
	// - AND operator
	patterns := []string{
		`(?i)\bSELECT\b`,
		`(?i)\bINSERT\b`,
		`(?i)\bUPDATE\b`,
		`(?i)\bDELETE\b`,
		`(?i)--`,
		`(?i)/\*.*?\*/`,
		`(?i)\bUNION\b`,
		`(?i)\bOR\b`,
		`(?i)\bAND\b`,
	}

	// Compile and match each pattern
	for _, pattern := range patterns {
		re := regexp.MustCompile(pattern)
		if re.MatchString(input) {
			return true
		}
	}

	return false
}

// SanitizeSQLi removes common SQL injection attack vectors from the input string
func SanitizeSQLi(input string) string {
	// Patterns to check for common SQL injection contexts:
	// - SQL keywords (e.g., SELECT, INSERT, UPDATE, DELETE)
	// - Comments (e.g., --, /* */)
	// - UNION operator
	// - OR operator
	// - AND operator
	patterns := []string{
		`(?i)\bSELECT\b`,
		`(?i)\bINSERT\b`,
		`(?i)\bUPDATE\b`,
		`(?i)\bDELETE\b`,
		`(?i)--`,
		`(?i)/\*.*?\*/`,
		`(?i)\bUNION\b`,
		`(?i)\bOR\b`,
		`(?i)\bAND\b`,
	}

	// Replace each pattern with an empty string
	for _, pattern := range patterns {
		re := regexp.MustCompile(pattern)
		input = re.ReplaceAllString(input, "")
	}

	return input
}
