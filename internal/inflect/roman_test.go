package inflect_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	inflect "github.com/cv/go-inflect/v2"
)

func TestIntToRoman(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  string
	}{
		// Basic numerals
		{name: "one", input: 1, want: "I"},
		{name: "two", input: 2, want: "II"},
		{name: "three", input: 3, want: "III"},
		{name: "four", input: 4, want: "IV"},
		{name: "five", input: 5, want: "V"},
		{name: "six", input: 6, want: "VI"},
		{name: "seven", input: 7, want: "VII"},
		{name: "eight", input: 8, want: "VIII"},
		{name: "nine", input: 9, want: "IX"},
		{name: "ten", input: 10, want: "X"},

		// Teens and twenties
		{name: "eleven", input: 11, want: "XI"},
		{name: "fourteen", input: 14, want: "XIV"},
		{name: "nineteen", input: 19, want: "XIX"},
		{name: "twenty", input: 20, want: "XX"},
		{name: "twenty-one", input: 21, want: "XXI"},
		{name: "twenty-four", input: 24, want: "XXIV"},
		{name: "twenty-nine", input: 29, want: "XXIX"},

		// Thirties to nineties
		{name: "thirty", input: 30, want: "XXX"},
		{name: "forty", input: 40, want: "XL"},
		{name: "forty-four", input: 44, want: "XLIV"},
		{name: "forty-nine", input: 49, want: "XLIX"},
		{name: "fifty", input: 50, want: "L"},
		{name: "sixty", input: 60, want: "LX"},
		{name: "seventy", input: 70, want: "LXX"},
		{name: "eighty", input: 80, want: "LXXX"},
		{name: "eighty-nine", input: 89, want: "LXXXIX"},
		{name: "ninety", input: 90, want: "XC"},
		{name: "ninety-nine", input: 99, want: "XCIX"},

		// Hundreds
		{name: "one hundred", input: 100, want: "C"},
		{name: "two hundred", input: 200, want: "CC"},
		{name: "three hundred", input: 300, want: "CCC"},
		{name: "four hundred", input: 400, want: "CD"},
		{name: "five hundred", input: 500, want: "D"},
		{name: "six hundred", input: 600, want: "DC"},
		{name: "seven hundred", input: 700, want: "DCC"},
		{name: "eight hundred", input: 800, want: "DCCC"},
		{name: "nine hundred", input: 900, want: "CM"},

		// Thousands
		{name: "one thousand", input: 1000, want: "M"},
		{name: "two thousand", input: 2000, want: "MM"},
		{name: "three thousand", input: 3000, want: "MMM"},

		// Complex numbers
		{name: "1984", input: 1984, want: "MCMLXXXIV"},
		{name: "2024", input: 2024, want: "MMXXIV"},
		{name: "2025", input: 2025, want: "MMXXV"},
		{name: "1666", input: 1666, want: "MDCLXVI"},
		{name: "3999", input: 3999, want: "MMMCMXCIX"},

		// Edge cases
		{name: "zero returns empty", input: 0, want: ""},
		{name: "negative returns empty", input: -1, want: ""},
		{name: "over 3999 returns empty", input: 4000, want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.IntToRoman(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int
		wantErr bool
	}{
		// Basic numerals
		{name: "I", input: "I", want: 1},
		{name: "II", input: "II", want: 2},
		{name: "III", input: "III", want: 3},
		{name: "IV", input: "IV", want: 4},
		{name: "V", input: "V", want: 5},
		{name: "VI", input: "VI", want: 6},
		{name: "VII", input: "VII", want: 7},
		{name: "VIII", input: "VIII", want: 8},
		{name: "IX", input: "IX", want: 9},
		{name: "X", input: "X", want: 10},

		// Teens and twenties
		{name: "XI", input: "XI", want: 11},
		{name: "XIV", input: "XIV", want: 14},
		{name: "XIX", input: "XIX", want: 19},
		{name: "XX", input: "XX", want: 20},
		{name: "XXI", input: "XXI", want: 21},
		{name: "XXIV", input: "XXIV", want: 24},
		{name: "XXIX", input: "XXIX", want: 29},

		// Larger numbers
		{name: "XL", input: "XL", want: 40},
		{name: "L", input: "L", want: 50},
		{name: "XC", input: "XC", want: 90},
		{name: "C", input: "C", want: 100},
		{name: "CD", input: "CD", want: 400},
		{name: "D", input: "D", want: 500},
		{name: "CM", input: "CM", want: 900},
		{name: "M", input: "M", want: 1000},

		// Complex numbers
		{name: "MCMLXXXIV", input: "MCMLXXXIV", want: 1984},
		{name: "MMXXIV", input: "MMXXIV", want: 2024},
		{name: "MMXXV", input: "MMXXV", want: 2025},
		{name: "MDCLXVI", input: "MDCLXVI", want: 1666},
		{name: "MMMCMXCIX", input: "MMMCMXCIX", want: 3999},

		// Lowercase input should work
		{name: "lowercase iv", input: "iv", want: 4},
		{name: "lowercase xiv", input: "xiv", want: 14},
		{name: "lowercase mcmlxxxiv", input: "mcmlxxxiv", want: 1984},
		{name: "mixed case MmXxIv", input: "MmXxIv", want: 2024},

		// Empty string
		{name: "empty string", input: "", want: 0, wantErr: true},

		// Invalid characters
		{name: "invalid char A", input: "A", want: 0, wantErr: true},
		{name: "invalid char in middle", input: "XIV2", want: 0, wantErr: true},
		{name: "space in middle", input: "X IV", want: 0, wantErr: true},

		// Invalid sequences - more than 3 consecutive same numerals
		{name: "IIII invalid", input: "IIII", want: 0, wantErr: true},
		{name: "XXXX invalid", input: "XXXX", want: 0, wantErr: true},
		{name: "CCCC invalid", input: "CCCC", want: 0, wantErr: true},
		{name: "MMMM invalid", input: "MMMM", want: 0, wantErr: true},

		// Invalid sequences - V, L, D cannot repeat
		{name: "VV invalid", input: "VV", want: 0, wantErr: true},
		{name: "LL invalid", input: "LL", want: 0, wantErr: true},
		{name: "DD invalid", input: "DD", want: 0, wantErr: true},

		// Invalid subtractive combinations
		{name: "IL invalid", input: "IL", want: 0, wantErr: true},
		{name: "IC invalid", input: "IC", want: 0, wantErr: true},
		{name: "ID invalid", input: "ID", want: 0, wantErr: true},
		{name: "IM invalid", input: "IM", want: 0, wantErr: true},
		{name: "VX invalid", input: "VX", want: 0, wantErr: true},
		{name: "XD invalid", input: "XD", want: 0, wantErr: true},
		{name: "XM invalid", input: "XM", want: 0, wantErr: true},
		{name: "LC invalid", input: "LC", want: 0, wantErr: true},
		{name: "LD invalid", input: "LD", want: 0, wantErr: true},
		{name: "LM invalid", input: "LM", want: 0, wantErr: true},
		{name: "DM invalid", input: "DM", want: 0, wantErr: true},

		// Invalid: subtractive pair repeated
		{name: "IXI invalid", input: "IXI", want: 0, wantErr: true},
		{name: "IXIX invalid", input: "IXIX", want: 0, wantErr: true},
		{name: "XCXC invalid", input: "XCXC", want: 0, wantErr: true},
		{name: "CMCM invalid", input: "CMCM", want: 0, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := inflect.RomanToInt(tt.input)
			if tt.wantErr {
				require.Error(t, err)
				assert.Equal(t, 0, got)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestIntToRomanRoundTrip(t *testing.T) {
	// Verify that all valid Roman numerals round-trip correctly
	for i := 1; i <= 3999; i++ {
		roman := inflect.IntToRoman(i)
		got, err := inflect.RomanToInt(roman)
		require.NoError(t, err, "Failed to parse Roman numeral for %d: %s", i, roman)
		assert.Equal(t, i, got, "Round-trip failed for %d -> %s -> %d", i, roman, got)
	}
}

func TestEngineIntToRoman(t *testing.T) {
	e := inflect.NewEngine()

	tests := []struct {
		input int
		want  string
	}{
		{1, "I"},
		{4, "IV"},
		{9, "IX"},
		{42, "XLII"},
		{1984, "MCMLXXXIV"},
		{2025, "MMXXV"},
	}

	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.input), func(t *testing.T) {
			got := e.IntToRoman(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestEngineRomanToInt(t *testing.T) {
	e := inflect.NewEngine()

	tests := []struct {
		input   string
		want    int
		wantErr bool
	}{
		{"I", 1, false},
		{"IV", 4, false},
		{"XIV", 14, false},
		{"MMXXV", 2025, false},
		{"invalid", 0, true},
		{"IIII", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := e.RomanToInt(tt.input)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

// Benchmark tests.
func BenchmarkIntToRoman(b *testing.B) {
	benchmarks := []struct {
		name  string
		input int
	}{
		{"small", 4},
		{"medium", 42},
		{"large", 1984},
		{"max", 3999},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for range b.N {
				inflect.IntToRoman(bm.input)
			}
		})
	}
}

func BenchmarkRomanToInt(b *testing.B) {
	benchmarks := []struct {
		name  string
		input string
	}{
		{"small", "IV"},
		{"medium", "XLII"},
		{"large", "MCMLXXXIV"},
		{"max", "MMMCMXCIX"},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for range b.N {
				_, _ = inflect.RomanToInt(bm.input)
			}
		})
	}
}
