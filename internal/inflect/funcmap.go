package inflect

import "text/template"

// FuncMap returns a template.FuncMap containing inflection functions for use
// with Go's text/template and html/template packages.
//
// The returned FuncMap includes the following functions (using camelCase names):
//
// Pluralization and Singularization:
//   - plural(word string, count ...int) string - Plural form of a noun
//   - pluralize(word string) string - Alias for plural
//   - singular(word string) string - Singular form of a noun
//   - singularize(word string) string - Alias for singular
//   - pluralNoun(word string, count ...int) string - Plural form with pronoun support
//   - pluralVerb(word string, count ...int) string - Plural form of a verb
//   - pluralAdj(word string, count ...int) string - Plural form of an adjective
//   - singularNoun(word string, count ...int) string - Singular form with pronoun support
//
// Articles:
//   - an(word string) string - Prefixes word with "a" or "an"
//   - a(word string) string - Alias for an()
//
// Numbers and Ordinals:
//   - ordinal(n int) string - Ordinal with suffix: 1 -> "1st"
//   - ordinalSuffix(n int) string - Just the suffix: 1 -> "st"
//   - ordinalWord(n int) string - Ordinal in words: 1 -> "first"
//   - ordinalToCardinal(s string) string - "first" -> "one"
//   - wordToOrdinal(s string) string - "one" -> "first"
//   - numberToWords(n int) string - Number in words: 42 -> "forty-two"
//   - numberToWordsWithAnd(n int) string - With "and": 123 -> "one hundred and twenty-three"
//   - formatNumber(n int) string - With commas: 1000 -> "1,000"
//   - countingWord(n int) string - 1 -> "once", 2 -> "twice", 3 -> "3 times"
//   - fractionToWords(num, denom int) string - 1,4 -> "one quarter"
//   - currencyToWords(amount float64, currency string) string - 1.50, "USD" -> "one dollar and fifty cents"
//   - no(word string, count int) string - 0 -> "no cats", 1 -> "1 cat"
//
// Verb Tenses:
//   - pastTense(verb string) string - Past tense: "walk" -> "walked"
//   - pastParticiple(verb string) string - Past participle: "take" -> "taken"
//   - presentParticiple(verb string) string - Present participle: "run" -> "running"
//   - futureTense(verb string) string - Future tense: "walk" -> "will walk"
//
// Adjectives and Adverbs:
//   - comparative(adj string) string - Comparative form: "big" -> "bigger"
//   - superlative(adj string) string - Superlative form: "big" -> "biggest"
//   - adverb(adj string) string - Adverb form: "quick" -> "quickly"
//
// Possessives:
//   - possessive(word string) string - Possessive form: "cat" -> "cat's"
//
// List Formatting:
//   - join(words []string) string - Join with Oxford comma: ["a","b","c"] -> "a, b, and c"
//   - joinWith(words []string, conj string) string - Join with custom conjunction
//
// Case Conversion:
//   - camelCase(s string) string - Convert to camelCase
//   - snakeCase(s string) string - Convert to snake_case
//   - underscore(s string) string - Alias for snakeCase
//   - kebabCase(s string) string - Convert to kebab-case
//   - dasherize(s string) string - Alias for kebabCase
//   - pascalCase(s string) string - Convert to PascalCase
//   - titleCase(s string) string - Alias for pascalCase
//   - camelize(s string) string - Alias for pascalCase
//   - camelizeDownFirst(s string) string - Alias for camelCase
//
// Text Transformation:
//   - capitalize(s string) string - Capitalize first letter: "hello" -> "Hello"
//   - titleize(s string) string - Capitalize each word: "hello world" -> "Hello World"
//   - humanize(s string) string - Human readable: "employee_salary" -> "Employee salary"
//
// Rails-style Helpers:
//   - tableize(word string) string - Type to table: "Person" -> "people"
//   - foreignKey(word string) string - Type to FK: "User" -> "user_id"
//   - typeify(word string) string - Table to type: "user_posts" -> "UserPost"
//   - parameterize(word string) string - URL slug: "Hello World" -> "hello-world"
//   - asciify(word string) string - Remove diacritics: "café" -> "cafe"
//
// Utility:
//   - wordCount(text string) int - Count words in text
//   - countSyllables(word string) int - Count syllables in word
//
// Example usage:
//
//	tmpl := template.New("example").Funcs(inflect.FuncMap())
//	tmpl.Parse(`I have {{plural "cat" .Count}} and {{an "apple"}}`)
//
//	// With count parameter:
//	tmpl.Parse(`There {{if eq .Count 1}}is{{else}}are{{end}} {{plural "item" .Count}}`)
//
// For custom engine configurations, use Engine.FuncMap() instead.
func FuncMap() template.FuncMap {
	return defaultEngine.FuncMap()
}

