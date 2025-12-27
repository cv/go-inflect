# Contributing to go-inflect

Thank you for your interest in contributing to go-inflect! This document provides guidelines and instructions for contributing.

## Reporting Bugs

If you find a bug, please create an issue with the following information:

- A clear, descriptive title
- Steps to reproduce the problem
- Expected behavior vs. actual behavior
- Go version (`go version`)
- Operating system and version
- Minimal code example that demonstrates the issue

Example:
```go
// Expected: "an hour"
// Actual: "a hour"
result := inflect.An("hour")
```

## Suggesting Features

Feature requests are welcome! When suggesting a feature:

1. Check existing issues to avoid duplicates
2. Create a new issue with the label "enhancement"
3. Describe the use case and why it would be valuable
4. If possible, reference the Python inflect library behavior for consistency
5. Provide example API usage showing how it might work

## Development Workflow

### 1. Fork and Clone

```bash
# Fork the repository on GitHub/GitLab, then:
git clone https://github.com/YOUR_USERNAME/go-inflect.git
cd go-inflect
```

### 2. Create a Branch

Create a descriptive branch name for your changes:

```bash
git checkout -b feature/add-new-pluralization-rule
# or
git checkout -b fix/silent-h-handling
```

### 3. Make Changes

- Write your code following the style guidelines below
- Add or update tests for your changes
- Update documentation if needed

### 4. Run Tests

Ensure all tests pass before submitting:

```bash
go test -v
```

For coverage information:

```bash
go test -cover
```

### 5. Run Linters

Format your code and check for issues:

```bash
gofmt -w .
go vet ./...
```

### 6. Submit a Pull Request

1. Commit your changes with a clear, descriptive message
2. Push your branch to your fork
3. Open a pull request against the main branch
4. Fill out the PR template with relevant details
5. Wait for review and address any feedback

## Code Style Guidelines

### Follow Effective Go

Adhere to the principles outlined in [Effective Go](https://go.dev/doc/effective_go):

- Use clear, concise naming (e.g., `Plural` not `ConvertWordToPluralForm`)
- Keep functions focused and small
- Handle errors appropriately
- Use Go idioms and patterns

### Use Table-Driven Tests

Structure tests as table-driven for better coverage and readability:

```go
func TestPlural(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected string
    }{
        {"regular noun", "cat", "cats"},
        {"sibilant ending", "box", "boxes"},
        {"consonant + y", "baby", "babies"},
        {"irregular", "child", "children"},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Plural(tt.input)
            if result != tt.expected {
                t.Errorf("Plural(%q) = %q, want %q", tt.input, result, tt.expected)
            }
        })
    }
}
```

### Add Doc Comments to Exported Functions

All exported functions, types, and constants must have documentation comments:

```go
// Plural converts a singular English noun to its plural form.
// It handles regular pluralization rules as well as common irregular nouns.
//
// Examples:
//   Plural("cat")   // returns "cats"
//   Plural("child") // returns "children"
func Plural(word string) string {
    // implementation
}
```

### Maintain High Test Coverage

- Aim for high test coverage on all new code
- Include edge cases (empty strings, special characters, etc.)
- Test both positive and negative cases
- Run `go test -cover` to check coverage percentages

## Code Review Process

1. All submissions require review before merging
2. Reviewers may request changes—please address feedback promptly
3. Keep pull requests focused; one feature or fix per PR
4. Ensure CI checks pass before requesting review

## License

By contributing to go-inflect, you agree that your contributions will be licensed under the [MIT License](LICENSE).

---

Questions? Feel free to open an issue for clarification. We appreciate your contributions!
