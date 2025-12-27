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

// Verb participles
inflect.PresentParticiple("run")  // "running"
inflect.PastParticiple("take")    // "taken"

// Word comparison
inflect.Compare("cat", "cats")    // "s:p" (singular to plural)
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

### Articles & Adjectives

| Function | Description |
|----------|-------------|
| `An(word)` / `A(word)` | Prefix with "a" or "an" |
| `PluralAdj(word, count...)` | "this" → "these", "a" → "some" |

### Numbers

| Function | Description |
|----------|-------------|
| `NumberToWords(n)` | 42 → "forty-two" |
| `NumberToWordsWithAnd(n)` | 101 → "one hundred and one" |
| `Ordinal(n)` | 1 → "1st", 2 → "2nd" |
| `OrdinalWord(n)` | 1 → "first", 2 → "second" |
| `No(word, count)` | "no errors" or "3 errors" |

### Lists

| Function | Description |
|----------|-------------|
| `Join(words)` | Join with "and" and Oxford comma |
| `JoinWithConj(words, conj)` | Custom conjunction ("or", "and/or") |
| `JoinNoOxford(words)` | British style without Oxford comma |

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
| `Capitalize(s)` / `Titleize(s)` | Case conversion |

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
```

## Features

- **Case preservation**: Input case is preserved ("CAT" → "CATS", "Child" → "Children")
- **Irregular forms**: Comprehensive irregular noun/verb handling
- **Latin/Greek plurals**: analysis/analyses, cactus/cacti, datum/data
- **Unchanged plurals**: sheep, fish, species, etc.
- **Abbreviations**: Handles acronyms based on pronunciation (FBI, YAML)

## Documentation

Full API documentation with examples: [pkg.go.dev/github.com/cv/go-inflect](https://pkg.go.dev/github.com/cv/go-inflect)

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## License

MIT License - see [LICENSE](LICENSE)

## Acknowledgments

Go port of the [Python inflect library](https://github.com/jaraco/inflect) by Jason R. Coombs.
