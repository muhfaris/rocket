package libcase

import (
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Enum untuk format penulisan
type CaseFormat string

const (
	CamelCase  CaseFormat = "camelCase"
	SnakeCase  CaseFormat = "snake_case"
	KebabCase  CaseFormat = "kebab-case"
	PascalCase CaseFormat = "PascalCase"
	Unknown    CaseFormat = "unknown"
)

// Fungsi untuk mendeteksi format penulisan path parameter
func Format(param string) (CaseFormat, string) {
	var (
		camelCaseRegex  = regexp.MustCompile(`^[a-z]+(?:[A-Z][a-z]+)*$`)
		snakeCaseRegex  = regexp.MustCompile(`^[a-z]+(?:_[a-z]+)*$`)
		kebabCaseRegex  = regexp.MustCompile(`^[a-z]+(?:-[a-z]+)*$`)
		pascalCaseRegex = regexp.MustCompile(`^[A-Z][a-z]+(?:[A-Z][a-z]+)*$`)
	)

	switch {
	case camelCaseRegex.MatchString(param):
		return CamelCase, toPascalCase(param)

	case snakeCaseRegex.MatchString(param):
		return SnakeCase, toPascalCase(param)

	case kebabCaseRegex.MatchString(param):
		return KebabCase, toPascalCase(param)

	case pascalCaseRegex.MatchString(param):
		return PascalCase, param

	default:
		return Unknown, param
	}
}

// Fungsi untuk mengonversi sebuah string ke PascalCase
func toPascalCase(param string) string {
	// Buat regex untuk memisahkan berdasarkan -, _, atau spasi
	re := regexp.MustCompile(`[_\-\s]+`)
	words := re.Split(param, -1)

	// Ubah huruf pertama setiap kata menjadi huruf besar
	for i, word := range words {
		if word == "id" {
			words[i] = "ID"
			continue
		}

		word := cases.Title(language.English).String(word)
		words[i] = word
	}

	// Gabungkan kembali menjadi satu string
	return strings.Join(words, "")
}

// Function to convert a string to snake_case
func ToSnakeCase(str string) string {
	var sb strings.Builder

	for i, r := range str {
		if unicode.IsUpper(r) {
			// Add an underscore before the uppercase letter if it's not the first character
			if i != 0 {
				sb.WriteRune('_')
			}
			// Convert the uppercase letter to lowercase
			sb.WriteRune(unicode.ToLower(r))
		} else {
			sb.WriteRune(r)
		}
	}

	return sb.String()
}

// ToTitleCase converts a string to title case (e.g., "hello" becomes "Hello").
func ToTitleCase(input string) string {
	if len(input) == 0 {
		return input
	}
	input = strings.ToLower(input)
	return strings.ToUpper(string(input[0])) + strings.ToLower(input[1:])
}
