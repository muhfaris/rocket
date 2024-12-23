package utils

import (
	"fmt"
	"regexp"
	"strings"
)

func ContainsSpaceOrSpecialChar(s string) bool {
	// Define the regex pattern to match spaces or special characters
	pattern := `[ !@#$%^&*(),.?":{}|<>]`

	// Compile the regex
	re := regexp.MustCompile(pattern)

	// Check if the string contains a match
	return re.MatchString(s)
}

// ConvertBracesToColon converts placeholders in the format {id} to :id
func ConvertBracesToColon(input string) string {
	// Replace { with :
	result := strings.Replace(input, "{", ":", -1)
	// Remove }
	result = strings.Replace(result, "}", "", -1)
	return result
}

// SanitizeString removes all special characters and spaces from a string.
func SanitizeString(input string) string {
	// Define a regular expression to match special characters and spaces
	regex := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	// Replace matches with an empty string
	sanitized := regex.ReplaceAllString(input, "")
	return sanitized
}

// ValidateImportPath checks if a string is a valid Go package import path.
func ValidateImportPath(path string) error {
	// Regular expression to validate typical Go package import paths.
	// Adjust as necessary for your specific requirements.
	regex := regexp.MustCompile(`^[a-zA-Z0-9\-_.~]+(\.[a-zA-Z]+)?(/[a-zA-Z0-9\-_.~]+)+$`)

	if !regex.MatchString(path) {
		return fmt.Errorf("invalid import path: %s", path)
	}

	// Optional: Check for specific domains (e.g., github.com).
	if !(regexp.MustCompile(`^([a-z0-9]+(\.[a-z0-9]+)+/).*`).MatchString(path)) {
		return fmt.Errorf("import path must include a domain: %s", path)
	}

	return nil
}
