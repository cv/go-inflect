package inflect_test

import (
	"testing"

	"gitlab-master.nvidia.com/urg/go-inflect"
)

func TestAn(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		// Basic cases
		{name: "consonant start", input: "cat", want: "a cat"},
		{name: "vowel start", input: "ant", want: "an ant"},

		// Single letters
		{name: "vowel letter", input: "a", want: "an a"},
		{name: "consonant letter", input: "b", want: "a b"},

		// Silent H
		{name: "silent h", input: "honest cat", want: "an honest cat"},
		{name: "regular h", input: "dishonest cat", want: "a dishonest cat"},
		{name: "h proper noun", input: "Honolulu sunset", want: "a Honolulu sunset"},

		// Special pronunciation cases
		{name: "mpeg abbreviation", input: "mpeg", want: "an mpeg"},
		{name: "onetime exception", input: "onetime holiday", want: "a onetime holiday"},

		// Vowels with consonant sounds (U variations)
		{name: "Ugandan", input: "Ugandan person", want: "a Ugandan person"},
		{name: "Ukrainian", input: "Ukrainian person", want: "a Ukrainian person"},
		{name: "Unabomber", input: "Unabomber", want: "a Unabomber"},
		{name: "unanimous", input: "unanimous decision", want: "a unanimous decision"},

		// Abbreviations and acronyms
		{name: "US abbreviation", input: "US farmer", want: "a US farmer"},
		{name: "uppercase word", input: "wild PIKACHU appeared", want: "a wild PIKACHU appeared"},
		{name: "YAML acronym", input: "YAML code block", want: "a YAML code block"},
		{name: "Core ML", input: "Core ML function", want: "a Core ML function"},
		{name: "JSON acronym", input: "JSON code block", want: "a JSON code block"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.An(tt.input)
			if got != tt.want {
				t.Errorf("An(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

// TestA verifies that A() is an alias for An()
func TestA(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"cat", "a cat"},
		{"ant", "an ant"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := inflect.A(tt.input)
			if got != tt.want {
				t.Errorf("A(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestPlural(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		// Empty string
		{name: "empty", input: "", want: ""},

		// Regular plurals - add s
		{name: "cat", input: "cat", want: "cats"},
		{name: "dog", input: "dog", want: "dogs"},
		{name: "book", input: "book", want: "books"},

		// Words ending in s, ss, sh, ch, x, z - add es
		{name: "bus", input: "bus", want: "buses"},
		{name: "class", input: "class", want: "classes"},
		{name: "bush", input: "bush", want: "bushes"},
		{name: "church", input: "church", want: "churches"},
		{name: "box", input: "box", want: "boxes"},
		{name: "buzz", input: "buzz", want: "buzzes"},

		// Consonant + y -> ies
		{name: "city", input: "city", want: "cities"},
		{name: "baby", input: "baby", want: "babies"},
		{name: "fly", input: "fly", want: "flies"},

		// Vowel + y -> ys
		{name: "boy", input: "boy", want: "boys"},
		{name: "day", input: "day", want: "days"},
		{name: "key", input: "key", want: "keys"},

		// Words ending in f/fe -> ves
		{name: "knife", input: "knife", want: "knives"},
		{name: "wife", input: "wife", want: "wives"},
		{name: "leaf", input: "leaf", want: "leaves"},
		{name: "wolf", input: "wolf", want: "wolves"},

		// Words ending in f that just take s
		{name: "roof", input: "roof", want: "roofs"},
		{name: "chief", input: "chief", want: "chiefs"},

		// Words ending in o
		{name: "hero", input: "hero", want: "heroes"},
		{name: "potato", input: "potato", want: "potatoes"},
		{name: "tomato", input: "tomato", want: "tomatoes"},
		{name: "echo", input: "echo", want: "echoes"},

		// Words ending in o that take s (vowel + o, exceptions)
		{name: "radio", input: "radio", want: "radios"},
		{name: "studio", input: "studio", want: "studios"},
		{name: "zoo", input: "zoo", want: "zoos"},
		{name: "piano", input: "piano", want: "pianos"},
		{name: "photo", input: "photo", want: "photos"},

		// Irregular plurals
		{name: "child", input: "child", want: "children"},
		{name: "foot", input: "foot", want: "feet"},
		{name: "tooth", input: "tooth", want: "teeth"},
		{name: "mouse", input: "mouse", want: "mice"},
		{name: "woman", input: "woman", want: "women"},
		{name: "man", input: "man", want: "men"},
		{name: "person", input: "person", want: "people"},
		{name: "ox", input: "ox", want: "oxen"},

		// Latin/Greek plurals
		{name: "analysis", input: "analysis", want: "analyses"},
		{name: "crisis", input: "crisis", want: "crises"},
		{name: "thesis", input: "thesis", want: "theses"},
		{name: "cactus", input: "cactus", want: "cacti"},
		{name: "fungus", input: "fungus", want: "fungi"},
		{name: "nucleus", input: "nucleus", want: "nuclei"},
		{name: "bacterium", input: "bacterium", want: "bacteria"},
		{name: "datum", input: "datum", want: "data"},
		{name: "medium", input: "medium", want: "media"},
		{name: "appendix", input: "appendix", want: "appendices"},
		{name: "index", input: "index", want: "indices"},

		// Unchanged plurals
		{name: "sheep", input: "sheep", want: "sheep"},
		{name: "deer", input: "deer", want: "deer"},
		{name: "fish", input: "fish", want: "fish"},
		{name: "species", input: "species", want: "species"},
		{name: "series", input: "series", want: "series"},
		{name: "aircraft", input: "aircraft", want: "aircraft"},

		// Words ending in -man -> -men
		{name: "fireman", input: "fireman", want: "firemen"},
		{name: "policeman", input: "policeman", want: "policemen"},
		{name: "spokesman", input: "spokesman", want: "spokesmen"},

		// Nationalities ending in -ese (unchanged)
		{name: "Chinese", input: "Chinese", want: "Chinese"},
		{name: "Japanese", input: "Japanese", want: "Japanese"},
		{name: "Portuguese", input: "Portuguese", want: "Portuguese"},

		// Case preservation
		{name: "CAT uppercase", input: "CAT", want: "CATS"},
		{name: "Cat titlecase", input: "Cat", want: "Cats"},
		{name: "Child titlecase", input: "Child", want: "Children"},
		{name: "CHILD uppercase", input: "CHILD", want: "CHILDREN"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Plural(tt.input)
			if got != tt.want {
				t.Errorf("Plural(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
