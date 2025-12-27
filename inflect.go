// Package inflect provides English language inflection utilities.
//
// It offers functions for pluralization, singularization, indefinite article
// selection (a/an), number-to-words conversion, ordinals, and more.
//
// This is a Go port of the Python inflect library.
package inflect

import (
	"fmt"
	"strings"
	"unicode"
)

// An returns the word prefixed with the appropriate indefinite article ("a" or "an").
//
// The selection follows standard English rules:
//   - Use "an" before vowel sounds: "an apple", "an hour"
//   - Use "a" before consonant sounds: "a cat", "a university"
//
// Special cases handled:
//   - Silent 'h': "an honest person"
//   - Vowels with consonant sounds: "a Ukrainian", "a unanimous decision"
//   - Abbreviations: "a YAML file", "a JSON object"
func An(word string) string {
	if word == "" {
		return ""
	}

	if needsAn(word) {
		return "an " + word
	}
	return "a " + word
}

// needsAn determines if a word/phrase should be preceded by "an" (vs "a").
func needsAn(text string) bool {
	// Get the first word to analyze
	firstWord := strings.Fields(text)[0]
	lower := strings.ToLower(firstWord)

	// Check for silent 'h' words that take "an"
	silentH := []string{"honest", "heir", "heiress", "heirloom", "honor", "honour", "hour", "hourly"}
	for _, h := range silentH {
		if strings.HasPrefix(lower, h) {
			return true
		}
	}

	// Check for abbreviations/acronyms (all uppercase or known patterns)
	if isAbbreviation(firstWord) {
		return abbreviationNeedsAn(firstWord)
	}

	// Check for known lowercase abbreviations pronounced letter-by-letter
	lowercaseAbbrevs := []string{"mpeg", "jpeg", "gif", "sql", "html", "xml", "fbi", "cia", "nsa"}
	for _, abbr := range lowercaseAbbrevs {
		if lower == abbr {
			return abbreviationNeedsAn(strings.ToUpper(abbr))
		}
	}

	// Check for special vowel patterns that sound like consonants (take "a")
	// "uni-" sounds like "yoo-", "eu-" sounds like "yoo-", etc.
	consonantVowelPatterns := []string{
		"uni", "upon", "use", "used", "user", "using", "usual",
		"usu", "uran", "uret", "euro", "ewe", "onc", "one",
		"onet", // onetime
	}
	for _, pat := range consonantVowelPatterns {
		if strings.HasPrefix(lower, pat) {
			return false
		}
	}

	// Check for "U" followed by consonant then vowel pattern (sounds like "you")
	// e.g., Ugandan, Ukrainian, Unabomber, unanimous
	if len(lower) >= 2 && lower[0] == 'u' {
		if isConsonantYSound(lower) {
			return false
		}
	}

	// Special case: single letters
	if len(firstWord) == 1 {
		return isVowelSound(rune(lower[0]))
	}

	// Default: check if first letter is a vowel
	first := rune(lower[0])
	return isVowelSound(first)
}

// isAbbreviation checks if a word appears to be an abbreviation/acronym.
func isAbbreviation(word string) bool {
	if len(word) < 2 {
		return false
	}

	// Check if all letters are uppercase
	allUpper := true
	for _, r := range word {
		if unicode.IsLetter(r) && !unicode.IsUpper(r) {
			allUpper = false
			break
		}
	}
	return allUpper
}

// abbreviationNeedsAn checks if an abbreviation should take "an".
// This depends on how the first letter is pronounced.
func abbreviationNeedsAn(abbrev string) bool {
	if len(abbrev) == 0 {
		return false
	}

	// Letters whose names start with vowel sounds:
	// A (ay), E (ee), F (eff), H (aitch), I (eye), L (ell), M (em),
	// N (en), O (oh), R (ar), S (ess), X (ex)
	vowelSoundLetters := "AEFHILMNORSX"
	first := unicode.ToUpper(rune(abbrev[0]))
	return strings.ContainsRune(vowelSoundLetters, first)
}

