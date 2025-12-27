package inflect_test

import (
	"testing"

	inflect "github.com/cv/go-inflect"
)

func TestPossessive(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		// Empty string
		{name: "empty", input: "", want: ""},

		// Regular singular nouns: add 's
		{name: "cat", input: "cat", want: "cat's"},
		{name: "dog", input: "dog", want: "dog's"},
		{name: "book", input: "book", want: "book's"},
		{name: "car", input: "car", want: "car's"},
		{name: "house", input: "house", want: "house's"},
		{name: "teacher", input: "teacher", want: "teacher's"},
		{name: "child", input: "child", want: "child's"},
		{name: "woman", input: "woman", want: "woman's"},
		{name: "man", input: "man", want: "man's"},

		// Plural nouns ending in s: add only '
		{name: "cats", input: "cats", want: "cats'"},
		{name: "dogs", input: "dogs", want: "dogs'"},
		{name: "books", input: "books", want: "books'"},
		{name: "teachers", input: "teachers", want: "teachers'"},
		{name: "boys", input: "boys", want: "boys'"},
		{name: "girls", input: "girls", want: "girls'"},
		{name: "parents", input: "parents", want: "parents'"},
		{name: "students", input: "students", want: "students'"},

		// Plural nouns not ending in s: add 's
		{name: "children", input: "children", want: "children's"},
		{name: "women", input: "women", want: "women's"},
		{name: "men", input: "men", want: "men's"},
		{name: "people", input: "people", want: "people's"},
		{name: "mice", input: "mice", want: "mice's"},
		{name: "geese", input: "geese", want: "geese's"},
		{name: "feet", input: "feet", want: "feet's"},
		{name: "teeth", input: "teeth", want: "teeth's"},
		{name: "oxen", input: "oxen", want: "oxen's"},

		// Singular nouns ending in s: add 's (default style)
		{name: "boss", input: "boss", want: "boss's"},
		{name: "class", input: "class", want: "class's"},
		{name: "bus", input: "bus", want: "bus's"},
		{name: "glass", input: "glass", want: "glass's"},
		{name: "dress", input: "dress", want: "dress's"},
		{name: "witness", input: "witness", want: "witness's"},

		// Proper names ending in s: add 's (default style)
		{name: "James", input: "James", want: "James's"},
		{name: "Charles", input: "Charles", want: "Charles's"},
		{name: "Jones", input: "Jones", want: "Jones's"},
		{name: "Williams", input: "Williams", want: "Williams's"},
		{name: "Thomas", input: "Thomas", want: "Thomas's"},
		{name: "Chris", input: "Chris", want: "Chris's"},
		{name: "Jess", input: "Jess", want: "Jess's"},
		{name: "Ross", input: "Ross", want: "Ross's"},

		// Classical/biblical names ending in s: commonly use just '
		// Note: by default we use 's, but these are commonly written with just '
		{name: "Jesus", input: "Jesus", want: "Jesus's"},
		{name: "Moses", input: "Moses", want: "Moses's"},
		{name: "Achilles", input: "Achilles", want: "Achilles's"},

		// Words ending in x, z: add 's
		{name: "fox", input: "fox", want: "fox's"},
		{name: "box", input: "box", want: "box's"},
		{name: "Max", input: "Max", want: "Max's"},
		{name: "Fritz", input: "Fritz", want: "Fritz's"},
		{name: "Sanchez", input: "Sanchez", want: "Sanchez's"},

		// Compound nouns: possessive on last word
		{name: "mother-in-law", input: "mother-in-law", want: "mother-in-law's"},
		{name: "sister-in-law", input: "sister-in-law", want: "sister-in-law's"},
		{name: "attorney general", input: "attorney general", want: "attorney general's"},
		{name: "passer-by", input: "passer-by", want: "passer-by's"},

		// Already possessive: return unchanged
		{name: "cat's already", input: "cat's", want: "cat's"},
		{name: "cats' already", input: "cats'", want: "cats'"},
		{name: "children's already", input: "children's", want: "children's"},

		// Case preservation
		{name: "CAT uppercase", input: "CAT", want: "CAT'S"},
		{name: "Cat titlecase", input: "Cat", want: "Cat's"},
		{name: "CATS uppercase", input: "CATS", want: "CATS'"},
		{name: "Cats titlecase", input: "Cats", want: "Cats'"},
		{name: "CHILDREN uppercase", input: "CHILDREN", want: "CHILDREN'S"},
		{name: "Children titlecase", input: "Children", want: "Children's"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Possessive(tt.input)
			if got != tt.want {
				t.Errorf("Possessive(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestPossessiveWithStyle(t *testing.T) {
	// Test the traditional style (just ' after s)
	tests := []struct {
		name  string
		input string
		want  string
	}{
		// Singular nouns ending in s: add only ' (traditional style)
		{name: "boss traditional", input: "boss", want: "boss'"},
		{name: "class traditional", input: "class", want: "class'"},
		{name: "James traditional", input: "James", want: "James'"},
		{name: "Charles traditional", input: "Charles", want: "Charles'"},
		{name: "Jones traditional", input: "Jones", want: "Jones'"},
		{name: "Jesus traditional", input: "Jesus", want: "Jesus'"},
		{name: "Moses traditional", input: "Moses", want: "Moses'"},

		// Non-s endings should still get 's
		{name: "cat traditional", input: "cat", want: "cat's"},
		{name: "children traditional", input: "children", want: "children's"},
	}

	// Enable traditional style
	inflect.PossessiveStyle(inflect.PossessiveTraditional)
	defer inflect.PossessiveStyle(inflect.PossessiveModern) // reset

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Possessive(tt.input)
			if got != tt.want {
				t.Errorf("Possessive(%q) with traditional style = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func BenchmarkPossessive(b *testing.B) {
	benchmarks := []struct {
		name  string
		input string
	}{
		{"singular", "cat"},
		{"plural_s", "cats"},
		{"plural_no_s", "children"},
		{"singular_s", "boss"},
		{"proper_name_s", "James"},
		{"compound", "mother-in-law"},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for range b.N {
				inflect.Possessive(bm.input)
			}
		})
	}
}
