package inflect_test

import (
	"fmt"
	"testing"

	inflect "github.com/cv/go-inflect"
)

func TestNumberToWords(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  string
	}{
		// Zero
		{name: "zero", input: 0, want: "zero"},

		// Basic numbers (1-9)
		{name: "one", input: 1, want: "one"},
		{name: "two", input: 2, want: "two"},
		{name: "three", input: 3, want: "three"},
		{name: "four", input: 4, want: "four"},
		{name: "five", input: 5, want: "five"},
		{name: "six", input: 6, want: "six"},
		{name: "seven", input: 7, want: "seven"},
		{name: "eight", input: 8, want: "eight"},
		{name: "nine", input: 9, want: "nine"},

		// Teens (10-19)
		{name: "ten", input: 10, want: "ten"},
		{name: "eleven", input: 11, want: "eleven"},
		{name: "twelve", input: 12, want: "twelve"},
		{name: "thirteen", input: 13, want: "thirteen"},
		{name: "fourteen", input: 14, want: "fourteen"},
		{name: "fifteen", input: 15, want: "fifteen"},
		{name: "sixteen", input: 16, want: "sixteen"},
		{name: "seventeen", input: 17, want: "seventeen"},
		{name: "eighteen", input: 18, want: "eighteen"},
		{name: "nineteen", input: 19, want: "nineteen"},

		// Tens (20, 30, ...)
		{name: "twenty", input: 20, want: "twenty"},
		{name: "thirty", input: 30, want: "thirty"},
		{name: "forty", input: 40, want: "forty"},
		{name: "fifty", input: 50, want: "fifty"},
		{name: "sixty", input: 60, want: "sixty"},
		{name: "seventy", input: 70, want: "seventy"},
		{name: "eighty", input: 80, want: "eighty"},
		{name: "ninety", input: 90, want: "ninety"},

		// Compound tens (21-99)
		{name: "twenty-one", input: 21, want: "twenty-one"},
		{name: "thirty-two", input: 32, want: "thirty-two"},
		{name: "forty-two", input: 42, want: "forty-two"},
		{name: "fifty-five", input: 55, want: "fifty-five"},
		{name: "sixty-seven", input: 67, want: "sixty-seven"},
		{name: "seventy-eight", input: 78, want: "seventy-eight"},
		{name: "eighty-nine", input: 89, want: "eighty-nine"},
		{name: "ninety-nine", input: 99, want: "ninety-nine"},

		// Hundreds
		{name: "one hundred", input: 100, want: "one hundred"},
		{name: "two hundred", input: 200, want: "two hundred"},
		{name: "one hundred one", input: 101, want: "one hundred one"},
		{name: "one hundred ten", input: 110, want: "one hundred ten"},
		{name: "one hundred eleven", input: 111, want: "one hundred eleven"},
		{name: "one hundred twenty", input: 120, want: "one hundred twenty"},
		{name: "one hundred twenty-one", input: 121, want: "one hundred twenty-one"},
		{name: "five hundred fifty-five", input: 555, want: "five hundred fifty-five"},
		{name: "nine hundred ninety-nine", input: 999, want: "nine hundred ninety-nine"},

		// Thousands
		{name: "one thousand", input: 1000, want: "one thousand"},
		{name: "two thousand", input: 2000, want: "two thousand"},
		{name: "one thousand one", input: 1001, want: "one thousand one"},
		{name: "one thousand ten", input: 1010, want: "one thousand ten"},
		{name: "one thousand one hundred", input: 1100, want: "one thousand one hundred"},
		{name: "one thousand two hundred thirty-four", input: 1234, want: "one thousand two hundred thirty-four"},
		{name: "twelve thousand", input: 12000, want: "twelve thousand"},
		{name: "twelve thousand three hundred forty-five", input: 12345, want: "twelve thousand three hundred forty-five"},
		{name: "twenty-one thousand", input: 21000, want: "twenty-one thousand"},
		{name: "one hundred twenty-three thousand four hundred fifty-six", input: 123456, want: "one hundred twenty-three thousand four hundred fifty-six"},

		// Millions
		{name: "one million", input: 1000000, want: "one million"},
		{name: "two million", input: 2000000, want: "two million"},
		{name: "one million one", input: 1000001, want: "one million one"},
		{name: "one million two hundred thirty-four thousand five hundred sixty-seven", input: 1234567, want: "one million two hundred thirty-four thousand five hundred sixty-seven"},
		{name: "twelve million three hundred forty-five thousand six hundred seventy-eight", input: 12345678, want: "twelve million three hundred forty-five thousand six hundred seventy-eight"},
		{name: "one hundred twenty-three million four hundred fifty-six thousand seven hundred eighty-nine", input: 123456789, want: "one hundred twenty-three million four hundred fifty-six thousand seven hundred eighty-nine"},

		// Billions
		{name: "one billion", input: 1000000000, want: "one billion"},
		{name: "two billion", input: 2000000000, want: "two billion"},
		{name: "one billion one", input: 1000000001, want: "one billion one"},

		// Negative numbers
		{name: "negative one", input: -1, want: "negative one"},
		{name: "negative five", input: -5, want: "negative five"},
		{name: "negative eleven", input: -11, want: "negative eleven"},
		{name: "negative twenty-one", input: -21, want: "negative twenty-one"},
		{name: "negative forty-two", input: -42, want: "negative forty-two"},
		{name: "negative one hundred", input: -100, want: "negative one hundred"},
		{name: "negative one thousand", input: -1000, want: "negative one thousand"},
		{name: "negative one million", input: -1000000, want: "negative one million"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.NumberToWords(tt.input)
			if got != tt.want {
				t.Errorf("NumberToWords(%d) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestNumberToWordsComprehensive(t *testing.T) {
	tests := []struct {
		input int
		want  string
	}{
		// ===== SECTION 1: All numbers 0-20 (unique word forms) =====
		{0, "zero"},
		{1, "one"},
		{2, "two"},
		{3, "three"},
		{4, "four"},
		{5, "five"},
		{6, "six"},
		{7, "seven"},
		{8, "eight"},
		{9, "nine"},
		{10, "ten"},
		{11, "eleven"},
		{12, "twelve"},
		{13, "thirteen"},
		{14, "fourteen"},
		{15, "fifteen"},
		{16, "sixteen"},
		{17, "seventeen"},
		{18, "eighteen"},
		{19, "nineteen"},
		{20, "twenty"},

		// ===== SECTION 2: All tens (20-90) =====
		{20, "twenty"},
		{30, "thirty"},
		{40, "forty"},
		{50, "fifty"},
		{60, "sixty"},
		{70, "seventy"},
		{80, "eighty"},
		{90, "ninety"},

		// ===== SECTION 3: Compound numbers 21-29 =====
		{21, "twenty-one"},
		{22, "twenty-two"},
		{23, "twenty-three"},
		{24, "twenty-four"},
		{25, "twenty-five"},
		{26, "twenty-six"},
		{27, "twenty-seven"},
		{28, "twenty-eight"},
		{29, "twenty-nine"},

		// ===== SECTION 4: Representative compounds from each decade =====
		// 30s
		{31, "thirty-one"},
		{33, "thirty-three"},
		{35, "thirty-five"},
		{37, "thirty-seven"},
		{39, "thirty-nine"},
		// 40s
		{41, "forty-one"},
		{42, "forty-two"},
		{44, "forty-four"},
		{46, "forty-six"},
		{48, "forty-eight"},
		// 50s
		{51, "fifty-one"},
		{53, "fifty-three"},
		{55, "fifty-five"},
		{57, "fifty-seven"},
		{59, "fifty-nine"},
		// 60s
		{61, "sixty-one"},
		{63, "sixty-three"},
		{65, "sixty-five"},
		{67, "sixty-seven"},
		{69, "sixty-nine"},
		// 70s
		{71, "seventy-one"},
		{73, "seventy-three"},
		{75, "seventy-five"},
		{77, "seventy-seven"},
		{79, "seventy-nine"},
		// 80s
		{81, "eighty-one"},
		{83, "eighty-three"},
		{85, "eighty-five"},
		{87, "eighty-seven"},
		{89, "eighty-nine"},
		// 90s
		{91, "ninety-one"},
		{93, "ninety-three"},
		{95, "ninety-five"},
		{97, "ninety-seven"},
		{99, "ninety-nine"},

		// ===== SECTION 5: Boundary values =====
		{99, "ninety-nine"},
		{100, "one hundred"},
		{101, "one hundred one"},
		{110, "one hundred ten"},
		{111, "one hundred eleven"},
		{119, "one hundred nineteen"},
		{120, "one hundred twenty"},
		{121, "one hundred twenty-one"},
		{199, "one hundred ninety-nine"},
		{200, "two hundred"},
		{500, "five hundred"},
		{900, "nine hundred"},
		{999, "nine hundred ninety-nine"},
		{1000, "one thousand"},
		{1001, "one thousand one"},
		{1010, "one thousand ten"},
		{1011, "one thousand eleven"},
		{1100, "one thousand one hundred"},
		{1111, "one thousand one hundred eleven"},
		{1234, "one thousand two hundred thirty-four"},
		{1999, "one thousand nine hundred ninety-nine"},
		{2000, "two thousand"},
		{9999, "nine thousand nine hundred ninety-nine"},
		{10000, "ten thousand"},
		{10001, "ten thousand one"},
		{99999, "ninety-nine thousand nine hundred ninety-nine"},
		{100000, "one hundred thousand"},
		{100001, "one hundred thousand one"},
		{999999, "nine hundred ninety-nine thousand nine hundred ninety-nine"},
		{1000000, "one million"},
		{1000001, "one million one"},

		// ===== SECTION 6: Large numbers - Thousands =====
		{2000, "two thousand"},
		{3000, "three thousand"},
		{10000, "ten thousand"},
		{11000, "eleven thousand"},
		{12000, "twelve thousand"},
		{15000, "fifteen thousand"},
		{20000, "twenty thousand"},
		{21000, "twenty-one thousand"},
		{50000, "fifty thousand"},
		{99000, "ninety-nine thousand"},
		{100000, "one hundred thousand"},
		{123000, "one hundred twenty-three thousand"},
		{500000, "five hundred thousand"},
		{999000, "nine hundred ninety-nine thousand"},

		// ===== SECTION 7: Large numbers - Thousands with remainders =====
		{12345, "twelve thousand three hundred forty-five"},
		{23456, "twenty-three thousand four hundred fifty-six"},
		{54321, "fifty-four thousand three hundred twenty-one"},
		{100100, "one hundred thousand one hundred"},
		{100101, "one hundred thousand one hundred one"},
		{123456, "one hundred twenty-three thousand four hundred fifty-six"},
		{500500, "five hundred thousand five hundred"},
		{999999, "nine hundred ninety-nine thousand nine hundred ninety-nine"},

		// ===== SECTION 8: Large numbers - Millions =====
		{1000000, "one million"},
		{2000000, "two million"},
		{10000000, "ten million"},
		{12000000, "twelve million"},
		{20000000, "twenty million"},
		{100000000, "one hundred million"},
		{123000000, "one hundred twenty-three million"},
		{500000000, "five hundred million"},
		{999000000, "nine hundred ninety-nine million"},

		// ===== SECTION 9: Large numbers - Millions with remainders =====
		{1000001, "one million one"},
		{1001000, "one million one thousand"},
		{1001001, "one million one thousand one"},
		{1234567, "one million two hundred thirty-four thousand five hundred sixty-seven"},
		{12345678, "twelve million three hundred forty-five thousand six hundred seventy-eight"},
		{123456789, "one hundred twenty-three million four hundred fifty-six thousand seven hundred eighty-nine"},

		// ===== SECTION 10: Large numbers - Billions =====
		{1000000000, "one billion"},
		{2000000000, "two billion"},
		{1000000001, "one billion one"},
		{1000001000, "one billion one thousand"},
		{1001000000, "one billion one million"},
		{1234567890, "one billion two hundred thirty-four million five hundred sixty-seven thousand eight hundred ninety"},

		// ===== SECTION 11: Negative numbers =====
		{-1, "negative one"},
		{-2, "negative two"},
		{-5, "negative five"},
		{-9, "negative nine"},
		{-10, "negative ten"},
		{-11, "negative eleven"},
		{-12, "negative twelve"},
		{-13, "negative thirteen"},
		{-19, "negative nineteen"},
		{-20, "negative twenty"},
		{-21, "negative twenty-one"},
		{-29, "negative twenty-nine"},
		{-42, "negative forty-two"},
		{-50, "negative fifty"},
		{-99, "negative ninety-nine"},
		{-100, "negative one hundred"},
		{-101, "negative one hundred one"},
		{-111, "negative one hundred eleven"},
		{-999, "negative nine hundred ninety-nine"},
		{-1000, "negative one thousand"},
		{-1001, "negative one thousand one"},
		{-1234, "negative one thousand two hundred thirty-four"},
		{-10000, "negative ten thousand"},
		{-100000, "negative one hundred thousand"},
		{-1000000, "negative one million"},
		{-1000000000, "negative one billion"},

		// ===== SECTION 12: Edge cases - Special patterns =====
		// Numbers with zeros in different positions
		{101, "one hundred one"},
		{110, "one hundred ten"},
		{1001, "one thousand one"},
		{1010, "one thousand ten"},
		{1100, "one thousand one hundred"},
		{10001, "ten thousand one"},
		{10010, "ten thousand ten"},
		{10100, "ten thousand one hundred"},
		{100001, "one hundred thousand one"},
		{100010, "one hundred thousand ten"},
		{100100, "one hundred thousand one hundred"},
		{1000001, "one million one"},
		{1000010, "one million ten"},
		{1000100, "one million one hundred"},
		{1001000, "one million one thousand"},
		{1010000, "one million ten thousand"},
		{1100000, "one million one hundred thousand"},

		// Numbers that look like repeating patterns
		{111, "one hundred eleven"},
		{222, "two hundred twenty-two"},
		{333, "three hundred thirty-three"},
		{444, "four hundred forty-four"},
		{555, "five hundred fifty-five"},
		{666, "six hundred sixty-six"},
		{777, "seven hundred seventy-seven"},
		{888, "eight hundred eighty-eight"},
		{999, "nine hundred ninety-nine"},
		{1111, "one thousand one hundred eleven"},
		{2222, "two thousand two hundred twenty-two"},
		{11111, "eleven thousand one hundred eleven"},
		{111111, "one hundred eleven thousand one hundred eleven"},

		// Powers of 10
		{1, "one"},
		{10, "ten"},
		{100, "one hundred"},
		{1000, "one thousand"},
		{10000, "ten thousand"},
		{100000, "one hundred thousand"},
		{1000000, "one million"},
		{10000000, "ten million"},
		{100000000, "one hundred million"},
		{1000000000, "one billion"},

		// Just under powers of 10
		{9, "nine"},
		{99, "ninety-nine"},
		{999, "nine hundred ninety-nine"},
		{9999, "nine thousand nine hundred ninety-nine"},
		{99999, "ninety-nine thousand nine hundred ninety-nine"},
		{999999, "nine hundred ninety-nine thousand nine hundred ninety-nine"},
		{9999999, "nine million nine hundred ninety-nine thousand nine hundred ninety-nine"},
		{99999999, "ninety-nine million nine hundred ninety-nine thousand nine hundred ninety-nine"},
		{999999999, "nine hundred ninety-nine million nine hundred ninety-nine thousand nine hundred ninety-nine"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("input_%d", tt.input), func(t *testing.T) {
			got := inflect.NumberToWords(tt.input)
			if got != tt.want {
				t.Errorf("NumberToWords(%d) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestNumberToWordsFloat(t *testing.T) {
	tests := []struct {
		name  string
		input float64
		want  string
	}{
		// Basic decimal numbers
		{name: "pi approximation", input: 3.14, want: "three point one four"},
		{name: "zero point five", input: 0.5, want: "zero point five"},
		{name: "negative e approximation", input: -2.718, want: "negative two point seven one eight"},

		// Whole numbers (no decimal part effectively)
		{name: "whole number with decimal", input: 5.0, want: "five"},
		{name: "zero", input: 0.0, want: "zero"},

		// Single decimal digit
		{name: "one point zero", input: 1.1, want: "one point one"},
		{name: "nine point nine", input: 9.9, want: "nine point nine"},

		// Multiple decimal digits
		{name: "long decimal", input: 1.23456, want: "one point two three four five six"},

		// Negative numbers
		{name: "negative simple", input: -1.5, want: "negative one point five"},
		{name: "negative zero point", input: -0.25, want: "negative zero point two five"},

		// Larger integer parts
		{name: "hundred point decimal", input: 100.01, want: "one hundred point zero one"},
		{name: "thousand point decimal", input: 1000.999, want: "one thousand point nine nine nine"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.NumberToWordsFloat(tt.input)
			if got != tt.want {
				t.Errorf("NumberToWordsFloat(%v) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestNumberToWordsFloatWithDecimal(t *testing.T) {
	tests := []struct {
		name    string
		input   float64
		decimal string
		want    string
	}{
		// Examples from requirements
		{name: "point decimal", input: 3.14, decimal: "point", want: "three point one four"},
		{name: "dot decimal", input: 3.14, decimal: "dot", want: "three dot one four"},
		{name: "and decimal", input: 3.14, decimal: "and", want: "three and one four"},

		// Different decimal words
		{name: "comma decimal", input: 2.5, decimal: "comma", want: "two comma five"},
		{name: "decimal word", input: 1.23, decimal: "decimal", want: "one decimal two three"},

		// Negative numbers with custom decimal
		{name: "negative with dot", input: -1.5, decimal: "dot", want: "negative one dot five"},
		{name: "negative with and", input: -0.25, decimal: "and", want: "negative zero and two five"},

		// Whole numbers (no decimal part)
		{name: "whole number with custom decimal", input: 5.0, decimal: "dot", want: "five"},
		{name: "zero with custom decimal", input: 0.0, decimal: "and", want: "zero"},

		// Larger numbers
		{name: "hundred with dot", input: 100.01, decimal: "dot", want: "one hundred dot zero one"},
		{name: "thousand with comma", input: 1000.999, decimal: "comma", want: "one thousand comma nine nine nine"},

		// Long decimal
		{name: "long decimal with dot", input: 1.23456, decimal: "dot", want: "one dot two three four five six"},

		// Edge cases with empty or unusual decimal words
		{name: "single char decimal", input: 7.89, decimal: ".", want: "seven . eight nine"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.NumberToWordsFloatWithDecimal(tt.input, tt.decimal)
			if got != tt.want {
				t.Errorf("NumberToWordsFloatWithDecimal(%v, %q) = %q, want %q", tt.input, tt.decimal, got, tt.want)
			}
		})
	}
}

func TestNumberToWordsThreshold(t *testing.T) {
	tests := []struct {
		name      string
		n         int
		threshold int
		want      string
	}{
		// Numbers below threshold - convert to words
		{name: "5 below 10", n: 5, threshold: 10, want: "five"},
		{name: "0 below 10", n: 0, threshold: 10, want: "zero"},
		{name: "9 below 10", n: 9, threshold: 10, want: "nine"},
		{name: "1 below 5", n: 1, threshold: 5, want: "one"},
		{name: "99 below 100", n: 99, threshold: 100, want: "ninety-nine"},

		// Numbers at threshold - return as string
		{name: "10 at 10", n: 10, threshold: 10, want: "10"},
		{name: "100 at 100", n: 100, threshold: 100, want: "100"},
		{name: "5 at 5", n: 5, threshold: 5, want: "5"},
		{name: "0 at 0", n: 0, threshold: 0, want: "0"},

		// Numbers above threshold - return as string
		{name: "15 above 10", n: 15, threshold: 10, want: "15"},
		{name: "100 above 10", n: 100, threshold: 10, want: "100"},
		{name: "1000 above 100", n: 1000, threshold: 100, want: "1000"},
		{name: "50 above 10", n: 50, threshold: 10, want: "50"},

		// Negative numbers below threshold - convert to words
		{name: "negative 3 below 10", n: -3, threshold: 10, want: "negative three"},
		{name: "negative 1 below 5", n: -1, threshold: 5, want: "negative one"},
		{name: "negative 99 below 100", n: -99, threshold: 100, want: "negative ninety-nine"},

		// Threshold of 1 (only convert 0 and negatives)
		{name: "0 below 1", n: 0, threshold: 1, want: "zero"},
		{name: "1 at 1", n: 1, threshold: 1, want: "1"},
		{name: "negative 5 below 1", n: -5, threshold: 1, want: "negative five"},

		// Large threshold
		{name: "999 below 1000", n: 999, threshold: 1000, want: "nine hundred ninety-nine"},
		{name: "1000 at 1000", n: 1000, threshold: 1000, want: "1000"},

		// Edge cases with negative threshold
		{name: "5 above negative threshold", n: 5, threshold: -10, want: "5"},
		{name: "negative 5 above negative 10", n: -5, threshold: -10, want: "-5"},
		{name: "negative 15 below negative 10", n: -15, threshold: -10, want: "negative fifteen"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.NumberToWordsThreshold(tt.n, tt.threshold)
			if got != tt.want {
				t.Errorf("NumberToWordsThreshold(%d, %d) = %q, want %q", tt.n, tt.threshold, got, tt.want)
			}
		})
	}
}

func TestNumberToWordsGrouped(t *testing.T) {
	tests := []struct {
		name      string
		n         int
		groupSize int
		want      string
	}{
		// Basic examples from requirements
		{name: "1234 grouped by 2", n: 1234, groupSize: 2, want: "twelve thirty-four"},
		{name: "123456 grouped by 2", n: 123456, groupSize: 2, want: "twelve thirty-four fifty-six"},
		{name: "1234 grouped by 3", n: 1234, groupSize: 3, want: "one two hundred thirty-four"},
		{name: "1234567890 grouped by 3", n: 1234567890, groupSize: 3, want: "one two hundred thirty-four five hundred sixty-seven eight hundred ninety"},

		// Phone number use cases (grouped by various sizes)
		{name: "phone 5551234 by 4", n: 5551234, groupSize: 4, want: "five hundred fifty-five one thousand two hundred thirty-four"},
		{name: "area code 415 by 3", n: 415, groupSize: 3, want: "four hundred fifteen"},

		// Credit card style (groups of 4)
		{name: "1234567812345678 by 4", n: 1234567812345678, groupSize: 4, want: "one thousand two hundred thirty-four five thousand six hundred seventy-eight one thousand two hundred thirty-four five thousand six hundred seventy-eight"},

		// Zero handling
		{name: "zero grouped by 2", n: 0, groupSize: 2, want: "zero"},
		{name: "zero grouped by 3", n: 0, groupSize: 3, want: "zero"},

		// Negative numbers
		{name: "negative 1234 by 2", n: -1234, groupSize: 2, want: "negative twelve thirty-four"},
		{name: "negative 123456 by 2", n: -123456, groupSize: 2, want: "negative twelve thirty-four fifty-six"},

		// Single digit groups
		{name: "12345 grouped by 1", n: 12345, groupSize: 1, want: "one two three four five"},
		{name: "9876 grouped by 1", n: 9876, groupSize: 1, want: "nine eight seven six"},

		// Groups with zeros (each group is converted as a number, so 00->0, 01->1)
		{name: "10001 grouped by 2", n: 10001, groupSize: 2, want: "one zero one"},
		{name: "100100 grouped by 2", n: 100100, groupSize: 2, want: "ten one zero"},
		{name: "1001001 grouped by 3", n: 1001001, groupSize: 3, want: "one one one"},

		// Small numbers with various group sizes
		{name: "5 grouped by 2", n: 5, groupSize: 2, want: "five"},
		{name: "12 grouped by 2", n: 12, groupSize: 2, want: "twelve"},
		{name: "123 grouped by 2", n: 123, groupSize: 2, want: "one twenty-three"},
		{name: "99 grouped by 3", n: 99, groupSize: 3, want: "ninety-nine"},

		// Larger group sizes
		{name: "123456 grouped by 4", n: 123456, groupSize: 4, want: "twelve three thousand four hundred fifty-six"},
		{name: "1234567 grouped by 4", n: 1234567, groupSize: 4, want: "one hundred twenty-three four thousand five hundred sixty-seven"},

		// Edge case: group size <= 0 falls back to regular NumberToWords
		{name: "1234 with group size 0", n: 1234, groupSize: 0, want: "one thousand two hundred thirty-four"},
		{name: "1234 with negative group size", n: -1, groupSize: -1, want: "negative one"},

		// Boundary cases
		{name: "10 grouped by 2", n: 10, groupSize: 2, want: "ten"},
		{name: "100 grouped by 2", n: 100, groupSize: 2, want: "one zero"},
		{name: "1000 grouped by 2", n: 1000, groupSize: 2, want: "ten zero"},
		{name: "10000 grouped by 2", n: 10000, groupSize: 2, want: "one zero zero"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.NumberToWordsGrouped(tt.n, tt.groupSize)
			if got != tt.want {
				t.Errorf("NumberToWordsGrouped(%d, %d) = %q, want %q", tt.n, tt.groupSize, got, tt.want)
			}
		})
	}
}

func TestFormatNumber(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  string
	}{
		// Basic cases from requirements
		{name: "1000", input: 1000, want: "1,000"},
		{name: "1000000", input: 1000000, want: "1,000,000"},
		{name: "123456789", input: 123456789, want: "123,456,789"},
		{name: "negative 1234", input: -1234, want: "-1,234"},
		{name: "999 no comma", input: 999, want: "999"},

		// Small numbers (no commas needed)
		{name: "zero", input: 0, want: "0"},
		{name: "single digit", input: 5, want: "5"},
		{name: "two digits", input: 42, want: "42"},
		{name: "three digits", input: 100, want: "100"},

		// Boundary cases around 1000
		{name: "999 boundary", input: 999, want: "999"},
		{name: "1000 boundary", input: 1000, want: "1,000"},
		{name: "1001", input: 1001, want: "1,001"},

		// Various digit groupings
		{name: "4 digits", input: 1234, want: "1,234"},
		{name: "5 digits", input: 12345, want: "12,345"},
		{name: "6 digits", input: 123456, want: "123,456"},
		{name: "7 digits", input: 1234567, want: "1,234,567"},
		{name: "8 digits", input: 12345678, want: "12,345,678"},
		{name: "9 digits", input: 123456789, want: "123,456,789"},
		{name: "10 digits", input: 1234567890, want: "1,234,567,890"},

		// Numbers with zeros
		{name: "10000", input: 10000, want: "10,000"},
		{name: "100000", input: 100000, want: "100,000"},
		{name: "1000000", input: 1000000, want: "1,000,000"},
		{name: "10000000", input: 10000000, want: "10,000,000"},
		{name: "1000000000", input: 1000000000, want: "1,000,000,000"},

		// Negative numbers
		{name: "negative small", input: -5, want: "-5"},
		{name: "negative three digits", input: -999, want: "-999"},
		{name: "negative 1000", input: -1000, want: "-1,000"},
		{name: "negative million", input: -1000000, want: "-1,000,000"},
		{name: "negative large", input: -123456789, want: "-123,456,789"},

		// Numbers with trailing zeros
		{name: "trailing zeros 4 digits", input: 1000, want: "1,000"},
		{name: "trailing zeros 7 digits", input: 1000000, want: "1,000,000"},
		{name: "mixed with zeros", input: 1002003, want: "1,002,003"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.FormatNumber(tt.input)
			if got != tt.want {
				t.Errorf("FormatNumber(%d) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestNo(t *testing.T) {
	tests := []struct {
		name  string
		word  string
		count int
		want  string
	}{
		// Zero count - uses "no" with plural
		{name: "zero errors", word: "error", count: 0, want: "no errors"},
		{name: "zero cats", word: "cat", count: 0, want: "no cats"},
		{name: "zero children", word: "child", count: 0, want: "no children"},
		{name: "zero mice", word: "mouse", count: 0, want: "no mice"},
		{name: "zero sheep", word: "sheep", count: 0, want: "no sheep"},
		{name: "zero boxes", word: "box", count: 0, want: "no boxes"},

		// Count of 1 - singular
		{name: "one error", word: "error", count: 1, want: "1 error"},
		{name: "one cat", word: "cat", count: 1, want: "1 cat"},
		{name: "one child", word: "child", count: 1, want: "1 child"},
		{name: "one mouse", word: "mouse", count: 1, want: "1 mouse"},
		{name: "one sheep", word: "sheep", count: 1, want: "1 sheep"},
		{name: "one box", word: "box", count: 1, want: "1 box"},

		// Count > 1 - plural
		{name: "two errors", word: "error", count: 2, want: "2 errors"},
		{name: "five cats", word: "cat", count: 5, want: "5 cats"},
		{name: "three children", word: "child", count: 3, want: "3 children"},
		{name: "ten mice", word: "mouse", count: 10, want: "10 mice"},
		{name: "hundred sheep", word: "sheep", count: 100, want: "100 sheep"},
		{name: "twenty boxes", word: "box", count: 20, want: "20 boxes"},

		// Large counts
		{name: "thousand errors", word: "error", count: 1000, want: "1000 errors"},
		{name: "million items", word: "item", count: 1000000, want: "1000000 items"},

		// Negative counts
		{name: "negative one error", word: "error", count: -1, want: "-1 error"},
		{name: "negative two errors", word: "error", count: -2, want: "-2 errors"},
		{name: "negative five cats", word: "cat", count: -5, want: "-5 cats"},

		// Empty word
		{name: "zero empty", word: "", count: 0, want: "no "},
		{name: "one empty", word: "", count: 1, want: "1 "},

		// Case preservation
		{name: "zero Errors titlecase", word: "Error", count: 0, want: "no Errors"},
		{name: "zero ERRORS uppercase", word: "ERROR", count: 0, want: "no ERRORS"},
		{name: "two Children titlecase", word: "Child", count: 2, want: "2 Children"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.No(tt.word, tt.count)
			if got != tt.want {
				t.Errorf("No(%q, %d) = %q, want %q", tt.word, tt.count, got, tt.want)
			}
		})
	}
}

func TestNum(t *testing.T) {
	defer inflect.Num()

	// Clear state before tests
	inflect.Num(0)

	tests := []struct {
		name       string
		args       []int
		wantReturn int
		wantGetNum int
	}{
		// Store positive values
		{name: "store 5", args: []int{5}, wantReturn: 5, wantGetNum: 5},
		{name: "store 1", args: []int{1}, wantReturn: 1, wantGetNum: 1},
		{name: "store 100", args: []int{100}, wantReturn: 100, wantGetNum: 100},
		{name: "store large number", args: []int{999999}, wantReturn: 999999, wantGetNum: 999999},

		// Clear with 0
		{name: "clear with 0", args: []int{0}, wantReturn: 0, wantGetNum: 0},

		// Clear with no args
		{name: "clear with no args", args: []int{}, wantReturn: 0, wantGetNum: 0},

		// Negative values
		{name: "store negative", args: []int{-5}, wantReturn: -5, wantGetNum: -5},
		{name: "store -1", args: []int{-1}, wantReturn: -1, wantGetNum: -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset state before each test
			inflect.Num(0)

			var got int
			if len(tt.args) == 0 {
				got = inflect.Num()
			} else {
				got = inflect.Num(tt.args[0])
			}

			if got != tt.wantReturn {
				t.Errorf("Num(%v) = %d, want %d", tt.args, got, tt.wantReturn)
			}

			gotNum := inflect.GetNum()
			if gotNum != tt.wantGetNum {
				t.Errorf("After Num(%v): GetNum() = %d, want %d", tt.args, gotNum, tt.wantGetNum)
			}
		})
	}
}

func TestGetNum(t *testing.T) {
	defer inflect.Num()

	tests := []struct {
		name    string
		setup   func()
		wantNum int
	}{
		{
			name:    "initial state",
			setup:   func() { inflect.Num(0) },
			wantNum: 0,
		},
		{
			name:    "after storing 5",
			setup:   func() { inflect.Num(5) },
			wantNum: 5,
		},
		{
			name:    "after storing 10",
			setup:   func() { inflect.Num(10) },
			wantNum: 10,
		},
		{
			name:    "after clearing with 0",
			setup:   func() { inflect.Num(5); inflect.Num(0) },
			wantNum: 0,
		},
		{
			name:    "after clearing with no args",
			setup:   func() { inflect.Num(5); inflect.Num() },
			wantNum: 0,
		},
		{
			name:    "after multiple stores",
			setup:   func() { inflect.Num(1); inflect.Num(2); inflect.Num(3) },
			wantNum: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			got := inflect.GetNum()
			if got != tt.wantNum {
				t.Errorf("GetNum() = %d, want %d", got, tt.wantNum)
			}
		})
	}
}

func TestNumIntegration(t *testing.T) {
	// Clean up after test
	defer inflect.Num(0)

	t.Run("complete workflow", func(t *testing.T) {
		// 1. Initial state should be 0
		inflect.Num(0)
		if got := inflect.GetNum(); got != 0 {
			t.Errorf("Initial GetNum() = %d, want 0", got)
		}

		// 2. Store a value
		ret := inflect.Num(5)
		if ret != 5 {
			t.Errorf("Num(5) = %d, want 5", ret)
		}
		if got := inflect.GetNum(); got != 5 {
			t.Errorf("After Num(5): GetNum() = %d, want 5", got)
		}

		// 3. Update the value
		ret = inflect.Num(10)
		if ret != 10 {
			t.Errorf("Num(10) = %d, want 10", ret)
		}
		if got := inflect.GetNum(); got != 10 {
			t.Errorf("After Num(10): GetNum() = %d, want 10", got)
		}

		// 4. Clear with 0
		ret = inflect.Num(0)
		if ret != 0 {
			t.Errorf("Num(0) = %d, want 0", ret)
		}
		if got := inflect.GetNum(); got != 0 {
			t.Errorf("After Num(0): GetNum() = %d, want 0", got)
		}

		// 5. Store again and clear with no args
		inflect.Num(7)
		ret = inflect.Num()
		if ret != 0 {
			t.Errorf("Num() = %d, want 0", ret)
		}
		if got := inflect.GetNum(); got != 0 {
			t.Errorf("After Num(): GetNum() = %d, want 0", got)
		}
	})
}

func BenchmarkNumberToWords(b *testing.B) {
	// Test with numbers of varying magnitudes
	benchmarks := []struct {
		name  string
		input int
	}{
		{"zero", 0},
		{"single_digit", 7},
		{"teen", 15},
		{"two_digit", 42},
		{"hundred", 123},
		{"thousand", 1234},
		{"million", 1234567},
		{"billion", 1234567890},
		{"negative", -42},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for range b.N {
				inflect.NumberToWords(bm.input)
			}
		})
	}
}
