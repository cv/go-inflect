package inflect

// FractionToWords converts a fraction to its English word representation.
//
// Special cases are handled as follows:
//   - Denominator 2: uses "half/halves"
//   - Denominator 4: uses "quarter/quarters"
//   - Other denominators: uses ordinal form (third, fifth, eighth, etc.)
//   - Numerator 1: singular form (one third)
//   - Numerator > 1: plural form (two thirds)
//   - Denominator 1: returns just the numerator in words
//   - Denominator 0: returns empty string (invalid fraction)
//
// Negative numbers are handled by prefixing "negative" when the overall
// fraction is negative (exactly one of numerator or denominator is negative).
//
// Examples:
//   - FractionToWords(1, 2) returns "one half"
//   - FractionToWords(3, 2) returns "three halves"
//   - FractionToWords(1, 4) returns "one quarter"
//   - FractionToWords(3, 4) returns "three quarters"
//   - FractionToWords(2, 3) returns "two thirds"
//   - FractionToWords(5, 8) returns "five eighths"
//   - FractionToWords(1, 100) returns "one hundredth"
//   - FractionToWords(-1, 2) returns "negative one half"
func FractionToWords(numerator, denominator int) string {
	return fractionToWordsInternal(numerator, denominator, true)
}

// FractionToWordsWithFourths converts a fraction to its English word representation,
// using "fourth/fourths" instead of "quarter/quarters" for denominator 4.
//
// This is an alternative style that some prefer for mathematical contexts.
//
// Examples:
//   - FractionToWordsWithFourths(1, 4) returns "one fourth"
//   - FractionToWordsWithFourths(3, 4) returns "three fourths"
func FractionToWordsWithFourths(numerator, denominator int) string {
	return fractionToWordsInternal(numerator, denominator, false)
}

// fractionToWordsInternal is the internal implementation that handles both
// quarter and fourth styles for denominator 4.
func fractionToWordsInternal(numerator, denominator int, useQuarters bool) string {
	// Handle invalid denominator
	if denominator == 0 {
		return ""
	}

	// Determine sign
	negative := false
	if numerator < 0 {
		negative = !negative
		numerator = -numerator
	}
	if denominator < 0 {
		negative = !negative
		denominator = -denominator
	}

	// Handle denominator 1 (whole numbers)
	if denominator == 1 {
		result := NumberToWords(numerator)
		if negative {
			return "negative " + result
		}
		return result
	}

	// Get numerator word
	numeratorWord := NumberToWords(numerator)

	// Determine if we need plural form (numerator != 1)
	plural := numerator != 1

	// Get denominator word based on special cases
	var denominatorWord string
	switch denominator {
	case 2:
		if plural {
			denominatorWord = "halves"
		} else {
			denominatorWord = "half"
		}
	case 4:
		denominatorWord = denominatorFour(useQuarters, plural)
	default:
		// Use ordinal form for other denominators
		ordinal := ordinalDenominator(denominator)
		if plural {
			denominatorWord = pluralizeOrdinal(ordinal)
		} else {
			denominatorWord = ordinal
		}
	}

	result := numeratorWord + " " + denominatorWord
	if negative {
		return "negative " + result
	}
	return result
}

// pluralizeOrdinal adds 's' to an ordinal word to make it plural.
// Handles compound ordinals like "twenty-first" → "twenty-firsts".
func pluralizeOrdinal(ordinal string) string {
	return ordinal + "s"
}

// denominatorFour returns the denominator word for 4 (quarter or fourth).
func denominatorFour(useQuarters, plural bool) string {
	if useQuarters {
		if plural {
			return "quarters"
		}
		return "quarter"
	}
	if plural {
		return "fourths"
	}
	return "fourth"
}

// ordinalDenominator returns the ordinal form of a number suitable for use as
// a fraction denominator. Unlike OrdinalWord, this returns just the ordinal
// suffix for exact powers (e.g., 100 → "hundredth", not "one hundredth").
func ordinalDenominator(n int) string {
	// Special cases for exact powers - return just the ordinal suffix
	if n == 100 {
		return "hundredth"
	}
	if n == 1000 {
		return "thousandth"
	}
	if n == 1000000 {
		return "millionth"
	}
	if n == 1000000000 {
		return "billionth"
	}

	// For other numbers, use the standard ordinal word
	return OrdinalWord(n)
}
