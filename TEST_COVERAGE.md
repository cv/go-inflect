# Test Coverage Overview

This document summarizes the test cases we need to port from Python's inflect library.

## Test Files Summary

### test_inflections.py (860+ test cases)
- **File**: `reference/tests/inflections.txt` (860 lines)
- **Coverage**: Core noun/verb/adjective pluralization
- **Format**: `singular -> plural | classical_plural # comment`
- **Examples**:
  - `abscissa -> abscissas|abscissae`
  - `brother -> brothers|brethren`
  - `formula -> formulas|formulae`

### test_numwords.py (100+ test cases)
- Basic numbers: 0-20, 21-99, 100-999, 1000+
- Decimal numbers
- Threshold parameter testing
- Comma formatting
- Group parameter
- Ordinal testing with roundtrips
- **Example cases**:
  - `0 -> "zero"`
  - `21 -> "twenty-one"`
  - `1000 -> "one thousand" or "1,000"`
  - `ordinal(1) -> "1st"`
  - `number_to_words(ordinal(1)) -> "first"`

### test_an.py (20+ test cases)
- Common words: cat, ant, honest cat, dishonest cat
- Silent H: Honolulu, honest
- Vowels with consonant sound: Ugandan, Ukrainian, unanimous
- Abbreviations: YAML, JSON, Core ML, mpeg
- Special cases: US farmer, wild PIKACHU

### test_join.py (15+ test cases)
- 1 word: "carrot" -> "carrot"
- 2 words: ["apple", "carrot"] -> "apple and carrot"
- 3 words: ["apple", "banana", "carrot"] -> "apple, banana, and carrot"
- Custom conjunction: "or" instead of "and"
- Custom final separator
- Semicolon detection (for items with commas)
- Spacing control (conj_spaced parameter)

### test_compounds.py (30+ test cases)
- Hyphenated compounds: "hello-out-there"
- Space-separated: "hello out there"
- Units with "degree": degree celsius, degree fahrenheit
- Units with "per": pound per square inch, metre per second
- Open compound nouns: high school, master genie, prima donna
- Classical mode for compounds: master genii, prime donne, Blood brethren

### test_classical_*.py (50+ test cases)
Six test files covering classical mode:

**test_classical_all.py**:
- Tests all classical flags together
- zero: "0 errors" vs "0 error"
- herd: "wildebeest" vs "wildebeests"
- names: proper name handling
- persons: "persons" vs "people"
- ancient: "formulae" vs "formulas"

**test_classical_zero.py**:
- classical(zero=True): "0 error" (classical)
- classical(zero=False): "0 errors" (modern)

**test_classical_herd.py**:
- classical(herd=True): "wildebeest" (classical)
- classical(herd=False): "wildebeests" (modern)

**test_classical_names.py**:
- classical(names=True): affects proper name pluralization
- Example: "Sally" -> "Sallys" vs "Sallies"

**test_classical_ancient.py**:
- classical(ancient=True): Latin/Greek plurals
- Examples: formula -> formulae, datum -> data

**test_classical_person.py**:
- classical(persons=True): "persons" instead of "people"

### test_pl_si.py
- Tests plural_noun and singular_noun roundtrip
- Ensures "I" -> "we" (plural_noun)

### test_pwd.py (1200+ lines, comprehensive)
- "Plurals, word-for-word" tests
- Most comprehensive test suite
- Tests all edge cases and special forms

### test_unicode.py
- Unicode character handling
- Ensures library works with non-ASCII text

## Total Test Case Count

| Test File | Approximate Cases |
|-----------|------------------|
| test_inflections.py | 860+ |
| test_numwords.py | 100+ |
| test_pwd.py | 500+ |
| test_compounds.py | 30+ |
| test_classical_*.py | 50+ |
| test_an.py | 20+ |
| test_join.py | 15+ |
| test_pl_si.py | 10+ |
| test_unicode.py | 5+ |
| **TOTAL** | **~1,600+** |

## Test Priority

### Phase 1 - Core Functionality (MVP)
1. Basic pluralization (test_inflections.py - subset)
2. Indefinite articles (test_an.py)
3. Number to words (test_numwords.py - basic cases)
4. Word joining (test_join.py)

### Phase 2 - Advanced Features
1. Classical mode (test_classical_*.py)
2. Compound words (test_compounds.py)
3. Singularization (test_pl_si.py)
4. Advanced number conversion (test_numwords.py - complete)

### Phase 3 - Comprehensive Coverage
1. All inflections (test_inflections.py - complete)
2. Edge cases (test_pwd.py)
3. Unicode support (test_unicode.py)

## Reference Files Location

- Python test files: `reference/tests/*.py`
- Test data: `reference/tests/inflections.txt`
- Python README: `reference/README.rst`
