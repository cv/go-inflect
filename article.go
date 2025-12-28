package inflect

import (
	"regexp"
	"slices"
	"strings"
	"unicode"
)

// customAWords stores words that should be forced to use "a" instead of "an".
// The keys are lowercase versions of the words/patterns.
var customAWords = make(map[string]bool)

// customAnWords stores words that should be forced to use "an" instead of "a".
// The keys are lowercase versions of the words/patterns.
var customAnWords = make(map[string]bool)

// customAPatterns stores regex patterns that force "a" instead of "an".
// Patterns are matched against the lowercase first word.
var customAPatterns []*regexp.Regexp

// customAnPatterns stores regex patterns that force "an" instead of "a".
// Patterns are matched against the lowercase first word.
var customAnPatterns []*regexp.Regexp

// silentHWords contains words starting with silent 'h' that take "an".
var silentHWords = map[string]bool{
	"honest": true, "heir": true, "heiress": true, "heirloom": true,
	"honor": true, "honour": true, "hour": true, "hourly": true,
}

// lowercaseAbbrevs contains lowercase abbreviations pronounced letter-by-letter.
var lowercaseAbbrevs = map[string]bool{
	"mpeg": true, "jpeg": true, "gif": true, "sql": true, "html": true,
	"xml": true, "fbi": true, "cia": true, "nsa": true,
}

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
//
// Custom patterns defined via DefA(), DefAn(), DefAPattern(), and DefAnPattern()
// take precedence over default rules.
func An(word string) string {
	if word == "" {
		return ""
	}

	// Get the first word for pattern matching
	fields := strings.Fields(word)
	if len(fields) == 0 {
		return word
	}
	firstWord := fields[0]
	lowerFirst := strings.ToLower(firstWord)

	// Check custom "a" exact words first (highest priority)
	if customAWords[lowerFirst] {
		return "a " + word
	}

	// Check custom "an" exact words second
	if customAnWords[lowerFirst] {
		return "an " + word
	}

	// Check custom "a" regex patterns third
	for _, pat := range customAPatterns {
		if pat.MatchString(lowerFirst) {
			return "a " + word
		}
	}

	// Check custom "an" regex patterns fourth
	for _, pat := range customAnPatterns {
		if pat.MatchString(lowerFirst) {
			return "an " + word
		}
	}

	// Fall back to default rules
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
	for h := range silentHWords {
		if strings.HasPrefix(lower, h) {
			return true
		}
	}

	// Check for abbreviations/acronyms (all uppercase or known patterns)
	if isAbbreviation(firstWord) {
		return abbreviationNeedsAn(firstWord)
	}

	// Check for known lowercase abbreviations pronounced letter-by-letter
	if lowercaseAbbrevs[lower] {
		return abbreviationNeedsAn(strings.ToUpper(lower))
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
	return isAllUpper(word)
}

// abbreviationNeedsAn checks if an abbreviation should take "an".
// This depends on how the first letter is pronounced.
func abbreviationNeedsAn(abbrev string) bool {
	if abbrev == "" {
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
		if slices.Contains(youPatterns, prefix) {
			return true
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

// DefA defines a custom pattern that forces "a" instead of "an" for a word.
//
// The pattern is matched against the first word of the input (case-insensitive).
// Custom "a" patterns take precedence over custom "an" patterns and default rules.
//
// Examples:
//
//	DefA("ape")
//	An("ape") // returns "a ape" instead of "an ape"
//	An("Ape") // returns "a Ape" (case-insensitive matching)
func DefA(word string) {
	lower := strings.ToLower(word)
	customAWords[lower] = true
	// Remove from customAnWords if present to avoid conflicts
	delete(customAnWords, lower)
}

// DefAn defines a custom pattern that forces "an" instead of "a" for a word.
//
// The pattern is matched against the first word of the input (case-insensitive).
// Custom "an" patterns take precedence over default rules but not custom "a" patterns.
//
// Examples:
//
//	DefAn("hero")
//	An("hero") // returns "an hero" instead of "a hero"
//	An("Hero") // returns "an Hero" (case-insensitive matching)
func DefAn(word string) {
	lower := strings.ToLower(word)
	customAnWords[lower] = true
	// Remove from customAWords if present to avoid conflicts
	delete(customAWords, lower)
}

// UndefA removes a custom "a" pattern.
//
// Returns true if the pattern was removed, false if it didn't exist.
//
// Examples:
//
//	DefA("ape")
//	An("ape") // returns "a ape"
//	UndefA("ape")
//	An("ape") // returns "an ape" (default rule)
func UndefA(word string) bool {
	lower := strings.ToLower(word)
	if customAWords[lower] {
		delete(customAWords, lower)
		return true
	}
	return false
}

// UndefAn removes a custom "an" pattern.
//
// Returns true if the pattern was removed, false if it didn't exist.
//
// Examples:
//
//	DefAn("hero")
//	An("hero") // returns "an hero"
//	UndefAn("hero")
//	An("hero") // returns "a hero" (default rule)
func UndefAn(word string) bool {
	lower := strings.ToLower(word)
	if customAnWords[lower] {
		delete(customAnWords, lower)
		return true
	}
	return false
}

// DefAPattern defines a regex pattern that forces "a" instead of "an".
//
// The pattern is matched against the lowercase first word of the input.
// The pattern must be a valid Go regex. Patterns are matched with full-string
// matching (automatically anchored with ^ and $).
//
// Pattern priorities (highest to lowest):
//  1. Exact word matches (DefA)
//  2. Exact word matches (DefAn)
//  3. Regex patterns (DefAPattern)
//  4. Regex patterns (DefAnPattern)
//  5. Default rules
//
// Returns an error if the pattern is invalid.
//
// Examples:
//
//	DefAPattern("euro.*")
//	An("euro")     // returns "a euro"
//	An("european") // returns "a european"
//	An("eurozone") // returns "a eurozone"
func DefAPattern(pattern string) error {
	// Anchor the pattern to match the full word
	anchored := "^(?:" + pattern + ")$"
	re, err := regexp.Compile(anchored)
	if err != nil {
		return err
	}
	customAPatterns = append(customAPatterns, re)
	return nil
}

// DefAnPattern defines a regex pattern that forces "an" instead of "a".
//
// The pattern is matched against the lowercase first word of the input.
// The pattern must be a valid Go regex. Patterns are matched with full-string
// matching (automatically anchored with ^ and $).
//
// Pattern priorities (highest to lowest):
//  1. Exact word matches (DefA)
//  2. Exact word matches (DefAn)
//  3. Regex patterns (DefAPattern)
//  4. Regex patterns (DefAnPattern)
//  5. Default rules
//
// Returns an error if the pattern is invalid.
//
// Examples:
//
//	DefAnPattern("honor.*")
//	An("honor")     // returns "an honor"
//	An("honorable") // returns "an honorable"
//	An("honorary")  // returns "an honorary"
func DefAnPattern(pattern string) error {
	// Anchor the pattern to match the full word
	anchored := "^(?:" + pattern + ")$"
	re, err := regexp.Compile(anchored)
	if err != nil {
		return err
	}
	customAnPatterns = append(customAnPatterns, re)
	return nil
}

// UndefAPattern removes a regex pattern from the "a" patterns list.
//
// The pattern string must match exactly as it was defined (before anchoring).
// Returns true if the pattern was found and removed, false otherwise.
//
// Examples:
//
//	DefAPattern("euro.*")
//	An("european") // returns "a european"
//	UndefAPattern("euro.*")
//	An("european") // returns "an european" (default rule)
func UndefAPattern(pattern string) bool {
	anchored := "^(?:" + pattern + ")$"
	for i, re := range customAPatterns {
		if re.String() == anchored {
			customAPatterns = append(customAPatterns[:i], customAPatterns[i+1:]...)
			return true
		}
	}
	return false
}

// UndefAnPattern removes a regex pattern from the "an" patterns list.
//
// The pattern string must match exactly as it was defined (before anchoring).
// Returns true if the pattern was found and removed, false otherwise.
//
// Examples:
//
//	DefAnPattern("honor.*")
//	An("honorable") // returns "an honorable"
//	UndefAnPattern("honor.*")
//	An("honorable") // returns "a honorable" (default rule)
func UndefAnPattern(pattern string) bool {
	anchored := "^(?:" + pattern + ")$"
	for i, re := range customAnPatterns {
		if re.String() == anchored {
			customAnPatterns = append(customAnPatterns[:i], customAnPatterns[i+1:]...)
			return true
		}
	}
	return false
}

// DefAReset resets all custom a/an patterns to defaults (empty).
//
// This removes all custom patterns added via DefA(), DefAn(), DefAPattern(),
// and DefAnPattern().
//
// Example:
//
//	DefA("ape")
//	DefAn("hero")
//	DefAPattern("euro.*")
//	DefAnPattern("honor.*")
//	DefAReset()
//	An("ape")       // returns "an ape" (default rule)
//	An("hero")      // returns "a hero" (default rule)
//	An("european")  // returns "an european" (default rule)
//	An("honorable") // returns "a honorable" (default rule)
func DefAReset() {
	customAWords = make(map[string]bool)
	customAnWords = make(map[string]bool)
	customAPatterns = nil
	customAnPatterns = nil
}