// isConsonantYSound checks if a word starting with 'u' has a consonant "y" sound.
// e.g., "unanimous" starts with "yoo-" sound.
func isConsonantYSound(lower string) bool {
	if len(lower) < 2 {
		return false
	}

	// Explicit patterns that have the "you" sound
	youPatterns := []string{
		"uga", "ukr", "ula", "ule", "uli", "ulo", "ulu",
		"una", "uni", "uno", "unu", "ura", "ure", "uri", "uro", "uru",
		"usa", "use", "usi", "uso", "usu", "uta", "ute", "uti", "uto", "utu",
	}

	if len(lower) >= 3 {
		prefix := lower[:3]
		for _, pat := range youPatterns {
			if prefix == pat {
				return true
			}
		}
	}

	return false
}

// isVowelSound checks if a letter represents a vowel sound.
func isVowelSound(r rune) bool {
	return strings.ContainsRune("aeiou", unicode.ToLower(r))
}

// A is an alias for An - returns word prefixed with appropriate indefinite article.
func A(word string) string {
	return An(word)
}

// Plural returns the plural form of an English noun.
//
// Examples:
//   - Plural("cat") returns "cats"
//   - Plural("box") returns "boxes"
//   - Plural("child") returns "children"
//   - Plural("sheep") returns "sheep"
func Plural(word string) string {
	if word == "" {
		return ""
	}

	lower := strings.ToLower(word)

	// Check for irregular plurals first
	if plural, ok := irregularPlurals[lower]; ok {
		return matchCase(word, plural)
	}

	// Check for uncountable/unchanged words
	for _, unchanged := range unchangedPlurals {
		if lower == unchanged {
			return word
		}
	}

	// Check for words ending in -ese, -ois (nationalities that don't change)
	if strings.HasSuffix(lower, "ese") || strings.HasSuffix(lower, "ois") {
		return word
	}

	// Apply suffix rules
	return applySuffixRules(word, lower)
}

// applySuffixRules applies standard English pluralization suffix rules.
func applySuffixRules(word, lower string) string {
	// Words ending in -man -> -men
	if strings.HasSuffix(lower, "man") && !strings.HasSuffix(lower, "human") {
		return word[:len(word)-3] + matchCase(word[len(word)-3:], "men")
	}

	// Words ending in -s, -ss, -sh, -ch, -x, -z -> add -es
	if strings.HasSuffix(lower, "s") || strings.HasSuffix(lower, "ss") ||
		strings.HasSuffix(lower, "sh") || strings.HasSuffix(lower, "ch") ||
		strings.HasSuffix(lower, "x") || strings.HasSuffix(lower, "z") {
		return word + matchSuffix(word, "es")
	}

	// Words ending in consonant + y -> -ies
	if strings.HasSuffix(lower, "y") && len(lower) > 1 {
		beforeY := lower[len(lower)-2]
		if !isVowel(rune(beforeY)) {
			return word[:len(word)-1] + matchSuffix(word, "ies")
		}
	}

	// Words ending in -f or -fe -> -ves (with exceptions)
	if strings.HasSuffix(lower, "fe") {
		if shouldChangeF(lower) {
			return word[:len(word)-2] + matchSuffix(word, "ves")
		}
	} else if strings.HasSuffix(lower, "f") && !strings.HasSuffix(lower, "ff") {
		if shouldChangeF(lower) {
			return word[:len(word)-1] + matchSuffix(word, "ves")
		}
	}

	// Words ending in -o -> -oes (with exceptions)
	if strings.HasSuffix(lower, "o") && len(lower) > 1 {
		beforeO := lower[len(lower)-2]
		// Vowel + o -> just add s (radio, studio, zoo)
		if isVowel(rune(beforeO)) {
			return word + matchSuffix(word, "s")
		}
		// Check if it's an exception that just takes -s
		if oExceptionTakesS(lower) {
			return word + matchSuffix(word, "s")
		}
		return word + matchSuffix(word, "es")
	}

	// Default: add -s
	return word + matchSuffix(word, "s")
}

