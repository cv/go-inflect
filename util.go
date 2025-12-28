package inflect

import (
	"maps"
	"strings"
	"unicode"
	"unicode/utf8"
)

// isAllUpper checks if all letters in a word are uppercase.
func isAllUpper(word string) bool {
	for _, r := range word {
		if unicode.IsLetter(r) && !unicode.IsUpper(r) {
			return false
		}
	}
	return true
}

// isProperName checks if a word is a proper name.
// A proper name is detected by having a capitalized first letter and not being
// all uppercase. Examples: "Jones", "Mary", "Smith".
func isProperName(word string) bool {
	if len(word) < 2 {
		return false
	}

	// Check if the first letter is uppercase (proper name indicator)
	firstRune, _ := utf8.DecodeRuneInString(word)
	if !unicode.IsUpper(firstRune) {
		return false
	}

	// If all uppercase, it's likely an acronym, not a proper name
	if isAllUpper(word) {
		return false
	}

	return true
}

// isProperNameEndingInS checks if a word is a proper name ending in 's'.
// A proper name is detected by having a capitalized first letter and not being
// all uppercase. Examples: "Jones", "Williams", "Hastings".
func isProperNameEndingInS(word string) bool {
	if !isProperName(word) {
		return false
	}

	// Check if the word ends in 's' or 'S'
	lastRune := rune(word[len(word)-1])
	return unicode.ToLower(lastRune) == 's'
}

// isVowel checks if a rune is a vowel.
func isVowel(r rune) bool {
	return strings.ContainsRune("aeiouAEIOU", r)
}

// matchCase adjusts the replacement to match the case pattern of the original.
func matchCase(original, replacement string) string {
	if original == "" || replacement == "" {
		return replacement
	}

	// Count letters to determine if it's a single-letter word
	letterCount := 0
	for _, r := range original {
		if unicode.IsLetter(r) {
			letterCount++
		}
	}

	// For single-letter words, just match the case of that letter
	if letterCount == 1 {
		firstRune, _ := utf8.DecodeRuneInString(original)
		if unicode.IsUpper(firstRune) {
			// Capitalize first letter of replacement
			runes := []rune(replacement)
			runes[0] = unicode.ToUpper(runes[0])
			return string(runes)
		}
		return replacement
	}

	// Check if original is all uppercase (multi-letter)
	if isAllUpper(original) {
		return strings.ToUpper(replacement)
	}

	// Check if original starts with uppercase
	firstRune, _ := utf8.DecodeRuneInString(original)
	if unicode.IsUpper(firstRune) {
		runes := []rune(replacement)
		runes[0] = unicode.ToUpper(runes[0])
		return string(runes)
	}

	return replacement
}

// matchSuffix returns the suffix in uppercase if the word is all uppercase.
func matchSuffix(word, suffix string) string {
	if isAllUpper(word) {
		return strings.ToUpper(suffix)
	}
	return suffix
}

// copyMap creates a shallow copy of a map.
func copyMap(src map[string]string) map[string]string {
	dst := make(map[string]string, len(src))
	maps.Copy(dst, src)
	return dst
}

// extractWhitespace returns the prefix (leading whitespace), trimmed word, and suffix (trailing whitespace).
// This safely handles the edge case where the word is all whitespace.
func extractWhitespace(word string) (prefix, trimmed, suffix string) {
	trimmed = strings.TrimSpace(word)
	if trimmed == "" {
		return word, "", ""
	}
	idx := strings.Index(word, trimmed)
	if idx < 0 {
		// Should never happen after TrimSpace, but handle gracefully
		return "", word, ""
	}
	prefix = word[:idx]
	suffix = word[idx+len(trimmed):]
	return
}

// IsPlural checks if a word appears to be in plural form.
//
// This function checks if the word is different from its singular form,
// indicating it's likely a plural. Note that this is heuristic and may
// not be accurate for all words, especially irregular forms.
//
// Examples:
//   - IsPlural("cats") returns true
//   - IsPlural("cat") returns false
//   - IsPlural("children") returns true
//   - IsPlural("child") returns false
//   - IsPlural("sheep") returns false (unchanged plurals are ambiguous)
func IsPlural(word string) bool {
	if word == "" {
		return false
	}

	lower := strings.ToLower(word)
	singularized := strings.ToLower(Singular(word))

	// If singularizing changes the word, it's likely plural
	return lower != singularized
}

// IsSingular checks if a word appears to be in singular form.
//
// This function returns true if the word is NOT plural.
// It's the logical inverse of IsPlural for most cases.
//
// Examples:
//   - IsSingular("cat") returns true
//   - IsSingular("cats") returns false
//   - IsSingular("child") returns true
//   - IsSingular("children") returns false
//   - IsSingular("sheep") returns true (unchanged plurals default to singular)
func IsSingular(word string) bool {
	if word == "" {
		return false
	}

	// A word is singular if it's not plural
	return !IsPlural(word)
}

// WordCount counts the number of words in a string.
//
// Words are separated by whitespace. This is a simple word count
// that doesn't account for punctuation or special cases.
//
// Examples:
//   - WordCount("hello world") returns 2
//   - WordCount("  one   two   three  ") returns 3
//   - WordCount("") returns 0
//   - WordCount("single") returns 1
func WordCount(text string) int {
	return len(strings.Fields(text))
}

// Capitalize capitalizes the first letter of a string.
//
// Examples:
//   - Capitalize("hello") returns "Hello"
//   - Capitalize("HELLO") returns "HELLO"
//   - Capitalize("") returns ""
//   - Capitalize("hello world") returns "Hello world"
func Capitalize(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// Titleize converts a string to title case.
//
// Each word's first letter is capitalized, rest are lowercased.
//
// Examples:
//   - Titleize("hello world") returns "Hello World"
//   - Titleize("HELLO WORLD") returns "Hello World"
//   - Titleize("hello-world") returns "Hello-World"
func Titleize(s string) string {
	if s == "" {
		return s
	}

	runes := []rune(strings.ToLower(s))
	capitalizeNext := true

	for i, r := range runes {
		if capitalizeNext && unicode.IsLetter(r) {
			runes[i] = unicode.ToUpper(r)
			capitalizeNext = false
		} else if unicode.IsSpace(r) || r == '-' {
			capitalizeNext = true
		}
	}

	return string(runes)
}
