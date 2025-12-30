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
    tmpl := template.New("example").Funcs(inflect.FuncMap())
    tmpl.Parse(`
You have {{.Count}} {{plural "message" .Count}}.
That's {{numberToWords .Count}} {{plural "message" .Count}}!
The {{ordinalWord .Position}} message is from {{possessive .Name}} account.
`)
    tmpl.Execute(os.Stdout, map[string]any{
        "Count":    5,
        "Position": 1, 
        "Name":     "James",
    })
}
// Output:
// You have 5 messages.
// That's five messages!
// The first message is from James's account.
```

50+ template functions available: `plural`, `singular`, `an`, `ordinal`, `ordinalWord`, `numberToWords`, `pastTense`, `presentParticiple`, `possessive`, `comparative`, `superlative`, `join`, `camelCase`, `snakeCase`, `tableize`, `humanize`, and [many more](https://pkg.go.dev/github.com/cv/go-inflect/v2#FuncMap).

### Direct Function Calls

All functions are also available for direct use:

```go
inflect.Plural("child")           // "children"
inflect.Singular("mice")          // "mouse"
inflect.An("apple")               // "an apple"
inflect.NumberToWords(42)         // "forty-two"
inflect.Ordinal(3)                // "3rd"
inflect.PastTense("go")           // "went"
inflect.Possessive("James")       // "James's"
inflect.Join([]string{"a","b","c"}) // "a, b, and c"
inflect.Comparative("big")        // "bigger"
inflect.CamelCase("hello_world")  // "helloWorld"
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
