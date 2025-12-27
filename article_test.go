package inflect_test

import (
	"testing"

	inflect "github.com/cv/go-inflect"
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

func TestDefA(t *testing.T) {
	// Reset to defaults after this test
	defer inflect.DefAReset()

	tests := []struct {
		name  string
		word  string
		input string
		want  string
	}{
		// Force "a" on words that normally take "an"
		{name: "ape forced a", word: "ape", input: "ape", want: "a ape"},
		{name: "apple forced a", word: "apple", input: "apple", want: "a apple"},
		{name: "eagle forced a", word: "eagle", input: "eagle", want: "a eagle"},
		{name: "hour forced a", word: "hour", input: "hour", want: "a hour"},

		// Case insensitive matching
		{name: "Ape titlecase", word: "ape", input: "Ape", want: "a Ape"},
		{name: "APE uppercase", word: "ape", input: "APE", want: "a APE"},
		{name: "define uppercase match lowercase", word: "APE", input: "ape", want: "a ape"},

		// Multi-word input (match first word)
		{name: "ape in phrase", word: "ape", input: "ape costume", want: "a ape costume"},
		{name: "eagle in phrase", word: "eagle", input: "eagle scout", want: "a eagle scout"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inflect.DefAReset()
			inflect.DefA(tt.word)

			got := inflect.An(tt.input)
			if got != tt.want {
				t.Errorf("After DefA(%q): An(%q) = %q, want %q", tt.word, tt.input, got, tt.want)
			}
		})
	}
}

func TestDefAn(t *testing.T) {
	// Reset to defaults after this test
	defer inflect.DefAReset()

	tests := []struct {
		name  string
		word  string
		input string
		want  string
	}{
		// Force "an" on words that normally take "a"
		{name: "hero forced an", word: "hero", input: "hero", want: "an hero"},
		{name: "historic forced an", word: "historic", input: "historic", want: "an historic"},
		{name: "unicorn forced an", word: "unicorn", input: "unicorn", want: "an unicorn"},
		{name: "cat forced an", word: "cat", input: "cat", want: "an cat"},

		// Case insensitive matching
		{name: "Hero titlecase", word: "hero", input: "Hero", want: "an Hero"},
		{name: "HERO uppercase", word: "hero", input: "HERO", want: "an HERO"},
		{name: "define uppercase match lowercase", word: "HERO", input: "hero", want: "an hero"},

		// Multi-word input (match first word)
		{name: "hero in phrase", word: "hero", input: "hero complex", want: "an hero complex"},
		{name: "historic in phrase", word: "historic", input: "historic event", want: "an historic event"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inflect.DefAReset()
			inflect.DefAn(tt.word)

			got := inflect.An(tt.input)
			if got != tt.want {
				t.Errorf("After DefAn(%q): An(%q) = %q, want %q", tt.word, tt.input, got, tt.want)
			}
		})
	}
}

func TestUndefA(t *testing.T) {
	// Reset to defaults after this test
	defer inflect.DefAReset()

	tests := []struct {
		name        string
		setupWord   string
		undefWord   string
		wantRemoved bool
		checkInput  string
		wantAfter   string
	}{
		{
			name:        "remove custom a rule",
			setupWord:   "ape",
			undefWord:   "ape",
			wantRemoved: true,
			checkInput:  "ape",
			wantAfter:   "an ape", // default rule restored
		},
		{
			name:        "remove nonexistent rule",
			setupWord:   "",
			undefWord:   "notdefined",
			wantRemoved: false,
			checkInput:  "ape",
			wantAfter:   "an ape",
		},
		{
			name:        "case insensitive removal",
			setupWord:   "ape",
			undefWord:   "APE",
			wantRemoved: true,
			checkInput:  "ape",
			wantAfter:   "an ape",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inflect.DefAReset()

			if tt.setupWord != "" {
				inflect.DefA(tt.setupWord)
			}

			removed := inflect.UndefA(tt.undefWord)
			if removed != tt.wantRemoved {
				t.Errorf("UndefA(%q) = %v, want %v", tt.undefWord, removed, tt.wantRemoved)
			}

			got := inflect.An(tt.checkInput)
			if got != tt.wantAfter {
				t.Errorf("After UndefA: An(%q) = %q, want %q", tt.checkInput, got, tt.wantAfter)
			}
		})
	}
}

