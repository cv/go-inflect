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
		{"negative", "num(-5)", "0"},
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

func BenchmarkInflect(b *testing.B) {
	benchmarks := []struct {
		name  string
		input string
	}{
		{"no_functions", "plain text without any functions"},
		{"single_plural", "I saw plural('cat')"},
		{"plural_with_count", "I saw plural('cat', 3)"},
		{"multiple_functions", "plural_noun('I') saw plural('cat', 2)"},
		{"ordinal", "This is the ordinal(1) item"},
		{"complex", "plural_noun('I') saw plural_adj('this') plural('cat', 5) ordinal(3) time"},
		{"nested_quotes", "an('apple') and an('orange')"},
		// Adverb function
		{"adverb", "He ran adverb('quick')"},
		// Case conversion functions
		{"snake_case", "Variable: snake_case('HelloWorld')"},
		{"camel_case", "Property: camel_case('hello_world')"},
		{"pascal_case", "Class: pascal_case('my_class')"},
		{"kebab_case", "URL: kebab_case('MyPage')"},
		{"humanize", "Label: humanize('employee_salary')"},
		{"case_conversion_multiple", "snake_case('Test') camel_case('Test') pascal_case('test')"},
		// Rails-style functions
		{"tableize", "Table: tableize('Person')"},
		{"foreign_key", "FK: foreign_key('User')"},
		{"typeify", "Type: typeify('users')"},
		{"parameterize", "Slug: parameterize('Hello World!')"},
		{"asciify", "ASCII: asciify('café')"},
		{"rails_multiple", "tableize('User') with foreign_key('Post')"},
		// Number formatting functions
		{"format_number", "Total: format_number(1000000)"},
		{"number_to_words_with_and", "Amount: number_to_words_with_and(123)"},
		{"number_to_words_threshold", "Count: number_to_words_threshold(5, 10)"},
		{"currency_to_words", "Price: currency_to_words(1.50, 'USD')"},
		{"fraction", "Portion: fraction(3, 4)"},
		// Comparison functions
		{"compare", "Result: compare('cat', 'cats')"},
		{"compare_nouns", "Result: compare_nouns('child', 'children')"},
		{"compare_verbs", "Result: compare_verbs('is', 'are')"},
		{"compare_adjs", "Result: compare_adjs('this', 'these')"},
		// Word count function
		{"word_count", "Words: word_count('hello world')"},
		{"word_count_long", "Words: word_count('the quick brown fox jumps over the lazy dog')"},
		// Capitalize and titleize
		{"capitalize", "Title: capitalize('hello')"},
		{"titleize", "Book: titleize('the great gatsby')"},
		// Word ordinals
		{"word_to_ordinal", "The word_to_ordinal('one') place"},
		{"ordinal_to_cardinal", "Number ordinal_to_cardinal('first')"},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for range b.N {
				inflect.Inflect(bm.input)
			}
		})
	}
}

