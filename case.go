package inflect

import (
	"strings"
	"unicode"
)

// Underscore converts a string to snake_case.
//
// It handles PascalCase, camelCase, kebab-case, and mixed inputs.
// Consecutive uppercase letters (like "HTTP") are kept together as one word.
//
// Examples:
//   - Underscore("HelloWorld") returns "hello_world"
//   - Underscore("hello-world") returns "hello_world"
//   - Underscore("HTTPServer") returns "http_server"
//   - Underscore("getHTTPResponse") returns "get_http_response"
//   - Underscore("already_snake") returns "already_snake"
func Underscore(s string) string {
	return convertCase(s, '_')
}

// SnakeCase is an alias for Underscore.
// It converts a string to snake_case.
//
// Examples:
//   - SnakeCase("HelloWorld") returns "hello_world"
//   - SnakeCase("HTTPServer") returns "http_server"
func SnakeCase(s string) string {
	return Underscore(s)
}

// Dasherize converts a string to kebab-case.
//
// It handles PascalCase, camelCase, snake_case, and mixed inputs.
// Consecutive uppercase letters (like "HTTP") are kept together as one word.
//
// Examples:
//   - Dasherize("HelloWorld") returns "hello-world"
//   - Dasherize("hello_world") returns "hello-world"
//   - Dasherize("HTTPServer") returns "http-server"
//   - Dasherize("getHTTPResponse") returns "get-http-response"
//   - Dasherize("already-kebab") returns "already-kebab"
func Dasherize(s string) string {
	return convertCase(s, '-')
}

// KebabCase is an alias for Dasherize.
// It converts a string to kebab-case.
//
// Examples:
//   - KebabCase("HelloWorld") returns "hello-world"
//   - KebabCase("HTTPServer") returns "http-server"
func KebabCase(s string) string {
	return Dasherize(s)
}

// PascalCase converts a string to PascalCase.
//
// It handles snake_case, kebab-case, and mixed inputs.
// Each word's first letter is capitalized, with no separators.
//
// Examples:
//   - PascalCase("hello_world") returns "HelloWorld"
//   - PascalCase("hello-world") returns "HelloWorld"
//   - PascalCase("hello world") returns "HelloWorld"
//   - PascalCase("HTTP_SERVER") returns "HttpServer"
func PascalCase(s string) string {
	return toCamelOrPascal(s, true)
}

// TitleCase is an alias for PascalCase.
// It converts a string to PascalCase (also known as TitleCase in some contexts).
//
// Note: This is different from Titleize which preserves word separators.
//
// Examples:
//   - TitleCase("hello_world") returns "HelloWorld"
//   - TitleCase("hello-world") returns "HelloWorld"
func TitleCase(s string) string {
	return PascalCase(s)
}

// CamelCase converts a string to camelCase.
//
// It handles snake_case, kebab-case, and mixed inputs.
// First word is lowercase, subsequent words are capitalized, with no separators.
//
// Examples:
//   - CamelCase("hello_world") returns "helloWorld"
//   - CamelCase("hello-world") returns "helloWorld"
//   - CamelCase("hello world") returns "helloWorld"
//   - CamelCase("HTTP_SERVER") returns "httpServer"
func CamelCase(s string) string {
	return toCamelOrPascal(s, false)
}

// charType represents the type of a character for case conversion.
type charType byte

const (
	charNone      charType = 0
	charLower     charType = 'l'
	charUpper     charType = 'u'
	charDigit     charType = 'd'
	charSeparator charType = 's'
)

// isSeparator checks if a rune is a word separator.
func isSeparator(r rune) bool {
	return r == '_' || r == '-' || unicode.IsSpace(r)
}

// convertCase converts a string to either snake_case or kebab-case.
// The separator parameter determines which format to use ('_' or '-').
func convertCase(s string, separator rune) string {
	if s == "" {
		return s
	}

	var result strings.Builder
	result.Grow(len(s) + 10)

	runes := []rune(s)
	var prevType charType

	for i, r := range runes {
		prevType = processConvertRune(r, i, runes, prevType, separator, &result)
	}

	return strings.Trim(result.String(), string(separator))
}

// processConvertRune processes a single rune for case conversion.
func processConvertRune(r rune, i int, runes []rune, prevType charType, separator rune, result *strings.Builder) charType {
	if isSeparator(r) {
		if result.Len() > 0 {
			return charSeparator
		}
		return prevType
	}

	switch {
	case unicode.IsUpper(r):
		return handleUpperRune(r, i, runes, prevType, separator, result)
	case unicode.IsLower(r):
		return handleLowerRune(r, prevType, separator, result)
	case unicode.IsDigit(r):
		return handleDigitRune(r, prevType, separator, result)
	}
	return prevType
}

