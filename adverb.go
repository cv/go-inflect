package inflect

import "strings"

// irregularAdverbs maps adjectives to their irregular adverb forms.
var irregularAdverbs = map[string]string{
	"good": "well",
}

// unchangedAdverbs contains adjectives that remain unchanged as adverbs.
var unchangedAdverbs = map[string]bool{
	"fast":     true,
	"hard":     true,
	"late":     true,
	"early":    true,
	"straight": true,
}

// Adverb returns the adverb form of an English adjective.
//
// Examples:
//   - Adverb("quick") returns "quickly"
//   - Adverb("happy") returns "happily"
//   - Adverb("gentle") returns "gently"
//   - Adverb("true") returns "truly"
//   - Adverb("basic") returns "basically"
//   - Adverb("good") returns "well"
//   - Adverb("fast") returns "fast"
func Adverb(adj string) string {
	if adj == "" {
		return ""
	}

	lower := strings.ToLower(adj)

	// Check irregular forms first
	if adv, ok := irregularAdverbs[lower]; ok {
		return matchCase(adj, adv)
	}

	// Check unchanged forms
	if unchangedAdverbs[lower] {
		return adj
	}

	return applyAdverbSuffix(adj, lower)
}

// applyAdverbSuffix applies the appropriate suffix to form an adverb.
func applyAdverbSuffix(adj, lower string) string {
	// Rule: public -> publicly (exception to -ic rule, just add -ly)
	if lower == "public" {
		return adj + matchSuffix(adj, "ly")
	}

	// Rule 5: Adjectives ending in -ic: add -ally (basic → basically)
	if strings.HasSuffix(lower, "ic") {
		return adj + matchSuffix(adj, "ally")
	}

	// Rule: Adjectives ending in -ll: add -y (full → fully)
	if strings.HasSuffix(lower, "ll") {
		return adj + matchSuffix(adj, "y")
	}

	// Rule 4: Adjectives ending in -ue: drop e, add -ly (true → truly)
	if strings.HasSuffix(lower, "ue") {
		return adj[:len(adj)-1] + matchSuffix(adj, "ly")
	}

	// Rule 3: Adjectives ending in -le: change -le to -ly (gentle → gently)
	if strings.HasSuffix(lower, "le") {
		return adj[:len(adj)-2] + matchSuffix(adj, "ly")
	}

	// Rule 2: Adjectives ending in consonant + y: change y to -ily (happy → happily)
	if strings.HasSuffix(lower, "y") && len(lower) > 1 {
		beforeY := lower[len(lower)-2]
		if !isVowel(rune(beforeY)) {
			return adj[:len(adj)-1] + matchSuffix(adj, "ily")
		}
	}

	// Rule 1: Most adjectives: add -ly (quick → quickly)
	return adj + matchSuffix(adj, "ly")
}
