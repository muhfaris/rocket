package utils

import (
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
