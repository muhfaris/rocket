package utils

import "regexp"

func ContainsSpaceOrSpecialChar(s string) bool {
	// Define the regex pattern to match spaces or special characters
	pattern := `[ !@#$%^&*(),.?":{}|<>]`

	// Compile the regex
	re := regexp.MustCompile(pattern)

	// Check if the string contains a match
	return re.MatchString(s)
}