func TestUndefAn(t *testing.T) {
	// Reset to defaults after this test
	defer inflect.DefAReset()

	tests := []struct {
		name        string
		setupWord   string
		undefWord   string
		wantRemoved bool
		checkInput  string
		wantAfter   string
	}{
		{
			name:        "remove custom an rule",
			setupWord:   "hero",
			undefWord:   "hero",
			wantRemoved: true,
			checkInput:  "hero",
			wantAfter:   "a hero", // default rule restored
		},
		{
			name:        "remove nonexistent rule",
			setupWord:   "",
			undefWord:   "notdefined",
			wantRemoved: false,
			checkInput:  "hero",
			wantAfter:   "a hero",
		},
		{
			name:        "case insensitive removal",
			setupWord:   "hero",
			undefWord:   "HERO",
			wantRemoved: true,
			checkInput:  "hero",
			wantAfter:   "a hero",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inflect.DefAReset()

			if tt.setupWord != "" {
				inflect.DefAn(tt.setupWord)
			}

			removed := inflect.UndefAn(tt.undefWord)
			if removed != tt.wantRemoved {
				t.Errorf("UndefAn(%q) = %v, want %v", tt.undefWord, removed, tt.wantRemoved)
			}

			got := inflect.An(tt.checkInput)
			if got != tt.wantAfter {
				t.Errorf("After UndefAn: An(%q) = %q, want %q", tt.checkInput, got, tt.wantAfter)
			}
		})
	}
}

func TestDefAReset(t *testing.T) {
	defer inflect.DefAReset()

	tests := []struct {
		name        string
		customA     []string // words to force "a"
		customAn    []string // words to force "an"
		checkInputs map[string]string
	}{
		{
			name:     "reset custom a rules",
			customA:  []string{"ape", "apple", "eagle"},
			customAn: nil,
			checkInputs: map[string]string{
				"ape":   "an ape",   // restored to default
				"apple": "an apple", // restored to default
				"eagle": "an eagle", // restored to default
				"cat":   "a cat",    // unchanged
			},
		},
		{
			name:     "reset custom an rules",
			customA:  nil,
			customAn: []string{"hero", "cat", "dog"},
			checkInputs: map[string]string{
				"hero": "a hero", // restored to default
				"cat":  "a cat",  // restored to default
				"dog":  "a dog",  // restored to default
				"ape":  "an ape", // unchanged
			},
		},
		{
			name:     "reset mixed custom rules",
			customA:  []string{"ape", "eagle"},
			customAn: []string{"hero", "cat"},
			checkInputs: map[string]string{
				"ape":   "an ape",   // restored to default
				"eagle": "an eagle", // restored to default
				"hero":  "a hero",   // restored to default
				"cat":   "a cat",    // restored to default
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inflect.DefAReset()

			// Add custom "a" rules
			for _, word := range tt.customA {
				inflect.DefA(word)
			}

			// Add custom "an" rules
			for _, word := range tt.customAn {
				inflect.DefAn(word)
			}

			// Reset
			inflect.DefAReset()

			// Check results
			for input, want := range tt.checkInputs {
				got := inflect.An(input)
				if got != want {
					t.Errorf("After reset: An(%q) = %q, want %q", input, got, want)
				}
			}
		})
	}
}

func TestDefADefAnPrecedence(t *testing.T) {
	// Reset to defaults after this test
	defer inflect.DefAReset()

	t.Run("DefA takes precedence over DefAn for same word", func(t *testing.T) {
		inflect.DefAReset()

		// First define as "an"
		inflect.DefAn("test")
		if got := inflect.An("test"); got != "an test" {
			t.Errorf("After DefAn: An(test) = %q, want %q", got, "an test")
		}

		// Then override with "a"
		inflect.DefA("test")
		if got := inflect.An("test"); got != "a test" {
			t.Errorf("After DefA override: An(test) = %q, want %q", got, "a test")
		}
	})

	t.Run("DefAn overrides previous DefA", func(t *testing.T) {
		inflect.DefAReset()

		// First define as "a"
		inflect.DefA("test")
		if got := inflect.An("test"); got != "a test" {
			t.Errorf("After DefA: An(test) = %q, want %q", got, "a test")
		}

		// Then override with "an"
		inflect.DefAn("test")
		if got := inflect.An("test"); got != "an test" {
			t.Errorf("After DefAn override: An(test) = %q, want %q", got, "an test")
		}
	})
}

