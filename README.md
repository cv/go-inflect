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

### Part-of-Speech Specific Pluralization

For more precise inflection, use part-of-speech specific functions that handle pronouns, verbs, and adjectives.

#### PluralNoun

Pluralizes nouns and pronouns with optional count parameter.

```go
// Pronouns - nominative case
inflect.PluralNoun("I")       // "We"
inflect.PluralNoun("he")      // "they"
inflect.PluralNoun("she")     // "they"
inflect.PluralNoun("it")      // "they"

// Pronouns - accusative case
inflect.PluralNoun("me")      // "us"
inflect.PluralNoun("him")     // "them"
inflect.PluralNoun("her")     // "them"

// Pronouns - possessive
inflect.PluralNoun("my")      // "our"
inflect.PluralNoun("mine")    // "ours"
inflect.PluralNoun("his")     // "their"
inflect.PluralNoun("hers")    // "theirs"

// Pronouns - reflexive
inflect.PluralNoun("myself")   // "ourselves"
inflect.PluralNoun("himself")  // "themselves"
inflect.PluralNoun("herself")  // "themselves"

// Regular nouns (delegates to Plural)
inflect.PluralNoun("cat")     // "cats"
inflect.PluralNoun("child")   // "children"

// With count parameter
inflect.PluralNoun("cat", 1)  // "cat" (singular when count=1)
inflect.PluralNoun("cat", 2)  // "cats"
inflect.PluralNoun("cat", 0)  // "cats"
```

#### PluralVerb

Pluralizes verbs (converts 3rd person singular to base form).

```go
// Auxiliary verbs
inflect.PluralVerb("is")      // "are"
inflect.PluralVerb("was")     // "were"
inflect.PluralVerb("has")     // "have"
inflect.PluralVerb("does")    // "do"

// Contractions
inflect.PluralVerb("isn't")   // "aren't"
inflect.PluralVerb("wasn't")  // "weren't"
inflect.PluralVerb("doesn't") // "don't"
inflect.PluralVerb("hasn't")  // "haven't"

// Modal verbs (unchanged)
inflect.PluralVerb("can")     // "can"
inflect.PluralVerb("will")    // "will"
inflect.PluralVerb("must")    // "must"

// Regular 3rd person singular verbs
inflect.PluralVerb("runs")    // "run"
inflect.PluralVerb("walks")   // "walk"
inflect.PluralVerb("tries")   // "try"
inflect.PluralVerb("watches") // "watch"

// With count parameter
inflect.PluralVerb("is", 1)   // "is" (singular when count=1)
inflect.PluralVerb("is", 2)   // "are"
```

#### PluralAdj

Pluralizes adjectives (demonstratives, articles, possessives).

```go
// Demonstrative adjectives
inflect.PluralAdj("this")     // "these"
inflect.PluralAdj("that")     // "those"

// Indefinite articles
inflect.PluralAdj("a")        // "some"
inflect.PluralAdj("an")       // "some"

// Possessive adjectives
inflect.PluralAdj("my")       // "our"
inflect.PluralAdj("his")      // "their"
inflect.PluralAdj("her")      // "their"
inflect.PluralAdj("its")      // "their"

// With count parameter
inflect.PluralAdj("this", 1)  // "this" (singular when count=1)
inflect.PluralAdj("this", 2)  // "these"
```

#### SingularNoun

Singularizes nouns and pronouns with gender support.

```go
// Pronouns - nominative case
inflect.SingularNoun("we")    // "I"
inflect.SingularNoun("they")  // "they" (default gender is singular they)

// Pronouns - accusative case
inflect.SingularNoun("us")    // "me"
inflect.SingularNoun("them")  // "them" (default gender)

// Pronouns - possessive
inflect.SingularNoun("our")   // "my"
inflect.SingularNoun("ours")  // "mine"
inflect.SingularNoun("their") // "their" (default gender)

// Pronouns - reflexive
inflect.SingularNoun("ourselves")   // "myself"
inflect.SingularNoun("themselves")  // "themself" (default gender)

// Regular nouns (delegates to Singular)
inflect.SingularNoun("cats")     // "cat"
inflect.SingularNoun("children") // "child"

// With count parameter
inflect.SingularNoun("cats", 1)  // "cat"
inflect.SingularNoun("cats", 2)  // "cats" (plural when count!=1)
```

### Gender for Pronouns

Set the gender for third-person singular pronouns returned by `SingularNoun()`.

