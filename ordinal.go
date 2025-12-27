package inflect

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

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

// cardinalToOrdinal maps cardinal word forms to ordinal word forms.
var cardinalToOrdinal = map[string]string{
	"zero":      "zeroth",
	"one":       "first",
	"two":       "second",
	"three":     "third",
	"four":      "fourth",
	"five":      "fifth",
	"six":       "sixth",
	"seven":     "seventh",
	"eight":     "eighth",
	"nine":      "ninth",
	"ten":       "tenth",
	"eleven":    "eleventh",
	"twelve":    "twelfth",
	"thirteen":  "thirteenth",
	"fourteen":  "fourteenth",
	"fifteen":   "fifteenth",
	"sixteen":   "sixteenth",
	"seventeen": "seventeenth",
	"eighteen":  "eighteenth",
	"nineteen":  "nineteenth",
	"twenty":    "twentieth",
	"thirty":    "thirtieth",
	"forty":     "fortieth",
	"fifty":     "fiftieth",
	"sixty":     "sixtieth",
	"seventy":   "seventieth",
	"eighty":    "eightieth",
	"ninety":    "ninetieth",
}

// WordToOrdinal converts a number word or numeric string to its ordinal form.
//
// If the input is a numeric string (e.g., "42"), it returns the numeric ordinal (e.g., "42nd").
// If the input is a word number (e.g., "forty-two"), it returns the word ordinal (e.g., "forty-second").
//
// The function preserves the case pattern of the input:
//   - "one" → "first"
//   - "One" → "First"
//   - "ONE" → "FIRST"
//   - "Twenty-One" → "Twenty-First"
//
// Examples:
//   - WordToOrdinal("1") returns "1st"
//   - WordToOrdinal("one") returns "first"
//   - WordToOrdinal("twenty-one") returns "twenty-first"
//   - WordToOrdinal("One") returns "First"
//   - WordToOrdinal("TWENTY") returns "TWENTIETH"
func WordToOrdinal(s string) string {
	// Try to parse as a number first
	if n, err := strconv.Atoi(s); err == nil {
		return Ordinal(n)
	}

	// Detect case pattern
	casePattern := detectCase(s)

	// Convert to lowercase for lookup
	lower := strings.ToLower(s)

	// Handle compound numbers (e.g., "twenty-one" → "twenty-first")
	if idx := strings.LastIndex(lower, "-"); idx >= 0 {
		prefix := lower[:idx]
		suffix := lower[idx+1:]

		// Check if suffix is already ordinal
		if strings.HasSuffix(suffix, "th") || strings.HasSuffix(suffix, "st") ||
			strings.HasSuffix(suffix, "nd") || strings.HasSuffix(suffix, "rd") {
			return s // Already ordinal
		}

		// Convert the suffix to ordinal
		if ordinal, ok := cardinalToOrdinal[suffix]; ok {
			result := prefix + "-" + ordinal
			return applyCase(result, casePattern)
		}
	}

	// Direct lookup for simple words
	if ordinal, ok := cardinalToOrdinal[lower]; ok {
		return applyCase(ordinal, casePattern)
	}

	// Return unchanged if not recognized
	return s
}

// casePattern represents the case pattern of a string.
type casePattern int

const (
	caseLower casePattern = iota
	caseUpper
	caseTitle
	caseMixed
)

// detectCase detects the case pattern of a string.
func detectCase(s string) casePattern {
	if s == "" {
		return caseLower
	}

	// Check if all uppercase
	allUpper := true
	allLower := true

	for _, r := range s {
		if unicode.IsLetter(r) {
			if !unicode.IsUpper(r) {
				allUpper = false
			}
			if !unicode.IsLower(r) {
				allLower = false
			}
		}
	}

	if allUpper {
		return caseUpper
	}
	if allLower {
		return caseLower
	}

	// Check for title case (first letter uppercase, rest lowercase per word)
	runes := []rune(s)
	if unicode.IsUpper(runes[0]) {
		return caseTitle
	}

	return caseMixed
}

// applyCase applies a case pattern to a string.
func applyCase(s string, pattern casePattern) string {
	switch pattern {
	case caseUpper:
		return strings.ToUpper(s)
	case caseTitle:
		return toTitleCase(s)
	default:
		return s
	}
}

// toTitleCase converts a string to title case (first letter of each word uppercase).
func toTitleCase(s string) string {
	runes := []rune(s)
	if len(runes) == 0 {
		return s
	}

	// Capitalize first letter
	runes[0] = unicode.ToUpper(runes[0])

	// Handle hyphenated words
	for i := 1; i < len(runes); i++ {
		if i > 0 && runes[i-1] == '-' {
			runes[i] = unicode.ToUpper(runes[i])
		}
	}

	return string(runes)
}
