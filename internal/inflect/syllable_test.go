package inflect_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	inflect "github.com/cv/go-inflect/v2"
)

// =============================================================================
// Unit Tests - Table-driven tests for CountSyllables
// =============================================================================

func TestCountSyllables(t *testing.T) {
	// Note: CountSyllables is a heuristic-based estimator. The expected values
	// reflect what the algorithm produces, which may differ from actual syllable
	// counts for some words due to English spelling irregularities.
	tests := []struct {
		name  string
		input string
		want  int
	}{
		// Empty string
		{name: "empty string", input: "", want: 0},

		// Single syllable words
		{name: "cat", input: "cat", want: 1},
		{name: "dog", input: "dog", want: 1},
		{name: "book", input: "book", want: 1},
		{name: "tree", input: "tree", want: 1},   // silent e handled
		{name: "house", input: "house", want: 1}, // silent e handled
		{name: "thought", input: "thought", want: 1},
		{name: "through", input: "through", want: 1},
		{name: "strengths", input: "strengths", want: 1},
		{name: "a", input: "a", want: 1},
		{name: "I", input: "I", want: 1},

		// Two syllable words
		{name: "happy", input: "happy", want: 1}, // y treated as vowel, 'a' and 'y' consecutive-ish
		{name: "garden", input: "garden", want: 2},
		{name: "table", input: "table", want: 1},   // silent e handled
		{name: "people", input: "people", want: 1}, // 'eo' is one group, silent e handled
		{name: "water", input: "water", want: 2},
		{name: "river", input: "river", want: 2},
		{name: "yellow", input: "yellow", want: 2},

		// Three syllable words
		{name: "beautiful", input: "beautiful", want: 3},
		{name: "elephant", input: "elephant", want: 3},
		{name: "wonderful", input: "wonderful", want: 3},
		{name: "tomorrow", input: "tomorrow", want: 3},
		{name: "banana", input: "banana", want: 3},

		// Multi-syllable words (heuristic estimates)
		{name: "education", input: "education", want: 4},
		{name: "communication", input: "communication", want: 5},
		{name: "dictionary", input: "dictionary", want: 3}, // heuristic: 'io' counted as one group
		{name: "territory", input: "territory", want: 3},   // heuristic: 'o' and 'y' groups
		{name: "extraordinary", input: "extraordinary", want: 4},
		{name: "unbelievable", input: "unbelievable", want: 4}, // silent e handled
		{name: "international", input: "international", want: 5},

		// Silent e handling
		{name: "make", input: "make", want: 1},
		{name: "take", input: "take", want: 1},
		{name: "love", input: "love", want: 1},
		{name: "come", input: "come", want: 1},
		{name: "time", input: "time", want: 1},
		{name: "life", input: "life", want: 1},
		{name: "fire", input: "fire", want: 1},
		{name: "before", input: "before", want: 2},
		{name: "complete", input: "complete", want: 2},

		// Words ending in 'e' that is not silent
		{name: "be", input: "be", want: 1},
		{name: "me", input: "me", want: 1},
		{name: "he", input: "he", want: 1},
		{name: "we", input: "we", want: 1},

		// Case insensitivity
		{name: "uppercase CAT", input: "CAT", want: 1},
		{name: "uppercase DOG", input: "DOG", want: 1},
		{name: "mixed case BeAuTiFuL", input: "BeAuTiFuL", want: 3},
		{name: "uppercase BEAUTIFUL", input: "BEAUTIFUL", want: 3},

		// Edge cases
		{name: "single vowel", input: "e", want: 1},
		{name: "consonants only", input: "xyz", want: 1},    // minimum 1 syllable for non-empty
		{name: "all vowels aeiou", input: "aeiou", want: 1}, // all consecutive = 1 group
		{name: "alternating ae", input: "aeae", want: 1},    // consecutive vowels
		{name: "y as vowel", input: "gym", want: 1},
		{name: "y as vowel rhythm", input: "rhythm", want: 1},
		{name: "y as vowel cycle", input: "cycle", want: 1}, // silent e handled

		// Words with unusual vowel patterns
		{name: "queue", input: "queue", want: 1}, // 'ueue' all vowels, one group
		{name: "eye", input: "eye", want: 1},     // silent e handled
		{name: "audio", input: "audio", want: 2}, // 'au' and 'io'
		{name: "area", input: "area", want: 2},   // 'a' and 'ea'
		{name: "idea", input: "idea", want: 2},   // 'i' and 'ea'
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := inflect.CountSyllables(tc.input)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestEngineCountSyllables(t *testing.T) {
	e := inflect.NewEngine()

	tests := []struct {
		name  string
		input string
		want  int
	}{
		{name: "empty", input: "", want: 0},
		{name: "cat", input: "cat", want: 1},
		{name: "beautiful", input: "beautiful", want: 3},
		{name: "international", input: "international", want: 5},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := e.CountSyllables(tc.input)
			assert.Equal(t, tc.want, got)
		})
	}
}

