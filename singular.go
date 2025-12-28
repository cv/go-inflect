package inflect

import "strings"

// singularIrregulars maps plural forms to their singular forms (reverse of irregularPlurals).
var singularIrregulars = buildSingularIrregulars()

// buildSingularIrregulars builds the reverse mapping from irregularPlurals.
func buildSingularIrregulars() map[string]string {
	m := make(map[string]string, len(irregularPlurals))
	for singular, plural := range irregularPlurals {
		m[plural] = singular
	}
	return m
}

// feWordBases contains base words whose singular ends in -fe (plural is -ves).
var feWordBases = map[string]bool{
	"kni": true, "wi": true, "li": true,
}

// Singular returns the singular form of an English noun.
//
// Examples:
//   - Singular("cats") returns "cat"
//   - Singular("boxes") returns "box"
//   - Singular("children") returns "child"
//   - Singular("sheep") returns "sheep"
func Singular(word string) string {
	if word == "" {
		return ""
	}

	lower := strings.ToLower(word)

	// Check for irregular plurals first
	if singular, ok := singularIrregulars[lower]; ok {
		return matchCase(word, singular)
	}

	// Check for uncountable/unchanged words
	if unchangedPlurals[lower] {
		return word
	}

	// Check for words ending in -ese, -ois (nationalities that don't change)
	if strings.HasSuffix(lower, "ese") || strings.HasSuffix(lower, "ois") {
		return word
	}

	// Apply suffix rules to singularize
	return applySingularSuffixRules(word, lower)
}

// applySingularSuffixRules applies standard English singularization suffix rules.
func applySingularSuffixRules(word, lower string) string {
	n := len(lower)

	// Words ending in -men -> -man (but not "women" which is irregular)
	if strings.HasSuffix(lower, "men") && n > 3 {
		return word[:len(word)-3] + matchCase(word[len(word)-3:], "man")
	}

	// Words ending in -ves -> -f or -fe
	if strings.HasSuffix(lower, "ves") && n > 3 {
		base := lower[:n-3]
		// Check if original was -fe (knives -> knife, wives -> wife)
		if singularEndsInFe(base) {
			return word[:len(word)-3] + matchSuffix(word, "fe")
		}
		// Otherwise was -f (wolves -> wolf, leaves -> leaf)
		return word[:len(word)-3] + matchSuffix(word, "f")
	}

	// Words ending in -ies (consonant + ies) -> -y
	if strings.HasSuffix(lower, "ies") && n > 3 {
		beforeIes := lower[n-4]
		if !isVowel(rune(beforeIes)) {
			return word[:len(word)-3] + matchSuffix(word, "y")
		}
	}

	// Words ending in -es after sibilants (s, ss, sh, ch, x, z)
	if strings.HasSuffix(lower, "es") && n > 2 {
		if result, ok := singularizeEsSuffix(word, lower[:n-2]); ok {
			return result
		}
	}

	// Default: remove trailing -s
	if strings.HasSuffix(lower, "s") && n > 1 {
		// Don't remove -s from words ending in -ss
		if strings.HasSuffix(lower, "ss") {
			return word
		}
		return word[:len(word)-1]
	}

	// Word doesn't appear to be plural
	return word
}

// singularEndsInFe checks if a base word's singular form ends in -fe.
func singularEndsInFe(base string) bool {
	return feWordBases[base]
}

// singularizeEsSuffix handles -es suffix singularization.
// Returns the singular form and true if a rule matched, empty and false otherwise.
func singularizeEsSuffix(word, base string) (string, bool) {
	// -sses -> -ss (classes -> class)
	if strings.HasSuffix(base, "ss") {
		return word[:len(word)-2], true
	}
	// -shes -> -sh (bushes -> bush)
	if strings.HasSuffix(base, "sh") {
		return word[:len(word)-2], true
	}
	// -ches -> -ch (churches -> church)
	if strings.HasSuffix(base, "ch") {
		return word[:len(word)-2], true
	}
	// -xes -> -x (boxes -> box)
	if strings.HasSuffix(base, "x") {
		return word[:len(word)-2], true
	}
	// -zes -> -z (buzzes -> buzz)
	if strings.HasSuffix(base, "zz") {
		return word[:len(word)-2], true
	}
	// -oes -> -o (heroes -> hero, potatoes -> potato)
	// But not words like "shoes" -> "shoe"
	if strings.HasSuffix(base, "o") && !oExceptionTakesS(base) {
		// Check if this looks like a word that would have taken -oes
		beforeO := base[len(base)-2]
		if !isVowel(rune(beforeO)) {
			return word[:len(word)-2], true
		}
	}
	// Single -s ending with -es: buses -> bus
	if strings.HasSuffix(base, "s") && !strings.HasSuffix(base, "ss") {
		return word[:len(word)-2], true
	}
	return "", false
}
