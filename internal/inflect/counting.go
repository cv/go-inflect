package inflect

import (
	"strconv"
)

// CountingWord converts an integer to its counting word representation.
//
// This provides frequency/repetition words:
//   - 1 returns "once"
//   - 2 returns "twice"
//   - 3 returns "thrice"
//   - 4+ returns the number word followed by "times" (e.g., "four times")
//   - 0 returns "zero times"
//   - Negative numbers are prefixed with "negative" (e.g., "negative once")
//
// Examples:
//   - CountingWord(1) returns "once"
//   - CountingWord(2) returns "twice"
//   - CountingWord(3) returns "thrice"
//   - CountingWord(4) returns "four times"
//   - CountingWord(10) returns "ten times"
//   - CountingWord(0) returns "zero times"
//   - CountingWord(-1) returns "negative once"
//   - CountingWord(-2) returns "negative twice"
func CountingWord(n int) string {
	return CountingWordWithOptions(n, true)
}

// CountingWordWithOptions converts an integer to its counting word representation
// with control over whether to use "thrice" for 3.
//
// When useThrice is true, 3 returns "thrice".
// When useThrice is false, 3 returns "three times".
//
// Examples:
//   - CountingWordWithOptions(3, true) returns "thrice"
//   - CountingWordWithOptions(3, false) returns "three times"
//   - CountingWordWithOptions(1, false) returns "once"
//   - CountingWordWithOptions(-3, true) returns "negative thrice"
//   - CountingWordWithOptions(-3, false) returns "negative three times"
func CountingWordWithOptions(n int, useThrice bool) string {
	// Handle negative numbers
	if n < 0 {
		return "negative " + countingWord(-n, useThrice)
	}
	return countingWord(n, useThrice)
}

// CountingWordThreshold converts an integer to its counting word representation
// only if the number is below the specified threshold. If the number is greater
// than or equal to the threshold, it returns the number as digits followed by "times".
//
// This is useful for making text more readable by spelling out small numbers
// while keeping larger numbers in digit form.
//
// Examples:
//   - CountingWordThreshold(5, 10) returns "five times" (5 < 10, convert to words)
//   - CountingWordThreshold(15, 10) returns "15 times" (15 >= 10, return as digits)
//   - CountingWordThreshold(1, 10) returns "once" (special word for 1)
//   - CountingWordThreshold(2, 10) returns "twice" (special word for 2)
//   - CountingWordThreshold(3, 10) returns "thrice" (special word for 3)
//   - CountingWordThreshold(100, 100) returns "100 times" (100 >= 100, return as digits)
//   - CountingWordThreshold(-3, 10) returns "negative thrice" (-3 < 10, convert to words)
func CountingWordThreshold(n, threshold int) string {
	// For special words (once, twice, thrice), always use words if below threshold
	absN := n
	if absN < 0 {
		absN = -absN
	}

	if absN < threshold {
		return CountingWord(n)
	}

	// At or above threshold: use digits
	return strconv.Itoa(n) + " times"
}

// countingWord converts a non-negative integer to its counting word form.
func countingWord(n int, useThrice bool) string {
	switch n {
	case 0:
		return "zero times"
	case 1:
		return "once"
	case 2:
		return "twice"
	case 3:
		if useThrice {
			return "thrice"
		}
		return "three times"
	default:
		return NumberToWords(n) + " times"
	}
}
