# go-inflect

A Go port of the [Python inflect](https://github.com/jaraco/inflect) library for English language inflection utilities.

`go-inflect` provides functions for pluralization, singularization, indefinite article selection (a/an), number-to-words conversion, ordinals, verb participles, and more.

## Installation

```bash
go get github.com/cv/go-inflect
```

## Usage

```go
import "github.com/cv/go-inflect"
```

### Indefinite Articles (An/A)

Select the appropriate indefinite article ("a" or "an") for a word based on its pronunciation.

```go
inflect.An("apple")      // "an apple"
inflect.An("banana")     // "a banana"
inflect.An("hour")       // "an hour" (silent 'h')
inflect.An("university") // "a university" (sounds like "yoo")
inflect.An("FBI agent")  // "an FBI agent" (F = "eff")
inflect.An("YAML file")  // "a YAML file" (Y = "why")

// A() is an alias for An()
inflect.A("cat")  // "a cat"
```

**Features:**
- Handles silent 'h' words: honest, heir, honor, hour
- Handles vowels with consonant sounds: university, unanimous, European
- Handles abbreviations/acronyms based on letter pronunciation

### Pluralization

Convert singular nouns to their plural forms.

```go
// Regular nouns
inflect.Plural("cat")      // "cats"
inflect.Plural("dog")      // "dogs"

// Sibilant endings (-s, -sh, -ch, -x, -z)
inflect.Plural("box")      // "boxes"
inflect.Plural("church")   // "churches"
inflect.Plural("buzz")     // "buzzes"

// Consonant + y
inflect.Plural("city")     // "cities"
inflect.Plural("baby")     // "babies"

// -f/-fe to -ves
inflect.Plural("knife")    // "knives"
inflect.Plural("wolf")     // "wolves"
inflect.Plural("leaf")     // "leaves"

// Irregular plurals
inflect.Plural("child")    // "children"
inflect.Plural("mouse")    // "mice"
inflect.Plural("person")   // "people"
inflect.Plural("foot")     // "feet"
inflect.Plural("analysis") // "analyses"
inflect.Plural("cactus")   // "cacti"
inflect.Plural("datum")    // "data"

// Unchanged plurals
inflect.Plural("sheep")    // "sheep"
inflect.Plural("fish")     // "fish"
inflect.Plural("species")  // "species"

// Case preservation
inflect.Plural("CAT")      // "CATS"
inflect.Plural("Child")    // "Children"
```

### Singularization

Convert plural nouns to their singular forms.

```go
// Regular nouns
inflect.Singular("cats")     // "cat"
inflect.Singular("dogs")     // "dog"

// Sibilant endings
inflect.Singular("boxes")    // "box"
inflect.Singular("churches") // "church"

// -ies to -y
inflect.Singular("cities")   // "city"
inflect.Singular("babies")   // "baby"

// -ves to -f/-fe
inflect.Singular("knives")   // "knife"
inflect.Singular("wolves")   // "wolf"

// Irregular plurals
inflect.Singular("children") // "child"
inflect.Singular("mice")     // "mouse"
inflect.Singular("people")   // "person"
inflect.Singular("feet")     // "foot"
inflect.Singular("analyses") // "analysis"
inflect.Singular("cacti")    // "cactus"

// Unchanged plurals
inflect.Singular("sheep")    // "sheep"
inflect.Singular("fish")     // "fish"
```

### Ordinals

Convert numbers to ordinal form (numeric or word).

```go
// Numeric ordinals
inflect.Ordinal(1)   // "1st"
inflect.Ordinal(2)   // "2nd"
inflect.Ordinal(3)   // "3rd"
inflect.Ordinal(4)   // "4th"
inflect.Ordinal(11)  // "11th"
inflect.Ordinal(12)  // "12th"
inflect.Ordinal(13)  // "13th"
inflect.Ordinal(21)  // "21st"
inflect.Ordinal(22)  // "22nd"
inflect.Ordinal(23)  // "23rd"
inflect.Ordinal(100) // "100th"
inflect.Ordinal(101) // "101st"
inflect.Ordinal(-1)  // "-1st"

// Word ordinals
inflect.OrdinalWord(1)   // "first"
inflect.OrdinalWord(2)   // "second"
inflect.OrdinalWord(3)   // "third"
inflect.OrdinalWord(11)  // "eleventh"
inflect.OrdinalWord(12)  // "twelfth"
inflect.OrdinalWord(21)  // "twenty-first"
inflect.OrdinalWord(100) // "one hundredth"
inflect.OrdinalWord(101) // "one hundred first"
inflect.OrdinalWord(-5)  // "negative fifth"
```

### Number to Words

Convert integers to their English word representation.

```go
inflect.NumberToWords(0)        // "zero"
inflect.NumberToWords(1)        // "one"
inflect.NumberToWords(12)       // "twelve"
inflect.NumberToWords(42)       // "forty-two"
inflect.NumberToWords(100)      // "one hundred"
inflect.NumberToWords(123)      // "one hundred twenty-three"
inflect.NumberToWords(1000)     // "one thousand"
inflect.NumberToWords(1234)     // "one thousand two hundred thirty-four"
inflect.NumberToWords(1000000)  // "one million"
inflect.NumberToWords(-5)       // "negative five"
```

### List Joining

Join lists of words with proper English grammar (Oxford comma).

```go
// Basic joining with "and"
inflect.Join([]string{})                    // ""
inflect.Join([]string{"apples"})            // "apples"
inflect.Join([]string{"apples", "oranges"}) // "apples and oranges"
inflect.Join([]string{"apples", "oranges", "bananas"})
// "apples, oranges, and bananas"

// Custom conjunction
inflect.JoinWithConj([]string{"tea", "coffee"}, "or")
// "tea or coffee"

inflect.JoinWithConj([]string{"tea", "coffee", "juice"}, "or")
// "tea, coffee, or juice"

inflect.JoinWithConj([]string{"read", "write", "execute"}, "and/or")
// "read, write, and/or execute"

// Custom separator (useful when items contain commas)
inflect.JoinWithSep([]string{"Jan 1, 2020", "Feb 2, 2021", "Mar 3, 2022"}, "and", "; ")
// "Jan 1, 2020; Feb 2, 2021; and Mar 3, 2022"
```

### Present Participle

Convert verbs to their present participle (-ing) form.

```go
// Simple verbs (just add -ing)
inflect.PresentParticiple("play")   // "playing"
inflect.PresentParticiple("walk")   // "walking"
inflect.PresentParticiple("eat")    // "eating"

// Drop silent 'e'
inflect.PresentParticiple("make")   // "making"
inflect.PresentParticiple("take")   // "taking"
inflect.PresentParticiple("write")  // "writing"

// Double final consonant (CVC pattern)
inflect.PresentParticiple("run")    // "running"
inflect.PresentParticiple("sit")    // "sitting"
inflect.PresentParticiple("stop")   // "stopping"
inflect.PresentParticiple("begin")  // "beginning"

// -ie becomes -ying
inflect.PresentParticiple("die")    // "dying"
inflect.PresentParticiple("lie")    // "lying"
inflect.PresentParticiple("tie")    // "tying"

// Keep -ee
inflect.PresentParticiple("see")    // "seeing"
inflect.PresentParticiple("flee")   // "fleeing"

// Add -k before -ing after -c
inflect.PresentParticiple("panic")  // "panicking"
inflect.PresentParticiple("picnic") // "picnicking"
```

### Word Comparison

Compare words to determine their singular/plural relationship.

```go
// Equal words
inflect.Compare("cat", "cat")       // "eq"
inflect.Compare("Cat", "CAT")       // "eq" (case-insensitive)

// Singular to plural
inflect.Compare("cat", "cats")      // "s:p"
inflect.Compare("child", "children") // "s:p"
inflect.Compare("mouse", "mice")    // "s:p"

// Plural to singular
inflect.Compare("cats", "cat")      // "p:s"
inflect.Compare("children", "child") // "p:s"

// Both are plural forms
inflect.Compare("indexes", "indices") // "p:p"

// Unrelated words
inflect.Compare("cat", "dog")       // ""

// CompareNouns is an alias for Compare
inflect.CompareNouns("dog", "dogs") // "s:p"
```

**Return values:**
- `"eq"` - Words are equal (case-insensitive)
- `"s:p"` - First word is singular, second is its plural
- `"p:s"` - First word is plural, second is its singular
- `"p:p"` - Both words are different plural forms of the same word
- `""` - Words are not related

## API Reference

| Function | Description |
|----------|-------------|
| `An(word string) string` | Returns word prefixed with "a" or "an" |
| `A(word string) string` | Alias for `An()` |
| `Plural(word string) string` | Returns the plural form of a noun |
| `Singular(word string) string` | Returns the singular form of a noun |
| `Ordinal(n int) string` | Returns numeric ordinal (1st, 2nd, 3rd, ...) |
| `OrdinalWord(n int) string` | Returns ordinal as word (first, second, ...) |
| `NumberToWords(n int) string` | Converts integer to English words |
| `Join(words []string) string` | Joins list with "and" and Oxford comma |
| `JoinWithConj(words []string, conj string) string` | Joins with custom conjunction |
| `JoinWithSep(words []string, conj string, sep string) string` | Joins with custom conjunction and separator |
| `PresentParticiple(verb string) string` | Returns present participle (-ing form) |
| `Compare(word1, word2 string) string` | Compares singular/plural relationship |
| `CompareNouns(noun1, noun2 string) string` | Alias for `Compare()` |

## Features

- **Case preservation**: Functions preserve the case pattern of input words (uppercase, title case, lowercase)
- **Irregular forms**: Extensive support for irregular plurals (child/children, mouse/mice, etc.)
- **Latin/Greek plurals**: Handles classical plurals (analysis/analyses, cactus/cacti, datum/data)
- **Unchanged plurals**: Recognizes words that don't change (sheep, fish, species, etc.)
- **Oxford comma**: List joining uses the serial comma for clarity
- **Negative numbers**: Number functions handle negative values

## Contributing

Contributions are welcome! Please follow these guidelines:

1. **Fork the repository** and create a feature branch
2. **Write tests** for any new functionality
3. **Run the test suite** before submitting:
   ```bash
   go test -v ./...
   ```
4. **Follow Go conventions**:
   - Run `go fmt` on your code
   - Run `go vet` to check for issues
   - Add documentation comments for exported functions
5. **Submit a merge request** with a clear description of changes

### Reporting Issues

When reporting bugs, please include:
- Go version (`go version`)
- Input that causes the issue
- Expected output
- Actual output

### Adding New Features

If you'd like to add features from the Python inflect library that aren't yet implemented:

1. Check the [Python inflect source](https://github.com/jaraco/inflect) for reference
2. Open an issue to discuss the feature before implementation
3. Include comprehensive tests covering edge cases

## License

See [LICENSE](LICENSE) for details.

## Acknowledgments

This library is a Go port of the excellent [Python inflect library](https://github.com/jaraco/inflect) by Jason R. Coombs.
