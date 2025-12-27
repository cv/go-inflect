package inflect

import "strings"

// PossessiveStyleType represents the style for forming possessives of words ending in s.
type PossessiveStyleType int

const (
	// PossessiveModern adds 's to all singular nouns, including those ending in s.
	// Example: James's, boss's, class's
	// This is the style recommended by most modern style guides.
	PossessiveModern PossessiveStyleType = iota

	// PossessiveTraditional adds only ' to singular nouns ending in s.
	// Example: James', boss', class'
	// This is the older, traditional style still used in some contexts.
	PossessiveTraditional
)

// possessiveStyle controls the style for forming possessives of words ending in s.
var possessiveStyle = PossessiveModern

// PossessiveStyle sets the style for forming possessives of words ending in s.
// Use PossessiveModern (default) for "James's" or PossessiveTraditional for "James'".
func PossessiveStyle(style PossessiveStyleType) {
	possessiveStyle = style
}

// GetPossessiveStyle returns the current possessive style setting.
func GetPossessiveStyle() PossessiveStyleType {
	return possessiveStyle
}

// Possessive returns the possessive form of an English noun.
//
// Rules applied:
//   - Singular nouns: add 's (cat → cat's)
//   - Plural nouns ending in s: add only ' (cats → cats')
//   - Plural nouns not ending in s: add 's (children → children's)
//   - Singular nouns ending in s: add 's or ' based on PossessiveStyle setting
//   - Words already in possessive form are returned unchanged
//
// Examples:
//   - Possessive("cat") returns "cat's"
//   - Possessive("cats") returns "cats'"
//   - Possessive("children") returns "children's"
//   - Possessive("James") returns "James's" (with PossessiveModern)
//   - Possessive("James") returns "James'" (with PossessiveTraditional)
func Possessive(word string) string {
	if word == "" {
		return ""
	}

	// Check if already possessive
	if isAlreadyPossessive(word) {
		return word
	}

	lower := strings.ToLower(word)

	// Check if the word ends in s
	endsInS := strings.HasSuffix(lower, "s")

	// Determine if the word is a "true" plural (not just a singular ending in s)
	// We check by seeing if pluralizing the singular form gives us the original word
	isPluralForm := isTruePlural(word, lower)

	// Handle plural nouns ending in s: add only '
	if isPluralForm && endsInS {
		return word + "'"
	}

	// Handle plural nouns not ending in s (children, men, etc.): add 's
	if isPluralForm {
		return word + matchSuffix(word, "'s")
	}

	// Singular noun ending in s
	if endsInS {
		if possessiveStyle == PossessiveTraditional {
			return word + "'"
		}
		return word + matchSuffix(word, "'s")
	}

	// Default: add 's
	return word + matchSuffix(word, "'s")
}

// isTruePlural determines if a word is actually a plural form (not just a singular ending in s).
func isTruePlural(word, lower string) bool {
	// Check for known irregular plurals that don't end in s
	irregularPluralNoS := map[string]bool{
		"children": true, "men": true, "women": true, "people": true,
		"mice": true, "geese": true, "feet": true, "teeth": true,
		"oxen": true, "lice": true, "dice": true,
	}
	if irregularPluralNoS[lower] {
		return true
	}

	// Words ending in double-s are typically singular (boss, class, glass, dress, etc.)
	if strings.HasSuffix(lower, "ss") {
		return false
	}

	// Proper names (capitalized) ending in s are typically singular
	// But we need to check if it might be a common noun made plural (like "Cats")
	if isProperName(word) && strings.HasSuffix(lower, "s") {
		// Check if this is a known common word in plural form
		singular := Singular(word)
		singularLower := strings.ToLower(singular)

		// Only consider it a plural if the singular is a known common word
		// We check this by seeing if the singular exists in common word patterns
		if singularLower != lower && isLikelyCommonNoun(singularLower) {
			return true // It's actually a plural of a common noun
		}
		return false // Proper name, not a plural
	}

	// Common singular words ending in s that shouldn't be treated as plural
	singularEndsInS := map[string]bool{
		"bus": true, "gas": true, "lens": true, "atlas": true,
		"iris": true, "plus": true, "minus": true, "virus": true,
		"bonus": true, "focus": true, "campus": true, "census": true,
		"corpus": true, "genius": true, "nexus": true, "oasis": true,
		"basis": true, "thesis": true, "crisis": true, "analysis": true,
		"diagnosis": true, "hypothesis": true, "parenthesis": true,
		"synopsis": true, "emphasis": true, "cosmos": true, "chaos": true,
		"ethos": true, "pathos": true, "logos": true, "status": true,
		"apparatus": true, "hiatus": true, "impetus": true, "radius": true,
		"nucleus": true, "syllabus": true, "stimulus": true, "fungus": true,
		"cactus": true, "octopus": true, "platypus": true, "walrus": true,
		"yes": true, "no": true, "us": true, "this": true, "thus": true,
	}
	if singularEndsInS[lower] {
		return false
	}

	// For words ending in s, check if they follow common plural patterns
	if strings.HasSuffix(lower, "s") {
		// Try to get the singular form
		singular := Singular(word)
		singularLower := strings.ToLower(singular)

		// If singular is the same, it's not a plural we can detect
		if singularLower == lower {
			return false
		}

		// Check if the singular form is a real word (not just stripping 's')
		// A "real" plural typically has a singular that differs by more than just the final 's'
		// or the singular-to-plural transformation follows known rules

		// If removing 's' gives a very short word (1-2 chars), likely not a valid singular
		if len(singularLower) <= 2 {
			return false
		}

		// Check for -es plurals (boxes, churches, etc.)
		if strings.HasSuffix(lower, "es") {
			baseWithoutEs := lower[:len(lower)-2]
			// If base ends in s, x, z, ch, sh, it's a valid -es plural
			if strings.HasSuffix(baseWithoutEs, "s") ||
				strings.HasSuffix(baseWithoutEs, "x") ||
				strings.HasSuffix(baseWithoutEs, "z") ||
				strings.HasSuffix(baseWithoutEs, "ch") ||
				strings.HasSuffix(baseWithoutEs, "sh") {
				return true
			}
		}

		// Check for -ies plurals (cities, babies, etc.)
		if strings.HasSuffix(lower, "ies") && len(lower) > 3 {
			return true
		}

		// Default: if singular differs and pluralizing it gives us back the word, it's plural
		pluralOfSingular := strings.ToLower(Plural(singular))
		if pluralOfSingular == lower && singularLower != lower {
			// Additional check: singular should look "complete" (not truncated)
			// Words like "bu" from "bus" don't look complete
			if !looksLikeCompleteWord(singularLower) {
				return false
			}
			return true
		}
	}

	return false
}

