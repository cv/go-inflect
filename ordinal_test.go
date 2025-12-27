package inflect_test

import (
	"testing"

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