```go
// Default is "t" (singular they)
inflect.SingularNoun("they")  // "they"
inflect.SingularNoun("them")  // "them"

// Masculine
inflect.Gender("m")
inflect.SingularNoun("they")  // "he"
inflect.SingularNoun("them")  // "him"
inflect.SingularNoun("their") // "his"

// Feminine
inflect.Gender("f")
inflect.SingularNoun("they")  // "she"
inflect.SingularNoun("them")  // "her"
inflect.SingularNoun("their") // "her"

// Neuter
inflect.Gender("n")
inflect.SingularNoun("they")  // "it"
inflect.SingularNoun("them")  // "it"
inflect.SingularNoun("their") // "its"

// Singular they (default)
inflect.Gender("t")
inflect.SingularNoun("they")  // "they"

// Check current gender
inflect.GetGender()  // returns "m", "f", "n", or "t"
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

// Convert word numbers to ordinals
inflect.WordToOrdinal("one")         // "first"
inflect.WordToOrdinal("twenty-one")  // "twenty-first"
inflect.WordToOrdinal("One")         // "First" (preserves case)
inflect.WordToOrdinal("TWENTY")      // "TWENTIETH"

// Numeric strings also work
inflect.WordToOrdinal("1")    // "1st"
inflect.WordToOrdinal("21")   // "21st"

// Get just the ordinal suffix
inflect.OrdinalSuffix(1)   // "st"
inflect.OrdinalSuffix(2)   // "nd"
inflect.OrdinalSuffix(3)   // "rd"
inflect.OrdinalSuffix(11)  // "th" (teens are special)
inflect.OrdinalSuffix(21)  // "st"

// Check if a string is an ordinal
inflect.IsOrdinal("1st")          // true
inflect.IsOrdinal("first")        // true
inflect.IsOrdinal("twenty-first") // true
inflect.IsOrdinal("one")          // false
inflect.IsOrdinal("42")           // false

// Convert ordinals back to cardinals
inflect.OrdinalToCardinal("1st")          // "1"
inflect.OrdinalToCardinal("first")        // "one"
inflect.OrdinalToCardinal("twenty-first") // "twenty-one"
inflect.OrdinalToCardinal("First")        // "One" (preserves case)
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

// British English style with "and"
inflect.NumberToWordsWithAnd(101)   // "one hundred and one"
inflect.NumberToWordsWithAnd(1001)  // "one thousand and one"
inflect.NumberToWordsWithAnd(1101)  // "one thousand one hundred and one"
inflect.NumberToWordsWithAnd(1234)  // "one thousand two hundred and thirty-four"
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

// Auto-detect separator based on content
inflect.JoinWithAutoSep([]string{"a", "b", "c"}, "and")
// "a, b, and c" (no commas in items, uses comma)
inflect.JoinWithAutoSep([]string{"Jan 1, 2020", "Feb 2, 2021"}, "and")
// "Jan 1, 2020; and Feb 2, 2021" (commas in items, uses semicolon)

// Different final separator (before conjunction)
inflect.JoinWithFinalSep([]string{"a", "b", "c"}, "and", ", ", "; ")
// "a, b; and c"

// Without Oxford comma (British style)
inflect.JoinNoOxford([]string{"a", "b", "c"})
// "a, b and c"
inflect.JoinNoOxfordWithConj([]string{"a", "b", "c"}, "or")
// "a, b or c"
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

### Past Participle

Convert verbs to their past participle form.

```go
// Regular verbs (-ed)
inflect.PastParticiple("walk")    // "walked"
inflect.PastParticiple("play")    // "played"
inflect.PastParticiple("stop")    // "stopped" (double consonant)
inflect.PastParticiple("try")     // "tried" (y -> ied)
inflect.PastParticiple("like")    // "liked" (just add -d)

// Irregular verbs
inflect.PastParticiple("go")      // "gone"
inflect.PastParticiple("take")    // "taken"
inflect.PastParticiple("run")     // "run" (unchanged)
inflect.PastParticiple("write")   // "written"
inflect.PastParticiple("think")   // "thought"
inflect.PastParticiple("buy")     // "bought"

// Case preservation
inflect.PastParticiple("Go")      // "Gone"
inflect.PastParticiple("WALK")    // "WALKED"

