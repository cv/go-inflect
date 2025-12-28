package inflect_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	inflect "github.com/cv/go-inflect"
)

func TestIsPlural(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		// Regular plurals
		{name: "cats is plural", input: "cats", want: true},
		{name: "dogs is plural", input: "dogs", want: true},
		{name: "boxes is plural", input: "boxes", want: true},
		{name: "cities is plural", input: "cities", want: true},

		// Irregular plurals
		{name: "children is plural", input: "children", want: true},
		{name: "mice is plural", input: "mice", want: true},
		{name: "feet is plural", input: "feet", want: true},
		{name: "teeth is plural", input: "teeth", want: true},
		{name: "geese is plural", input: "geese", want: true},
		{name: "people is plural", input: "people", want: true},

		// Singular words
		{name: "cat is not plural", input: "cat", want: false},
		{name: "dog is not plural", input: "dog", want: false},
		{name: "child is not plural", input: "child", want: false},
		{name: "mouse is not plural", input: "mouse", want: false},
		{name: "foot is not plural", input: "foot", want: false},
		{name: "person is not plural", input: "person", want: false},

		// Unchanged plurals (ambiguous - default to not plural)
		{name: "sheep is ambiguous", input: "sheep", want: false},
		{name: "fish is ambiguous", input: "fish", want: false},
		{name: "deer is ambiguous", input: "deer", want: false},

		// Edge cases
		{name: "empty string", input: "", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.IsPlural(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestIsSingular(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		// Singular words
		{name: "cat is singular", input: "cat", want: true},
		{name: "dog is singular", input: "dog", want: true},
		{name: "child is singular", input: "child", want: true},
		{name: "mouse is singular", input: "mouse", want: true},
		{name: "foot is singular", input: "foot", want: true},
		{name: "person is singular", input: "person", want: true},

		// Plural words
		{name: "cats is not singular", input: "cats", want: false},
		{name: "dogs is not singular", input: "dogs", want: false},
		{name: "children is not singular", input: "children", want: false},
		{name: "mice is not singular", input: "mice", want: false},
		{name: "feet is not singular", input: "feet", want: false},
		{name: "people is not singular", input: "people", want: false},

		// Unchanged plurals (default to singular)
		{name: "sheep defaults to singular", input: "sheep", want: true},
		{name: "fish defaults to singular", input: "fish", want: true},
		{name: "deer defaults to singular", input: "deer", want: true},

		// Edge cases
		{name: "empty string", input: "", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.IsSingular(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestWordCount(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{name: "empty string", input: "", want: 0},
		{name: "single word", input: "hello", want: 1},
		{name: "two words", input: "hello world", want: 2},
		{name: "three words", input: "one two three", want: 3},
		{name: "extra spaces", input: "  one   two   three  ", want: 3},
		{name: "tabs and newlines", input: "one\ttwo\nthree", want: 3},
		{name: "just whitespace", input: "   ", want: 0},
		{name: "sentence", input: "The quick brown fox jumps", want: 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.WordCount(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCapitalize(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "empty string", input: "", want: ""},
		{name: "lowercase", input: "hello", want: "Hello"},
		{name: "uppercase", input: "HELLO", want: "HELLO"},
		{name: "already capitalized", input: "Hello", want: "Hello"},
		{name: "sentence", input: "hello world", want: "Hello world"},
		{name: "single letter", input: "a", want: "A"},
		{name: "unicode", input: "über", want: "Über"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Capitalize(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestTitleize(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "empty string", input: "", want: ""},
		{name: "lowercase", input: "hello world", want: "Hello World"},
		{name: "uppercase", input: "HELLO WORLD", want: "Hello World"},
		{name: "mixed case", input: "hELLO wORLD", want: "Hello World"},
		{name: "single word", input: "hello", want: "Hello"},
		{name: "hyphenated", input: "hello-world", want: "Hello-World"},
		{name: "extra spaces", input: "hello  world", want: "Hello  World"},
		{name: "sentence", input: "the quick brown fox", want: "The Quick Brown Fox"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Titleize(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func BenchmarkIsPlural(b *testing.B) {
	inputs := []string{"cats", "children", "cat", "sheep"}
	for i := range b.N {
		inflect.IsPlural(inputs[i%len(inputs)])
	}
}

func BenchmarkIsSingular(b *testing.B) {
	inputs := []string{"cat", "child", "cats", "sheep"}
	for i := range b.N {
		inflect.IsSingular(inputs[i%len(inputs)])
	}
}

func BenchmarkWordCount(b *testing.B) {
	input := "The quick brown fox jumps over the lazy dog"
	for range b.N {
		inflect.WordCount(input)
	}
}

func BenchmarkCapitalize(b *testing.B) {
	for range b.N {
		inflect.Capitalize("hello world")
	}
}

func BenchmarkTitleize(b *testing.B) {
	for range b.N {
		inflect.Titleize("the quick brown fox")
	}
}
