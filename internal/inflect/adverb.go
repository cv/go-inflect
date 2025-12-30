package inflect

import "strings"

// irregularAdverbs maps adjectives to their irregular adverb forms.
var irregularAdverbs = map[string]string{
	"good":  "well",
	"whole": "wholly",
	"day":   "daily",
	"gay":   "gaily",
}

// unchangedAdverbs contains adjectives that remain unchanged as adverbs.
// These are "flat adverbs" that can function as both adjective and adverb.
var unchangedAdverbs = map[string]bool{
	// Common flat adverbs
	"fast":     true,
	"hard":     true,
	"late":     true,
	"early":    true,
	"straight": true,
	// Words that are already adverbs
	"well":   true,
	"ill":    true,
	"just":   true,
	"only":   true,
	"still":  true,
	"much":   true,
	"far":    true,
	"long":   true,
	"likely": true,
	"even":   true,
	// A-prefix words that already function as adverbs
	"ahead":  true,
	"away":   true,
	"alone":  true,
	"aloud":  true,
	"apart":  true,
	"abroad": true,
	"afoot":  true,
	"afloat": true,
	"ashore": true,
	"asleep": true,
	"awake":  true,
	"alive":  true,
	"askew":  true,
	"awry":   true,
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

// leKeepsE contains -le words that keep the e when forming adverbs.
// Most -le words drop the e (gentle → gently), but these keep it.
var leKeepsE = map[string]bool{
	"sole": true,
}

// shortYAdverbs contains short monosyllabic adjectives ending in consonant + y
// that just add -ly instead of changing y to i (shy → shyly, not shily).
var shortYAdverbs = map[string]bool{
	"shy":  true,
	"sly":  true,
	"dry":  true,
	"wry":  true,
	"coy":  true,
	"spry": true,
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

	// Rule 3a: Adjectives ending in -ile: keep the e, add -ly (agile → agilely)
	if strings.HasSuffix(lower, "ile") {
		return adj + matchSuffix(adj, "ly")
	}

	// Rule 3b: Adjectives ending in -le (not -ile): usually change -le to -ly (gentle → gently)
	// Exception: some words keep the e (sole → solely)
	if strings.HasSuffix(lower, "le") {
		if leKeepsE[lower] {
			return adj + matchSuffix(adj, "ly")
		}
		return adj[:len(adj)-2] + matchSuffix(adj, "ly")
	}

	// Rule 2: Adjectives ending in consonant + y
	if strings.HasSuffix(lower, "y") && len(lower) > 1 {
		beforeY := lower[len(lower)-2]
		if !isVowel(rune(beforeY)) {
			// Short monosyllabic words just add -ly (shy → shyly)
			if shortYAdverbs[lower] {
				return adj + matchSuffix(adj, "ly")
			}
			// Most words change y to -ily (happy → happily)
			return adj[:len(adj)-1] + matchSuffix(adj, "ily")
		}
	}

	// Rule 1: Most adjectives: add -ly (quick → quickly)
	return adj + matchSuffix(adj, "ly")
}
