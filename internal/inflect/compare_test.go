package inflect_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	inflect "github.com/cv/go-inflect/v2"
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
			assert.Equal(t, tt.want, got)
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
			assert.Equal(t, tt.want, got)
			// Verify it matches Compare() behavior
			compareGot := inflect.Compare(tt.noun1, tt.noun2)
			assert.Equal(t, compareGot, got)
		})
	}
}

func TestCompareVerbs(t *testing.T) {
	tests := []struct {
		name  string
		verb1 string
		verb2 string
		want  string
	}{
		// Empty strings
		{name: "empty strings", verb1: "", verb2: "", want: "eq"},
		{name: "empty and word", verb1: "", verb2: "run", want: ""},
		{name: "word and empty", verb1: "run", verb2: "", want: ""},

		// Equal verbs
		{name: "same verb", verb1: "run", verb2: "run", want: "eq"},
		{name: "same verb uppercase", verb1: "RUN", verb2: "run", want: "eq"},

		// 3rd person singular to base form (s:p)
		{name: "runs to run", verb1: "runs", verb2: "run", want: "s:p"},
		{name: "walks to walk", verb1: "walks", verb2: "walk", want: "s:p"},
		{name: "watches to watch", verb1: "watches", verb2: "watch", want: "s:p"},
		{name: "tries to try", verb1: "tries", verb2: "try", want: "s:p"},
		{name: "goes to go", verb1: "goes", verb2: "go", want: "s:p"},
		{name: "does to do", verb1: "does", verb2: "do", want: "s:p"},
		{name: "has to have", verb1: "has", verb2: "have", want: "s:p"},
		{name: "is to are", verb1: "is", verb2: "are", want: "s:p"},
		{name: "was to were", verb1: "was", verb2: "were", want: "s:p"},

		// Base form to 3rd person singular (p:s)
		{name: "run to runs", verb1: "run", verb2: "runs", want: "p:s"},
		{name: "walk to walks", verb1: "walk", verb2: "walks", want: "p:s"},
		{name: "watch to watches", verb1: "watch", verb2: "watches", want: "p:s"},
		{name: "try to tries", verb1: "try", verb2: "tries", want: "p:s"},
		{name: "go to goes", verb1: "go", verb2: "goes", want: "p:s"},
		{name: "do to does", verb1: "do", verb2: "does", want: "p:s"},
		{name: "have to has", verb1: "have", verb2: "has", want: "p:s"},
		{name: "are to is", verb1: "are", verb2: "is", want: "p:s"},
		{name: "were to was", verb1: "were", verb2: "was", want: "p:s"},

		// Contractions
		{name: "isn't to aren't", verb1: "isn't", verb2: "aren't", want: "s:p"},
		{name: "doesn't to don't", verb1: "doesn't", verb2: "don't", want: "s:p"},
		{name: "hasn't to haven't", verb1: "hasn't", verb2: "haven't", want: "s:p"},

		// Unrelated verbs
		{name: "different verbs", verb1: "run", verb2: "walk", want: ""},
		{name: "runs to walking", verb1: "runs", verb2: "walking", want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.CompareVerbs(tt.verb1, tt.verb2)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCompareAdjs(t *testing.T) {
	tests := []struct {
		name string
		adj1 string
		adj2 string
		want string
	}{
		// Empty strings
		{name: "empty strings", adj1: "", adj2: "", want: "eq"},
		{name: "empty and word", adj1: "", adj2: "this", want: ""},
		{name: "word and empty", adj1: "this", adj2: "", want: ""},

		// Equal adjectives
		{name: "same adjective", adj1: "this", adj2: "this", want: "eq"},
		{name: "same adjective uppercase", adj1: "THIS", adj2: "this", want: "eq"},

		// Demonstrative adjectives (s:p)
		{name: "this to these", adj1: "this", adj2: "these", want: "s:p"},
		{name: "that to those", adj1: "that", adj2: "those", want: "s:p"},

		// Demonstrative adjectives (p:s)
		{name: "these to this", adj1: "these", adj2: "this", want: "p:s"},
		{name: "those to that", adj1: "those", adj2: "that", want: "p:s"},

		// Articles
		{name: "a to some", adj1: "a", adj2: "some", want: "s:p"},
		{name: "an to some", adj1: "an", adj2: "some", want: "s:p"},
		{name: "some to a", adj1: "some", adj2: "a", want: "p:s"},

		// Possessive adjectives
		{name: "my to our", adj1: "my", adj2: "our", want: "s:p"},
		{name: "his to their", adj1: "his", adj2: "their", want: "s:p"},
		{name: "her to their", adj1: "her", adj2: "their", want: "s:p"},
		{name: "our to my", adj1: "our", adj2: "my", want: "p:s"},

		// Unrelated adjectives
		{name: "different adjectives", adj1: "big", adj2: "small", want: ""},
		{name: "this to that", adj1: "this", adj2: "that", want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.CompareAdjs(tt.adj1, tt.adj2)
			assert.Equal(t, tt.want, got)
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

func BenchmarkCompareNouns(b *testing.B) {
	benchmarks := []struct {
		name  string
		word1 string
		word2 string
	}{
		{"equal", "cat", "cat"},
		{"singular_plural", "cat", "cats"},
		{"plural_singular", "cats", "cat"},
		{"irregular", "child", "children"},
		{"unrelated", "cat", "dog"},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for range b.N {
				inflect.CompareNouns(bm.word1, bm.word2)
			}
		})
	}
}

func BenchmarkCompareVerbs(b *testing.B) {
	benchmarks := []struct {
		name  string
		word1 string
		word2 string
	}{
		{"equal", "run", "run"},
		{"singular_plural", "runs", "run"},
		{"is_are", "is", "are"},
		{"has_have", "has", "have"},
		{"unrelated", "run", "walk"},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for range b.N {
				inflect.CompareVerbs(bm.word1, bm.word2)
			}
		})
	}
}

func BenchmarkCompareAdjs(b *testing.B) {
	benchmarks := []struct {
		name  string
		word1 string
		word2 string
	}{
		{"equal", "big", "big"},
		{"this_these", "this", "these"},
		{"that_those", "that", "those"},
		{"a_some", "a", "some"},
		{"unrelated", "big", "small"},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for range b.N {
				inflect.CompareAdjs(bm.word1, bm.word2)
			}
		})
	}
}
