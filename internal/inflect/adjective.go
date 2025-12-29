package inflect

import "strings"

// irregularComparatives maps adjectives to their irregular comparative forms.
var irregularComparatives = map[string]string{
	"good":   "better",
	"well":   "better", // as adjective (feeling well)
	"bad":    "worse",
	"ill":    "worse",
	"far":    "farther",
	"little": "less",
	"much":   "more",
	"many":   "more",
}

// irregularSuperlatives maps adjectives to their irregular superlative forms.
var irregularSuperlatives = map[string]string{
	"good":   "best",
	"well":   "best", // as adjective (feeling well)
	"bad":    "worst",
	"ill":    "worst",
	"far":    "farthest",
	"little": "least",
	"much":   "most",
	"many":   "most",
}

// twoSyllableWithSuffix contains two-syllable adjectives that take -er/-est
// instead of more/most. These typically end in -y, -le, -ow, -er.
var twoSyllableWithSuffix = map[string]bool{
	"simple":  true,
	"gentle":  true,
	"narrow":  true,
	"shallow": true,
	"quiet":   true,
	"clever":  true,
	"common":  true,
	"hollow":  true,
	"mellow":  true,
	"yellow":  true,
	"feeble":  true,
	"humble":  true,
	"noble":   true,
	"able":    true,
	"tender":  true,
	"bitter":  true,
	"slender": true,
}

// Comparative returns the comparative form of an English adjective.
//
// Examples:
//   - Comparative("big") returns "bigger"
//   - Comparative("happy") returns "happier"
//   - Comparative("beautiful") returns "more beautiful"
//   - Comparative("good") returns "better"
func Comparative(adj string) string {
	if adj == "" {
		return ""
	}

	lower := strings.ToLower(adj)

	// Check irregular forms first
	if comp, ok := irregularComparatives[lower]; ok {
		return matchCase(adj, comp)
	}

	// Determine if we should use -er or "more"
	if useSuffix := shouldUseSuffix(lower); useSuffix {
		return applyComparativeSuffix(adj, lower)
	}

	// Use "more" for longer adjectives
	return applyMore(adj)
}

// Superlative returns the superlative form of an English adjective.
//
// Examples:
//   - Superlative("big") returns "biggest"
//   - Superlative("happy") returns "happiest"
//   - Superlative("beautiful") returns "most beautiful"
//   - Superlative("good") returns "best"
func Superlative(adj string) string {
	if adj == "" {
		return ""
	}

	lower := strings.ToLower(adj)

	// Check irregular forms first
	if sup, ok := irregularSuperlatives[lower]; ok {
		return matchCase(adj, sup)
	}

	// Determine if we should use -est or "most"
	if useSuffix := shouldUseSuffix(lower); useSuffix {
		return applySuperlativeSuffix(adj, lower)
	}

	// Use "most" for longer adjectives
	return applyMost(adj)
}

// shouldUseSuffix determines if an adjective should use -er/-est suffixes
// rather than more/most.
func shouldUseSuffix(lower string) bool {
	syllables := countSyllables(lower)

	// One-syllable adjectives always use -er/-est
	if syllables == 1 {
		return true
	}

	// Two-syllable adjectives ending in -y use -er/-est
	if syllables == 2 && strings.HasSuffix(lower, "y") {
		return true
	}

	// Check for specific two-syllable words that take -er/-est
	if twoSyllableWithSuffix[lower] {
		return true
	}

	// Three or more syllables use more/most
	return false
}

// countSyllables estimates the number of syllables in a word.
// This is a simplified heuristic based on vowel groups.
func countSyllables(word string) int {
	if word == "" {
		return 0
	}

	word = strings.ToLower(word)
	count := 0
	prevVowel := false

	for i, r := range word {
		isCurrentVowel := isVowel(r)

		// Count vowel groups (consecutive vowels count as one)
		if isCurrentVowel && !prevVowel {
			count++
		}

		prevVowel = isCurrentVowel

		// Handle silent e at end
		if i == len(word)-1 && r == 'e' && count > 1 {
			count--
		}
	}

	// Every word has at least one syllable
	if count == 0 {
		count = 1
	}

	return count
}

// applyComparativeSuffix adds the -er suffix with appropriate modifications.
func applyComparativeSuffix(adj, lower string) string {
	// Words ending in -e: just add -r
	if strings.HasSuffix(lower, "e") {
		return adj + matchSuffix(adj, "r")
	}

	// Words ending in consonant + y: change y to -ier
	if strings.HasSuffix(lower, "y") && len(lower) > 1 {
		beforeY := lower[len(lower)-2]
		if !isVowel(rune(beforeY)) {
			return adj[:len(adj)-1] + matchSuffix(adj, "ier")
		}
	}

	// CVC pattern: double the final consonant
	if shouldDoubleFinalConsonant(lower) {
		lastChar := string(lower[len(lower)-1])
		return adj + matchSuffix(adj, lastChar+"er")
	}

	// Default: just add -er
	return adj + matchSuffix(adj, "er")
}

// applySuperlativeSuffix adds the -est suffix with appropriate modifications.
func applySuperlativeSuffix(adj, lower string) string {
	// Words ending in -e: just add -st
	if strings.HasSuffix(lower, "e") {
		return adj + matchSuffix(adj, "st")
	}

	// Words ending in consonant + y: change y to -iest
	if strings.HasSuffix(lower, "y") && len(lower) > 1 {
		beforeY := lower[len(lower)-2]
		if !isVowel(rune(beforeY)) {
			return adj[:len(adj)-1] + matchSuffix(adj, "iest")
		}
	}

	// CVC pattern: double the final consonant
	if shouldDoubleFinalConsonant(lower) {
		lastChar := string(lower[len(lower)-1])
		return adj + matchSuffix(adj, lastChar+"est")
	}

	// Default: just add -est
	return adj + matchSuffix(adj, "est")
}

// shouldDoubleFinalConsonant checks if the final consonant should be doubled.
// This applies to short words with a CVC (consonant-vowel-consonant) pattern.
func shouldDoubleFinalConsonant(lower string) bool {
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

// applyMore prepends "more" to the adjective, preserving case.
func applyMore(adj string) string {
	if isAllUpper(adj) {
		return "MORE " + adj
	}
	if adj != "" && adj[0] >= 'A' && adj[0] <= 'Z' {
		return "More " + toTitleCaseWord(adj)
	}
	return "more " + adj
}

// applyMost prepends "most" to the adjective, preserving case.
func applyMost(adj string) string {
	if isAllUpper(adj) {
		return "MOST " + adj
	}
	if adj != "" && adj[0] >= 'A' && adj[0] <= 'Z' {
		return "Most " + toTitleCaseWord(adj)
	}
	return "most " + adj
}

// toTitleCaseWord converts a word to title case (first letter uppercase, rest lowercase).
func toTitleCaseWord(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(string(s[0])) + strings.ToLower(s[1:])
}