func TestInflectPastTense(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "regular verb walk", input: "past_tense('walk')", expected: "walked"},
		{name: "regular verb play", input: "past_tense('play')", expected: "played"},
		{name: "regular verb try", input: "past_tense('try')", expected: "tried"},
		{name: "regular verb stop", input: "past_tense('stop')", expected: "stopped"},
		{name: "irregular verb go", input: "past_tense('go')", expected: "went"},
		{name: "irregular verb see", input: "past_tense('see')", expected: "saw"},
		{name: "irregular verb take", input: "past_tense('take')", expected: "took"},
		{name: "irregular verb be", input: "past_tense('be')", expected: "was"},
		{name: "in sentence", input: "She past_tense('walk') home", expected: "She walked home"},
		{name: "double quotes", input: `He past_tense("run") fast`, expected: "He ran fast"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestInflectPastParticiple(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "regular verb walk", input: "past_participle('walk')", expected: "walked"},
		{name: "regular verb play", input: "past_participle('play')", expected: "played"},
		{name: "irregular verb take", input: "past_participle('take')", expected: "taken"},
		{name: "irregular verb go", input: "past_participle('go')", expected: "gone"},
		{name: "irregular verb see", input: "past_participle('see')", expected: "seen"},
		{name: "irregular verb write", input: "past_participle('write')", expected: "written"},
		{name: "in sentence", input: "I have past_participle('take') it", expected: "I have taken it"},
		{name: "double quotes", input: `She has past_participle("write") a book`, expected: "She has written a book"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestInflectPresentParticiple(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "regular verb walk", input: "present_participle('walk')", expected: "walking"},
		{name: "regular verb play", input: "present_participle('play')", expected: "playing"},
		{name: "verb ending in e", input: "present_participle('make')", expected: "making"},
		{name: "verb doubling consonant", input: "present_participle('run')", expected: "running"},
		{name: "verb doubling consonant stop", input: "present_participle('stop')", expected: "stopping"},
		{name: "verb ending in ie", input: "present_participle('lie')", expected: "lying"},
		{name: "in sentence", input: "He is present_participle('run')", expected: "He is running"},
		{name: "double quotes", input: `She is present_participle("swim")`, expected: "She is swimming"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestInflectFutureTense(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "regular verb walk", input: "future_tense('walk')", expected: "will walk"},
		{name: "regular verb play", input: "future_tense('play')", expected: "will play"},
		{name: "irregular verb go", input: "future_tense('go')", expected: "will go"},
		{name: "verb be", input: "future_tense('be')", expected: "will be"},
		{name: "in sentence", input: "She future_tense('walk') home", expected: "She will walk home"},
		{name: "double quotes", input: `He future_tense("run") fast`, expected: "He will run fast"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestInflectPossessive(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "singular noun", input: "possessive('cat')", expected: "cat's"},
		{name: "singular noun dog", input: "possessive('dog')", expected: "dog's"},
		{name: "name", input: "possessive('John')", expected: "John's"},
		{name: "name ending in s", input: "possessive('James')", expected: "James's"},
		{name: "plural noun", input: "possessive('cats')", expected: "cats'"},
		{name: "in sentence", input: "The possessive('cat') toy", expected: "The cat's toy"},
		{name: "double quotes", input: `The possessive("dog") bone`, expected: "The dog's bone"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestInflectComparative(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "short adjective big", input: "comparative('big')", expected: "bigger"},
		{name: "short adjective small", input: "comparative('small')", expected: "smaller"},
		{name: "short adjective tall", input: "comparative('tall')", expected: "taller"},
		{name: "adjective ending in y", input: "comparative('happy')", expected: "happier"},
		{name: "adjective ending in e", input: "comparative('large')", expected: "larger"},
		{name: "long adjective", input: "comparative('beautiful')", expected: "more beautiful"},
		{name: "irregular good", input: "comparative('good')", expected: "better"},
		{name: "irregular bad", input: "comparative('bad')", expected: "worse"},
		{name: "in sentence", input: "This is comparative('big')", expected: "This is bigger"},
		{name: "double quotes", input: `She is comparative("smart")`, expected: "She is smarter"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestInflectSuperlative(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "short adjective big", input: "superlative('big')", expected: "biggest"},
		{name: "short adjective small", input: "superlative('small')", expected: "smallest"},
		{name: "short adjective tall", input: "superlative('tall')", expected: "tallest"},
		{name: "adjective ending in y", input: "superlative('happy')", expected: "happiest"},
		{name: "adjective ending in e", input: "superlative('large')", expected: "largest"},
		{name: "long adjective", input: "superlative('beautiful')", expected: "most beautiful"},
		{name: "irregular good", input: "superlative('good')", expected: "best"},
		{name: "irregular bad", input: "superlative('bad')", expected: "worst"},
		{name: "in sentence", input: "This is the superlative('big')", expected: "This is the biggest"},
		{name: "double quotes", input: `She is the superlative("smart")`, expected: "She is the smartest"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestInflectOrdinalWord(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "first", input: "ordinal_word(1)", expected: "first"},
		{name: "second", input: "ordinal_word(2)", expected: "second"},
		{name: "third", input: "ordinal_word(3)", expected: "third"},
		{name: "fourth", input: "ordinal_word(4)", expected: "fourth"},
		{name: "fifth", input: "ordinal_word(5)", expected: "fifth"},
		{name: "eleventh", input: "ordinal_word(11)", expected: "eleventh"},
		{name: "twelfth", input: "ordinal_word(12)", expected: "twelfth"},
		{name: "twentieth", input: "ordinal_word(20)", expected: "twentieth"},
		{name: "twenty-first", input: "ordinal_word(21)", expected: "twenty-first"},
		{name: "hundredth", input: "ordinal_word(100)", expected: "one hundredth"},
		{name: "in sentence", input: "The ordinal_word(1) place", expected: "The first place"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestInflectNumberToWords(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "zero", input: "number_to_words(0)", expected: "zero"},
		{name: "one", input: "number_to_words(1)", expected: "one"},
		{name: "ten", input: "number_to_words(10)", expected: "ten"},
		{name: "twelve", input: "number_to_words(12)", expected: "twelve"},
		{name: "twenty", input: "number_to_words(20)", expected: "twenty"},
		{name: "forty-two", input: "number_to_words(42)", expected: "forty-two"},
		{name: "one hundred", input: "number_to_words(100)", expected: "one hundred"},
		{name: "one hundred twenty-three", input: "number_to_words(123)", expected: "one hundred twenty-three"},
		{name: "one thousand", input: "number_to_words(1000)", expected: "one thousand"},
		{name: "negative", input: "number_to_words(-5)", expected: "negative five"},
		{name: "in sentence", input: "I have number_to_words(42) apples", expected: "I have forty-two apples"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestInflectCountingWord(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "once", input: "counting_word(1)", expected: "once"},
		{name: "twice", input: "counting_word(2)", expected: "twice"},
		{name: "thrice", input: "counting_word(3)", expected: "thrice"},
		{name: "four times", input: "counting_word(4)", expected: "four times"},
		{name: "ten times", input: "counting_word(10)", expected: "ten times"},
		{name: "zero times", input: "counting_word(0)", expected: "zero times"},
		{name: "in sentence", input: "I saw it counting_word(2)", expected: "I saw it twice"},
		{name: "in sentence once", input: "I did it counting_word(1)", expected: "I did it once"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestInflectNo(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "no errors", input: "no('error', 0)", expected: "no errors"},
		{name: "one error", input: "no('error', 1)", expected: "1 error"},
		{name: "two errors", input: "no('error', 2)", expected: "2 errors"},
		{name: "three errors", input: "no('error', 3)", expected: "3 errors"},
		{name: "no items", input: "no('item', 0)", expected: "no items"},
		{name: "one item", input: "no('item', 1)", expected: "1 item"},
		{name: "five items", input: "no('item', 5)", expected: "5 items"},
		{name: "in sentence zero", input: "There are no('error', 0)", expected: "There are no errors"},
		{name: "in sentence nonzero", input: "There are no('error', 3)", expected: "There are 3 errors"},
		{name: "double quotes", input: `Found no("match", 0)`, expected: "Found no matches"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestInflectNewFunctionsEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		// Empty args - should return original
		{name: "past_tense no args", input: "past_tense()", expected: "past_tense()"},
		{name: "past_participle no args", input: "past_participle()", expected: "past_participle()"},
		{name: "present_participle no args", input: "present_participle()", expected: "present_participle()"},
		{name: "possessive no args", input: "possessive()", expected: "possessive()"},
		{name: "comparative no args", input: "comparative()", expected: "comparative()"},
		{name: "superlative no args", input: "superlative()", expected: "superlative()"},
		{name: "ordinal_word no args", input: "ordinal_word()", expected: "ordinal_word()"},
		{name: "number_to_words no args", input: "number_to_words()", expected: "number_to_words()"},
		{name: "counting_word no args", input: "counting_word()", expected: "counting_word()"},
		{name: "no single arg", input: "no('error')", expected: "no('error')"},

		// Invalid numeric args
		{name: "ordinal_word non-numeric", input: "ordinal_word('abc')", expected: "ordinal_word('abc')"},
		{name: "number_to_words non-numeric", input: "number_to_words('abc')", expected: "number_to_words('abc')"},
		{name: "counting_word non-numeric", input: "counting_word('abc')", expected: "counting_word('abc')"},
		{name: "no non-numeric count", input: "no('error', 'abc')", expected: "no('error', 'abc')"},

		// Multiple new functions in one string
		{name: "multiple new functions", input: "The ordinal_word(1) time I past_tense('see') it", expected: "The first time I saw it"},
		{name: "comparative and superlative", input: "This is comparative('good') but that is superlative('good')", expected: "This is better but that is best"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestInflectAdverb(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "quick -> quickly", input: "adverb('quick')", expected: "quickly"},
		{name: "slow -> slowly", input: "adverb('slow')", expected: "slowly"},
		{name: "happy -> happily", input: "adverb('happy')", expected: "happily"},
		{name: "gentle -> gently", input: "adverb('gentle')", expected: "gently"},
		{name: "in sentence", input: "He ran adverb('quick')", expected: "He ran quickly"},
		{name: "double quotes", input: `She spoke adverb("soft")`, expected: "She spoke softly"},
		{name: "no args", input: "adverb()", expected: "adverb()"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestInflectCapitalize(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "hello -> Hello", input: "capitalize('hello')", expected: "Hello"},
		{name: "WORLD stays WORLD", input: "capitalize('WORLD')", expected: "WORLD"},
		{name: "already capitalized", input: "capitalize('Hello')", expected: "Hello"},
		{name: "in sentence", input: "The word is capitalize('test')", expected: "The word is Test"},
		{name: "no args", input: "capitalize()", expected: "capitalize()"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestInflectTitleize(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "hello world -> Hello World", input: "titleize('hello world')", expected: "Hello World"},
		{name: "the quick brown fox", input: "titleize('the quick brown fox')", expected: "The Quick Brown Fox"},
		{name: "in sentence", input: "Title: titleize('my book title')", expected: "Title: My Book Title"},
		{name: "no args", input: "titleize()", expected: "titleize()"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestInflectWordToOrdinal(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "one -> first", input: "word_to_ordinal('one')", expected: "first"},
		{name: "two -> second", input: "word_to_ordinal('two')", expected: "second"},
		{name: "three -> third", input: "word_to_ordinal('three')", expected: "third"},
		{name: "twenty -> twentieth", input: "word_to_ordinal('twenty')", expected: "twentieth"},
		{name: "in sentence", input: "The word_to_ordinal('one') place", expected: "The first place"},
		{name: "no args", input: "word_to_ordinal()", expected: "word_to_ordinal()"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestInflectOrdinalToCardinal(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "first -> one", input: "ordinal_to_cardinal('first')", expected: "one"},
		{name: "second -> two", input: "ordinal_to_cardinal('second')", expected: "two"},
		{name: "third -> three", input: "ordinal_to_cardinal('third')", expected: "three"},
		{name: "twentieth -> twenty", input: "ordinal_to_cardinal('twentieth')", expected: "twenty"},
		{name: "in sentence", input: "The ordinal_to_cardinal('first') item", expected: "The one item"},
		{name: "no args", input: "ordinal_to_cardinal()", expected: "ordinal_to_cardinal()"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestInflectFraction(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "1/4 -> one quarter", input: "fraction(1, 4)", expected: "one quarter"},
		{name: "1/2 -> one half", input: "fraction(1, 2)", expected: "one half"},
		{name: "3/4 -> three quarters", input: "fraction(3, 4)", expected: "three quarters"},
		{name: "2/3 -> two thirds", input: "fraction(2, 3)", expected: "two thirds"},
		{name: "in sentence", input: "I ate fraction(1, 4) of the pie", expected: "I ate one quarter of the pie"},
		{name: "single arg", input: "fraction(1)", expected: "fraction(1)"},
		{name: "no args", input: "fraction()", expected: "fraction()"},
		{name: "non-numeric", input: "fraction('a', 'b')", expected: "fraction('a', 'b')"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestInflectFormatNumber(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "1000 -> 1,000", input: "format_number(1000)", expected: "1,000"},
		{name: "1000000 -> 1,000,000", input: "format_number(1000000)", expected: "1,000,000"},
		{name: "123 -> 123", input: "format_number(123)", expected: "123"},
		{name: "in sentence", input: "The total is format_number(1000000)", expected: "The total is 1,000,000"},
		{name: "no args", input: "format_number()", expected: "format_number()"},
		{name: "non-numeric", input: "format_number('abc')", expected: "format_number('abc')"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestInflectSnakeCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "HelloWorld -> hello_world", input: "snake_case('HelloWorld')", expected: "hello_world"},
		{name: "hello world -> hello_world", input: "snake_case('hello world')", expected: "hello_world"},
		{name: "camelCase -> camel_case", input: "snake_case('camelCase')", expected: "camel_case"},
		{name: "in sentence", input: "Variable: snake_case('MyVariable')", expected: "Variable: my_variable"},
		{name: "no args", input: "snake_case()", expected: "snake_case()"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestInflectCamelCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "hello_world -> helloWorld", input: "camel_case('hello_world')", expected: "helloWorld"},
		{name: "hello world -> helloWorld", input: "camel_case('hello world')", expected: "helloWorld"},
		{name: "HelloWorld -> helloWorld", input: "camel_case('HelloWorld')", expected: "helloWorld"},
		{name: "in sentence", input: "Variable: camel_case('my_variable')", expected: "Variable: myVariable"},
		{name: "no args", input: "camel_case()", expected: "camel_case()"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestInflectPascalCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "hello_world -> HelloWorld", input: "pascal_case('hello_world')", expected: "HelloWorld"},
		{name: "hello world -> HelloWorld", input: "pascal_case('hello world')", expected: "HelloWorld"},
		{name: "camelCase -> CamelCase", input: "pascal_case('camelCase')", expected: "CamelCase"},
		{name: "in sentence", input: "Class: pascal_case('my_class')", expected: "Class: MyClass"},
		{name: "no args", input: "pascal_case()", expected: "pascal_case()"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestInflectKebabCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "HelloWorld -> hello-world", input: "kebab_case('HelloWorld')", expected: "hello-world"},
		{name: "hello world -> hello-world", input: "kebab_case('hello world')", expected: "hello-world"},
		{name: "camelCase -> camel-case", input: "kebab_case('camelCase')", expected: "camel-case"},
		{name: "in sentence", input: "URL: kebab_case('MyPage')", expected: "URL: my-page"},
		{name: "no args", input: "kebab_case()", expected: "kebab_case()"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestInflectHumanize(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "hello_world -> Hello world", input: "humanize('hello_world')", expected: "Hello world"},
		{name: "employee_salary -> Employee salary", input: "humanize('employee_salary')", expected: "Employee salary"},
		{name: "in sentence", input: "Label: humanize('user_name')", expected: "Label: User name"},
		{name: "no args", input: "humanize()", expected: "humanize()"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestInflectTableize(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "Person -> people", input: "tableize('Person')", expected: "people"},
		{name: "UserPost -> user_posts", input: "tableize('UserPost')", expected: "user_posts"},
		{name: "in sentence", input: "Table: tableize('BlogPost')", expected: "Table: blog_posts"},
		{name: "no args", input: "tableize()", expected: "tableize()"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestInflectForeignKey(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "Person -> person_id", input: "foreign_key('Person')", expected: "person_id"},
		{name: "UserPost -> user_post_id", input: "foreign_key('UserPost')", expected: "user_post_id"},
		{name: "in sentence", input: "FK: foreign_key('Message')", expected: "FK: message_id"},
		{name: "no args", input: "foreign_key()", expected: "foreign_key()"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestInflectTypeify(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "user_post -> UserPost", input: "typeify('user_post')", expected: "UserPost"},
		{name: "blog_posts -> BlogPost", input: "typeify('blog_posts')", expected: "BlogPost"},
		{name: "in sentence", input: "Type: typeify('admin_users')", expected: "Type: AdminUser"},
		{name: "no args", input: "typeify()", expected: "typeify()"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestInflectParameterize(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "Hello World -> hello-world", input: "parameterize('Hello World')", expected: "hello-world"},
		{name: "My Blog Post -> my-blog-post", input: "parameterize('My Blog Post')", expected: "my-blog-post"},
		{name: "in sentence", input: "Slug: parameterize('My Article Title')", expected: "Slug: my-article-title"},
		{name: "no args", input: "parameterize()", expected: "parameterize()"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestInflectAsciify(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "café -> cafe", input: "asciify('café')", expected: "cafe"},
		{name: "naïve -> naive", input: "asciify('naïve')", expected: "naive"},
		{name: "in sentence", input: "ASCII: asciify('résumé')", expected: "ASCII: resume"},
		{name: "no args", input: "asciify()", expected: "asciify()"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestInflectNumberToWordsWithAnd(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "123 with and", input: "number_to_words_with_and(123)", expected: "one hundred and twenty-three"},
		{name: "1001 with and", input: "number_to_words_with_and(1001)", expected: "one thousand and one"},
		{name: "in sentence", input: "Total: number_to_words_with_and(101)", expected: "Total: one hundred and one"},
		{name: "no args", input: "number_to_words_with_and()", expected: "number_to_words_with_and()"},
		{name: "non-numeric", input: "number_to_words_with_and('abc')", expected: "number_to_words_with_and('abc')"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestInflectNumberToWordsThreshold(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "5 threshold 10 -> five", input: "number_to_words_threshold(5, 10)", expected: "five"},
		{name: "15 threshold 10 -> 15", input: "number_to_words_threshold(15, 10)", expected: "15"},
		{name: "in sentence", input: "Count: number_to_words_threshold(3, 10)", expected: "Count: three"},
		{name: "single arg", input: "number_to_words_threshold(5)", expected: "number_to_words_threshold(5)"},
		{name: "no args", input: "number_to_words_threshold()", expected: "number_to_words_threshold()"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestInflectCurrencyToWords(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "1.50 USD", input: "currency_to_words(1.50, 'USD')", expected: "one dollar and fifty cents"},
		{name: "100.00 USD", input: "currency_to_words(100.00, 'USD')", expected: "one hundred dollars"},
		{name: "0.50 USD", input: "currency_to_words(0.50, 'USD')", expected: "fifty cents"},
		{name: "1.00 GBP", input: "currency_to_words(1.00, 'GBP')", expected: "one pound"},
		{name: "in sentence", input: "Price: currency_to_words(25.99, 'USD')", expected: "Price: twenty-five dollars and ninety-nine cents"},
		{name: "single arg", input: "currency_to_words(100)", expected: "currency_to_words(100)"},
		{name: "no args", input: "currency_to_words()", expected: "currency_to_words()"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestInflectCompare(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "cat cat -> eq", input: "compare('cat', 'cat')", expected: "eq"},
		{name: "cat cats -> s:p", input: "compare('cat', 'cats')", expected: "s:p"},
		{name: "cats cat -> p:s", input: "compare('cats', 'cat')", expected: "p:s"},
		{name: "cat dog -> empty", input: "compare('cat', 'dog')", expected: ""},
		{name: "single arg", input: "compare('cat')", expected: "compare('cat')"},
		{name: "no args", input: "compare()", expected: "compare()"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestInflectCompareNouns(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "child children -> s:p", input: "compare_nouns('child', 'children')", expected: "s:p"},
		{name: "children child -> p:s", input: "compare_nouns('children', 'child')", expected: "p:s"},
		{name: "single arg", input: "compare_nouns('cat')", expected: "compare_nouns('cat')"},
		{name: "no args", input: "compare_nouns()", expected: "compare_nouns()"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestInflectCompareVerbs(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "is are -> s:p", input: "compare_verbs('is', 'are')", expected: "s:p"},
		{name: "are is -> p:s", input: "compare_verbs('are', 'is')", expected: "p:s"},
		{name: "single arg", input: "compare_verbs('is')", expected: "compare_verbs('is')"},
		{name: "no args", input: "compare_verbs()", expected: "compare_verbs()"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestInflectCompareAdjs(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "this these -> s:p", input: "compare_adjs('this', 'these')", expected: "s:p"},
		{name: "these this -> p:s", input: "compare_adjs('these', 'this')", expected: "p:s"},
		{name: "single arg", input: "compare_adjs('this')", expected: "compare_adjs('this')"},
		{name: "no args", input: "compare_adjs()", expected: "compare_adjs()"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestInflectWordCount(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "hello world -> 2", input: "word_count('hello world')", expected: "2"},
		{name: "one -> 1", input: "word_count('one')", expected: "1"},
		{name: "empty string returns original", input: "word_count('')", expected: "word_count('')"},
		{name: "multiple spaces", input: "word_count('one  two  three')", expected: "3"},
		{name: "in sentence", input: "Words: word_count('hello world')", expected: "Words: 2"},
		{name: "no args", input: "word_count()", expected: "word_count()"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestInflectNewFunctionsEdgeCasesExtended(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		// New empty args cases
		{name: "adverb no args", input: "adverb()", expected: "adverb()"},
		{name: "capitalize no args", input: "capitalize()", expected: "capitalize()"},
		{name: "titleize no args", input: "titleize()", expected: "titleize()"},
		{name: "word_to_ordinal no args", input: "word_to_ordinal()", expected: "word_to_ordinal()"},
		{name: "ordinal_to_cardinal no args", input: "ordinal_to_cardinal()", expected: "ordinal_to_cardinal()"},
		{name: "fraction no args", input: "fraction()", expected: "fraction()"},
		{name: "format_number no args", input: "format_number()", expected: "format_number()"},
		{name: "snake_case no args", input: "snake_case()", expected: "snake_case()"},
		{name: "camel_case no args", input: "camel_case()", expected: "camel_case()"},
		{name: "pascal_case no args", input: "pascal_case()", expected: "pascal_case()"},
		{name: "kebab_case no args", input: "kebab_case()", expected: "kebab_case()"},
		{name: "humanize no args", input: "humanize()", expected: "humanize()"},
		{name: "tableize no args", input: "tableize()", expected: "tableize()"},
		{name: "foreign_key no args", input: "foreign_key()", expected: "foreign_key()"},
		{name: "typeify no args", input: "typeify()", expected: "typeify()"},
		{name: "parameterize no args", input: "parameterize()", expected: "parameterize()"},
		{name: "asciify no args", input: "asciify()", expected: "asciify()"},
		{name: "word_count no args", input: "word_count()", expected: "word_count()"},

		// Multiple new functions in one string
		{name: "multiple case functions", input: "snake_case('HelloWorld') and kebab_case('HelloWorld')", expected: "hello_world and hello-world"},
		{name: "adverb and capitalize", input: "capitalize(adverb('quick'))", expected: "Adverb(quick)"}, // outer function processes inner as literal string
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestInflectJoin(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "three items", input: "join('a', 'b', 'c')", expected: "a, b, and c"},
		{name: "two items", input: "join('a', 'b')", expected: "a and b"},
		{name: "one item", input: "join('a')", expected: "a"},
		{name: "no args", input: "join()", expected: "join()"},
		{name: "in sentence", input: "I like join('apples', 'oranges', 'bananas')", expected: "I like apples, oranges, and bananas"},
		{name: "four items", input: "join('w', 'x', 'y', 'z')", expected: "w, x, y, and z"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestInflectJoinWith(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "or conjunction", input: "join_with('or', 'a', 'b', 'c')", expected: "a, b, or c"},
		{name: "and conjunction", input: "join_with('and', 'a', 'b', 'c')", expected: "a, b, and c"},
		{name: "two items", input: "join_with('or', 'a', 'b')", expected: "a or b"},
		{name: "one item", input: "join_with('or', 'a')", expected: "a"},
		{name: "no args", input: "join_with()", expected: "join_with()"},
		{name: "only conjunction", input: "join_with('or')", expected: "join_with('or')"},
		{name: "in sentence", input: "Choose join_with('or', 'red', 'blue', 'green')", expected: "Choose red, blue, or green"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Inflect(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}