// matchSuffix returns the suffix in uppercase if the word is all uppercase.
func matchSuffix(word, suffix string) string {
	allUpper := true
	for _, r := range word {
		if unicode.IsLetter(r) && !unicode.IsUpper(r) {
			allUpper = false
			break
		}
	}
	if allUpper {
		return strings.ToUpper(suffix)
	}
	return suffix
}

// isVowel checks if a rune is a vowel.
func isVowel(r rune) bool {
	return strings.ContainsRune("aeiouAEIOU", r)
}

// shouldChangeF determines if a word ending in -f/-fe should change to -ves.
func shouldChangeF(lower string) bool {
	// Words that change -f/-fe to -ves
	changeToVes := []string{
		"calf", "elf", "half", "knife", "leaf", "life", "loaf",
		"self", "sheaf", "shelf", "thief", "wife", "wolf",
	}
	for _, w := range changeToVes {
		if lower == w {
			return true
		}
	}
	return false
}

// oExceptionTakesS returns true if a word ending in -o just takes -s.
func oExceptionTakesS(lower string) bool {
	// Musical terms, abbreviations, and other exceptions
	exceptions := []string{
		"alto", "auto", "basso", "canto", "casino", "combo", "contralto",
		"disco", "dynamo", "embryo", "espresso", "euro", "fiasco", "ghetto",
		"inferno", "kilo", "limo", "maestro", "memo", "metro", "piano",
		"photo", "pimento", "polo", "poncho", "pro", "ratio", "rhino",
		"silo", "solo", "soprano", "stiletto", "studio", "taco", "tattoo",
		"tempo", "tornado", "torso", "tuxedo", "video", "virtuoso", "zero",
		"albino", "archipelago", "armadillo", "commando", "dodo", "flamingo",
		"grotto", "magneto", "manifesto", "mosquito", "motto", "otto",
		"placebo", "portfolio", "quarto", "stucco", "tobacco", "volcano",
	}
	for _, w := range exceptions {
		if lower == w {
			return true
		}
	}
	return false
}

// matchCase adjusts the replacement to match the case pattern of the original.
func matchCase(original, replacement string) string {
	if len(original) == 0 || len(replacement) == 0 {
		return replacement
	}

	// Check if original is all uppercase
	allUpper := true
	for _, r := range original {
		if unicode.IsLetter(r) && !unicode.IsUpper(r) {
			allUpper = false
			break
		}
	}
	if allUpper {
		return strings.ToUpper(replacement)
	}

	// Check if original starts with uppercase
	firstRune := []rune(original)[0]
	if unicode.IsUpper(firstRune) {
		runes := []rune(replacement)
		runes[0] = unicode.ToUpper(runes[0])
		return string(runes)
	}

	return replacement
}

// irregularPlurals maps singular forms to their irregular plural forms.
var irregularPlurals = map[string]string{
	"child":       "children",
	"foot":        "feet",
	"goose":       "geese",
	"louse":       "lice",
	"man":         "men",
	"mouse":       "mice",
	"ox":          "oxen",
	"person":      "people",
	"tooth":       "teeth",
	"woman":       "women",
	"die":         "dice",
	"criterion":   "criteria",
	"phenomenon":  "phenomena",
	"analysis":    "analyses",
	"basis":       "bases",
	"crisis":      "crises",
	"diagnosis":   "diagnoses",
	"hypothesis":  "hypotheses",
	"oasis":       "oases",
	"parenthesis": "parentheses",
	"synopsis":    "synopses",
	"thesis":      "theses",
	"alumnus":     "alumni",
	"cactus":      "cacti",
	"focus":       "foci",
	"fungus":      "fungi",
	"nucleus":     "nuclei",
	"radius":      "radii",
	"stimulus":    "stimuli",
	"syllabus":    "syllabi",
	"bacterium":   "bacteria",
	"curriculum":  "curricula",
	"datum":       "data",
	"medium":      "media",
	"memorandum":  "memoranda",
	"millennium":  "millennia",
	"stadium":     "stadia",
	"stratum":     "strata",
	"appendix":    "appendices",
	"index":       "indices",
	"matrix":      "matrices",
	"vertex":      "vertices",
	"apex":        "apices",
}

