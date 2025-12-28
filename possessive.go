package inflect

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

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

// Package-level maps to avoid allocation on each call.
var (
	// irregularPluralNoS contains known irregular plurals that don't end in s.
	irregularPluralNoS = map[string]bool{
		"children": true, "men": true, "women": true, "people": true,
		"mice": true, "geese": true, "feet": true, "teeth": true,
		"oxen": true, "lice": true, "dice": true,
	}

	// singularEndsInS contains common singular words ending in s that shouldn't be treated as plural.
	singularEndsInS = map[string]bool{
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

	// commonNouns contains common nouns we know about (a sample - most validation is heuristic).
	commonNouns = map[string]bool{
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

	// truncatedNames contains known truncated name patterns that should NOT be considered common nouns.
	// These are singularized forms of common names.
	truncatedNames = map[string]bool{
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

	// validShortA contains valid short words ending in 'a'.
	validShortA = map[string]bool{
		"data": true, "sofa": true, "mega": true, "soda": true, "mama": true,
		"papa": true, "diva": true, "yoga": true, "cola": true, "area": true,
		"idea": true, "lava": true, "toga": true, "tuna": true, "visa": true,
		"beta": true, "meta": true, "aqua": true, "aura": true, "era": true,
	}
)

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

	// Check if already possessive (no allocation needed)
	if isAlreadyPossessive(word) {
		return word
	}

	// Check if the word ends in s (case-insensitive, no allocation)
	endsInS := endsWithS(word)

	// Fast path for simple singular words not ending in s
	if !endsInS {
		// Check for irregular plurals that don't end in s
		lower := strings.ToLower(word)
		if irregularPluralNoS[lower] {
			// Plural noun not ending in s: add 's
			return word + matchSuffix(word, "'s")
		}
		// Simple singular: add 's
		return word + matchSuffix(word, "'s")
	}

	// Word ends in s - need to determine if plural or singular
	lower := strings.ToLower(word)

	// Words ending in double-s are typically singular (boss, class, glass, dress, etc.)
	if strings.HasSuffix(lower, "ss") {
		if possessiveStyle == PossessiveTraditional {
			return word + "'"
		}
		return word + matchSuffix(word, "'s")
	}

	// Check for known singular words ending in s
	if singularEndsInS[lower] {
		if possessiveStyle == PossessiveTraditional {
			return word + "'"
		}
		return word + matchSuffix(word, "'s")
	}

	// Proper names ending in s are typically singular
	if isProperName(word) {
		// Check if this might be a plural of a common noun (like "Cats")
		singular := Singular(word)
		singularLower := strings.ToLower(singular)
		if singularLower != lower && isLikelyCommonNoun(singularLower) {
			// It's actually a plural of a common noun
			return word + "'"
		}
		// Proper name, not a plural - treat as singular ending in s
		if possessiveStyle == PossessiveTraditional {
			return word + "'"
		}
		return word + matchSuffix(word, "'s")
	}

	// Check if it's a true plural
	if isTruePluralEndsInS(word, lower) {
		return word + "'"
	}

	// Default: singular ending in s
	if possessiveStyle == PossessiveTraditional {
		return word + "'"
	}
	return word + matchSuffix(word, "'s")
}

// endsWithS checks if a word ends with 's' or 'S' without allocation.
func endsWithS(word string) bool {
	if word == "" {
		return false
	}
	lastByte := word[len(word)-1]
	return lastByte == 's' || lastByte == 'S'
}

// isTruePluralEndsInS determines if a word ending in s is actually a plural form.
// This is called only for words that end in 's' and are not proper names.
func isTruePluralEndsInS(word, lower string) bool {
	// Try to get the singular form
	singular := Singular(word)
	singularLower := strings.ToLower(singular)

	// If singular is the same, it's not a plural we can detect
	if singularLower == lower {
		return false
	}

	// If removing 's' gives a very short word (1-2 chars), likely not a valid singular
	if len(singularLower) <= 2 {
		return false
	}

	// Check for -es plurals (boxes, churches, etc.)
	if isEsPlural(lower) {
		return true
	}

	// Check for -ies plurals (cities, babies, etc.)
	if strings.HasSuffix(lower, "ies") && len(lower) > 3 {
		return true
	}

	// Default: if singular differs and pluralizing it gives us back the word, it's plural
	return isValidPluralOfSingular(lower, singular, singularLower)
}

// isEsPlural checks if a word is a valid -es plural (boxes, churches, etc.).
func isEsPlural(lower string) bool {
	if !strings.HasSuffix(lower, "es") {
		return false
	}
	baseWithoutEs := lower[:len(lower)-2]
	// If base ends in s, x, z, ch, sh, it's a valid -es plural
	return strings.HasSuffix(baseWithoutEs, "s") ||
		strings.HasSuffix(baseWithoutEs, "x") ||
		strings.HasSuffix(baseWithoutEs, "z") ||
		strings.HasSuffix(baseWithoutEs, "ch") ||
		strings.HasSuffix(baseWithoutEs, "sh")
}

// isValidPluralOfSingular checks if lower is a valid plural of singular.
func isValidPluralOfSingular(lower, singular, singularLower string) bool {
	pluralOfSingular := strings.ToLower(Plural(singular))
	if pluralOfSingular != lower || singularLower == lower {
		return false
	}
	// Additional check: singular should look "complete" (not truncated)
	// Words like "bu" from "bus" don't look complete
	return looksLikeCompleteWord(singularLower)
}

// isLikelyCommonNoun checks if a word is likely a common noun (not a proper name).
// This uses heuristics to identify common words vs. truncated proper names.
func isLikelyCommonNoun(word string) bool {
	if len(word) < 2 {
		return false
	}

	// Check known common nouns
	if commonNouns[word] {
		return true
	}

	// Check known truncated name patterns
	if truncatedNames[word] {
		return false
	}

	// Check for truncated patterns that indicate non-words
	lastRune, _ := utf8.DecodeLastRuneInString(word)

	// Words ending in 'i' or 'u' alone are rare in English common nouns
	if lastRune == 'i' || lastRune == 'u' {
		return false
	}

	// Short words ending in 'a' are often truncated names (thoma, anna → ann)
	if lastRune == 'a' && len(word) <= 5 {
		if !validShortA[word] {
			return false
		}
	}

	// Words ending in 'e' with consonant clusters are often truncated
	if lastRune == 'e' && len(word) >= 4 {
		// Get the second and third last runes
		remaining := word[:len(word)-1]
		secondLastRune, size := utf8.DecodeLastRuneInString(remaining)
		if size > 0 {
			remaining = remaining[:len(remaining)-size]
			thirdLastRune, _ := utf8.DecodeLastRuneInString(remaining)
			// Consonant + consonant + e is suspicious (charle, achille)
			if !isVowel(secondLastRune) && !isVowel(thirdLastRune) {
				return false
			}
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

	// Get second last character
	remaining := word[:len(word)-1]
	secondLastRune, _ := utf8.DecodeLastRuneInString(remaining)

	// Get last character
	lastRune, _ := utf8.DecodeLastRuneInString(word)
	lastLower := unicode.ToLower(lastRune)

	// Words typically don't end in these consonants without a vowel before them
	switch lastLower {
	case 'b', 'c', 'd', 'f', 'g', 'h', 'j', 'k', 'm', 'p', 'q', 'v', 'w', 'z':
		// Check if the word has a vowel before the ending
		if !isVowel(secondLastRune) {
			return false
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