// isLikelyCommonNoun checks if a word is likely a common noun (not a proper name).
// This uses heuristics to identify common words vs. truncated proper names.
func isLikelyCommonNoun(word string) bool {
	if len(word) < 2 {
		return false
	}

	// Common nouns we know about (a sample - most validation is heuristic)
	commonNouns := map[string]bool{
		"cat": true, "dog": true, "book": true, "car": true, "house": true,
		"boy": true, "girl": true, "man": true, "woman": true, "child": true,
		"tree": true, "bird": true, "fish": true, "day": true, "night": true,
		"hand": true, "foot": true, "head": true, "eye": true, "ear": true,
		"door": true, "window": true, "table": true, "chair": true, "bed": true,
		"teacher": true, "student": true, "parent": true, "friend": true,
		"city": true, "country": true, "state": true, "street": true,
		"box": true, "bag": true, "ball": true, "cup": true, "glass": true,
		"paper": true, "pen": true, "key": true, "phone": true, "computer": true,
	}
	if commonNouns[word] {
		return true
	}

	// Known truncated name patterns that should NOT be considered common nouns
	// These are singularized forms of common names
	truncatedNames := map[string]bool{
		// From -es names
		"jame": true, "charle": true, "jone": true, "mose": true, "jesu": true,
		"thoma": true, "jess": true, "ross": true, "walle": true, "jule": true,
		"mile": true, "gile": true, "style": true, "kyle": true, "achille": true,
		// From -s names
		"william": true, "adam": true, "lewi": true, "davi": true, "elli": true,
		"harri": true, "morri": true, "denni": true, "franci": true, "chri": true,
		// Short truncated forms
		"mos": true, "ros": true, "gus": true,
	}
	if truncatedNames[word] {
		return false
	}

	// Check for truncated patterns that indicate non-words
	lastChar := rune(word[len(word)-1])

	// Words ending in 'i' or 'u' alone are rare in English common nouns
	if lastChar == 'i' || lastChar == 'u' {
		return false
	}

	// Short words ending in 'a' are often truncated names (thoma, anna → ann)
	if lastChar == 'a' && len(word) <= 5 {
		validShortA := map[string]bool{
			"data": true, "sofa": true, "mega": true, "soda": true, "mama": true,
			"papa": true, "diva": true, "yoga": true, "cola": true, "area": true,
			"idea": true, "lava": true, "toga": true, "tuna": true, "visa": true,
			"beta": true, "meta": true, "aqua": true, "aura": true, "era": true,
		}
		if !validShortA[word] {
			return false
		}
	}

	// Words ending in 'e' with consonant clusters are often truncated
	if lastChar == 'e' && len(word) >= 4 {
		secondLastChar := rune(word[len(word)-2])
		thirdLastChar := rune(word[len(word)-3])
		// Consonant + consonant + e is suspicious (charle, achille)
		if !isVowel(secondLastChar) && !isVowel(thirdLastChar) {
			return false
		}
	}

	return true
}

// looksLikeCompleteWord checks if a word looks like a valid English word.
// This is a heuristic to filter out truncated forms like "bu" from "bus".
func looksLikeCompleteWord(word string) bool {
	if len(word) < 2 {
		return false
	}

	secondLastChar := rune(word[len(word)-2])

	// Words typically don't end in these consonants without a vowel before them
	invalidEndings := []string{"b", "c", "d", "f", "g", "h", "j", "k", "m", "p", "q", "v", "w", "z"}
	for _, ending := range invalidEndings {
		if strings.HasSuffix(word, ending) {
			// Check if the word has a vowel before the ending
			if !isVowel(secondLastChar) {
				return false
			}
		}
	}

	return true
}

// isAlreadyPossessive checks if a word is already in possessive form.
func isAlreadyPossessive(word string) bool {
	// Check for 's or s' at the end
	if strings.HasSuffix(word, "'s") || strings.HasSuffix(word, "'S") {
		return true
	}
	if strings.HasSuffix(word, "s'") || strings.HasSuffix(word, "S'") {
		return true
	}
	return false
}
