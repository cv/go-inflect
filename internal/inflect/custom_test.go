package inflect_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	inflect "github.com/cv/go-inflect"
)

func TestDefNoun(t *testing.T) {
	// Reset to defaults after this test
	defer inflect.DefNounReset()

	tests := []struct {
		name     string
		singular string
		plural   string
	}{
		{name: "simple custom", singular: "foo", plural: "foos"},
		{name: "irregular custom", singular: "gizmo", plural: "gizmata"},
		{name: "override builtin", singular: "child", plural: "childs"},
		{name: "unicode custom", singular: "café", plural: "cafés"},
		{name: "empty plural", singular: "widget", plural: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset before each test
			inflect.DefNounReset()

			// Define the custom rule
			inflect.DefNoun(tt.singular, tt.plural)

			// Test Plural()
			got := inflect.Plural(tt.singular)
			assert.Equal(t, tt.plural, got)

			// Test Singular() for reverse lookup
			if tt.plural != "" {
				gotSingular := inflect.Singular(tt.plural)
				assert.Equal(t, tt.singular, gotSingular)
			}
		})
	}
}

func TestDefNounCasePreservation(t *testing.T) {
	// Reset to defaults after this test
	defer inflect.DefNounReset()

	tests := []struct {
		name         string
		singular     string
		plural       string
		inputWord    string
		wantPlural   string
		inputPlural  string
		wantSingular string
	}{
		{
			name:         "lowercase to lowercase",
			singular:     "foo",
			plural:       "foos",
			inputWord:    "foo",
			wantPlural:   "foos",
			inputPlural:  "foos",
			wantSingular: "foo",
		},
		{
			name:         "titlecase input",
			singular:     "foo",
			plural:       "foos",
			inputWord:    "Foo",
			wantPlural:   "Foos",
			inputPlural:  "Foos",
			wantSingular: "Foo",
		},
		{
			name:         "uppercase input",
			singular:     "foo",
			plural:       "foos",
			inputWord:    "FOO",
			wantPlural:   "FOOS",
			inputPlural:  "FOOS",
			wantSingular: "FOO",
		},
		{
			name:         "irregular custom titlecase",
			singular:     "gizmo",
			plural:       "gizmata",
			inputWord:    "Gizmo",
			wantPlural:   "Gizmata",
			inputPlural:  "Gizmata",
			wantSingular: "Gizmo",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inflect.DefNounReset()
			inflect.DefNoun(tt.singular, tt.plural)

			got := inflect.Plural(tt.inputWord)
			assert.Equal(t, tt.wantPlural, got)

			gotSingular := inflect.Singular(tt.inputPlural)
			assert.Equal(t, tt.wantSingular, gotSingular)
		})
	}
}

func TestUndefNoun(t *testing.T) {
	// Reset to defaults after this test
	defer inflect.DefNounReset()

	tests := []struct {
		name            string
		setupSingular   string
		setupPlural     string
		undefSingular   string
		wantRemoved     bool
		checkWord       string
		wantPluralAfter string // what Plural() returns after undef
	}{
		{
			name:            "remove custom rule",
			setupSingular:   "foo",
			setupPlural:     "foos",
			undefSingular:   "foo",
			wantRemoved:     true,
			checkWord:       "foo",
			wantPluralAfter: "foos", // standard rule applies
		},
		{
			name:            "remove nonexistent rule",
			setupSingular:   "",
			setupPlural:     "",
			undefSingular:   "notdefined",
			wantRemoved:     false,
			checkWord:       "cat",
			wantPluralAfter: "cats",
		},
		{
			name:            "cannot remove builtin",
			setupSingular:   "",
			setupPlural:     "",
			undefSingular:   "child",
			wantRemoved:     false,
			checkWord:       "child",
			wantPluralAfter: "children",
		},
		{
			name:            "case insensitive removal",
			setupSingular:   "bar",
			setupPlural:     "barz",
			undefSingular:   "BAR",
			wantRemoved:     true,
			checkWord:       "bar",
			wantPluralAfter: "bars", // standard rule applies
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inflect.DefNounReset()

			// Setup custom rule if specified
			if tt.setupSingular != "" {
				inflect.DefNoun(tt.setupSingular, tt.setupPlural)
			}

			// Attempt to remove
			removed := inflect.UndefNoun(tt.undefSingular)
			assert.Equal(t, tt.wantRemoved, removed)

			// Check pluralization after removal
			got := inflect.Plural(tt.checkWord)
			assert.Equal(t, tt.wantPluralAfter, got)
		})
	}
}

