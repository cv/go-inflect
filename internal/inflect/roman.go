package inflect

import (
	"errors"
	"strings"
)

// romanNumerals maps values to their Roman numeral representations.
// Ordered from largest to smallest for conversion.
var romanNumerals = []struct {
	value  int
	symbol string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

// romanValues maps Roman numeral characters to their values.
var romanValues = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

// validSubtractive maps each character to its valid subtractive targets.
var validSubtractive = map[byte]string{
	'I': "VX",
	'X': "LC",
	'C': "DM",
}

// ErrInvalidRoman is returned when a Roman numeral string is malformed.
var ErrInvalidRoman = errors.New("invalid Roman numeral")

// IntToRoman converts an integer to its Roman numeral representation.
//
// Roman numerals are only defined for integers from 1 to 3999.
// For values outside this range, an empty string is returned.
//
// Examples:
//   - IntToRoman(4) returns "IV"
//   - IntToRoman(9) returns "IX"
//   - IntToRoman(1984) returns "MCMLXXXIV"
//   - IntToRoman(2025) returns "MMXXV"
//   - IntToRoman(0) returns ""
//   - IntToRoman(4000) returns ""
func IntToRoman(n int) string {
	if n < 1 || n > 3999 {
		return ""
	}

	var result strings.Builder
	for _, rn := range romanNumerals {
		for n >= rn.value {
			result.WriteString(rn.symbol)
			n -= rn.value
		}
	}
	return result.String()
}

// RomanToInt converts a Roman numeral string to its integer value.
//
// The function accepts both uppercase and lowercase input.
// It validates the input and returns an error for malformed numerals.
//
// Validation rules:
//   - Only valid Roman numeral characters (I, V, X, L, C, D, M)
//   - No more than 3 consecutive identical numerals (except M)
//   - V, L, D cannot repeat
//   - Valid subtractive combinations only (IV, IX, XL, XC, CD, CM)
//
// Examples:
//   - RomanToInt("XIV") returns (14, nil)
//   - RomanToInt("MMXXV") returns (2025, nil)
//   - RomanToInt("iv") returns (4, nil)
//   - RomanToInt("IIII") returns (0, error)
//   - RomanToInt("ABC") returns (0, error)
func RomanToInt(s string) (int, error) {
	if s == "" {
		return 0, ErrInvalidRoman
	}

	// Convert to uppercase for processing
	s = strings.ToUpper(s)

	// Validate characters and structure
	if err := validateRoman(s); err != nil {
		return 0, err
	}

	// Calculate value using subtractive notation
	return calculateRomanValue(s), nil
}

// calculateRomanValue computes the integer value of a validated Roman numeral.
func calculateRomanValue(s string) int {
	total := 0
	for i := range len(s) {
		curr := romanValues[s[i]]
		if i+1 < len(s) && curr < romanValues[s[i+1]] {
			total -= curr
		} else {
			total += curr
		}
	}
	return total
}

// IntToRoman converts an integer to its Roman numeral representation.
func (e *Engine) IntToRoman(n int) string {
	return IntToRoman(n)
}

// RomanToInt converts a Roman numeral string to its integer value.
func (e *Engine) RomanToInt(s string) (int, error) {
	return RomanToInt(s)
}

// validateRoman checks if a Roman numeral string follows valid formation rules.
func validateRoman(s string) error {
	if err := validateCharacters(s); err != nil {
		return err
	}
	if err := validateRepetitions(s); err != nil {
		return err
	}
	return validateSubtractions(s)
}

// validateCharacters ensures all characters are valid Roman numerals.
func validateCharacters(s string) error {
	for i := range len(s) {
		if _, ok := romanValues[s[i]]; !ok {
			return ErrInvalidRoman
		}
	}
	return nil
}

// validateRepetitions checks for invalid repetition patterns.
func validateRepetitions(s string) error {
	for i := range len(s) {
		ch := s[i]

		// V, L, D cannot repeat
		if (ch == 'V' || ch == 'L' || ch == 'D') && i+1 < len(s) && s[i+1] == ch {
			return ErrInvalidRoman
		}

		// I, X, C, M can appear at most 3 times consecutively
		if ch == 'I' || ch == 'X' || ch == 'C' || ch == 'M' {
			count := 1
			for j := i + 1; j < len(s) && s[j] == ch; j++ {
				count++
			}
			if count > 3 {
				return ErrInvalidRoman
			}
		}
	}
	return nil
}

// validateSubtractions checks for invalid subtractive combinations.
func validateSubtractions(s string) error {
	for i := range len(s) - 1 {
		curr := romanValues[s[i]]
		next := romanValues[s[i+1]]

		if curr >= next {
			continue
		}

		// Check if this is a valid subtractive combination
		validTargets, canSubtract := validSubtractive[s[i]]
		if !canSubtract || !strings.ContainsRune(validTargets, rune(s[i+1])) {
			return ErrInvalidRoman
		}

		// After a subtractive pair, the next character must be smaller
		// than the smaller value of the pair (e.g., IXI is invalid)
		if i+2 < len(s) && romanValues[s[i+2]] >= curr {
			return ErrInvalidRoman
		}
	}
	return nil
}
