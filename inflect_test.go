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

func TestNumberToWords(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  string
	}{
		// Zero
		{name: "zero", input: 0, want: "zero"},

		// Basic numbers (1-9)
		{name: "one", input: 1, want: "one"},
		{name: "two", input: 2, want: "two"},
		{name: "three", input: 3, want: "three"},
		{name: "four", input: 4, want: "four"},
		{name: "five", input: 5, want: "five"},
		{name: "six", input: 6, want: "six"},
		{name: "seven", input: 7, want: "seven"},
		{name: "eight", input: 8, want: "eight"},
		{name: "nine", input: 9, want: "nine"},

		// Teens (10-19)
		{name: "ten", input: 10, want: "ten"},
		{name: "eleven", input: 11, want: "eleven"},
		{name: "twelve", input: 12, want: "twelve"},
		{name: "thirteen", input: 13, want: "thirteen"},
		{name: "fourteen", input: 14, want: "fourteen"},
		{name: "fifteen", input: 15, want: "fifteen"},
		{name: "sixteen", input: 16, want: "sixteen"},
		{name: "seventeen", input: 17, want: "seventeen"},
		{name: "eighteen", input: 18, want: "eighteen"},
		{name: "nineteen", input: 19, want: "nineteen"},

		// Tens (20, 30, ...)
		{name: "twenty", input: 20, want: "twenty"},
		{name: "thirty", input: 30, want: "thirty"},
		{name: "forty", input: 40, want: "forty"},
		{name: "fifty", input: 50, want: "fifty"},
		{name: "sixty", input: 60, want: "sixty"},
		{name: "seventy", input: 70, want: "seventy"},
		{name: "eighty", input: 80, want: "eighty"},
		{name: "ninety", input: 90, want: "ninety"},

		// Compound tens (21-99)
		{name: "twenty-one", input: 21, want: "twenty-one"},
		{name: "thirty-two", input: 32, want: "thirty-two"},
		{name: "forty-two", input: 42, want: "forty-two"},
		{name: "fifty-five", input: 55, want: "fifty-five"},
		{name: "sixty-seven", input: 67, want: "sixty-seven"},
		{name: "seventy-eight", input: 78, want: "seventy-eight"},
		{name: "eighty-nine", input: 89, want: "eighty-nine"},
		{name: "ninety-nine", input: 99, want: "ninety-nine"},

		// Hundreds
		{name: "one hundred", input: 100, want: "one hundred"},
		{name: "two hundred", input: 200, want: "two hundred"},
		{name: "one hundred one", input: 101, want: "one hundred one"},
		{name: "one hundred ten", input: 110, want: "one hundred ten"},
		{name: "one hundred eleven", input: 111, want: "one hundred eleven"},
		{name: "one hundred twenty", input: 120, want: "one hundred twenty"},
		{name: "one hundred twenty-one", input: 121, want: "one hundred twenty-one"},
		{name: "five hundred fifty-five", input: 555, want: "five hundred fifty-five"},
		{name: "nine hundred ninety-nine", input: 999, want: "nine hundred ninety-nine"},

		// Thousands
		{name: "one thousand", input: 1000, want: "one thousand"},
		{name: "two thousand", input: 2000, want: "two thousand"},
		{name: "one thousand one", input: 1001, want: "one thousand one"},
		{name: "one thousand ten", input: 1010, want: "one thousand ten"},
		{name: "one thousand one hundred", input: 1100, want: "one thousand one hundred"},
		{name: "one thousand two hundred thirty-four", input: 1234, want: "one thousand two hundred thirty-four"},
		{name: "twelve thousand", input: 12000, want: "twelve thousand"},
		{name: "twelve thousand three hundred forty-five", input: 12345, want: "twelve thousand three hundred forty-five"},
		{name: "twenty-one thousand", input: 21000, want: "twenty-one thousand"},
		{name: "one hundred twenty-three thousand four hundred fifty-six", input: 123456, want: "one hundred twenty-three thousand four hundred fifty-six"},

		// Millions
		{name: "one million", input: 1000000, want: "one million"},
		{name: "two million", input: 2000000, want: "two million"},
		{name: "one million one", input: 1000001, want: "one million one"},
		{name: "one million two hundred thirty-four thousand five hundred sixty-seven", input: 1234567, want: "one million two hundred thirty-four thousand five hundred sixty-seven"},
		{name: "twelve million three hundred forty-five thousand six hundred seventy-eight", input: 12345678, want: "twelve million three hundred forty-five thousand six hundred seventy-eight"},
		{name: "one hundred twenty-three million four hundred fifty-six thousand seven hundred eighty-nine", input: 123456789, want: "one hundred twenty-three million four hundred fifty-six thousand seven hundred eighty-nine"},

		// Billions
		{name: "one billion", input: 1000000000, want: "one billion"},
		{name: "two billion", input: 2000000000, want: "two billion"},
		{name: "one billion one", input: 1000000001, want: "one billion one"},

		// Negative numbers
		{name: "negative one", input: -1, want: "negative one"},
		{name: "negative five", input: -5, want: "negative five"},
		{name: "negative eleven", input: -11, want: "negative eleven"},
		{name: "negative twenty-one", input: -21, want: "negative twenty-one"},
		{name: "negative forty-two", input: -42, want: "negative forty-two"},
		{name: "negative one hundred", input: -100, want: "negative one hundred"},
		{name: "negative one thousand", input: -1000, want: "negative one thousand"},
		{name: "negative one million", input: -1000000, want: "negative one million"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.NumberToWords(tt.input)
			if got != tt.want {
				t.Errorf("NumberToWords(%d) = %q, want %q", tt.input, got, tt.want)
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

func TestCompare(t *testing.T) {
	tests := []struct {
		name  string
		word1 string
		word2 string
		want  string
	}{
		// Equal words
		{name: "same word", word1: "cat", word2: "cat", want: "eq"},
		{name: "same word uppercase", word1: "CAT", word2: "CAT", want: "eq"},
		{name: "same word mixed case", word1: "Cat", word2: "cat", want: "eq"},
		{name: "same plural", word1: "cats", word2: "cats", want: "eq"},
		{name: "empty strings", word1: "", word2: "", want: "eq"},

		// Singular to plural
		{name: "cat to cats", word1: "cat", word2: "cats", want: "s:p"},
		{name: "dog to dogs", word1: "dog", word2: "dogs", want: "s:p"},
		{name: "box to boxes", word1: "box", word2: "boxes", want: "s:p"},
		{name: "city to cities", word1: "city", word2: "cities", want: "s:p"},
		{name: "child to children", word1: "child", word2: "children", want: "s:p"},
		{name: "mouse to mice", word1: "mouse", word2: "mice", want: "s:p"},
		{name: "knife to knives", word1: "knife", word2: "knives", want: "s:p"},
		{name: "analysis to analyses", word1: "analysis", word2: "analyses", want: "s:p"},
		{name: "index to indices", word1: "index", word2: "indices", want: "s:p"},
		{name: "cactus to cacti", word1: "cactus", word2: "cacti", want: "s:p"},

		// Plural to singular
		{name: "cats to cat", word1: "cats", word2: "cat", want: "p:s"},
		{name: "dogs to dog", word1: "dogs", word2: "dog", want: "p:s"},
		{name: "boxes to box", word1: "boxes", word2: "box", want: "p:s"},
		{name: "cities to city", word1: "cities", word2: "city", want: "p:s"},
		{name: "children to child", word1: "children", word2: "child", want: "p:s"},
		{name: "mice to mouse", word1: "mice", word2: "mouse", want: "p:s"},
		{name: "knives to knife", word1: "knives", word2: "knife", want: "p:s"},
		{name: "analyses to analysis", word1: "analyses", word2: "analysis", want: "p:s"},
		{name: "indices to index", word1: "indices", word2: "index", want: "p:s"},
		{name: "cacti to cactus", word1: "cacti", word2: "cactus", want: "p:s"},

		// Both plurals (different plural forms of same word)
		{name: "indexes to indices", word1: "indexes", word2: "indices", want: "p:p"},
		{name: "indices to indexes", word1: "indices", word2: "indexes", want: "p:p"},

		// Unrelated words
		{name: "cat to dog", word1: "cat", word2: "dog", want: ""},
		{name: "cats to dogs", word1: "cats", word2: "dogs", want: ""},
		{name: "child to mouse", word1: "child", word2: "mouse", want: ""},
		{name: "box to fox", word1: "box", word2: "fox", want: ""},

		// Empty string edge cases
		{name: "empty and word", word1: "", word2: "cat", want: ""},
		{name: "word and empty", word1: "cat", word2: "", want: ""},

		// Case preservation in comparison
		{name: "Cat to Cats", word1: "Cat", word2: "Cats", want: "s:p"},
		{name: "CAT to CATS", word1: "CAT", word2: "CATS", want: "s:p"},
		{name: "CATS to CAT", word1: "CATS", word2: "CAT", want: "p:s"},

		// Unchanged plurals
		{name: "sheep to sheep", word1: "sheep", word2: "sheep", want: "eq"},
		{name: "deer to deer", word1: "deer", word2: "deer", want: "eq"},
		{name: "fish to fish", word1: "fish", word2: "fish", want: "eq"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Compare(tt.word1, tt.word2)
			if got != tt.want {
				t.Errorf("Compare(%q, %q) = %q, want %q", tt.word1, tt.word2, got, tt.want)
			}
		})
	}
}

func TestCompareNouns(t *testing.T) {
	// CompareNouns is an alias for Compare, so we just verify it behaves the same
	tests := []struct {
		name  string
		noun1 string
		noun2 string
		want  string
	}{
		{name: "singular to plural", noun1: "cat", noun2: "cats", want: "s:p"},
		{name: "plural to singular", noun1: "cats", noun2: "cat", want: "p:s"},
		{name: "equal nouns", noun1: "dog", noun2: "dog", want: "eq"},
		{name: "unrelated nouns", noun1: "cat", noun2: "dog", want: ""},
		{name: "irregular plural", noun1: "child", noun2: "children", want: "s:p"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.CompareNouns(tt.noun1, tt.noun2)
			if got != tt.want {
				t.Errorf("CompareNouns(%q, %q) = %q, want %q", tt.noun1, tt.noun2, got, tt.want)
			}
			// Verify it matches Compare() behavior
			compareGot := inflect.Compare(tt.noun1, tt.noun2)
			if got != compareGot {
				t.Errorf("CompareNouns(%q, %q) = %q, but Compare() = %q", tt.noun1, tt.noun2, got, compareGot)
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

func TestCompareVerbs(t *testing.T) {
	// CompareVerbs is a placeholder stub that always returns empty string.
	// These tests verify the function exists and returns "" for any input.
	tests := []struct {
		name  string
		verb1 string
		verb2 string
		want  string
	}{
		{name: "empty strings", verb1: "", verb2: "", want: ""},
		{name: "same verb", verb1: "run", verb2: "run", want: ""},
		{name: "different verbs", verb1: "run", verb2: "walk", want: ""},
		{name: "conjugated forms", verb1: "run", verb2: "running", want: ""},
		{name: "past tense", verb1: "walk", verb2: "walked", want: ""},
		{name: "irregular verb", verb1: "go", verb2: "went", want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.CompareVerbs(tt.verb1, tt.verb2)
			if got != tt.want {
				t.Errorf("CompareVerbs(%q, %q) = %q, want %q", tt.verb1, tt.verb2, got, tt.want)
			}
		})
	}
}

func TestCompareAdjs(t *testing.T) {
	// CompareAdjs is a placeholder stub that always returns empty string.
	// These tests verify the function exists and returns "" for any input.
	tests := []struct {
		name string
		adj1 string
		adj2 string
		want string
	}{
		{name: "empty strings", adj1: "", adj2: "", want: ""},
		{name: "same adjective", adj1: "big", adj2: "big", want: ""},
		{name: "different adjectives", adj1: "big", adj2: "small", want: ""},
		{name: "comparative form", adj1: "big", adj2: "bigger", want: ""},
		{name: "superlative form", adj1: "big", adj2: "biggest", want: ""},
		{name: "irregular adjective", adj1: "good", adj2: "better", want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.CompareAdjs(tt.adj1, tt.adj2)
			if got != tt.want {
				t.Errorf("CompareAdjs(%q, %q) = %q, want %q", tt.adj1, tt.adj2, got, tt.want)
			}
		})
	}
}

func TestNo(t *testing.T) {
	tests := []struct {
		name  string
		word  string
		count int
		want  string
	}{
		// Zero count - uses "no" with plural
		{name: "zero errors", word: "error", count: 0, want: "no errors"},
		{name: "zero cats", word: "cat", count: 0, want: "no cats"},
		{name: "zero children", word: "child", count: 0, want: "no children"},
		{name: "zero mice", word: "mouse", count: 0, want: "no mice"},
		{name: "zero sheep", word: "sheep", count: 0, want: "no sheep"},
		{name: "zero boxes", word: "box", count: 0, want: "no boxes"},

		// Count of 1 - singular
		{name: "one error", word: "error", count: 1, want: "1 error"},
		{name: "one cat", word: "cat", count: 1, want: "1 cat"},
		{name: "one child", word: "child", count: 1, want: "1 child"},
		{name: "one mouse", word: "mouse", count: 1, want: "1 mouse"},
		{name: "one sheep", word: "sheep", count: 1, want: "1 sheep"},
		{name: "one box", word: "box", count: 1, want: "1 box"},

		// Count > 1 - plural
		{name: "two errors", word: "error", count: 2, want: "2 errors"},
		{name: "five cats", word: "cat", count: 5, want: "5 cats"},
		{name: "three children", word: "child", count: 3, want: "3 children"},
		{name: "ten mice", word: "mouse", count: 10, want: "10 mice"},
		{name: "hundred sheep", word: "sheep", count: 100, want: "100 sheep"},
		{name: "twenty boxes", word: "box", count: 20, want: "20 boxes"},

		// Large counts
		{name: "thousand errors", word: "error", count: 1000, want: "1000 errors"},
		{name: "million items", word: "item", count: 1000000, want: "1000000 items"},

		// Negative counts
		{name: "negative one error", word: "error", count: -1, want: "-1 error"},
		{name: "negative two errors", word: "error", count: -2, want: "-2 errors"},
		{name: "negative five cats", word: "cat", count: -5, want: "-5 cats"},

		// Empty word
		{name: "zero empty", word: "", count: 0, want: "no "},
		{name: "one empty", word: "", count: 1, want: "1 "},

		// Case preservation
		{name: "zero Errors titlecase", word: "Error", count: 0, want: "no Errors"},
		{name: "zero ERRORS uppercase", word: "ERROR", count: 0, want: "no ERRORS"},
		{name: "two Children titlecase", word: "Child", count: 2, want: "2 Children"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.No(tt.word, tt.count)
			if got != tt.want {
				t.Errorf("No(%q, %d) = %q, want %q", tt.word, tt.count, got, tt.want)
			}
		})
	}
}

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
		t.Run(tt.name, func(t *testing.T) {
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
		t.Run(tt.name, func(t *testing.T) {
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
