package inflect_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	inflect "github.com/cv/go-inflect"
)

func TestInflect(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		// Basic plural function
		{
			name:     "simple plural",
			input:    "The plural of cat is plural('cat')",
			expected: "The plural of cat is cats",
		},
		{
			name:     "plural with double quotes",
			input:    `The plural of box is plural("box")`,
			expected: "The plural of box is boxes",
		},
		{
			name:     "plural irregular word",
			input:    "One child, many plural('child')",
			expected: "One child, many children",
		},

		// Plural with count
		{
			name:     "plural with count 1",
			input:    "There is plural('error', 1)",
			expected: "There is error",
		},
		{
			name:     "plural with count 3",
			input:    "There are plural('error', 3)",
			expected: "There are errors",
		},
		{
			name:     "plural with count 0",
			input:    "There are plural('item', 0)",
			expected: "There are items",
		},

		// Singular function
		{
			name:     "simple singular",
			input:    "The singular of cats is singular('cats')",
			expected: "The singular of cats is cat",
		},
		{
			name:     "singular irregular",
			input:    "One of the singular('children')",
			expected: "One of the child",
		},

		// Article functions (an/a)
		{
			name:     "an with vowel",
			input:    "I saw an('apple')",
			expected: "I saw an apple",
		},
		{
			name:     "an with consonant",
			input:    "I saw an('banana')",
			expected: "I saw a banana",
		},
		{
			name:     "a function",
			input:    "This is a('umbrella')",
			expected: "This is an umbrella",
		},
		{
			name:     "a with consonant",
			input:    "This is a('cat')",
			expected: "This is a cat",
		},

		// Ordinal function
		{
			name:     "ordinal 1",
			input:    "This is the ordinal(1) item",
			expected: "This is the 1st item",
		},
		{
			name:     "ordinal 2",
			input:    "This is the ordinal(2) place",
			expected: "This is the 2nd place",
		},
		{
			name:     "ordinal 3",
			input:    "The ordinal(3) time",
			expected: "The 3rd time",
		},
		{
			name:     "ordinal 11",
			input:    "The ordinal(11) hour",
			expected: "The 11th hour",
		},
		{
			name:     "ordinal 21",
			input:    "The ordinal(21) century",
			expected: "The 21st century",
		},

		// Num function
		{
			name:     "num simple",
			input:    "There are num(3) items",
			expected: "There are 3 items",
		},
		{
			name:     "num zero",
			input:    "There are num(0) items",
			expected: "There are 0 items",
		},
		{
			name:     "num large",
			input:    "Count: num(1000)",
			expected: "Count: 1000",
		},

		// Combined examples
		{
			name:     "num with plural",
			input:    "There are num(3) plural('error', 3)",
			expected: "There are 3 errors",
		},
		{
			name:     "num with plural singular case",
			input:    "There is num(1) plural('error', 1)",
			expected: "There is 1 error",
		},
		{
			name:     "multiple functions",
			input:    "I have an('apple') and plural('orange', 2)",
			expected: "I have an apple and oranges",
		},
		{
			name:     "ordinal with plural",
			input:    "The ordinal(1) of plural('mouse', 3)",
			expected: "The 1st of mice",
		},

		// Edge cases
		{
			name:     "no functions",
			input:    "Just plain text",
			expected: "Just plain text",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "unknown function preserved",
			input:    "This unknown('test') stays",
			expected: "This unknown('test') stays",
		},
		{
			name:     "function at start",
			input:    "plural('cat') are cute",
			expected: "cats are cute",
		},
		{
			name:     "function at end",
			input:    "I love plural('dog')",
			expected: "I love dogs",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestInflectPlural(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"basic noun", "plural('cat')", "cats"},
		{"irregular noun", "plural('mouse')", "mice"},
		{"with count 0", "plural('dog', 0)", "dogs"},
		{"with count 1", "plural('dog', 1)", "dog"},
		{"with count 2", "plural('dog', 2)", "dogs"},
		{"with count -1", "plural('dog', -1)", "dog"},
		{"s ending", "plural('bus')", "buses"},
		{"y ending", "plural('baby')", "babies"},
		{"unchanged", "plural('sheep')", "sheep"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestInflectSingular(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"basic noun", "singular('cats')", "cat"},
		{"irregular noun", "singular('mice')", "mouse"},
		{"es ending", "singular('boxes')", "box"},
		{"ies ending", "singular('babies')", "baby"},
		{"unchanged", "singular('sheep')", "sheep"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestInflectArticle(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"an with a", "an('apple')", "an apple"},
		{"an with e", "an('elephant')", "an elephant"},
		{"an with consonant", "an('cat')", "a cat"},
		{"a with vowel", "a('orange')", "an orange"},
		{"a with consonant", "a('dog')", "a dog"},
		{"silent h", "an('hour')", "an hour"},
		{"aspirated h", "an('house')", "a house"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestInflectOrdinal(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"1st", "ordinal(1)", "1st"},
		{"2nd", "ordinal(2)", "2nd"},
		{"3rd", "ordinal(3)", "3rd"},
		{"4th", "ordinal(4)", "4th"},
		{"11th", "ordinal(11)", "11th"},
		{"12th", "ordinal(12)", "12th"},
		{"13th", "ordinal(13)", "13th"},
		{"21st", "ordinal(21)", "21st"},
		{"22nd", "ordinal(22)", "22nd"},
		{"23rd", "ordinal(23)", "23rd"},
		{"100th", "ordinal(100)", "100th"},
		{"101st", "ordinal(101)", "101st"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestInflectNum(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"zero", "num(0)", "0"},
		{"positive", "num(42)", "42"},
		{"large", "num(12345)", "12345"},
		{"negative", "num(-5)", "-5"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestPluralNoun(t *testing.T) {
	tests := []struct {
		name     string
		word     string
		count    []int
		expected string
	}{
		// Nominative pronouns
		{name: "I -> We", word: "I", expected: "We"},
		{name: "i -> we (lowercase)", word: "i", expected: "we"},
		{name: "he -> they", word: "he", expected: "they"},
		{name: "she -> they", word: "she", expected: "they"},
		{name: "it -> they (nominative)", word: "it", expected: "they"},
		{name: "He -> They (case preservation)", word: "He", expected: "They"},
		{name: "SHE -> THEY (uppercase)", word: "SHE", expected: "THEY"},

		// Accusative pronouns
		{name: "me -> us", word: "me", expected: "us"},
		{name: "him -> them", word: "him", expected: "them"},
		{name: "her -> them (accusative)", word: "her", expected: "them"},
		{name: "Me -> Us (case preservation)", word: "Me", expected: "Us"},

		// Possessive pronouns
		{name: "my -> our", word: "my", expected: "our"},
		{name: "mine -> ours", word: "mine", expected: "ours"},
		{name: "his -> their", word: "his", expected: "their"},
		{name: "hers -> theirs", word: "hers", expected: "theirs"},
		{name: "its -> their", word: "its", expected: "their"},
		{name: "My -> Our (case preservation)", word: "My", expected: "Our"},

		// Reflexive pronouns
		{name: "myself -> ourselves", word: "myself", expected: "ourselves"},
		{name: "yourself -> yourselves", word: "yourself", expected: "yourselves"},
		{name: "himself -> themselves", word: "himself", expected: "themselves"},
		{name: "herself -> themselves", word: "herself", expected: "themselves"},
		{name: "itself -> themselves", word: "itself", expected: "themselves"},

		// Regular nouns (should delegate to Plural)
		{name: "cat -> cats", word: "cat", expected: "cats"},
		{name: "box -> boxes", word: "box", expected: "boxes"},
		{name: "child -> children", word: "child", expected: "children"},
		{name: "sheep -> sheep", word: "sheep", expected: "sheep"},

		// Count parameter tests
		{name: "cat count=1 singular", word: "cat", count: []int{1}, expected: "cat"},
		{name: "cat count=2 plural", word: "cat", count: []int{2}, expected: "cats"},
		{name: "cat count=0 plural", word: "cat", count: []int{0}, expected: "cats"},
		{name: "cat count=-1 singular", word: "cat", count: []int{-1}, expected: "cat"},
		{name: "I count=1 singular", word: "I", count: []int{1}, expected: "I"},
		{name: "I count=2 plural", word: "I", count: []int{2}, expected: "We"},

		// Whitespace preservation
		{name: "leading space", word: " cat", expected: " cats"},
		{name: "trailing space", word: "cat ", expected: "cats "},
		{name: "both spaces", word: " cat ", expected: " cats "},
		{name: "pronoun with space", word: " I ", expected: " We "},

		// Empty string
		{name: "empty string", word: "", expected: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got string
			if len(tt.count) > 0 {
				got = inflect.PluralNoun(tt.word, tt.count[0])
			} else {
				got = inflect.PluralNoun(tt.word)
			}
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestPluralVerb(t *testing.T) {
	tests := []struct {
		name     string
		word     string
		count    []int
		expected string
	}{
		// Auxiliary verbs
		{name: "is -> are", word: "is", expected: "are"},
		{name: "was -> were", word: "was", expected: "were"},
		{name: "has -> have", word: "has", expected: "have"},
		{name: "does -> do", word: "does", expected: "do"},
		{name: "goes -> go", word: "goes", expected: "go"},

		// Case preservation
		{name: "Is -> Are", word: "Is", expected: "Are"},
		{name: "WAS -> WERE", word: "WAS", expected: "WERE"},
		{name: "Has -> Have", word: "Has", expected: "Have"},

		// Contractions
		{name: "isn't -> aren't", word: "isn't", expected: "aren't"},
		{name: "wasn't -> weren't", word: "wasn't", expected: "weren't"},
		{name: "hasn't -> haven't", word: "hasn't", expected: "haven't"},
		{name: "doesn't -> don't", word: "doesn't", expected: "don't"},
		{name: "Isn't -> Aren't", word: "Isn't", expected: "Aren't"},

		// Modal verbs (unchanged)
		{name: "can unchanged", word: "can", expected: "can"},
		{name: "could unchanged", word: "could", expected: "could"},
		{name: "may unchanged", word: "may", expected: "may"},
		{name: "might unchanged", word: "might", expected: "might"},
		{name: "must unchanged", word: "must", expected: "must"},
		{name: "shall unchanged", word: "shall", expected: "shall"},
		{name: "should unchanged", word: "should", expected: "should"},
		{name: "will unchanged", word: "will", expected: "will"},
		{name: "would unchanged", word: "would", expected: "would"},
		{name: "can't unchanged", word: "can't", expected: "can't"},
		{name: "won't unchanged", word: "won't", expected: "won't"},

		// Regular third person singular verbs
		{name: "runs -> run", word: "runs", expected: "run"},
		{name: "walks -> walk", word: "walks", expected: "walk"},
		{name: "eats -> eat", word: "eats", expected: "eat"},
		{name: "sees -> see", word: "sees", expected: "see"},
		{name: "tries -> try", word: "tries", expected: "try"},
		{name: "flies -> fly", word: "flies", expected: "fly"},
		{name: "watches -> watch", word: "watches", expected: "watch"},
		{name: "pushes -> push", word: "pushes", expected: "push"},
		{name: "fixes -> fix", word: "fixes", expected: "fix"},

		// Count parameter tests
		{name: "is count=1 singular", word: "is", count: []int{1}, expected: "is"},
		{name: "is count=2 plural", word: "is", count: []int{2}, expected: "are"},
		{name: "was count=1 singular", word: "was", count: []int{1}, expected: "was"},
		{name: "was count=0 plural", word: "was", count: []int{0}, expected: "were"},
		{name: "has count=-1 singular", word: "has", count: []int{-1}, expected: "has"},
		{name: "are count=1 -> is", word: "are", count: []int{1}, expected: "is"},
		{name: "were count=1 -> was", word: "were", count: []int{1}, expected: "was"},
		{name: "have count=1 -> has", word: "have", count: []int{1}, expected: "has"},

		// Whitespace preservation
		{name: "leading space", word: " is", expected: " are"},
		{name: "trailing space", word: "was ", expected: "were "},
		{name: "both spaces", word: " has ", expected: " have "},

		// Empty string
		{name: "empty string", word: "", expected: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got string
			if len(tt.count) > 0 {
				got = inflect.PluralVerb(tt.word, tt.count[0])
			} else {
				got = inflect.PluralVerb(tt.word)
			}
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestPluralAdj(t *testing.T) {
	tests := []struct {
		name     string
		word     string
		count    []int
		expected string
	}{
		// Demonstrative adjectives
		{name: "this -> these", word: "this", expected: "these"},
		{name: "that -> those", word: "that", expected: "those"},
		{name: "This -> These", word: "This", expected: "These"},
		{name: "THAT -> THOSE", word: "THAT", expected: "THOSE"},

		// Indefinite articles
		{name: "a -> some", word: "a", expected: "some"},
		{name: "an -> some", word: "an", expected: "some"},
		{name: "A -> Some", word: "A", expected: "Some"},

		// Possessive adjectives
		{name: "my -> our", word: "my", expected: "our"},
		{name: "your -> your (unchanged)", word: "your", expected: "your"},
		{name: "his -> their", word: "his", expected: "their"},
		{name: "her -> their", word: "her", expected: "their"},
		{name: "its -> their", word: "its", expected: "their"},
		{name: "My -> Our", word: "My", expected: "Our"},

		// Regular adjectives (unchanged)
		{name: "big unchanged", word: "big", expected: "big"},
		{name: "small unchanged", word: "small", expected: "small"},
		{name: "happy unchanged", word: "happy", expected: "happy"},

		// Count parameter tests
		{name: "this count=1 singular", word: "this", count: []int{1}, expected: "this"},
		{name: "this count=2 plural", word: "this", count: []int{2}, expected: "these"},
		{name: "that count=0 plural", word: "that", count: []int{0}, expected: "those"},
		{name: "a count=1 singular", word: "a", count: []int{1}, expected: "a"},
		{name: "a count=2 plural", word: "a", count: []int{2}, expected: "some"},
		{name: "these count=1 -> this", word: "these", count: []int{1}, expected: "this"},
		{name: "those count=1 -> that", word: "those", count: []int{1}, expected: "that"},
		{name: "some count=1 -> a", word: "some", count: []int{1}, expected: "a"},

		// Whitespace preservation
		{name: "leading space", word: " this", expected: " these"},
		{name: "trailing space", word: "that ", expected: "those "},
		{name: "both spaces", word: " a ", expected: " some "},

		// Empty string
		{name: "empty string", word: "", expected: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got string
			if len(tt.count) > 0 {
				got = inflect.PluralAdj(tt.word, tt.count[0])
			} else {
				got = inflect.PluralAdj(tt.word)
			}
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestSingularNoun(t *testing.T) {
	// Test with default gender (t = they)
	tests := []struct {
		name     string
		word     string
		count    []int
		expected string
	}{
		// Nominative pronouns
		{name: "we -> I", word: "we", expected: "I"},
		{name: "We -> I (case)", word: "We", expected: "I"},
		{name: "WE -> I (uppercase)", word: "WE", expected: "I"},
		{name: "they -> they (default gender)", word: "they", expected: "they"},

		// Accusative pronouns
		{name: "us -> me", word: "us", expected: "me"},
		{name: "Us -> Me (case)", word: "Us", expected: "Me"},
		{name: "them -> them (default gender)", word: "them", expected: "them"},

		// Possessive pronouns
		{name: "our -> my", word: "our", expected: "my"},
		{name: "ours -> mine", word: "ours", expected: "mine"},
		{name: "Our -> My (case)", word: "Our", expected: "My"},
		{name: "their -> their (default gender)", word: "their", expected: "their"},
		{name: "theirs -> theirs (default gender)", word: "theirs", expected: "theirs"},

		// Reflexive pronouns
		{name: "ourselves -> myself", word: "ourselves", expected: "myself"},
		{name: "yourselves -> yourself", word: "yourselves", expected: "yourself"},
		{name: "themselves -> themself (default gender)", word: "themselves", expected: "themself"},

		// Regular nouns (should delegate to Singular)
		{name: "cats -> cat", word: "cats", expected: "cat"},
		{name: "boxes -> box", word: "boxes", expected: "box"},
		{name: "children -> child", word: "children", expected: "child"},
		{name: "sheep -> sheep", word: "sheep", expected: "sheep"},

		// Count parameter tests
		{name: "cats count=1 singular", word: "cats", count: []int{1}, expected: "cat"},
		{name: "cats count=2 plural", word: "cats", count: []int{2}, expected: "cats"},
		{name: "cats count=0 plural", word: "cats", count: []int{0}, expected: "cats"},
		{name: "we count=1 singular", word: "we", count: []int{1}, expected: "I"},
		{name: "we count=2 plural", word: "we", count: []int{2}, expected: "we"},

		// Whitespace preservation
		{name: "leading space", word: " cats", expected: " cat"},
		{name: "trailing space", word: "boxes ", expected: "box "},
		{name: "both spaces", word: " we ", expected: " I "},

		// Empty string
		{name: "empty string", word: "", expected: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got string
			if len(tt.count) > 0 {
				got = inflect.SingularNoun(tt.word, tt.count[0])
			} else {
				got = inflect.SingularNoun(tt.word)
			}
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestSingularNounWithGender(t *testing.T) {
	tests := []struct {
		name     string
		gender   string
		word     string
		expected string
	}{
		// Masculine gender
		{name: "they -> he (masculine)", gender: "m", word: "they", expected: "he"},
		{name: "them -> him (masculine)", gender: "m", word: "them", expected: "him"},
		{name: "their -> his (masculine)", gender: "m", word: "their", expected: "his"},
		{name: "theirs -> his (masculine)", gender: "m", word: "theirs", expected: "his"},
		{name: "themselves -> himself (masculine)", gender: "m", word: "themselves", expected: "himself"},

		// Feminine gender
		{name: "they -> she (feminine)", gender: "f", word: "they", expected: "she"},
		{name: "them -> her (feminine)", gender: "f", word: "them", expected: "her"},
		{name: "their -> her (feminine)", gender: "f", word: "their", expected: "her"},
		{name: "theirs -> hers (feminine)", gender: "f", word: "theirs", expected: "hers"},
		{name: "themselves -> herself (feminine)", gender: "f", word: "themselves", expected: "herself"},

		// Neuter gender
		{name: "they -> it (neuter)", gender: "n", word: "they", expected: "it"},
		{name: "them -> it (neuter)", gender: "n", word: "them", expected: "it"},
		{name: "their -> its (neuter)", gender: "n", word: "their", expected: "its"},
		{name: "theirs -> its (neuter)", gender: "n", word: "theirs", expected: "its"},
		{name: "themselves -> itself (neuter)", gender: "n", word: "themselves", expected: "itself"},

		// Singular they
		{name: "they -> they (they)", gender: "t", word: "they", expected: "they"},
		{name: "them -> them (they)", gender: "t", word: "them", expected: "them"},
		{name: "their -> their (they)", gender: "t", word: "their", expected: "their"},
		{name: "theirs -> theirs (they)", gender: "t", word: "theirs", expected: "theirs"},
		{name: "themselves -> themself (they)", gender: "t", word: "themselves", expected: "themself"},

		// Non-gendered pronouns shouldn't change with gender
		{name: "we -> I (any gender)", gender: "m", word: "we", expected: "I"},
		{name: "us -> me (any gender)", gender: "f", word: "us", expected: "me"},
		{name: "our -> my (any gender)", gender: "n", word: "our", expected: "my"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Save current gender and set test gender
			originalGender := inflect.GetGender()
			inflect.Gender(tt.gender)
			defer inflect.Gender(originalGender)

			got := inflect.SingularNoun(tt.word)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestPluralVerbWithCustomDef(t *testing.T) {
	// Test DefVerb integration
	inflect.DefVerb("foobar", "feebar")
	defer inflect.UndefVerb("foobar")

	got := inflect.PluralVerb("foobar")
	assert.Equal(t, "feebar", got, "PluralVerb with custom def")
}

func TestPluralAdjWithCustomDef(t *testing.T) {
	// Test DefAdj integration
	inflect.DefAdj("red", "reds")
	defer inflect.UndefAdj("red")

	got := inflect.PluralAdj("red")
	assert.Equal(t, "reds", got, "PluralAdj with custom def")
}

func TestInflectPluralNoun(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "pronoun I", input: "plural_noun('I')", expected: "We"},
		{name: "pronoun me", input: "plural_noun('me')", expected: "us"},
		{name: "pronoun my", input: "plural_noun('my')", expected: "our"},
		{name: "regular noun", input: "plural_noun('cat')", expected: "cats"},
		{name: "with count 1", input: "plural_noun('cat', 1)", expected: "cat"},
		{name: "with count 2", input: "plural_noun('cat', 2)", expected: "cats"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestInflectPluralVerb(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "is -> are", input: "plural_verb('is')", expected: "are"},
		{name: "was -> were", input: "plural_verb('was')", expected: "were"},
		{name: "has -> have", input: "plural_verb('has')", expected: "have"},
		{name: "with count 1", input: "plural_verb('is', 1)", expected: "is"},
		{name: "with count 2", input: "plural_verb('was', 2)", expected: "were"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestInflectPluralAdj(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "this -> these", input: "plural_adj('this')", expected: "these"},
		{name: "that -> those", input: "plural_adj('that')", expected: "those"},
		{name: "a -> some", input: "plural_adj('a')", expected: "some"},
		{name: "with count 1", input: "plural_adj('this', 1)", expected: "this"},
		{name: "with count 2", input: "plural_adj('that', 2)", expected: "those"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestInflectSingularNoun(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "we -> I", input: "singular_noun('we')", expected: "I"},
		{name: "us -> me", input: "singular_noun('us')", expected: "me"},
		{name: "regular noun", input: "singular_noun('cats')", expected: "cat"},
		{name: "with count 1", input: "singular_noun('cats', 1)", expected: "cat"},
		{name: "with count 2", input: "singular_noun('cats', 2)", expected: "cats"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}
