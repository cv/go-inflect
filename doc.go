// Package inflect provides English word inflection functions.
//
// # Concurrency
//
// All package-level functions are safe for concurrent use. The package uses
// a shared default [Engine] protected by a read-write mutex:
//
//   - Inflection methods (Plural, Singular, etc.) acquire read locks
//   - Configuration methods (Classical, DefNoun, etc.) acquire write locks
//
// For isolated configurations, use [NewEngine] to create separate instances.
// Each Engine has its own mutex, so operations on different engines don't
// block each other.
//
// This package re-exports all functionality from internal/inflect.
// To regenerate the exports after modifying the internal package, run:
//
//	go generate
//
//go:generate go run ./tools/gen-exports
package inflect
