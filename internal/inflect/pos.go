package inflect

import "strings"

// PluralNoun returns the plural form of an English noun or pronoun.
//
// This function handles:
//   - Pronouns in nominative case: "I" -> "we", "he"/"she"/"it" -> "they"
//   - Pronouns in accusative case: "me" -> "us", "him"/"her" -> "them"
//   - Possessive pronouns: "my" -> "our", "mine" -> "ours", "his"/"her"/"its" -> "their"
//   - Reflexive pronouns: "myself" -> "ourselves", "himself"/"herself"/"itself" -> "themselves"
//   - Regular nouns: defers to Plural()
//
// If count is provided and equals 1 or -1, returns the singular form.
// If count is not 1, returns the plural form.
// If no count is provided, returns the plural form.
//
// Examples:
//
//	PluralNoun("I") returns "we"
//	PluralNoun("me") returns "us"
//	PluralNoun("my") returns "our"
//	PluralNoun("cat") returns "cats"
//	PluralNoun("cat", 1) returns "cat"
//	PluralNoun("cat", 2) returns "cats"
func PluralNoun(word string, count ...int) string {
	return defaultEngine.PluralNoun(word, count...)
}

// PluralNoun returns the plural form of an English noun or pronoun.
//
// This method handles:
//   - Pronouns in nominative case: "I" -> "we", "he"/"she"/"it" -> "they"
//   - Pronouns in accusative case: "me" -> "us", "him"/"her" -> "them"
//   - Possessive pronouns: "my" -> "our", "mine" -> "ours", "his"/"her"/"its" -> "their"
//   - Reflexive pronouns: "myself" -> "ourselves", "himself"/"herself"/"itself" -> "themselves"
//   - Regular nouns: defers to Plural()
//
// If count is provided and equals 1 or -1, returns the singular form.
// If count is not 1, returns the plural form.
// If no count is provided, returns the plural form.
//
// Examples:
//
//	e.PluralNoun("I") returns "we"
//	e.PluralNoun("me") returns "us"
//	e.PluralNoun("my") returns "our"
//	e.PluralNoun("cat") returns "cats"
//	e.PluralNoun("cat", 1) returns "cat"
//	e.PluralNoun("cat", 2) returns "cats"
func (e *Engine) PluralNoun(word string, count ...int) string {
	if word == "" {
		return ""
	}

	// Preserve leading/trailing whitespace
	prefix, trimmed, suffix := extractWhitespace(word)
	if trimmed == "" {
		return word
	}

	// Handle count parameter
	if len(count) > 0 && (count[0] == 1 || count[0] == -1) {
		return word // Return singular form as-is
	}

	lower := strings.ToLower(trimmed)

	// Check for pronouns first
	if plural, ok := allPronounsToPlural[lower]; ok {
		return prefix + matchCase(trimmed, plural) + suffix
	}

	// Fall back to regular Plural() for nouns
	return prefix + e.Plural(trimmed) + suffix
}

// PluralVerb returns the plural form of an English verb.
//
// This function handles:
//   - Auxiliary verbs: "is" -> "are", "was" -> "were", "has" -> "have"
//   - Contractions: "isn't" -> "aren't", "doesn't" -> "don't", "hasn't" -> "haven't"
//   - Modal verbs (unchanged): "can", "could", "may", "might", "must", "shall", "should", "will", "would"
//   - Regular verbs in third person singular: removes -s/-es suffix
//
// If count is provided and equals 1 or -1, returns the singular form.
// If count is not 1, returns the plural form.
// If no count is provided, returns the plural form.
//
// Examples:
//
//	PluralVerb("is") returns "are"
//	PluralVerb("was") returns "were"
//	PluralVerb("has") returns "have"
//	PluralVerb("doesn't") returns "don't"
//	PluralVerb("runs") returns "run"
//	PluralVerb("is", 1) returns "is"
//	PluralVerb("is", 2) returns "are"
func PluralVerb(word string, count ...int) string {
	return defaultEngine.PluralVerb(word, count...)
}

