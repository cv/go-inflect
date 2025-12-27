# go-inflect Roadmap

This is a port of Python's `inflect` library to Go. The library provides tools for generating plurals, singulars, ordinals, indefinite articles, and converting numbers to words.

## Progress Summary

Run `bd list` to see all tasks, or `bd list --type epic` to see just the epics.

## Feature Epics

### 1. Pluralization (go-inflect-tw5)
Core pluralization functionality:
- `plural()` - unconditional pluralization
- `plural_noun()` - pluralize nouns
- `plural_verb()` - pluralize verbs
- `plural_adj()` - pluralize adjectives
- 860+ test cases from inflections.txt
- Compound word support (hyphenated & spaced)
- Unit handling (degree celsius, per square inch, etc)
- Open compound nouns (high school, prima donna)

### 2. Singularization (go-inflect-tn5)
Converting plurals back to singular form:
- `singular_noun()` - convert plural to singular
- Comprehensive test coverage with all plural forms

### 3. Indefinite Articles (go-inflect-cew)
Smart a/an selection:
- `a()` and `an()` functions
- Common words (cat → "a cat", ant → "an ant")
- Abbreviations (YAML, JSON, Core ML, etc)
- Silent letters and special cases (honest → "an honest")

### 4. Number Conversion (go-inflect-upi)
Convert numbers to English words:
- `number_to_words()` core function
- Basic numbers (0-20, 21-99, 100s, 1000s, millions, etc)
- Decimal number support
- Threshold parameter (numbers above threshold stay numeric)
- Comma formatting
- Group and decimal parameters
- 100+ test cases

### 5. Ordinals (go-inflect-6t7)
Ordinal number generation:
- `ordinal()` function
- Numeric ordinals (1st, 2nd, 3rd, 21st, etc)
- Word ordinals (first, second, third, etc)
- Roundtrip testing with number_to_words()

### 6. Word Joining (go-inflect-apj)
Grammatically correct list joining:
- `join()` function
- Customizable conjunction (and/or)
- Customizable final separator
- Smart comma vs semicolon detection
- Tests for 1, 2, 3+ word lists

### 7. Classical Mode (go-inflect-va4)
Toggle between modern and classical English:
- `classical()` mode toggle
- `all=True/False` for global setting
- `zero=True` - "0 errors" vs "0 error"
- `herd=True` - "wildebeest" vs "wildebeests"
- `names=True` - proper name pluralization
- `ancient=True` - Latin/Greek plurals (formulae)
- `persons=True` - "persons" vs "people"
- Comprehensive mode combination testing

### 8. Present Participle (go-inflect-hea)
Verb conjugation:
- `present_participle()` - verb to -ing form

### 9. Comparison Functions (go-inflect-o1g)
Compare word forms:
- `compare()` - compare singular/plural forms
- `compare_nouns()` - compare noun forms
- `compare_verbs()` - compare verb forms
- `compare_adjs()` - compare adjective forms

### 10. Custom Definitions (go-inflect-dbg)
User-defined inflection rules:
- `defnoun()` - custom noun pluralization
- `defverb()` - custom verb conjugation
- `defadj()` - custom adjective pluralization
- `defa()` - custom a/an pattern
- `defan()` - custom a/an pattern
- Regex pattern support

### 11. Utility Functions (go-inflect-x25)
Helper functions:
- `num()` - store/retrieve default count
- `no()` - "no N" vs "N Xs"
- `inflect()` - parse and inflect inline text
- `gender()` - set pronoun gender

### 12. Project Setup (go-inflect-4nb)
Infrastructure:
- Go module initialization
- Testing framework
- README with examples
- CI/CD setup
- Contribution guidelines

### 13. Documentation (go-inflect-d1u)
Documentation and examples:
- Public API documentation
- Comprehensive examples
- Ported README examples from Python
- Performance benchmarks

## Getting Started

To see what's ready to work on:
```bash
bd ready
```

To view a specific epic and its tasks:
```bash
bd show go-inflect-tw5  # Pluralization epic
```

To see the dependency graph:
```bash
bd graph
```
