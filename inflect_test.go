package inflect_test

import (
	"testing"

	inflect "gitlab-master.nvidia.com/urg/go-inflect"
)

func TestJoin(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  string
	}{
		// Empty slice
		{name: "empty slice", input: []string{}, want: ""},
		{name: "nil slice", input: nil, want: ""},

		// Single word
		{name: "single word", input: []string{"apple"}, want: "apple"},
		{name: "single empty string", input: []string{""}, want: ""},

		// Two words
		{name: "two words", input: []string{"apple", "banana"}, want: "apple and banana"},
		{name: "two short words", input: []string{"a", "b"}, want: "a and b"},

		// Three+ words (Oxford comma)
		{name: "three words", input: []string{"a", "b", "c"}, want: "a, b, and c"},
		{name: "four words", input: []string{"apple", "banana", "cherry", "date"}, want: "apple, banana, cherry, and date"},
		{name: "five words", input: []string{"one", "two", "three", "four", "five"}, want: "one, two, three, four, and five"},

		// Words with special characters
		{name: "words with commas", input: []string{"red, blue", "green"}, want: "red, blue and green"},
		{name: "words with quotes", input: []string{`"hello"`, `"world"`}, want: `"hello" and "world"`},
		{name: "words with unicode", input: []string{"café", "naïve", "résumé"}, want: "café, naïve, and résumé"},
		{name: "words with numbers", input: []string{"item1", "item2", "item3"}, want: "item1, item2, and item3"},
		{name: "words with spaces", input: []string{"New York", "Los Angeles", "Chicago"}, want: "New York, Los Angeles, and Chicago"},
		{name: "words with ampersand", input: []string{"Tom & Jerry", "Mickey"}, want: "Tom & Jerry and Mickey"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Join(tt.input)
			if got != tt.want {
				t.Errorf("Join(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestJoinWithConj(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		conj  string
		want  string
	}{
		// Empty slice
		{name: "empty slice with or", input: []string{}, conj: "or", want: ""},
		{name: "nil slice with or", input: nil, conj: "or", want: ""},

		// Single word
		{name: "single word with or", input: []string{"apple"}, conj: "or", want: "apple"},
		{name: "single empty string with or", input: []string{""}, conj: "or", want: ""},

		// Two words with "or"
		{name: "two words with or", input: []string{"apple", "banana"}, conj: "or", want: "apple or banana"},
		{name: "two short words with or", input: []string{"a", "b"}, conj: "or", want: "a or b"},

		// Three+ words with "or" (Oxford comma)
		{name: "three words with or", input: []string{"a", "b", "c"}, conj: "or", want: "a, b, or c"},
		{name: "four words with or", input: []string{"apple", "banana", "cherry", "date"}, conj: "or", want: "apple, banana, cherry, or date"},

		// Other conjunctions
		{name: "two words with and/or", input: []string{"a", "b"}, conj: "and/or", want: "a and/or b"},
		{name: "three words with and/or", input: []string{"a", "b", "c"}, conj: "and/or", want: "a, b, and/or c"},
		{name: "two words with nor", input: []string{"this", "that"}, conj: "nor", want: "this nor that"},
		{name: "three words with nor", input: []string{"this", "that", "other"}, conj: "nor", want: "this, that, nor other"},

		// Verify "and" conjunction matches Join() behavior
		{name: "two words with and", input: []string{"a", "b"}, conj: "and", want: "a and b"},
		{name: "three words with and", input: []string{"a", "b", "c"}, conj: "and", want: "a, b, and c"},

		// Words with special characters
		{name: "words with spaces and or", input: []string{"New York", "Los Angeles"}, conj: "or", want: "New York or Los Angeles"},
		{name: "words with unicode and or", input: []string{"café", "thé", "chocolat"}, conj: "or", want: "café, thé, or chocolat"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.JoinWithConj(tt.input, tt.conj)
			if got != tt.want {
				t.Errorf("JoinWithConj(%q, %q) = %q, want %q", tt.input, tt.conj, got, tt.want)
			}
		})
	}
}

func TestJoinWithSep(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		conj  string
		sep   string
		want  string
	}{
		// Empty slice
		{name: "empty slice", input: []string{}, conj: "and", sep: "; ", want: ""},
		{name: "nil slice", input: nil, conj: "and", sep: "; ", want: ""},

		// Single word
		{name: "single word", input: []string{"apple"}, conj: "and", sep: "; ", want: "apple"},

		// Two words (separator not used between two items)
		{name: "two words", input: []string{"a", "b"}, conj: "and", sep: "; ", want: "a and b"},
		{name: "two words with or", input: []string{"a", "b"}, conj: "or", sep: "; ", want: "a or b"},

		// Three+ words with semicolon separator
		{name: "three words semicolon", input: []string{"a", "b", "c"}, conj: "and", sep: "; ", want: "a; b; and c"},
		{name: "four words semicolon", input: []string{"a", "b", "c", "d"}, conj: "and", sep: "; ", want: "a; b; c; and d"},
		{name: "three words semicolon or", input: []string{"a", "b", "c"}, conj: "or", sep: "; ", want: "a; b; or c"},

		// Items containing commas (primary use case)
		{name: "dates with commas", input: []string{"Jan 1, 2020", "Feb 2, 2021", "Mar 3, 2022"}, conj: "and", sep: "; ", want: "Jan 1, 2020; Feb 2, 2021; and Mar 3, 2022"},
		{name: "locations with commas", input: []string{"New York, NY", "Los Angeles, CA", "Chicago, IL"}, conj: "or", sep: "; ", want: "New York, NY; Los Angeles, CA; or Chicago, IL"},
		{name: "names with titles", input: []string{"Smith, John", "Doe, Jane"}, conj: "and", sep: "; ", want: "Smith, John and Doe, Jane"},

		// Custom separators
		{name: "pipe separator", input: []string{"a", "b", "c"}, conj: "and", sep: " | ", want: "a | b | and c"},
		{name: "dash separator", input: []string{"x", "y", "z"}, conj: "or", sep: " - ", want: "x - y - or z"},
		{name: "newline separator", input: []string{"line1", "line2", "line3"}, conj: "and", sep: "\n", want: "line1\nline2\nand line3"},

		// Verify comma separator matches JoinWithConj behavior
		{name: "comma separator three", input: []string{"a", "b", "c"}, conj: "and", sep: ", ", want: "a, b, and c"},
		{name: "comma separator four", input: []string{"a", "b", "c", "d"}, conj: "or", sep: ", ", want: "a, b, c, or d"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.JoinWithSep(tt.input, tt.conj, tt.sep)
			if got != tt.want {
				t.Errorf("JoinWithSep(%q, %q, %q) = %q, want %q", tt.input, tt.conj, tt.sep, got, tt.want)
			}
		})
	}
}

