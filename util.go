package inflect

import (
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
	for k, v := range src {
		dst[k] = v
	}
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