// handleUpperRune handles an uppercase rune in case conversion.
func handleUpperRune(r rune, i int, runes []rune, prevType charType, separator rune, result *strings.Builder) charType {
	if result.Len() > 0 {
		needsSep := false
		switch prevType {
		case charLower:
			needsSep = true
		case charUpper:
			needsSep = i+1 < len(runes) && unicode.IsLower(runes[i+1])
		case charDigit:
			needsSep = true
		case charSeparator:
			needsSep = true
		}
		if needsSep {
			result.WriteRune(separator)
		}
	}
	result.WriteRune(unicode.ToLower(r))
	return charUpper
}

// handleLowerRune handles a lowercase rune in case conversion.
func handleLowerRune(r rune, prevType charType, separator rune, result *strings.Builder) charType {
	if result.Len() > 0 && (prevType == charSeparator || prevType == charDigit) {
		result.WriteRune(separator)
	}
	result.WriteRune(r)
	return charLower
}

// handleDigitRune handles a digit rune in case conversion.
func handleDigitRune(r rune, prevType charType, separator rune, result *strings.Builder) charType {
	if result.Len() > 0 {
		if prevType == charSeparator || prevType == charLower || prevType == charUpper {
			result.WriteRune(separator)
		}
	}
	result.WriteRune(r)
	return charDigit
}

// toCamelOrPascal converts a string to camelCase or PascalCase.
// If pascal is true, the first letter is capitalized; otherwise it's lowercase.
func toCamelOrPascal(s string, pascal bool) string {
	if s == "" {
		return s
	}

	words := splitIntoWords(s)
	if len(words) == 0 {
		return ""
	}

	var result strings.Builder
	result.Grow(len(s))

	for i, word := range words {
		if word == "" {
			continue
		}
		word = strings.ToLower(word)
		if i == 0 && !pascal {
			result.WriteString(word)
		} else {
			result.WriteString(capitalizeWord(word))
		}
	}

	return result.String()
}

// splitIntoWords splits a string into words, handling various formats.
func splitIntoWords(s string) []string {
	var words []string
	var current strings.Builder

	runes := []rune(s)

	for i, r := range runes {
		words = processSplitRune(r, i, runes, &current, words)
	}

	if current.Len() > 0 {
		words = append(words, current.String())
	}

	return words
}

// processSplitRune processes a single rune for word splitting.
func processSplitRune(r rune, i int, runes []rune, current *strings.Builder, words []string) []string {
	if isSeparator(r) {
		if current.Len() > 0 {
			words = append(words, current.String())
			current.Reset()
		}
		return words
	}

	switch {
	case unicode.IsUpper(r):
		words = handleSplitUpper(r, i, runes, current, words)
	case unicode.IsDigit(r):
		words = handleSplitDigit(r, current, words)
	case unicode.IsLetter(r):
		words = handleSplitLetter(r, current, words)
	}
	return words
}

// handleSplitUpper handles an uppercase rune when splitting words.
func handleSplitUpper(r rune, i int, runes []rune, current *strings.Builder, words []string) []string {
	if current.Len() > 0 {
		lastRune := getLastRune(current)
		switch {
		case unicode.IsLower(lastRune):
			words = append(words, current.String())
			current.Reset()
		case unicode.IsUpper(lastRune):
			if i+1 < len(runes) && unicode.IsLower(runes[i+1]) {
				words = append(words, current.String())
				current.Reset()
			}
		case unicode.IsDigit(lastRune):
			words = append(words, current.String())
			current.Reset()
		}
	}
	current.WriteRune(r)
	return words
}

// handleSplitDigit handles a digit rune when splitting words.
func handleSplitDigit(r rune, current *strings.Builder, words []string) []string {
	if current.Len() > 0 {
		lastRune := getLastRune(current)
		if unicode.IsLetter(lastRune) {
			words = append(words, current.String())
			current.Reset()
		}
	}
	current.WriteRune(r)
	return words
}

// handleSplitLetter handles a lowercase letter when splitting words.
func handleSplitLetter(r rune, current *strings.Builder, words []string) []string {
	if current.Len() > 0 {
		lastRune := getLastRune(current)
		if unicode.IsDigit(lastRune) {
			words = append(words, current.String())
			current.Reset()
		}
	}
	current.WriteRune(r)
	return words
}

// getLastRune returns the last rune from a strings.Builder.
func getLastRune(b *strings.Builder) rune {
	s := b.String()
	runes := []rune(s)
	return runes[len(runes)-1]
}

// capitalizeWord capitalizes the first letter of a word.
func capitalizeWord(word string) string {
	if word == "" {
		return word
	}
	runes := []rune(word)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
