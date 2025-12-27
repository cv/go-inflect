# Claude Workflow for go-inflect

## Project Overview

This is a Go port of the Python [inflect](https://pypi.org/project/inflect/) library for English word inflection. The goal is to provide idiomatic Go equivalents for pluralization, singularization, indefinite articles, number-to-words conversion, and more.

## Development Workflow

### 1. Find Work
```bash
bd ready              # See available issues
bd show <id>          # View issue details
```

### 2. Claim & Implement
```bash
bd update <id> --status in_progress  # Claim the work
```

Then:
1. Read the test file first (tests define requirements)
2. Check `reference/` for Python source behavior
3. Implement in `inflect.go`
4. Run tests: `go test -v`

### 3. Quality Gates
```bash
gofmt -w .            # Format code
go vet ./...          # Static analysis
go test -v            # Run tests
go test -cover        # Check coverage
```

### 4. Commit & Close
```bash
git add -A
git commit -m "Implement feature X..."
bd close <id> -r "Reason"
```

### 5. Push (MANDATORY before session end)
```bash
git pull --rebase
bd sync
git push
```

## Code Style

- Follow [Effective Go](https://go.dev/doc/effective_go)
- Package comment at top of `inflect.go`
- Doc comments on all exported functions
- Table-driven tests
- Handle edge cases (empty strings, single characters)

## Reference Materials

- `reference/tests/` - Python test files showing expected behavior
- `reference/README.rst` - Python library documentation
- `features.md` - Feature list to implement
- `ROADMAP.md` - Development phases
- `TEST_COVERAGE.md` - Test case reference

## Architecture Decisions

### Pronunciation-Based Article Selection (An/A)
The `An()` function uses a multi-layered approach:
1. Check for silent 'h' words (honest, hour, heir)
2. Check for abbreviations (uppercase = letter pronunciation)
3. Check for known lowercase abbreviations (mpeg, jpeg)
4. Check for consonant-Y sounds in U-words (Ukrainian, unanimous)
5. Default: vowel letters get "an", consonants get "a"

This matches the Python inflect behavior while being idiomatic Go.

## Common Patterns

### Adding New Inflection Functions

1. Define the exported function with doc comment
2. Create helper functions for rule matching
3. Use word lists for exceptions
4. Write table-driven tests first

Example structure:
```go
// Plural returns the plural form of an English noun.
func Plural(word string) string {
    if word == "" {
        return ""
    }
    // Check irregular forms first
    // Apply suffix rules
    // Default behavior
}
```

## Session Handoff Notes

Last session (2025-12-26):
- Implemented `An()` and `A()` functions
- All 21 tests passing
- Next: Pluralization epic (go-inflect-tw5)
