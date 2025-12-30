package inflect_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	inflect "github.com/cv/go-inflect/v2"
)

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

		// Words ending in -man that should NOT become -men
		{name: "German", input: "German", want: "Germans"},
		{name: "Roman", input: "Roman", want: "Romans"},
		{name: "Ottoman", input: "Ottoman", want: "Ottomans"},
		{name: "Norman", input: "Norman", want: "Normans"},
		{name: "shaman", input: "shaman", want: "shamans"},
		{name: "talisman", input: "talisman", want: "talismans"},
		{name: "human", input: "human", want: "humans"},

		// Additional Latin neuter (-um -> -a)
		{name: "addendum", input: "addendum", want: "addenda"},
		{name: "erratum", input: "erratum", want: "errata"},
		{name: "symposium", input: "symposium", want: "symposia"},
		{name: "atrium", input: "atrium", want: "atria"},

		// Greek neuter (-on -> -a)
		{name: "automaton", input: "automaton", want: "automata"},
		{name: "polyhedron", input: "polyhedron", want: "polyhedra"},

		// Hebrew plurals
		{name: "seraph", input: "seraph", want: "seraphim"},
		{name: "cherub", input: "cherub", want: "cherubim"},
		{name: "kibbutz", input: "kibbutz", want: "kibbutzim"},

		// Italian plurals
		{name: "graffito", input: "graffito", want: "graffiti"},
		{name: "virtuoso", input: "virtuoso", want: "virtuosi"},
		{name: "libretto", input: "libretto", want: "libretti"},

		// Words ending in -f that now use -ves
		{name: "hoof", input: "hoof", want: "hooves"},
		{name: "scarf", input: "scarf", want: "scarves"},
		{name: "wharf", input: "wharf", want: "wharves"},

		// Additional -is -> -es words
		{name: "axis", input: "axis", want: "axes"},
		{name: "ellipsis", input: "ellipsis", want: "ellipses"},
		{name: "nemesis", input: "nemesis", want: "nemeses"},
		{name: "synthesis", input: "synthesis", want: "syntheses"},

		// Additional -us -> -i words
		{name: "calculus", input: "calculus", want: "calculi"},
		{name: "locus", input: "locus", want: "loci"},
		{name: "bacillus", input: "bacillus", want: "bacilli"},

		// Additional -ex/-ix -> -ices words
		{name: "cortex", input: "cortex", want: "cortices"},
		{name: "vortex", input: "vortex", want: "vortices"},
		{name: "helix", input: "helix", want: "helices"},

		// French -eau -> -eaux
		{name: "bureau", input: "bureau", want: "bureaux"},
		{name: "plateau", input: "plateau", want: "plateaux"},
		{name: "chateau", input: "chateau", want: "chateaux"},

		// Unchanged plurals (French loanwords, etc.)
		{name: "corps", input: "corps", want: "corps"},
		{name: "chassis", input: "chassis", want: "chassis"},
		{name: "means", input: "means", want: "means"},
		{name: "gallows", input: "gallows", want: "gallows"},
		{name: "barracks", input: "barracks", want: "barracks"},
		{name: "headquarters", input: "headquarters", want: "headquarters"},

		// Compound -foot -> -feet
		{name: "bigfoot", input: "bigfoot", want: "bigfeet"},
		{name: "clubfoot", input: "clubfoot", want: "clubfeet"},

		// Compound -tooth -> -teeth
		{name: "eyetooth", input: "eyetooth", want: "eyeteeth"},
		{name: "sabertooth", input: "sabertooth", want: "saberteeth"},

		// Brand names and other -man exceptions
		{name: "Walkman", input: "Walkman", want: "Walkmans"},
		{name: "leman", input: "leman", want: "lemans"},

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
			assert.Equal(t, tt.want, got)
		})
	}
}

func BenchmarkPlural(b *testing.B) {
	// Test with representative inputs covering different pluralization rules
	benchmarks := []struct {
		name  string
		input string
	}{
		{"regular", "cat"},        // Simple: add -s
		{"sibilant", "box"},       // Add -es (ends in x)
		{"consonant_y", "city"},   // Change y to ies
		{"irregular", "child"},    // Irregular: children
		{"latin", "analysis"},     // Latin: -is to -es
		{"unchanged", "sheep"},    // Unchanged plural
		{"f_to_ves", "knife"},     // f/fe to ves
		{"man_to_men", "fireman"}, // -man to -men
		{"o_adds_es", "hero"},     // consonant + o adds -es
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for range b.N {
				inflect.Plural(bm.input)
			}
		})
	}
}

func BenchmarkPluralNoun(b *testing.B) {
	benchmarks := []struct {
		name  string
		word  string
		count []int
	}{
		{"regular_no_count", "cat", nil},
		{"regular_count_2", "cat", []int{2}},
		{"regular_count_1", "cat", []int{1}},
		{"pronoun_I", "I", nil},
		{"pronoun_they", "they", nil},
		{"irregular", "child", nil},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for range b.N {
				inflect.PluralNoun(bm.word, bm.count...)
			}
		})
	}
}

func BenchmarkPluralVerb(b *testing.B) {
	benchmarks := []struct {
		name  string
		word  string
		count []int
	}{
		{"is_to_are", "is", nil},
		{"has_to_have", "has", nil},
		{"runs_to_run", "runs", nil},
		{"walks_to_walk", "walks", nil},
		{"with_count_1", "is", []int{1}},
		{"with_count_2", "is", []int{2}},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for range b.N {
				inflect.PluralVerb(bm.word, bm.count...)
			}
		})
	}
}
