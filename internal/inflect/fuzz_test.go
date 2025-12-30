package inflect

import (
	"math"
	"testing"
	"unicode/utf8"
)

// Fuzz tests for string-processing functions.
// These test that functions don't panic on arbitrary input.
// Run with: go test -fuzz=FuzzPlural -fuzztime=30s

func FuzzPlural(f *testing.F) {
	// Seed corpus with interesting cases from unit tests
	seeds := []string{
		// Regular plurals
		"cat", "dog", "book",
		// Sibilants (add -es)
		"bus", "class", "bush", "church", "box", "buzz",
		// Consonant + y -> ies
		"city", "baby", "fly",
		// Vowel + y -> ys
		"boy", "day", "key",
		// Words ending in f/fe -> ves
		"knife", "wife", "leaf", "wolf", "roof", "chief",
		// Words ending in o
		"hero", "potato", "tomato", "echo", "radio", "studio", "zoo", "piano", "photo",
		// Irregular plurals
		"child", "foot", "tooth", "mouse", "woman", "man", "person", "ox",
		// Latin/Greek plurals
		"analysis", "crisis", "thesis", "cactus", "fungus", "nucleus",
		"bacterium", "datum", "medium", "appendix", "index",
		// Unchanged plurals
		"sheep", "deer", "fish", "species", "series", "aircraft",
		// Words ending in -man -> -men
		"fireman", "policeman", "spokesman",
		// Nationalities (unchanged)
		"Chinese", "Japanese", "Portuguese",
		// Classical mode words
		"formula", "antenna", "vertebra", "alumna", "larva", "nebula", "nova",
		"octopus", "opus", "corpus", "genus",
		// Proper names ending in s
		"Jones", "Williams", "Hastings", "Ross", "Burns",
		// Edge cases
		"", " ", "  ", "\t", "\n",
		"UPPERCASE", "MixedCase", "lowercase",
		"cat's", "dogs'", "children's",
		"a", "I", "the",
		"123", "test123", "123test",
		"café", "naïve", "日本語",
		"a b c", "one-two-three",
	}
	for _, s := range seeds {
		f.Add(s)
	}

	f.Fuzz(func(_ *testing.T, input string) {
		if !utf8.ValidString(input) {
			return
		}
		// Should not panic
		_ = Plural(input)
	})
}

func FuzzSingular(f *testing.F) {
	seeds := []string{
		// Regular plurals - remove s
		"cats", "dogs", "books",
		// Words ending in -es after sibilants
		"buses", "classes", "bushes", "churches", "boxes", "buzzes",
		// -ies -> -y
		"cities", "babies", "flies",
		// -ys (vowel + y)
		"boys", "days", "keys",
		// Words ending in -ves -> -f or -fe
		"knives", "wives", "lives", "leaves", "wolves", "calves", "halves",
		// Words ending in -oes -> -o
		"heroes", "potatoes", "tomatoes", "echoes",
		// Words ending in -os
		"radios", "studios", "zoos", "pianos", "photos",
		// Irregular plurals
		"children", "feet", "teeth", "mice", "women", "men", "people", "oxen", "geese", "lice", "dice",
		// Latin/Greek plurals
		"analyses", "crises", "theses", "cacti", "fungi", "nuclei",
		"bacteria", "data", "media", "appendices", "indices", "criteria", "phenomena",
		// Unchanged plurals
		"sheep", "deer", "fish", "species", "series", "aircraft", "moose",
		// Words ending in -men -> -man
		"firemen", "policemen", "spokesmen",
		// Nationalities
		"Chinese", "Japanese", "Portuguese",
		// Case preservation
		"CATS", "Cats", "Children", "CHILDREN", "BOXES", "Cities", "MICE",
		// Already singular
		"cat", "class",
		// Edge cases
		"", " ",
	}
	for _, s := range seeds {
		f.Add(s)
	}

	f.Fuzz(func(_ *testing.T, input string) {
		if !utf8.ValidString(input) {
			return
		}
		_ = Singular(input)
	})
}

// Covers: An, A.
func FuzzAn(f *testing.F) {
	seeds := []string{
		// Basic cases
		"cat", "ant", "a", "b",
		// Silent H
		"honest cat", "dishonest cat", "Honolulu sunset",
		// Special pronunciation
		"mpeg", "onetime holiday",
		// Vowels with consonant sounds (U variations)
		"Ugandan person", "Ukrainian person", "Unabomber", "unanimous decision",
		// Abbreviations and acronyms
		"US farmer", "wild PIKACHU appeared", "YAML code block",
		"Core ML function", "JSON code block", "FBI", "URL", "XML",
		// Words that might need forcing
		"ape", "apple", "eagle", "hour", "hero", "historic",
		// Pattern-matching cases
		"euro", "european", "eurozone", "eurocentric",
		"honor", "honorable", "honorary", "honored",
		"heir", "heirloom", "heiress",
		// Edge cases
		"", " ", "a", "an", "the",
		"European", "one", "once", "unicorn",
		"herb", "hotel",
		"11", "8", "18", "80",
	}
	for _, s := range seeds {
		f.Add(s)
	}

	f.Fuzz(func(_ *testing.T, input string) {
		if !utf8.ValidString(input) {
			return
		}
		_ = An(input)
		_ = A(input)
	})
}

