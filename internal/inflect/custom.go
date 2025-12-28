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
	defaultEngine.DefNoun(singular, plural)
}

// DefNoun defines a custom noun pluralization rule.
//
// The singular and plural forms are stored in lowercase, and subsequent calls
// to Plural() and Singular() will use this custom rule with case preservation.
//
// Examples:
//
//	e := NewEngine()
//	e.DefNoun("foo", "foos")
//	e.Plural("foo") // returns "foos"
//	e.Plural("Foo") // returns "Foos"
//	e.Singular("foos") // returns "foo"
func (e *Engine) DefNoun(singular, plural string) {
	e.mu.Lock()
	defer e.mu.Unlock()
	lower := strings.ToLower(singular)
	lowerPlural := strings.ToLower(plural)
	e.irregularPlurals[lower] = lowerPlural
	e.singularIrregulars[lowerPlural] = lower
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
	return defaultEngine.UndefNoun(singular)
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
//	e := NewEngine()
//	e.DefNoun("foo", "foos")
//	e.Plural("foo") // returns "foos"
//	e.UndefNoun("foo")
//	e.Plural("foo") // returns "foos" (standard rule)
func (e *Engine) UndefNoun(singular string) bool {
	e.mu.Lock()
	defer e.mu.Unlock()
	lower := strings.ToLower(singular)

	// Check if this is a built-in rule
	if _, isBuiltIn := defaultIrregularPlurals[lower]; isBuiltIn {
		return false
	}

	// Check if the rule exists
	plural, exists := e.irregularPlurals[lower]
	if !exists {
		return false
	}

	// Remove from both maps
	delete(e.irregularPlurals, lower)
	delete(e.singularIrregulars, plural)
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
	defaultEngine.DefNounReset()
}

// DefNounReset resets all noun pluralization rules to their defaults.
//
// This removes all custom rules added via DefNoun() and restores any
// built-in rules that may have been overwritten.
//
// Example:
//
//	e := NewEngine()
//	e.DefNoun("child", "childs")  // override built-in
//	e.DefNoun("foo", "foos")      // add custom
//	e.DefNounReset()
//	e.Plural("child") // returns "children" (restored)
//	e.Plural("foo")   // returns "foos" (standard rule, custom removed)
func (e *Engine) DefNounReset() {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.irregularPlurals = copyMap(defaultIrregularPlurals)
	// Build singularIrregulars as reverse of irregularPlurals
	e.singularIrregulars = make(map[string]string, len(e.irregularPlurals))
	for singular, plural := range e.irregularPlurals {
		e.singularIrregulars[plural] = singular
	}
}

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
	defaultEngine.DefVerb(singular, plural)
}

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
//	e := NewEngine()
//	e.DefVerb("run", "runs")
//	e.DefVerb("be", "are")
func (e *Engine) DefVerb(singular, plural string) {
	e.mu.Lock()
	defer e.mu.Unlock()
	lower := strings.ToLower(singular)
	lowerPlural := strings.ToLower(plural)
	e.customVerbs[lower] = lowerPlural
	e.customVerbsReverse[lowerPlural] = lower
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
	return defaultEngine.UndefVerb(singular)
}

// UndefVerb removes a custom verb conjugation rule.
//
// NOTE: This is a placeholder stub for future implementation.
//
// Returns true if the rule was removed, false if it didn't exist.
//
// Examples:
//
//	e := NewEngine()
//	e.DefVerb("run", "runs")
//	e.UndefVerb("run") // returns true
//	e.UndefVerb("walk") // returns false (not defined)
func (e *Engine) UndefVerb(singular string) bool {
	e.mu.Lock()
	defer e.mu.Unlock()
	lower := strings.ToLower(singular)
	plural, exists := e.customVerbs[lower]
	if !exists {
		return false
	}
	delete(e.customVerbs, lower)
	delete(e.customVerbsReverse, plural)
	return true
}

// DefVerbReset resets all custom verb conjugation rules.
//
// NOTE: This is a placeholder stub for future implementation.
//
// This removes all custom rules added via DefVerb().
func DefVerbReset() {
	defaultEngine.DefVerbReset()
}

