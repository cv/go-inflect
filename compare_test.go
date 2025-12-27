package inflect_test

import (
	"testing"

	inflect "github.com/cv/go-inflect"
)

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
			for range b.N {
				inflect.Compare(bm.word1, bm.word2)
			}
		})
	}
}
