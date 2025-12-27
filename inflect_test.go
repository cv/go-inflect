package inflect_test

import (
	"testing"

	inflect "github.com/cv/go-inflect"
)

func TestJoin(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  string
	}{
		// Empty slice
		{name: "empty slice", input: []string{}, want: ""},
		{name: "nil slice", input: nil, want: ""},

		// Single word
		{name: "single word", input: []string{"apple"}, want: "apple"},
		{name: "single empty string", input: []string{""}, want: ""},

		// Two words
		{name: "two words", input: []string{"apple", "banana"}, want: "apple and banana"},
		{name: "two short words", input: []string{"a", "b"}, want: "a and b"},

		// Three+ words (Oxford comma)
		{name: "three words", input: []string{"a", "b", "c"}, want: "a, b, and c"},
		{name: "four words", input: []string{"apple", "banana", "cherry", "date"}, want: "apple, banana, cherry, and date"},
		{name: "five words", input: []string{"one", "two", "three", "four", "five"}, want: "one, two, three, four, and five"},

		// Words with special characters
		{name: "words with commas", input: []string{"red, blue", "green"}, want: "red, blue and green"},
		{name: "words with quotes", input: []string{`"hello"`, `"world"`}, want: `"hello" and "world"`},
		{name: "words with unicode", input: []string{"café", "naïve", "résumé"}, want: "café, naïve, and résumé"},
		{name: "words with numbers", input: []string{"item1", "item2", "item3"}, want: "item1, item2, and item3"},
		{name: "words with spaces", input: []string{"New York", "Los Angeles", "Chicago"}, want: "New York, Los Angeles, and Chicago"},
		{name: "words with ampersand", input: []string{"Tom & Jerry", "Mickey"}, want: "Tom & Jerry and Mickey"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Join(tt.input)
			if got != tt.want {
				t.Errorf("Join(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestJoinWithConj(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		conj  string
		want  string
	}{
		// Empty slice
		{name: "empty slice with or", input: []string{}, conj: "or", want: ""},
		{name: "nil slice with or", input: nil, conj: "or", want: ""},

		// Single word
		{name: "single word with or", input: []string{"apple"}, conj: "or", want: "apple"},
		{name: "single empty string with or", input: []string{""}, conj: "or", want: ""},

		// Two words with "or"
		{name: "two words with or", input: []string{"apple", "banana"}, conj: "or", want: "apple or banana"},
		{name: "two short words with or", input: []string{"a", "b"}, conj: "or", want: "a or b"},

		// Three+ words with "or" (Oxford comma)
		{name: "three words with or", input: []string{"a", "b", "c"}, conj: "or", want: "a, b, or c"},
		{name: "four words with or", input: []string{"apple", "banana", "cherry", "date"}, conj: "or", want: "apple, banana, cherry, or date"},

		// Other conjunctions
		{name: "two words with and/or", input: []string{"a", "b"}, conj: "and/or", want: "a and/or b"},
		{name: "three words with and/or", input: []string{"a", "b", "c"}, conj: "and/or", want: "a, b, and/or c"},
		{name: "two words with nor", input: []string{"this", "that"}, conj: "nor", want: "this nor that"},
		{name: "three words with nor", input: []string{"this", "that", "other"}, conj: "nor", want: "this, that, nor other"},

		// Verify "and" conjunction matches Join() behavior
		{name: "two words with and", input: []string{"a", "b"}, conj: "and", want: "a and b"},
		{name: "three words with and", input: []string{"a", "b", "c"}, conj: "and", want: "a, b, and c"},

		// Words with special characters
		{name: "words with spaces and or", input: []string{"New York", "Los Angeles"}, conj: "or", want: "New York or Los Angeles"},
		{name: "words with unicode and or", input: []string{"café", "thé", "chocolat"}, conj: "or", want: "café, thé, or chocolat"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.JoinWithConj(tt.input, tt.conj)
			if got != tt.want {
				t.Errorf("JoinWithConj(%q, %q) = %q, want %q", tt.input, tt.conj, got, tt.want)
			}
		})
	}
}

func TestJoinWithSep(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		conj  string
		sep   string
		want  string
	}{
		// Empty slice
		{name: "empty slice", input: []string{}, conj: "and", sep: "; ", want: ""},
		{name: "nil slice", input: nil, conj: "and", sep: "; ", want: ""},

		// Single word
		{name: "single word", input: []string{"apple"}, conj: "and", sep: "; ", want: "apple"},

		// Two words (separator not used between two items)
		{name: "two words", input: []string{"a", "b"}, conj: "and", sep: "; ", want: "a and b"},
		{name: "two words with or", input: []string{"a", "b"}, conj: "or", sep: "; ", want: "a or b"},

		// Three+ words with semicolon separator
		{name: "three words semicolon", input: []string{"a", "b", "c"}, conj: "and", sep: "; ", want: "a; b; and c"},
		{name: "four words semicolon", input: []string{"a", "b", "c", "d"}, conj: "and", sep: "; ", want: "a; b; c; and d"},
		{name: "three words semicolon or", input: []string{"a", "b", "c"}, conj: "or", sep: "; ", want: "a; b; or c"},

		// Items containing commas (primary use case)
		{name: "dates with commas", input: []string{"Jan 1, 2020", "Feb 2, 2021", "Mar 3, 2022"}, conj: "and", sep: "; ", want: "Jan 1, 2020; Feb 2, 2021; and Mar 3, 2022"},
		{name: "locations with commas", input: []string{"New York, NY", "Los Angeles, CA", "Chicago, IL"}, conj: "or", sep: "; ", want: "New York, NY; Los Angeles, CA; or Chicago, IL"},
		{name: "names with titles", input: []string{"Smith, John", "Doe, Jane"}, conj: "and", sep: "; ", want: "Smith, John and Doe, Jane"},

		// Custom separators
		{name: "pipe separator", input: []string{"a", "b", "c"}, conj: "and", sep: " | ", want: "a | b | and c"},
		{name: "dash separator", input: []string{"x", "y", "z"}, conj: "or", sep: " - ", want: "x - y - or z"},
		{name: "newline separator", input: []string{"line1", "line2", "line3"}, conj: "and", sep: "\n", want: "line1\nline2\nand line3"},

		// Verify comma separator matches JoinWithConj behavior
		{name: "comma separator three", input: []string{"a", "b", "c"}, conj: "and", sep: ", ", want: "a, b, and c"},
		{name: "comma separator four", input: []string{"a", "b", "c", "d"}, conj: "or", sep: ", ", want: "a, b, c, or d"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.JoinWithSep(tt.input, tt.conj, tt.sep)
			if got != tt.want {
				t.Errorf("JoinWithSep(%q, %q, %q) = %q, want %q", tt.input, tt.conj, tt.sep, got, tt.want)
			}
		})
	}
}

func TestJoinWithAutoSep(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		conj  string
		want  string
	}{
		// Empty slice
		{name: "empty slice", input: []string{}, conj: "and", want: ""},
		{name: "nil slice", input: nil, conj: "and", want: ""},

		// Single word (no commas)
		{name: "single word", input: []string{"apple"}, conj: "and", want: "apple"},
		{name: "single word with comma", input: []string{"Jan 1, 2020"}, conj: "and", want: "Jan 1, 2020"},

		// Two words without commas -> uses comma separator
		{name: "two words no commas", input: []string{"a", "b"}, conj: "and", want: "a and b"},
		{name: "two words no commas or", input: []string{"a", "b"}, conj: "or", want: "a or b"},

		// Two words with commas -> uses semicolon separator
		{name: "two words with commas", input: []string{"Jan 1, 2020", "Feb 2, 2021"}, conj: "and", want: "Jan 1, 2020; and Feb 2, 2021"},
		{name: "two words one has comma", input: []string{"Jan 1, 2020", "March"}, conj: "and", want: "Jan 1, 2020; and March"},

		// Three+ words without commas -> uses comma separator
		{name: "three words no commas", input: []string{"a", "b", "c"}, conj: "and", want: "a, b, and c"},
		{name: "four words no commas", input: []string{"a", "b", "c", "d"}, conj: "or", want: "a, b, c, or d"},

		// Three+ words with commas -> uses semicolon separator
		{name: "three words with commas", input: []string{"Jan 1, 2020", "Feb 2, 2021", "Mar 3, 2022"}, conj: "and", want: "Jan 1, 2020; Feb 2, 2021; and Mar 3, 2022"},
		{name: "three words one has comma", input: []string{"apple", "Jan 1, 2020", "banana"}, conj: "and", want: "apple; Jan 1, 2020; and banana"},
		{name: "locations with commas", input: []string{"New York, NY", "Los Angeles, CA", "Chicago, IL"}, conj: "or", want: "New York, NY; Los Angeles, CA; or Chicago, IL"},

		// Names in last, first format
		{name: "names with commas", input: []string{"Smith, John", "Doe, Jane", "Brown, Bob"}, conj: "and", want: "Smith, John; Doe, Jane; and Brown, Bob"},

		// Edge cases
		{name: "comma only in last item", input: []string{"red", "green", "blue, dark"}, conj: "and", want: "red; green; and blue, dark"},
		{name: "comma only in first item", input: []string{"red, light", "green", "blue"}, conj: "and", want: "red, light; green; and blue"},
		{name: "multiple commas in items", input: []string{"a, b, c", "x, y, z"}, conj: "and", want: "a, b, c; and x, y, z"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.JoinWithAutoSep(tt.input, tt.conj)
			if got != tt.want {
				t.Errorf("JoinWithAutoSep(%q, %q) = %q, want %q", tt.input, tt.conj, got, tt.want)
			}
		})
	}
}

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

