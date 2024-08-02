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