func FuzzInflect(f *testing.F) {
	seeds := []string{
		"I saw plural('cat')",
		"plural('cat', 3)",
		"an('apple') and an('orange')",
		"ordinal(1) place",
		"plural_noun('I') saw plural('cat')",
		"", "no functions here",
		"plural('", "plural()", "plural('cat",
		"nested plural('plural('cat')')",
		"plural('cat', 'not a number')",
		"unknown_func('test')",
		// adverb function
		"adverb('quick')",
		"adverb('slow')",
		"He ran adverb('fast')",
		// capitalize and titleize
		"capitalize('hello')",
		"capitalize('WORLD')",
		"titleize('hello world')",
		"titleize('the quick brown fox')",
		// word ordinals
		"word_to_ordinal('one')",
		"word_to_ordinal('twenty')",
		"ordinal_to_cardinal('first')",
		"ordinal_to_cardinal('twentieth')",
		// fraction
		"fraction(1, 4)",
		"fraction(3, 4)",
		"fraction(1, 2)",
		"fraction(2, 3)",
		// format_number
		"format_number(1000)",
		"format_number(1000000)",
		"format_number(123456789)",
		// case conversion functions
		"snake_case('HelloWorld')",
		"snake_case('camelCase')",
		"camel_case('hello_world')",
		"camel_case('HelloWorld')",
		"pascal_case('hello_world')",
		"pascal_case('camelCase')",
		"kebab_case('HelloWorld')",
		"kebab_case('camelCase')",
		"humanize('employee_salary')",
		"humanize('user_name')",
		// Rails-style functions
		"tableize('Person')",
		"tableize('UserPost')",
		"foreign_key('User')",
		"foreign_key('AdminUser')",
		"typeify('users')",
		"typeify('user_posts')",
		"parameterize('Hello World!')",
		"parameterize('My Blog Post')",
		"asciify('café')",
		"asciify('naïve')",
		// number formatting functions
		"number_to_words_with_and(123)",
		"number_to_words_with_and(1001)",
		"number_to_words_threshold(5, 10)",
		"number_to_words_threshold(15, 10)",
		"currency_to_words(1.50, 'USD')",
		"currency_to_words(100.00, 'GBP')",
		"currency_to_words(0.50, 'EUR')",
		// compare functions
		"compare('cat', 'cats')",
		"compare('cats', 'cat')",
		"compare('cat', 'cat')",
		"compare('cat', 'dog')",
		"compare_nouns('child', 'children')",
		"compare_nouns('mouse', 'mice')",
		"compare_verbs('is', 'are')",
		"compare_verbs('was', 'were')",
		"compare_adjs('this', 'these')",
		"compare_adjs('that', 'those')",
		// word_count function
		"word_count('hello world')",
		"word_count('one two three four five')",
		"word_count('')",
		"word_count('single')",
	}
	for _, s := range seeds {
		f.Add(s)
	}

	f.Fuzz(func(_ *testing.T, input string) {
		if !utf8.ValidString(input) {
			return
		}
		_ = Inflect(input)
	})
}

// Covers: NumberToWords, NumberToWordsWithAnd.
func FuzzNumberToWords(f *testing.F) {
	seeds := []int{
		// Zero
		0,
		// Basic numbers (1-9)
		1, 2, 3, 4, 5, 6, 7, 8, 9,
		// Teens (10-19)
		10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
		// Tens
		20, 30, 40, 50, 60, 70, 80, 90,
		// Compound tens
		21, 32, 42, 55, 67, 78, 89, 99,
		// Hundreds
		100, 101, 110, 111, 120, 121, 200, 555, 999,
		// Thousands
		1000, 1001, 1010, 1100, 1234, 12000, 21000, 12345, 123456,
		// Millions
		1000000, 1000001, 1234567, 12345678, 123456789,
		// Billions
		1000000000, 1000000001, 1234567890,
		// Negative numbers
		-1, -5, -11, -21, -42, -100, -1000, -1000000,
		// Edge cases
		2147483647, -2147483648, -999999999,
	}
	for _, n := range seeds {
		f.Add(n)
	}

	f.Fuzz(func(_ *testing.T, input int) {
		_ = NumberToWords(input)
		_ = NumberToWordsWithAnd(input)
	})
}

// Covers: Ordinal, OrdinalWord.
func FuzzOrdinal(f *testing.F) {
	seeds := []int{
		// Basic cases (1-10)
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		// Teens (special cases: 11th, 12th, 13th)
		11, 12, 13, 14, 19,
		// 21st, 22nd, 23rd pattern
		21, 22, 23, 24,
		// Hundreds
		100, 101, 102, 103, 111, 112, 113, 121, 122, 123,
		// Thousands
		1000, 1001, 1011, 1021,
		// Negative numbers
		-1, -2, -3, -11, -12, -13, -21,
		// Large numbers
		12345, 1000000, 1000000000,
	}
	for _, n := range seeds {
		f.Add(n)
	}

	f.Fuzz(func(_ *testing.T, input int) {
		_ = Ordinal(input)
		_ = OrdinalWord(input)
	})
}

func FuzzPresentParticiple(f *testing.F) {
	seeds := []string{
		// Single letter verbs
		"a",
		// Already ending in -ing
		"running", "sing",
		// Double consonant (CVC pattern)
		"run", "sit", "hit", "cut", "stop", "drop", "plan", "skip",
		"begin", "occur", "prefer", "admit", "commit", "regret",
		// Drop silent e
		"make", "take", "come", "give", "have", "write", "live", "move", "hope", "dance",
		// Just add -ing
		"play", "stay", "enjoy", "show", "follow", "fix", "mix", "go", "do",
		"eat", "read", "think", "walk", "talk", "open", "listen", "visit",
		// ie -> ying
		"die", "lie", "tie",
		// ee -> eeing
		"see", "flee", "agree", "free",
		// be -> being
		"be",
		// Words ending in -c (add k)
		"panic", "picnic", "traffic", "mimic", "frolic",
		// Words ending in -ye, -oe
		"dye", "hoe", "toe",
		// Words ending in -nge/-inge
		"singe",
		// Case preservation
		"RUN", "Run", "MAKE", "Make", "DIE", "PANIC",
		// Edge cases
		"", " ", "123",
	}
	for _, s := range seeds {
		f.Add(s)
	}

	f.Fuzz(func(_ *testing.T, input string) {
		if !utf8.ValidString(input) {
			return
		}
		_ = PresentParticiple(input)
	})
}

