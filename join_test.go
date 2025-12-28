package inflect_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

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
			assert.Equal(t, tt.want, got)
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
			assert.Equal(t, tt.want, got)
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
			assert.Equal(t, tt.want, got)
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
			assert.Equal(t, tt.want, got)
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
			for range b.N {
				inflect.Join(bm.input)
			}
		})
	}
}

func TestJoinWithFinalSep(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		conj     string
		sep      string
		finalSep string
		want     string
	}{
		// Empty and single
		{name: "empty slice", input: []string{}, conj: "and", sep: ", ", finalSep: "; ", want: ""},
		{name: "single word", input: []string{"a"}, conj: "and", sep: ", ", finalSep: "; ", want: "a"},

		// Two words (no separator used)
		{name: "two words", input: []string{"a", "b"}, conj: "and", sep: ", ", finalSep: "; ", want: "a and b"},

		// Three+ words with different final separator
		{name: "three words comma then semicolon", input: []string{"a", "b", "c"}, conj: "and", sep: ", ", finalSep: "; ", want: "a, b; and c"},
		{name: "four words comma then semicolon", input: []string{"a", "b", "c", "d"}, conj: "and", sep: ", ", finalSep: "; ", want: "a, b, c; and d"},

		// Same separator and final separator (should match JoinWithSep)
		{name: "same separators", input: []string{"a", "b", "c"}, conj: "and", sep: ", ", finalSep: ", ", want: "a, b, and c"},

		// No final separator (space only for clean look)
		{name: "space final separator", input: []string{"a", "b", "c"}, conj: "and", sep: ", ", finalSep: " ", want: "a, b and c"},

		// Different conjunctions
		{name: "or conjunction", input: []string{"a", "b", "c"}, conj: "or", sep: ", ", finalSep: "; ", want: "a, b; or c"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.JoinWithFinalSep(tt.input, tt.conj, tt.sep, tt.finalSep)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestJoinNoOxford(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  string
	}{
		// Empty and single
		{name: "empty slice", input: []string{}, want: ""},
		{name: "single word", input: []string{"a"}, want: "a"},

		// Two words (same as Join)
		{name: "two words", input: []string{"a", "b"}, want: "a and b"},

		// Three+ words without Oxford comma
		{name: "three words", input: []string{"a", "b", "c"}, want: "a, b and c"},
		{name: "four words", input: []string{"a", "b", "c", "d"}, want: "a, b, c and d"},
		{name: "five words", input: []string{"one", "two", "three", "four", "five"}, want: "one, two, three, four and five"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.JoinNoOxford(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestJoinNoOxfordWithConj(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		conj  string
		want  string
	}{
		// Empty and single
		{name: "empty slice", input: []string{}, conj: "or", want: ""},
		{name: "single word", input: []string{"a"}, conj: "or", want: "a"},

		// Two words
		{name: "two words", input: []string{"a", "b"}, conj: "or", want: "a or b"},

		// Three+ words without Oxford comma
		{name: "three words or", input: []string{"a", "b", "c"}, conj: "or", want: "a, b or c"},
		{name: "three words and", input: []string{"a", "b", "c"}, conj: "and", want: "a, b and c"},
		{name: "four words nor", input: []string{"a", "b", "c", "d"}, conj: "nor", want: "a, b, c nor d"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.JoinNoOxfordWithConj(tt.input, tt.conj)
			assert.Equal(t, tt.want, got)
		})
	}
}

func BenchmarkJoinWithFinalSep(b *testing.B) {
	input := []string{"a", "b", "c", "d", "e"}
	for b.Loop() {
		inflect.JoinWithFinalSep(input, "and", ", ", "; ")
	}
}

func BenchmarkJoinNoOxford(b *testing.B) {
	input := []string{"a", "b", "c", "d", "e"}
	for b.Loop() {
		inflect.JoinNoOxford(input)
	}
}
