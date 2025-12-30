package inflect_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	inflect "github.com/cv/go-inflect/v2"
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
			assert.Equal(t, tt.enabled, got)
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
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestClassicalIntegration(t *testing.T) {
	// Clean up after test
	defer inflect.Classical(false)

	t.Run("complete workflow", func(t *testing.T) {
		// 1. Initial/default state should be false
		inflect.Classical(false)
		assert.False(t, inflect.IsClassical(), "Initial IsClassical()")

		// 2. Enable classical mode
		inflect.Classical(true)
		assert.True(t, inflect.IsClassical(), "After Classical(true): IsClassical()")

		// 3. Disable classical mode
		inflect.Classical(false)
		assert.False(t, inflect.IsClassical(), "After Classical(false): IsClassical()")

		// 4. Toggle multiple times
		inflect.Classical(true)
		inflect.Classical(true) // Setting to same value should work
		assert.True(t, inflect.IsClassical(), "After double Classical(true): IsClassical()")

		inflect.Classical(false)
		inflect.Classical(false) // Setting to same value should work
		assert.False(t, inflect.IsClassical(), "After double Classical(false): IsClassical()")
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
			assert.Equal(t, tt.want, got)
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
			assert.Equal(t, tt.want, got)
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
		assert.True(t, inflect.IsClassicalAncient(), "IsClassicalAncient() should be true after ClassicalAncient(true)")

		// Verify formula -> formulae
		assert.Equal(t, "formulae", inflect.Plural("formula"), "Plural(\"formula\")")
	})

	t.Run("ClassicalAncient can be disabled while ClassicalAll was enabled", func(t *testing.T) {
		// Enable all classical options
		inflect.ClassicalAll(true)

		// Verify it's enabled
		assert.True(t, inflect.IsClassicalAncient(), "IsClassicalAncient() should be true after ClassicalAll(true)")

		// Disable only ClassicalAncient
		inflect.ClassicalAncient(false)

		// Verify ClassicalAncient is now disabled
		assert.False(t, inflect.IsClassicalAncient(), "IsClassicalAncient() should be false after ClassicalAncient(false)")

		// Verify formula -> formulas (modern form)
		assert.Equal(t, "formulas", inflect.Plural("formula"), "Plural(\"formula\")")
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
			assert.Equal(t, tt.want, got)
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
			assert.Equal(t, tt.want, got)
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
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestClassicalAllIntegration(t *testing.T) {
	// Reset to defaults after this test
	defer inflect.ClassicalAll(false)

	t.Run("complete workflow", func(t *testing.T) {
		// 1. Start with default (modern) pluralization
		inflect.ClassicalAll(false)
		assert.Equal(t, "formulas", inflect.Plural("formula"), "Default: Plural(formula)")
		assert.False(t, inflect.IsClassicalAll(), "Default: IsClassicalAll()")
		assert.False(t, inflect.IsClassical(), "Default: IsClassical()")

		// 2. Enable classical mode
		inflect.ClassicalAll(true)
		assert.Equal(t, "formulae", inflect.Plural("formula"), "Classical: Plural(formula)")
		assert.True(t, inflect.IsClassicalAll(), "Classical: IsClassicalAll()")
		assert.True(t, inflect.IsClassical(), "Classical: IsClassical()")

		// 3. Verify regular words still work
		assert.Equal(t, "cats", inflect.Plural("cat"), "Classical: Plural(cat)")

		// 4. Verify irregular words still work
		assert.Equal(t, "children", inflect.Plural("child"), "Classical: Plural(child)")

		// 5. Disable classical mode
		inflect.ClassicalAll(false)
		assert.Equal(t, "formulas", inflect.Plural("formula"), "After disable: Plural(formula)")
		assert.False(t, inflect.IsClassicalAll(), "After disable: IsClassicalAll()")

		// 6. Use Classical() alias
		inflect.Classical(true)
		assert.Equal(t, "antennae", inflect.Plural("antenna"), "Via alias: Plural(antenna)")
		assert.True(t, inflect.IsClassicalAll(), "Via alias: IsClassicalAll()")
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
			assert.Equal(t, tt.want, got)
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
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestClassicalPersonsIntegration(t *testing.T) {
	// Reset to defaults after this test
	defer inflect.ClassicalAll(false)

	t.Run("complete workflow", func(t *testing.T) {
		// 1. Start with default (modern) pluralization
		inflect.ClassicalAll(false)
		assert.Equal(t, "people", inflect.Plural("person"), "Default: Plural(person)")
		assert.False(t, inflect.IsClassicalPersons(), "Default: IsClassicalPersons()")

		// 2. Enable classical persons only
		inflect.ClassicalPersons(true)
		assert.Equal(t, "persons", inflect.Plural("person"), "ClassicalPersons: Plural(person)")
		assert.True(t, inflect.IsClassicalPersons(), "ClassicalPersons: IsClassicalPersons()")

		// 3. Verify classical ancient is still false
		assert.False(t, inflect.IsClassicalAncient(), "ClassicalPersons only: IsClassicalAncient()")
		assert.Equal(t, "formulas", inflect.Plural("formula"), "ClassicalPersons only: Plural(formula)")

		// 4. Enable ClassicalAll
		inflect.ClassicalAll(true)
		assert.Equal(t, "persons", inflect.Plural("person"), "ClassicalAll: Plural(person)")
		assert.Equal(t, "formulae", inflect.Plural("formula"), "ClassicalAll: Plural(formula)")

		// 5. Disable persons but keep ancient
		inflect.ClassicalPersons(false)
		assert.Equal(t, "people", inflect.Plural("person"), "Persons off, Ancient on: Plural(person)")
		assert.Equal(t, "formulae", inflect.Plural("formula"), "Persons off, Ancient on: Plural(formula)")

		// 6. Reset all
		inflect.ClassicalAll(false)
		assert.Equal(t, "people", inflect.Plural("person"), "After reset: Plural(person)")
		assert.Equal(t, "formulas", inflect.Plural("formula"), "After reset: Plural(formula)")
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
			assert.Equal(t, tt.wantGetter, inflect.IsClassicalNames())

			// Test pluralization
			got := inflect.Plural(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestClassicalNamesIntegration(t *testing.T) {
	// Reset to defaults after this test
	defer inflect.ClassicalAll(false)

	t.Run("complete workflow", func(t *testing.T) {
		// 1. Start with default (modern) pluralization
		inflect.ClassicalAll(false)
		assert.Equal(t, "Joneses", inflect.Plural("Jones"), "Default: Plural(Jones)")
		assert.False(t, inflect.IsClassicalNames(), "Default: IsClassicalNames()")

		// 2. Enable classical names only
		inflect.ClassicalNames(true)
		assert.Equal(t, "Jones", inflect.Plural("Jones"), "ClassicalNames: Plural(Jones)")
		assert.True(t, inflect.IsClassicalNames(), "ClassicalNames: IsClassicalNames()")

		// 3. Verify classical ancient is still false
		assert.False(t, inflect.IsClassicalAncient(), "ClassicalNames only: IsClassicalAncient()")
		assert.Equal(t, "formulas", inflect.Plural("formula"), "ClassicalNames only: Plural(formula)")

		// 4. Regular nouns ending in 's' should still pluralize normally
		assert.Equal(t, "buses", inflect.Plural("bus"), "ClassicalNames: Plural(bus)")

		// 5. Proper names NOT ending in 's' should still pluralize normally
		assert.Equal(t, "Smiths", inflect.Plural("Smith"), "ClassicalNames: Plural(Smith)")

		// 6. Enable ClassicalAll
		inflect.ClassicalAll(true)
		assert.Equal(t, "Jones", inflect.Plural("Jones"), "ClassicalAll: Plural(Jones)")
		assert.Equal(t, "formulae", inflect.Plural("formula"), "ClassicalAll: Plural(formula)")

		// 7. Disable names but keep ancient
		inflect.ClassicalNames(false)
		assert.Equal(t, "Joneses", inflect.Plural("Jones"), "Names off, Ancient on: Plural(Jones)")
		assert.Equal(t, "formulae", inflect.Plural("formula"), "Names off, Ancient on: Plural(formula)")

		// 8. Reset all
		inflect.ClassicalAll(false)
		assert.Equal(t, "Joneses", inflect.Plural("Jones"), "After reset: Plural(Jones)")
		assert.Equal(t, "formulas", inflect.Plural("formula"), "After reset: Plural(formula)")
	})
}

func TestClassicalHerd(t *testing.T) {
	defer inflect.ClassicalHerd(false)

	tests := []struct {
		name    string
		enabled bool
		input   string
		want    string
	}{
		{name: "bison classical", enabled: true, input: "bison", want: "bison"},
		{name: "bison modern", enabled: false, input: "bison", want: "bisons"},
		{name: "buffalo classical", enabled: true, input: "buffalo", want: "buffalo"},
		{name: "buffalo modern", enabled: false, input: "buffalo", want: "buffaloes"},
		{name: "wildebeest classical", enabled: true, input: "wildebeest", want: "wildebeest"},
		{name: "wildebeest modern", enabled: false, input: "wildebeest", want: "wildebeests"},
		// Regular animals unaffected
		{name: "cat classical", enabled: true, input: "cat", want: "cats"},
		{name: "cat modern", enabled: false, input: "cat", want: "cats"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inflect.ClassicalHerd(tt.enabled)
			got := inflect.Plural(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestIsClassicalHerd(t *testing.T) {
	defer inflect.ClassicalAll(false)

	tests := []struct {
		name  string
		setup func()
		want  bool
	}{
		{name: "default is false", setup: func() { inflect.ClassicalHerd(false) }, want: false},
		{name: "enabled", setup: func() { inflect.ClassicalHerd(true) }, want: true},
		{name: "enabled via ClassicalAll", setup: func() { inflect.ClassicalAll(true) }, want: true},
		{name: "disabled after enabled", setup: func() { inflect.ClassicalHerd(true); inflect.ClassicalHerd(false) }, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inflect.ClassicalAll(false)
			tt.setup()
			assert.Equal(t, tt.want, inflect.IsClassicalHerd())
		})
	}
}

func TestClassicalZero(t *testing.T) {
	defer inflect.ClassicalZero(false)

	tests := []struct {
		name    string
		enabled bool
		word    string
		count   int
		want    string
	}{
		{name: "zero classical", enabled: true, word: "cat", count: 0, want: "no cat"},
		{name: "zero modern", enabled: false, word: "cat", count: 0, want: "no cats"},
		{name: "one classical", enabled: true, word: "cat", count: 1, want: "1 cat"},
		{name: "one modern", enabled: false, word: "cat", count: 1, want: "1 cat"},
		{name: "two classical", enabled: true, word: "cat", count: 2, want: "2 cats"},
		{name: "two modern", enabled: false, word: "cat", count: 2, want: "2 cats"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inflect.ClassicalZero(tt.enabled)
			got := inflect.No(tt.word, tt.count)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestIsClassicalZero(t *testing.T) {
	defer inflect.ClassicalAll(false)

	tests := []struct {
		name  string
		setup func()
		want  bool
	}{
		{name: "default is false", setup: func() { inflect.ClassicalZero(false) }, want: false},
		{name: "enabled", setup: func() { inflect.ClassicalZero(true) }, want: true},
		{name: "enabled via ClassicalAll", setup: func() { inflect.ClassicalAll(true) }, want: true},
		{name: "disabled after enabled", setup: func() { inflect.ClassicalZero(true); inflect.ClassicalZero(false) }, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inflect.ClassicalAll(false)
			tt.setup()
			assert.Equal(t, tt.want, inflect.IsClassicalZero())
		})
	}
}

func TestIsClassicalNames(t *testing.T) {
	defer inflect.ClassicalAll(false)

	tests := []struct {
		name  string
		setup func()
		want  bool
	}{
		{name: "default is false", setup: func() { inflect.ClassicalNames(false) }, want: false},
		{name: "enabled", setup: func() { inflect.ClassicalNames(true) }, want: true},
		{name: "enabled via ClassicalAll", setup: func() { inflect.ClassicalAll(true) }, want: true},
		{name: "disabled after enabled", setup: func() { inflect.ClassicalNames(true); inflect.ClassicalNames(false) }, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inflect.ClassicalAll(false)
			tt.setup()
			assert.Equal(t, tt.want, inflect.IsClassicalNames())
		})
	}
}
