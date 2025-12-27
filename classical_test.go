package inflect_test

import (
	"testing"

	inflect "github.com/cv/go-inflect"
)

func TestClassical(t *testing.T) {
	// Clean up after test
	defer inflect.Classical(false)

	tests := []struct {
		name    string
		enabled bool
	}{
		{name: "enable classical mode", enabled: true},
		{name: "disable classical mode", enabled: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inflect.Classical(tt.enabled)
			got := inflect.IsClassical()
			if got != tt.enabled {
				t.Errorf("after Classical(%v): IsClassical() = %v, want %v", tt.enabled, got, tt.enabled)
			}
		})
	}
}

func TestIsClassical(t *testing.T) {
	// Clean up after test
	defer inflect.Classical(false)

	tests := []struct {
		name  string
		setup func()
		want  bool
	}{
		{
			name:  "default state is false",
			setup: func() { inflect.Classical(false) },
			want:  false,
		},
		{
			name:  "after enabling",
			setup: func() { inflect.Classical(true) },
			want:  true,
		},
		{
			name:  "after enabling then disabling",
			setup: func() { inflect.Classical(true); inflect.Classical(false) },
			want:  false,
		},
		{
			name: "after multiple toggles ending enabled",
			setup: func() {
				inflect.Classical(false)
				inflect.Classical(true)
				inflect.Classical(false)
				inflect.Classical(true)
			},
			want: true,
		},
		{
			name: "after multiple toggles ending disabled",
			setup: func() {
				inflect.Classical(true)
				inflect.Classical(false)
				inflect.Classical(true)
				inflect.Classical(false)
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			got := inflect.IsClassical()
			if got != tt.want {
				t.Errorf("IsClassical() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClassicalIntegration(t *testing.T) {
	// Clean up after test
	defer inflect.Classical(false)

	t.Run("complete workflow", func(t *testing.T) {
		// 1. Initial/default state should be false
		inflect.Classical(false)
		if got := inflect.IsClassical(); got != false {
			t.Errorf("Initial IsClassical() = %v, want false", got)
		}

		// 2. Enable classical mode
		inflect.Classical(true)
		if got := inflect.IsClassical(); got != true {
			t.Errorf("After Classical(true): IsClassical() = %v, want true", got)
		}

		// 3. Disable classical mode
		inflect.Classical(false)
		if got := inflect.IsClassical(); got != false {
			t.Errorf("After Classical(false): IsClassical() = %v, want false", got)
		}

		// 4. Toggle multiple times
		inflect.Classical(true)
		inflect.Classical(true) // Setting to same value should work
		if got := inflect.IsClassical(); got != true {
			t.Errorf("After double Classical(true): IsClassical() = %v, want true", got)
		}

		inflect.Classical(false)
		inflect.Classical(false) // Setting to same value should work
		if got := inflect.IsClassical(); got != false {
			t.Errorf("After double Classical(false): IsClassical() = %v, want false", got)
		}
	})
}

func TestClassicalAncient(t *testing.T) {
	// Clean up after test
	defer inflect.ClassicalAncient(false)

	tests := []struct {
		name    string
		enabled bool
		input   string
		want    string
	}{
		{
			name:    "enabled formula becomes formulae",
			enabled: true,
			input:   "formula",
			want:    "formulae",
		},
		{
			name:    "disabled formula becomes formulas",
			enabled: false,
			input:   "formula",
			want:    "formulas",
		},
		{
			name:    "enabled antenna becomes antennae",
			enabled: true,
			input:   "antenna",
			want:    "antennae",
		},
		{
			name:    "disabled antenna becomes antennas",
			enabled: false,
			input:   "antenna",
			want:    "antennas",
		},
		{
			name:    "enabled nebula becomes nebulae",
			enabled: true,
			input:   "nebula",
			want:    "nebulae",
		},
		{
			name:    "disabled nebula becomes nebulas",
			enabled: false,
			input:   "nebula",
			want:    "nebulas",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inflect.ClassicalAncient(tt.enabled)
			got := inflect.Plural(tt.input)
			if got != tt.want {
				t.Errorf("ClassicalAncient(%v): Plural(%q) = %q, want %q", tt.enabled, tt.input, got, tt.want)
			}
		})
	}
}

func TestIsClassicalAncient(t *testing.T) {
	// Clean up after test
	defer inflect.ClassicalAncient(false)

	tests := []struct {
		name  string
		setup func()
		want  bool
	}{
		{
			name:  "default is false",
			setup: func() { inflect.ClassicalAncient(false) },
			want:  false,
		},
		{
			name:  "true after enabling",
			setup: func() { inflect.ClassicalAncient(true) },
			want:  true,
		},
		{
			name:  "false after disabling",
			setup: func() { inflect.ClassicalAncient(true); inflect.ClassicalAncient(false) },
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			got := inflect.IsClassicalAncient()
			if got != tt.want {
				t.Errorf("IsClassicalAncient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClassicalAncientIndependentOfClassicalAll(t *testing.T) {
	// Clean up after test
	defer func() {
		inflect.ClassicalAll(false)
		inflect.ClassicalAncient(false)
	}()

	t.Run("ClassicalAncient works independently of ClassicalAll", func(t *testing.T) {
		// Start with ClassicalAll disabled
		inflect.ClassicalAll(false)

		// Enable only ClassicalAncient
		inflect.ClassicalAncient(true)

		// Verify ClassicalAncient is enabled
		if !inflect.IsClassicalAncient() {
			t.Error("IsClassicalAncient() should be true after ClassicalAncient(true)")
		}

		// Verify formula -> formulae
		if got := inflect.Plural("formula"); got != "formulae" {
			t.Errorf("Plural(\"formula\") = %q, want \"formulae\"", got)
		}
	})

	t.Run("ClassicalAncient can be disabled while ClassicalAll was enabled", func(t *testing.T) {
		// Enable all classical options
		inflect.ClassicalAll(true)

		// Verify it's enabled
		if !inflect.IsClassicalAncient() {
			t.Error("IsClassicalAncient() should be true after ClassicalAll(true)")
		}

		// Disable only ClassicalAncient
		inflect.ClassicalAncient(false)

		// Verify ClassicalAncient is now disabled
		if inflect.IsClassicalAncient() {
			t.Error("IsClassicalAncient() should be false after ClassicalAncient(false)")
		}

		// Verify formula -> formulas (modern form)
		if got := inflect.Plural("formula"); got != "formulas" {
			t.Errorf("Plural(\"formula\") = %q, want \"formulas\"", got)
		}
	})
}

func TestClassicalAll(t *testing.T) {
	// Reset to defaults after this test
	defer inflect.ClassicalAll(false)

	tests := []struct {
		name    string
		enabled bool
		word    string
		want    string
	}{
		// Classical mode enabled - Latin/Greek plurals
		{name: "formula classical", enabled: true, word: "formula", want: "formulae"},
		{name: "antenna classical", enabled: true, word: "antenna", want: "antennae"},
		{name: "vertebra classical", enabled: true, word: "vertebra", want: "vertebrae"},
		{name: "alumna classical", enabled: true, word: "alumna", want: "alumnae"},
		{name: "larva classical", enabled: true, word: "larva", want: "larvae"},
		{name: "nebula classical", enabled: true, word: "nebula", want: "nebulae"},
		{name: "nova classical", enabled: true, word: "nova", want: "novae"},
		{name: "supernova classical", enabled: true, word: "supernova", want: "supernovae"},
		{name: "octopus classical", enabled: true, word: "octopus", want: "octopodes"},
		{name: "opus classical", enabled: true, word: "opus", want: "opera"},
		{name: "corpus classical", enabled: true, word: "corpus", want: "corpora"},
		{name: "genus classical", enabled: true, word: "genus", want: "genera"},

		// Classical mode disabled - modern English plurals
		{name: "formula modern", enabled: false, word: "formula", want: "formulas"},
		{name: "antenna modern", enabled: false, word: "antenna", want: "antennas"},
		{name: "vertebra modern", enabled: false, word: "vertebra", want: "vertebras"},
		{name: "alumna modern", enabled: false, word: "alumna", want: "alumnas"},
		{name: "larva modern", enabled: false, word: "larva", want: "larvas"},
		{name: "nebula modern", enabled: false, word: "nebula", want: "nebulas"},
		{name: "nova modern", enabled: false, word: "nova", want: "novas"},

		// Regular words should not be affected
		{name: "cat classical", enabled: true, word: "cat", want: "cats"},
		{name: "cat modern", enabled: false, word: "cat", want: "cats"},
		{name: "box classical", enabled: true, word: "box", want: "boxes"},
		{name: "box modern", enabled: false, word: "box", want: "boxes"},

		// Irregular plurals should still work
		{name: "child classical", enabled: true, word: "child", want: "children"},
		{name: "child modern", enabled: false, word: "child", want: "children"},
		{name: "mouse classical", enabled: true, word: "mouse", want: "mice"},
		{name: "mouse modern", enabled: false, word: "mouse", want: "mice"},

		// Case preservation
		{name: "Formula titlecase classical", enabled: true, word: "Formula", want: "Formulae"},
		{name: "FORMULA uppercase classical", enabled: true, word: "FORMULA", want: "FORMULAE"},
		{name: "Formula titlecase modern", enabled: false, word: "Formula", want: "Formulas"},
		{name: "FORMULA uppercase modern", enabled: false, word: "FORMULA", want: "FORMULAS"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inflect.ClassicalAll(tt.enabled)
			got := inflect.Plural(tt.word)
			if got != tt.want {
				t.Errorf("ClassicalAll(%v): Plural(%q) = %q, want %q", tt.enabled, tt.word, got, tt.want)
			}
		})
	}
}

func TestClassicalAliasForClassicalAll(t *testing.T) {
	// Reset to defaults after this test
	defer inflect.ClassicalAll(false)

	// Verify Classical() is an alias for ClassicalAll()
	tests := []struct {
		name    string
		enabled bool
		word    string
		want    string
	}{
		{name: "formula classical via Classical()", enabled: true, word: "formula", want: "formulae"},
		{name: "formula modern via Classical()", enabled: false, word: "formula", want: "formulas"},
		{name: "antenna classical via Classical()", enabled: true, word: "antenna", want: "antennae"},
		{name: "antenna modern via Classical()", enabled: false, word: "antenna", want: "antennas"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inflect.Classical(tt.enabled) // Use Classical(), not ClassicalAll()
			got := inflect.Plural(tt.word)
			if got != tt.want {
				t.Errorf("Classical(%v): Plural(%q) = %q, want %q", tt.enabled, tt.word, got, tt.want)
			}
		})
	}
}

func TestIsClassicalAll(t *testing.T) {
	// Reset to defaults after this test
	defer inflect.ClassicalAll(false)

	tests := []struct {
		name  string
		setup func()
		want  bool
	}{
		{
			name:  "default is false",
			setup: func() { inflect.ClassicalAll(false) },
			want:  false,
		},
		{
			name:  "enabled via ClassicalAll",
			setup: func() { inflect.ClassicalAll(true) },
			want:  true,
		},
		{
			name:  "enabled via Classical alias",
			setup: func() { inflect.Classical(true) },
			want:  true,
		},
		{
			name: "disabled after being enabled",
			setup: func() {
				inflect.ClassicalAll(true)
				inflect.ClassicalAll(false)
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			got := inflect.IsClassicalAll()
			if got != tt.want {
				t.Errorf("IsClassicalAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClassicalAllIntegration(t *testing.T) {
	// Reset to defaults after this test
	defer inflect.ClassicalAll(false)

	t.Run("complete workflow", func(t *testing.T) {
		// 1. Start with default (modern) pluralization
		inflect.ClassicalAll(false)
		if got := inflect.Plural("formula"); got != "formulas" {
			t.Errorf("Default: Plural(formula) = %q, want %q", got, "formulas")
		}
		if inflect.IsClassicalAll() {
			t.Error("Default: IsClassicalAll() should be false")
		}
		if inflect.IsClassical() {
			t.Error("Default: IsClassical() should be false")
		}

		// 2. Enable classical mode
		inflect.ClassicalAll(true)
		if got := inflect.Plural("formula"); got != "formulae" {
			t.Errorf("Classical: Plural(formula) = %q, want %q", got, "formulae")
		}
		if !inflect.IsClassicalAll() {
			t.Error("Classical: IsClassicalAll() should be true")
		}
		if !inflect.IsClassical() {
			t.Error("Classical: IsClassical() should be true")
		}

		// 3. Verify regular words still work
		if got := inflect.Plural("cat"); got != "cats" {
			t.Errorf("Classical: Plural(cat) = %q, want %q", got, "cats")
		}

		// 4. Verify irregular words still work
		if got := inflect.Plural("child"); got != "children" {
			t.Errorf("Classical: Plural(child) = %q, want %q", got, "children")
		}

		// 5. Disable classical mode
		inflect.ClassicalAll(false)
		if got := inflect.Plural("formula"); got != "formulas" {
			t.Errorf("After disable: Plural(formula) = %q, want %q", got, "formulas")
		}
		if inflect.IsClassicalAll() {
			t.Error("After disable: IsClassicalAll() should be false")
		}

		// 6. Use Classical() alias
		inflect.Classical(true)
		if got := inflect.Plural("antenna"); got != "antennae" {
			t.Errorf("Via alias: Plural(antenna) = %q, want %q", got, "antennae")
		}
		if !inflect.IsClassicalAll() {
			t.Error("Via alias: IsClassicalAll() should be true")
		}
	})
}

func TestClassicalPersons(t *testing.T) {
	// Reset to defaults after this test
	defer inflect.ClassicalPersons(false)

	tests := []struct {
		name    string
		enabled bool
		input   string
		want    string
	}{
		// Classical persons enabled: person -> persons
		{name: "persons enabled lowercase", enabled: true, input: "person", want: "persons"},
		{name: "persons enabled titlecase", enabled: true, input: "Person", want: "Persons"},
		{name: "persons enabled uppercase", enabled: true, input: "PERSON", want: "PERSONS"},

		// Classical persons disabled: person -> people (default)
		{name: "people default lowercase", enabled: false, input: "person", want: "people"},
		{name: "people default titlecase", enabled: false, input: "Person", want: "People"},
		{name: "people default uppercase", enabled: false, input: "PERSON", want: "PEOPLE"},

		// Other words should not be affected
		{name: "cat unaffected enabled", enabled: true, input: "cat", want: "cats"},
		{name: "cat unaffected disabled", enabled: false, input: "cat", want: "cats"},
		{name: "child unaffected enabled", enabled: true, input: "child", want: "children"},
		{name: "child unaffected disabled", enabled: false, input: "child", want: "children"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inflect.ClassicalPersons(tt.enabled)
			got := inflect.Plural(tt.input)
			if got != tt.want {
				t.Errorf("ClassicalPersons(%v): Plural(%q) = %q, want %q",
					tt.enabled, tt.input, got, tt.want)
			}
		})
	}
}

func TestIsClassicalPersons(t *testing.T) {
	// Reset to defaults after this test
	defer inflect.ClassicalPersons(false)

	tests := []struct {
		name  string
		setup func()
		want  bool
	}{
		{
			name:  "default is false",
			setup: func() { inflect.ClassicalPersons(false) },
			want:  false,
		},
		{
			name:  "enabled via ClassicalPersons",
			setup: func() { inflect.ClassicalPersons(true) },
			want:  true,
		},
		{
			name:  "enabled via ClassicalAll",
			setup: func() { inflect.ClassicalAll(true) },
			want:  true,
		},
		{
			name: "disabled after being enabled",
			setup: func() {
				inflect.ClassicalPersons(true)
				inflect.ClassicalPersons(false)
			},
			want: false,
		},
		{
			name: "independent of ClassicalAncient",
			setup: func() {
				inflect.ClassicalAncient(true)
				inflect.ClassicalPersons(false)
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset all classical flags before each test
			inflect.ClassicalAll(false)
			tt.setup()
			got := inflect.IsClassicalPersons()
			if got != tt.want {
				t.Errorf("IsClassicalPersons() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClassicalPersonsIntegration(t *testing.T) {
	// Reset to defaults after this test
	defer inflect.ClassicalAll(false)

	t.Run("complete workflow", func(t *testing.T) {
		// 1. Start with default (modern) pluralization
		inflect.ClassicalAll(false)
		if got := inflect.Plural("person"); got != "people" {
			t.Errorf("Default: Plural(person) = %q, want %q", got, "people")
		}
		if inflect.IsClassicalPersons() {
			t.Error("Default: IsClassicalPersons() should be false")
		}

		// 2. Enable classical persons only
		inflect.ClassicalPersons(true)
		if got := inflect.Plural("person"); got != "persons" {
			t.Errorf("ClassicalPersons: Plural(person) = %q, want %q", got, "persons")
		}
		if !inflect.IsClassicalPersons() {
			t.Error("ClassicalPersons: IsClassicalPersons() should be true")
		}

		// 3. Verify classical ancient is still false
		if inflect.IsClassicalAncient() {
			t.Error("ClassicalPersons only: IsClassicalAncient() should be false")
		}
		if got := inflect.Plural("formula"); got != "formulas" {
			t.Errorf("ClassicalPersons only: Plural(formula) = %q, want %q", got, "formulas")
		}

		// 4. Enable ClassicalAll
		inflect.ClassicalAll(true)
		if got := inflect.Plural("person"); got != "persons" {
			t.Errorf("ClassicalAll: Plural(person) = %q, want %q", got, "persons")
		}
		if got := inflect.Plural("formula"); got != "formulae" {
			t.Errorf("ClassicalAll: Plural(formula) = %q, want %q", got, "formulae")
		}

		// 5. Disable persons but keep ancient
		inflect.ClassicalPersons(false)
		if got := inflect.Plural("person"); got != "people" {
			t.Errorf("Persons off, Ancient on: Plural(person) = %q, want %q", got, "people")
		}
		if got := inflect.Plural("formula"); got != "formulae" {
			t.Errorf("Persons off, Ancient on: Plural(formula) = %q, want %q", got, "formulae")
		}

		// 6. Reset all
		inflect.ClassicalAll(false)
		if got := inflect.Plural("person"); got != "people" {
			t.Errorf("After reset: Plural(person) = %q, want %q", got, "people")
		}
		if got := inflect.Plural("formula"); got != "formulas" {
			t.Errorf("After reset: Plural(formula) = %q, want %q", got, "formulas")
		}
	})
}

func TestClassicalNames(t *testing.T) {
	// Reset to defaults after this test
	defer inflect.ClassicalAll(false)

	tests := []struct {
		name       string
		enabled    bool
		input      string
		want       string
		wantGetter bool
	}{
		// ClassicalNames(false) - regular pluralization for names ending in 's'
		{name: "Jones regular", enabled: false, input: "Jones", want: "Joneses", wantGetter: false},
		{name: "Williams regular", enabled: false, input: "Williams", want: "Williamses", wantGetter: false},
		{name: "Hastings regular", enabled: false, input: "Hastings", want: "Hastingses", wantGetter: false},
		{name: "Ross regular", enabled: false, input: "Ross", want: "Rosses", wantGetter: false},
		{name: "Burns regular", enabled: false, input: "Burns", want: "Burnses", wantGetter: false},

		// ClassicalNames(true) - proper names ending in 's' remain unchanged
		{name: "Jones classical", enabled: true, input: "Jones", want: "Jones", wantGetter: true},
		{name: "Williams classical", enabled: true, input: "Williams", want: "Williams", wantGetter: true},
		{name: "Hastings classical", enabled: true, input: "Hastings", want: "Hastings", wantGetter: true},
		{name: "Ross classical", enabled: true, input: "Ross", want: "Ross", wantGetter: true},
		{name: "Burns classical", enabled: true, input: "Burns", want: "Burns", wantGetter: true},

		// Names not ending in 's' should still pluralize normally
		{name: "Mary classical", enabled: true, input: "Mary", want: "Marys", wantGetter: true},
		{name: "Smith classical", enabled: true, input: "Smith", want: "Smiths", wantGetter: true},
		{name: "Johnson classical", enabled: true, input: "Johnson", want: "Johnsons", wantGetter: true},
		{name: "Mary regular", enabled: false, input: "Mary", want: "Marys", wantGetter: false},
		{name: "Smith regular", enabled: false, input: "Smith", want: "Smiths", wantGetter: false},

		// Lowercase words ending in 's' should NOT be treated as proper names
		{name: "bus classical", enabled: true, input: "bus", want: "buses", wantGetter: true},
		{name: "class classical", enabled: true, input: "class", want: "classes", wantGetter: true},
		{name: "boss classical", enabled: true, input: "boss", want: "bosses", wantGetter: true},

		// All uppercase words (acronyms) should NOT be treated as proper names
		{name: "CBS classical", enabled: true, input: "CBS", want: "CBSES", wantGetter: true},
		{name: "GPS classical", enabled: true, input: "GPS", want: "GPSES", wantGetter: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inflect.ClassicalNames(tt.enabled)

			// Test the getter function
			if got := inflect.IsClassicalNames(); got != tt.wantGetter {
				t.Errorf("IsClassicalNames() = %v, want %v", got, tt.wantGetter)
			}

			// Test pluralization
			got := inflect.Plural(tt.input)
			if got != tt.want {
				t.Errorf("Plural(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestClassicalNamesIntegration(t *testing.T) {
	// Reset to defaults after this test
	defer inflect.ClassicalAll(false)

	t.Run("complete workflow", func(t *testing.T) {
		// 1. Start with default (modern) pluralization
		inflect.ClassicalAll(false)
		if got := inflect.Plural("Jones"); got != "Joneses" {
			t.Errorf("Default: Plural(Jones) = %q, want %q", got, "Joneses")
		}
		if inflect.IsClassicalNames() {
			t.Error("Default: IsClassicalNames() should be false")
		}

		// 2. Enable classical names only
		inflect.ClassicalNames(true)
		if got := inflect.Plural("Jones"); got != "Jones" {
			t.Errorf("ClassicalNames: Plural(Jones) = %q, want %q", got, "Jones")
		}
		if !inflect.IsClassicalNames() {
			t.Error("ClassicalNames: IsClassicalNames() should be true")
		}

		// 3. Verify classical ancient is still false
		if inflect.IsClassicalAncient() {
			t.Error("ClassicalNames only: IsClassicalAncient() should be false")
		}
		if got := inflect.Plural("formula"); got != "formulas" {
			t.Errorf("ClassicalNames only: Plural(formula) = %q, want %q", got, "formulas")
		}

		// 4. Regular nouns ending in 's' should still pluralize normally
		if got := inflect.Plural("bus"); got != "buses" {
			t.Errorf("ClassicalNames: Plural(bus) = %q, want %q", got, "buses")
		}

		// 5. Proper names NOT ending in 's' should still pluralize normally
		if got := inflect.Plural("Smith"); got != "Smiths" {
			t.Errorf("ClassicalNames: Plural(Smith) = %q, want %q", got, "Smiths")
		}

		// 6. Enable ClassicalAll
		inflect.ClassicalAll(true)
		if got := inflect.Plural("Jones"); got != "Jones" {
			t.Errorf("ClassicalAll: Plural(Jones) = %q, want %q", got, "Jones")
		}
		if got := inflect.Plural("formula"); got != "formulae" {
			t.Errorf("ClassicalAll: Plural(formula) = %q, want %q", got, "formulae")
		}

		// 7. Disable names but keep ancient
		inflect.ClassicalNames(false)
		if got := inflect.Plural("Jones"); got != "Joneses" {
			t.Errorf("Names off, Ancient on: Plural(Jones) = %q, want %q", got, "Joneses")
		}
		if got := inflect.Plural("formula"); got != "formulae" {
			t.Errorf("Names off, Ancient on: Plural(formula) = %q, want %q", got, "formulae")
		}

		// 8. Reset all
		inflect.ClassicalAll(false)
		if got := inflect.Plural("Jones"); got != "Joneses" {
			t.Errorf("After reset: Plural(Jones) = %q, want %q", got, "Joneses")
		}
		if got := inflect.Plural("formula"); got != "formulas" {
			t.Errorf("After reset: Plural(formula) = %q, want %q", got, "formulas")
		}
	})
}
