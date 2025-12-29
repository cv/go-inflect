package inflect

import "strings"

// PastTense returns the simple past tense form of an English verb.
//
// Examples:
//   - PastTense("walk") returns "walked"
//   - PastTense("go") returns "went"
//   - PastTense("try") returns "tried"
//   - PastTense("stop") returns "stopped"
func PastTense(verb string) string {
	if verb == "" {
		return ""
	}

	lower := strings.ToLower(verb)

	// Check verbs with same past/participle first (most common)
	if past, ok := irregularVerbsSame[lower]; ok {
		return matchCase(verb, past)
	}

	// Check verbs with different past tense
	if past, ok := irregularPastTenseOnly[lower]; ok {
		return matchCase(verb, past)
	}

	// Apply regular rules
	return applyPastTenseRules(verb, lower)
}

// applyPastTenseRules applies regular past tense formation rules.
func applyPastTenseRules(verb, lower string) string {
	// Verbs ending in -e: add -d
	if strings.HasSuffix(lower, "e") {
		return verb + matchSuffix(verb, "d")
	}

	// Verbs ending in consonant + y: change y to -ied
	if strings.HasSuffix(lower, "y") && len(lower) > 1 {
		beforeY := lower[len(lower)-2]
		if !isVowel(rune(beforeY)) {
			return verb[:len(verb)-1] + matchSuffix(verb, "ied")
		}
	}

	// CVC pattern: double the final consonant and add -ed
	if shouldDoubleFinalConsonantForPast(lower) {
		lastChar := string(lower[len(lower)-1])
		return verb + matchSuffix(verb, lastChar+"ed")
	}

	// Default: add -ed
	return verb + matchSuffix(verb, "ed")
}

// shouldDoubleFinalConsonantForPast checks if the final consonant should be doubled
// when forming the past tense. This applies to short verbs with a CVC pattern.
func shouldDoubleFinalConsonantForPast(lower string) bool {
	if len(lower) < 2 {
		return false
	}

	// Only for short (one-syllable) words
	if countSyllables(lower) != 1 {
		return false
	}

	lastChar := rune(lower[len(lower)-1])
	secondLastChar := rune(lower[len(lower)-2])

	// Last char must be a consonant (not w, x, or y)
	if isVowel(lastChar) || lastChar == 'w' || lastChar == 'x' || lastChar == 'y' {
		return false
	}

	// Second-to-last must be a single vowel
	if !isVowel(secondLastChar) {
		return false
	}

	// Check that there's a consonant before the vowel (CVC pattern)
	if len(lower) >= 3 {
		thirdLastChar := rune(lower[len(lower)-3])
		if isVowel(thirdLastChar) {
			return false // VVC pattern, don't double
		}
	}

	return true
}
