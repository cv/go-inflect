package inflect_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

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
			assert.Equal(t, tt.want, got, "An(%q)", tt.input)
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
			assert.Equal(t, tt.want, got, "A(%q)", tt.input)
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
			assert.Equal(t, tt.want, got, "After DefA(%q): An(%q)", tt.word, tt.input)
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
			assert.Equal(t, tt.want, got, "After DefAn(%q): An(%q)", tt.word, tt.input)
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
			assert.Equal(t, tt.wantRemoved, removed, "UndefA(%q)", tt.undefWord)

			got := inflect.An(tt.checkInput)
			assert.Equal(t, tt.wantAfter, got, "After UndefA: An(%q)", tt.checkInput)
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
			assert.Equal(t, tt.wantRemoved, removed, "UndefAn(%q)", tt.undefWord)

			got := inflect.An(tt.checkInput)
			assert.Equal(t, tt.wantAfter, got, "After UndefAn: An(%q)", tt.checkInput)
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
				assert.Equal(t, want, got, "After reset: An(%q)", input)
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
		assert.Equal(t, "an test", inflect.An("test"), "After DefAn")

		// Then override with "a"
		inflect.DefA("test")
		assert.Equal(t, "a test", inflect.An("test"), "After DefA override")
	})

	t.Run("DefAn overrides previous DefA", func(t *testing.T) {
		inflect.DefAReset()

		// First define as "a"
		inflect.DefA("test")
		assert.Equal(t, "a test", inflect.An("test"), "After DefA")

		// Then override with "an"
		inflect.DefAn("test")
		assert.Equal(t, "an test", inflect.An("test"), "After DefAn override")
	})
}

func TestDefAIntegration(t *testing.T) {
	// Reset to defaults after this test
	defer inflect.DefAReset()

	t.Run("complete workflow", func(t *testing.T) {
		inflect.DefAReset()

		// 1. Verify default behavior
		assert.Equal(t, "an ape", inflect.An("ape"), "Default: An(ape)")
		assert.Equal(t, "a hero", inflect.An("hero"), "Default: An(hero)")

		// 2. Add custom "a" rule
		inflect.DefA("ape")
		assert.Equal(t, "a ape", inflect.An("ape"), "After DefA: An(ape)")

		// 3. Add custom "an" rule
		inflect.DefAn("hero")
		assert.Equal(t, "an hero", inflect.An("hero"), "After DefAn: An(hero)")

		// 4. Remove custom "a" rule
		assert.True(t, inflect.UndefA("ape"), "UndefA(ape) should return true")
		assert.Equal(t, "an ape", inflect.An("ape"), "After UndefA: An(ape)")

		// 5. Remove custom "an" rule
		assert.True(t, inflect.UndefAn("hero"), "UndefAn(hero) should return true")
		assert.Equal(t, "a hero", inflect.An("hero"), "After UndefAn: An(hero)")

		// 6. Add multiple rules and reset
		inflect.DefA("ape")
		inflect.DefA("eagle")
		inflect.DefAn("hero")
		inflect.DefAn("cat")

		inflect.DefAReset()

		assert.Equal(t, "an ape", inflect.An("ape"), "After reset: An(ape)")
		assert.Equal(t, "an eagle", inflect.An("eagle"), "After reset: An(eagle)")
		assert.Equal(t, "a hero", inflect.An("hero"), "After reset: An(hero)")
		assert.Equal(t, "a cat", inflect.An("cat"), "After reset: An(cat)")
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
			assert.NoError(t, err, "DefAPattern(%q)", tt.pattern)

			for _, input := range tt.inputs {
				got := inflect.An(input)
				want := tt.want + " " + input
				assert.Equal(t, want, got, "An(%q)", input)
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
			assert.NoError(t, err, "DefAnPattern(%q)", tt.pattern)

			for _, input := range tt.inputs {
				got := inflect.An(input)
				want := tt.want + " " + input
				assert.Equal(t, want, got, "An(%q)", input)
			}
		})
	}
}

func TestDefAPatternInvalidRegex(t *testing.T) {
	defer inflect.DefAReset()

	err := inflect.DefAPattern("[invalid")
	assert.Error(t, err, "DefAPattern with invalid regex should return error")

	err = inflect.DefAnPattern("[invalid")
	assert.Error(t, err, "DefAnPattern with invalid regex should return error")
}