func TestOrdinalWord(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  string
	}{
		// Basic cases (1-10)
		{name: "first", input: 1, want: "first"},
		{name: "second", input: 2, want: "second"},
		{name: "third", input: 3, want: "third"},
		{name: "fourth", input: 4, want: "fourth"},
		{name: "fifth", input: 5, want: "fifth"},
		{name: "sixth", input: 6, want: "sixth"},
		{name: "seventh", input: 7, want: "seventh"},
		{name: "eighth", input: 8, want: "eighth"},
		{name: "ninth", input: 9, want: "ninth"},
		{name: "tenth", input: 10, want: "tenth"},

		// Teens (11-19)
		{name: "eleventh", input: 11, want: "eleventh"},
		{name: "twelfth", input: 12, want: "twelfth"},
		{name: "thirteenth", input: 13, want: "thirteenth"},
		{name: "fourteenth", input: 14, want: "fourteenth"},
		{name: "fifteenth", input: 15, want: "fifteenth"},
		{name: "sixteenth", input: 16, want: "sixteenth"},
		{name: "seventeenth", input: 17, want: "seventeenth"},
		{name: "eighteenth", input: 18, want: "eighteenth"},
		{name: "nineteenth", input: 19, want: "nineteenth"},

		// Exact tens (20, 30, ...)
		{name: "twentieth", input: 20, want: "twentieth"},
		{name: "thirtieth", input: 30, want: "thirtieth"},
		{name: "fortieth", input: 40, want: "fortieth"},
		{name: "fiftieth", input: 50, want: "fiftieth"},
		{name: "sixtieth", input: 60, want: "sixtieth"},
		{name: "seventieth", input: 70, want: "seventieth"},
		{name: "eightieth", input: 80, want: "eightieth"},
		{name: "ninetieth", input: 90, want: "ninetieth"},

		// Compound tens (21-29, etc.)
		{name: "twenty-first", input: 21, want: "twenty-first"},
		{name: "twenty-second", input: 22, want: "twenty-second"},
		{name: "twenty-third", input: 23, want: "twenty-third"},
		{name: "twenty-fourth", input: 24, want: "twenty-fourth"},
		{name: "thirty-fifth", input: 35, want: "thirty-fifth"},
		{name: "forty-second", input: 42, want: "forty-second"},
		{name: "ninety-ninth", input: 99, want: "ninety-ninth"},

		// Hundreds
		{name: "one hundredth", input: 100, want: "one hundredth"},
		{name: "one hundred first", input: 101, want: "one hundred first"},
		{name: "one hundred second", input: 102, want: "one hundred second"},
		{name: "one hundred tenth", input: 110, want: "one hundred tenth"},
		{name: "one hundred eleventh", input: 111, want: "one hundred eleventh"},
		{name: "one hundred twentieth", input: 120, want: "one hundred twentieth"},
		{name: "one hundred twenty-first", input: 121, want: "one hundred twenty-first"},
		{name: "two hundredth", input: 200, want: "two hundredth"},
		{name: "five hundred fifty-fifth", input: 555, want: "five hundred fifty-fifth"},
		{name: "nine hundred ninety-ninth", input: 999, want: "nine hundred ninety-ninth"},

		// Thousands
		{name: "one thousandth", input: 1000, want: "one thousandth"},
		{name: "one thousand first", input: 1001, want: "one thousand first"},
		{name: "one thousand tenth", input: 1010, want: "one thousand tenth"},
		{name: "one thousand one hundredth", input: 1100, want: "one thousand one hundredth"},
		{name: "two thousandth", input: 2000, want: "two thousandth"},
		{name: "twelve thousandth", input: 12000, want: "twelve thousandth"},
		{name: "twenty-one thousandth", input: 21000, want: "twenty-one thousandth"},

		// Large numbers
		{name: "one millionth", input: 1000000, want: "one millionth"},
		{name: "one million first", input: 1000001, want: "one million first"},
		{name: "one billionth", input: 1000000000, want: "one billionth"},

		// Edge cases
		{name: "zeroth", input: 0, want: "zeroth"},
		{name: "negative first", input: -1, want: "negative first"},
		{name: "negative second", input: -2, want: "negative second"},
		{name: "negative eleventh", input: -11, want: "negative eleventh"},
		{name: "negative twenty-first", input: -21, want: "negative twenty-first"},
		{name: "negative one hundredth", input: -100, want: "negative one hundredth"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.OrdinalWord(tt.input)
			if got != tt.want {
				t.Errorf("OrdinalWord(%d) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestOrdinal(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  string
	}{
		// Basic cases (1st, 2nd, 3rd, 4th-9th)
		{name: "first", input: 1, want: "1st"},
		{name: "second", input: 2, want: "2nd"},
		{name: "third", input: 3, want: "3rd"},
		{name: "fourth", input: 4, want: "4th"},
		{name: "fifth", input: 5, want: "5th"},
		{name: "sixth", input: 6, want: "6th"},
		{name: "seventh", input: 7, want: "7th"},
		{name: "eighth", input: 8, want: "8th"},
		{name: "ninth", input: 9, want: "9th"},
		{name: "tenth", input: 10, want: "10th"},

		// Teens (11th, 12th, 13th) - special cases
		{name: "eleventh", input: 11, want: "11th"},
		{name: "twelfth", input: 12, want: "12th"},
		{name: "thirteenth", input: 13, want: "13th"},
		{name: "fourteenth", input: 14, want: "14th"},
		{name: "nineteenth", input: 19, want: "19th"},

		// Larger numbers (21st, 22nd, 23rd, 100th, 101st)
		{name: "twenty-first", input: 21, want: "21st"},
		{name: "twenty-second", input: 22, want: "22nd"},
		{name: "twenty-third", input: 23, want: "23rd"},
		{name: "twenty-fourth", input: 24, want: "24th"},
		{name: "hundredth", input: 100, want: "100th"},
		{name: "hundred-first", input: 101, want: "101st"},
		{name: "hundred-second", input: 102, want: "102nd"},
		{name: "hundred-third", input: 103, want: "103rd"},
		{name: "hundred-eleventh", input: 111, want: "111th"},
		{name: "hundred-twelfth", input: 112, want: "112th"},
		{name: "hundred-thirteenth", input: 113, want: "113th"},
		{name: "thousandth", input: 1000, want: "1000th"},
		{name: "thousand-first", input: 1001, want: "1001st"},
		{name: "thousand-eleventh", input: 1011, want: "1011th"},

		// Edge cases
		{name: "zero", input: 0, want: "0th"},
		{name: "negative one", input: -1, want: "-1st"},
		{name: "negative two", input: -2, want: "-2nd"},
		{name: "negative three", input: -3, want: "-3rd"},
		{name: "negative eleven", input: -11, want: "-11th"},
		{name: "negative twenty-one", input: -21, want: "-21st"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Ordinal(tt.input)
			if got != tt.want {
				t.Errorf("Ordinal(%d) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestAn(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		// Basic cases
		{name: "consonant start", input: "cat", want: "a cat"},
		{name: "vowel start", input: "ant", want: "an ant"},

		// Single letters
		{name: "vowel letter", input: "a", want: "an a"},
		{name: "consonant letter", input: "b", want: "a b"},

		// Silent H
		{name: "silent h", input: "honest cat", want: "an honest cat"},
		{name: "regular h", input: "dishonest cat", want: "a dishonest cat"},
		{name: "h proper noun", input: "Honolulu sunset", want: "a Honolulu sunset"},

		// Special pronunciation cases
		{name: "mpeg abbreviation", input: "mpeg", want: "an mpeg"},
		{name: "onetime exception", input: "onetime holiday", want: "a onetime holiday"},

		// Vowels with consonant sounds (U variations)
		{name: "Ugandan", input: "Ugandan person", want: "a Ugandan person"},
		{name: "Ukrainian", input: "Ukrainian person", want: "a Ukrainian person"},
		{name: "Unabomber", input: "Unabomber", want: "a Unabomber"},
		{name: "unanimous", input: "unanimous decision", want: "a unanimous decision"},

		// Abbreviations and acronyms
		{name: "US abbreviation", input: "US farmer", want: "a US farmer"},
		{name: "uppercase word", input: "wild PIKACHU appeared", want: "a wild PIKACHU appeared"},
		{name: "YAML acronym", input: "YAML code block", want: "a YAML code block"},
		{name: "Core ML", input: "Core ML function", want: "a Core ML function"},
		{name: "JSON acronym", input: "JSON code block", want: "a JSON code block"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.An(tt.input)
			if got != tt.want {
				t.Errorf("An(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

// TestA verifies that A() is an alias for An()
func TestA(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"cat", "a cat"},
		{"ant", "an ant"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := inflect.A(tt.input)
			if got != tt.want {
				t.Errorf("A(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestPresentParticiple(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		// Empty string
		{name: "empty", input: "", want: ""},

		// Single letter verbs
		{name: "single letter", input: "a", want: "aing"},

		// Already ending in -ing
		{name: "already -ing", input: "running", want: "running"},
		{name: "already -ing sing", input: "sing", want: "singing"},

		// Double consonant (CVC pattern)
		{name: "run", input: "run", want: "running"},
		{name: "sit", input: "sit", want: "sitting"},
		{name: "hit", input: "hit", want: "hitting"},
		{name: "cut", input: "cut", want: "cutting"},
		{name: "stop", input: "stop", want: "stopping"},
		{name: "drop", input: "drop", want: "dropping"},
		{name: "plan", input: "plan", want: "planning"},
		{name: "skip", input: "skip", want: "skipping"},
		{name: "begin", input: "begin", want: "beginning"},
		{name: "occur", input: "occur", want: "occurring"},
		{name: "prefer", input: "prefer", want: "preferring"},
		{name: "admit", input: "admit", want: "admitting"},
		{name: "commit", input: "commit", want: "committing"},
		{name: "regret", input: "regret", want: "regretting"},

		// Drop silent e
		{name: "make", input: "make", want: "making"},
		{name: "take", input: "take", want: "taking"},
		{name: "come", input: "come", want: "coming"},
		{name: "give", input: "give", want: "giving"},
		{name: "have", input: "have", want: "having"},
		{name: "write", input: "write", want: "writing"},
		{name: "live", input: "live", want: "living"},
		{name: "move", input: "move", want: "moving"},
		{name: "hope", input: "hope", want: "hoping"},
		{name: "dance", input: "dance", want: "dancing"},

		// Just add -ing (no changes needed)
		{name: "play", input: "play", want: "playing"},
		{name: "stay", input: "stay", want: "staying"},
		{name: "enjoy", input: "enjoy", want: "enjoying"},
		{name: "show", input: "show", want: "showing"},
		{name: "follow", input: "follow", want: "following"},
		{name: "fix", input: "fix", want: "fixing"},
		{name: "mix", input: "mix", want: "mixing"},
		{name: "go", input: "go", want: "going"},
		{name: "do", input: "do", want: "doing"},
		{name: "eat", input: "eat", want: "eating"},
		{name: "read", input: "read", want: "reading"},
		{name: "think", input: "think", want: "thinking"},
		{name: "walk", input: "walk", want: "walking"},
		{name: "talk", input: "talk", want: "talking"},
		{name: "open", input: "open", want: "opening"},
		{name: "listen", input: "listen", want: "listening"},
		{name: "visit", input: "visit", want: "visiting"},

		// ie -> ying
		{name: "die", input: "die", want: "dying"},
		{name: "lie", input: "lie", want: "lying"},
		{name: "tie", input: "tie", want: "tying"},

		// ee -> eeing
		{name: "see", input: "see", want: "seeing"},
		{name: "flee", input: "flee", want: "fleeing"},
		{name: "agree", input: "agree", want: "agreeing"},
		{name: "free", input: "free", want: "freeing"},

		// be -> being (special vowel + e case)
		{name: "be", input: "be", want: "being"},

		// Words ending in -c (add k)
		{name: "panic", input: "panic", want: "panicking"},
		{name: "picnic", input: "picnic", want: "picnicking"},
		{name: "traffic", input: "traffic", want: "trafficking"},
		{name: "mimic", input: "mimic", want: "mimicking"},
		{name: "frolic", input: "frolic", want: "frolicking"},

		// Words ending in -ye, -oe (keep e)
		{name: "dye", input: "dye", want: "dyeing"},
		{name: "hoe", input: "hoe", want: "hoeing"},
		{name: "toe", input: "toe", want: "toeing"},

		// Words ending in -nge/-inge (keep e)
		{name: "singe", input: "singe", want: "singeing"},

		// Case preservation
		{name: "RUN uppercase", input: "RUN", want: "RUNNING"},
		{name: "Run titlecase", input: "Run", want: "Running"},
		{name: "MAKE uppercase", input: "MAKE", want: "MAKING"},
		{name: "Make titlecase", input: "Make", want: "Making"},
		{name: "DIE uppercase", input: "DIE", want: "DYING"},
		{name: "PANIC uppercase", input: "PANIC", want: "PANICKING"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.PresentParticiple(tt.input)
			if got != tt.want {
				t.Errorf("PresentParticiple(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestSingular(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		// Empty string
		{name: "empty", input: "", want: ""},

		// Regular plurals - remove s
		{name: "cats", input: "cats", want: "cat"},
		{name: "dogs", input: "dogs", want: "dog"},
		{name: "books", input: "books", want: "book"},

		// Words ending in -es after sibilants
		{name: "buses", input: "buses", want: "bus"},
		{name: "classes", input: "classes", want: "class"},
		{name: "bushes", input: "bushes", want: "bush"},
		{name: "churches", input: "churches", want: "church"},
		{name: "boxes", input: "boxes", want: "box"},
		{name: "buzzes", input: "buzzes", want: "buzz"},

		// -ies -> -y (consonant + ies)
		{name: "cities", input: "cities", want: "city"},
		{name: "babies", input: "babies", want: "baby"},
		{name: "flies", input: "flies", want: "fly"},

		// -ys (vowel + y) - just remove s
		{name: "boys", input: "boys", want: "boy"},
		{name: "days", input: "days", want: "day"},
		{name: "keys", input: "keys", want: "key"},

		// Words ending in -ves -> -f or -fe
		{name: "knives", input: "knives", want: "knife"},
		{name: "wives", input: "wives", want: "wife"},
		{name: "lives", input: "lives", want: "life"},
		{name: "leaves", input: "leaves", want: "leaf"},
		{name: "wolves", input: "wolves", want: "wolf"},
		{name: "calves", input: "calves", want: "calf"},
		{name: "halves", input: "halves", want: "half"},

		// Words ending in -oes -> -o
		{name: "heroes", input: "heroes", want: "hero"},
		{name: "potatoes", input: "potatoes", want: "potato"},
		{name: "tomatoes", input: "tomatoes", want: "tomato"},
		{name: "echoes", input: "echoes", want: "echo"},

		// Words ending in -os (just remove s)
		{name: "radios", input: "radios", want: "radio"},
		{name: "studios", input: "studios", want: "studio"},
		{name: "zoos", input: "zoos", want: "zoo"},
		{name: "pianos", input: "pianos", want: "piano"},
		{name: "photos", input: "photos", want: "photo"},

		// Irregular plurals
		{name: "children", input: "children", want: "child"},
		{name: "feet", input: "feet", want: "foot"},
		{name: "teeth", input: "teeth", want: "tooth"},
		{name: "mice", input: "mice", want: "mouse"},
		{name: "women", input: "women", want: "woman"},
		{name: "men", input: "men", want: "man"},
		{name: "people", input: "people", want: "person"},
		{name: "oxen", input: "oxen", want: "ox"},
		{name: "geese", input: "geese", want: "goose"},
		{name: "lice", input: "lice", want: "louse"},
		{name: "dice", input: "dice", want: "die"},

		// Latin/Greek plurals
		{name: "analyses", input: "analyses", want: "analysis"},
		{name: "crises", input: "crises", want: "crisis"},
		{name: "theses", input: "theses", want: "thesis"},
		{name: "cacti", input: "cacti", want: "cactus"},
		{name: "fungi", input: "fungi", want: "fungus"},
		{name: "nuclei", input: "nuclei", want: "nucleus"},
		{name: "bacteria", input: "bacteria", want: "bacterium"},
		{name: "data", input: "data", want: "datum"},
		{name: "media", input: "media", want: "medium"},
		{name: "appendices", input: "appendices", want: "appendix"},
		{name: "indices", input: "indices", want: "index"},
		{name: "criteria", input: "criteria", want: "criterion"},
		{name: "phenomena", input: "phenomena", want: "phenomenon"},

		// Unchanged plurals
		{name: "sheep", input: "sheep", want: "sheep"},
		{name: "deer", input: "deer", want: "deer"},
		{name: "fish", input: "fish", want: "fish"},
		{name: "species", input: "species", want: "species"},
		{name: "series", input: "series", want: "series"},
		{name: "aircraft", input: "aircraft", want: "aircraft"},
		{name: "moose", input: "moose", want: "moose"},

		// Words ending in -men -> -man
		{name: "firemen", input: "firemen", want: "fireman"},
		{name: "policemen", input: "policemen", want: "policeman"},
		{name: "spokesmen", input: "spokesmen", want: "spokesman"},

		// Nationalities ending in -ese (unchanged)
		{name: "Chinese", input: "Chinese", want: "Chinese"},
		{name: "Japanese", input: "Japanese", want: "Japanese"},
		{name: "Portuguese", input: "Portuguese", want: "Portuguese"},

		// Case preservation
		{name: "CATS uppercase", input: "CATS", want: "CAT"},
		{name: "Cats titlecase", input: "Cats", want: "Cat"},
		{name: "Children titlecase", input: "Children", want: "Child"},
		{name: "CHILDREN uppercase", input: "CHILDREN", want: "CHILD"},
		{name: "BOXES uppercase", input: "BOXES", want: "BOX"},
		{name: "Cities titlecase", input: "Cities", want: "City"},
		{name: "MICE uppercase", input: "MICE", want: "MOUSE"},

		// Already singular (should return unchanged)
		{name: "already singular cat", input: "cat", want: "cat"},
		{name: "already singular class", input: "class", want: "class"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Singular(tt.input)
			if got != tt.want {
				t.Errorf("Singular(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestCompare(t *testing.T) {
	tests := []struct {
		name  string
		word1 string
		word2 string
		want  string
	}{
		// Equal words
		{name: "same word", word1: "cat", word2: "cat", want: "eq"},
		{name: "same word uppercase", word1: "CAT", word2: "CAT", want: "eq"},
		{name: "same word mixed case", word1: "Cat", word2: "cat", want: "eq"},
		{name: "same plural", word1: "cats", word2: "cats", want: "eq"},
		{name: "empty strings", word1: "", word2: "", want: "eq"},

		// Singular to plural
		{name: "cat to cats", word1: "cat", word2: "cats", want: "s:p"},
		{name: "dog to dogs", word1: "dog", word2: "dogs", want: "s:p"},
		{name: "box to boxes", word1: "box", word2: "boxes", want: "s:p"},
		{name: "city to cities", word1: "city", word2: "cities", want: "s:p"},
		{name: "child to children", word1: "child", word2: "children", want: "s:p"},
		{name: "mouse to mice", word1: "mouse", word2: "mice", want: "s:p"},
		{name: "knife to knives", word1: "knife", word2: "knives", want: "s:p"},
		{name: "analysis to analyses", word1: "analysis", word2: "analyses", want: "s:p"},
		{name: "index to indices", word1: "index", word2: "indices", want: "s:p"},
		{name: "cactus to cacti", word1: "cactus", word2: "cacti", want: "s:p"},

		// Plural to singular
		{name: "cats to cat", word1: "cats", word2: "cat", want: "p:s"},
		{name: "dogs to dog", word1: "dogs", word2: "dog", want: "p:s"},
		{name: "boxes to box", word1: "boxes", word2: "box", want: "p:s"},
		{name: "cities to city", word1: "cities", word2: "city", want: "p:s"},
		{name: "children to child", word1: "children", word2: "child", want: "p:s"},
		{name: "mice to mouse", word1: "mice", word2: "mouse", want: "p:s"},
		{name: "knives to knife", word1: "knives", word2: "knife", want: "p:s"},
		{name: "analyses to analysis", word1: "analyses", word2: "analysis", want: "p:s"},
		{name: "indices to index", word1: "indices", word2: "index", want: "p:s"},
		{name: "cacti to cactus", word1: "cacti", word2: "cactus", want: "p:s"},

		// Both plurals (different plural forms of same word)
		{name: "indexes to indices", word1: "indexes", word2: "indices", want: "p:p"},
		{name: "indices to indexes", word1: "indices", word2: "indexes", want: "p:p"},

		// Unrelated words
		{name: "cat to dog", word1: "cat", word2: "dog", want: ""},
		{name: "cats to dogs", word1: "cats", word2: "dogs", want: ""},
		{name: "child to mouse", word1: "child", word2: "mouse", want: ""},
		{name: "box to fox", word1: "box", word2: "fox", want: ""},

		// Empty string edge cases
		{name: "empty and word", word1: "", word2: "cat", want: ""},
		{name: "word and empty", word1: "cat", word2: "", want: ""},

		// Case preservation in comparison
		{name: "Cat to Cats", word1: "Cat", word2: "Cats", want: "s:p"},
		{name: "CAT to CATS", word1: "CAT", word2: "CATS", want: "s:p"},
		{name: "CATS to CAT", word1: "CATS", word2: "CAT", want: "p:s"},

		// Unchanged plurals
		{name: "sheep to sheep", word1: "sheep", word2: "sheep", want: "eq"},
		{name: "deer to deer", word1: "deer", word2: "deer", want: "eq"},
		{name: "fish to fish", word1: "fish", word2: "fish", want: "eq"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Compare(tt.word1, tt.word2)
			if got != tt.want {
				t.Errorf("Compare(%q, %q) = %q, want %q", tt.word1, tt.word2, got, tt.want)
			}
		})
	}
}

func TestCompareNouns(t *testing.T) {
	// CompareNouns is an alias for Compare, so we just verify it behaves the same
	tests := []struct {
		name  string
		noun1 string
		noun2 string
		want  string
	}{
		{name: "singular to plural", noun1: "cat", noun2: "cats", want: "s:p"},
		{name: "plural to singular", noun1: "cats", noun2: "cat", want: "p:s"},
		{name: "equal nouns", noun1: "dog", noun2: "dog", want: "eq"},
		{name: "unrelated nouns", noun1: "cat", noun2: "dog", want: ""},
		{name: "irregular plural", noun1: "child", noun2: "children", want: "s:p"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.CompareNouns(tt.noun1, tt.noun2)
			if got != tt.want {
				t.Errorf("CompareNouns(%q, %q) = %q, want %q", tt.noun1, tt.noun2, got, tt.want)
			}
			// Verify it matches Compare() behavior
			compareGot := inflect.Compare(tt.noun1, tt.noun2)
			if got != compareGot {
				t.Errorf("CompareNouns(%q, %q) = %q, but Compare() = %q", tt.noun1, tt.noun2, got, compareGot)
			}
		})
	}
}

func TestPlural(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		// Empty string
		{name: "empty", input: "", want: ""},

		// Regular plurals - add s
		{name: "cat", input: "cat", want: "cats"},
		{name: "dog", input: "dog", want: "dogs"},
		{name: "book", input: "book", want: "books"},

		// Words ending in s, ss, sh, ch, x, z - add es
		{name: "bus", input: "bus", want: "buses"},
		{name: "class", input: "class", want: "classes"},
		{name: "bush", input: "bush", want: "bushes"},
		{name: "church", input: "church", want: "churches"},
		{name: "box", input: "box", want: "boxes"},
		{name: "buzz", input: "buzz", want: "buzzes"},

		// Consonant + y -> ies
		{name: "city", input: "city", want: "cities"},
		{name: "baby", input: "baby", want: "babies"},
		{name: "fly", input: "fly", want: "flies"},

		// Vowel + y -> ys
		{name: "boy", input: "boy", want: "boys"},
		{name: "day", input: "day", want: "days"},
		{name: "key", input: "key", want: "keys"},

		// Words ending in f/fe -> ves
		{name: "knife", input: "knife", want: "knives"},
		{name: "wife", input: "wife", want: "wives"},
		{name: "leaf", input: "leaf", want: "leaves"},
		{name: "wolf", input: "wolf", want: "wolves"},

		// Words ending in f that just take s
		{name: "roof", input: "roof", want: "roofs"},
		{name: "chief", input: "chief", want: "chiefs"},

		// Words ending in o
		{name: "hero", input: "hero", want: "heroes"},
		{name: "potato", input: "potato", want: "potatoes"},
		{name: "tomato", input: "tomato", want: "tomatoes"},
		{name: "echo", input: "echo", want: "echoes"},

		// Words ending in o that take s (vowel + o, exceptions)
		{name: "radio", input: "radio", want: "radios"},
		{name: "studio", input: "studio", want: "studios"},
		{name: "zoo", input: "zoo", want: "zoos"},
		{name: "piano", input: "piano", want: "pianos"},
		{name: "photo", input: "photo", want: "photos"},

		// Irregular plurals
		{name: "child", input: "child", want: "children"},
		{name: "foot", input: "foot", want: "feet"},
		{name: "tooth", input: "tooth", want: "teeth"},
		{name: "mouse", input: "mouse", want: "mice"},
		{name: "woman", input: "woman", want: "women"},
		{name: "man", input: "man", want: "men"},
		{name: "person", input: "person", want: "people"},
		{name: "ox", input: "ox", want: "oxen"},

		// Latin/Greek plurals
		{name: "analysis", input: "analysis", want: "analyses"},
		{name: "crisis", input: "crisis", want: "crises"},
		{name: "thesis", input: "thesis", want: "theses"},
		{name: "cactus", input: "cactus", want: "cacti"},
		{name: "fungus", input: "fungus", want: "fungi"},
		{name: "nucleus", input: "nucleus", want: "nuclei"},
		{name: "bacterium", input: "bacterium", want: "bacteria"},
		{name: "datum", input: "datum", want: "data"},
		{name: "medium", input: "medium", want: "media"},
		{name: "appendix", input: "appendix", want: "appendices"},
		{name: "index", input: "index", want: "indices"},

		// Unchanged plurals
		{name: "sheep", input: "sheep", want: "sheep"},
		{name: "deer", input: "deer", want: "deer"},
		{name: "fish", input: "fish", want: "fish"},
		{name: "species", input: "species", want: "species"},
		{name: "series", input: "series", want: "series"},
		{name: "aircraft", input: "aircraft", want: "aircraft"},

		// Words ending in -man -> -men
		{name: "fireman", input: "fireman", want: "firemen"},
		{name: "policeman", input: "policeman", want: "policemen"},
		{name: "spokesman", input: "spokesman", want: "spokesmen"},

		// Nationalities ending in -ese (unchanged)
		{name: "Chinese", input: "Chinese", want: "Chinese"},
		{name: "Japanese", input: "Japanese", want: "Japanese"},
		{name: "Portuguese", input: "Portuguese", want: "Portuguese"},

		// Case preservation
		{name: "CAT uppercase", input: "CAT", want: "CATS"},
		{name: "Cat titlecase", input: "Cat", want: "Cats"},
		{name: "Child titlecase", input: "Child", want: "Children"},
		{name: "CHILD uppercase", input: "CHILD", want: "CHILDREN"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Plural(tt.input)
			if got != tt.want {
				t.Errorf("Plural(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestCompareVerbs(t *testing.T) {
	// CompareVerbs is a placeholder stub that always returns empty string.
	// These tests verify the function exists and returns "" for any input.
	tests := []struct {
		name  string
		verb1 string
		verb2 string
		want  string
	}{
		{name: "empty strings", verb1: "", verb2: "", want: ""},
		{name: "same verb", verb1: "run", verb2: "run", want: ""},
		{name: "different verbs", verb1: "run", verb2: "walk", want: ""},
		{name: "conjugated forms", verb1: "run", verb2: "running", want: ""},
		{name: "past tense", verb1: "walk", verb2: "walked", want: ""},
		{name: "irregular verb", verb1: "go", verb2: "went", want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.CompareVerbs(tt.verb1, tt.verb2)
			if got != tt.want {
				t.Errorf("CompareVerbs(%q, %q) = %q, want %q", tt.verb1, tt.verb2, got, tt.want)
			}
		})
	}
}

func TestCompareAdjs(t *testing.T) {
	// CompareAdjs is a placeholder stub that always returns empty string.
	// These tests verify the function exists and returns "" for any input.
	tests := []struct {
		name string
		adj1 string
		adj2 string
		want string
	}{
		{name: "empty strings", adj1: "", adj2: "", want: ""},
		{name: "same adjective", adj1: "big", adj2: "big", want: ""},
		{name: "different adjectives", adj1: "big", adj2: "small", want: ""},
		{name: "comparative form", adj1: "big", adj2: "bigger", want: ""},
		{name: "superlative form", adj1: "big", adj2: "biggest", want: ""},
		{name: "irregular adjective", adj1: "good", adj2: "better", want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.CompareAdjs(tt.adj1, tt.adj2)
			if got != tt.want {
				t.Errorf("CompareAdjs(%q, %q) = %q, want %q", tt.adj1, tt.adj2, got, tt.want)
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

func TestDefNoun(t *testing.T) {
	// Reset to defaults after this test
	defer inflect.DefNounReset()

	tests := []struct {
		name     string
		singular string
		plural   string
	}{
		{name: "simple custom", singular: "foo", plural: "foos"},
		{name: "irregular custom", singular: "gizmo", plural: "gizmata"},
		{name: "override builtin", singular: "child", plural: "childs"},
		{name: "unicode custom", singular: "café", plural: "cafés"},
		{name: "empty plural", singular: "widget", plural: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset before each test
			inflect.DefNounReset()

			// Define the custom rule
			inflect.DefNoun(tt.singular, tt.plural)

			// Test Plural()
			got := inflect.Plural(tt.singular)
			if got != tt.plural {
				t.Errorf("After DefNoun(%q, %q): Plural(%q) = %q, want %q",
					tt.singular, tt.plural, tt.singular, got, tt.plural)
			}

			// Test Singular() for reverse lookup
			if tt.plural != "" {
				gotSingular := inflect.Singular(tt.plural)
				if gotSingular != tt.singular {
					t.Errorf("After DefNoun(%q, %q): Singular(%q) = %q, want %q",
						tt.singular, tt.plural, tt.plural, gotSingular, tt.singular)
				}
			}
		})
	}
}

func TestDefNounCasePreservation(t *testing.T) {
	// Reset to defaults after this test
	defer inflect.DefNounReset()

	tests := []struct {
		name         string
		singular     string
		plural       string
		inputWord    string
		wantPlural   string
		inputPlural  string
		wantSingular string
	}{
		{
			name:         "lowercase to lowercase",
			singular:     "foo",
			plural:       "foos",
			inputWord:    "foo",
			wantPlural:   "foos",
			inputPlural:  "foos",
			wantSingular: "foo",
		},
		{
			name:         "titlecase input",
			singular:     "foo",
			plural:       "foos",
			inputWord:    "Foo",
			wantPlural:   "Foos",
			inputPlural:  "Foos",
			wantSingular: "Foo",
		},
		{
			name:         "uppercase input",
			singular:     "foo",
			plural:       "foos",
			inputWord:    "FOO",
			wantPlural:   "FOOS",
			inputPlural:  "FOOS",
			wantSingular: "FOO",
		},
		{
			name:         "irregular custom titlecase",
			singular:     "gizmo",
			plural:       "gizmata",
			inputWord:    "Gizmo",
			wantPlural:   "Gizmata",
			inputPlural:  "Gizmata",
			wantSingular: "Gizmo",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inflect.DefNounReset()
			inflect.DefNoun(tt.singular, tt.plural)

			got := inflect.Plural(tt.inputWord)
			if got != tt.wantPlural {
				t.Errorf("Plural(%q) = %q, want %q", tt.inputWord, got, tt.wantPlural)
			}

			gotSingular := inflect.Singular(tt.inputPlural)
			if gotSingular != tt.wantSingular {
				t.Errorf("Singular(%q) = %q, want %q", tt.inputPlural, gotSingular, tt.wantSingular)
			}
		})
	}
}

func TestUndefNoun(t *testing.T) {
	// Reset to defaults after this test
	defer inflect.DefNounReset()

	tests := []struct {
		name            string
		setupSingular   string
		setupPlural     string
		undefSingular   string
		wantRemoved     bool
		checkWord       string
		wantPluralAfter string // what Plural() returns after undef
	}{
		{
			name:            "remove custom rule",
			setupSingular:   "foo",
			setupPlural:     "foos",
			undefSingular:   "foo",
			wantRemoved:     true,
			checkWord:       "foo",
			wantPluralAfter: "foos", // standard rule applies
		},
		{
			name:            "remove nonexistent rule",
			setupSingular:   "",
			setupPlural:     "",
			undefSingular:   "notdefined",
			wantRemoved:     false,
			checkWord:       "cat",
			wantPluralAfter: "cats",
		},
		{
			name:            "cannot remove builtin",
			setupSingular:   "",
			setupPlural:     "",
			undefSingular:   "child",
			wantRemoved:     false,
			checkWord:       "child",
			wantPluralAfter: "children",
		},
		{
			name:            "case insensitive removal",
			setupSingular:   "bar",
			setupPlural:     "barz",
			undefSingular:   "BAR",
			wantRemoved:     true,
			checkWord:       "bar",
			wantPluralAfter: "bars", // standard rule applies
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inflect.DefNounReset()

			// Setup custom rule if specified
			if tt.setupSingular != "" {
				inflect.DefNoun(tt.setupSingular, tt.setupPlural)
			}

			// Attempt to remove
			removed := inflect.UndefNoun(tt.undefSingular)
			if removed != tt.wantRemoved {
				t.Errorf("UndefNoun(%q) = %v, want %v", tt.undefSingular, removed, tt.wantRemoved)
			}

			// Check pluralization after removal
			got := inflect.Plural(tt.checkWord)
			if got != tt.wantPluralAfter {
				t.Errorf("After UndefNoun: Plural(%q) = %q, want %q", tt.checkWord, got, tt.wantPluralAfter)
			}
		})
	}
}

func TestDefNounReset(t *testing.T) {
	defer inflect.DefNounReset()

	tests := []struct {
		name          string
		customRules   map[string]string // singular -> plural
		overrideRules map[string]string // override built-in rules
		checkWords    map[string]string // word -> expected plural after reset
	}{
		{
			name: "reset custom rules",
			customRules: map[string]string{
				"foo":    "foos",
				"bar":    "barz",
				"widget": "widgetz",
			},
			checkWords: map[string]string{
				"foo":    "foos", // standard rule
				"bar":    "bars", // standard rule
				"widget": "widgets",
				"child":  "children", // builtin preserved
			},
		},
		{
			name: "reset overridden builtins",
			overrideRules: map[string]string{
				"child": "childs",
				"mouse": "mouses",
			},
			checkWords: map[string]string{
				"child": "children", // builtin restored
				"mouse": "mice",     // builtin restored
			},
		},
		{
			name: "reset mixed custom and overrides",
			customRules: map[string]string{
				"gizmo": "gizmata",
			},
			overrideRules: map[string]string{
				"foot": "foots",
			},
			checkWords: map[string]string{
				"gizmo": "gizmoes", // standard rule (consonant + o -> oes)
				"foot":  "feet",    // builtin restored
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Start fresh
			inflect.DefNounReset()

			// Add custom rules
			for singular, plural := range tt.customRules {
				inflect.DefNoun(singular, plural)
			}

			// Override builtins
			for singular, plural := range tt.overrideRules {
				inflect.DefNoun(singular, plural)
			}

			// Reset
			inflect.DefNounReset()

			// Check results
			for word, wantPlural := range tt.checkWords {
				got := inflect.Plural(word)
				if got != wantPlural {
					t.Errorf("After reset: Plural(%q) = %q, want %q", word, got, wantPlural)
				}
			}
		})
	}
}

func TestDefNounIntegration(t *testing.T) {
	// Reset to defaults after this test
	defer inflect.DefNounReset()

	// Test a complete workflow
	t.Run("complete workflow", func(t *testing.T) {
		inflect.DefNounReset()

		// 1. Verify default behavior
		if got := inflect.Plural("child"); got != "children" {
			t.Errorf("Default: Plural(child) = %q, want %q", got, "children")
		}

		// 2. Add custom rule
		inflect.DefNoun("foo", "foozles")
		if got := inflect.Plural("foo"); got != "foozles" {
			t.Errorf("After DefNoun: Plural(foo) = %q, want %q", got, "foozles")
		}
		if got := inflect.Singular("foozles"); got != "foo" {
			t.Errorf("After DefNoun: Singular(foozles) = %q, want %q", got, "foo")
		}

		// 3. Override builtin
		inflect.DefNoun("child", "childs")
		if got := inflect.Plural("child"); got != "childs" {
			t.Errorf("After override: Plural(child) = %q, want %q", got, "childs")
		}

		// 4. Remove custom rule (but not builtin)
		if removed := inflect.UndefNoun("foo"); !removed {
			t.Error("UndefNoun(foo) should return true")
		}
		if got := inflect.Plural("foo"); got != "foos" {
			t.Errorf("After UndefNoun: Plural(foo) = %q, want %q", got, "foos")
		}

		// 5. Cannot remove builtin (even if overridden)
		if removed := inflect.UndefNoun("child"); removed {
			t.Error("UndefNoun(child) should return false for builtin")
		}

		// 6. Reset everything
		inflect.DefNounReset()
		if got := inflect.Plural("child"); got != "children" {
			t.Errorf("After reset: Plural(child) = %q, want %q", got, "children")
		}
		if got := inflect.Plural("foo"); got != "foos" {
			t.Errorf("After reset: Plural(foo) = %q, want %q", got, "foos")
		}
	})
}

func TestDefA(t *testing.T) {
	// Reset to defaults after this test
	defer inflect.DefAReset()

	tests := []struct {
		name  string
		word  string
		input string
		want  string
	}{
		// Force "a" on words that normally take "an"
		{name: "ape forced a", word: "ape", input: "ape", want: "a ape"},
		{name: "apple forced a", word: "apple", input: "apple", want: "a apple"},
		{name: "eagle forced a", word: "eagle", input: "eagle", want: "a eagle"},
		{name: "hour forced a", word: "hour", input: "hour", want: "a hour"},

		// Case insensitive matching
		{name: "Ape titlecase", word: "ape", input: "Ape", want: "a Ape"},
		{name: "APE uppercase", word: "ape", input: "APE", want: "a APE"},
		{name: "define uppercase match lowercase", word: "APE", input: "ape", want: "a ape"},

		// Multi-word input (match first word)
		{name: "ape in phrase", word: "ape", input: "ape costume", want: "a ape costume"},
		{name: "eagle in phrase", word: "eagle", input: "eagle scout", want: "a eagle scout"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inflect.DefAReset()
			inflect.DefA(tt.word)

			got := inflect.An(tt.input)
			if got != tt.want {
				t.Errorf("After DefA(%q): An(%q) = %q, want %q", tt.word, tt.input, got, tt.want)
			}
		})
	}
}

func TestDefAn(t *testing.T) {
	// Reset to defaults after this test
	defer inflect.DefAReset()

	tests := []struct {
		name  string
		word  string
		input string
		want  string
	}{
		// Force "an" on words that normally take "a"
		{name: "hero forced an", word: "hero", input: "hero", want: "an hero"},
		{name: "historic forced an", word: "historic", input: "historic", want: "an historic"},
		{name: "unicorn forced an", word: "unicorn", input: "unicorn", want: "an unicorn"},
		{name: "cat forced an", word: "cat", input: "cat", want: "an cat"},

		// Case insensitive matching
		{name: "Hero titlecase", word: "hero", input: "Hero", want: "an Hero"},
		{name: "HERO uppercase", word: "hero", input: "HERO", want: "an HERO"},
		{name: "define uppercase match lowercase", word: "HERO", input: "hero", want: "an hero"},

		// Multi-word input (match first word)
		{name: "hero in phrase", word: "hero", input: "hero complex", want: "an hero complex"},
		{name: "historic in phrase", word: "historic", input: "historic event", want: "an historic event"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inflect.DefAReset()
			inflect.DefAn(tt.word)

			got := inflect.An(tt.input)
			if got != tt.want {
				t.Errorf("After DefAn(%q): An(%q) = %q, want %q", tt.word, tt.input, got, tt.want)
			}
		})
	}
}

func TestUndefA(t *testing.T) {
	// Reset to defaults after this test
	defer inflect.DefAReset()

	tests := []struct {
		name        string
		setupWord   string
		undefWord   string
		wantRemoved bool
		checkInput  string
		wantAfter   string
	}{
		{
			name:        "remove custom a rule",
			setupWord:   "ape",
			undefWord:   "ape",
			wantRemoved: true,
			checkInput:  "ape",
			wantAfter:   "an ape", // default rule restored
		},
		{
			name:        "remove nonexistent rule",
			setupWord:   "",
			undefWord:   "notdefined",
			wantRemoved: false,
			checkInput:  "ape",
			wantAfter:   "an ape",
		},
		{
			name:        "case insensitive removal",
			setupWord:   "ape",
			undefWord:   "APE",
			wantRemoved: true,
			checkInput:  "ape",
			wantAfter:   "an ape",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inflect.DefAReset()

			if tt.setupWord != "" {
				inflect.DefA(tt.setupWord)
			}

			removed := inflect.UndefA(tt.undefWord)
			if removed != tt.wantRemoved {
				t.Errorf("UndefA(%q) = %v, want %v", tt.undefWord, removed, tt.wantRemoved)
			}

			got := inflect.An(tt.checkInput)
			if got != tt.wantAfter {
				t.Errorf("After UndefA: An(%q) = %q, want %q", tt.checkInput, got, tt.wantAfter)
			}
		})
	}
}

func TestUndefAn(t *testing.T) {
	// Reset to defaults after this test
	defer inflect.DefAReset()

	tests := []struct {
		name        string
		setupWord   string
		undefWord   string
		wantRemoved bool
		checkInput  string
		wantAfter   string
	}{
		{
			name:        "remove custom an rule",
			setupWord:   "hero",
			undefWord:   "hero",
			wantRemoved: true,
			checkInput:  "hero",
			wantAfter:   "a hero", // default rule restored
		},
		{
			name:        "remove nonexistent rule",
			setupWord:   "",
			undefWord:   "notdefined",
			wantRemoved: false,
			checkInput:  "hero",
			wantAfter:   "a hero",
		},
		{
			name:        "case insensitive removal",
			setupWord:   "hero",
			undefWord:   "HERO",
			wantRemoved: true,
			checkInput:  "hero",
			wantAfter:   "a hero",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inflect.DefAReset()

			if tt.setupWord != "" {
				inflect.DefAn(tt.setupWord)
			}

			removed := inflect.UndefAn(tt.undefWord)
			if removed != tt.wantRemoved {
				t.Errorf("UndefAn(%q) = %v, want %v", tt.undefWord, removed, tt.wantRemoved)
			}

			got := inflect.An(tt.checkInput)
			if got != tt.wantAfter {
				t.Errorf("After UndefAn: An(%q) = %q, want %q", tt.checkInput, got, tt.wantAfter)
			}
		})
	}
}

func TestDefAReset(t *testing.T) {
	defer inflect.DefAReset()

	tests := []struct {
		name        string
		customA     []string // words to force "a"
		customAn    []string // words to force "an"
		checkInputs map[string]string
	}{
		{
			name:     "reset custom a rules",
			customA:  []string{"ape", "apple", "eagle"},
			customAn: nil,
			checkInputs: map[string]string{
				"ape":   "an ape",   // restored to default
				"apple": "an apple", // restored to default
				"eagle": "an eagle", // restored to default
				"cat":   "a cat",    // unchanged
			},
		},
		{
			name:     "reset custom an rules",
			customA:  nil,
			customAn: []string{"hero", "cat", "dog"},
			checkInputs: map[string]string{
				"hero": "a hero", // restored to default
				"cat":  "a cat",  // restored to default
				"dog":  "a dog",  // restored to default
				"ape":  "an ape", // unchanged
			},
		},
		{
			name:     "reset mixed custom rules",
			customA:  []string{"ape", "eagle"},
			customAn: []string{"hero", "cat"},
			checkInputs: map[string]string{
				"ape":   "an ape",   // restored to default
				"eagle": "an eagle", // restored to default
				"hero":  "a hero",   // restored to default
				"cat":   "a cat",    // restored to default
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inflect.DefAReset()

			// Add custom "a" rules
			for _, word := range tt.customA {
				inflect.DefA(word)
			}

			// Add custom "an" rules
			for _, word := range tt.customAn {
				inflect.DefAn(word)
			}

			// Reset
			inflect.DefAReset()

			// Check results
			for input, want := range tt.checkInputs {
				got := inflect.An(input)
				if got != want {
					t.Errorf("After reset: An(%q) = %q, want %q", input, got, want)
				}
			}
		})
	}
}

func TestDefADefAnPrecedence(t *testing.T) {
	// Reset to defaults after this test
	defer inflect.DefAReset()

	t.Run("DefA takes precedence over DefAn for same word", func(t *testing.T) {
		inflect.DefAReset()

		// First define as "an"
		inflect.DefAn("test")
		if got := inflect.An("test"); got != "an test" {
			t.Errorf("After DefAn: An(test) = %q, want %q", got, "an test")
		}

		// Then override with "a"
		inflect.DefA("test")
		if got := inflect.An("test"); got != "a test" {
			t.Errorf("After DefA override: An(test) = %q, want %q", got, "a test")
		}
	})

	t.Run("DefAn overrides previous DefA", func(t *testing.T) {
		inflect.DefAReset()

		// First define as "a"
		inflect.DefA("test")
		if got := inflect.An("test"); got != "a test" {
			t.Errorf("After DefA: An(test) = %q, want %q", got, "a test")
		}

		// Then override with "an"
		inflect.DefAn("test")
		if got := inflect.An("test"); got != "an test" {
			t.Errorf("After DefAn override: An(test) = %q, want %q", got, "an test")
		}
	})
}

func TestDefAIntegration(t *testing.T) {
	// Reset to defaults after this test
	defer inflect.DefAReset()

	t.Run("complete workflow", func(t *testing.T) {
		inflect.DefAReset()

		// 1. Verify default behavior
		if got := inflect.An("ape"); got != "an ape" {
			t.Errorf("Default: An(ape) = %q, want %q", got, "an ape")
		}
		if got := inflect.An("hero"); got != "a hero" {
			t.Errorf("Default: An(hero) = %q, want %q", got, "a hero")
		}

		// 2. Add custom "a" rule
		inflect.DefA("ape")
		if got := inflect.An("ape"); got != "a ape" {
			t.Errorf("After DefA: An(ape) = %q, want %q", got, "a ape")
		}

		// 3. Add custom "an" rule
		inflect.DefAn("hero")
		if got := inflect.An("hero"); got != "an hero" {
			t.Errorf("After DefAn: An(hero) = %q, want %q", got, "an hero")
		}

		// 4. Remove custom "a" rule
		if removed := inflect.UndefA("ape"); !removed {
			t.Error("UndefA(ape) should return true")
		}
		if got := inflect.An("ape"); got != "an ape" {
			t.Errorf("After UndefA: An(ape) = %q, want %q", got, "an ape")
		}

		// 5. Remove custom "an" rule
		if removed := inflect.UndefAn("hero"); !removed {
			t.Error("UndefAn(hero) should return true")
		}
		if got := inflect.An("hero"); got != "a hero" {
			t.Errorf("After UndefAn: An(hero) = %q, want %q", got, "a hero")
		}

		// 6. Add multiple rules and reset
		inflect.DefA("ape")
		inflect.DefA("eagle")
		inflect.DefAn("hero")
		inflect.DefAn("cat")

		inflect.DefAReset()

		if got := inflect.An("ape"); got != "an ape" {
			t.Errorf("After reset: An(ape) = %q, want %q", got, "an ape")
		}
		if got := inflect.An("eagle"); got != "an eagle" {
			t.Errorf("After reset: An(eagle) = %q, want %q", got, "an eagle")
		}
		if got := inflect.An("hero"); got != "a hero" {
			t.Errorf("After reset: An(hero) = %q, want %q", got, "a hero")
		}
		if got := inflect.An("cat"); got != "a cat" {
			t.Errorf("After reset: An(cat) = %q, want %q", got, "a cat")
		}
	})
}

func TestDefAPattern(t *testing.T) {
	defer inflect.DefAReset()

	tests := []struct {
		name    string
		pattern string
		inputs  []string
		want    string // "a" for all matches
	}{
		{
			name:    "euro prefix pattern",
			pattern: "euro.*",
			inputs:  []string{"euro", "european", "eurozone", "eurocentric"},
			want:    "a",
		},
		{
			name:    "uni prefix pattern",
			pattern: "uni.*",
			inputs:  []string{"unit", "uniform", "universe"},
			want:    "a",
		},
		{
			name:    "exact match pattern",
			pattern: "apple",
			inputs:  []string{"apple"},
			want:    "a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inflect.DefAReset()

			err := inflect.DefAPattern(tt.pattern)
			if err != nil {
				t.Fatalf("DefAPattern(%q) returned error: %v", tt.pattern, err)
			}

			for _, input := range tt.inputs {
				got := inflect.An(input)
				want := tt.want + " " + input
				if got != want {
					t.Errorf("An(%q) = %q, want %q", input, got, want)
				}
			}
		})
	}
}

func TestDefAnPattern(t *testing.T) {
	defer inflect.DefAReset()

	tests := []struct {
		name    string
		pattern string
		inputs  []string
		want    string // "an" for all matches
	}{
		{
			name:    "honor prefix pattern",
			pattern: "honor.*",
			inputs:  []string{"honor", "honorable", "honorary", "honored"},
			want:    "an",
		},
		{
			name:    "heir prefix pattern",
			pattern: "heir.*",
			inputs:  []string{"heir", "heirloom", "heiress"},
			want:    "an",
		},
		{
			name:    "exact match pattern",
			pattern: "cat",
			inputs:  []string{"cat"},
			want:    "an",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inflect.DefAReset()

			err := inflect.DefAnPattern(tt.pattern)
			if err != nil {
				t.Fatalf("DefAnPattern(%q) returned error: %v", tt.pattern, err)
			}

			for _, input := range tt.inputs {
				got := inflect.An(input)
				want := tt.want + " " + input
				if got != want {
					t.Errorf("An(%q) = %q, want %q", input, got, want)
				}
			}
		})
	}
}

func TestDefAPatternInvalidRegex(t *testing.T) {
	defer inflect.DefAReset()

	err := inflect.DefAPattern("[invalid")
	if err == nil {
		t.Error("DefAPattern with invalid regex should return error")
	}

	err = inflect.DefAnPattern("[invalid")
	if err == nil {
		t.Error("DefAnPattern with invalid regex should return error")
	}
}

func TestUndefAPattern(t *testing.T) {
	defer inflect.DefAReset()

	t.Run("remove existing pattern", func(t *testing.T) {
		inflect.DefAReset()

		// Use "apple" which normally takes "an" but pattern forces "a"
		if err := inflect.DefAPattern("apple.*"); err != nil {
			t.Fatalf("DefAPattern failed: %v", err)
		}
		if got := inflect.An("appleton"); got != "a appleton" {
			t.Errorf("Before UndefAPattern: An(appleton) = %q, want %q", got, "a appleton")
		}

		// Remove pattern
		if removed := inflect.UndefAPattern("apple.*"); !removed {
			t.Error("UndefAPattern should return true for existing pattern")
		}

		// Verify default behavior restored (words starting with vowel get "an")
		if got := inflect.An("appleton"); got != "an appleton" {
			t.Errorf("After UndefAPattern: An(appleton) = %q, want %q", got, "an appleton")
		}
	})

	t.Run("remove non-existing pattern", func(t *testing.T) {
		inflect.DefAReset()

		if removed := inflect.UndefAPattern("nonexistent.*"); removed {
			t.Error("UndefAPattern should return false for non-existing pattern")
		}
	})
}

func TestUndefAnPattern(t *testing.T) {
	defer inflect.DefAReset()

	t.Run("remove existing pattern", func(t *testing.T) {
		inflect.DefAReset()

		// Add pattern and verify it works
		if err := inflect.DefAnPattern("hero.*"); err != nil {
			t.Fatalf("DefAnPattern failed: %v", err)
		}
		if got := inflect.An("heroic"); got != "an heroic" {
			t.Errorf("Before UndefAnPattern: An(heroic) = %q, want %q", got, "an heroic")
		}

		// Remove pattern
		if removed := inflect.UndefAnPattern("hero.*"); !removed {
			t.Error("UndefAnPattern should return true for existing pattern")
		}

		// Verify default behavior restored
		if got := inflect.An("heroic"); got != "a heroic" {
			t.Errorf("After UndefAnPattern: An(heroic) = %q, want %q", got, "a heroic")
		}
	})

	t.Run("remove non-existing pattern", func(t *testing.T) {
		inflect.DefAReset()

		if removed := inflect.UndefAnPattern("nonexistent.*"); removed {
			t.Error("UndefAnPattern should return false for non-existing pattern")
		}
	})
}

func TestDefAResetClearsPatterns(t *testing.T) {
	defer inflect.DefAReset()

	// Add both exact matches and patterns
	inflect.DefA("apple")
	inflect.DefAn("cat")
	if err := inflect.DefAPattern("euro.*"); err != nil {
		t.Fatalf("DefAPattern failed: %v", err)
	}
	if err := inflect.DefAnPattern("honor.*"); err != nil {
		t.Fatalf("DefAnPattern failed: %v", err)
	}

	// Verify patterns are working
	if got := inflect.An("apple"); got != "a apple" {
		t.Errorf("Before reset: An(apple) = %q, want %q", got, "a apple")
	}
	if got := inflect.An("european"); got != "a european" {
		t.Errorf("Before reset: An(european) = %q, want %q", got, "a european")
	}
	if got := inflect.An("honorable"); got != "an honorable" {
		t.Errorf("Before reset: An(honorable) = %q, want %q", got, "an honorable")
	}

	// Reset
	inflect.DefAReset()

	// Verify all patterns are cleared (back to defaults)
	if got := inflect.An("apple"); got != "an apple" {
		t.Errorf("After reset: An(apple) = %q, want %q", got, "an apple")
	}
	// "european" defaults to "a" because "eu" sounds like "you"
	if got := inflect.An("european"); got != "a european" {
		t.Errorf("After reset: An(european) = %q, want %q", got, "a european")
	}
	// "honorable" defaults to "an" because the "h" is silent
	if got := inflect.An("honorable"); got != "an honorable" {
		t.Errorf("After reset: An(honorable) = %q, want %q", got, "an honorable")
	}
}

func TestPatternPrecedence(t *testing.T) {
	defer inflect.DefAReset()

	t.Run("exact word takes precedence over pattern", func(t *testing.T) {
		inflect.DefAReset()

		// Add pattern first
		if err := inflect.DefAnPattern("euro.*"); err != nil {
			t.Fatalf("DefAnPattern failed: %v", err)
		}
		if got := inflect.An("euro"); got != "an euro" {
			t.Errorf("With pattern only: An(euro) = %q, want %q", got, "an euro")
		}

		// Add exact word match - should take precedence
		inflect.DefA("euro")
		if got := inflect.An("euro"); got != "a euro" {
			t.Errorf("With exact word override: An(euro) = %q, want %q", got, "a euro")
		}

		// Other words matching pattern still work
		if got := inflect.An("european"); got != "an european" {
			t.Errorf("Pattern still matches: An(european) = %q, want %q", got, "an european")
		}
	})

	t.Run("DefAPattern takes precedence over DefAnPattern", func(t *testing.T) {
		inflect.DefAReset()

		// Both patterns match "european"
		if err := inflect.DefAnPattern("euro.*"); err != nil {
			t.Fatalf("DefAnPattern failed: %v", err)
		}
		if err := inflect.DefAPattern("europ.*"); err != nil {
			t.Fatalf("DefAPattern failed: %v", err)
		}

		// DefAPattern should take precedence
		if got := inflect.An("european"); got != "a european" {
			t.Errorf("An(european) = %q, want %q", got, "a european")
		}

		// "euro" only matches DefAnPattern
		if got := inflect.An("euro"); got != "an euro" {
			t.Errorf("An(euro) = %q, want %q", got, "an euro")
		}
	})
}

func TestPatternCaseInsensitive(t *testing.T) {
	defer inflect.DefAReset()

	if err := inflect.DefAPattern("euro.*"); err != nil {
		t.Fatalf("DefAPattern failed: %v", err)
	}

	// Pattern should match regardless of case
	tests := []struct {
		input string
		want  string
	}{
		{"euro", "a euro"},
		{"Euro", "a Euro"},
		{"EURO", "a EURO"},
		{"european", "a european"},
		{"European", "a European"},
		{"EUROPEAN", "a EUROPEAN"},
	}

	for _, tt := range tests {
		if got := inflect.An(tt.input); got != tt.want {
			t.Errorf("An(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestDefVerb(t *testing.T) {
	// Reset to defaults after this test
	defer inflect.DefVerbReset()

	tests := []struct {
		name     string
		singular string
		plural   string
	}{
		{name: "simple verb", singular: "run", plural: "runs"},
		{name: "irregular verb", singular: "be", plural: "are"},
		{name: "custom verb", singular: "foo", plural: "foos"},
		{name: "unicode verb", singular: "café", plural: "cafés"},
		{name: "empty plural", singular: "widget", plural: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inflect.DefVerbReset()

			// Define the custom rule - this should not panic
			inflect.DefVerb(tt.singular, tt.plural)

			// Since DefVerb is a stub, we just verify it doesn't panic
			// and the function completes successfully
		})
	}
}

func TestDefVerbCaseInsensitive(t *testing.T) {
	defer inflect.DefVerbReset()

	// DefVerb should store in lowercase
	inflect.DefVerb("Run", "Runs")

	// Verify the lowercase key can be undefined
	if removed := inflect.UndefVerb("run"); !removed {
		t.Error("UndefVerb(run) should return true after DefVerb(Run, Runs)")
	}
}

func TestUndefVerb(t *testing.T) {
	defer inflect.DefVerbReset()

	tests := []struct {
		name          string
		setupSingular string
		setupPlural   string
		undefSingular string
		wantRemoved   bool
	}{
		{
			name:          "remove existing rule",
			setupSingular: "run",
			setupPlural:   "runs",
			undefSingular: "run",
			wantRemoved:   true,
		},
		{
			name:          "remove nonexistent rule",
			setupSingular: "",
			setupPlural:   "",
			undefSingular: "notdefined",
			wantRemoved:   false,
		},
		{
			name:          "case insensitive removal",
			setupSingular: "walk",
			setupPlural:   "walks",
			undefSingular: "WALK",
			wantRemoved:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inflect.DefVerbReset()

			if tt.setupSingular != "" {
				inflect.DefVerb(tt.setupSingular, tt.setupPlural)
			}

			removed := inflect.UndefVerb(tt.undefSingular)
			if removed != tt.wantRemoved {
				t.Errorf("UndefVerb(%q) = %v, want %v", tt.undefSingular, removed, tt.wantRemoved)
			}
		})
	}
}

func TestDefVerbReset(t *testing.T) {
	defer inflect.DefVerbReset()

	// Add some custom rules
	inflect.DefVerb("foo", "foos")
	inflect.DefVerb("bar", "bars")
	inflect.DefVerb("baz", "bazzes")

	// Reset
	inflect.DefVerbReset()

	// Verify rules are gone - UndefVerb should return false
	if removed := inflect.UndefVerb("foo"); removed {
		t.Error("After reset: UndefVerb(foo) should return false")
	}
	if removed := inflect.UndefVerb("bar"); removed {
		t.Error("After reset: UndefVerb(bar) should return false")
	}
	if removed := inflect.UndefVerb("baz"); removed {
		t.Error("After reset: UndefVerb(baz) should return false")
	}
}

func TestDefVerbIntegration(t *testing.T) {
	defer inflect.DefVerbReset()

	t.Run("complete workflow", func(t *testing.T) {
		inflect.DefVerbReset()

		// 1. Add custom rule
		inflect.DefVerb("run", "runs")

		// 2. Verify rule exists (via UndefVerb returning true)
		// Need to re-add since UndefVerb removes it
		inflect.DefVerb("run", "runs")

		// 3. Add more rules
		inflect.DefVerb("walk", "walks")
		inflect.DefVerb("be", "are")

		// 4. Remove one rule
		if removed := inflect.UndefVerb("walk"); !removed {
			t.Error("UndefVerb(walk) should return true")
		}

		// 5. Verify removed rule is gone
		if removed := inflect.UndefVerb("walk"); removed {
			t.Error("UndefVerb(walk) should return false after removal")
		}

		// 6. Other rules still exist
		if removed := inflect.UndefVerb("run"); !removed {
			t.Error("UndefVerb(run) should return true")
		}

		// 7. Reset and verify all gone
		inflect.DefVerb("test", "tests")
		inflect.DefVerbReset()

		if removed := inflect.UndefVerb("test"); removed {
			t.Error("After reset: UndefVerb(test) should return false")
		}
		if removed := inflect.UndefVerb("be"); removed {
			t.Error("After reset: UndefVerb(be) should return false")
		}
	})
}

func TestDefAdj(t *testing.T) {
	// Reset to defaults after this test
	defer inflect.DefAdjReset()

	tests := []struct {
		name     string
		singular string
		plural   string
	}{
		{name: "simple adjective", singular: "big", plural: "bigs"},
		{name: "irregular adjective", singular: "happy", plural: "happies"},
		{name: "custom adjective", singular: "foo", plural: "foos"},
		{name: "unicode adjective", singular: "café", plural: "cafés"},
		{name: "empty plural", singular: "widget", plural: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inflect.DefAdjReset()

			// Define the custom rule - this should not panic
			inflect.DefAdj(tt.singular, tt.plural)

			// Since DefAdj is a stub, we just verify it doesn't panic
			// and the function completes successfully
		})
	}
}

func TestDefAdjCaseInsensitive(t *testing.T) {
	defer inflect.DefAdjReset()

	// DefAdj should store in lowercase
	inflect.DefAdj("Big", "Bigs")

	// Verify the lowercase key can be undefined
	if removed := inflect.UndefAdj("big"); !removed {
		t.Error("UndefAdj(big) should return true after DefAdj(Big, Bigs)")
	}
}

func TestUndefAdj(t *testing.T) {
	defer inflect.DefAdjReset()

	tests := []struct {
		name          string
		setupSingular string
		setupPlural   string
		undefSingular string
		wantRemoved   bool
	}{
		{
			name:          "remove existing rule",
			setupSingular: "big",
			setupPlural:   "bigs",
			undefSingular: "big",
			wantRemoved:   true,
		},
		{
			name:          "remove nonexistent rule",
			setupSingular: "",
			setupPlural:   "",
			undefSingular: "notdefined",
			wantRemoved:   false,
		},
		{
			name:          "case insensitive removal",
			setupSingular: "small",
			setupPlural:   "smalls",
			undefSingular: "SMALL",
			wantRemoved:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inflect.DefAdjReset()

			if tt.setupSingular != "" {
				inflect.DefAdj(tt.setupSingular, tt.setupPlural)
			}

			removed := inflect.UndefAdj(tt.undefSingular)
			if removed != tt.wantRemoved {
				t.Errorf("UndefAdj(%q) = %v, want %v", tt.undefSingular, removed, tt.wantRemoved)
			}
		})
	}
}

func TestDefAdjReset(t *testing.T) {
	defer inflect.DefAdjReset()

	// Add some custom rules
	inflect.DefAdj("foo", "foos")
	inflect.DefAdj("bar", "bars")
	inflect.DefAdj("baz", "bazzes")

	// Reset
	inflect.DefAdjReset()

	// Verify rules are gone - UndefAdj should return false
	if removed := inflect.UndefAdj("foo"); removed {
		t.Error("After reset: UndefAdj(foo) should return false")
	}
	if removed := inflect.UndefAdj("bar"); removed {
		t.Error("After reset: UndefAdj(bar) should return false")
	}
	if removed := inflect.UndefAdj("baz"); removed {
		t.Error("After reset: UndefAdj(baz) should return false")
	}
}

func TestDefAdjIntegration(t *testing.T) {
	defer inflect.DefAdjReset()

	t.Run("complete workflow", func(t *testing.T) {
		inflect.DefAdjReset()

		// 1. Add custom rule
		inflect.DefAdj("big", "bigs")

		// 2. Verify rule exists (via UndefAdj returning true)
		// Need to re-add since UndefAdj removes it
		inflect.DefAdj("big", "bigs")

		// 3. Add more rules
		inflect.DefAdj("small", "smalls")
		inflect.DefAdj("happy", "happies")

		// 4. Remove one rule
		if removed := inflect.UndefAdj("small"); !removed {
			t.Error("UndefAdj(small) should return true")
		}

		// 5. Verify removed rule is gone
		if removed := inflect.UndefAdj("small"); removed {
			t.Error("UndefAdj(small) should return false after removal")
		}

		// 6. Other rules still exist
		if removed := inflect.UndefAdj("big"); !removed {
			t.Error("UndefAdj(big) should return true")
		}

		// 7. Reset and verify all gone
		inflect.DefAdj("test", "tests")
		inflect.DefAdjReset()

		if removed := inflect.UndefAdj("test"); removed {
			t.Error("After reset: UndefAdj(test) should return false")
		}
		if removed := inflect.UndefAdj("happy"); removed {
			t.Error("After reset: UndefAdj(happy) should return false")
		}
	})
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

func TestClassical(t *testing.T) {
	// Clean up after test
	defer inflect.Classical(false)

	tests := []struct {
		name    string
		enabled bool
	}{
		{name: "enable classical mode", enabled: true},
		{name: "disable classical mode", enabled: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inflect.Classical(tt.enabled)
			got := inflect.IsClassical()
			if got != tt.enabled {
				t.Errorf("after Classical(%v): IsClassical() = %v, want %v", tt.enabled, got, tt.enabled)
			}
		})
	}
}

func TestIsClassical(t *testing.T) {
	// Clean up after test
	defer inflect.Classical(false)

	tests := []struct {
		name  string
		setup func()
		want  bool
	}{
		{
			name:  "default state is false",
			setup: func() { inflect.Classical(false) },
			want:  false,
		},
		{
			name:  "after enabling",
			setup: func() { inflect.Classical(true) },
			want:  true,
		},
		{
			name:  "after enabling then disabling",
			setup: func() { inflect.Classical(true); inflect.Classical(false) },
			want:  false,
		},
		{
			name: "after multiple toggles ending enabled",
			setup: func() {
				inflect.Classical(false)
				inflect.Classical(true)
				inflect.Classical(false)
				inflect.Classical(true)
			},
			want: true,
		},
		{
			name: "after multiple toggles ending disabled",
			setup: func() {
				inflect.Classical(true)
				inflect.Classical(false)
				inflect.Classical(true)
				inflect.Classical(false)
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			got := inflect.IsClassical()
			if got != tt.want {
				t.Errorf("IsClassical() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClassicalIntegration(t *testing.T) {
	// Clean up after test
	defer inflect.Classical(false)

	t.Run("complete workflow", func(t *testing.T) {
		// 1. Initial/default state should be false
		inflect.Classical(false)
		if got := inflect.IsClassical(); got != false {
			t.Errorf("Initial IsClassical() = %v, want false", got)
		}

		// 2. Enable classical mode
		inflect.Classical(true)
		if got := inflect.IsClassical(); got != true {
			t.Errorf("After Classical(true): IsClassical() = %v, want true", got)
		}

		// 3. Disable classical mode
		inflect.Classical(false)
		if got := inflect.IsClassical(); got != false {
			t.Errorf("After Classical(false): IsClassical() = %v, want false", got)
		}

		// 4. Toggle multiple times
		inflect.Classical(true)
		inflect.Classical(true) // Setting to same value should work
		if got := inflect.IsClassical(); got != true {
			t.Errorf("After double Classical(true): IsClassical() = %v, want true", got)
		}

		inflect.Classical(false)
		inflect.Classical(false) // Setting to same value should work
		if got := inflect.IsClassical(); got != false {
			t.Errorf("After double Classical(false): IsClassical() = %v, want false", got)
		}
	})
}

func TestClassicalAncient(t *testing.T) {
	// Clean up after test
	defer inflect.ClassicalAncient(false)

	tests := []struct {
		name    string
		enabled bool
		input   string
		want    string
	}{
		{
			name:    "enabled formula becomes formulae",
			enabled: true,
			input:   "formula",
			want:    "formulae",
		},
		{
			name:    "disabled formula becomes formulas",
			enabled: false,
			input:   "formula",
			want:    "formulas",
		},
		{
			name:    "enabled antenna becomes antennae",
			enabled: true,
			input:   "antenna",
			want:    "antennae",
		},
		{
			name:    "disabled antenna becomes antennas",
			enabled: false,
			input:   "antenna",
			want:    "antennas",
		},
		{
			name:    "enabled nebula becomes nebulae",
			enabled: true,
			input:   "nebula",
			want:    "nebulae",
		},
		{
			name:    "disabled nebula becomes nebulas",
			enabled: false,
			input:   "nebula",
			want:    "nebulas",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inflect.ClassicalAncient(tt.enabled)
			got := inflect.Plural(tt.input)
			if got != tt.want {
				t.Errorf("ClassicalAncient(%v): Plural(%q) = %q, want %q", tt.enabled, tt.input, got, tt.want)
			}
		})
	}
}

func TestIsClassicalAncient(t *testing.T) {
	// Clean up after test
	defer inflect.ClassicalAncient(false)

	tests := []struct {
		name  string
		setup func()
		want  bool
	}{
		{
			name:  "default is false",
			setup: func() { inflect.ClassicalAncient(false) },
			want:  false,
		},
		{
			name:  "true after enabling",
			setup: func() { inflect.ClassicalAncient(true) },
			want:  true,
		},
		{
			name:  "false after disabling",
			setup: func() { inflect.ClassicalAncient(true); inflect.ClassicalAncient(false) },
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			got := inflect.IsClassicalAncient()
			if got != tt.want {
				t.Errorf("IsClassicalAncient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClassicalAncientIndependentOfClassicalAll(t *testing.T) {
	// Clean up after test
	defer func() {
		inflect.ClassicalAll(false)
		inflect.ClassicalAncient(false)
	}()

	t.Run("ClassicalAncient works independently of ClassicalAll", func(t *testing.T) {
		// Start with ClassicalAll disabled
		inflect.ClassicalAll(false)

		// Enable only ClassicalAncient
		inflect.ClassicalAncient(true)

		// Verify ClassicalAncient is enabled
		if !inflect.IsClassicalAncient() {
			t.Error("IsClassicalAncient() should be true after ClassicalAncient(true)")
		}

		// Verify formula -> formulae
		if got := inflect.Plural("formula"); got != "formulae" {
			t.Errorf("Plural(\"formula\") = %q, want \"formulae\"", got)
		}
	})

	t.Run("ClassicalAncient can be disabled while ClassicalAll was enabled", func(t *testing.T) {
		// Enable all classical options
		inflect.ClassicalAll(true)

		// Verify it's enabled
		if !inflect.IsClassicalAncient() {
			t.Error("IsClassicalAncient() should be true after ClassicalAll(true)")
		}

		// Disable only ClassicalAncient
		inflect.ClassicalAncient(false)

		// Verify ClassicalAncient is now disabled
		if inflect.IsClassicalAncient() {
			t.Error("IsClassicalAncient() should be false after ClassicalAncient(false)")
		}

		// Verify formula -> formulas (modern form)
		if got := inflect.Plural("formula"); got != "formulas" {
			t.Errorf("Plural(\"formula\") = %q, want \"formulas\"", got)
		}
	})
}

// =============================================================================
// Benchmarks
// =============================================================================

func BenchmarkPlural(b *testing.B) {
	// Test with representative inputs covering different pluralization rules
	benchmarks := []struct {
		name  string
		input string
	}{
		{"regular", "cat"},        // Simple: add -s
		{"sibilant", "box"},       // Add -es (ends in x)
		{"consonant_y", "city"},   // Change y to ies
		{"irregular", "child"},    // Irregular: children
		{"latin", "analysis"},     // Latin: -is to -es
		{"unchanged", "sheep"},    // Unchanged plural
		{"f_to_ves", "knife"},     // f/fe to ves
		{"man_to_men", "fireman"}, // -man to -men
		{"o_adds_es", "hero"},     // consonant + o adds -es
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				inflect.Plural(bm.input)
			}
		})
	}
}

func BenchmarkSingular(b *testing.B) {
	// Test with representative inputs covering different singularization rules
	benchmarks := []struct {
		name  string
		input string
	}{
		{"regular", "cats"},       // Simple: remove -s
		{"sibilant", "boxes"},     // Remove -es
		{"ies_to_y", "cities"},    // Change ies to y
		{"irregular", "children"}, // Irregular: child
		{"latin", "analyses"},     // Latin: -es to -is
		{"unchanged", "sheep"},    // Unchanged
		{"ves_to_f", "knives"},    // ves to f/fe
		{"men_to_man", "firemen"}, // -men to -man
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				inflect.Singular(bm.input)
			}
		})
	}
}

func BenchmarkAn(b *testing.B) {
	// Test with inputs covering different article selection rules
	benchmarks := []struct {
		name  string
		input string
	}{
		{"consonant", "cat"},           // a cat
		{"vowel", "apple"},             // an apple
		{"silent_h", "honest person"},  // an honest
		{"consonant_u", "university"},  // a university (y sound)
		{"abbreviation", "FBI agent"},  // an FBI
		{"phrase", "elegant solution"}, // an elegant
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				inflect.An(bm.input)
			}
		})
	}
}

func BenchmarkOrdinal(b *testing.B) {
	// Test with numbers covering different ordinal suffix rules
	benchmarks := []struct {
		name  string
		input int
	}{
		{"first", 1},         // 1st
		{"second", 2},        // 2nd
		{"third", 3},         // 3rd
		{"fourth", 4},        // 4th
		{"eleventh", 11},     // 11th (special teen)
		{"twenty_first", 21}, // 21st
		{"hundredth", 100},   // 100th
		{"large", 12345},     // 12345th
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				inflect.Ordinal(bm.input)
			}
		})
	}
}

func BenchmarkOrdinalWord(b *testing.B) {
	// Test with numbers of varying complexity
	benchmarks := []struct {
		name  string
		input int
	}{
		{"first", 1},         // first
		{"twelfth", 12},      // twelfth
		{"twenty_first", 21}, // twenty-first
		{"hundredth", 100},   // one hundredth
		{"thousand", 1000},   // one thousandth
		{"complex", 12345},   // complex ordinal word
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				inflect.OrdinalWord(bm.input)
			}
		})
	}
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
			for i := 0; i < b.N; i++ {
				inflect.NumberToWords(bm.input)
			}
		})
	}
}

func BenchmarkJoin(b *testing.B) {
	// Test with slices of varying lengths
	benchmarks := []struct {
		name  string
		input []string
	}{
		{"empty", []string{}},
		{"single", []string{"apple"}},
		{"two", []string{"apple", "banana"}},
		{"three", []string{"apple", "banana", "cherry"}},
		{"five", []string{"one", "two", "three", "four", "five"}},
		{"ten", []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				inflect.Join(bm.input)
			}
		})
	}
}

