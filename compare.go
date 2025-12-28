package inflect

import "strings"

// Compare result constants.
const (
	compareEq           = "eq"
	compareSingToPlural = "s:p"
	comparePluralToSing = "p:s"
	comparePluralPlural = "p:p"
)

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
	return defaultEngine.Compare(word1, word2)
}

// Compare compares two words for singular/plural equality using this engine's settings.
//
// It returns:
//   - "eq" if the words are equal (case-insensitive)
//   - "s:p" if word1 is singular and word2 is its plural form
//   - "p:s" if word1 is plural and word2 is its singular form
//   - "p:p" if both words are different plural forms of the same word
//   - "" if the words are not related
//
// Examples:
//   - e.Compare("cat", "cat") returns "eq"
//   - e.Compare("cat", "cats") returns "s:p"
//   - e.Compare("cats", "cat") returns "p:s"
//   - e.Compare("indexes", "indices") returns "p:p"
//   - e.Compare("cat", "dog") returns ""
func (e *Engine) Compare(word1, word2 string) string {
	// Handle empty strings
	if word1 == "" || word2 == "" {
		if word1 == "" && word2 == "" {
			return compareEq
		}
		return ""
	}

	// Normalize for comparison
	lower1 := strings.ToLower(word1)
	lower2 := strings.ToLower(word2)

	// Same word
	if lower1 == lower2 {
		return compareEq
	}

	// Check if word1 is singular and word2 is its plural
	// Use Plural() for verification since Singular() has edge cases
	if strings.ToLower(e.Plural(word1)) == lower2 {
		return compareSingToPlural
	}

	// Check if word2 is singular and word1 is its plural
	if strings.ToLower(e.Plural(word2)) == lower1 {
		return comparePluralToSing
	}

	// Check if both are different plural forms of the same singular word
	singular1 := strings.ToLower(e.Singular(word1))
	singular2 := strings.ToLower(e.Singular(word2))

	if singular1 == singular2 {
		// Verify both are actually plurals (different from their singular form)
		// by checking that pluralizing the singular gives us something related
		pluralOfSingular := strings.ToLower(e.Plural(singular1))
		// If both words singularize to the same thing, and that singular
		// can be pluralized, they're both plural forms
		if lower1 != singular1 && lower2 != singular2 {
			// Additional verification: ensure the singular is valid
			if pluralOfSingular == lower1 || pluralOfSingular == lower2 {
				return comparePluralPlural
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
	return defaultEngine.CompareNouns(noun1, noun2)
}

// CompareNouns compares two nouns for singular/plural equality using this engine's settings.
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
//   - e.CompareNouns("cat", "cats") returns "s:p"
//   - e.CompareNouns("mice", "mouse") returns "p:s"
func (e *Engine) CompareNouns(noun1, noun2 string) string {
	return e.Compare(noun1, noun2)
}

// CompareVerbs compares two verbs for singular/plural equality.
//
// This compares verbs using verb pluralization rules (3rd person singular vs base form).
//
// It returns:
//   - "eq" if the verbs are equal (case-insensitive)
//   - "s:p" if verb1 is singular (3rd person) and verb2 is its plural (base form)
//   - "p:s" if verb1 is plural (base form) and verb2 is its singular (3rd person)
//   - "" if the verbs are not related
//
// Examples:
//   - CompareVerbs("runs", "run") returns "s:p" (3rd person to base)
//   - CompareVerbs("run", "runs") returns "p:s" (base to 3rd person)
//   - CompareVerbs("is", "are") returns "s:p"
//   - CompareVerbs("has", "have") returns "s:p"
func CompareVerbs(verb1, verb2 string) string {
	return defaultEngine.CompareVerbs(verb1, verb2)
}

// CompareVerbs compares two verbs for singular/plural equality using this engine's settings.
//
// This compares verbs using verb pluralization rules (3rd person singular vs base form).
//
// It returns:
//   - "eq" if the verbs are equal (case-insensitive)
//   - "s:p" if verb1 is singular (3rd person) and verb2 is its plural (base form)
//   - "p:s" if verb1 is plural (base form) and verb2 is its singular (3rd person)
//   - "" if the verbs are not related
//
// Examples:
//   - e.CompareVerbs("runs", "run") returns "s:p" (3rd person to base)
//   - e.CompareVerbs("run", "runs") returns "p:s" (base to 3rd person)
//   - e.CompareVerbs("is", "are") returns "s:p"
//   - e.CompareVerbs("has", "have") returns "s:p"
func (e *Engine) CompareVerbs(verb1, verb2 string) string {
	// Handle empty strings
	if verb1 == "" || verb2 == "" {
		if verb1 == "" && verb2 == "" {
			return compareEq
		}
		return ""
	}

	// Normalize for comparison
	lower1 := strings.ToLower(verb1)
	lower2 := strings.ToLower(verb2)

	// Same verb
	if lower1 == lower2 {
		return compareEq
	}

	// Check if verb1 is singular (3rd person) and verb2 is its plural (base form)
	// PluralVerb converts 3rd person singular to base form
	if strings.ToLower(e.PluralVerb(verb1)) == lower2 {
		return compareSingToPlural
	}

	// Check if verb2 is singular and verb1 is its plural
	if strings.ToLower(e.PluralVerb(verb2)) == lower1 {
		return comparePluralToSing
	}

	return ""
}

// CompareAdjs compares two adjectives for singular/plural equality.
//
// This compares adjectives using adjective pluralization rules (demonstratives, articles, possessives).
//
// It returns:
//   - "eq" if the adjectives are equal (case-insensitive)
//   - "s:p" if adj1 is singular and adj2 is its plural form
//   - "p:s" if adj1 is plural and adj2 is its singular form
//   - "" if the adjectives are not related
//
// Examples:
//   - CompareAdjs("this", "these") returns "s:p"
//   - CompareAdjs("that", "those") returns "s:p"
//   - CompareAdjs("these", "this") returns "p:s"
//   - CompareAdjs("a", "some") returns "s:p"
func CompareAdjs(adj1, adj2 string) string {
	return defaultEngine.CompareAdjs(adj1, adj2)
}

// CompareAdjs compares two adjectives for singular/plural equality using this engine's settings.
//
// This compares adjectives using adjective pluralization rules (demonstratives, articles, possessives).
//
// It returns:
//   - "eq" if the adjectives are equal (case-insensitive)
//   - "s:p" if adj1 is singular and adj2 is its plural form
//   - "p:s" if adj1 is plural and adj2 is its singular form
//   - "" if the adjectives are not related
//
// Examples:
//   - e.CompareAdjs("this", "these") returns "s:p"
//   - e.CompareAdjs("that", "those") returns "s:p"
//   - e.CompareAdjs("these", "this") returns "p:s"
//   - e.CompareAdjs("a", "some") returns "s:p"
func (e *Engine) CompareAdjs(adj1, adj2 string) string {
	// Handle empty strings
	if adj1 == "" || adj2 == "" {
		if adj1 == "" && adj2 == "" {
			return compareEq
		}
		return ""
	}

	// Normalize for comparison
	lower1 := strings.ToLower(adj1)
	lower2 := strings.ToLower(adj2)

	// Same adjective
	if lower1 == lower2 {
		return compareEq
	}

	// Check if adj1 is singular and adj2 is its plural form
	if strings.ToLower(e.PluralAdj(adj1)) == lower2 {
		return compareSingToPlural
	}

	// Check if adj2 is singular and adj1 is its plural form
	if strings.ToLower(e.PluralAdj(adj2)) == lower1 {
		return comparePluralToSing
	}

	return ""
}