func TestUndefAPattern(t *testing.T) {
	defer inflect.DefAReset()

	t.Run("remove existing pattern", func(t *testing.T) {
		inflect.DefAReset()

		// Use "apple" which normally takes "an" but pattern forces "a"
		err := inflect.DefAPattern("apple.*")
		assert.NoError(t, err, "DefAPattern failed")
		assert.Equal(t, "a appleton", inflect.An("appleton"), "Before UndefAPattern")

		// Remove pattern
		assert.True(t, inflect.UndefAPattern("apple.*"), "UndefAPattern should return true for existing pattern")

		// Verify default behavior restored (words starting with vowel get "an")
		assert.Equal(t, "an appleton", inflect.An("appleton"), "After UndefAPattern")
	})

	t.Run("remove non-existing pattern", func(t *testing.T) {
		inflect.DefAReset()

		assert.False(t, inflect.UndefAPattern("nonexistent.*"), "UndefAPattern should return false for non-existing pattern")
	})
}

func TestUndefAnPattern(t *testing.T) {
	defer inflect.DefAReset()

	t.Run("remove existing pattern", func(t *testing.T) {
		inflect.DefAReset()

		// Add pattern and verify it works
		err := inflect.DefAnPattern("hero.*")
		assert.NoError(t, err, "DefAnPattern failed")
		assert.Equal(t, "an heroic", inflect.An("heroic"), "Before UndefAnPattern")

		// Remove pattern
		assert.True(t, inflect.UndefAnPattern("hero.*"), "UndefAnPattern should return true for existing pattern")

		// Verify default behavior restored
		assert.Equal(t, "a heroic", inflect.An("heroic"), "After UndefAnPattern")
	})

	t.Run("remove non-existing pattern", func(t *testing.T) {
		inflect.DefAReset()

		assert.False(t, inflect.UndefAnPattern("nonexistent.*"), "UndefAnPattern should return false for non-existing pattern")
	})
}

func TestDefAResetClearsPatterns(t *testing.T) {
	defer inflect.DefAReset()

	// Add both exact matches and patterns
	inflect.DefA("apple")
	inflect.DefAn("cat")
	err := inflect.DefAPattern("euro.*")
	assert.NoError(t, err, "DefAPattern failed")
	err = inflect.DefAnPattern("honor.*")
	assert.NoError(t, err, "DefAnPattern failed")

	// Verify patterns are working
	assert.Equal(t, "a apple", inflect.An("apple"), "Before reset: An(apple)")
	assert.Equal(t, "a european", inflect.An("european"), "Before reset: An(european)")
	assert.Equal(t, "an honorable", inflect.An("honorable"), "Before reset: An(honorable)")

	// Reset
	inflect.DefAReset()

	// Verify all patterns are cleared (back to defaults)
	assert.Equal(t, "an apple", inflect.An("apple"), "After reset: An(apple)")
	// "european" defaults to "a" because "eu" sounds like "you"
	assert.Equal(t, "a european", inflect.An("european"), "After reset: An(european)")
	// "honorable" defaults to "an" because the "h" is silent
	assert.Equal(t, "an honorable", inflect.An("honorable"), "After reset: An(honorable)")
}

func TestPatternPrecedence(t *testing.T) {
	defer inflect.DefAReset()

	t.Run("exact word takes precedence over pattern", func(t *testing.T) {
		inflect.DefAReset()

		// Add pattern first
		err := inflect.DefAnPattern("euro.*")
		assert.NoError(t, err, "DefAnPattern failed")
		assert.Equal(t, "an euro", inflect.An("euro"), "With pattern only")

		// Add exact word match - should take precedence
		inflect.DefA("euro")
		assert.Equal(t, "a euro", inflect.An("euro"), "With exact word override")

		// Other words matching pattern still work
		assert.Equal(t, "an european", inflect.An("european"), "Pattern still matches")
	})

	t.Run("DefAPattern takes precedence over DefAnPattern", func(t *testing.T) {
		inflect.DefAReset()

		// Both patterns match "european"
		err := inflect.DefAnPattern("euro.*")
		assert.NoError(t, err, "DefAnPattern failed")
		err = inflect.DefAPattern("europ.*")
		assert.NoError(t, err, "DefAPattern failed")

		// DefAPattern should take precedence
		assert.Equal(t, "a european", inflect.An("european"))

		// "euro" only matches DefAnPattern
		assert.Equal(t, "an euro", inflect.An("euro"))
	})
}

func TestPatternCaseInsensitive(t *testing.T) {
	defer inflect.DefAReset()

	err := inflect.DefAPattern("euro.*")
	assert.NoError(t, err, "DefAPattern failed")

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
		assert.Equal(t, tt.want, inflect.An(tt.input), "An(%q)", tt.input)
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
