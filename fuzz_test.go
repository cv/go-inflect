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
	// Seed corpus with interesting cases
	seeds := []string{
		"cat", "dog", "child", "mouse", "fish",
		"analysis", "cactus", "datum", "phenomenon",
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
		"cats", "dogs", "children", "mice", "fish",
		"analyses", "cacti", "data", "phenomena",
		"", " ", "boxes", "buses", "churches",
		"CATS", "Cats", "cats",
		"potatoes", "heroes", "photos",
		"leaves", "wolves", "knives",
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

func FuzzAn(f *testing.F) {
	seeds := []string{
		"apple", "banana", "hour", "honest", "university",
		"FBI", "URL", "YAML", "XML",
		"", " ", "a", "an", "the",
		"European", "one", "once", "unicorn",
		"heir", "herb", "hotel",
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

func FuzzNumberToWords(f *testing.F) {
	seeds := []int{
		0, 1, -1, 10, 100, 1000,
		12, 21, 99, 100, 101, 111,
		1000000, 1000000000,
		-999999999, 2147483647, -2147483648,
	}
	for _, n := range seeds {
		f.Add(n)
	}

	f.Fuzz(func(_ *testing.T, input int) {
		_ = NumberToWords(input)
		_ = NumberToWordsWithAnd(input)
	})
}

func FuzzOrdinal(f *testing.F) {
	seeds := []int{
		0, 1, 2, 3, 4, 11, 12, 13, 21, 22, 23,
		100, 101, 102, 103, 111, 112, 113,
		-1, -2, -3, -11, -12, -13,
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
		"run", "walk", "go", "be", "have",
		"stop", "hop", "sit", "cut",
		"love", "hate", "make", "take",
		"try", "cry", "fly", "die",
		"", " ", "a", "123",
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
		"walk", "run", "go", "be", "have",
		"stop", "try", "play", "stay",
		"love", "hate", "make", "take",
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
		"walk", "run", "go", "be", "have",
		"take", "give", "see", "do",
		"break", "speak", "write", "drive",
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

func FuzzComparative(f *testing.F) {
	seeds := []string{
		"big", "small", "beautiful", "good", "bad",
		"happy", "hot", "large", "nice",
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

func FuzzWordToOrdinal(f *testing.F) {
	seeds := []string{
		"one", "two", "three", "first", "second", "third",
		"1", "2", "3", "1st", "2nd", "3rd",
		"twenty-one", "twenty-first",
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
		"cat", "dog", "child", "mouse", "fish",
		"box", "bus", "church", "potato", "hero",
		// Pronouns
		"I", "me", "my", "mine", "myself",
		"you", "your", "yours", "yourself",
		"he", "she", "it", "they",
		"him", "her", "them",
		"his", "hers", "its", "their", "theirs",
		"we", "us", "our", "ours", "ourselves",
		// Irregular plurals
		"analysis", "cactus", "datum", "phenomenon",
		"leaf", "wolf", "knife",
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
		// Auxiliary verbs
		"is", "was", "has", "does", "am", "are", "were", "have", "do",
		// Contractions
		"isn't", "wasn't", "hasn't", "doesn't", "aren't", "weren't", "haven't", "don't",
		// Modal verbs (unchanged)
		"can", "could", "may", "might", "must", "shall", "should", "will", "would",
		// Regular verbs (third person singular)
		"runs", "walks", "goes", "sees", "flies", "tries",
		"passes", "pushes", "watches", "fixes", "buzzes",
		// Base form verbs
		"run", "walk", "go", "see", "fly", "try",
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
		"cats", "dogs", "children", "mice", "fish",
		"boxes", "buses", "churches", "potatoes", "heroes",
		// Pronouns (plural)
		"we", "us", "our", "ours", "ourselves",
		"they", "them", "their", "theirs", "themselves",
		// Pronouns (singular)
		"I", "me", "my", "mine", "myself",
		"he", "she", "it", "him", "her",
		// Irregular plurals
		"analyses", "cacti", "data", "phenomena",
		"leaves", "wolves", "knives",
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
