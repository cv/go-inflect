package inflect_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cv/go-inflect"
)

func TestREADMEExamples(t *testing.T) {
	tests := []struct {
		name string
		got  string
		want string
	}{
		// Quick Start examples
		{"An(apple)", inflect.An("apple"), "an apple"},
		{"An(banana)", inflect.An("banana"), "a banana"},
		{"An(hour)", inflect.An("hour"), "an hour"},
		{"An(FBI agent)", inflect.An("FBI agent"), "an FBI agent"},

		{"Plural(cat)", inflect.Plural("cat"), "cats"},
		{"Plural(child)", inflect.Plural("child"), "children"},
		{"Plural(analysis)", inflect.Plural("analysis"), "analyses"},

		{"Singular(boxes)", inflect.Singular("boxes"), "box"},
		{"Singular(mice)", inflect.Singular("mice"), "mouse"},

		{"NumberToWords(42)", inflect.NumberToWords(42), "forty-two"},
		{"Ordinal(3)", inflect.Ordinal(3), "3rd"},
		{"OrdinalWord(3)", inflect.OrdinalWord(3), "third"},

		{"Join", inflect.Join([]string{"a", "b", "c"}), "a, b, and c"},

		{"PresentParticiple(run)", inflect.PresentParticiple("run"), "running"},
		{"PastParticiple(take)", inflect.PastParticiple("take"), "taken"},

		{"Compare(cat,cats)", inflect.Compare("cat", "cats"), "s:p"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.got)
		})
	}
}

func TestREADMEClassical(t *testing.T) {
	inflect.ClassicalAll(true)
	defer inflect.ClassicalAll(false)

	assert.Equal(t, "formulae", inflect.Plural("formula"), "Classical Plural(formula)")
	assert.Equal(t, "cacti", inflect.Plural("cactus"), "Classical Plural(cactus)")
}

func TestREADMEGender(t *testing.T) {
	defer inflect.Gender("t")

	inflect.Gender("m")
	assert.Equal(t, "he", inflect.SingularNoun("they"), "Gender m SingularNoun(they)")

	inflect.Gender("f")
	assert.Equal(t, "she", inflect.SingularNoun("they"), "Gender f SingularNoun(they)")

	inflect.Gender("t")
	assert.Equal(t, "they", inflect.SingularNoun("they"), "Gender t SingularNoun(they)")
}
