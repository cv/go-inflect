package inflect_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	inflect "github.com/cv/go-inflect"
)

func TestSingular(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		// Empty string
		{name: "empty", input: "", want: ""},

		// Regular plurals - remove s
		{name: "cats", input: "cats", want: "cat"},
		{name: "dogs", input: "dogs", want: "dog"},
		{name: "books", input: "books", want: "book"},

		// Words ending in -es after sibilants
		{name: "buses", input: "buses", want: "bus"},
		{name: "classes", input: "classes", want: "class"},
		{name: "bushes", input: "bushes", want: "bush"},
		{name: "churches", input: "churches", want: "church"},
		{name: "boxes", input: "boxes", want: "box"},
		{name: "buzzes", input: "buzzes", want: "buzz"},

		// -ies -> -y (consonant + ies)
		{name: "cities", input: "cities", want: "city"},
		{name: "babies", input: "babies", want: "baby"},
		{name: "flies", input: "flies", want: "fly"},

		// -ys (vowel + y) - just remove s
		{name: "boys", input: "boys", want: "boy"},
		{name: "days", input: "days", want: "day"},
		{name: "keys", input: "keys", want: "key"},

		// Words ending in -ves -> -f or -fe
		{name: "knives", input: "knives", want: "knife"},
		{name: "wives", input: "wives", want: "wife"},
		{name: "lives", input: "lives", want: "life"},
		{name: "leaves", input: "leaves", want: "leaf"},
		{name: "wolves", input: "wolves", want: "wolf"},
		{name: "calves", input: "calves", want: "calf"},
		{name: "halves", input: "halves", want: "half"},

		// Words ending in -oes -> -o
		{name: "heroes", input: "heroes", want: "hero"},
		{name: "potatoes", input: "potatoes", want: "potato"},
		{name: "tomatoes", input: "tomatoes", want: "tomato"},
		{name: "echoes", input: "echoes", want: "echo"},

		// Words ending in -os (just remove s)
		{name: "radios", input: "radios", want: "radio"},
		{name: "studios", input: "studios", want: "studio"},
		{name: "zoos", input: "zoos", want: "zoo"},
		{name: "pianos", input: "pianos", want: "piano"},
		{name: "photos", input: "photos", want: "photo"},

		// Irregular plurals
		{name: "children", input: "children", want: "child"},
		{name: "feet", input: "feet", want: "foot"},
		{name: "teeth", input: "teeth", want: "tooth"},
		{name: "mice", input: "mice", want: "mouse"},
		{name: "women", input: "women", want: "woman"},
		{name: "men", input: "men", want: "man"},
		{name: "people", input: "people", want: "person"},
		{name: "oxen", input: "oxen", want: "ox"},
		{name: "geese", input: "geese", want: "goose"},
		{name: "lice", input: "lice", want: "louse"},
		{name: "dice", input: "dice", want: "die"},

		// Latin/Greek plurals
		{name: "analyses", input: "analyses", want: "analysis"},
		{name: "crises", input: "crises", want: "crisis"},
		{name: "theses", input: "theses", want: "thesis"},
		{name: "cacti", input: "cacti", want: "cactus"},
		{name: "fungi", input: "fungi", want: "fungus"},
		{name: "nuclei", input: "nuclei", want: "nucleus"},
		{name: "bacteria", input: "bacteria", want: "bacterium"},
		{name: "data", input: "data", want: "datum"},
		{name: "media", input: "media", want: "medium"},
		{name: "appendices", input: "appendices", want: "appendix"},
		{name: "indices", input: "indices", want: "index"},
		{name: "criteria", input: "criteria", want: "criterion"},
		{name: "phenomena", input: "phenomena", want: "phenomenon"},

		// Unchanged plurals
		{name: "sheep", input: "sheep", want: "sheep"},
		{name: "deer", input: "deer", want: "deer"},
		{name: "fish", input: "fish", want: "fish"},
		{name: "species", input: "species", want: "species"},
		{name: "series", input: "series", want: "series"},
		{name: "aircraft", input: "aircraft", want: "aircraft"},
		{name: "moose", input: "moose", want: "moose"},

		// Words ending in -men -> -man
		{name: "firemen", input: "firemen", want: "fireman"},
		{name: "policemen", input: "policemen", want: "policeman"},
		{name: "spokesmen", input: "spokesmen", want: "spokesman"},

		// Nationalities ending in -ese (unchanged)
		{name: "Chinese", input: "Chinese", want: "Chinese"},
		{name: "Japanese", input: "Japanese", want: "Japanese"},
		{name: "Portuguese", input: "Portuguese", want: "Portuguese"},

		// Case preservation
		{name: "CATS uppercase", input: "CATS", want: "CAT"},
		{name: "Cats titlecase", input: "Cats", want: "Cat"},
		{name: "Children titlecase", input: "Children", want: "Child"},
		{name: "CHILDREN uppercase", input: "CHILDREN", want: "CHILD"},
		{name: "BOXES uppercase", input: "BOXES", want: "BOX"},
		{name: "Cities titlecase", input: "Cities", want: "City"},
		{name: "MICE uppercase", input: "MICE", want: "MOUSE"},

		// Already singular (should return unchanged)
		{name: "already singular cat", input: "cat", want: "cat"},
		{name: "already singular class", input: "class", want: "class"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Singular(tt.input)
			assert.Equal(t, tt.want, got, "Singular(%q)", tt.input)
		})
	}
}

func BenchmarkSingular(b *testing.B) {
	// Test with representative inputs covering different singularization rules
	benchmarks := []struct {
		name  string
		input string
	}{
		{"regular", "cats"},       // Simple: remove -s
		{"sibilant", "boxes"},     // Remove -es
		{"ies_to_y", "cities"},    // Change ies to y
		{"irregular", "children"}, // Irregular: child
		{"latin", "analyses"},     // Latin: -es to -is
		{"unchanged", "sheep"},    // Unchanged
		{"ves_to_f", "knives"},    // ves to f/fe
		{"men_to_man", "firemen"}, // -men to -man
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for range b.N {
				inflect.Singular(bm.input)
			}
		})
	}
}