func FuzzPastTense(f *testing.F) {
	seeds := []string{
		// Regular verbs: add -ed
		"walk", "talk", "work", "play", "stay", "enjoy", "destroy",
		"help", "start", "finish", "watch", "wash", "push", "pull",
		"open", "close", "need", "want", "ask", "answer", "clean", "cook", "look",
		// Verbs ending in -e: add -d
		"love", "like", "live", "move", "change", "create", "use", "hope",
		"smile", "dance", "arrive", "decide", "believe", "receive",
		// Consonant + y: change y to -ied
		"try", "cry", "carry", "study", "hurry", "worry", "marry",
		"copy", "apply", "reply", "supply", "occupy", "deny", "rely",
		// CVC pattern: double final consonant
		"stop", "drop", "shop", "plan", "rob", "rub", "hug", "jog",
		"grab", "trip", "slip", "step", "beg", "nod", "chat",
		// Don't double w, x, y
		"show", "fix", "box", "mix",
		// Irregular verbs
		"go", "be", "have", "do", "say", "make", "get", "see", "come", "take",
		"know", "think", "find", "give", "tell", "become", "leave", "put", "keep", "let",
		"begin", "run", "write", "read", "bring", "buy", "catch", "teach", "fight",
		"build", "send", "spend", "lose", "feel", "meet", "sit", "stand", "hear", "hold",
		"speak", "break", "choose", "grow", "throw", "blow", "fly", "draw", "drive", "ride",
		"rise", "hide", "eat", "fall", "swim", "sing", "ring", "drink", "sink", "win",
		"hit", "cut", "shut", "set", "hurt", "cost", "sleep", "wake", "wear", "tear",
		"bear", "swear", "steal", "freeze", "forget", "forgive", "bite", "shake",
		"mistake", "undertake", "shine", "lie", "lay", "pay", "mean",
		// Case preservation
		"WALK", "Walk", "GO", "Go", "TRY", "Try",
		// Edge cases
		"", " ", "123",
	}
	for _, s := range seeds {
		f.Add(s)
	}

	f.Fuzz(func(_ *testing.T, input string) {
		if !utf8.ValidString(input) {
			return
		}
		_ = PastTense(input)
	})
}

func FuzzPastParticiple(f *testing.F) {
	seeds := []string{
		// Regular verbs (-ed)
		"walk", "talk", "play", "stay", "work", "help", "ask", "call", "open", "listen",
		// Verbs ending in -e
		"like", "love", "dance", "hope", "use", "close",
		// Verbs ending in consonant + y
		"try", "cry", "study", "carry", "worry",
		// Verbs ending in vowel + y
		"enjoy", "delay",
		// CVC pattern - double consonant
		"stop", "drop", "plan", "skip", "admit", "occur", "prefer", "regret",
		// Don't double w, x, y
		"fix", "mix", "show",
		// Verbs ending in -c
		"panic", "picnic", "traffic",
		// Irregular verbs
		"go", "be", "have", "do", "say", "get", "make", "know", "think", "take",
		"see", "come", "give", "find", "tell", "write", "run", "eat", "drink", "sing",
		"swim", "begin", "break", "choose", "speak", "steal", "forget", "drive", "ride",
		"hide", "bite", "fly", "grow", "throw", "draw", "fall", "buy", "bring", "catch",
		"teach", "fight", "seek", "feel", "keep", "sleep", "leave", "meet", "read", "lead",
		"sit", "stand", "lose", "win", "put", "cut", "hit", "let", "set", "shut", "hurt",
		"cost", "build", "send", "spend", "lend", "bend",
		// Case preservation
		"WALK", "Walk", "GO", "Go",
		// Edge cases
		"", " ", "123",
	}
	for _, s := range seeds {
		f.Add(s)
	}

	f.Fuzz(func(_ *testing.T, input string) {
		if !utf8.ValidString(input) {
			return
		}
		_ = PastParticiple(input)
	})
}

// Covers: Comparative, Superlative.
func FuzzComparative(f *testing.F) {
	seeds := []string{
		// Irregular forms
		"good", "bad", "far", "little", "much", "many", "old",
		// One-syllable adjectives: add -er
		"tall", "short", "fast", "slow", "young", "long", "strong", "weak",
		"cheap", "deep", "high", "low", "new", "poor", "rich", "warm", "cold",
		"dark", "light", "hard", "soft", "clean", "loud",
		// One-syllable ending in -e: add -r
		"large", "wide", "close", "late", "nice", "safe", "wise", "rude", "rare", "pale", "fine", "cute", "pure",
		// CVC pattern: double final consonant
		"big", "hot", "thin", "fat", "wet", "sad", "red", "dim", "fit",
		// Consonant + y: change y to -ier
		"happy", "easy", "busy", "funny", "pretty", "heavy", "dirty", "angry", "crazy", "lazy", "tiny", "ugly", "early", "noisy",
		// Two-syllable adjectives that take -er
		"simple", "gentle", "narrow", "shallow", "quiet", "clever",
		// Long adjectives: use "more"
		"beautiful", "dangerous", "expensive", "important", "interesting", "comfortable",
		"difficult", "intelligent", "wonderful", "terrible", "horrible", "incredible",
		"successful", "popular", "famous", "nervous",
		// Case preservation
		"BIG", "Big", "GOOD", "Good", "BEAUTIFUL", "Beautiful",
		// Edge cases
		"", " ", "123", "a",
	}
	for _, s := range seeds {
		f.Add(s)
	}

	f.Fuzz(func(_ *testing.T, input string) {
		if !utf8.ValidString(input) {
			return
		}
		_ = Comparative(input)
		_ = Superlative(input)
	})
}