func TestDefAIntegration(t *testing.T) {
	// Reset to defaults after this test
	defer inflect.DefAReset()

	t.Run("complete workflow", func(t *testing.T) {
		inflect.DefAReset()

		// 1. Verify default behavior
		if got := inflect.An("ape"); got != "an ape" {
			t.Errorf("Default: An(ape) = %q, want %q", got, "an ape")
		}
		if got := inflect.An("hero"); got != "a hero" {
			t.Errorf("Default: An(hero) = %q, want %q", got, "a hero")
		}

		// 2. Add custom "a" rule
		inflect.DefA("ape")
		if got := inflect.An("ape"); got != "a ape" {
			t.Errorf("After DefA: An(ape) = %q, want %q", got, "a ape")
		}

		// 3. Add custom "an" rule
		inflect.DefAn("hero")
		if got := inflect.An("hero"); got != "an hero" {
			t.Errorf("After DefAn: An(hero) = %q, want %q", got, "an hero")
		}

		// 4. Remove custom "a" rule
		if removed := inflect.UndefA("ape"); !removed {
			t.Error("UndefA(ape) should return true")
		}
		if got := inflect.An("ape"); got != "an ape" {
			t.Errorf("After UndefA: An(ape) = %q, want %q", got, "an ape")
		}

		// 5. Remove custom "an" rule
		if removed := inflect.UndefAn("hero"); !removed {
			t.Error("UndefAn(hero) should return true")
		}
		if got := inflect.An("hero"); got != "a hero" {
			t.Errorf("After UndefAn: An(hero) = %q, want %q", got, "a hero")
		}

		// 6. Add multiple rules and reset
		inflect.DefA("ape")
		inflect.DefA("eagle")
		inflect.DefAn("hero")
		inflect.DefAn("cat")

		inflect.DefAReset()

		if got := inflect.An("ape"); got != "an ape" {
			t.Errorf("After reset: An(ape) = %q, want %q", got, "an ape")
		}
		if got := inflect.An("eagle"); got != "an eagle" {
			t.Errorf("After reset: An(eagle) = %q, want %q", got, "an eagle")
		}
		if got := inflect.An("hero"); got != "a hero" {
			t.Errorf("After reset: An(hero) = %q, want %q", got, "a hero")
		}
		if got := inflect.An("cat"); got != "a cat" {
			t.Errorf("After reset: An(cat) = %q, want %q", got, "a cat")
		}
	})
}

func TestDefAPattern(t *testing.T) {
	defer inflect.DefAReset()

	tests := []struct {
		name    string
		pattern string
		inputs  []string
		want    string // "a" for all matches
	}{
		{
			name:    "euro prefix pattern",
			pattern: "euro.*",
			inputs:  []string{"euro", "european", "eurozone", "eurocentric"},
			want:    "a",
		},
		{
			name:    "uni prefix pattern",
			pattern: "uni.*",
			inputs:  []string{"unit", "uniform", "universe"},
			want:    "a",
		},
		{
			name:    "exact match pattern",
			pattern: "apple",
			inputs:  []string{"apple"},
			want:    "a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inflect.DefAReset()

			err := inflect.DefAPattern(tt.pattern)
			if err != nil {
				t.Fatalf("DefAPattern(%q) returned error: %v", tt.pattern, err)
			}

			for _, input := range tt.inputs {
				got := inflect.An(input)
				want := tt.want + " " + input
				if got != want {
					t.Errorf("An(%q) = %q, want %q", input, got, want)
				}
			}
		})
	}
}