func TestOrdinalWord(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  string
	}{
		// Basic cases (1-10)
		{name: "first", input: 1, want: "first"},
		{name: "second", input: 2, want: "second"},
		{name: "third", input: 3, want: "third"},
		{name: "fourth", input: 4, want: "fourth"},
		{name: "fifth", input: 5, want: "fifth"},
		{name: "sixth", input: 6, want: "sixth"},
		{name: "seventh", input: 7, want: "seventh"},
		{name: "eighth", input: 8, want: "eighth"},
		{name: "ninth", input: 9, want: "ninth"},
		{name: "tenth", input: 10, want: "tenth"},

		// Teens (11-19)
		{name: "eleventh", input: 11, want: "eleventh"},
		{name: "twelfth", input: 12, want: "twelfth"},
		{name: "thirteenth", input: 13, want: "thirteenth"},
		{name: "fourteenth", input: 14, want: "fourteenth"},
		{name: "fifteenth", input: 15, want: "fifteenth"},
		{name: "sixteenth", input: 16, want: "sixteenth"},
		{name: "seventeenth", input: 17, want: "seventeenth"},
		{name: "eighteenth", input: 18, want: "eighteenth"},
		{name: "nineteenth", input: 19, want: "nineteenth"},

		// Exact tens (20, 30, ...)
		{name: "twentieth", input: 20, want: "twentieth"},
		{name: "thirtieth", input: 30, want: "thirtieth"},
		{name: "fortieth", input: 40, want: "fortieth"},
		{name: "fiftieth", input: 50, want: "fiftieth"},
		{name: "sixtieth", input: 60, want: "sixtieth"},
		{name: "seventieth", input: 70, want: "seventieth"},
		{name: "eightieth", input: 80, want: "eightieth"},
		{name: "ninetieth", input: 90, want: "ninetieth"},

		// Compound tens (21-29, etc.)
		{name: "twenty-first", input: 21, want: "twenty-first"},
		{name: "twenty-second", input: 22, want: "twenty-second"},
		{name: "twenty-third", input: 23, want: "twenty-third"},
		{name: "twenty-fourth", input: 24, want: "twenty-fourth"},
		{name: "thirty-fifth", input: 35, want: "thirty-fifth"},
		{name: "forty-second", input: 42, want: "forty-second"},
		{name: "ninety-ninth", input: 99, want: "ninety-ninth"},

		// Hundreds
		{name: "one hundredth", input: 100, want: "one hundredth"},
		{name: "one hundred first", input: 101, want: "one hundred first"},
		{name: "one hundred second", input: 102, want: "one hundred second"},
		{name: "one hundred tenth", input: 110, want: "one hundred tenth"},
		{name: "one hundred eleventh", input: 111, want: "one hundred eleventh"},
		{name: "one hundred twentieth", input: 120, want: "one hundred twentieth"},
		{name: "one hundred twenty-first", input: 121, want: "one hundred twenty-first"},
		{name: "two hundredth", input: 200, want: "two hundredth"},
		{name: "five hundred fifty-fifth", input: 555, want: "five hundred fifty-fifth"},
		{name: "nine hundred ninety-ninth", input: 999, want: "nine hundred ninety-ninth"},

		// Thousands
		{name: "one thousandth", input: 1000, want: "one thousandth"},
		{name: "one thousand first", input: 1001, want: "one thousand first"},
		{name: "one thousand tenth", input: 1010, want: "one thousand tenth"},
		{name: "one thousand one hundredth", input: 1100, want: "one thousand one hundredth"},
		{name: "two thousandth", input: 2000, want: "two thousandth"},
		{name: "twelve thousandth", input: 12000, want: "twelve thousandth"},
		{name: "twenty-one thousandth", input: 21000, want: "twenty-one thousandth"},

		// Large numbers
		{name: "one millionth", input: 1000000, want: "one millionth"},
		{name: "one million first", input: 1000001, want: "one million first"},
		{name: "one billionth", input: 1000000000, want: "one billionth"},

		// Edge cases
		{name: "zeroth", input: 0, want: "zeroth"},
		{name: "negative first", input: -1, want: "negative first"},
		{name: "negative second", input: -2, want: "negative second"},
		{name: "negative eleventh", input: -11, want: "negative eleventh"},
		{name: "negative twenty-first", input: -21, want: "negative twenty-first"},
		{name: "negative one hundredth", input: -100, want: "negative one hundredth"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.OrdinalWord(tt.input)
			if got != tt.want {
				t.Errorf("OrdinalWord(%d) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestOrdinal(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  string
	}{
		// Basic cases (1st, 2nd, 3rd, 4th-9th)
		{name: "first", input: 1, want: "1st"},
		{name: "second", input: 2, want: "2nd"},
		{name: "third", input: 3, want: "3rd"},
		{name: "fourth", input: 4, want: "4th"},
		{name: "fifth", input: 5, want: "5th"},
		{name: "sixth", input: 6, want: "6th"},
		{name: "seventh", input: 7, want: "7th"},
		{name: "eighth", input: 8, want: "8th"},
		{name: "ninth", input: 9, want: "9th"},
		{name: "tenth", input: 10, want: "10th"},

		// Teens (11th, 12th, 13th) - special cases
		{name: "eleventh", input: 11, want: "11th"},
		{name: "twelfth", input: 12, want: "12th"},
		{name: "thirteenth", input: 13, want: "13th"},
		{name: "fourteenth", input: 14, want: "14th"},
		{name: "nineteenth", input: 19, want: "19th"},

		// Larger numbers (21st, 22nd, 23rd, 100th, 101st)
		{name: "twenty-first", input: 21, want: "21st"},
		{name: "twenty-second", input: 22, want: "22nd"},
		{name: "twenty-third", input: 23, want: "23rd"},
		{name: "twenty-fourth", input: 24, want: "24th"},
		{name: "hundredth", input: 100, want: "100th"},
		{name: "hundred-first", input: 101, want: "101st"},
		{name: "hundred-second", input: 102, want: "102nd"},
		{name: "hundred-third", input: 103, want: "103rd"},
		{name: "hundred-eleventh", input: 111, want: "111th"},
		{name: "hundred-twelfth", input: 112, want: "112th"},
		{name: "hundred-thirteenth", input: 113, want: "113th"},
		{name: "thousandth", input: 1000, want: "1000th"},
		{name: "thousand-first", input: 1001, want: "1001st"},
		{name: "thousand-eleventh", input: 1011, want: "1011th"},

		// Edge cases
		{name: "zero", input: 0, want: "0th"},
		{name: "negative one", input: -1, want: "-1st"},
		{name: "negative two", input: -2, want: "-2nd"},
		{name: "negative three", input: -3, want: "-3rd"},
		{name: "negative eleven", input: -11, want: "-11th"},
		{name: "negative twenty-one", input: -21, want: "-21st"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Ordinal(tt.input)
			if got != tt.want {
				t.Errorf("Ordinal(%d) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

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

// TestA verifies that A() is an alias for An()
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

func TestPresentParticiple(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		// Empty string
		{name: "empty", input: "", want: ""},

		// Single letter verbs
		{name: "single letter", input: "a", want: "aing"},

		// Already ending in -ing
		{name: "already -ing", input: "running", want: "running"},
		{name: "already -ing sing", input: "sing", want: "singing"},

		// Double consonant (CVC pattern)
		{name: "run", input: "run", want: "running"},
		{name: "sit", input: "sit", want: "sitting"},
		{name: "hit", input: "hit", want: "hitting"},
		{name: "cut", input: "cut", want: "cutting"},
		{name: "stop", input: "stop", want: "stopping"},
		{name: "drop", input: "drop", want: "dropping"},
		{name: "plan", input: "plan", want: "planning"},
		{name: "skip", input: "skip", want: "skipping"},
		{name: "begin", input: "begin", want: "beginning"},
		{name: "occur", input: "occur", want: "occurring"},
		{name: "prefer", input: "prefer", want: "preferring"},
		{name: "admit", input: "admit", want: "admitting"},
		{name: "commit", input: "commit", want: "committing"},
		{name: "regret", input: "regret", want: "regretting"},

		// Drop silent e
		{name: "make", input: "make", want: "making"},
		{name: "take", input: "take", want: "taking"},
		{name: "come", input: "come", want: "coming"},
		{name: "give", input: "give", want: "giving"},
		{name: "have", input: "have", want: "having"},
		{name: "write", input: "write", want: "writing"},
		{name: "live", input: "live", want: "living"},
		{name: "move", input: "move", want: "moving"},
		{name: "hope", input: "hope", want: "hoping"},
		{name: "dance", input: "dance", want: "dancing"},

		// Just add -ing (no changes needed)
		{name: "play", input: "play", want: "playing"},
		{name: "stay", input: "stay", want: "staying"},
		{name: "enjoy", input: "enjoy", want: "enjoying"},
		{name: "show", input: "show", want: "showing"},
		{name: "follow", input: "follow", want: "following"},
		{name: "fix", input: "fix", want: "fixing"},
		{name: "mix", input: "mix", want: "mixing"},
		{name: "go", input: "go", want: "going"},
		{name: "do", input: "do", want: "doing"},
		{name: "eat", input: "eat", want: "eating"},
		{name: "read", input: "read", want: "reading"},
		{name: "think", input: "think", want: "thinking"},
		{name: "walk", input: "walk", want: "walking"},
		{name: "talk", input: "talk", want: "talking"},
		{name: "open", input: "open", want: "opening"},
		{name: "listen", input: "listen", want: "listening"},
		{name: "visit", input: "visit", want: "visiting"},

		// ie -> ying
		{name: "die", input: "die", want: "dying"},
		{name: "lie", input: "lie", want: "lying"},
		{name: "tie", input: "tie", want: "tying"},

		// ee -> eeing
		{name: "see", input: "see", want: "seeing"},
		{name: "flee", input: "flee", want: "fleeing"},
		{name: "agree", input: "agree", want: "agreeing"},
		{name: "free", input: "free", want: "freeing"},

		// be -> being (special vowel + e case)
		{name: "be", input: "be", want: "being"},

		// Words ending in -c (add k)
		{name: "panic", input: "panic", want: "panicking"},
		{name: "picnic", input: "picnic", want: "picnicking"},
		{name: "traffic", input: "traffic", want: "trafficking"},
		{name: "mimic", input: "mimic", want: "mimicking"},
		{name: "frolic", input: "frolic", want: "frolicking"},

		// Words ending in -ye, -oe (keep e)
		{name: "dye", input: "dye", want: "dyeing"},
		{name: "hoe", input: "hoe", want: "hoeing"},
		{name: "toe", input: "toe", want: "toeing"},

		// Words ending in -nge/-inge (keep e)
		{name: "singe", input: "singe", want: "singeing"},

		// Case preservation
		{name: "RUN uppercase", input: "RUN", want: "RUNNING"},
		{name: "Run titlecase", input: "Run", want: "Running"},
		{name: "MAKE uppercase", input: "MAKE", want: "MAKING"},
		{name: "Make titlecase", input: "Make", want: "Making"},
		{name: "DIE uppercase", input: "DIE", want: "DYING"},
		{name: "PANIC uppercase", input: "PANIC", want: "PANICKING"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.PresentParticiple(tt.input)
			if got != tt.want {
				t.Errorf("PresentParticiple(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestSingular(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		// Empty string
		{name: "empty", input: "", want: ""},

		// Regular plurals - remove s
		{name: "cats", input: "cats", want: "cat"},
		{name: "dogs", input: "dogs", want: "dog"},
		{name: "books", input: "books", want: "book"},

		// Words ending in -es after sibilants
		{name: "buses", input: "buses", want: "bus"},
		{name: "classes", input: "classes", want: "class"},
		{name: "bushes", input: "bushes", want: "bush"},
		{name: "churches", input: "churches", want: "church"},
		{name: "boxes", input: "boxes", want: "box"},
		{name: "buzzes", input: "buzzes", want: "buzz"},

		// -ies -> -y (consonant + ies)
		{name: "cities", input: "cities", want: "city"},
		{name: "babies", input: "babies", want: "baby"},
		{name: "flies", input: "flies", want: "fly"},

		// -ys (vowel + y) - just remove s
		{name: "boys", input: "boys", want: "boy"},
		{name: "days", input: "days", want: "day"},
		{name: "keys", input: "keys", want: "key"},

		// Words ending in -ves -> -f or -fe
		{name: "knives", input: "knives", want: "knife"},
		{name: "wives", input: "wives", want: "wife"},
		{name: "lives", input: "lives", want: "life"},
		{name: "leaves", input: "leaves", want: "leaf"},
		{name: "wolves", input: "wolves", want: "wolf"},
		{name: "calves", input: "calves", want: "calf"},
		{name: "halves", input: "halves", want: "half"},

		// Words ending in -oes -> -o
		{name: "heroes", input: "heroes", want: "hero"},
		{name: "potatoes", input: "potatoes", want: "potato"},
		{name: "tomatoes", input: "tomatoes", want: "tomato"},
		{name: "echoes", input: "echoes", want: "echo"},

		// Words ending in -os (just remove s)
		{name: "radios", input: "radios", want: "radio"},
		{name: "studios", input: "studios", want: "studio"},
		{name: "zoos", input: "zoos", want: "zoo"},
		{name: "pianos", input: "pianos", want: "piano"},
		{name: "photos", input: "photos", want: "photo"},

		// Irregular plurals
		{name: "children", input: "children", want: "child"},
		{name: "feet", input: "feet", want: "foot"},
		{name: "teeth", input: "teeth", want: "tooth"},
		{name: "mice", input: "mice", want: "mouse"},
		{name: "women", input: "women", want: "woman"},
		{name: "men", input: "men", want: "man"},
		{name: "people", input: "people", want: "person"},
		{name: "oxen", input: "oxen", want: "ox"},
		{name: "geese", input: "geese", want: "goose"},
		{name: "lice", input: "lice", want: "louse"},
		{name: "dice", input: "dice", want: "die"},

		// Latin/Greek plurals
		{name: "analyses", input: "analyses", want: "analysis"},
		{name: "crises", input: "crises", want: "crisis"},
		{name: "theses", input: "theses", want: "thesis"},
		{name: "cacti", input: "cacti", want: "cactus"},
		{name: "fungi", input: "fungi", want: "fungus"},
		{name: "nuclei", input: "nuclei", want: "nucleus"},
		{name: "bacteria", input: "bacteria", want: "bacterium"},
		{name: "data", input: "data", want: "datum"},
		{name: "media", input: "media", want: "medium"},
		{name: "appendices", input: "appendices", want: "appendix"},
		{name: "indices", input: "indices", want: "index"},
		{name: "criteria", input: "criteria", want: "criterion"},
		{name: "phenomena", input: "phenomena", want: "phenomenon"},

		// Unchanged plurals
		{name: "sheep", input: "sheep", want: "sheep"},
		{name: "deer", input: "deer", want: "deer"},
		{name: "fish", input: "fish", want: "fish"},
		{name: "species", input: "species", want: "species"},
		{name: "series", input: "series", want: "series"},
		{name: "aircraft", input: "aircraft", want: "aircraft"},
		{name: "moose", input: "moose", want: "moose"},

		// Words ending in -men -> -man
		{name: "firemen", input: "firemen", want: "fireman"},
		{name: "policemen", input: "policemen", want: "policeman"},
		{name: "spokesmen", input: "spokesmen", want: "spokesman"},

		// Nationalities ending in -ese (unchanged)
		{name: "Chinese", input: "Chinese", want: "Chinese"},
		{name: "Japanese", input: "Japanese", want: "Japanese"},
		{name: "Portuguese", input: "Portuguese", want: "Portuguese"},

		// Case preservation
		{name: "CATS uppercase", input: "CATS", want: "CAT"},
		{name: "Cats titlecase", input: "Cats", want: "Cat"},
		{name: "Children titlecase", input: "Children", want: "Child"},
		{name: "CHILDREN uppercase", input: "CHILDREN", want: "CHILD"},
		{name: "BOXES uppercase", input: "BOXES", want: "BOX"},
		{name: "Cities titlecase", input: "Cities", want: "City"},
		{name: "MICE uppercase", input: "MICE", want: "MOUSE"},

		// Already singular (should return unchanged)
		{name: "already singular cat", input: "cat", want: "cat"},
		{name: "already singular class", input: "class", want: "class"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Singular(tt.input)
			if got != tt.want {
				t.Errorf("Singular(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

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
			if got != tt.want {
				t.Errorf("Plural(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