// Covers: CamelCase, PascalCase, SnakeCase, KebabCase, Dasherize, Underscore.
func FuzzCaseConversion(f *testing.F) {
	seeds := []string{
		"camelCase", "PascalCase", "snake_case", "kebab-case",
		"XMLHttpRequest", "getHTTPResponse", "IOError",
		"", " ", "a", "ABC", "abc",
		"hello world", "HELLO_WORLD",
		"123test", "test123",
	}
	for _, s := range seeds {
		f.Add(s)
	}

	f.Fuzz(func(_ *testing.T, input string) {
		if !utf8.ValidString(input) {
			return
		}
		_ = CamelCase(input)
		_ = PascalCase(input)
		_ = SnakeCase(input)
		_ = KebabCase(input)
		_ = Dasherize(input)
		_ = Underscore(input)
	})
}

// Covers: Join, JoinWithConj, JoinNoOxford.
func FuzzJoin(f *testing.F) {
	// Fuzz with varying number of items encoded as newline-separated
	seeds := []string{
		"",
		"one",
		"one\ntwo",
		"one\ntwo\nthree",
		"apple\nbanana\ncherry\ndate",
		"a, b\nc, d\ne, f",
		"\n\n\n",
	}
	for _, s := range seeds {
		f.Add(s)
	}

	f.Fuzz(func(_ *testing.T, input string) {
		if !utf8.ValidString(input) {
			return
		}
		// Split on newlines to get items
		var items []string
		start := 0
		for i := range len(input) {
			if input[i] == '\n' {
				items = append(items, input[start:i])
				start = i + 1
			}
		}
		if start < len(input) {
			items = append(items, input[start:])
		}

		_ = Join(items)
		_ = JoinWithConj(items, "or")
		_ = JoinNoOxford(items)
	})
}

func FuzzPossessive(f *testing.F) {
	seeds := []string{
		"cat", "cats", "child", "children",
		"James", "boss", "class",
		"", " ", "s", "ss", "'s",
	}
	for _, s := range seeds {
		f.Add(s)
	}

	f.Fuzz(func(_ *testing.T, input string) {
		if !utf8.ValidString(input) {
			return
		}
		_ = Possessive(input)
	})
}

// Covers: Compare, CompareNouns, CompareVerbs, CompareAdjs.
func FuzzCompare(f *testing.F) {
	seeds := []struct {
		a, b string
	}{
		{"cat", "cats"},
		{"mouse", "mice"},
		{"child", "children"},
		{"", ""},
		{"cat", "dog"},
		{"run", "runs"},
	}
	for _, s := range seeds {
		f.Add(s.a, s.b)
	}

	f.Fuzz(func(_ *testing.T, a, b string) {
		if !utf8.ValidString(a) || !utf8.ValidString(b) {
			return
		}
		_ = Compare(a, b)
		_ = CompareNouns(a, b)
		_ = CompareVerbs(a, b)
		_ = CompareAdjs(a, b)
	})
}

// Covers: WordToOrdinal, OrdinalToCardinal.
func FuzzWordToOrdinal(f *testing.F) {
	seeds := []string{
		// Basic word numbers
		"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten",
		"eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen",
		// Tens
		"twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety",
		// Compound numbers
		"twenty-one", "thirty-two", "forty-three", "ninety-nine",
		// Special cases
		"zero",
		// Already ordinals (for OrdinalToCardinal)
		"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth",
		"eleventh", "twelfth", "twentieth", "twenty-first", "zeroth",
		// Case preservation
		"One", "TWO", "Twenty-One", "First", "SECOND", "Twenty-First",
		// Numeric strings
		"1", "2", "3", "11", "21", "100", "1st", "2nd", "3rd", "11th", "21st", "100th",
		// Not ordinals
		"cat", "north", "month", "earth",
		// Edge cases
		"", " ", "invalid",
	}
	for _, s := range seeds {
		f.Add(s)
	}

	f.Fuzz(func(_ *testing.T, input string) {
		if !utf8.ValidString(input) {
			return
		}
		_ = WordToOrdinal(input)
		_ = OrdinalToCardinal(input)
	})
}

func FuzzCurrency(f *testing.F) {
	f.Add(0.0, "USD")
	f.Add(1.0, "USD")
	f.Add(1.50, "USD")
	f.Add(123.45, "USD")
	f.Add(-50.00, "USD")
	f.Add(1000000.99, "USD")
	f.Add(0.01, "GBP")
	f.Add(100.00, "EUR")
	f.Add(1.00, "XXX") // unknown currency

	f.Fuzz(func(_ *testing.T, amount float64, currency string) {
		if !utf8.ValidString(currency) {
			return
		}
		_ = CurrencyToWords(amount, currency)
	})
}

// Covers: FractionToWords, FractionToWordsWithFourths.
func FuzzFraction(f *testing.F) {
	seeds := []struct {
		num, den int
	}{
		{1, 2}, {1, 3}, {1, 4}, {2, 3}, {3, 4},
		{0, 1}, {1, 1}, {5, 5},
		{1, 0}, {0, 0}, // edge cases
		{-1, 2}, {1, -2}, {-1, -2},
		{100, 3}, {1, 1000},
	}
	for _, s := range seeds {
		f.Add(s.num, s.den)
	}

	f.Fuzz(func(_ *testing.T, num, den int) {
		_ = FractionToWords(num, den)
		_ = FractionToWordsWithFourths(num, den)
	})
}

func FuzzAdverb(f *testing.F) {
	seeds := []string{
		"quick", "slow", "happy", "easy",
		"good", "bad", "fast", "hard",
		"", " ", "123",
	}
	for _, s := range seeds {
		f.Add(s)
	}

	f.Fuzz(func(_ *testing.T, input string) {
		if !utf8.ValidString(input) {
			return
		}
		_ = Adverb(input)
	})
}

