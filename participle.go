package inflect

import "strings"

// doubleConsonantWords contains multi-syllable words that double the final consonant.
var doubleConsonantWords = map[string]bool{
	"admit": true, "begin": true, "commit": true, "compel": true,
	"confer": true, "control": true, "defer": true, "deter": true,
	"equip": true, "excel": true, "expel": true, "forget": true,
	"incur": true, "occur": true, "omit": true, "patrol": true,
	"permit": true, "prefer": true, "propel": true, "rebel": true,
	"recur": true, "refer": true, "regret": true, "repel": true,
	"submit": true, "transfer": true, "transmit": true, "upset": true,
}

// PresentParticiple converts a verb to its present participle (-ing) form.
//
// Examples:
//   - PresentParticiple("run") returns "running" (double consonant)
//   - PresentParticiple("make") returns "making" (drop silent e)
//   - PresentParticiple("play") returns "playing" (just add -ing)
//   - PresentParticiple("die") returns "dying" (ie -> ying)
//   - PresentParticiple("see") returns "seeing" (ee -> eeing)
//   - PresentParticiple("panic") returns "panicking" (c -> ck)
func PresentParticiple(verb string) string {
	if verb == "" {
		return ""
	}

	lower := strings.ToLower(verb)
	n := len(lower)

	// Already a present participle (ends in doubled consonant + ing, like "running")
	if isAlreadyParticiple(lower) {
		return verb
	}

	// Single letter verbs - just add -ing
	if n == 1 {
		return verb + matchSuffix(verb, "ing")
	}

	// Words ending in -ie: change to -ying (die -> dying, lie -> lying)
	if strings.HasSuffix(lower, "ie") {
		return verb[:len(verb)-2] + matchSuffix(verb, "ying")
	}

	// Words ending in -ee: just add -ing (see -> seeing, flee -> fleeing)
	if strings.HasSuffix(lower, "ee") {
		return verb + matchSuffix(verb, "ing")
	}

	// Words ending in -ye, -oe: just add -ing (dye -> dyeing, hoe -> hoeing)
	if strings.HasSuffix(lower, "ye") || strings.HasSuffix(lower, "oe") {
		return verb + matchSuffix(verb, "ing")
	}

	// Words ending in -nge/-inge: keep the e (singe -> singeing)
	if strings.HasSuffix(lower, "nge") || strings.HasSuffix(lower, "inge") {
		return verb + matchSuffix(verb, "ing")
	}

	// Words ending in -c: add k before -ing (panic -> panicking)
	if strings.HasSuffix(lower, "c") {
		return verb + matchSuffix(verb, "king")
	}

	// Words ending in consonant + e (silent e): drop e, add -ing
	// But keep e if it's the only vowel in the word (be -> being)
	if strings.HasSuffix(lower, "e") && n >= 2 {
		beforeE := rune(lower[n-2])
		if !isVowel(beforeE) {
			// Check if 'e' is the only vowel (not a silent e)
			vowelCount := countVowels(lower[:n-1]) // count vowels excluding final 'e'
			if vowelCount == 0 {
				// 'e' is the only vowel, keep it (be -> being)
				return verb + matchSuffix(verb, "ing")
			}
			return verb[:len(verb)-1] + matchSuffix(verb, "ing")
		}
		// vowel + e: just add -ing
		return verb + matchSuffix(verb, "ing")
	}

	// Check for CVC pattern that requires doubling the final consonant
	if shouldDoubleConsonant(lower) {
		lastChar := verb[len(verb)-1:]
		return verb + matchSuffix(verb, strings.ToLower(lastChar)+"ing")
	}

	// Default: just add -ing
	return verb + matchSuffix(verb, "ing")
}

// isAlreadyParticiple checks if a word is already a present participle.
// This catches words like "running", "sitting" but not base verbs like "sing".
func isAlreadyParticiple(lower string) bool {
	if !strings.HasSuffix(lower, "ing") {
		return false
	}
	n := len(lower)
	if n < 5 {
		return false
	}

	// Check for doubled consonant before -ing (running, sitting, hitting)
	beforeIng := lower[n-4]
	beforeThat := lower[n-5]
	if beforeIng == beforeThat && !isVowel(rune(beforeIng)) {
		return true
	}

	// Check for common participle patterns
	// Words ending in -ting, -ning, -ping, etc. after a consonant
	// But not "sing", "ring", "bring" which are base verbs

	return false
}

// shouldDoubleConsonant checks if the final consonant should be doubled.
// This applies to CVC (consonant-vowel-consonant) patterns in stressed syllables.
func shouldDoubleConsonant(lower string) bool {
	n := len(lower)
	if n < 3 {
		return false
	}

	lastChar := rune(lower[n-1])

	// Don't double w, x, y
	if lastChar == 'w' || lastChar == 'x' || lastChar == 'y' {
		return false
	}

	// Must end in a consonant
	if isVowel(lastChar) {
		return false
	}

	// Check for single vowel before the final consonant
	beforeLast := rune(lower[n-2])
	if !isVowel(beforeLast) {
		return false
	}

	// Don't double if there's a vowel digraph (two vowels in a row before consonant)
	// Examples: eat, read, beat, lead - these have "ea" before the final consonant
	if n >= 3 && isVowel(rune(lower[n-3])) {
		return false
	}

	// At this point we know the word ends in consonant + single vowel + consonant

	// For short words (3 letters): double if CVC pattern
	// Examples: run, sit, hit, cut
	if n == 3 {
		return true
	}

	// For 4-letter words: double only if there's a single vowel cluster
	// "stop", "drop", "skip", "plan" -> double (single vowel)
	// "open" -> don't double (two separate vowels = multi-syllable)
	if n == 4 {
		// Count distinct vowel clusters
		if countVowels(lower) == 1 {
			return true
		}
		return false
	}

	// For multi-syllable words, check for common patterns that double
	// Words ending in stressed syllables typically double
	return doubleConsonantWords[lower]
}

// countVowels counts the number of vowels in a string.
func countVowels(s string) int {
	count := 0
	for _, r := range s {
		if isVowel(r) {
			count++
		}
	}
	return count
}
