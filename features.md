# go-inflect Port - Feature Tracking

## Core Features

### Pluralization
- title: Implement plural() - unconditional pluralization
- title: Implement plural_noun() - pluralize nouns
- title: Implement plural_verb() - pluralize verbs
- title: Implement plural_adj() - pluralize adjectives
- title: Test 860+ inflection test cases from inflections.txt
- title: Implement compound word pluralization (hyphenated and spaced)
- title: Implement unit handling (degree celsius, per square inch, etc)
- title: Implement open compound nouns (high school, prima donna)

### Singularization
- title: Implement singular_noun() - convert plural to singular
- title: Test singular_noun() with all plural forms

### Indefinite Articles
- title: Implement a() and an() - indefinite article selection
- title: Test a/an with common words (cat, ant, honest, etc)
- title: Test a/an with abbreviations (YAML, JSON, Core ML, etc)
- title: Test a/an with silent letters and special cases

### Number Conversion
- title: Implement number_to_words() - convert numbers to words
- title: Support basic numbers (0-20, 21-99, 100s, 1000s, millions, etc)
- title: Support decimal numbers
- title: Support threshold parameter (don't convert above threshold)
- title: Support comma parameter (format numbers with commas)
- title: Support group parameter (group digits by N)
- title: Support decimal parameter (decimal point character)
- title: Test 100+ number_to_words test cases

### Ordinals
- title: Implement ordinal() - convert number to ordinal (1st, 2nd, 3rd)
- title: Support numeric ordinals (1st, 2nd, 3rd, 21st, etc)
- title: Support word ordinals (first, second, third, etc)
- title: Test ordinal() with number_to_words() roundtrip

### Word Joining
- title: Implement join() - join words with proper grammar
- title: Support customizable conjunction (and/or)
- title: Support customizable final separator
- title: Support comma vs semicolon detection
- title: Test join with 1, 2, 3+ words

### Classical Mode
- title: Implement classical() mode toggle
- title: Support classical(all=True/False)
- title: Support classical(zero=True) - "0 errors" vs "0 error"
- title: Support classical(herd=True) - "wildebeest" vs "wildebeests"
- title: Support classical(names=True) - proper name pluralization
- title: Support classical(ancient=True) - Latin/Greek plurals (formulae)
- title: Support classical(persons=True) - "persons" vs "people"
- title: Test all classical mode combinations

### Present Participle
- title: Implement present_participle() - verb to -ing form

### Comparison Functions
- title: Implement compare() - compare singular/plural forms
- title: Implement compare_nouns() - compare noun forms
- title: Implement compare_verbs() - compare verb forms
- title: Implement compare_adjs() - compare adjective forms

### Custom Definitions
- title: Implement defnoun() - define custom noun pluralization
- title: Implement defverb() - define custom verb conjugation
- title: Implement defadj() - define custom adjective pluralization
- title: Implement defa() - define custom a/an pattern
- title: Implement defan() - define custom a/an pattern
- title: Support regex patterns in custom definitions

### Utility Functions
- title: Implement num() - store/retrieve default count
- title: Implement no() - "no N" vs "N Xs"
- title: Implement inflect() - parse and inflect inline text
- title: Implement gender() - set pronoun gender

### Project Setup
- title: Initialize Go module and project structure
- title: Set up testing framework
- title: Create README with examples
- title: Set up CI/CD
- title: Write contribution guidelines

### Documentation
- title: Document all public APIs
- title: Create comprehensive examples
- title: Port README examples from Python
- title: Create performance benchmarks