func FuzzPluralNoun(f *testing.F) {
	seeds := []string{
		// Regular nouns
		"cat", "dog", "book",
		// Sibilants
		"box", "bus", "church", "class", "bush", "buzz",
		// Words ending in o
		"potato", "hero", "radio", "photo",
		// Pronouns (comprehensive from pronoun handling)
		"I", "me", "my", "mine", "myself",
		"you", "your", "yours", "yourself",
		"he", "she", "it", "they",
		"him", "her", "them",
		"his", "hers", "its", "their", "theirs",
		"we", "us", "our", "ours", "ourselves",
		"who", "whom", "whose", "whoever", "whomever",
		// Irregular plurals
		"child", "mouse", "foot", "tooth", "woman", "man", "person", "ox",
		"analysis", "cactus", "datum", "phenomenon", "criterion",
		"leaf", "wolf", "knife", "wife",
		// Unchanged plurals
		"sheep", "deer", "fish", "species", "series", "aircraft",
		// Edge cases
		"", " ", "  ", "\t", "\n",
		"UPPERCASE", "MixedCase", "lowercase",
		"123", "test123",
		"café", "naïve",
	}
	for _, s := range seeds {
		f.Add(s)
	}

	f.Fuzz(func(_ *testing.T, input string) {
		if !utf8.ValidString(input) {
			return
		}
		_ = PluralNoun(input)
	})
}

func FuzzPluralVerb(f *testing.F) {
	seeds := []string{
		// Auxiliary verbs (from verb conjugation tests)
		"is", "was", "has", "does", "am", "are", "were", "have", "do",
		// Contractions
		"isn't", "wasn't", "hasn't", "doesn't", "aren't", "weren't", "haven't", "don't",
		// Modal verbs (unchanged)
		"can", "could", "may", "might", "must", "shall", "should", "will", "would",
		// Regular verbs (third person singular -> base form)
		"runs", "walks", "goes", "sees", "flies", "tries",
		"passes", "pushes", "watches", "fixes", "buzzes",
		// Irregular past tense verbs
		"went", "came", "took", "gave", "saw", "knew", "thought",
		// Base form verbs
		"run", "walk", "go", "see", "fly", "try", "be", "have", "do",
		// Edge cases
		"", " ", "  ", "\t", "\n",
		"UPPERCASE", "MixedCase", "lowercase",
		"123", "test123",
	}
	for _, s := range seeds {
		f.Add(s)
	}

	f.Fuzz(func(_ *testing.T, input string) {
		if !utf8.ValidString(input) {
			return
		}
		_ = PluralVerb(input)
	})
}

func FuzzPluralAdj(f *testing.F) {
	seeds := []string{
		// Demonstrative adjectives
		"this", "that", "these", "those",
		// Indefinite articles
		"a", "an", "some",
		// Possessive adjectives
		"my", "your", "his", "her", "its", "our", "their",
		// Regular adjectives (unchanged)
		"big", "small", "beautiful", "happy", "red", "blue",
		// Edge cases
		"", " ", "  ", "\t", "\n",
		"UPPERCASE", "MixedCase", "lowercase",
		"123", "test123",
	}
	for _, s := range seeds {
		f.Add(s)
	}

	f.Fuzz(func(_ *testing.T, input string) {
		if !utf8.ValidString(input) {
			return
		}
		_ = PluralAdj(input)
	})
}

func FuzzSingularNoun(f *testing.F) {
	seeds := []string{
		// Plural nouns to singularize
		"cats", "dogs", "books",
		"boxes", "buses", "churches", "classes", "bushes", "buzzes",
		"cities", "babies", "flies",
		"boys", "days", "keys",
		"knives", "wives", "lives", "leaves", "wolves", "calves", "halves",
		"heroes", "potatoes", "tomatoes", "echoes",
		"radios", "studios", "zoos", "pianos", "photos",
		// Irregular plurals
		"children", "feet", "teeth", "mice", "women", "men", "people", "oxen", "geese", "lice", "dice",
		"analyses", "crises", "theses", "cacti", "fungi", "nuclei",
		"bacteria", "data", "media", "appendices", "indices", "criteria", "phenomena",
		// Pronouns (plural)
		"we", "us", "our", "ours", "ourselves",
		"they", "them", "their", "theirs", "themselves",
		// Pronouns (singular)
		"I", "me", "my", "mine", "myself",
		"he", "she", "it", "him", "her",
		// Unchanged plurals
		"sheep", "deer", "fish", "species", "series", "aircraft", "moose",
		// Words ending in -men
		"firemen", "policemen", "spokesmen",
		// Edge cases
		"", " ", "  ", "\t", "\n",
		"UPPERCASE", "MixedCase", "lowercase",
		"123", "test123",
		"café", "naïve",
	}
	for _, s := range seeds {
		f.Add(s)
	}

	f.Fuzz(func(_ *testing.T, input string) {
		if !utf8.ValidString(input) {
			return
		}
		_ = SingularNoun(input)
	})
}

func FuzzNo(f *testing.F) {
	// Seeds: word (string), count (int)
	seeds := []struct {
		word  string
		count int
	}{
		// Regular nouns
		{"cat", 0}, {"cat", 1}, {"cat", 2}, {"cat", -1},
		{"error", 0}, {"error", 1}, {"error", 5},
		{"child", 0}, {"child", 1}, {"child", 3},
		// Edge cases
		{"", 0}, {"", 1}, {"", -1},
		{" ", 0}, {"  ", 1},
		// Irregular plurals
		{"mouse", 0}, {"mouse", 1}, {"mouse", 5},
		{"datum", 0}, {"datum", 1}, {"datum", 10},
		// Large numbers
		{"item", 1000000}, {"item", -1000000},
		// Special characters
		{"café", 0}, {"naïve", 2},
	}
	for _, s := range seeds {
		f.Add(s.word, s.count)
	}

	f.Fuzz(func(_ *testing.T, word string, count int) {
		if !utf8.ValidString(word) {
			return
		}
		_ = No(word, count)
	})
}