func BenchmarkPresentParticiple(b *testing.B) {
	// Test with verbs covering different transformation rules
	benchmarks := []struct {
		name  string
		input string
	}{
		{"add_ing", "play"},         // playing (just add -ing)
		{"double_consonant", "run"}, // running (double final consonant)
		{"drop_e", "make"},          // making (drop silent e)
		{"ie_to_y", "die"},          // dying (ie -> y)
		{"ee_keep", "see"},          // seeing (keep ee)
		{"add_k", "panic"},          // panicking (add k before -ing)
		{"already_ing", "sing"},     // singing
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				inflect.PresentParticiple(bm.input)
			}
		})
	}
}

func BenchmarkCompare(b *testing.B) {
	// Test with different comparison scenarios
	benchmarks := []struct {
		name  string
		word1 string
		word2 string
	}{
		{"equal", "cat", "cat"},
		{"singular_to_plural", "cat", "cats"},
		{"plural_to_singular", "cats", "cat"},
		{"irregular_s_to_p", "child", "children"},
		{"irregular_p_to_s", "mice", "mouse"},
		{"unrelated", "cat", "dog"},
		{"unchanged", "sheep", "sheep"},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				inflect.Compare(bm.word1, bm.word2)
			}
		})
	}
}

func TestClassicalAll(t *testing.T) {
	// Reset to defaults after this test
	defer inflect.ClassicalAll(false)

	tests := []struct {
		name    string
		enabled bool
		word    string
		want    string
	}{
		// Classical mode enabled - Latin/Greek plurals
		{name: "formula classical", enabled: true, word: "formula", want: "formulae"},
		{name: "antenna classical", enabled: true, word: "antenna", want: "antennae"},
		{name: "vertebra classical", enabled: true, word: "vertebra", want: "vertebrae"},
		{name: "alumna classical", enabled: true, word: "alumna", want: "alumnae"},
		{name: "larva classical", enabled: true, word: "larva", want: "larvae"},
		{name: "nebula classical", enabled: true, word: "nebula", want: "nebulae"},
		{name: "nova classical", enabled: true, word: "nova", want: "novae"},
		{name: "supernova classical", enabled: true, word: "supernova", want: "supernovae"},
		{name: "octopus classical", enabled: true, word: "octopus", want: "octopodes"},
		{name: "opus classical", enabled: true, word: "opus", want: "opera"},
		{name: "corpus classical", enabled: true, word: "corpus", want: "corpora"},
		{name: "genus classical", enabled: true, word: "genus", want: "genera"},

		// Classical mode disabled - modern English plurals
		{name: "formula modern", enabled: false, word: "formula", want: "formulas"},
		{name: "antenna modern", enabled: false, word: "antenna", want: "antennas"},
		{name: "vertebra modern", enabled: false, word: "vertebra", want: "vertebras"},
		{name: "alumna modern", enabled: false, word: "alumna", want: "alumnas"},
		{name: "larva modern", enabled: false, word: "larva", want: "larvas"},
		{name: "nebula modern", enabled: false, word: "nebula", want: "nebulas"},
		{name: "nova modern", enabled: false, word: "nova", want: "novas"},

		// Regular words should not be affected
		{name: "cat classical", enabled: true, word: "cat", want: "cats"},
		{name: "cat modern", enabled: false, word: "cat", want: "cats"},
		{name: "box classical", enabled: true, word: "box", want: "boxes"},
		{name: "box modern", enabled: false, word: "box", want: "boxes"},

		// Irregular plurals should still work
		{name: "child classical", enabled: true, word: "child", want: "children"},
		{name: "child modern", enabled: false, word: "child", want: "children"},
		{name: "mouse classical", enabled: true, word: "mouse", want: "mice"},
		{name: "mouse modern", enabled: false, word: "mouse", want: "mice"},

		// Case preservation
		{name: "Formula titlecase classical", enabled: true, word: "Formula", want: "Formulae"},
		{name: "FORMULA uppercase classical", enabled: true, word: "FORMULA", want: "FORMULAE"},
		{name: "Formula titlecase modern", enabled: false, word: "Formula", want: "Formulas"},
		{name: "FORMULA uppercase modern", enabled: false, word: "FORMULA", want: "FORMULAS"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inflect.ClassicalAll(tt.enabled)
			got := inflect.Plural(tt.word)
			if got != tt.want {
				t.Errorf("ClassicalAll(%v): Plural(%q) = %q, want %q", tt.enabled, tt.word, got, tt.want)
			}
		})
	}
}

