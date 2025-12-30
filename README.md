# go-inflect

A Go port of the Python [inflect](https://github.com/jaraco/inflect) library for English word inflection.

- **100+ functions** — pluralization, articles, numbers, verbs, and more
- **Case-preserving** — "CAT" → "CATS", "Child" → "Children"
- **Thread-safe** — all functions safe for concurrent use
- **Minimal dependencies** — only `golang.org/x/text` for Unicode normalization
- **Extensively tested** — benchmarks, fuzz tests, and high coverage
- **Drop-in compatible** with `jinzhu/inflection` and `go-openapi/inflect`

## Installation

```bash
go get github.com/cv/go-inflect
```

## Quick Start

```go
import "github.com/cv/go-inflect"

// Articles
inflect.An("apple")       // "an apple"
inflect.An("hour")        // "an hour" (silent h)

// Nouns
inflect.Plural("child")   // "children"
inflect.Singular("mice")  // "mouse"

// Verbs
inflect.PastTense("go")           // "went"
inflect.PresentParticiple("run")  // "running"

// Numbers
inflect.NumberToWords(42)  // "forty-two"
inflect.Ordinal(3)         // "3rd"

// Lists
inflect.Join([]string{"a", "b", "c"})  // "a, b, and c"

// Adjectives
inflect.Comparative("big")     // "bigger"
inflect.Superlative("big")     // "biggest"

// And more: possessives, fractions, currency, case conversion...
```

## API Reference

### Nouns

| Function | Example |
|----------|---------|
| `Plural(word)` | "cat" → "cats", "child" → "children" |
| `Singular(word)` | "cats" → "cat", "mice" → "mouse" |
| `PluralNoun(word, count...)` | Pluralize with optional count |
| `SingularNoun(word, count...)` | Singularize with optional count |

### Verbs

| Function | Example |
|----------|---------|
| `PluralVerb(word, count...)` | "is" → "are", "runs" → "run" |
| `PresentParticiple(verb)` | "run" → "running" |
| `PastParticiple(verb)` | "take" → "taken" |
| `PastTense(verb)` | "walk" → "walked", "go" → "went" |

### Articles & Adjectives

| Function | Example |
|----------|---------|
| `An(word)` / `A(word)` | "an apple", "a banana" |
| `PluralAdj(word, count...)` | "this" → "these" |
| `Comparative(adj)` | "big" → "bigger" |
| `Superlative(adj)` | "big" → "biggest" |
| `Adverb(adj)` | "quick" → "quickly" |
| `Possessive(noun)` | "cat" → "cat's", "cats" → "cats'" |

### Numbers

| Function | Example |
|----------|---------|
| `NumberToWords(n)` | 42 → "forty-two" |
| `NumberToWordsWithAnd(n)` | 101 → "one hundred and one" |
| `NumberToWordsFloat(f)` | 3.14 → "three point one four" |
| `FormatNumber(n)` | 1000 → "1,000" |
| `Ordinal(n)` | 1 → "1st" |
| `OrdinalWord(n)` | 1 → "first" |
| `FractionToWords(num, den)` | (3, 4) → "three quarters" |
| `CurrencyToWords(amt, code)` | (1.50, "USD") → "one dollar and fifty cents" |
| `CountingWord(n)` | 1 → "once", 2 → "twice" |

### Lists

| Function | Example |
|----------|---------|
| `Join(words)` | ["a", "b", "c"] → "a, b, and c" |
| `JoinWithConj(words, conj)` | Custom conjunction |
| `JoinNoOxford(words)` | British style (no Oxford comma) |

### Comparison

| Function | Description |
|----------|-------------|
| `Compare(w1, w2)` | Returns "eq", "s:p", "p:s", "p:p", or "" |
| `CompareNouns(n1, n2)` | Compare noun forms |
| `CompareVerbs(v1, v2)` | Compare verb forms |

### Case Conversion

| Function | Example |
|----------|---------|
| `CamelCase(s)` | "hello_world" → "helloWorld" |
| `PascalCase(s)` | "hello_world" → "HelloWorld" |
| `SnakeCase(s)` | "HelloWorld" → "hello_world" |
| `KebabCase(s)` | "HelloWorld" → "hello-world" |

### Rails-Style Helpers

| Function | Example |
|----------|---------|
| `Humanize(word)` | "employee_salary" → "Employee salary" |
| `Tableize(word)` | "RawScaledScorer" → "raw_scaled_scorers" |
| `Parameterize(word)` | "Hello World!" → "hello-world" |
| `ForeignKey(word)` | "Person" → "person_id" |
| `Typeify(word)` | "users" → "User" |
| `Asciify(word)` | "café" → "cafe" |

### Utilities

| Function | Description |
|----------|-------------|
| `IsPlural(word)` / `IsSingular(word)` | Check word form |
| `WordCount(text)` | Count words in text |
| `Capitalize(s)` / `Titleize(s)` | Case helpers |

## Advanced Features

### Classical Mode

Enable classical/formal pluralization:

```go
inflect.ClassicalAll(true)
inflect.Plural("formula")  // "formulae"
inflect.Plural("cactus")   // "cacti"

// Fine-grained control
inflect.ClassicalAncient(true)  // Latin/Greek plurals
inflect.ClassicalPersons(true)  // "persons" vs "people"
```

### Custom Definitions

```go
inflect.DefNoun("regex", "regexen")
inflect.DefAn("herb")  // US pronunciation
```

### Isolated Configurations

Package functions share a default engine. For isolated settings, create separate engines:

```go
classical := inflect.NewEngine()
classical.Classical(true)
classical.Plural("formula")  // "formulae"

modern := inflect.NewEngine()
modern.Plural("formula")     // "formulas"
```

### Gender for Pronouns

```go
inflect.Gender("m")
inflect.SingularNoun("they")  // "he"

inflect.Gender("f")
inflect.SingularNoun("they")  // "she"
```

## Migration

### From jinzhu/inflection or go-openapi/inflect

Core functions work identically:

```go
inflect.Plural("cat")       // "cats"
inflect.Singular("cats")    // "cat"
inflect.Underscore("Hello") // "hello"
```

Compatibility aliases provided:

| Alias | Target |
|-------|--------|
| `Pluralize` | `Plural` |
| `Singularize` | `Singular` |
| `Camelize` | `PascalCase` |
| `CamelizeDownFirst` | `CamelCase` |
| `AddIrregular` | `DefNoun` |
| `AddUncountable` | marks word as uncountable |

## Development

```bash
make help  # see all targets
```

See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## Documentation

- [pkg.go.dev/github.com/cv/go-inflect](https://pkg.go.dev/github.com/cv/go-inflect)

## License

MIT — see [LICENSE](LICENSE)

## Acknowledgments

- Port of [Python inflect](https://github.com/jaraco/inflect) by Jason R. Coombs
- Compatibility APIs from [jinzhu/inflection](https://github.com/jinzhu/inflection) and [go-openapi/inflect](https://github.com/go-openapi/inflect)
- Rails-style helpers inspired by [ActiveSupport::Inflector](https://api.rubyonrails.org/classes/ActiveSupport/Inflector.html)
- Written with [Claude](https://claude.ai) and [Nemotron](https://developer.nvidia.com/nemotron)
