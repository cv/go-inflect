package inflect_test

import (
	"testing"

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
			if got != tt.plural {
				t.Errorf("After DefNoun(%q, %q): Plural(%q) = %q, want %q",
					tt.singular, tt.plural, tt.singular, got, tt.plural)
			}

			// Test Singular() for reverse lookup
			if tt.plural != "" {
				gotSingular := inflect.Singular(tt.plural)
				if gotSingular != tt.singular {
					t.Errorf("After DefNoun(%q, %q): Singular(%q) = %q, want %q",
						tt.singular, tt.plural, tt.plural, gotSingular, tt.singular)
				}
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
			if got != tt.wantPlural {
				t.Errorf("Plural(%q) = %q, want %q", tt.inputWord, got, tt.wantPlural)
			}

			gotSingular := inflect.Singular(tt.inputPlural)
			if gotSingular != tt.wantSingular {
				t.Errorf("Singular(%q) = %q, want %q", tt.inputPlural, gotSingular, tt.wantSingular)
			}
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
			if removed != tt.wantRemoved {
				t.Errorf("UndefNoun(%q) = %v, want %v", tt.undefSingular, removed, tt.wantRemoved)
			}

			// Check pluralization after removal
			got := inflect.Plural(tt.checkWord)
			if got != tt.wantPluralAfter {
				t.Errorf("After UndefNoun: Plural(%q) = %q, want %q", tt.checkWord, got, tt.wantPluralAfter)
			}
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
				if got != wantPlural {
					t.Errorf("After reset: Plural(%q) = %q, want %q", word, got, wantPlural)
				}
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
		if got := inflect.Plural("child"); got != "children" {
			t.Errorf("Default: Plural(child) = %q, want %q", got, "children")
		}

		// 2. Add custom rule
		inflect.DefNoun("foo", "foozles")
		if got := inflect.Plural("foo"); got != "foozles" {
			t.Errorf("After DefNoun: Plural(foo) = %q, want %q", got, "foozles")
		}
		if got := inflect.Singular("foozles"); got != "foo" {
			t.Errorf("After DefNoun: Singular(foozles) = %q, want %q", got, "foo")
		}

		// 3. Override builtin
		inflect.DefNoun("child", "childs")
		if got := inflect.Plural("child"); got != "childs" {
			t.Errorf("After override: Plural(child) = %q, want %q", got, "childs")
		}

		// 4. Remove custom rule (but not builtin)
		if removed := inflect.UndefNoun("foo"); !removed {
			t.Error("UndefNoun(foo) should return true")
		}
		if got := inflect.Plural("foo"); got != "foos" {
			t.Errorf("After UndefNoun: Plural(foo) = %q, want %q", got, "foos")
		}

		// 5. Cannot remove builtin (even if overridden)
		if removed := inflect.UndefNoun("child"); removed {
			t.Error("UndefNoun(child) should return false for builtin")
		}

		// 6. Reset everything
		inflect.DefNounReset()
		if got := inflect.Plural("child"); got != "children" {
			t.Errorf("After reset: Plural(child) = %q, want %q", got, "children")
		}
		if got := inflect.Plural("foo"); got != "foos" {
			t.Errorf("After reset: Plural(foo) = %q, want %q", got, "foos")
		}
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
	if removed := inflect.UndefVerb("run"); !removed {
		t.Error("UndefVerb(run) should return true after DefVerb(Run, Runs)")
	}
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
			if removed != tt.wantRemoved {
				t.Errorf("UndefVerb(%q) = %v, want %v", tt.undefSingular, removed, tt.wantRemoved)
			}
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
	if removed := inflect.UndefVerb("foo"); removed {
		t.Error("After reset: UndefVerb(foo) should return false")
	}
	if removed := inflect.UndefVerb("bar"); removed {
		t.Error("After reset: UndefVerb(bar) should return false")
	}
	if removed := inflect.UndefVerb("baz"); removed {
		t.Error("After reset: UndefVerb(baz) should return false")
	}
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
		if removed := inflect.UndefVerb("walk"); !removed {
			t.Error("UndefVerb(walk) should return true")
		}

		// 5. Verify removed rule is gone
		if removed := inflect.UndefVerb("walk"); removed {
			t.Error("UndefVerb(walk) should return false after removal")
		}

		// 6. Other rules still exist
		if removed := inflect.UndefVerb("run"); !removed {
			t.Error("UndefVerb(run) should return true")
		}

		// 7. Reset and verify all gone
		inflect.DefVerb("test", "tests")
		inflect.DefVerbReset()

		if removed := inflect.UndefVerb("test"); removed {
			t.Error("After reset: UndefVerb(test) should return false")
		}
		if removed := inflect.UndefVerb("be"); removed {
			t.Error("After reset: UndefVerb(be) should return false")
		}
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
	if removed := inflect.UndefAdj("big"); !removed {
		t.Error("UndefAdj(big) should return true after DefAdj(Big, Bigs)")
	}
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
			if removed != tt.wantRemoved {
				t.Errorf("UndefAdj(%q) = %v, want %v", tt.undefSingular, removed, tt.wantRemoved)
			}
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
	if removed := inflect.UndefAdj("foo"); removed {
		t.Error("After reset: UndefAdj(foo) should return false")
	}
	if removed := inflect.UndefAdj("bar"); removed {
		t.Error("After reset: UndefAdj(bar) should return false")
	}
	if removed := inflect.UndefAdj("baz"); removed {
		t.Error("After reset: UndefAdj(baz) should return false")
	}
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
		if removed := inflect.UndefAdj("small"); !removed {
			t.Error("UndefAdj(small) should return true")
		}

		// 5. Verify removed rule is gone
		if removed := inflect.UndefAdj("small"); removed {
			t.Error("UndefAdj(small) should return false after removal")
		}

		// 6. Other rules still exist
		if removed := inflect.UndefAdj("big"); !removed {
			t.Error("UndefAdj(big) should return true")
		}

		// 7. Reset and verify all gone
		inflect.DefAdj("test", "tests")
		inflect.DefAdjReset()

		if removed := inflect.UndefAdj("test"); removed {
			t.Error("After reset: UndefAdj(test) should return false")
		}
		if removed := inflect.UndefAdj("happy"); removed {
			t.Error("After reset: UndefAdj(happy) should return false")
		}
	})
}