func FuzzCountingWord(f *testing.F) {
	seeds := []int{
		// Special words
		0, 1, 2, 3,
		// Regular small numbers
		4, 5, 6, 7, 8, 9, 10,
		// Teens
		11, 12, 13, 14, 15,
		// Larger numbers
		20, 21, 50, 99, 100, 101,
		1000, 1000000, 1000000000,
		// Negative numbers
		-1, -2, -3, -10, -100,
		// Edge cases
		2147483647, -2147483648,
	}
	for _, n := range seeds {
		f.Add(n)
	}

	f.Fuzz(func(_ *testing.T, n int) {
		_ = CountingWord(n)
	})
}

func FuzzCountingWordWithOptions(f *testing.F) {
	// Seeds: n (int), useThrice (bool)
	seeds := []struct {
		n         int
		useThrice bool
	}{
		// Special words with useThrice variations
		{1, true}, {1, false},
		{2, true}, {2, false},
		{3, true}, {3, false},
		// Zero
		{0, true}, {0, false},
		// Regular numbers
		{4, true}, {4, false},
		{10, true}, {10, false},
		{100, true}, {100, false},
		// Negative numbers
		{-1, true}, {-1, false},
		{-2, true}, {-2, false},
		{-3, true}, {-3, false},
		{-10, true}, {-10, false},
		// Edge cases
		{2147483647, true}, {-2147483648, false},
	}
	for _, s := range seeds {
		f.Add(s.n, s.useThrice)
	}

	f.Fuzz(func(_ *testing.T, n int, useThrice bool) {
		_ = CountingWordWithOptions(n, useThrice)
	})
}

func FuzzCountingWordThreshold(f *testing.F) {
	// Seeds: n (int), threshold (int)
	seeds := []struct {
		n         int
		threshold int
	}{
		// Below threshold
		{1, 10}, {2, 10}, {3, 10}, {5, 10}, {9, 10},
		// At threshold
		{10, 10}, {100, 100},
		// Above threshold
		{15, 10}, {100, 10}, {1000, 10},
		// Zero
		{0, 10}, {0, 0},
		// Negative numbers
		{-1, 10}, {-5, 10}, {-15, 10},
		// Negative threshold (edge case)
		{5, -10}, {-5, -10},
		// Edge cases
		{2147483647, 10}, {-2147483648, 10},
		{10, 2147483647}, {10, -2147483648},
	}
	for _, s := range seeds {
		f.Add(s.n, s.threshold)
	}

	f.Fuzz(func(_ *testing.T, n, threshold int) {
		_ = CountingWordThreshold(n, threshold)
	})
}

func FuzzCapitalize(f *testing.F) {
	seeds := []string{
		// Empty and whitespace
		"", " ", "  ", "\t", "\n", "\r\n",
		// Single characters
		"a", "A", "1", "!", "é", "日",
		// Mixed case
		"hello", "HELLO", "Hello", "hELLO", "HeLLo",
		// With numbers
		"123", "abc123", "123abc", "a1b2c3",
		// Unicode
		"café", "naïve", "日本語", "Привет", "مرحبا",
		"über", "ÜBER", "Äpfel", "ñoño",
		// Special characters
		"hello world", "hello-world", "hello_world",
		"'quoted'", "\"double\"", "(parens)",
		// Leading whitespace
		" hello", "  world", "\thello",
		// Only whitespace
		"   ", "\t\t", "\n\n",
	}
	for _, s := range seeds {
		f.Add(s)
	}

	f.Fuzz(func(_ *testing.T, input string) {
		if !utf8.ValidString(input) {
			return
		}
		_ = Capitalize(input)
	})
}

func FuzzTitleize(f *testing.F) {
	seeds := []string{
		// Empty and whitespace
		"", " ", "  ", "\t", "\n", "\r\n",
		// Single words
		"hello", "HELLO", "Hello", "hELLO",
		// Multiple words
		"hello world", "HELLO WORLD", "Hello World",
		"the quick brown fox",
		// With hyphens
		"hello-world", "one-two-three", "HELLO-WORLD",
		// Mixed separators
		"hello world-test", "one two-three four",
		// With numbers
		"123", "abc123", "123abc", "test 123 value",
		// Unicode
		"café au lait", "naïve approach", "日本語 テスト",
		"über alles", "ÜBER ALLES", "Äpfel und Birnen",
		// Edge cases
		"a", "A", " a ", "-a-", "a-",
		// Multiple spaces/hyphens
		"hello  world", "hello--world", "  hello  ",
		// Only whitespace
		"   ", "\t\t", "\n\n",
	}
	for _, s := range seeds {
		f.Add(s)
	}

	f.Fuzz(func(_ *testing.T, input string) {
		if !utf8.ValidString(input) {
			return
		}
		_ = Titleize(input)
	})
}

