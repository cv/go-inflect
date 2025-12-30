// Package inflect provides English language inflection utilities.
//
// This package offers 100+ functions for pluralization, singularization,
// articles, numbers, ordinals, verb tenses, possessives, and more.
//
// All package-level functions are safe for concurrent use. For isolated
// configurations (e.g., different classical mode settings), create separate
// [Engine] instances with [NewEngine].
//
// The easiest way to use this package is with Go templates via [FuncMap]:
//
//	tmpl := template.New("example").Funcs(inflect.FuncMap())
//	tmpl.Parse(`I have {{plural "cat" .Count}}`)
//
// All functions are also available for direct use:
//
//	inflect.Plural("child")      // "children"
//	inflect.An("apple")          // "an apple"
//	inflect.NumberToWords(42)    // "forty-two"
//
//go:generate go run ./tools/gen-exports.go
package inflect