// unchangedPlurals lists words that don't change in plural form.
var unchangedPlurals = []string{
	"aircraft", "bison", "buffalo", "cod", "deer", "fish", "moose",
	"offspring", "pike", "salmon", "series", "sheep", "shrimp",
	"species", "squid", "swine", "trout", "tuna",
}

// Ordinal converts an integer to its ordinal string representation.
//
// Examples:
//   - Ordinal(1) returns "1st"
//   - Ordinal(2) returns "2nd"
//   - Ordinal(3) returns "3rd"
//   - Ordinal(11) returns "11th"
//   - Ordinal(21) returns "21st"
//   - Ordinal(-1) returns "-1st"
func Ordinal(n int) string {
	suffix := ordinalSuffix(n)
	return fmt.Sprintf("%d%s", n, suffix)
}

// ordinalSuffix returns the ordinal suffix for a number.
func ordinalSuffix(n int) string {
	// Handle negative numbers by using absolute value
	if n < 0 {
		n = -n
	}

	// Special case: 11, 12, 13 always use "th"
	// Check the last two digits to handle 111, 112, 113, etc.
	lastTwo := n % 100
	if lastTwo >= 11 && lastTwo <= 13 {
		return "th"
	}

	// Otherwise, check the last digit
	switch n % 10 {
	case 1:
		return "st"
	case 2:
		return "nd"
	case 3:
		return "rd"
	default:
		return "th"
	}
}

// OrdinalWord converts an integer to its ordinal word representation.
//
// Examples:
//   - OrdinalWord(1) returns "first"
//   - OrdinalWord(2) returns "second"
//   - OrdinalWord(11) returns "eleventh"
//   - OrdinalWord(21) returns "twenty-first"
//   - OrdinalWord(100) returns "one hundredth"
//   - OrdinalWord(101) returns "one hundred first"
//   - OrdinalWord(-1) returns "negative first"
func OrdinalWord(n int) string {
	if n == 0 {
		return "zeroth"
	}

	if n < 0 {
		return "negative " + OrdinalWord(-n)
	}

	return convertToOrdinalWord(n)
}

// convertToOrdinalWord converts a positive integer to its ordinal word form.
func convertToOrdinalWord(n int) string {
	// Handle numbers 1-19 with direct lookup
	if n <= 19 {
		return onesOrdinal[n]
	}

	// Handle exact tens (20, 30, 40, ...)
	if n < 100 && n%10 == 0 {
		return tensOrdinal[n/10]
	}

	// Handle 20-99 with compound form (twenty-first, etc.)
	if n < 100 {
		return tensCardinal[n/10] + "-" + onesOrdinal[n%10]
	}

	// Handle exact hundreds (100, 200, ...)
	if n < 1000 && n%100 == 0 {
		return onesCardinal[n/100] + " hundredth"
	}

	// Handle 100-999
	if n < 1000 {
		return onesCardinal[n/100] + " hundred " + convertToOrdinalWord(n%100)
	}

	// Handle exact thousands (1000, 2000, ...)
	if n < 1000000 && n%1000 == 0 {
		return cardinalWord(n/1000) + " thousandth"
	}

	// Handle 1000-999999
	if n < 1000000 {
		return cardinalWord(n/1000) + " thousand " + convertToOrdinalWord(n%1000)
	}

	// Handle exact millions (1000000, 2000000, ...)
	if n < 1000000000 && n%1000000 == 0 {
		return cardinalWord(n/1000000) + " millionth"
	}

	// Handle 1000000-999999999
	if n < 1000000000 {
		return cardinalWord(n/1000000) + " million " + convertToOrdinalWord(n%1000000)
	}

	// Handle exact billions
	if n%1000000000 == 0 {
		return cardinalWord(n/1000000000) + " billionth"
	}

	// Handle billions and above
	return cardinalWord(n/1000000000) + " billion " + convertToOrdinalWord(n%1000000000)
}

