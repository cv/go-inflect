package inflect

import "strings"

// DefNoun defines a custom noun pluralization rule.
//
// The singular and plural forms are stored in lowercase, and subsequent calls
// to Plural() and Singular() will use this custom rule with case preservation.
//
// Examples:
//
//	DefNoun("foo", "foos")
//	Plural("foo") // returns "foos"
//	Plural("Foo") // returns "Foos"
//	Singular("foos") // returns "foo"
func DefNoun(singular, plural string) {
	lower := strings.ToLower(singular)
	lowerPlural := strings.ToLower(plural)
	irregularPlurals[lower] = lowerPlural
	singularIrregulars[lowerPlural] = lower
}

// UndefNoun removes a custom noun pluralization rule.
//
// This removes only user-defined rules; it cannot remove built-in irregular
// plurals. To restore a built-in rule that was overwritten, use DefNounReset().
//
// Returns true if the rule was removed, false if it didn't exist or was a
// built-in rule.
//
// Examples:
//
//	DefNoun("foo", "foos")
//	Plural("foo") // returns "foos"
//	UndefNoun("foo")
//	Plural("foo") // returns "foos" (standard rule)
func UndefNoun(singular string) bool {
	lower := strings.ToLower(singular)

	// Check if this is a built-in rule
	if _, isBuiltIn := defaultIrregularPlurals[lower]; isBuiltIn {
		return false
	}

	// Check if the rule exists
	plural, exists := irregularPlurals[lower]
	if !exists {
		return false
	}

	// Remove from both maps
	delete(irregularPlurals, lower)
	delete(singularIrregulars, plural)
	return true
}

// DefNounReset resets all noun pluralization rules to their defaults.
//
// This removes all custom rules added via DefNoun() and restores any
// built-in rules that may have been overwritten.
//
// Example:
//
//	DefNoun("child", "childs")  // override built-in
//	DefNoun("foo", "foos")      // add custom
//	DefNounReset()
//	Plural("child") // returns "children" (restored)
//	Plural("foo")   // returns "foos" (standard rule, custom removed)
func DefNounReset() {
	irregularPlurals = copyMap(defaultIrregularPlurals)
	singularIrregulars = buildSingularIrregulars()
}

// customVerbs stores custom verb conjugation rules (singular -> plural).
// This is a placeholder map for future verb conjugation support.
var customVerbs = make(map[string]string)

// customVerbsReverse stores reverse verb conjugation rules (plural -> singular).
var customVerbsReverse = make(map[string]string)

// DefVerb defines a custom verb conjugation rule.
//
// NOTE: This is a placeholder stub for future implementation.
// Full verb conjugation is not yet implemented; this function only stores
// the singular/plural pair in internal maps for future use.
//
// The singular and plural forms are stored in lowercase.
//
// Examples:
//
//	DefVerb("run", "runs")
//	DefVerb("be", "are")
func DefVerb(singular, plural string) {
	lower := strings.ToLower(singular)
	lowerPlural := strings.ToLower(plural)
	customVerbs[lower] = lowerPlural
	customVerbsReverse[lowerPlural] = lower
}

// UndefVerb removes a custom verb conjugation rule.
//
// NOTE: This is a placeholder stub for future implementation.
//
// Returns true if the rule was removed, false if it didn't exist.
//
// Examples:
//
//	DefVerb("run", "runs")
//	UndefVerb("run") // returns true
//	UndefVerb("walk") // returns false (not defined)
func UndefVerb(singular string) bool {
	lower := strings.ToLower(singular)
	plural, exists := customVerbs[lower]
	if !exists {
		return false
	}
	delete(customVerbs, lower)
	delete(customVerbsReverse, plural)
	return true
}

// DefVerbReset resets all custom verb conjugation rules.
//
// NOTE: This is a placeholder stub for future implementation.
//
// This removes all custom rules added via DefVerb().
func DefVerbReset() {
	customVerbs = make(map[string]string)
	customVerbsReverse = make(map[string]string)
}

// customAdjs stores custom adjective pluralization rules (singular -> plural).
// This is a placeholder map for future adjective pluralization support.
var customAdjs = make(map[string]string)

// customAdjsReverse stores reverse adjective rules (plural -> singular).
var customAdjsReverse = make(map[string]string)

// DefAdj defines a custom adjective pluralization rule.
//
// NOTE: This is a placeholder stub for future implementation.
// Full adjective pluralization is not yet implemented; this function only stores
// the singular/plural pair in internal maps for future use.
//
// The singular and plural forms are stored in lowercase.
//
// Examples:
//
//	DefAdj("big", "bigs")
//	DefAdj("happy", "happies")
func DefAdj(singular, plural string) {
	lower := strings.ToLower(singular)
	lowerPlural := strings.ToLower(plural)
	customAdjs[lower] = lowerPlural
	customAdjsReverse[lowerPlural] = lower
}

// UndefAdj removes a custom adjective pluralization rule.
//
// NOTE: This is a placeholder stub for future implementation.
//
// Returns true if the rule was removed, false if it didn't exist.
//
// Examples:
//
//	DefAdj("big", "bigs")
//	UndefAdj("big") // returns true
//	UndefAdj("small") // returns false (not defined)
func UndefAdj(singular string) bool {
	lower := strings.ToLower(singular)
	plural, exists := customAdjs[lower]
	if !exists {
		return false
	}
	delete(customAdjs, lower)
	delete(customAdjsReverse, plural)
	return true
}

// DefAdjReset resets all custom adjective pluralization rules.
//
// NOTE: This is a placeholder stub for future implementation.
//
// This removes all custom rules added via DefAdj().
func DefAdjReset() {
	customAdjs = make(map[string]string)
	customAdjsReverse = make(map[string]string)
}

// AddIrregular is an alias for DefNoun, provided for compatibility with
// github.com/jinzhu/inflection.
//
// Example:
//
//	AddIrregular("person", "people")
//	Plural("person") // returns "people"
func AddIrregular(singular, plural string) {
	DefNoun(singular, plural)
}

// AddUncountable marks words as uncountable (same singular and plural form),
// provided for compatibility with github.com/jinzhu/inflection and
// github.com/go-openapi/inflect.
//
// Example:
//
//	AddUncountable("fish", "sheep")
//	Plural("fish")  // returns "fish"
//	Plural("sheep") // returns "sheep"
func AddUncountable(words ...string) {
	for _, w := range words {
		DefNoun(w, w)
	}
}

// Pluralize is an alias for Plural, provided for compatibility with
// github.com/go-openapi/inflect.
//
// Example:
//
//	Pluralize("cat") // returns "cats"
func Pluralize(word string) string {
	return Plural(word)
}

// Singularize is an alias for Singular, provided for compatibility with
// github.com/go-openapi/inflect.
//
// Example:
//
//	Singularize("cats") // returns "cat"
func Singularize(word string) string {
	return Singular(word)
}

// Camelize is an alias for PascalCase, provided for compatibility with
// github.com/go-openapi/inflect.
//
// Example:
//
//	Camelize("hello_world") // returns "HelloWorld"
func Camelize(word string) string {
	return PascalCase(word)
}

// CamelizeDownFirst is an alias for CamelCase, provided for compatibility with
// github.com/go-openapi/inflect.
//
// Example:
//
//	CamelizeDownFirst("hello_world") // returns "helloWorld"
func CamelizeDownFirst(word string) string {
	return CamelCase(word)
}
