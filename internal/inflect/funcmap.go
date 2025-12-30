package inflect

import "text/template"

// FuncMap returns a template.FuncMap containing inflection functions for use
// with Go's text/template and html/template packages.
//
// The returned FuncMap includes the following functions (using camelCase names):
//
// Pluralization and Singularization:
//   - plural(word string, count ...int) string - Plural form of a noun
//   - singular(word string) string - Singular form of a noun
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
//   - ordinalWord(n int) string - Ordinal in words: 1 -> "first"
//   - numberToWords(n int) string - Number in words: 42 -> "forty-two"
//
// Verb Tenses:
//   - pastTense(verb string) string - Past tense: "walk" -> "walked"
//   - presentParticiple(verb string) string - Present participle: "run" -> "running"
//
// Adjective Comparison:
//   - comparative(adj string) string - Comparative form: "big" -> "bigger"
//   - superlative(adj string) string - Superlative form: "big" -> "biggest"
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
//   - kebabCase(s string) string - Convert to kebab-case
//   - pascalCase(s string) string - Convert to PascalCase
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
		"singular":     e.Singular,
		"pluralNoun":   e.templatePluralNoun,
		"pluralVerb":   e.templatePluralVerb,
		"pluralAdj":    e.templatePluralAdj,
		"singularNoun": e.templateSingularNoun,

		// Articles
		"an": e.An,
		"a":  e.An, // alias

		// Numbers and Ordinals
		"ordinal":       Ordinal,
		"ordinalWord":   OrdinalWord,
		"numberToWords": NumberToWords,

		// Verb Tenses
		"pastTense":         PastTense,
		"presentParticiple": PresentParticiple,

		// Adjective Comparison
		"comparative": Comparative,
		"superlative": Superlative,

		// Possessives
		"possessive": e.Possessive,

		// List Formatting
		"join":     Join,
		"joinWith": JoinWithConj,

		// Case Conversion
		"camelCase":  CamelCase,
		"snakeCase":  SnakeCase,
		"kebabCase":  KebabCase,
		"pascalCase": PascalCase,
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