// cardinalWord converts a positive integer to its cardinal word form.
func cardinalWord(n int) string {
	if n == 0 {
		return "zero"
	}

	if n <= 19 {
		return onesCardinal[n]
	}

	if n < 100 && n%10 == 0 {
		return tensCardinal[n/10]
	}

	if n < 100 {
		return tensCardinal[n/10] + "-" + onesCardinal[n%10]
	}

	if n < 1000 && n%100 == 0 {
		return onesCardinal[n/100] + " hundred"
	}

	if n < 1000 {
		return onesCardinal[n/100] + " hundred " + cardinalWord(n%100)
	}

	if n < 1000000 && n%1000 == 0 {
		return cardinalWord(n/1000) + " thousand"
	}

	if n < 1000000 {
		return cardinalWord(n/1000) + " thousand " + cardinalWord(n%1000)
	}

	if n < 1000000000 && n%1000000 == 0 {
		return cardinalWord(n/1000000) + " million"
	}

	if n < 1000000000 {
		return cardinalWord(n/1000000) + " million " + cardinalWord(n%1000000)
	}

	if n%1000000000 == 0 {
		return cardinalWord(n/1000000000) + " billion"
	}

	return cardinalWord(n/1000000000) + " billion " + cardinalWord(n%1000000000)
}

// onesCardinal maps 1-19 to their cardinal word forms.
var onesCardinal = []string{
	"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen",
	"seventeen", "eighteen", "nineteen",
}

// onesOrdinal maps 1-19 to their ordinal word forms.
var onesOrdinal = []string{
	"", "first", "second", "third", "fourth", "fifth", "sixth", "seventh",
	"eighth", "ninth", "tenth", "eleventh", "twelfth", "thirteenth",
	"fourteenth", "fifteenth", "sixteenth", "seventeenth", "eighteenth",
	"nineteenth",
}

// tensCardinal maps tens (2-9 representing 20-90) to their cardinal word forms.
var tensCardinal = []string{
	"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety",
}

// tensOrdinal maps tens (2-9 representing 20-90) to their ordinal word forms.
var tensOrdinal = []string{
	"", "", "twentieth", "thirtieth", "fortieth", "fiftieth", "sixtieth",
	"seventieth", "eightieth", "ninetieth",
}

// Join combines a slice of strings into a grammatically correct English list.
//
// The function uses the Oxford comma (serial comma) for lists of three or more items.
// It uses "and" as the conjunction. For custom conjunctions, use JoinWithConj.
//
// Examples:
//   - Join([]string{}) returns ""
//   - Join([]string{"a"}) returns "a"
//   - Join([]string{"a", "b"}) returns "a and b"
//   - Join([]string{"a", "b", "c"}) returns "a, b, and c"
func Join(words []string) string {
	return JoinWithConj(words, "and")
}

// JoinWithConj combines a slice of strings into a grammatically correct English list
// with a custom conjunction.
//
// The function uses the Oxford comma (serial comma) for lists of three or more items.
//
// Examples:
//   - JoinWithConj([]string{"a", "b"}, "or") returns "a or b"
//   - JoinWithConj([]string{"a", "b", "c"}, "or") returns "a, b, or c"
//   - JoinWithConj([]string{"a", "b", "c"}, "and/or") returns "a, b, and/or c"
func JoinWithConj(words []string, conj string) string {
	switch len(words) {
	case 0:
		return ""
	case 1:
		return words[0]
	case 2:
		return words[0] + " " + conj + " " + words[1]
	default:
		return strings.Join(words[:len(words)-1], ", ") + ", " + conj + " " + words[len(words)-1]
	}
}