func TestClassicalAliasForClassicalAll(t *testing.T) {
	// Reset to defaults after this test
	defer inflect.ClassicalAll(false)

	// Verify Classical() is an alias for ClassicalAll()
	tests := []struct {
		name    string
		enabled bool
		word    string
		want    string
	}{
		{name: "formula classical via Classical()", enabled: true, word: "formula", want: "formulae"},
		{name: "formula modern via Classical()", enabled: false, word: "formula", want: "formulas"},
		{name: "antenna classical via Classical()", enabled: true, word: "antenna", want: "antennae"},
		{name: "antenna modern via Classical()", enabled: false, word: "antenna", want: "antennas"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inflect.Classical(tt.enabled) // Use Classical(), not ClassicalAll()
			got := inflect.Plural(tt.word)
			if got != tt.want {
				t.Errorf("Classical(%v): Plural(%q) = %q, want %q", tt.enabled, tt.word, got, tt.want)
			}
		})
	}
}

func TestIsClassicalAll(t *testing.T) {
	// Reset to defaults after this test
	defer inflect.ClassicalAll(false)

	tests := []struct {
		name  string
		setup func()
		want  bool
	}{
		{
			name:  "default is false",
			setup: func() { inflect.ClassicalAll(false) },
			want:  false,
		},
		{
			name:  "enabled via ClassicalAll",
			setup: func() { inflect.ClassicalAll(true) },
			want:  true,
		},
		{
			name:  "enabled via Classical alias",
			setup: func() { inflect.Classical(true) },
			want:  true,
		},
		{
			name: "disabled after being enabled",
			setup: func() {
				inflect.ClassicalAll(true)
				inflect.ClassicalAll(false)
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			got := inflect.IsClassicalAll()
			if got != tt.want {
				t.Errorf("IsClassicalAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClassicalAllIntegration(t *testing.T) {
	// Reset to defaults after this test
	defer inflect.ClassicalAll(false)

	t.Run("complete workflow", func(t *testing.T) {
		// 1. Start with default (modern) pluralization
		inflect.ClassicalAll(false)
		if got := inflect.Plural("formula"); got != "formulas" {
			t.Errorf("Default: Plural(formula) = %q, want %q", got, "formulas")
		}
		if inflect.IsClassicalAll() {
			t.Error("Default: IsClassicalAll() should be false")
		}
		if inflect.IsClassical() {
			t.Error("Default: IsClassical() should be false")
		}

		// 2. Enable classical mode
		inflect.ClassicalAll(true)
		if got := inflect.Plural("formula"); got != "formulae" {
			t.Errorf("Classical: Plural(formula) = %q, want %q", got, "formulae")
		}
		if !inflect.IsClassicalAll() {
			t.Error("Classical: IsClassicalAll() should be true")
		}
		if !inflect.IsClassical() {
			t.Error("Classical: IsClassical() should be true")
		}

		// 3. Verify regular words still work
		if got := inflect.Plural("cat"); got != "cats" {
			t.Errorf("Classical: Plural(cat) = %q, want %q", got, "cats")
		}

		// 4. Verify irregular words still work
		if got := inflect.Plural("child"); got != "children" {
			t.Errorf("Classical: Plural(child) = %q, want %q", got, "children")
		}

		// 5. Disable classical mode
		inflect.ClassicalAll(false)
		if got := inflect.Plural("formula"); got != "formulas" {
			t.Errorf("After disable: Plural(formula) = %q, want %q", got, "formulas")
		}
		if inflect.IsClassicalAll() {
			t.Error("After disable: IsClassicalAll() should be false")
		}

		// 6. Use Classical() alias
		inflect.Classical(true)
		if got := inflect.Plural("antenna"); got != "antennae" {
			t.Errorf("Via alias: Plural(antenna) = %q, want %q", got, "antennae")
		}
		if !inflect.IsClassicalAll() {
			t.Error("Via alias: IsClassicalAll() should be true")
		}
	})
}

func TestClassicalPersons(t *testing.T) {
	// Reset to defaults after this test
	defer inflect.ClassicalPersons(false)

	tests := []struct {
		name    string
		enabled bool
		input   string
		want    string
	}{
		// Classical persons enabled: person -> persons
		{name: "persons enabled lowercase", enabled: true, input: "person", want: "persons"},
		{name: "persons enabled titlecase", enabled: true, input: "Person", want: "Persons"},
		{name: "persons enabled uppercase", enabled: true, input: "PERSON", want: "PERSONS"},

		// Classical persons disabled: person -> people (default)
		{name: "people default lowercase", enabled: false, input: "person", want: "people"},
		{name: "people default titlecase", enabled: false, input: "Person", want: "People"},
		{name: "people default uppercase", enabled: false, input: "PERSON", want: "PEOPLE"},

		// Other words should not be affected
		{name: "cat unaffected enabled", enabled: true, input: "cat", want: "cats"},
		{name: "cat unaffected disabled", enabled: false, input: "cat", want: "cats"},
		{name: "child unaffected enabled", enabled: true, input: "child", want: "children"},
		{name: "child unaffected disabled", enabled: false, input: "child", want: "children"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inflect.ClassicalPersons(tt.enabled)
			got := inflect.Plural(tt.input)
			if got != tt.want {
				t.Errorf("ClassicalPersons(%v): Plural(%q) = %q, want %q",
					tt.enabled, tt.input, got, tt.want)
			}
		})
	}
}

func TestIsClassicalPersons(t *testing.T) {
	// Reset to defaults after this test
	defer inflect.ClassicalPersons(false)

	tests := []struct {
		name  string
		setup func()
		want  bool
	}{
		{
			name:  "default is false",
			setup: func() { inflect.ClassicalPersons(false) },
			want:  false,
		},
		{
			name:  "enabled via ClassicalPersons",
			setup: func() { inflect.ClassicalPersons(true) },
			want:  true,
		},
		{
			name:  "enabled via ClassicalAll",
			setup: func() { inflect.ClassicalAll(true) },
			want:  true,
		},
		{
			name: "disabled after being enabled",
			setup: func() {
				inflect.ClassicalPersons(true)
				inflect.ClassicalPersons(false)
			},
			want: false,
		},
		{
			name: "independent of ClassicalAncient",
			setup: func() {
				inflect.ClassicalAncient(true)
				inflect.ClassicalPersons(false)
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset all classical flags before each test
			inflect.ClassicalAll(false)
			tt.setup()
			got := inflect.IsClassicalPersons()
			if got != tt.want {
				t.Errorf("IsClassicalPersons() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClassicalPersonsIntegration(t *testing.T) {
	// Reset to defaults after this test
	defer inflect.ClassicalAll(false)

	t.Run("complete workflow", func(t *testing.T) {
		// 1. Start with default (modern) pluralization
		inflect.ClassicalAll(false)
		if got := inflect.Plural("person"); got != "people" {
			t.Errorf("Default: Plural(person) = %q, want %q", got, "people")
		}
		if inflect.IsClassicalPersons() {
			t.Error("Default: IsClassicalPersons() should be false")
		}

		// 2. Enable classical persons only
		inflect.ClassicalPersons(true)
		if got := inflect.Plural("person"); got != "persons" {
			t.Errorf("ClassicalPersons: Plural(person) = %q, want %q", got, "persons")
		}
		if !inflect.IsClassicalPersons() {
			t.Error("ClassicalPersons: IsClassicalPersons() should be true")
		}

		// 3. Verify classical ancient is still false
		if inflect.IsClassicalAncient() {
			t.Error("ClassicalPersons only: IsClassicalAncient() should be false")
		}
		if got := inflect.Plural("formula"); got != "formulas" {
			t.Errorf("ClassicalPersons only: Plural(formula) = %q, want %q", got, "formulas")
		}

		// 4. Enable ClassicalAll
		inflect.ClassicalAll(true)
		if got := inflect.Plural("person"); got != "persons" {
			t.Errorf("ClassicalAll: Plural(person) = %q, want %q", got, "persons")
		}
		if got := inflect.Plural("formula"); got != "formulae" {
			t.Errorf("ClassicalAll: Plural(formula) = %q, want %q", got, "formulae")
		}

		// 5. Disable persons but keep ancient
		inflect.ClassicalPersons(false)
		if got := inflect.Plural("person"); got != "people" {
			t.Errorf("Persons off, Ancient on: Plural(person) = %q, want %q", got, "people")
		}
		if got := inflect.Plural("formula"); got != "formulae" {
			t.Errorf("Persons off, Ancient on: Plural(formula) = %q, want %q", got, "formulae")
		}

		// 6. Reset all
		inflect.ClassicalAll(false)
		if got := inflect.Plural("person"); got != "people" {
			t.Errorf("After reset: Plural(person) = %q, want %q", got, "people")
		}
		if got := inflect.Plural("formula"); got != "formulas" {
			t.Errorf("After reset: Plural(formula) = %q, want %q", got, "formulas")
		}
	})
}

func TestClassicalNames(t *testing.T) {
	// Reset to defaults after this test
	defer inflect.ClassicalAll(false)

	tests := []struct {
		name       string
		enabled    bool
		input      string
		want       string
		wantGetter bool
	}{
		// ClassicalNames(false) - regular pluralization for names ending in 's'
		{name: "Jones regular", enabled: false, input: "Jones", want: "Joneses", wantGetter: false},
		{name: "Williams regular", enabled: false, input: "Williams", want: "Williamses", wantGetter: false},
		{name: "Hastings regular", enabled: false, input: "Hastings", want: "Hastingses", wantGetter: false},
		{name: "Ross regular", enabled: false, input: "Ross", want: "Rosses", wantGetter: false},
		{name: "Burns regular", enabled: false, input: "Burns", want: "Burnses", wantGetter: false},

		// ClassicalNames(true) - proper names ending in 's' remain unchanged
		{name: "Jones classical", enabled: true, input: "Jones", want: "Jones", wantGetter: true},
		{name: "Williams classical", enabled: true, input: "Williams", want: "Williams", wantGetter: true},
		{name: "Hastings classical", enabled: true, input: "Hastings", want: "Hastings", wantGetter: true},
		{name: "Ross classical", enabled: true, input: "Ross", want: "Ross", wantGetter: true},
		{name: "Burns classical", enabled: true, input: "Burns", want: "Burns", wantGetter: true},

		// Names not ending in 's' should still pluralize normally
		{name: "Mary classical", enabled: true, input: "Mary", want: "Marys", wantGetter: true},
		{name: "Smith classical", enabled: true, input: "Smith", want: "Smiths", wantGetter: true},
		{name: "Johnson classical", enabled: true, input: "Johnson", want: "Johnsons", wantGetter: true},
		{name: "Mary regular", enabled: false, input: "Mary", want: "Marys", wantGetter: false},
		{name: "Smith regular", enabled: false, input: "Smith", want: "Smiths", wantGetter: false},

		// Lowercase words ending in 's' should NOT be treated as proper names
		{name: "bus classical", enabled: true, input: "bus", want: "buses", wantGetter: true},
		{name: "class classical", enabled: true, input: "class", want: "classes", wantGetter: true},
		{name: "boss classical", enabled: true, input: "boss", want: "bosses", wantGetter: true},

		// All uppercase words (acronyms) should NOT be treated as proper names
		{name: "CBS classical", enabled: true, input: "CBS", want: "CBSES", wantGetter: true},
		{name: "GPS classical", enabled: true, input: "GPS", want: "GPSES", wantGetter: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inflect.ClassicalNames(tt.enabled)

			// Test the getter function
			if got := inflect.IsClassicalNames(); got != tt.wantGetter {
				t.Errorf("IsClassicalNames() = %v, want %v", got, tt.wantGetter)
			}

			// Test pluralization
			got := inflect.Plural(tt.input)
			if got != tt.want {
				t.Errorf("Plural(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestClassicalNamesIntegration(t *testing.T) {
	// Reset to defaults after this test
	defer inflect.ClassicalAll(false)

	t.Run("complete workflow", func(t *testing.T) {
		// 1. Start with default (modern) pluralization
		inflect.ClassicalAll(false)
		if got := inflect.Plural("Jones"); got != "Joneses" {
			t.Errorf("Default: Plural(Jones) = %q, want %q", got, "Joneses")
		}
		if inflect.IsClassicalNames() {
			t.Error("Default: IsClassicalNames() should be false")
		}

		// 2. Enable classical names only
		inflect.ClassicalNames(true)
		if got := inflect.Plural("Jones"); got != "Jones" {
			t.Errorf("ClassicalNames: Plural(Jones) = %q, want %q", got, "Jones")
		}
		if !inflect.IsClassicalNames() {
			t.Error("ClassicalNames: IsClassicalNames() should be true")
		}

		// 3. Verify classical ancient is still false
		if inflect.IsClassicalAncient() {
			t.Error("ClassicalNames only: IsClassicalAncient() should be false")
		}
		if got := inflect.Plural("formula"); got != "formulas" {
			t.Errorf("ClassicalNames only: Plural(formula) = %q, want %q", got, "formulas")
		}

		// 4. Regular nouns ending in 's' should still pluralize normally
		if got := inflect.Plural("bus"); got != "buses" {
			t.Errorf("ClassicalNames: Plural(bus) = %q, want %q", got, "buses")
		}

		// 5. Proper names NOT ending in 's' should still pluralize normally
		if got := inflect.Plural("Smith"); got != "Smiths" {
			t.Errorf("ClassicalNames: Plural(Smith) = %q, want %q", got, "Smiths")
		}

		// 6. Enable ClassicalAll
		inflect.ClassicalAll(true)
		if got := inflect.Plural("Jones"); got != "Jones" {
			t.Errorf("ClassicalAll: Plural(Jones) = %q, want %q", got, "Jones")
		}
		if got := inflect.Plural("formula"); got != "formulae" {
			t.Errorf("ClassicalAll: Plural(formula) = %q, want %q", got, "formulae")
		}

		// 7. Disable names but keep ancient
		inflect.ClassicalNames(false)
		if got := inflect.Plural("Jones"); got != "Joneses" {
			t.Errorf("Names off, Ancient on: Plural(Jones) = %q, want %q", got, "Joneses")
		}
		if got := inflect.Plural("formula"); got != "formulae" {
			t.Errorf("Names off, Ancient on: Plural(formula) = %q, want %q", got, "formulae")
		}

		// 8. Reset all
		inflect.ClassicalAll(false)
		if got := inflect.Plural("Jones"); got != "Joneses" {
			t.Errorf("After reset: Plural(Jones) = %q, want %q", got, "Joneses")
		}
		if got := inflect.Plural("formula"); got != "formulas" {
			t.Errorf("After reset: Plural(formula) = %q, want %q", got, "formulas")
		}
	})
}

func TestGender(t *testing.T) {
	// Reset to default before and after test
	inflect.Gender("t")
	defer inflect.Gender("t")

	tests := []struct {
		name     string
		setTo    string
		expected string
	}{
		{name: "masculine", setTo: "m", expected: "m"},
		{name: "feminine", setTo: "f", expected: "f"},
		{name: "neuter", setTo: "n", expected: "n"},
		{name: "they", setTo: "t", expected: "t"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inflect.Gender(tt.setTo)
			if got := inflect.GetGender(); got != tt.expected {
				t.Errorf("Gender(%q): GetGender() = %q, want %q", tt.setTo, got, tt.expected)
			}
		})
	}
}

func TestGetGenderDefault(t *testing.T) {
	// Reset to default before and after test
	inflect.Gender("t")
	defer inflect.Gender("t")

	// Default should be "t" (singular they)
	if got := inflect.GetGender(); got != "t" {
		t.Errorf("GetGender() default = %q, want %q", got, "t")
	}
}

func TestGenderInvalidValues(t *testing.T) {
	// Reset to default before and after test
	inflect.Gender("t")
	defer inflect.Gender("t")

	// Set a valid gender first
	inflect.Gender("m")
	if got := inflect.GetGender(); got != "m" {
		t.Errorf("Gender(m): GetGender() = %q, want %q", got, "m")
	}

	// Invalid values should be ignored
	invalidValues := []string{
		"",       // empty
		"x",      // single invalid char
		"male",   // full word
		"female", // full word
		"M",      // uppercase
		"F",      // uppercase
		"N",      // uppercase
		"T",      // uppercase
		"mm",     // repeated char
		" m",     // with space
		"m ",     // with space
	}

	for _, invalid := range invalidValues {
		t.Run("invalid:"+invalid, func(t *testing.T) {
			inflect.Gender(invalid)
			if got := inflect.GetGender(); got != "m" {
				t.Errorf("Gender(%q): GetGender() = %q, want %q (unchanged)", invalid, got, "m")
			}
		})
	}
}

func TestGenderSequence(t *testing.T) {
	// Reset to default before and after test
	inflect.Gender("t")
	defer inflect.Gender("t")

	// Test setting multiple genders in sequence
	sequence := []struct {
		setTo    string
		expected string
	}{
		{"m", "m"},
		{"f", "f"},
		{"n", "n"},
		{"t", "t"},
		{"invalid", "t"}, // should stay "t"
		{"m", "m"},
		{"", "m"}, // should stay "m"
		{"f", "f"},
	}

	for i, step := range sequence {
		inflect.Gender(step.setTo)
		if got := inflect.GetGender(); got != step.expected {
			t.Errorf("Step %d: Gender(%q): GetGender() = %q, want %q", i, step.setTo, got, step.expected)
		}
	}
}

func TestInflect(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		// Basic plural function
		{
			name:     "simple plural",
			input:    "The plural of cat is plural('cat')",
			expected: "The plural of cat is cats",
		},
		{
			name:     "plural with double quotes",
			input:    `The plural of box is plural("box")`,
			expected: "The plural of box is boxes",
		},
		{
			name:     "plural irregular word",
			input:    "One child, many plural('child')",
			expected: "One child, many children",
		},

		// Plural with count
		{
			name:     "plural with count 1",
			input:    "There is plural('error', 1)",
			expected: "There is error",
		},
		{
			name:     "plural with count 3",
			input:    "There are plural('error', 3)",
			expected: "There are errors",
		},
		{
			name:     "plural with count 0",
			input:    "There are plural('item', 0)",
			expected: "There are items",
		},

		// Singular function
		{
			name:     "simple singular",
			input:    "The singular of cats is singular('cats')",
			expected: "The singular of cats is cat",
		},
		{
			name:     "singular irregular",
			input:    "One of the singular('children')",
			expected: "One of the child",
		},

		// Article functions (an/a)
		{
			name:     "an with vowel",
			input:    "I saw an('apple')",
			expected: "I saw an apple",
		},
		{
			name:     "an with consonant",
			input:    "I saw an('banana')",
			expected: "I saw a banana",
		},
		{
			name:     "a function",
			input:    "This is a('umbrella')",
			expected: "This is an umbrella",
		},
		{
			name:     "a with consonant",
			input:    "This is a('cat')",
			expected: "This is a cat",
		},

		// Ordinal function
		{
			name:     "ordinal 1",
			input:    "This is the ordinal(1) item",
			expected: "This is the 1st item",
		},
		{
			name:     "ordinal 2",
			input:    "This is the ordinal(2) place",
			expected: "This is the 2nd place",
		},
		{
			name:     "ordinal 3",
			input:    "The ordinal(3) time",
			expected: "The 3rd time",
		},
		{
			name:     "ordinal 11",
			input:    "The ordinal(11) hour",
			expected: "The 11th hour",
		},
		{
			name:     "ordinal 21",
			input:    "The ordinal(21) century",
			expected: "The 21st century",
		},

		// Num function
		{
			name:     "num simple",
			input:    "There are num(3) items",
			expected: "There are 3 items",
		},
		{
			name:     "num zero",
			input:    "There are num(0) items",
			expected: "There are 0 items",
		},
		{
			name:     "num large",
			input:    "Count: num(1000)",
			expected: "Count: 1000",
		},

		// Combined examples
		{
			name:     "num with plural",
			input:    "There are num(3) plural('error', 3)",
			expected: "There are 3 errors",
		},
		{
			name:     "num with plural singular case",
			input:    "There is num(1) plural('error', 1)",
			expected: "There is 1 error",
		},
		{
			name:     "multiple functions",
			input:    "I have an('apple') and plural('orange', 2)",
			expected: "I have an apple and oranges",
		},
		{
			name:     "ordinal with plural",
			input:    "The ordinal(1) of plural('mouse', 3)",
			expected: "The 1st of mice",
		},

		// Edge cases
		{
			name:     "no functions",
			input:    "Just plain text",
			expected: "Just plain text",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "unknown function preserved",
			input:    "This unknown('test') stays",
			expected: "This unknown('test') stays",
		},
		{
			name:     "function at start",
			input:    "plural('cat') are cute",
			expected: "cats are cute",
		},
		{
			name:     "function at end",
			input:    "I love plural('dog')",
			expected: "I love dogs",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			if got != tt.expected {
				t.Errorf("Inflect(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestInflectPlural(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"basic noun", "plural('cat')", "cats"},
		{"irregular noun", "plural('mouse')", "mice"},
		{"with count 0", "plural('dog', 0)", "dogs"},
		{"with count 1", "plural('dog', 1)", "dog"},
		{"with count 2", "plural('dog', 2)", "dogs"},
		{"with count -1", "plural('dog', -1)", "dog"},
		{"s ending", "plural('bus')", "buses"},
		{"y ending", "plural('baby')", "babies"},
		{"unchanged", "plural('sheep')", "sheep"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			if got != tt.expected {
				t.Errorf("Inflect(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestInflectSingular(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"basic noun", "singular('cats')", "cat"},
		{"irregular noun", "singular('mice')", "mouse"},
		{"es ending", "singular('boxes')", "box"},
		{"ies ending", "singular('babies')", "baby"},
		{"unchanged", "singular('sheep')", "sheep"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			if got != tt.expected {
				t.Errorf("Inflect(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestInflectArticle(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"an with a", "an('apple')", "an apple"},
		{"an with e", "an('elephant')", "an elephant"},
		{"an with consonant", "an('cat')", "a cat"},
		{"a with vowel", "a('orange')", "an orange"},
		{"a with consonant", "a('dog')", "a dog"},
		{"silent h", "an('hour')", "an hour"},
		{"aspirated h", "an('house')", "a house"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			if got != tt.expected {
				t.Errorf("Inflect(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestInflectOrdinal(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"1st", "ordinal(1)", "1st"},
		{"2nd", "ordinal(2)", "2nd"},
		{"3rd", "ordinal(3)", "3rd"},
		{"4th", "ordinal(4)", "4th"},
		{"11th", "ordinal(11)", "11th"},
		{"12th", "ordinal(12)", "12th"},
		{"13th", "ordinal(13)", "13th"},
		{"21st", "ordinal(21)", "21st"},
		{"22nd", "ordinal(22)", "22nd"},
		{"23rd", "ordinal(23)", "23rd"},
		{"100th", "ordinal(100)", "100th"},
		{"101st", "ordinal(101)", "101st"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			if got != tt.expected {
				t.Errorf("Inflect(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestInflectNum(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"zero", "num(0)", "0"},
		{"positive", "num(42)", "42"},
		{"large", "num(12345)", "12345"},
		{"negative", "num(-5)", "-5"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			if got != tt.expected {
				t.Errorf("Inflect(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}
