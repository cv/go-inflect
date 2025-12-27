package inflect

import "fmt"

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