func TestDefNounReset(t *testing.T) {
	defer inflect.DefNounReset()

	tests := []struct {
		name          string
		customRules   map[string]string // singular -> plural
		overrideRules map[string]string // override built-in rules
		checkWords    map[string]string // word -> expected plural after reset
	}{
		{
			name: "reset custom rules",
			customRules: map[string]string{
				"foo":    "foos",
				"bar":    "barz",
				"widget": "widgetz",
			},
			checkWords: map[string]string{
				"foo":    "foos", // standard rule
				"bar":    "bars", // standard rule
				"widget": "widgets",
				"child":  "children", // builtin preserved
			},
		},
		{
			name: "reset overridden builtins",
			overrideRules: map[string]string{
				"child": "childs",
				"mouse": "mouses",
			},
			checkWords: map[string]string{
				"child": "children", // builtin restored
				"mouse": "mice",     // builtin restored
			},
		},
		{
			name: "reset mixed custom and overrides",
			customRules: map[string]string{
				"gizmo": "gizmata",
			},
			overrideRules: map[string]string{
				"foot": "foots",
			},
			checkWords: map[string]string{
				"gizmo": "gizmoes", // standard rule (consonant + o -> oes)
				"foot":  "feet",    // builtin restored
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Start fresh
			inflect.DefNounReset()

			// Add custom rules
			for singular, plural := range tt.customRules {
				inflect.DefNoun(singular, plural)
			}

			// Override builtins
			for singular, plural := range tt.overrideRules {
				inflect.DefNoun(singular, plural)
			}

			// Reset
			inflect.DefNounReset()

			// Check results
			for word, wantPlural := range tt.checkWords {
				got := inflect.Plural(word)
				assert.Equal(t, wantPlural, got)
			}
		})
	}
}

func TestDefNounIntegration(t *testing.T) {
	// Reset to defaults after this test
	defer inflect.DefNounReset()

	// Test a complete workflow
	t.Run("complete workflow", func(t *testing.T) {
		inflect.DefNounReset()

		// 1. Verify default behavior
		assert.Equal(t, "children", inflect.Plural("child"), "Default: Plural(child)")

		// 2. Add custom rule
		inflect.DefNoun("foo", "foozles")
		assert.Equal(t, "foozles", inflect.Plural("foo"), "After DefNoun: Plural(foo)")
		assert.Equal(t, "foo", inflect.Singular("foozles"), "After DefNoun: Singular(foozles)")

		// 3. Override builtin
		inflect.DefNoun("child", "childs")
		assert.Equal(t, "childs", inflect.Plural("child"), "After override: Plural(child)")

		// 4. Remove custom rule (but not builtin)
		assert.True(t, inflect.UndefNoun("foo"), "UndefNoun(foo) should return true")
		assert.Equal(t, "foos", inflect.Plural("foo"), "After UndefNoun: Plural(foo)")

		// 5. Cannot remove builtin (even if overridden)
		assert.False(t, inflect.UndefNoun("child"), "UndefNoun(child) should return false for builtin")

		// 6. Reset everything
		inflect.DefNounReset()
		assert.Equal(t, "children", inflect.Plural("child"), "After reset: Plural(child)")
		assert.Equal(t, "foos", inflect.Plural("foo"), "After reset: Plural(foo)")
	})
}

func TestDefVerb(t *testing.T) {
	// Reset to defaults after this test
	defer inflect.DefVerbReset()

	tests := []struct {
		name     string
		singular string
		plural   string
	}{
		{name: "simple verb", singular: "run", plural: "runs"},
		{name: "irregular verb", singular: "be", plural: "are"},
		{name: "custom verb", singular: "foo", plural: "foos"},
		{name: "unicode verb", singular: "café", plural: "cafés"},
		{name: "empty plural", singular: "widget", plural: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			inflect.DefVerbReset()

			// Define the custom rule - this should not panic
			inflect.DefVerb(tt.singular, tt.plural)

			// Since DefVerb is a stub, we just verify it doesn't panic
			// and the function completes successfully
		})
	}
}

func TestDefVerbCaseInsensitive(t *testing.T) {
	defer inflect.DefVerbReset()

	// DefVerb should store in lowercase
	inflect.DefVerb("Run", "Runs")

	// Verify the lowercase key can be undefined
	assert.True(t, inflect.UndefVerb("run"), "UndefVerb(run) should return true after DefVerb(Run, Runs)")
}

func TestUndefVerb(t *testing.T) {
	defer inflect.DefVerbReset()

	tests := []struct {
		name          string
		setupSingular string
		setupPlural   string
		undefSingular string
		wantRemoved   bool
	}{
		{
			name:          "remove existing rule",
			setupSingular: "run",
			setupPlural:   "runs",
			undefSingular: "run",
			wantRemoved:   true,
		},
		{
			name:          "remove nonexistent rule",
			setupSingular: "",
			setupPlural:   "",
			undefSingular: "notdefined",
			wantRemoved:   false,
		},
		{
			name:          "case insensitive removal",
			setupSingular: "walk",
			setupPlural:   "walks",
			undefSingular: "WALK",
			wantRemoved:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inflect.DefVerbReset()

			if tt.setupSingular != "" {
				inflect.DefVerb(tt.setupSingular, tt.setupPlural)
			}

			removed := inflect.UndefVerb(tt.undefSingular)
			assert.Equal(t, tt.wantRemoved, removed)
		})
	}
}

