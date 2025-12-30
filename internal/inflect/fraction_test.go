package inflect_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	inflect "github.com/cv/go-inflect/v2"
)

func TestFractionToWords(t *testing.T) {
	tests := []struct {
		name        string
		numerator   int
		denominator int
		want        string
	}{
		// Halves (denominator 2)
		{name: "one half", numerator: 1, denominator: 2, want: "one half"},
		{name: "two halves", numerator: 2, denominator: 2, want: "two halves"},
		{name: "three halves", numerator: 3, denominator: 2, want: "three halves"},
		{name: "five halves", numerator: 5, denominator: 2, want: "five halves"},

		// Thirds (denominator 3)
		{name: "one third", numerator: 1, denominator: 3, want: "one third"},
		{name: "two thirds", numerator: 2, denominator: 3, want: "two thirds"},
		{name: "four thirds", numerator: 4, denominator: 3, want: "four thirds"},

		// Quarters (denominator 4)
		{name: "one quarter", numerator: 1, denominator: 4, want: "one quarter"},
		{name: "two quarters", numerator: 2, denominator: 4, want: "two quarters"},
		{name: "three quarters", numerator: 3, denominator: 4, want: "three quarters"},
		{name: "five quarters", numerator: 5, denominator: 4, want: "five quarters"},

		// Fifths and beyond (use ordinal forms)
		{name: "one fifth", numerator: 1, denominator: 5, want: "one fifth"},
		{name: "two fifths", numerator: 2, denominator: 5, want: "two fifths"},
		{name: "three fifths", numerator: 3, denominator: 5, want: "three fifths"},

		{name: "one sixth", numerator: 1, denominator: 6, want: "one sixth"},
		{name: "five sixths", numerator: 5, denominator: 6, want: "five sixths"},

		{name: "one seventh", numerator: 1, denominator: 7, want: "one seventh"},
		{name: "three sevenths", numerator: 3, denominator: 7, want: "three sevenths"},

		{name: "one eighth", numerator: 1, denominator: 8, want: "one eighth"},
		{name: "five eighths", numerator: 5, denominator: 8, want: "five eighths"},
		{name: "seven eighths", numerator: 7, denominator: 8, want: "seven eighths"},

		{name: "one ninth", numerator: 1, denominator: 9, want: "one ninth"},
		{name: "four ninths", numerator: 4, denominator: 9, want: "four ninths"},

		{name: "one tenth", numerator: 1, denominator: 10, want: "one tenth"},
		{name: "three tenths", numerator: 3, denominator: 10, want: "three tenths"},
		{name: "seven tenths", numerator: 7, denominator: 10, want: "seven tenths"},

		// Larger denominators
		{name: "one eleventh", numerator: 1, denominator: 11, want: "one eleventh"},
		{name: "five elevenths", numerator: 5, denominator: 11, want: "five elevenths"},

		{name: "one twelfth", numerator: 1, denominator: 12, want: "one twelfth"},
		{name: "seven twelfths", numerator: 7, denominator: 12, want: "seven twelfths"},

		{name: "one twentieth", numerator: 1, denominator: 20, want: "one twentieth"},
		{name: "three twentieths", numerator: 3, denominator: 20, want: "three twentieths"},

		{name: "one twenty-first", numerator: 1, denominator: 21, want: "one twenty-first"},
		{name: "five twenty-firsts", numerator: 5, denominator: 21, want: "five twenty-firsts"},

		// Hundredths
		{name: "one hundredth", numerator: 1, denominator: 100, want: "one hundredth"},
		{name: "seven hundredths", numerator: 7, denominator: 100, want: "seven hundredths"},
		{name: "twenty-three hundredths", numerator: 23, denominator: 100, want: "twenty-three hundredths"},
		{name: "ninety-nine hundredths", numerator: 99, denominator: 100, want: "ninety-nine hundredths"},

		// Thousandths
		{name: "one thousandth", numerator: 1, denominator: 1000, want: "one thousandth"},
		{name: "three thousandths", numerator: 3, denominator: 1000, want: "three thousandths"},

		// Edge case: denominator 1 (whole numbers)
		{name: "one over one", numerator: 1, denominator: 1, want: "one"},
		{name: "five over one", numerator: 5, denominator: 1, want: "five"},
		{name: "twenty-three over one", numerator: 23, denominator: 1, want: "twenty-three"},
		{name: "one hundred over one", numerator: 100, denominator: 1, want: "one hundred"},

		// Edge case: denominator 0 (error case)
		{name: "one over zero", numerator: 1, denominator: 0, want: ""},
		{name: "five over zero", numerator: 5, denominator: 0, want: ""},

		// Negative numerators
		{name: "negative one half", numerator: -1, denominator: 2, want: "negative one half"},
		{name: "negative three quarters", numerator: -3, denominator: 4, want: "negative three quarters"},
		{name: "negative two thirds", numerator: -2, denominator: 3, want: "negative two thirds"},
		{name: "negative five eighths", numerator: -5, denominator: 8, want: "negative five eighths"},

		// Negative denominators
		{name: "one over negative two", numerator: 1, denominator: -2, want: "negative one half"},
		{name: "three over negative four", numerator: 3, denominator: -4, want: "negative three quarters"},

		// Both negative (should be positive)
		{name: "negative one over negative two", numerator: -1, denominator: -2, want: "one half"},
		{name: "negative three over negative four", numerator: -3, denominator: -4, want: "three quarters"},

		// Zero numerator
		{name: "zero halves", numerator: 0, denominator: 2, want: "zero halves"},
		{name: "zero thirds", numerator: 0, denominator: 3, want: "zero thirds"},
		{name: "zero quarters", numerator: 0, denominator: 4, want: "zero quarters"},
		{name: "zero fifths", numerator: 0, denominator: 5, want: "zero fifths"},

		// Large numerators
		{name: "eleven twelfths", numerator: 11, denominator: 12, want: "eleven twelfths"},
		{name: "twenty-one thirty-seconds", numerator: 21, denominator: 32, want: "twenty-one thirty-seconds"},
		{name: "ninety-nine hundredths", numerator: 99, denominator: 100, want: "ninety-nine hundredths"},
		{name: "one hundred one thousandths", numerator: 101, denominator: 1000, want: "one hundred one thousandths"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.FractionToWords(tt.numerator, tt.denominator)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFractionToWordsWithFourths(t *testing.T) {
	tests := []struct {
		name        string
		numerator   int
		denominator int
		want        string
	}{
		{name: "one fourth", numerator: 1, denominator: 4, want: "one fourth"},
		{name: "two fourths", numerator: 2, denominator: 4, want: "two fourths"},
		{name: "three fourths", numerator: 3, denominator: 4, want: "three fourths"},
		{name: "five fourths", numerator: 5, denominator: 4, want: "five fourths"},
		{name: "negative three fourths", numerator: -3, denominator: 4, want: "negative three fourths"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.FractionToWordsWithFourths(tt.numerator, tt.denominator)
			assert.Equal(t, tt.want, got)
		})
	}
}

// TestFractionToWordsExamples tests the specific examples from the issue.
func TestFractionToWordsExamples(t *testing.T) {
	examples := []struct {
		numerator   int
		denominator int
		want        string
	}{
		{1, 2, "one half"},
		{3, 2, "three halves"},
		{1, 4, "one quarter"},
		{3, 4, "three quarters"},
		{2, 3, "two thirds"},
		{5, 8, "five eighths"},
		{1, 100, "one hundredth"},
		{7, 100, "seven hundredths"},
	}

	for _, ex := range examples {
		t.Run("", func(t *testing.T) {
			got := inflect.FractionToWords(ex.numerator, ex.denominator)
			assert.Equal(t, ex.want, got)
		})
	}
}

func BenchmarkFractionToWords(b *testing.B) {
	benchmarks := []struct {
		name string
		num  int
		den  int
	}{
		{"half", 1, 2},
		{"quarter", 1, 4},
		{"three_quarters", 3, 4},
		{"third", 1, 3},
		{"two_thirds", 2, 3},
		{"fifth", 1, 5},
		{"complex", 7, 16},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for range b.N {
				inflect.FractionToWords(bm.num, bm.den)
			}
		})
	}
}

func BenchmarkFractionToWordsWithFourths(b *testing.B) {
	benchmarks := []struct {
		name string
		num  int
		den  int
	}{
		{"half", 1, 2},
		{"fourth", 1, 4},
		{"three_fourths", 3, 4},
		{"third", 1, 3},
		{"fifth", 1, 5},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for range b.N {
				inflect.FractionToWordsWithFourths(bm.num, bm.den)
			}
		})
	}
}