// FuncMap returns a template.FuncMap containing inflection functions that use
// this Engine's configuration.
//
// This allows template functions to respect custom noun definitions, classical
// modes, and other engine settings.
//
// Example:
//
//	e := inflect.NewEngine()
//	e.DefNoun("foo", "fooz")
//	e.Classical(inflect.ClassicalAll, true)
//
//	tmpl := template.New("example").Funcs(e.FuncMap())
//	tmpl.Parse(`The plural of foo is {{plural "foo"}}`)
func (e *Engine) FuncMap() template.FuncMap {
	return template.FuncMap{
		// Pluralization and Singularization
		"plural":       e.templatePlural,
		"pluralize":    e.Plural, // alias
		"singular":     e.Singular,
		"singularize":  e.Singular, // alias
		"pluralNoun":   e.templatePluralNoun,
		"pluralVerb":   e.templatePluralVerb,
		"pluralAdj":    e.templatePluralAdj,
		"singularNoun": e.templateSingularNoun,

		// Articles
		"an": e.An,
		"a":  e.An, // alias

		// Numbers and Ordinals
		"ordinal":              Ordinal,
		"ordinalSuffix":        OrdinalSuffix,
		"ordinalWord":          OrdinalWord,
		"ordinalToCardinal":    OrdinalToCardinal,
		"wordToOrdinal":        WordToOrdinal,
		"numberToWords":        NumberToWords,
		"numberToWordsWithAnd": NumberToWordsWithAnd,
		"formatNumber":         FormatNumber,
		"countingWord":         CountingWord,
		"fractionToWords":      FractionToWords,
		"currencyToWords":      CurrencyToWords,
		"no":                   e.templateNo,

		// Verb Tenses
		"pastTense":         PastTense,
		"pastParticiple":    PastParticiple,
		"presentParticiple": PresentParticiple,
		"futureTense":       FutureTense,

		// Adjectives and Adverbs
		"comparative": Comparative,
		"superlative": Superlative,
		"adverb":      Adverb,

		// Possessives
		"possessive": e.Possessive,

		// List Formatting
		"join":     Join,
		"joinWith": JoinWithConj,

		// Case Conversion
		"camelCase":         CamelCase,
		"snakeCase":         SnakeCase,
		"underscore":        Underscore, // alias
		"kebabCase":         KebabCase,
		"dasherize":         Dasherize, // alias
		"pascalCase":        PascalCase,
		"titleCase":         TitleCase,         // alias
		"camelize":          Camelize,          // alias
		"camelizeDownFirst": CamelizeDownFirst, // alias

		// Text Transformation
		"capitalize": Capitalize,
		"titleize":   Titleize,
		"humanize":   Humanize,

		// Rails-style Helpers
		"tableize":     Tableize,
		"foreignKey":   ForeignKey,
		"typeify":      Typeify,
		"parameterize": Parameterize,
		"asciify":      Asciify,

		// Utility
		"wordCount":      WordCount,
		"countSyllables": CountSyllables,
	}
}

// templatePlural wraps Plural for template use, handling the variadic count parameter.
// Templates can call it as {{plural "cat"}} or {{plural "cat" .Count}}.
func (e *Engine) templatePlural(word string, count ...int) string {
	if len(count) > 0 && (count[0] == 1 || count[0] == -1) {
		return word
	}
	return e.Plural(word)
}

// templatePluralNoun wraps PluralNoun for template use.
func (e *Engine) templatePluralNoun(word string, count ...int) string {
	return e.PluralNoun(word, count...)
}

// templatePluralVerb wraps PluralVerb for template use.
func (e *Engine) templatePluralVerb(word string, count ...int) string {
	return e.PluralVerb(word, count...)
}

// templatePluralAdj wraps PluralAdj for template use.
func (e *Engine) templatePluralAdj(word string, count ...int) string {
	return e.PluralAdj(word, count...)
}

// templateSingularNoun wraps SingularNoun for template use.
func (e *Engine) templateSingularNoun(word string, count ...int) string {
	return e.SingularNoun(word, count...)
}

// templateNo wraps No for template use.
func (e *Engine) templateNo(word string, count int) string {
	return e.No(word, count)
}