func TestDefVerbReset(t *testing.T) {
	defer inflect.DefVerbReset()

	// Add some custom rules
	inflect.DefVerb("foo", "foos")
	inflect.DefVerb("bar", "bars")
	inflect.DefVerb("baz", "bazzes")

	// Reset
	inflect.DefVerbReset()

	// Verify rules are gone - UndefVerb should return false
	assert.False(t, inflect.UndefVerb("foo"), "After reset: UndefVerb(foo) should return false")
	assert.False(t, inflect.UndefVerb("bar"), "After reset: UndefVerb(bar) should return false")
	assert.False(t, inflect.UndefVerb("baz"), "After reset: UndefVerb(baz) should return false")
}

func TestDefVerbIntegration(t *testing.T) {
	defer inflect.DefVerbReset()

	t.Run("complete workflow", func(t *testing.T) {
		inflect.DefVerbReset()

		// 1. Add custom rule
		inflect.DefVerb("run", "runs")

		// 2. Verify rule exists (via UndefVerb returning true)
		// Need to re-add since UndefVerb removes it
		inflect.DefVerb("run", "runs")

		// 3. Add more rules
		inflect.DefVerb("walk", "walks")
		inflect.DefVerb("be", "are")

		// 4. Remove one rule
		assert.True(t, inflect.UndefVerb("walk"), "UndefVerb(walk) should return true")

		// 5. Verify removed rule is gone
		assert.False(t, inflect.UndefVerb("walk"), "UndefVerb(walk) should return false after removal")

		// 6. Other rules still exist
		assert.True(t, inflect.UndefVerb("run"), "UndefVerb(run) should return true")

		// 7. Reset and verify all gone
		inflect.DefVerb("test", "tests")
		inflect.DefVerbReset()

		assert.False(t, inflect.UndefVerb("test"), "After reset: UndefVerb(test) should return false")
		assert.False(t, inflect.UndefVerb("be"), "After reset: UndefVerb(be) should return false")
	})
}

func TestDefAdj(t *testing.T) {
	// Reset to defaults after this test
	defer inflect.DefAdjReset()

	tests := []struct {
		name     string
		singular string
		plural   string
	}{
		{name: "simple adjective", singular: "big", plural: "bigs"},
		{name: "irregular adjective", singular: "happy", plural: "happies"},
		{name: "custom adjective", singular: "foo", plural: "foos"},
		{name: "unicode adjective", singular: "café", plural: "cafés"},
		{name: "empty plural", singular: "widget", plural: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			inflect.DefAdjReset()

			// Define the custom rule - this should not panic
			inflect.DefAdj(tt.singular, tt.plural)

			// Since DefAdj is a stub, we just verify it doesn't panic
			// and the function completes successfully
		})
	}
}

func TestDefAdjCaseInsensitive(t *testing.T) {
	defer inflect.DefAdjReset()

	// DefAdj should store in lowercase
	inflect.DefAdj("Big", "Bigs")

	// Verify the lowercase key can be undefined
	assert.True(t, inflect.UndefAdj("big"), "UndefAdj(big) should return true after DefAdj(Big, Bigs)")
}

func TestUndefAdj(t *testing.T) {
	defer inflect.DefAdjReset()

	tests := []struct {
		name          string
		setupSingular string
		setupPlural   string
		undefSingular string
		wantRemoved   bool
	}{
		{
			name:          "remove existing rule",
			setupSingular: "big",
			setupPlural:   "bigs",
			undefSingular: "big",
			wantRemoved:   true,
		},
		{
			name:          "remove nonexistent rule",
			setupSingular: "",
			setupPlural:   "",
			undefSingular: "notdefined",
			wantRemoved:   false,
		},
		{
			name:          "case insensitive removal",
			setupSingular: "small",
			setupPlural:   "smalls",
			undefSingular: "SMALL",
			wantRemoved:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inflect.DefAdjReset()

			if tt.setupSingular != "" {
				inflect.DefAdj(tt.setupSingular, tt.setupPlural)
			}

			removed := inflect.UndefAdj(tt.undefSingular)
			assert.Equal(t, tt.wantRemoved, removed)
		})
	}
}