// DefVerbReset resets all custom verb conjugation rules.
//
// NOTE: This is a placeholder stub for future implementation.
//
// This removes all custom rules added via DefVerb().
//
// Example:
//
//	e := NewEngine()
//	e.DefVerb("run", "runs")
//	e.DefVerbReset()
//	e.UndefVerb("run") // returns false (rule was reset)
func (e *Engine) DefVerbReset() {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.customVerbs = make(map[string]string)
	e.customVerbsReverse = make(map[string]string)
}

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
	defaultEngine.DefAdj(singular, plural)
}

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
//	e := NewEngine()
//	e.DefAdj("big", "bigs")
//	e.DefAdj("happy", "happies")
func (e *Engine) DefAdj(singular, plural string) {
	e.mu.Lock()
	defer e.mu.Unlock()
	lower := strings.ToLower(singular)
	lowerPlural := strings.ToLower(plural)
	e.customAdjs[lower] = lowerPlural
	e.customAdjsReverse[lowerPlural] = lower
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
	return defaultEngine.UndefAdj(singular)
}

// UndefAdj removes a custom adjective pluralization rule.
//
// NOTE: This is a placeholder stub for future implementation.
//
// Returns true if the rule was removed, false if it didn't exist.
//
// Examples:
//
//	e := NewEngine()
//	e.DefAdj("big", "bigs")
//	e.UndefAdj("big") // returns true
//	e.UndefAdj("small") // returns false (not defined)
func (e *Engine) UndefAdj(singular string) bool {
	e.mu.Lock()
	defer e.mu.Unlock()
	lower := strings.ToLower(singular)
	plural, exists := e.customAdjs[lower]
	if !exists {
		return false
	}
	delete(e.customAdjs, lower)
	delete(e.customAdjsReverse, plural)
	return true
}

// DefAdjReset resets all custom adjective pluralization rules.
//
// NOTE: This is a placeholder stub for future implementation.
//
// This removes all custom rules added via DefAdj().
func DefAdjReset() {
	defaultEngine.DefAdjReset()
}

// DefAdjReset resets all custom adjective pluralization rules.
//
// NOTE: This is a placeholder stub for future implementation.
//
// This removes all custom rules added via DefAdj().
//
// Example:
//
//	e := NewEngine()
//	e.DefAdj("big", "bigs")
//	e.DefAdjReset()
//	e.UndefAdj("big") // returns false (rule was reset)
func (e *Engine) DefAdjReset() {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.customAdjs = make(map[string]string)
	e.customAdjsReverse = make(map[string]string)
}

// AddIrregular is an alias for DefNoun, provided for compatibility with
// github.com/jinzhu/inflection.
//
// Example:
//
//	AddIrregular("person", "people")
//	Plural("person") // returns "people"
func AddIrregular(singular, plural string) {
	defaultEngine.DefNoun(singular, plural)
}

// AddIrregular is an alias for DefNoun, provided for compatibility with
// github.com/jinzhu/inflection.
//
// Example:
//
//	e := NewEngine()
//	e.AddIrregular("person", "people")
//	e.Plural("person") // returns "people"
func (e *Engine) AddIrregular(singular, plural string) {
	e.DefNoun(singular, plural)
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
	defaultEngine.AddUncountable(words...)
}

// AddUncountable marks words as uncountable (same singular and plural form),
// provided for compatibility with github.com/jinzhu/inflection and
// github.com/go-openapi/inflect.
//
// Example:
//
//	e := NewEngine()
//	e.AddUncountable("fish", "sheep")
//	e.Plural("fish")  // returns "fish"
//	e.Plural("sheep") // returns "sheep"
func (e *Engine) AddUncountable(words ...string) {
	for _, w := range words {
		e.DefNoun(w, w)
	}
}

// Pluralize is an alias for Plural, provided for compatibility with
// github.com/go-openapi/inflect.
//
// Example:
//
//	Pluralize("cat") // returns "cats"
func Pluralize(word string) string {
	return defaultEngine.Plural(word)
}

// Pluralize is an alias for Plural, provided for compatibility with
// github.com/go-openapi/inflect.
//
// Example:
//
//	e := NewEngine()
//	e.Pluralize("cat") // returns "cats"
func (e *Engine) Pluralize(word string) string {
	return e.Plural(word)
}

// Singularize is an alias for Singular, provided for compatibility with
// github.com/go-openapi/inflect.
//
// Example:
//
//	Singularize("cats") // returns "cat"
func Singularize(word string) string {
	return defaultEngine.Singular(word)
}

// Singularize is an alias for Singular, provided for compatibility with
// github.com/go-openapi/inflect.
//
// Example:
//
//	e := NewEngine()
//	e.Singularize("cats") // returns "cat"
func (e *Engine) Singularize(word string) string {
	return e.Singular(word)
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
