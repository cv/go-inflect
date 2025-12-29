// Package inflect provides English word inflection functions.
//
// All package-level functions are safe for concurrent use. For isolated
// configurations (e.g., different classical mode settings), create separate
// [Engine] instances with [NewEngine].
//
//go:generate go run ./tools/gen-exports
package inflect