func FuzzTitleCase(f *testing.F) {
	seeds := []string{
		// Empty and whitespace
		"", " ", "  ", "\t", "\n", "\r\n",
		// snake_case
		"hello_world", "one_two_three", "HELLO_WORLD",
		// kebab-case
		"hello-world", "one-two-three", "HELLO-WORLD",
		// Space separated
		"hello world", "one two three",
		// Mixed case inputs
		"camelCase", "PascalCase", "snake_case", "kebab-case",
		// With numbers
		"123", "abc123", "123abc", "test_123_value",
		// Acronyms
		"XMLHttpRequest", "getHTTPResponse", "IOError",
		"HTTP_SERVER", "xml_http_request",
		// Unicode
		"café_au_lait", "naïve-approach", "日本語_テスト",
		// Edge cases
		"a", "A", "_a_", "-a-", "a_", "_a",
		// Multiple separators
		"hello__world", "hello--world", "hello_ _world",
		// Only separators
		"___", "---", "_ _", "- -",
		// Only whitespace
		"   ", "\t\t", "\n\n",
	}
	for _, s := range seeds {
		f.Add(s)
	}

	f.Fuzz(func(_ *testing.T, input string) {
		if !utf8.ValidString(input) {
			return
		}
		_ = TitleCase(input)
	})
}

func FuzzNumberToWordsFloat(f *testing.F) {
	// Seed with uint64 bits representation of interesting float64 values
	// because Go's fuzz testing doesn't natively support float64
	seeds := []uint64{
		// Common values
		0x0000000000000000, // 0.0
		0x3FF0000000000000, // 1.0
		0xBFF0000000000000, // -1.0
		0x4024000000000000, // 10.0
		0x40091EB851EB851F, // 3.14
		0x4005BF0A8B145769, // 2.718
		0x3FE0000000000000, // 0.5
		0x3FB999999999999A, // 0.1
		0x3FF199999999999A, // 1.1
		0x412E848000000000, // 1000000.0
		0xC12E848000000000, // -1000000.0
		// Special values
		0x7FF0000000000000, // +Inf
		0xFFF0000000000000, // -Inf
		0x7FF8000000000000, // NaN
		// Very small values
		0x3E70000000000000, // 1e-8
		0x0010000000000000, // smallest normal positive
		// Very large values
		0x7FEFFFFFFFFFFFFF, // largest finite positive
		0xFFEFFFFFFFFFFFFF, // largest finite negative
		// Subnormal values
		0x0000000000000001, // smallest subnormal positive
	}
	for _, bits := range seeds {
		f.Add(bits)
	}

	f.Fuzz(func(_ *testing.T, bits uint64) {
		n := math.Float64frombits(bits)
		// Skip NaN and Inf as they may cause panics or undefined behavior
		if math.IsNaN(n) || math.IsInf(n, 0) {
			return
		}
		// Skip values that are too large to convert to int
		if n > float64(math.MaxInt64) || n < float64(math.MinInt64) {
			return
		}
		// Should not panic
		_ = NumberToWordsFloat(n)
	})
}

func FuzzNumberToWordsFloatWithDecimal(f *testing.F) {
	// Seeds: uint64 bits for float64, string for decimal word
	seeds := []struct {
		bits    uint64
		decimal string
	}{
		{0x40091EB851EB851F, "point"},   // 3.14 point
		{0x40091EB851EB851F, "dot"},     // 3.14 dot
		{0x40091EB851EB851F, "and"},     // 3.14 and
		{0x3FE0000000000000, "point"},   // 0.5 point
		{0x0000000000000000, "point"},   // 0.0 point
		{0xBFF0000000000000, "point"},   // -1.0 point
		{0x3FF0000000000000, ""},        // 1.0 with empty decimal word
		{0x40091EB851EB851F, " "},       // 3.14 with space
		{0x3FB999999999999A, "decimal"}, // 0.1 decimal
		{0x412E848000000000, "point"},   // 1000000.0 point
		{0x40091EB851EB851F, "café"},    // 3.14 with unicode
	}
	for _, s := range seeds {
		f.Add(s.bits, s.decimal)
	}

	f.Fuzz(func(_ *testing.T, bits uint64, decimal string) {
		if !utf8.ValidString(decimal) {
			return
		}
		n := math.Float64frombits(bits)
		// Skip NaN and Inf
		if math.IsNaN(n) || math.IsInf(n, 0) {
			return
		}
		// Skip values that are too large to convert to int
		if n > float64(math.MaxInt64) || n < float64(math.MinInt64) {
			return
		}
		// Should not panic
		_ = NumberToWordsFloatWithDecimal(n, decimal)
	})
}

func FuzzNumberToWordsGrouped(f *testing.F) {
	// Seeds: n (int), groupSize (int)
	seeds := []struct {
		n         int
		groupSize int
	}{
		// Basic cases
		{1234, 2}, {1234, 3}, {1234, 4},
		{123456, 2}, {123456, 3},
		{1234567890, 3}, {1234567890, 4},
		// Zero
		{0, 2}, {0, 1}, {0, 0},
		// Negative numbers
		{-1234, 2}, {-1234, 3},
		{-123456, 2}, {-1234567890, 3},
		// Edge case group sizes
		{1234, 0}, {1234, -1}, {1234, 1},
		{1234, 100}, // group size larger than number
		// Single digit
		{5, 2}, {5, 1}, {5, 0},
		// Edge case ints
		{2147483647, 3}, {-2147483648, 3},
		{2147483647, 2}, {-2147483648, 2},
		{999999999, 3}, {-999999999, 3},
		// Phone number style
		{5551234567, 3}, {5551234567, 4},
	}
	for _, s := range seeds {
		f.Add(s.n, s.groupSize)
	}

	f.Fuzz(func(_ *testing.T, n, groupSize int) {
		// Should not panic
		_ = NumberToWordsGrouped(n, groupSize)
	})
}

