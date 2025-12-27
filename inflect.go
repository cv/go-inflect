// Package inflect provides English language inflection utilities.
//
// It offers functions for pluralization, singularization, indefinite article
// selection (a/an), number-to-words conversion, ordinals, and more.
//
// This is a Go port of the Python inflect library.
package inflect

import (
	"strings"
	"unicode"
)

// An returns the word prefixed with the appropriate indefinite article ("a" or "an").
//
// The selection follows standard English rules:
//   - Use "an" before vowel sounds: "an apple", "an hour"
//   - Use "a" before consonant sounds: "a cat", "a university"
//
// Special cases handled:
//   - Silent 'h': "an honest person"
//   - Vowels with consonant sounds: "a Ukrainian", "a unanimous decision"
//   - Abbreviations: "a YAML file", "a JSON object"
func An(word string) string {
	if word == "" {
		return ""
	}

	if needsAn(word) {
		return "an " + word
	}
	return "a " + word
}

// needsAn determines if a word/phrase should be preceded by "an" (vs "a").
func needsAn(text string) bool {
	// Get the first word to analyze
	firstWord := strings.Fields(text)[0]
	lower := strings.ToLower(firstWord)

	// Check for silent 'h' words that take "an"
	silentH := []string{"honest", "heir", "heiress", "heirloom", "honor", "honour", "hour", "hourly"}
	for _, h := range silentH {
		if strings.HasPrefix(lower, h) {
			return true
		}
	}

	// Check for abbreviations/acronyms (all uppercase or known patterns)
	if isAbbreviation(firstWord) {
		return abbreviationNeedsAn(firstWord)
	}

	// Check for known lowercase abbreviations pronounced letter-by-letter
	lowercaseAbbrevs := []string{"mpeg", "jpeg", "gif", "sql", "html", "xml", "fbi", "cia", "nsa"}
	for _, abbr := range lowercaseAbbrevs {
		if lower == abbr {
			return abbreviationNeedsAn(strings.ToUpper(abbr))
		}
	}

	// Check for special vowel patterns that sound like consonants (take "a")
	// "uni-" sounds like "yoo-", "eu-" sounds like "yoo-", etc.
	consonantVowelPatterns := []string{
		"uni", "upon", "use", "used", "user", "using", "usual",
		"usu", "uran", "uret", "euro", "ewe", "onc", "one",
		"onet", // onetime
	}
	for _, pat := range consonantVowelPatterns {
		if strings.HasPrefix(lower, pat) {
			return false
		}
	}

	// Check for "U" followed by consonant then vowel pattern (sounds like "you")
	// e.g., Ugandan, Ukrainian, Unabomber, unanimous
	if len(lower) >= 2 && lower[0] == 'u' {
		if isConsonantYSound(lower) {
			return false
		}
	}

	// Special case: single letters
	if len(firstWord) == 1 {
		return isVowelSound(rune(lower[0]))
	}

	// Default: check if first letter is a vowel
	first := rune(lower[0])
	return isVowelSound(first)
}

// isAbbreviation checks if a word appears to be an abbreviation/acronym.
func isAbbreviation(word string) bool {
	if len(word) < 2 {
		return false
	}

	// Check if all letters are uppercase
	allUpper := true
	for _, r := range word {
		if unicode.IsLetter(r) && !unicode.IsUpper(r) {
			allUpper = false
			break
		}
	}
	return allUpper
}

// abbreviationNeedsAn checks if an abbreviation should take "an".
// This depends on how the first letter is pronounced.
func abbreviationNeedsAn(abbrev string) bool {
	if len(abbrev) == 0 {
		return false
	}

	// Letters whose names start with vowel sounds:
	// A (ay), E (ee), F (eff), H (aitch), I (eye), L (ell), M (em),
	// N (en), O (oh), R (ar), S (ess), X (ex)
	vowelSoundLetters := "AEFHILMNORSX"
	first := unicode.ToUpper(rune(abbrev[0]))
	return strings.ContainsRune(vowelSoundLetters, first)
}

// isConsonantYSound checks if a word starting with 'u' has a consonant "y" sound.
// e.g., "unanimous" starts with "yoo-" sound.
func isConsonantYSound(lower string) bool {
	if len(lower) < 2 {
		return false
	}

	// Explicit patterns that have the "you" sound
	youPatterns := []string{
		"uga", "ukr", "ula", "ule", "uli", "ulo", "ulu",
		"una", "uni", "uno", "unu", "ura", "ure", "uri", "uro", "uru",
		"usa", "use", "usi", "uso", "usu", "uta", "ute", "uti", "uto", "utu",
	}

	if len(lower) >= 3 {
		prefix := lower[:3]
		for _, pat := range youPatterns {
			if prefix == pat {
				return true
			}
		}
	}

	return false
}

// isVowelSound checks if a letter represents a vowel sound.
func isVowelSound(r rune) bool {
	return strings.ContainsRune("aeiou", unicode.ToLower(r))
}

// A is an alias for An - returns word prefixed with appropriate indefinite article.
func A(word string) string {
	return An(word)
}
