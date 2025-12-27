package inflect_test

import (
	"testing"

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
			if tt.got != tt.want {
				t.Errorf("got %q, want %q", tt.got, tt.want)
			}
		})
	}
}

func TestREADMEClassical(t *testing.T) {
	inflect.ClassicalAll(true)
	defer inflect.ClassicalAll(false)

	if got := inflect.Plural("formula"); got != "formulae" {
		t.Errorf("Classical Plural(formula) = %q, want %q", got, "formulae")
	}
	if got := inflect.Plural("cactus"); got != "cacti" {
		t.Errorf("Classical Plural(cactus) = %q, want %q", got, "cacti")
	}
}

func TestREADMEGender(t *testing.T) {
	defer inflect.Gender("t")

	inflect.Gender("m")
	if got := inflect.SingularNoun("they"); got != "he" {
		t.Errorf("Gender m SingularNoun(they) = %q, want %q", got, "he")
	}

	inflect.Gender("f")
	if got := inflect.SingularNoun("they"); got != "she" {
		t.Errorf("Gender f SingularNoun(they) = %q, want %q", got, "she")
	}

	inflect.Gender("t")
	if got := inflect.SingularNoun("they"); got != "they" {
		t.Errorf("Gender t SingularNoun(they) = %q, want %q", got, "they")
	}
}

func TestREADMEInflect(t *testing.T) {
	if got := inflect.Inflect("I saw plural('cat', 3)"); got != "I saw cats" {
		t.Errorf("Inflect plural = %q, want %q", got, "I saw cats")
	}
	if got := inflect.Inflect("This is the ordinal(1) item"); got != "This is the 1st item" {
		t.Errorf("Inflect ordinal = %q, want %q", got, "This is the 1st item")
	}
}