// PluralVerb returns the plural form of an English verb.
//
// This method handles:
//   - Auxiliary verbs: "is" -> "are", "was" -> "were", "has" -> "have"
//   - Contractions: "isn't" -> "aren't", "doesn't" -> "don't", "hasn't" -> "haven't"
//   - Modal verbs (unchanged): "can", "could", "may", "might", "must", "shall", "should", "will", "would"
//   - Regular verbs in third person singular: removes -s/-es suffix
//
// If count is provided and equals 1 or -1, returns the singular form.
// If count is not 1, returns the plural form.
// If no count is provided, returns the plural form.
//
// Examples:
//
//	e.PluralVerb("is") returns "are"
//	e.PluralVerb("was") returns "were"
//	e.PluralVerb("has") returns "have"
//	e.PluralVerb("doesn't") returns "don't"
//	e.PluralVerb("runs") returns "run"
//	e.PluralVerb("is", 1) returns "is"
//	e.PluralVerb("is", 2) returns "are"
func (e *Engine) PluralVerb(word string, count ...int) string {
	if word == "" {
		return ""
	}

	// Preserve leading/trailing whitespace
	prefix, trimmed, suffix := extractWhitespace(word)
	if trimmed == "" {
		return word
	}

	// Handle count parameter - if singular count, return singular form
	if len(count) > 0 && (count[0] == 1 || count[0] == -1) {
		// Return appropriate singular form
		lower := strings.ToLower(trimmed)
		if singular, ok := verbPluralToSingular[lower]; ok {
			return prefix + matchCase(trimmed, singular) + suffix
		}
		return word
	}

	lower := strings.ToLower(trimmed)

	// Check for unchanged modal verbs
	if verbUnchanged[lower] {
		return word
	}

	// Check for known irregular verb mappings
	if plural, ok := verbSingularToPlural[lower]; ok {
		return prefix + matchCase(trimmed, plural) + suffix
	}

	// Check custom verb definitions
	e.mu.RLock()
	plural, ok := e.customVerbs[lower]
	e.mu.RUnlock()
	if ok {
		return prefix + matchCase(trimmed, plural) + suffix
	}

	// For regular verbs in third person singular (ends in -s/-es),
	// convert to base form (which is the plural form)

	// Handle -ies -> -y first (tries -> try, flies -> fly)
	if strings.HasSuffix(lower, "ies") && len(lower) > 3 {
		return prefix + trimmed[:len(trimmed)-3] + matchSuffix(trimmed, "y") + suffix
	}

	// Handle -es after sibilants: -sses, -shes, -ches, -xes, -zes
	if strings.HasSuffix(lower, "es") && len(lower) > 2 {
		base := lower[:len(lower)-2]
		if strings.HasSuffix(base, "ss") ||
			strings.HasSuffix(base, "sh") ||
			strings.HasSuffix(base, "ch") ||
			strings.HasSuffix(base, "x") ||
			strings.HasSuffix(base, "z") {
			return prefix + trimmed[:len(trimmed)-2] + suffix
		}
		// -oes -> -o (goes -> go) - but goes is in the irregular list
		if strings.HasSuffix(base, "o") {
			return prefix + trimmed[:len(trimmed)-2] + suffix
		}
		// For other -es endings like "sees", just remove -s (not -es)
		// "sees" -> "see", "flees" -> "flee"
	}

	// Regular -s ending: just remove the -s
	if strings.HasSuffix(lower, "s") && len(lower) > 1 {
		// Don't remove -s from words ending in -ss
		if !strings.HasSuffix(lower, "ss") {
			// Regular third person singular: runs -> run, sees -> see
			return prefix + trimmed[:len(trimmed)-1] + suffix
		}
	}

	// Return unchanged if no rule matches (already plural or base form)
	return word
}

