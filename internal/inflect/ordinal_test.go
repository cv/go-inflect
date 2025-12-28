package inflect_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	inflect "github.com/cv/go-inflect"
)

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
			assert.Equal(t, tt.want, got)
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
			assert.Equal(t, tt.want, got)
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
			for range b.N {
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
			for range b.N {
				inflect.OrdinalWord(bm.input)
			}
		})
	}
}

func TestOrdinalSuffix(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  string
	}{
		// Basic suffixes
		{name: "1 -> st", input: 1, want: "st"},
		{name: "2 -> nd", input: 2, want: "nd"},
		{name: "3 -> rd", input: 3, want: "rd"},
		{name: "4 -> th", input: 4, want: "th"},
		{name: "5 -> th", input: 5, want: "th"},
		{name: "9 -> th", input: 9, want: "th"},
		{name: "10 -> th", input: 10, want: "th"},

		// Special teens (11, 12, 13 always use "th")
		{name: "11 -> th", input: 11, want: "th"},
		{name: "12 -> th", input: 12, want: "th"},
		{name: "13 -> th", input: 13, want: "th"},
		{name: "14 -> th", input: 14, want: "th"},

		// 21, 22, 23 follow regular rules
		{name: "21 -> st", input: 21, want: "st"},
		{name: "22 -> nd", input: 22, want: "nd"},
		{name: "23 -> rd", input: 23, want: "rd"},
		{name: "24 -> th", input: 24, want: "th"},

		// 111, 112, 113 use "th" (like 11, 12, 13)
		{name: "111 -> th", input: 111, want: "th"},
		{name: "112 -> th", input: 112, want: "th"},
		{name: "113 -> th", input: 113, want: "th"},

		// 121, 122, 123 follow regular rules
		{name: "121 -> st", input: 121, want: "st"},
		{name: "122 -> nd", input: 122, want: "nd"},
		{name: "123 -> rd", input: 123, want: "rd"},

		// Zero
		{name: "0 -> th", input: 0, want: "th"},

		// Negative numbers (use absolute value for suffix)
		{name: "-1 -> st", input: -1, want: "st"},
		{name: "-2 -> nd", input: -2, want: "nd"},
		{name: "-11 -> th", input: -11, want: "th"},
		{name: "-21 -> st", input: -21, want: "st"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.OrdinalSuffix(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestIsOrdinal(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		// Numeric ordinals
		{name: "1st is ordinal", input: "1st", want: true},
		{name: "2nd is ordinal", input: "2nd", want: true},
		{name: "3rd is ordinal", input: "3rd", want: true},
		{name: "4th is ordinal", input: "4th", want: true},
		{name: "11th is ordinal", input: "11th", want: true},
		{name: "21st is ordinal", input: "21st", want: true},
		{name: "100th is ordinal", input: "100th", want: true},

		// Word ordinals
		{name: "first is ordinal", input: "first", want: true},
		{name: "second is ordinal", input: "second", want: true},
		{name: "third is ordinal", input: "third", want: true},
		{name: "fourth is ordinal", input: "fourth", want: true},
		{name: "fifth is ordinal", input: "fifth", want: true},
		{name: "eleventh is ordinal", input: "eleventh", want: true},
		{name: "twelfth is ordinal", input: "twelfth", want: true},
		{name: "twentieth is ordinal", input: "twentieth", want: true},
		{name: "twenty-first is ordinal", input: "twenty-first", want: true},
		{name: "zeroth is ordinal", input: "zeroth", want: true},

		// Case variations
		{name: "First is ordinal", input: "First", want: true},
		{name: "SECOND is ordinal", input: "SECOND", want: true},
		{name: "Twenty-First is ordinal", input: "Twenty-First", want: true},

		// Not ordinals
		{name: "1 is not ordinal", input: "1", want: false},
		{name: "one is not ordinal", input: "one", want: false},
		{name: "two is not ordinal", input: "two", want: false},
		{name: "twenty is not ordinal", input: "twenty", want: false},
		{name: "twenty-one is not ordinal", input: "twenty-one", want: false},
		{name: "hundred is not ordinal", input: "hundred", want: false},
		{name: "cat is not ordinal", input: "cat", want: false},
		{name: "empty is not ordinal", input: "", want: false},

		// Edge cases - words that end in ordinal-like suffixes but aren't
		{name: "north is not ordinal", input: "north", want: false},
		{name: "month is not ordinal", input: "month", want: false},
		{name: "earth is not ordinal", input: "earth", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.IsOrdinal(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestOrdinalToCardinal(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		// Numeric ordinals to numbers
		{name: "1st -> 1", input: "1st", want: "1"},
		{name: "2nd -> 2", input: "2nd", want: "2"},
		{name: "3rd -> 3", input: "3rd", want: "3"},
		{name: "4th -> 4", input: "4th", want: "4"},
		{name: "11th -> 11", input: "11th", want: "11"},
		{name: "21st -> 21", input: "21st", want: "21"},
		{name: "100th -> 100", input: "100th", want: "100"},
		{name: "1000th -> 1000", input: "1000th", want: "1000"},

		// Word ordinals to cardinal words
		{name: "first -> one", input: "first", want: "one"},
		{name: "second -> two", input: "second", want: "two"},
		{name: "third -> three", input: "third", want: "three"},
		{name: "fourth -> four", input: "fourth", want: "four"},
		{name: "fifth -> five", input: "fifth", want: "five"},
		{name: "sixth -> six", input: "sixth", want: "six"},
		{name: "seventh -> seven", input: "seventh", want: "seven"},
		{name: "eighth -> eight", input: "eighth", want: "eight"},
		{name: "ninth -> nine", input: "ninth", want: "nine"},
		{name: "tenth -> ten", input: "tenth", want: "ten"},
		{name: "eleventh -> eleven", input: "eleventh", want: "eleven"},
		{name: "twelfth -> twelve", input: "twelfth", want: "twelve"},
		{name: "twentieth -> twenty", input: "twentieth", want: "twenty"},
		{name: "twenty-first -> twenty-one", input: "twenty-first", want: "twenty-one"},
		{name: "zeroth -> zero", input: "zeroth", want: "zero"},

		// Case preservation
		{name: "First -> One", input: "First", want: "One"},
		{name: "SECOND -> TWO", input: "SECOND", want: "TWO"},
		{name: "Twenty-First -> Twenty-One", input: "Twenty-First", want: "Twenty-One"},

		// Not ordinals - return unchanged
		{name: "one unchanged", input: "one", want: "one"},
		{name: "twenty unchanged", input: "twenty", want: "twenty"},
		{name: "cat unchanged", input: "cat", want: "cat"},
		{name: "42 unchanged", input: "42", want: "42"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.OrdinalToCardinal(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func BenchmarkOrdinalSuffix(b *testing.B) {
	inputs := []int{1, 2, 3, 11, 12, 13, 21, 22, 23, 100, 111, 121}
	b.ResetTimer()
	for i := range b.N {
		inflect.OrdinalSuffix(inputs[i%len(inputs)])
	}
}

func BenchmarkIsOrdinal(b *testing.B) {
	inputs := []string{"1st", "first", "twenty-first", "one", "cat"}
	b.ResetTimer()
	for i := range b.N {
		inflect.IsOrdinal(inputs[i%len(inputs)])
	}
}

func BenchmarkOrdinalToCardinal(b *testing.B) {
	inputs := []string{"1st", "first", "twenty-first", "SECOND"}
	b.ResetTimer()
	for i := range b.N {
		inflect.OrdinalToCardinal(inputs[i%len(inputs)])
	}
}
