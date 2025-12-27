package inflect

import "strings"

// Compare compares two words for singular/plural equality.
//
// It returns:
//   - "eq" if the words are equal (case-insensitive)
//   - "s:p" if word1 is singular and word2 is its plural form
//   - "p:s" if word1 is plural and word2 is its singular form
//   - "p:p" if both words are different plural forms of the same word
//   - "" if the words are not related
//
// Examples:
//   - Compare("cat", "cat") returns "eq"
//   - Compare("cat", "cats") returns "s:p"
//   - Compare("cats", "cat") returns "p:s"
//   - Compare("indexes", "indices") returns "p:p"
//   - Compare("cat", "dog") returns ""
func Compare(word1, word2 string) string {
	// Handle empty strings
	if word1 == "" || word2 == "" {
		if word1 == "" && word2 == "" {
			return "eq"
		}
		return ""
	}

	// Normalize for comparison
	lower1 := strings.ToLower(word1)
	lower2 := strings.ToLower(word2)

	// Same word
	if lower1 == lower2 {
		return "eq"
	}

	// Check if word1 is singular and word2 is its plural
	// Use Plural() for verification since Singular() has edge cases
	if strings.ToLower(Plural(word1)) == lower2 {
		return "s:p"
	}

	// Check if word2 is singular and word1 is its plural
	if strings.ToLower(Plural(word2)) == lower1 {
		return "p:s"
	}

	// Check if both are different plural forms of the same singular word
	singular1 := strings.ToLower(Singular(word1))
	singular2 := strings.ToLower(Singular(word2))

	if singular1 == singular2 {
		// Verify both are actually plurals (different from their singular form)
		// by checking that pluralizing the singular gives us something related
		pluralOfSingular := strings.ToLower(Plural(singular1))
		// If both words singularize to the same thing, and that singular
		// can be pluralized, they're both plural forms
		if lower1 != singular1 && lower2 != singular2 {
			// Additional verification: ensure the singular is valid
			if pluralOfSingular == lower1 || pluralOfSingular == lower2 {
				return "p:p"
			}
		}
	}

	return ""
}

// CompareNouns compares two nouns for singular/plural equality.
//
// This is an alias for Compare that makes the intent explicit when working
// specifically with nouns.
//
// It returns:
//   - "eq" if the nouns are equal (case-insensitive)
//   - "s:p" if noun1 is singular and noun2 is its plural form
//   - "p:s" if noun1 is plural and noun2 is its singular form
//   - "p:p" if both nouns are different plural forms of the same word
//   - "" if the nouns are not related
//
// Examples:
//   - CompareNouns("cat", "cats") returns "s:p"
//   - CompareNouns("mice", "mouse") returns "p:s"
func CompareNouns(noun1, noun2 string) string {
	return Compare(noun1, noun2)
}

// CompareVerbs compares two verbs for conjugation equality.
//
// NOTE: This is a placeholder stub for future implementation.
// Verb conjugation comparison is not yet implemented.
//
// Currently always returns an empty string.
//
// Future implementation will return:
//   - "eq" if the verbs are equal (case-insensitive)
//   - Comparison codes for different conjugation relationships
//   - "" if the verbs are not related
func CompareVerbs(_, _ string) string {
	// TODO: Implement verb conjugation comparison
	return ""
}

// CompareAdjs compares two adjectives for comparative/superlative equality.
//
// NOTE: This is a placeholder stub for future implementation.
// Adjective comparison is not yet implemented.
//
// Currently always returns an empty string.
//
// Future implementation will return:
//   - "eq" if the adjectives are equal (case-insensitive)
//   - Comparison codes for different adjective relationships (e.g., base/comparative/superlative)
//   - "" if the adjectives are not related
func CompareAdjs(_, _ string) string {
	// TODO: Implement adjective comparison
	return ""
}