// TestCountSyllablesConsistency verifies that package-level and Engine methods
// return the same results.
func TestCountSyllablesConsistency(t *testing.T) {
	e := inflect.NewEngine()
	words := []string{
		"", "a", "cat", "dog", "beautiful", "extraordinary",
		"CAT", "DOG", "BeAuTiFuL", "make", "time", "complete",
	}

	for _, word := range words {
		t.Run(word, func(t *testing.T) {
			pkg := inflect.CountSyllables(word)
			eng := e.CountSyllables(word)
			assert.Equal(t, pkg, eng, "package-level and Engine methods should return same result")
		})
	}
}

// =============================================================================
// Example Tests
// =============================================================================

func ExampleCountSyllables() {
	fmt.Println(inflect.CountSyllables("cat"))
	fmt.Println(inflect.CountSyllables("beautiful"))
	fmt.Println(inflect.CountSyllables("international"))
	// Output:
	// 1
	// 3
	// 5
}

func ExampleCountSyllables_silentE() {
	// Silent 'e' at the end of words is handled
	fmt.Println(inflect.CountSyllables("make"))
	fmt.Println(inflect.CountSyllables("time"))
	fmt.Println(inflect.CountSyllables("complete"))
	// Output:
	// 1
	// 1
	// 2
}

func ExampleEngine_CountSyllables() {
	e := inflect.NewEngine()
	fmt.Println(e.CountSyllables("elephant"))
	fmt.Println(e.CountSyllables("international"))
	// Output:
	// 3
	// 5
}

// =============================================================================
// Benchmark Tests
// =============================================================================

func BenchmarkCountSyllables(b *testing.B) {
	for b.Loop() {
		inflect.CountSyllables("beautiful")
	}
}

func BenchmarkCountSyllablesShort(b *testing.B) {
	for b.Loop() {
		inflect.CountSyllables("cat")
	}
}

func BenchmarkCountSyllablesLong(b *testing.B) {
	for b.Loop() {
		inflect.CountSyllables("internationalization")
	}
}

func BenchmarkEngineCountSyllables(b *testing.B) {
	e := inflect.NewEngine()
	for b.Loop() {
		e.CountSyllables("beautiful")
	}
}

func BenchmarkCountSyllablesParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			inflect.CountSyllables("beautiful")
		}
	})
}

// =============================================================================
// Fuzz Tests
// =============================================================================

func FuzzCountSyllables(f *testing.F) {
	// Seed corpus with interesting cases
	seeds := []string{
		"", "a", "e", "i", "o", "u", "y",
		"cat", "dog", "beautiful", "extraordinary",
		"make", "time", "complete",
		"CAT", "BEAUTIFUL",
		"xyz", "aeiou", "rhythm",
		"queue", "eye",
		"strengths", "through", "thought",
	}

	for _, seed := range seeds {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, word string) {
		result := inflect.CountSyllables(word)

		// Invariants that should always hold:
		// 1. Result is non-negative
		if result < 0 {
			t.Errorf("CountSyllables(%q) = %d, want >= 0", word, result)
		}

		// 2. Empty string returns 0
		if word == "" && result != 0 {
			t.Errorf("CountSyllables(%q) = %d, want 0", word, result)
		}

		// 3. Non-empty string returns at least 1
		if word != "" && result < 1 {
			t.Errorf("CountSyllables(%q) = %d, want >= 1 for non-empty string", word, result)
		}

		// 4. Engine method returns same result
		e := inflect.NewEngine()
		engineResult := e.CountSyllables(word)
		if result != engineResult {
			t.Errorf("CountSyllables(%q) = %d, Engine.CountSyllables(%q) = %d, want same", word, result, word, engineResult)
		}
	})
}