func TestDefAnPattern(t *testing.T) {
	defer inflect.DefAReset()

	tests := []struct {
		name    string
		pattern string
		inputs  []string
		want    string // "an" for all matches
	}{
		{
			name:    "honor prefix pattern",
			pattern: "honor.*",
			inputs:  []string{"honor", "honorable", "honorary", "honored"},
			want:    "an",
		},
		{
			name:    "heir prefix pattern",
			pattern: "heir.*",
			inputs:  []string{"heir", "heirloom", "heiress"},
			want:    "an",
		},
		{
			name:    "exact match pattern",
			pattern: "cat",
			inputs:  []string{"cat"},
			want:    "an",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inflect.DefAReset()

			err := inflect.DefAnPattern(tt.pattern)
			if err != nil {
				t.Fatalf("DefAnPattern(%q) returned error: %v", tt.pattern, err)
			}

			for _, input := range tt.inputs {
				got := inflect.An(input)
				want := tt.want + " " + input
				if got != want {
					t.Errorf("An(%q) = %q, want %q", input, got, want)
				}
			}
		})
	}
}

func TestDefAPatternInvalidRegex(t *testing.T) {
	defer inflect.DefAReset()

	err := inflect.DefAPattern("[invalid")
	if err == nil {
		t.Error("DefAPattern with invalid regex should return error")
	}

	err = inflect.DefAnPattern("[invalid")
	if err == nil {
		t.Error("DefAnPattern with invalid regex should return error")
	}
}

func TestUndefAPattern(t *testing.T) {
	defer inflect.DefAReset()

	t.Run("remove existing pattern", func(t *testing.T) {
		inflect.DefAReset()

		// Use "apple" which normally takes "an" but pattern forces "a"
		if err := inflect.DefAPattern("apple.*"); err != nil {
			t.Fatalf("DefAPattern failed: %v", err)
		}
		if got := inflect.An("appleton"); got != "a appleton" {
			t.Errorf("Before UndefAPattern: An(appleton) = %q, want %q", got, "a appleton")
		}

		// Remove pattern
		if removed := inflect.UndefAPattern("apple.*"); !removed {
			t.Error("UndefAPattern should return true for existing pattern")
		}

		// Verify default behavior restored (words starting with vowel get "an")
		if got := inflect.An("appleton"); got != "an appleton" {
			t.Errorf("After UndefAPattern: An(appleton) = %q, want %q", got, "an appleton")
		}
	})

	t.Run("remove non-existing pattern", func(t *testing.T) {
		inflect.DefAReset()

		if removed := inflect.UndefAPattern("nonexistent.*"); removed {
			t.Error("UndefAPattern should return false for non-existing pattern")
		}
	})
}

func TestUndefAnPattern(t *testing.T) {
	defer inflect.DefAReset()

	t.Run("remove existing pattern", func(t *testing.T) {
		inflect.DefAReset()

		// Add pattern and verify it works
		if err := inflect.DefAnPattern("hero.*"); err != nil {
			t.Fatalf("DefAnPattern failed: %v", err)
		}
		if got := inflect.An("heroic"); got != "an heroic" {
			t.Errorf("Before UndefAnPattern: An(heroic) = %q, want %q", got, "an heroic")
		}

		// Remove pattern
		if removed := inflect.UndefAnPattern("hero.*"); !removed {
			t.Error("UndefAnPattern should return true for existing pattern")
		}

		// Verify default behavior restored
		if got := inflect.An("heroic"); got != "a heroic" {
			t.Errorf("After UndefAnPattern: An(heroic) = %q, want %q", got, "a heroic")
		}
	})

	t.Run("remove non-existing pattern", func(t *testing.T) {
		inflect.DefAReset()

		if removed := inflect.UndefAnPattern("nonexistent.*"); removed {
			t.Error("UndefAnPattern should return false for non-existing pattern")
		}
	})
}

