package inflect

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

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

// wordZero is the word representation of zero.
const wordZero = "zero"

// NumberToWords converts an integer to its English word representation.
//
// Examples:
//   - NumberToWords(0) returns "zero"
//   - NumberToWords(1) returns "one"
//   - NumberToWords(42) returns "forty-two"
//   - NumberToWords(100) returns "one hundred"
//   - NumberToWords(1000) returns "one thousand"
//   - NumberToWords(-5) returns "negative five"
func NumberToWords(n int) string {
	if n < 0 {
		return "negative " + cardinalWord(-n)
	}
	return cardinalWord(n)
}

// NumberToWordsWithAnd converts an integer to its English word representation
// using British English style with "and" before the final part.
//
// This style inserts "and" after hundreds when followed by tens or ones,
// and after thousands/millions/billions when followed by a number less than 100.
//
// Examples:
//   - NumberToWordsWithAnd(101) returns "one hundred and one"
//   - NumberToWordsWithAnd(121) returns "one hundred and twenty-one"
//   - NumberToWordsWithAnd(1001) returns "one thousand and one"
//   - NumberToWordsWithAnd(1101) returns "one thousand one hundred and one"
//   - NumberToWordsWithAnd(-101) returns "negative one hundred and one"
func NumberToWordsWithAnd(n int) string {
	if n < 0 {
		return "negative " + cardinalWordWithAnd(-n)
	}
	return cardinalWordWithAnd(n)
}

// cardinalWordWithAnd converts a positive integer to its cardinal word form
// using British English style with "and".
func cardinalWordWithAnd(n int) string {
	if n == 0 {
		return wordZero
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
		// Insert "and" after hundreds
		return onesCardinal[n/100] + " hundred and " + cardinalWordWithAnd(n%100)
	}

	if n < 1000000 && n%1000 == 0 {
		return cardinalWordWithAnd(n/1000) + " thousand"
	}

	if n < 1000000 {
		remainder := n % 1000
		prefix := cardinalWordWithAnd(n/1000) + " thousand"
		// If remainder is less than 100, add "and"
		if remainder < 100 {
			return prefix + " and " + cardinalWordWithAnd(remainder)
		}
		return prefix + " " + cardinalWordWithAnd(remainder)
	}

	if n < 1000000000 && n%1000000 == 0 {
		return cardinalWordWithAnd(n/1000000) + " million"
	}

	if n < 1000000000 {
		remainder := n % 1000000
		prefix := cardinalWordWithAnd(n/1000000) + " million"
		// If remainder is less than 100, add "and"
		if remainder < 100 {
			return prefix + " and " + cardinalWordWithAnd(remainder)
		}
		return prefix + " " + cardinalWordWithAnd(remainder)
	}

	if n%1000000000 == 0 {
		return cardinalWordWithAnd(n/1000000000) + " billion"
	}

	remainder := n % 1000000000
	prefix := cardinalWordWithAnd(n/1000000000) + " billion"
	// If remainder is less than 100, add "and"
	if remainder < 100 {
		return prefix + " and " + cardinalWordWithAnd(remainder)
	}
	return prefix + " " + cardinalWordWithAnd(remainder)
}

// NumberToWordsFloat converts a floating-point number to its English word representation.
//
// The integer part is converted using NumberToWords, followed by "point",
// then each digit after the decimal point is converted individually.
//
// Examples:
//   - NumberToWordsFloat(3.14) returns "three point one four"
//   - NumberToWordsFloat(0.5) returns "zero point five"
//   - NumberToWordsFloat(-2.718) returns "negative two point seven one eight"
func NumberToWordsFloat(f float64) string {
	return NumberToWordsFloatWithDecimal(f, "point")
}

// NumberToWordsFloatWithDecimal converts a floating-point number to its English word representation
// using a custom word for the decimal point.
//
// The integer part is converted using NumberToWords, followed by the specified decimal word,
// then each digit after the decimal point is converted individually.
//
// Examples:
//   - NumberToWordsFloatWithDecimal(3.14, "point") returns "three point one four"
//   - NumberToWordsFloatWithDecimal(3.14, "dot") returns "three dot one four"
//   - NumberToWordsFloatWithDecimal(3.14, "and") returns "three and one four"
func NumberToWordsFloatWithDecimal(f float64, decimal string) string {
	// Handle negative numbers
	prefix := ""
	if f < 0 {
		prefix = "negative "
		f = math.Abs(f)
	}

	// Get integer part
	intPart := int(f)

	// Convert to string to extract decimal digits
	str := strconv.FormatFloat(f, 'f', -1, 64)

	// Find the decimal point
	dotIdx := strings.Index(str, ".")
	if dotIdx == -1 {
		// No decimal point, just convert as integer
		return prefix + cardinalWord(intPart)
	}

	// Get decimal digits
	decimalDigits := str[dotIdx+1:]

	// Build the result: integer word + "point" + each decimal digit word
	parts := make([]string, 0, 2+len(decimalDigits))
	parts = append(parts, prefix+cardinalWord(intPart), decimal)

	// Convert each decimal digit individually
	digitWords := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for _, ch := range decimalDigits {
		digit := int(ch - '0')
		parts = append(parts, digitWords[digit])
	}

	return strings.Join(parts, " ")
}

