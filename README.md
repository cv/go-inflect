# go-inflect

A Go port of the Python [inflect](https://github.com/jaraco/inflect) library for English word inflection.

- **100+ functions** — pluralization, articles, numbers, verbs, and more
- **Case-preserving** — "CAT" → "CATS", "Child" → "Children"
- **Thread-safe** — all functions safe for concurrent use
- **Template-ready** — `FuncMap()` for `text/template` and `html/template`
- **Minimal dependencies** — only `golang.org/x/text` for Unicode normalization
- **Drop-in compatible** with `jinzhu/inflection` and `go-openapi/inflect`

## Installation

```bash
go get github.com/cv/go-inflect/v2
```

## Quick Start

The easiest way to use go-inflect is with Go templates:

```go
import (
    "os"
    "text/template"
    "github.com/cv/go-inflect/v2"
)

func main() {
    tmpl := template.New("report").Funcs(inflect.FuncMap())
    tmpl.Parse(`
📊 Build Report for {{.Project}}

{{.Tests}} {{plural "test" .Tests}} {{pastTense "run"}} in {{.Duration}} {{plural "second" .Duration}}.
Found {{no "error" .Errors}} and {{no "warning" .Warnings}}.

This is the {{ordinalWord .BuildNum}} build, and it {{pastTense "complete"}} {{comparative "fast"}}
than {{numberToWords .Improvement}}% of previous builds — {{possessive .Author}} {{superlative "good"}} yet!

Summary: {{join .Tags}}
`)
    tmpl.Execute(os.Stdout, map[string]any{
        "Project":     "go-inflect",
        "Tests":       142,
        "Duration":    3,
        "Errors":      0,
        "Warnings":    2,
        "BuildNum":    47,
        "Improvement": 85,
        "Author":      "James",
        "Tags":        []string{"fast", "stable", "passing"},
    })
}
// Output:
// 📊 Build Report for go-inflect
// 
// 142 tests ran in 3 seconds.
// Found no errors and 2 warnings.
// 
// This is the forty-seventh build, and it completed faster
// than eighty-five% of previous builds — James's best yet!
// 
// Summary: fast, stable, and passing
```

50+ template functions available — see [FuncMap documentation](https://pkg.go.dev/github.com/cv/go-inflect/v2#FuncMap) for the complete list.

### Direct Function Calls

All functions are also available for direct use:

```go
// Pluralization handles irregular forms
inflect.Plural("person")             // "people"
inflect.Plural("criterion")          // "criteria"  
inflect.Singular("phenomena")        // "phenomenon"

// Verb conjugation
inflect.PastTense("run")             // "ran"
inflect.PastParticiple("write")      // "written"
inflect.PresentParticiple("swim")    // "swimming"

// Numbers in words
inflect.NumberToWords(42)            // "forty-two"
inflect.OrdinalWord(3)               // "third"
inflect.FractionToWords(3, 4)        // "three quarters"
inflect.CurrencyToWords(99.99, "USD") // "ninety-nine dollars and ninety-nine cents"

// Adjectives and adverbs
inflect.Comparative("beautiful")     // "more beautiful"
inflect.Superlative("fast")          // "fastest"
inflect.Adverb("quick")              // "quickly"

// Handy utilities
inflect.An("hour")                   // "an hour"
inflect.Join([]string{"a","b","c"})  // "a, b, and c"
inflect.No("error", 0)               // "no errors"
inflect.Possessive("James")          // "James's"
```

See [pkg.go.dev](https://pkg.go.dev/github.com/cv/go-inflect/v2) for the complete API.

## Custom Engine

For isolated configurations (different classical modes, custom definitions), create a separate engine:

```go
// Classical Latin/Greek plurals
classical := inflect.NewEngine()
classical.Classical(true)
classical.Plural("formula")  // "formulae"
classical.Plural("cactus")   // "cacti"

// Modern English (default)
modern := inflect.NewEngine()
modern.Plural("formula")     // "formulas"

// Custom definitions
eng := inflect.NewEngine()
eng.DefNoun("regex", "regexen")
eng.DefNoun("pokemon", "pokemon")

// Use custom engine with templates
tmpl := template.New("custom").Funcs(eng.FuncMap())
```

## Migration from jinzhu/inflection

Core functions work identically:

```go
inflect.Plural("cat")       // "cats"
inflect.Singular("cats")    // "cat"
inflect.Underscore("Hello") // "hello"
```

Compatibility aliases: `Pluralize`, `Singularize`, `Camelize`, `CamelizeDownFirst`, `AddIrregular`, `AddUncountable`.

## Documentation

Full API documentation: [pkg.go.dev/github.com/cv/go-inflect/v2](https://pkg.go.dev/github.com/cv/go-inflect/v2)

## License

MIT — see [LICENSE](LICENSE)

## Acknowledgments

- Port of [Python inflect](https://github.com/jaraco/inflect) by Jason R. Coombs
- Compatibility APIs from [jinzhu/inflection](https://github.com/jinzhu/inflection) and [go-openapi/inflect](https://github.com/go-openapi/inflect)
- Rails-style helpers inspired by [ActiveSupport::Inflector](https://api.rubyonrails.org/classes/ActiveSupport/Inflector.html)
