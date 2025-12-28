# go-inflect

A Go port of the [Python inflect](https://github.com/jaraco/inflect) library for English language inflection.

## Installation

```bash
go get github.com/cv/go-inflect
```

## Quick Start

```go
import "github.com/cv/go-inflect"

// Indefinite articles
inflect.An("apple")       // "an apple"
inflect.An("banana")      // "a banana"
inflect.An("hour")        // "an hour" (silent h)
inflect.An("FBI agent")   // "an FBI agent"

// Pluralization
inflect.Plural("cat")     // "cats"
inflect.Plural("child")   // "children"
inflect.Plural("analysis")// "analyses"

// Singularization
inflect.Singular("boxes") // "box"
inflect.Singular("mice")  // "mouse"

// Numbers to words
inflect.NumberToWords(42) // "forty-two"
inflect.Ordinal(3)        // "3rd"
inflect.OrdinalWord(3)    // "third"

// List joining (Oxford comma)
inflect.Join([]string{"a", "b", "c"}) // "a, b, and c"

// Verb forms
inflect.PresentParticiple("run")  // "running"
inflect.PastParticiple("take")    // "taken"
inflect.PastTense("walk")         // "walked"
inflect.PastTense("go")           // "went"

// Word comparison
inflect.Compare("cat", "cats")    // "s:p" (singular to plural)

// Adjective forms
inflect.Comparative("big")        // "bigger"
inflect.Superlative("big")        // "biggest"
inflect.Comparative("beautiful")  // "more beautiful"
inflect.Superlative("beautiful")  // "most beautiful"
inflect.Adverb("quick")           // "quickly"
inflect.Adverb("happy")           // "happily"

// Possessive forms
inflect.Possessive("cat")         // "cat's"
inflect.Possessive("cats")        // "cats'"
inflect.Possessive("children")    // "children's"
inflect.Possessive("James")       // "James's" (modern style)

// Fractions
inflect.FractionToWords(1, 2)     // "one half"
inflect.FractionToWords(3, 4)     // "three quarters"
inflect.FractionToWords(2, 3)     // "two thirds"

// Currency
inflect.CurrencyToWords(123.45, "USD") // "one hundred twenty-three dollars and forty-five cents"
inflect.CurrencyToWords(1.00, "GBP")   // "one pound"

// Counting words
inflect.CountingWord(1)           // "once"
inflect.CountingWord(2)           // "twice"
inflect.CountingWord(5)           // "five times"
```

## Core Functions

### Nouns

| Function | Description |
|----------|-------------|
| `Plural(word)` | Pluralize a noun: "cat" → "cats" |
| `Singular(word)` | Singularize a noun: "cats" → "cat" |
| `PluralNoun(word, count...)` | Pluralize nouns/pronouns with optional count |
| `SingularNoun(word, count...)` | Singularize nouns/pronouns with optional count |

### Verbs

| Function | Description |
|----------|-------------|
| `PluralVerb(word, count...)` | "is" → "are", "runs" → "run" |
| `PresentParticiple(verb)` | "run" → "running" |
| `PastParticiple(verb)` | "take" → "taken", "walk" → "walked" |
| `PastTense(verb)` | "walk" → "walked", "go" → "went" |

### Articles & Adjectives

| Function | Description |
|----------|-------------|
| `An(word)` / `A(word)` | Prefix with "a" or "an" |
| `PluralAdj(word, count...)` | "this" → "these", "a" → "some" |
| `Comparative(adj)` | "big" → "bigger", "beautiful" → "more beautiful" |
| `Superlative(adj)` | "big" → "biggest", "beautiful" → "most beautiful" |
| `Adverb(adj)` | "quick" → "quickly", "happy" → "happily" |
| `Possessive(noun)` | "cat" → "cat's", "cats" → "cats'" |

### Numbers

| Function | Description |
|----------|-------------|
| `NumberToWords(n)` | 42 → "forty-two" |
| `NumberToWordsWithAnd(n)` | 101 → "one hundred and one" |
| `NumberToWordsFloat(f)` | 3.14 → "three point one four" |
| `NumberToWordsGrouped(n, size)` | Group digits: 1234 with size 2 → "twelve thirty-four" |
| `NumberToWordsThreshold(n, t)` | Words if n < t, else digits |
| `FormatNumber(n)` | 1000 → "1,000" (thousand separators) |
| `Ordinal(n)` | 1 → "1st", 2 → "2nd" |
| `OrdinalWord(n)` | 1 → "first", 2 → "second" |
| `OrdinalSuffix(n)` | 1 → "st", 2 → "nd", 3 → "rd" |
| `WordToOrdinal(s)` | "one" → "first", "1" → "1st" |
| `OrdinalToCardinal(s)` | "1st" → "1", "first" → "one" |
| `CountingWord(n)` | 1 → "once", 2 → "twice", 5 → "five times" |
| `CountingWordThreshold(n, t)` | Words if n < t, else "n times" |
| `FractionToWords(num, den)` | (3, 4) → "three quarters" |
| `FractionToWordsWithFourths(n, d)` | (1, 4) → "one fourth" |
| `CurrencyToWords(amount, code)` | (1.50, "USD") → "one dollar and fifty cents" |
| `No(word, count)` | "no errors" or "3 errors" |

### Lists

| Function | Description |
|----------|-------------|
| `Join(words)` | Join with "and" and Oxford comma |
| `JoinWithConj(words, conj)` | Custom conjunction ("or", "and/or") |
| `JoinNoOxford(words)` | British style without Oxford comma |
| `JoinWithSep(words, conj, sep)` | Custom separator |
| `JoinWithAutoSep(words, conj)` | Auto `;` if items contain commas |

### Comparison

| Function | Description |
|----------|-------------|
| `Compare(w1, w2)` | Returns "eq", "s:p", "p:s", "p:p", or "" |
| `CompareNouns(n1, n2)` | Compare noun forms |
| `CompareVerbs(v1, v2)` | Compare verb forms |
| `CompareAdjs(a1, a2)` | Compare adjective forms |

### Utilities

| Function | Description |
|----------|-------------|
| `IsPlural(word)` / `IsSingular(word)` | Check word form |
| `IsOrdinal(s)` / `IsParticiple(word)` | Check word type |
| `WordCount(text)` | Count words in text |
| `Capitalize(s)` / `Titleize(s)` | Capitalize first letter / title case |

### Case Conversion

| Function | Description |
|----------|-------------|
| `CamelCase(s)` | Convert to camelCase |
| `PascalCase(s)` | Convert to PascalCase |
| `SnakeCase(s)` | Convert to snake_case |
| `KebabCase(s)` | Convert to kebab-case |
| `Underscore(s)` | Alias for SnakeCase |
| `Dasherize(s)` | Alias for KebabCase |

## Advanced Features

### Classical Mode

Enable classical/formal pluralization rules:

```go
inflect.ClassicalAll(true)
inflect.Plural("formula")  // "formulae" instead of "formulas"
inflect.Plural("cactus")   // "cacti"

// Fine-grained control
inflect.ClassicalAncient(true)  // Latin/Greek plurals
inflect.ClassicalPersons(true)  // "persons" vs "people"
inflect.ClassicalHerd(true)     // "wildebeest" vs "wildebeests"
inflect.ClassicalZero(true)     // "no error" vs "no errors"
```

### Custom Definitions

Override or extend default behavior:

```go
// Custom noun pluralization
inflect.DefNoun("regex", "regexen")
inflect.Plural("regex")  // "regexen"

// Custom article selection
inflect.DefAn("herb")    // US pronunciation: "an herb"
inflect.An("herb")       // "an herb"

// Pattern-based rules
inflect.DefAPattern("euro.*")  // "a euro", "a european"
```

### Gender for Pronouns

Control third-person singular pronoun resolution:

```go
inflect.Gender("m")
inflect.SingularNoun("they")  // "he"

inflect.Gender("f")
inflect.SingularNoun("they")  // "she"

inflect.Gender("t")  // default: singular they
inflect.SingularNoun("they")  // "they"
```

### Inline Inflection

Parse and inflect text with embedded function calls:

```go
inflect.Inflect("I saw plural('cat', 3)")
// "I saw cats"

inflect.Inflect("This is the ordinal(1) item")
// "This is the 1st item"

inflect.Inflect("plural_noun('I') saw plural_adj('this') plural('cat')")
// "We saw these cats"
```

Supported functions: `plural`, `singular`, `an`/`a`, `ordinal`, `num`,
`plural_noun`, `plural_verb`, `plural_adj`, `singular_noun` (all with optional count parameter).

## Features

- **Case preservation**: Input case is preserved ("CAT" → "CATS", "Child" → "Children")
- **Irregular forms**: Comprehensive irregular noun/verb handling
- **Latin/Greek plurals**: analysis/analyses, cactus/cacti, datum/data
- **Unchanged plurals**: sheep, fish, species, etc.
- **Abbreviations**: Handles acronyms based on pronunciation (FBI, YAML)

## Documentation

- **pkg.go.dev**: [pkg.go.dev/github.com/cv/go-inflect](https://pkg.go.dev/github.com/cv/go-inflect)
- **GitHub Pages**: [lixo.org/go-inflect](https://lixo.org/go-inflect/)

## Development

Run `make` to see available targets:

```
bench                          Run benchmarks
bench-compare                  Compare against baseline
bench-save                     Save benchmark baseline
build                          Build the project
deps                           Download dependencies
help                           Print help message
lint                           Run linter
test                           Run tests with race detection and coverage
```

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## License

MIT License - see [LICENSE](LICENSE)

## Acknowledgments

Go port of the [Python inflect library](https://github.com/jaraco/inflect) by Jason R. Coombs.