func TestDefAdjReset(t *testing.T) {
	defer inflect.DefAdjReset()

	// Add some custom rules
	inflect.DefAdj("foo", "foos")
	inflect.DefAdj("bar", "bars")
	inflect.DefAdj("baz", "bazzes")

	// Reset
	inflect.DefAdjReset()

	// Verify rules are gone - UndefAdj should return false
	assert.False(t, inflect.UndefAdj("foo"), "After reset: UndefAdj(foo) should return false")
	assert.False(t, inflect.UndefAdj("bar"), "After reset: UndefAdj(bar) should return false")
	assert.False(t, inflect.UndefAdj("baz"), "After reset: UndefAdj(baz) should return false")
}

func TestDefAdjIntegration(t *testing.T) {
	defer inflect.DefAdjReset()

	t.Run("complete workflow", func(t *testing.T) {
		inflect.DefAdjReset()

		// 1. Add custom rule
		inflect.DefAdj("big", "bigs")

		// 2. Verify rule exists (via UndefAdj returning true)
		// Need to re-add since UndefAdj removes it
		inflect.DefAdj("big", "bigs")

		// 3. Add more rules
		inflect.DefAdj("small", "smalls")
		inflect.DefAdj("happy", "happies")

		// 4. Remove one rule
		assert.True(t, inflect.UndefAdj("small"), "UndefAdj(small) should return true")

		// 5. Verify removed rule is gone
		assert.False(t, inflect.UndefAdj("small"), "UndefAdj(small) should return false after removal")

		// 6. Other rules still exist
		assert.True(t, inflect.UndefAdj("big"), "UndefAdj(big) should return true")

		// 7. Reset and verify all gone
		inflect.DefAdj("test", "tests")
		inflect.DefAdjReset()

		assert.False(t, inflect.UndefAdj("test"), "After reset: UndefAdj(test) should return false")
		assert.False(t, inflect.UndefAdj("happy"), "After reset: UndefAdj(happy) should return false")
	})
}

// Tests for jinzhu/inflection compatibility aliases

func TestAddIrregular(t *testing.T) {
	defer inflect.DefNounReset()

	// AddIrregular should behave exactly like DefNoun
	inflect.AddIrregular("person", "people")
	assert.Equal(t, "people", inflect.Plural("person"))
	assert.Equal(t, "person", inflect.Singular("people"))

	// Case preservation
	assert.Equal(t, "People", inflect.Plural("Person"))
	assert.Equal(t, "PEOPLE", inflect.Plural("PERSON"))
}

func TestAddUncountable(t *testing.T) {
	defer inflect.DefNounReset()

	// Single word
	inflect.AddUncountable("rice")
	assert.Equal(t, "rice", inflect.Plural("rice"))
	assert.Equal(t, "rice", inflect.Singular("rice"))

	// Multiple words
	inflect.AddUncountable("fish", "sheep", "deer")
	assert.Equal(t, "fish", inflect.Plural("fish"))
	assert.Equal(t, "sheep", inflect.Plural("sheep"))
	assert.Equal(t, "deer", inflect.Plural("deer"))

	// Case preservation
	assert.Equal(t, "Fish", inflect.Plural("Fish"))
	assert.Equal(t, "SHEEP", inflect.Plural("SHEEP"))
}

// Tests for go-openapi/inflect compatibility aliases

func TestPluralize(t *testing.T) {
	// Pluralize should behave exactly like Plural
	assert.Equal(t, inflect.Plural("cat"), inflect.Pluralize("cat"))
	assert.Equal(t, inflect.Plural("child"), inflect.Pluralize("child"))
	assert.Equal(t, inflect.Plural("Person"), inflect.Pluralize("Person"))
}

func TestSingularize(t *testing.T) {
	// Singularize should behave exactly like Singular
	assert.Equal(t, inflect.Singular("cats"), inflect.Singularize("cats"))
	assert.Equal(t, inflect.Singular("children"), inflect.Singularize("children"))
	assert.Equal(t, inflect.Singular("People"), inflect.Singularize("People"))
}

func TestCamelize(t *testing.T) {
	// Camelize should behave exactly like PascalCase
	assert.Equal(t, inflect.PascalCase("hello_world"), inflect.Camelize("hello_world"))
	assert.Equal(t, inflect.PascalCase("foo-bar"), inflect.Camelize("foo-bar"))
	assert.Equal(t, inflect.PascalCase("some_thing"), inflect.Camelize("some_thing"))
}

func TestCamelizeDownFirst(t *testing.T) {
	// CamelizeDownFirst should behave exactly like CamelCase
	assert.Equal(t, inflect.CamelCase("hello_world"), inflect.CamelizeDownFirst("hello_world"))
	assert.Equal(t, inflect.CamelCase("foo-bar"), inflect.CamelizeDownFirst("foo-bar"))
	assert.Equal(t, inflect.CamelCase("some_thing"), inflect.CamelizeDownFirst("some_thing"))
}
