package inflect_test

import (
	"testing"

	inflect "github.com/cv/go-inflect"
)

func TestGender(t *testing.T) {
	// Reset to default before and after test
	inflect.Gender("t")
	defer inflect.Gender("t")

	tests := []struct {
		name     string
		setTo    string
		expected string
	}{
		{name: "masculine", setTo: "m", expected: "m"},
		{name: "feminine", setTo: "f", expected: "f"},
		{name: "neuter", setTo: "n", expected: "n"},
		{name: "they", setTo: "t", expected: "t"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inflect.Gender(tt.setTo)
			if got := inflect.GetGender(); got != tt.expected {
				t.Errorf("Gender(%q): GetGender() = %q, want %q", tt.setTo, got, tt.expected)
			}
		})
	}
}

func TestGetGenderDefault(t *testing.T) {
	// Reset to default before and after test
	inflect.Gender("t")
	defer inflect.Gender("t")

	// Default should be "t" (singular they)
	if got := inflect.GetGender(); got != "t" {
		t.Errorf("GetGender() default = %q, want %q", got, "t")
	}
}

func TestGenderInvalidValues(t *testing.T) {
	// Reset to default before and after test
	inflect.Gender("t")
	defer inflect.Gender("t")

	// Set a valid gender first
	inflect.Gender("m")
	if got := inflect.GetGender(); got != "m" {
		t.Errorf("Gender(m): GetGender() = %q, want %q", got, "m")
	}

	// Invalid values should be ignored
	invalidValues := []string{
		"",       // empty
		"x",      // single invalid char
		"male",   // full word
		"female", // full word
		"M",      // uppercase
		"F",      // uppercase
		"N",      // uppercase
		"T",      // uppercase
		"mm",     // repeated char
		" m",     // with space
		"m ",     // with space
	}

	for _, invalid := range invalidValues {
		t.Run("invalid:"+invalid, func(t *testing.T) {
			inflect.Gender(invalid)
			if got := inflect.GetGender(); got != "m" {
				t.Errorf("Gender(%q): GetGender() = %q, want %q (unchanged)", invalid, got, "m")
			}
		})
	}
}

func TestGenderSequence(t *testing.T) {
	// Reset to default before and after test
	inflect.Gender("t")
	defer inflect.Gender("t")

	// Test setting multiple genders in sequence
	sequence := []struct {
		setTo    string
		expected string
	}{
		{"m", "m"},
		{"f", "f"},
		{"n", "n"},
		{"t", "t"},
		{"invalid", "t"}, // should stay "t"
		{"m", "m"},
		{"", "m"}, // should stay "m"
		{"f", "f"},
	}

	for i, step := range sequence {
		inflect.Gender(step.setTo)
		if got := inflect.GetGender(); got != step.expected {
			t.Errorf("Step %d: Gender(%q): GetGender() = %q, want %q", i, step.setTo, got, step.expected)
		}
	}
}