// Check if a word is a participle
inflect.IsParticiple("running")   // true (present participle)
inflect.IsParticiple("walked")    // true (past participle)
inflect.IsParticiple("taken")     // true (irregular past participle)
inflect.IsParticiple("walk")      // false (base verb)
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

// Compare verbs (3rd person singular vs base form)
inflect.CompareVerbs("runs", "run")     // "s:p"
inflect.CompareVerbs("is", "are")       // "s:p"
inflect.CompareVerbs("has", "have")     // "s:p"
inflect.CompareVerbs("run", "runs")     // "p:s"
inflect.CompareVerbs("doesn't", "don't") // "s:p"

// Compare adjectives (demonstratives, articles, possessives)
inflect.CompareAdjs("this", "these")    // "s:p"
inflect.CompareAdjs("that", "those")    // "s:p"
inflect.CompareAdjs("a", "some")        // "s:p"
inflect.CompareAdjs("my", "our")        // "s:p"
inflect.CompareAdjs("these", "this")    // "p:s"
```

**Return values:**
- `"eq"` - Words are equal (case-insensitive)
- `"s:p"` - First word is singular, second is its plural
- `"p:s"` - First word is plural, second is its singular
- `"p:p"` - Both words are different plural forms of the same word
- `""` - Words are not related

## API Reference

### Core Functions

| Function | Description |
|----------|-------------|
| `An(word string) string` | Returns word prefixed with "a" or "an" |
| `A(word string) string` | Alias for `An()` |
| `Plural(word string) string` | Returns the plural form of a noun |
| `Singular(word string) string` | Returns the singular form of a noun |

### Part-of-Speech Specific Functions

| Function | Description |
|----------|-------------|
| `PluralNoun(word string, count ...int) string` | Pluralizes nouns and pronouns |
| `PluralVerb(word string, count ...int) string` | Pluralizes verbs (3rd person → base form) |
| `PluralAdj(word string, count ...int) string` | Pluralizes adjectives (this→these, a→some) |
| `SingularNoun(word string, count ...int) string` | Singularizes nouns and pronouns |

### Gender Control

| Function | Description |
|----------|-------------|
| `Gender(g string)` | Sets gender for 3rd person pronouns ("m", "f", "n", "t") |
| `GetGender() string` | Returns current gender setting |

### Number Functions

| Function | Description |
|----------|-------------|
| `Ordinal(n int) string` | Returns numeric ordinal (1st, 2nd, 3rd, ...) |
| `OrdinalWord(n int) string` | Returns ordinal as word (first, second, ...) |
| `OrdinalSuffix(n int) string` | Returns just the ordinal suffix (st, nd, rd, th) |
| `WordToOrdinal(s string) string` | Converts word/numeric string to ordinal |
| `IsOrdinal(s string) bool` | Checks if a string is an ordinal |
| `OrdinalToCardinal(s string) string` | Converts ordinal to cardinal (first → one) |
| `NumberToWords(n int) string` | Converts integer to English words |
| `NumberToWordsWithAnd(n int) string` | Converts integer to words (British style with "and") |

### List Functions

| Function | Description |
|----------|-------------|
| `Join(words []string) string` | Joins list with "and" and Oxford comma |
| `JoinWithConj(words []string, conj string) string` | Joins with custom conjunction |
| `JoinWithSep(words []string, conj, sep string) string` | Joins with custom separator |
| `JoinWithFinalSep(words []string, conj, sep, finalSep string) string` | Joins with different final separator |
| `JoinWithAutoSep(words []string, conj string) string` | Auto-detects separator based on content |
| `JoinNoOxford(words []string) string` | Joins without Oxford comma |
| `JoinNoOxfordWithConj(words []string, conj string) string` | Joins without Oxford comma, custom conjunction |

### Verb Functions

| Function | Description |
|----------|-------------|
| `PresentParticiple(verb string) string` | Returns present participle (-ing form) |
| `PastParticiple(verb string) string` | Returns past participle (walked, gone, taken) |
| `IsParticiple(word string) bool` | Checks if a word is a participle |

### Comparison Functions

| Function | Description |
|----------|-------------|
| `Compare(word1, word2 string) string` | Compares singular/plural relationship |
| `CompareNouns(noun1, noun2 string) string` | Alias for `Compare()` |
| `CompareVerbs(verb1, verb2 string) string` | Compares verb forms (3rd person vs base) |
| `CompareAdjs(adj1, adj2 string) string` | Compares adjective forms (this/these, a/some) |

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