func FuzzNumberToWordsThreshold(f *testing.F) {
	// Seeds: n (int), threshold (int)
	seeds := []struct {
		n         int
		threshold int
	}{
		// Below threshold
		{1, 10}, {2, 10}, {3, 10}, {5, 10}, {9, 10},
		// At threshold
		{10, 10}, {100, 100}, {0, 0},
		// Above threshold
		{15, 10}, {100, 10}, {1000, 10},
		// Zero
		{0, 10}, {0, 1},
		// Negative numbers
		{-1, 10}, {-5, 10}, {-15, 10},
		{-1, -10}, {-5, -10},
		// Negative threshold
		{5, -10}, {-5, -10}, {0, -1},
		// Edge cases
		{2147483647, 10}, {-2147483648, 10},
		{10, 2147483647}, {10, -2147483648},
		{2147483647, 2147483647}, {-2147483648, -2147483648},
		// Large threshold
		{1000000, 1000000}, {999999, 1000000},
	}
	for _, s := range seeds {
		f.Add(s.n, s.threshold)
	}

	f.Fuzz(func(_ *testing.T, n, threshold int) {
		// Should not panic
		_ = NumberToWordsThreshold(n, threshold)
	})
}

// Fuzz tests for Rails-style helper functions

func FuzzHumanize(f *testing.F) {
	seeds := []string{
		"employee_salary", "author_id", "author_ID", "authorID",
		"hello-world", "XMLParser", "user_name", "firstName",
		"", " ", "a", "already humanized", "multiple__underscores",
		"trailing_id", "UPPERCASE", "MixedCase",
		"café_au_lait", "naïve_approach",
	}
	for _, s := range seeds {
		f.Add(s)
	}

	f.Fuzz(func(_ *testing.T, input string) {
		if !utf8.ValidString(input) {
			return
		}
		_ = Humanize(input)
	})
}

// Covers: ForeignKey, ForeignKeyCondensed.
func FuzzForeignKey(f *testing.F) {
	seeds := []string{
		"Person", "Message", "AdminUser", "user", "XMLParser",
		"", " ", "a", "UPPERCASE", "MixedCase", "snake_case",
		"café", "naïve",
	}
	for _, s := range seeds {
		f.Add(s)
	}

	f.Fuzz(func(_ *testing.T, input string) {
		if !utf8.ValidString(input) {
			return
		}
		_ = ForeignKey(input)
		_ = ForeignKeyCondensed(input)
	})
}

func FuzzTableize(f *testing.F) {
	seeds := []string{
		"Person", "RawScaledScorer", "MouseTrap", "User", "Child",
		"admin_user", "", " ", "a", "UPPERCASE", "MixedCase",
		"café", "naïve",
	}
	for _, s := range seeds {
		f.Add(s)
	}

	f.Fuzz(func(_ *testing.T, input string) {
		if !utf8.ValidString(input) {
			return
		}
		_ = Tableize(input)
	})
}

func FuzzParameterize(f *testing.F) {
	seeds := []string{
		"Hello World!", "Hello, World!", "  Multiple   Spaces  ",
		"Special!@#$%Characters", "Already-Dashed", "under_scored",
		"MixedCase", "", " ", "café au lait",
		"日本語", "über", "Crème brûlée",
	}
	for _, s := range seeds {
		f.Add(s)
	}

	f.Fuzz(func(_ *testing.T, input string) {
		if !utf8.ValidString(input) {
			return
		}
		_ = Parameterize(input)
	})
}

func FuzzParameterizeJoin(f *testing.F) {
	seeds := []struct {
		word string
		sep  string
	}{
		{"Hello World!", "_"}, {"Hello World!", "-"}, {"Hello World!", ""},
		{"Multiple   Spaces", "_"}, {"café au lait", "-"},
		{"", "-"}, {" ", "_"}, {"test", ""},
	}
	for _, s := range seeds {
		f.Add(s.word, s.sep)
	}

	f.Fuzz(func(_ *testing.T, word, sep string) {
		if !utf8.ValidString(word) || !utf8.ValidString(sep) {
			return
		}
		_ = ParameterizeJoin(word, sep)
	})
}

func FuzzTypeify(f *testing.F) {
	seeds := []string{
		"users", "raw_scaled_scorers", "people", "mice", "admin_users",
		"categories", "", " ", "a", "UPPERCASE", "MixedCase",
		"café", "naïve",
	}
	for _, s := range seeds {
		f.Add(s)
	}

	f.Fuzz(func(_ *testing.T, input string) {
		if !utf8.ValidString(input) {
			return
		}
		_ = Typeify(input)
	})
}

func FuzzAsciify(f *testing.F) {
	seeds := []string{
		"café", "naïve", "résumé", "日本語", "hello", "über",
		"Ångström", "", " ", "Crème brûlée",
		"Привет", "مرحبا", "你好", "🎉",
	}
	for _, s := range seeds {
		f.Add(s)
	}

	f.Fuzz(func(_ *testing.T, input string) {
		if !utf8.ValidString(input) {
			return
		}
		_ = Asciify(input)
	})
}

// Fuzz tests for compatibility aliases

// Covers: Pluralize, Singularize.
func FuzzPluralizeSingularize(f *testing.F) {
	seeds := []string{
		"cat", "cats", "child", "children", "person", "people",
		"analysis", "analyses", "", " ", "a",
	}
	for _, s := range seeds {
		f.Add(s)
	}

	f.Fuzz(func(_ *testing.T, input string) {
		if !utf8.ValidString(input) {
			return
		}
		_ = Pluralize(input)
		_ = Singularize(input)
	})
}

// Covers: Camelize, CamelizeDownFirst.
func FuzzCamelizeVariants(f *testing.F) {
	seeds := []string{
		"hello_world", "foo-bar", "some_thing", "XMLParser",
		"camelCase", "PascalCase", "snake_case", "kebab-case",
		"", " ", "a", "ABC", "abc",
	}
	for _, s := range seeds {
		f.Add(s)
	}

	f.Fuzz(func(_ *testing.T, input string) {
		if !utf8.ValidString(input) {
			return
		}
		_ = Camelize(input)
		_ = CamelizeDownFirst(input)
	})
}
