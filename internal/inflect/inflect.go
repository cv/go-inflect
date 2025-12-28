// Package inflect provides English language inflection utilities.
//
// It offers functions for pluralization, singularization, indefinite article
// selection (a/an), number-to-words conversion, ordinals, verb participles,
// list joining, and word comparison.
//
// This is a Go port of the Python inflect library
// (https://github.com/jaraco/inflect).
//
// # Basic Usage
//
// The most commonly used functions work on single words:
//
//	inflect.Plural("cat")        // "cats"
//	inflect.Singular("boxes")    // "box"
//	inflect.An("apple")          // "an apple"
//	inflect.NumberToWords(42)    // "forty-two"
//	inflect.Ordinal(3)           // "3rd"
//
// # Part-of-Speech Functions
//
// For more precise inflection, use part-of-speech specific functions:
//
//	inflect.PluralNoun("I")      // "we" (handles pronouns)
//	inflect.PluralVerb("is")     // "are"
//	inflect.PluralAdj("this")    // "these"
//
// These functions accept an optional count parameter:
//
//	inflect.PluralNoun("cat", 1) // "cat" (singular when count=1)
//	inflect.PluralNoun("cat", 2) // "cats"
//
// # Classical Mode
//
// Enable classical/formal pluralization for Latin and Greek forms:
//
//	inflect.ClassicalAll(true)
//	inflect.Plural("formula")    // "formulae" instead of "formulas"
//
// # Custom Definitions
//
// Override default behavior for specific words:
//
//	inflect.DefNoun("regex", "regexen")
//	inflect.Plural("regex")      // "regexen"
//
// # Case Preservation
//
// All functions preserve the case pattern of input words:
//
//	inflect.Plural("CAT")        // "CATS"
//	inflect.Plural("Child")      // "Children"
package inflect