// PluralAdj returns the plural form of an English adjective.
//
// This function handles:
//   - Demonstrative adjectives: "this" -> "these", "that" -> "those"
//   - Indefinite articles: "a" -> "some", "an" -> "some"
//   - Possessive adjectives: "my" -> "our", "his"/"her"/"its" -> "their"
//
// If count is provided and equals 1 or -1, returns the singular form.
// If count is not 1, returns the plural form.
// If no count is provided, returns the plural form.
//
// Examples:
//
//	PluralAdj("this") returns "these"
//	PluralAdj("that") returns "those"
//	PluralAdj("a") returns "some"
//	PluralAdj("an") returns "some"
//	PluralAdj("my") returns "our"
//	PluralAdj("this", 1) returns "this"
//	PluralAdj("this", 2) returns "these"
func PluralAdj(word string, count ...int) string {
	return defaultEngine.PluralAdj(word, count...)
}

// PluralAdj returns the plural form of an English adjective.
//
// This method handles:
//   - Demonstrative adjectives: "this" -> "these", "that" -> "those"
//   - Indefinite articles: "a" -> "some", "an" -> "some"
//   - Possessive adjectives: "my" -> "our", "his"/"her"/"its" -> "their"
//
// If count is provided and equals 1 or -1, returns the singular form.
// If count is not 1, returns the plural form.
// If no count is provided, returns the plural form.
//
// Examples:
//
//	e.PluralAdj("this") returns "these"
//	e.PluralAdj("that") returns "those"
//	e.PluralAdj("a") returns "some"
//	e.PluralAdj("an") returns "some"
//	e.PluralAdj("my") returns "our"
//	e.PluralAdj("this", 1) returns "this"
//	e.PluralAdj("this", 2) returns "these"
func (e *Engine) PluralAdj(word string, count ...int) string {
	if word == "" {
		return ""
	}

	// Preserve leading/trailing whitespace
	prefix, trimmed, suffix := extractWhitespace(word)
	if trimmed == "" {
		return word
	}

	// Handle count parameter - if singular count, return singular form
	if len(count) > 0 && (count[0] == 1 || count[0] == -1) {
		// Return appropriate singular form
		lower := strings.ToLower(trimmed)
		if singular, ok := adjPluralToSingular[lower]; ok {
			return prefix + matchCase(trimmed, singular) + suffix
		}
		e.mu.RLock()
		g := e.gender
		e.mu.RUnlock()
		if genderMap, ok := adjPluralToSingularByGender[lower]; ok {
			if singular, ok := genderMap[g]; ok {
				return prefix + matchCase(trimmed, singular) + suffix
			}
		}
		return word
	}

	lower := strings.ToLower(trimmed)

	// Check for known adjective mappings
	if plural, ok := adjSingularToPlural[lower]; ok {
		return prefix + matchCase(trimmed, plural) + suffix
	}

	// Check custom adjective definitions
	e.mu.RLock()
	plural, ok := e.customAdjs[lower]
	e.mu.RUnlock()
	if ok {
		return prefix + matchCase(trimmed, plural) + suffix
	}

	// Most adjectives don't change between singular and plural in English
	return word
}