func TestDefAResetClearsPatterns(t *testing.T) {
	defer inflect.DefAReset()

	// Add both exact matches and patterns
	inflect.DefA("apple")
	inflect.DefAn("cat")
	if err := inflect.DefAPattern("euro.*"); err != nil {
		t.Fatalf("DefAPattern failed: %v", err)
	}
	if err := inflect.DefAnPattern("honor.*"); err != nil {
		t.Fatalf("DefAnPattern failed: %v", err)
	}

	// Verify patterns are working
	if got := inflect.An("apple"); got != "a apple" {
		t.Errorf("Before reset: An(apple) = %q, want %q", got, "a apple")
	}
	if got := inflect.An("european"); got != "a european" {
		t.Errorf("Before reset: An(european) = %q, want %q", got, "a european")
	}
	if got := inflect.An("honorable"); got != "an honorable" {
		t.Errorf("Before reset: An(honorable) = %q, want %q", got, "an honorable")
	}

	// Reset
	inflect.DefAReset()

	// Verify all patterns are cleared (back to defaults)
	if got := inflect.An("apple"); got != "an apple" {
		t.Errorf("After reset: An(apple) = %q, want %q", got, "an apple")
	}
	// "european" defaults to "a" because "eu" sounds like "you"
	if got := inflect.An("european"); got != "a european" {
		t.Errorf("After reset: An(european) = %q, want %q", got, "a european")
	}
	// "honorable" defaults to "an" because the "h" is silent
	if got := inflect.An("honorable"); got != "an honorable" {
		t.Errorf("After reset: An(honorable) = %q, want %q", got, "an honorable")
	}
}

func TestPatternPrecedence(t *testing.T) {
	defer inflect.DefAReset()

	t.Run("exact word takes precedence over pattern", func(t *testing.T) {
		inflect.DefAReset()

		// Add pattern first
		if err := inflect.DefAnPattern("euro.*"); err != nil {
			t.Fatalf("DefAnPattern failed: %v", err)
		}
		if got := inflect.An("euro"); got != "an euro" {
			t.Errorf("With pattern only: An(euro) = %q, want %q", got, "an euro")
		}

		// Add exact word match - should take precedence
		inflect.DefA("euro")
		if got := inflect.An("euro"); got != "a euro" {
			t.Errorf("With exact word override: An(euro) = %q, want %q", got, "a euro")
		}

		// Other words matching pattern still work
		if got := inflect.An("european"); got != "an european" {
			t.Errorf("Pattern still matches: An(european) = %q, want %q", got, "an european")
		}
	})

	t.Run("DefAPattern takes precedence over DefAnPattern", func(t *testing.T) {
		inflect.DefAReset()

		// Both patterns match "european"
		if err := inflect.DefAnPattern("euro.*"); err != nil {
			t.Fatalf("DefAnPattern failed: %v", err)
		}
		if err := inflect.DefAPattern("europ.*"); err != nil {
			t.Fatalf("DefAPattern failed: %v", err)
		}

		// DefAPattern should take precedence
		if got := inflect.An("european"); got != "a european" {
			t.Errorf("An(european) = %q, want %q", got, "a european")
		}

		// "euro" only matches DefAnPattern
		if got := inflect.An("euro"); got != "an euro" {
			t.Errorf("An(euro) = %q, want %q", got, "an euro")
		}
	})
}

func TestPatternCaseInsensitive(t *testing.T) {
	defer inflect.DefAReset()

	if err := inflect.DefAPattern("euro.*"); err != nil {
		t.Fatalf("DefAPattern failed: %v", err)
	}

	// Pattern should match regardless of case
	tests := []struct {
		input string
		want  string
	}{
		{"euro", "a euro"},
		{"Euro", "a Euro"},
		{"EURO", "a EURO"},
		{"european", "a european"},
		{"European", "a European"},
		{"EUROPEAN", "a EUROPEAN"},
	}

	for _, tt := range tests {
		if got := inflect.An(tt.input); got != tt.want {
			t.Errorf("An(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func BenchmarkAn(b *testing.B) {
	// Test with inputs covering different article selection rules
	benchmarks := []struct {
		name  string
		input string
	}{
		{"consonant", "cat"},           // a cat
		{"vowel", "apple"},             // an apple
		{"silent_h", "honest person"},  // an honest
		{"consonant_u", "university"},  // a university (y sound)
		{"abbreviation", "FBI agent"},  // an FBI
		{"phrase", "elegant solution"}, // an elegant
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for range b.N {
				inflect.An(bm.input)
			}
		})
	}
}