// NumberToWordsThreshold converts an integer to its English word representation
// only if the number is below the specified threshold. If the number is greater
// than or equal to the threshold, it returns the number as a string.
//
// This is useful for making text more readable by spelling out small numbers
// while keeping larger numbers in digit form.
//
// Examples:
//   - NumberToWordsThreshold(5, 10) returns "five" (5 < 10, convert to words)
//   - NumberToWordsThreshold(15, 10) returns "15" (15 >= 10, return as string)
//   - NumberToWordsThreshold(100, 100) returns "100" (100 >= 100, return as string)
//   - NumberToWordsThreshold(-3, 10) returns "negative three" (-3 < 10, convert to words)
func NumberToWordsThreshold(n, threshold int) string {
	if n < threshold {
		return NumberToWords(n)
	}
	return strconv.Itoa(n)
}

// NumberToWordsGrouped converts an integer to English words by splitting it into
// groups of the specified size and converting each group independently.
//
// This is useful for reading phone numbers, credit card numbers, and other
// digit sequences where each group should be pronounced as a separate number.
//
// The number is split from right to left, so the leftmost group may have
// fewer digits than the specified group size.
//
// Examples:
//   - NumberToWordsGrouped(1234, 2) returns "twelve thirty-four"
//   - NumberToWordsGrouped(123456, 2) returns "twelve thirty-four fifty-six"
//   - NumberToWordsGrouped(1234, 3) returns "one two hundred thirty-four"
//   - NumberToWordsGrouped(1234567890, 3) returns "one two hundred thirty-four five hundred sixty-seven eight hundred ninety"
//   - NumberToWordsGrouped(0, 2) returns "zero"
//   - NumberToWordsGrouped(-1234, 2) returns "negative twelve thirty-four"
func NumberToWordsGrouped(n, groupSize int) string {
	if groupSize <= 0 {
		return NumberToWords(n)
	}

	// Handle negative numbers
	prefix := ""
	if n < 0 {
		prefix = "negative "
		n = -n
	}

	// Handle zero
	if n == 0 {
		return wordZero
	}

	// Convert number to string
	s := strconv.Itoa(n)

	// Split into groups from right to left
	var groups []string
	for s != "" {
		start := max(len(s)-groupSize, 0)
		groups = append([]string{s[start:]}, groups...)
		s = s[:start]
	}

	// Convert each group to words
	words := make([]string, 0, len(groups))
	for _, g := range groups {
		num, _ := strconv.Atoi(g)
		words = append(words, cardinalWord(num))
	}

	return prefix + strings.Join(words, " ")
}

// cardinalWord converts a positive integer to its cardinal word form.
func cardinalWord(n int) string {
	if n == 0 {
		return wordZero
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

// FormatNumber formats an integer with commas as thousand separators.
//
// Examples:
//   - FormatNumber(1000) returns "1,000"
//   - FormatNumber(1000000) returns "1,000,000"
//   - FormatNumber(123456789) returns "123,456,789"
//   - FormatNumber(-1234) returns "-1,234"
//   - FormatNumber(999) returns "999" (no comma needed)
func FormatNumber(n int) string {
	// Handle negative numbers
	if n < 0 {
		return "-" + FormatNumber(-n)
	}

	// Convert to string
	s := strconv.Itoa(n)

	// No formatting needed for numbers with 3 or fewer digits
	if len(s) <= 3 {
		return s
	}

	// Build result with commas inserted every 3 digits from the right
	var result strings.Builder
	result.Grow(len(s) + (len(s)-1)/3) // pre-allocate space for digits + commas

	// Calculate the position of the first comma
	firstGroup := len(s) % 3
	if firstGroup == 0 {
		firstGroup = 3
	}

	// Write first group (1-3 digits)
	result.WriteString(s[:firstGroup])

	// Write remaining groups with preceding commas
	for i := firstGroup; i < len(s); i += 3 {
		result.WriteByte(',')
		result.WriteString(s[i : i+3])
	}

	return result.String()
}

// No returns a count and noun phrase in English, using "no" for zero counts.
//
// The function handles pluralization automatically:
//   - For count 0 with classicalZero=false (default): returns "no" + plural form
//   - For count 0 with classicalZero=true: returns "no" + singular form
//   - For count 1: returns "1" + singular form
//   - For count > 1: returns count + plural form
//
// Examples:
//   - No("error", 0) returns "no errors" (default)
//   - No("error", 1) returns "1 error"
//   - No("error", 2) returns "2 errors"
//   - No("child", 0) returns "no children" (default)
//   - No("child", 1) returns "1 child"
//   - No("child", 3) returns "3 children"
//
// With ClassicalZero(true):
//   - No("error", 0) returns "no error"
//   - No("child", 0) returns "no child"
func No(word string, count int) string {
	if count == 0 {
		if defaultEngine.IsClassicalZero() {
			return "no " + word
		}
		return "no " + Plural(word)
	}
	if count == 1 || count == -1 {
		return fmt.Sprintf("%d %s", count, word)
	}
	return fmt.Sprintf("%d %s", count, Plural(word))
}

// defaultNum stores the default count for number-related functions.
// A value of 0 indicates no default is set.
var defaultNum int

// Num stores and retrieves a default count for number-related operations.
//
// When called with a positive integer, it stores that value as the default
// count and returns it. When called with 0 or no arguments, it clears the
// default count and returns 0.
//
// Examples:
//   - Num(5) stores 5 as default count, returns 5
//   - Num(0) clears the default count, returns 0
//   - Num() clears the default count, returns 0
func Num(n ...int) int {
	if len(n) == 0 || n[0] == 0 {
		defaultNum = 0
		return 0
	}
	defaultNum = n[0]
	return defaultNum
}

// GetNum retrieves the current default count.
//
// Returns 0 if no default has been set or if it was cleared.
//
// Examples:
//   - After Num(5): GetNum() returns 5
//   - After Num(0) or Num(): GetNum() returns 0
//   - Before any Num() call: GetNum() returns 0
func GetNum() int {
	return defaultNum
}