// SingularNoun returns the singular form of an English noun or pronoun.
//
// This function handles:
//   - Pronouns in nominative case: "we" -> "I", "they" -> he/she/it/they (depends on gender)
//   - Pronouns in accusative case: "us" -> "me", "them" -> him/her/it/them (depends on gender)
//   - Possessive pronouns: "our" -> "my", "ours" -> "mine", "their" -> his/her/its/their
//   - Reflexive pronouns: "ourselves" -> "myself", "themselves" -> himself/herself/itself/themself
//   - Regular nouns: defers to Singular()
//
// Third-person singular pronouns use the gender set by Gender():
//   - Gender("m"): masculine - "they" -> "he"
//   - Gender("f"): feminine - "they" -> "she"
//   - Gender("n"): neuter - "they" -> "it"
//   - Gender("t"): they (singular they) - "they" -> "they"
//
// If count is provided and equals 1 or -1, returns the singular form.
// If count is not 1, returns the plural form.
// If no count is provided, returns the singular form.
//
// Examples:
//
//	SingularNoun("we") returns "I"
//	SingularNoun("us") returns "me"
//	SingularNoun("our") returns "my"
//	SingularNoun("they") returns "it" (or he/she/they based on gender)
//	SingularNoun("cats") returns "cat"
//	SingularNoun("cats", 1) returns "cat"
//	SingularNoun("cats", 2) returns "cats"
func SingularNoun(word string, count ...int) string {
	return defaultEngine.SingularNoun(word, count...)
}

// SingularNoun returns the singular form of an English noun or pronoun.
//
// This method handles:
//   - Pronouns in nominative case: "we" -> "I", "they" -> he/she/it/they (depends on gender)
//   - Pronouns in accusative case: "us" -> "me", "them" -> him/her/it/them (depends on gender)
//   - Possessive pronouns: "our" -> "my", "ours" -> "mine", "their" -> his/her/its/their
//   - Reflexive pronouns: "ourselves" -> "myself", "themselves" -> himself/herself/itself/themself
//   - Regular nouns: defers to Singular()
//
// Third-person singular pronouns use the gender set by SetGender():
//   - SetGender("m"): masculine - "they" -> "he"
//   - SetGender("f"): feminine - "they" -> "she"
//   - SetGender("n"): neuter - "they" -> "it"
//   - SetGender("t"): they (singular they) - "they" -> "they"
//
// If count is provided and equals 1 or -1, returns the singular form.
// If count is not 1, returns the plural form.
// If no count is provided, returns the singular form.
//
// Examples:
//
//	e.SingularNoun("we") returns "I"
//	e.SingularNoun("us") returns "me"
//	e.SingularNoun("our") returns "my"
//	e.SingularNoun("they") returns "it" (or he/she/they based on gender)
//	e.SingularNoun("cats") returns "cat"
//	e.SingularNoun("cats", 1) returns "cat"
//	e.SingularNoun("cats", 2) returns "cats"
func (e *Engine) SingularNoun(word string, count ...int) string {
	if word == "" {
		return ""
	}

	// Preserve leading/trailing whitespace
	prefix, trimmed, suffix := extractWhitespace(word)
	if trimmed == "" {
		return word
	}

	// Handle count parameter - if plural count, return plural form
	if len(count) > 0 && count[0] != 1 && count[0] != -1 {
		return word // Return plural form as-is
	}

	lower := strings.ToLower(trimmed)

	// Read gender once with lock
	e.mu.RLock()
	g := e.gender
	e.mu.RUnlock()

	// Check for nominative pronouns
	if genderMap, ok := pronounNominativeSingularByGender[lower]; ok {
		if singular, ok := genderMap[g]; ok {
			return prefix + matchCase(trimmed, singular) + suffix
		}
	}

	// Check for accusative pronouns
	if genderMap, ok := pronounAccusativeSingularByGender[lower]; ok {
		if singular, ok := genderMap[g]; ok {
			return prefix + matchCase(trimmed, singular) + suffix
		}
	}

	// Check for possessive pronouns
	if genderMap, ok := pronounPossessiveSingularByGender[lower]; ok {
		if singular, ok := genderMap[g]; ok {
			return prefix + matchCase(trimmed, singular) + suffix
		}
	}

	// Check for reflexive pronouns
	if genderMap, ok := pronounReflexiveSingularByGender[lower]; ok {
		if singular, ok := genderMap[g]; ok {
			return prefix + matchCase(trimmed, singular) + suffix
		}
	}

	// Fall back to regular Singular() for nouns
	return prefix + e.Singular(trimmed) + suffix
}
