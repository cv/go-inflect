package inflect

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountingWord(t *testing.T) {
	tests := []struct {
		n    int
		want string
	}{
		// Special words
		{1, "once"},
		{2, "twice"},
		{3, "thrice"},

		// Zero
		{0, "zero times"},

		// Regular numbers (4+)
		{4, "four times"},
		{5, "five times"},
		{6, "six times"},
		{7, "seven times"},
		{8, "eight times"},
		{9, "nine times"},
		{10, "ten times"},
		{11, "eleven times"},
		{12, "twelve times"},
		{13, "thirteen times"},
		{19, "nineteen times"},
		{20, "twenty times"},
		{21, "twenty-one times"},
		{42, "forty-two times"},
		{100, "one hundred times"},
		{101, "one hundred one times"},
		{1000, "one thousand times"},

		// Negative numbers
		{-1, "negative once"},
		{-2, "negative twice"},
		{-3, "negative thrice"},
		{-4, "negative four times"},
		{-10, "negative ten times"},
		{-100, "negative one hundred times"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			got := CountingWord(tt.n)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCountingWordWithOptions(t *testing.T) {
	tests := []struct {
		n         int
		useThrice bool
		want      string
	}{
		// With useThrice = true (default behavior)
		{1, true, "once"},
		{2, true, "twice"},
		{3, true, "thrice"},
		{4, true, "four times"},
		{0, true, "zero times"},
		{-1, true, "negative once"},
		{-2, true, "negative twice"},
		{-3, true, "negative thrice"},

		// With useThrice = false
		{1, false, "once"},
		{2, false, "twice"},
		{3, false, "three times"},
		{4, false, "four times"},
		{0, false, "zero times"},
		{-1, false, "negative once"},
		{-2, false, "negative twice"},
		{-3, false, "negative three times"},

		// Larger numbers (useThrice doesn't affect them)
		{10, true, "ten times"},
		{10, false, "ten times"},
		{100, true, "one hundred times"},
		{100, false, "one hundred times"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			got := CountingWordWithOptions(tt.n, tt.useThrice)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCountingWordThreshold(t *testing.T) {
	tests := []struct {
		n         int
		threshold int
		want      string
	}{
		// Below threshold: use words
		{1, 10, "once"},
		{2, 10, "twice"},
		{3, 10, "thrice"},
		{4, 10, "four times"},
		{5, 10, "five times"},
		{9, 10, "nine times"},

		// At or above threshold: use digits
		{10, 10, "10 times"},
		{11, 10, "11 times"},
		{15, 10, "15 times"},
		{100, 10, "100 times"},
		{1000, 10, "1000 times"},

		// Threshold of 100
		{50, 100, "fifty times"},
		{99, 100, "ninety-nine times"},
		{100, 100, "100 times"},
		{101, 100, "101 times"},

		// Threshold of 1 (always use digits except for negative)
		{1, 1, "1 times"},
		{2, 1, "2 times"},
		{0, 1, "zero times"}, // 0 < 1, so words

		// Zero threshold (always use digits)
		{0, 0, "0 times"},
		{1, 0, "1 times"},

		// Negative threshold (effectively always use digits for positive)
		{1, -1, "1 times"},
		{5, -5, "5 times"},

		// Negative numbers below threshold
		{-1, 10, "negative once"},
		{-2, 10, "negative twice"},
		{-3, 10, "negative thrice"},
		{-5, 10, "negative five times"},
		{-9, 10, "negative nine times"},

		// Negative numbers at or above threshold
		{-10, 10, "-10 times"},
		{-15, 10, "-15 times"},
		{-100, 10, "-100 times"},

		// Edge case: threshold equals special numbers
		{1, 2, "once"},
		{2, 2, "2 times"},
		{3, 3, "3 times"},
		{3, 4, "thrice"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			got := CountingWordThreshold(tt.n, tt.threshold)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCountingWordLargeNumbers(t *testing.T) {
	tests := []struct {
		n    int
		want string
	}{
		{1000, "one thousand times"},
		{1001, "one thousand one times"},
		{1234, "one thousand two hundred thirty-four times"},
		{1000000, "one million times"},
		{1000001, "one million one times"},
		{1000000000, "one billion times"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			got := CountingWord(tt.n)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCountingWordEdgeCases(t *testing.T) {
	// Test that CountingWord with default options equals CountingWordWithOptions(n, true)
	for n := -10; n <= 10; n++ {
		got := CountingWord(n)
		want := CountingWordWithOptions(n, true)
		assert.Equal(t, want, got, "CountingWord(%d) should equal CountingWordWithOptions(%d, true)", n, n)
	}
}
