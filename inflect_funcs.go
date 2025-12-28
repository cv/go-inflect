package inflect

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// inflectFuncPattern matches function calls in the format:
// functionName('arg') or functionName('arg', num) or functionName(num)
// Supports single quotes, double quotes, or no quotes for string args.
var inflectFuncPattern = regexp.MustCompile(`(\w+)\(([^)]*)\)`)

// Inflect parses text containing inline function calls and replaces them
// with their inflected results.
//
// Supported function calls:
//
// Basic inflection:
//   - plural('word') - returns the plural form of the word
//   - plural('word', n) - returns plural if n != 1, singular otherwise
//   - singular('word') - returns the singular form of the word
//   - an('word') - returns the word with appropriate article ("a" or "an")
//   - a('word') - alias for an()
//
// Part-of-speech specific:
//   - plural_noun('word') - pluralizes nouns/pronouns ("I" -> "we")
//   - plural_noun('word', n) - with count
//   - plural_verb('word') - pluralizes verbs ("is" -> "are")
//   - plural_verb('word', n) - with count
//   - plural_adj('word') - pluralizes adjectives ("this" -> "these")
//   - plural_adj('word', n) - with count
//   - singular_noun('word') - singularizes nouns/pronouns
//   - singular_noun('word', n) - with count
//
// Numbers:
//   - ordinal(n) - returns ordinal form like "1st", "2nd", "3rd"
//   - ordinal_word(n) - returns ordinal word like "first", "second", "third"
//   - num(n) - returns the number as a string
//   - number_to_words(n) - converts number to words like "forty-two"
//   - number_to_words_with_and(n) - converts number to words with "and" like "one hundred and twenty-three"
//   - number_to_words_threshold(n, threshold) - returns words if n < threshold, else digits
//   - counting_word(n) - returns counting words: 1 → "once", 2 → "twice", 3+ → "3 times"
//   - no(word, count) - returns "no word" for 0, or "count words" otherwise
//   - format_number(n) - formats number with commas like "1,000"
//   - fraction(num, denom) - converts fraction to words like "one quarter"
//   - currency_to_words(amount, currency) - converts currency to words like "one dollar and fifty cents"
//
// Verb tenses:
//   - past_tense(verb) - returns past tense like "walk" → "walked"
//   - past_participle(verb) - returns past participle like "take" → "taken"
//   - present_participle(verb) - returns present participle like "run" → "running"
//
// Other inflections:
//   - possessive(noun) - returns possessive form like "cat" → "cat's"
//   - comparative(adj) - returns comparative form like "big" → "bigger"
//   - superlative(adj) - returns superlative form like "big" → "biggest"
//   - adverb(adj) - returns adverb form like "quick" → "quickly"
//
// Word ordinals:
//   - word_to_ordinal(word) - converts cardinal to ordinal word like "one" → "first"
//   - ordinal_to_cardinal(word) - converts ordinal to cardinal word like "first" → "one"
//
// Capitalization:
//   - capitalize(word) - capitalizes first letter like "hello" → "Hello"
//   - titleize(text) - capitalizes each word like "hello world" → "Hello World"
//
// Case conversion:
//   - snake_case(text) - converts to snake_case
//   - camel_case(text) - converts to camelCase
//   - pascal_case(text) - converts to PascalCase
//   - kebab_case(text) - converts to kebab-case
//   - humanize(text) - converts to human-readable form
//
// Rails-style functions:
//   - tableize(word) - converts to table name like "Person" → "people"
//   - foreign_key(word) - converts to foreign key like "Person" → "person_id"
//   - typeify(word) - converts to type name like "user_post" → "UserPost"
//   - parameterize(word) - converts to URL slug like "Hello World" → "hello-world"
//   - asciify(word) - converts to ASCII like "café" → "cafe"
//
// Word comparison:
//   - compare(word1, word2) - compares words for singular/plural equality
//   - compare_nouns(noun1, noun2) - compares nouns for singular/plural equality
//   - compare_verbs(verb1, verb2) - compares verbs for singular/plural equality
//   - compare_adjs(adj1, adj2) - compares adjectives for singular/plural equality
//
// List joining:
//   - join('a', 'b', 'c') - joins with Oxford comma: "a, b, and c"
//   - join_with('or', 'a', 'b', 'c') - joins with custom conjunction: "a, b, or c"
//
// Other:
//   - word_count(text) - counts words in text, returns count as string
//
// Examples:
//   - Inflect("The plural of cat is plural('cat')") -> "The plural of cat is cats"
//   - Inflect("I saw an('apple')") -> "I saw an apple"
//   - Inflect("There are num(3) plural('error', 3)") -> "There are 3 errors"
//   - Inflect("This is the ordinal(1) item") -> "This is the 1st item"
//   - Inflect("plural_noun('I') saw it") -> "We saw it"
//   - Inflect("The cat plural_verb('is') happy") -> "The cat are happy"
//   - Inflect("She past_tense('walk') home") -> "She walked home"
//   - Inflect("I have past_participle('take') it") -> "I have taken it"
//   - Inflect("He is present_participle('run')") -> "He is running"
//   - Inflect("The possessive('cat') toy") -> "The cat's toy"
//   - Inflect("This is comparative('big')") -> "This is bigger"
//   - Inflect("This is the superlative('big')") -> "This is the biggest"
//   - Inflect("The ordinal_word(1) place") -> "The first place"
//   - Inflect("I have number_to_words(42) apples") -> "I have forty-two apples"
//   - Inflect("I saw it counting_word(2)") -> "I saw it twice"
//   - Inflect("There are no('error', 0)") -> "There are no errors"
func Inflect(text string) string {
	return inflectFuncPattern.ReplaceAllStringFunc(text, processInflectCall)
}

// inflectFunc is a function that processes arguments and returns the result.
type inflectFunc func(args []string, original string) string

// inflectFuncs maps function names to their handlers.
var inflectFuncs = map[string]inflectFunc{
	// Basic inflection
	"plural":        processPlural,
	"plural_noun":   processPluralNoun,
	"plural_verb":   processPluralVerb,
	"plural_adj":    processPluralAdj,
	"singular":      processSingularNoun,
	"singular_noun": processSingularNoun,
	"an":            processArticle,
	"a":             processArticle,

	// Numbers
	"ordinal":                   processOrdinal,
	"ordinal_word":              processOrdinalWord,
	"num":                       processNum,
	"number_to_words":           processNumberToWords,
	"number_to_words_with_and":  processNumberToWordsWithAnd,
	"number_to_words_threshold": processNumberToWordsThreshold,
	"counting_word":             processCountingWord,
	"no":                        processNo,
	"format_number":             processFormatNumber,
	"fraction":                  processFraction,
	"currency_to_words":         processCurrencyToWords,

	// Verb tenses
	"past_tense":         processPastTense,
	"past_participle":    processPastParticiple,
	"present_participle": processPresentParticiple,

	// Other inflections
	"possessive":  processPossessive,
	"comparative": processComparative,
	"superlative": processSuperlative,
	"adverb":      processAdverb,

	// Word ordinals
	"word_to_ordinal":     processWordToOrdinal,
	"ordinal_to_cardinal": processOrdinalToCardinal,

	// Capitalization
	"capitalize": processCapitalize,
	"titleize":   processTitleize,

	// Case conversion
	"snake_case":  processSnakeCase,
	"camel_case":  processCamelCase,
	"pascal_case": processPascalCase,
	"kebab_case":  processKebabCase,
	"humanize":    processHumanize,

	// Rails-style functions
	"tableize":     processTableize,
	"foreign_key":  processForeignKey,
	"typeify":      processTypeify,
	"parameterize": processParameterize,
	"asciify":      processAsciify,

	// Word comparison
	"compare":       processCompare,
	"compare_nouns": processCompareNouns,
	"compare_verbs": processCompareVerbs,
	"compare_adjs":  processCompareAdjs,

	// Other
	"word_count": processWordCount,

	// List joining
	"join":      processJoin,
	"join_with": processJoinWith,
}

// processInflectCall processes a single function call match and returns
// the inflected result.
func processInflectCall(match string) string {
	// Parse the function name and arguments
	submatches := inflectFuncPattern.FindStringSubmatch(match)
	if len(submatches) < 3 {
		return match
	}

	funcName := strings.ToLower(submatches[1])
	argsStr := strings.TrimSpace(submatches[2])

	// Look up and call the handler function
	if handler, ok := inflectFuncs[funcName]; ok {
		args := parseInflectArgs(argsStr)
		return handler(args, match)
	}

	return match
}

// parseInflectArgs parses the arguments string into a slice of strings.
// Handles quoted strings and numeric arguments.
func parseInflectArgs(argsStr string) []string {
	if argsStr == "" {
		return nil
	}

	var args []string
	var current strings.Builder
	inQuote := false
	quoteChar := rune(0)

	for i, r := range argsStr {
		switch {
		case (r == '\'' || r == '"') && !inQuote:
			inQuote = true
			quoteChar = r
		case r == quoteChar && inQuote:
			inQuote = false
			quoteChar = 0
		case r == ',' && !inQuote:
			arg := strings.TrimSpace(current.String())
			if arg != "" {
				args = append(args, arg)
			}
			current.Reset()
		default:
			// Skip leading/trailing whitespace in arguments
			if !inQuote && current.Len() == 0 && unicode.IsSpace(r) {
				continue
			}
			// Check if we're at end and it's trailing whitespace
			if !inQuote && unicode.IsSpace(r) {
				// Look ahead to see if there's more non-whitespace
				hasMore := false
				for j := i + 1; j < len(argsStr); j++ {
					if !unicode.IsSpace(rune(argsStr[j])) && argsStr[j] != ',' {
						hasMore = true
						break
					} else if argsStr[j] == ',' {
						break
					}
				}
				if !hasMore {
					continue
				}
			}
			current.WriteRune(r)
		}
	}

	// Add the last argument
	arg := strings.TrimSpace(current.String())
	if arg != "" {
		args = append(args, arg)
	}

	return args
}

// processPlural handles plural('word') and plural('word', n).
func processPlural(args []string, original string) string {
	if len(args) == 0 {
		return original
	}

	word := args[0]

	// If there's a count argument, use it to determine singular/plural
	if len(args) >= 2 {
		count, err := strconv.Atoi(args[1])
		if err != nil {
			return original
		}
		if count == 1 || count == -1 {
			return word
		}
		return Plural(word)
	}

	return Plural(word)
}

// processPluralNoun handles plural_noun('word') and plural_noun('word', n).
func processPluralNoun(args []string, original string) string {
	if len(args) == 0 {
		return original
	}

	word := args[0]

	// If there's a count argument, use it
	if len(args) >= 2 {
		count, err := strconv.Atoi(args[1])
		if err != nil {
			return original
		}
		return PluralNoun(word, count)
	}

	return PluralNoun(word)
}

// processPluralVerb handles plural_verb('word') and plural_verb('word', n).
func processPluralVerb(args []string, original string) string {
	if len(args) == 0 {
		return original
	}

	word := args[0]

	// If there's a count argument, use it
	if len(args) >= 2 {
		count, err := strconv.Atoi(args[1])
		if err != nil {
			return original
		}
		return PluralVerb(word, count)
	}

	return PluralVerb(word)
}

// processPluralAdj handles plural_adj('word') and plural_adj('word', n).
func processPluralAdj(args []string, original string) string {
	if len(args) == 0 {
		return original
	}

	word := args[0]

	// If there's a count argument, use it
	if len(args) >= 2 {
		count, err := strconv.Atoi(args[1])
		if err != nil {
			return original
		}
		return PluralAdj(word, count)
	}

	return PluralAdj(word)
}

// processSingularNoun handles singular_noun('word') and singular_noun('word', n).
func processSingularNoun(args []string, original string) string {
	if len(args) == 0 {
		return original
	}

	word := args[0]

	// If there's a count argument, use it
	if len(args) >= 2 {
		count, err := strconv.Atoi(args[1])
		if err != nil {
			return original
		}
		return SingularNoun(word, count)
	}

	return SingularNoun(word)
}

// processArticle handles an('word') and a('word').
func processArticle(args []string, original string) string {
	if len(args) == 0 {
		return original
	}
	return An(args[0])
}

// processOrdinal handles ordinal(n).
func processOrdinal(args []string, original string) string {
	if len(args) == 0 {
		return original
	}
	n, err := strconv.Atoi(args[0])
	if err != nil {
		return original
	}
	return Ordinal(n)
}

// processNum handles num(n).
func processNum(args []string, original string) string {
	if len(args) == 0 {
		return original
	}
	// Just return the number as-is (it's already a string)
	return args[0]
}

// processOrdinalWord handles ordinal_word(n).
func processOrdinalWord(args []string, original string) string {
	if len(args) == 0 {
		return original
	}
	n, err := strconv.Atoi(args[0])
	if err != nil {
		return original
	}
	return OrdinalWord(n)
}

// processNumberToWords handles number_to_words(n).
func processNumberToWords(args []string, original string) string {
	if len(args) == 0 {
		return original
	}
	n, err := strconv.Atoi(args[0])
	if err != nil {
		return original
	}
	return NumberToWords(n)
}

// processCountingWord handles counting_word(n).
func processCountingWord(args []string, original string) string {
	if len(args) == 0 {
		return original
	}
	n, err := strconv.Atoi(args[0])
	if err != nil {
		return original
	}
	return CountingWord(n)
}

// processNo handles no('word', count).
func processNo(args []string, original string) string {
	if len(args) < 2 {
		return original
	}
	word := args[0]
	count, err := strconv.Atoi(args[1])
	if err != nil {
		return original
	}
	return No(word, count)
}

// processPastTense handles past_tense('verb').
func processPastTense(args []string, original string) string {
	if len(args) == 0 {
		return original
	}
	return PastTense(args[0])
}

// processPastParticiple handles past_participle('verb').
func processPastParticiple(args []string, original string) string {
	if len(args) == 0 {
		return original
	}
	return PastParticiple(args[0])
}

// processPresentParticiple handles present_participle('verb').
func processPresentParticiple(args []string, original string) string {
	if len(args) == 0 {
		return original
	}
	return PresentParticiple(args[0])
}

// processPossessive handles possessive('noun').
func processPossessive(args []string, original string) string {
	if len(args) == 0 {
		return original
	}
	return Possessive(args[0])
}

// processComparative handles comparative('adj').
func processComparative(args []string, original string) string {
	if len(args) == 0 {
		return original
	}
	return Comparative(args[0])
}

// processSuperlative handles superlative('adj').
func processSuperlative(args []string, original string) string {
	if len(args) == 0 {
		return original
	}
	return Superlative(args[0])
}

// PluralNoun returns the plural form of an English noun or pronoun.
//
// This function handles:
//   - Pronouns in nominative case: "I" -> "we", "he"/"she"/"it" -> "they"
//   - Pronouns in accusative case: "me" -> "us", "him"/"her" -> "them"
//   - Possessive pronouns: "my" -> "our", "mine" -> "ours", "his"/"her"/"its" -> "their"
//   - Reflexive pronouns: "myself" -> "ourselves", "himself"/"herself"/"itself" -> "themselves"
//   - Regular nouns: defers to Plural()
//
// If count is provided and equals 1 or -1, returns the singular form.
// If count is not 1, returns the plural form.
// If no count is provided, returns the plural form.
//
// Examples:
//
//	PluralNoun("I") returns "we"
//	PluralNoun("me") returns "us"
//	PluralNoun("my") returns "our"
//	PluralNoun("cat") returns "cats"
//	PluralNoun("cat", 1) returns "cat"
//	PluralNoun("cat", 2) returns "cats"
func PluralNoun(word string, count ...int) string {
	if word == "" {
		return ""
	}

	// Preserve leading/trailing whitespace
	prefix, trimmed, suffix := extractWhitespace(word)
	if trimmed == "" {
		return word
	}

	// Handle count parameter
	if len(count) > 0 && (count[0] == 1 || count[0] == -1) {
		return word // Return singular form as-is
	}

	lower := strings.ToLower(trimmed)

	// Check for pronouns first
	if plural, ok := allPronounsToPlural[lower]; ok {
		return prefix + matchCase(trimmed, plural) + suffix
	}

	// Fall back to regular Plural() for nouns
	return prefix + Plural(trimmed) + suffix
}

// PluralVerb returns the plural form of an English verb.
//
// This function handles:
//   - Auxiliary verbs: "is" -> "are", "was" -> "were", "has" -> "have"
//   - Contractions: "isn't" -> "aren't", "doesn't" -> "don't", "hasn't" -> "haven't"
//   - Modal verbs (unchanged): "can", "could", "may", "might", "must", "shall", "should", "will", "would"
//   - Regular verbs in third person singular: removes -s/-es suffix
//
// If count is provided and equals 1 or -1, returns the singular form.
// If count is not 1, returns the plural form.
// If no count is provided, returns the plural form.
//
// Examples:
//
//	PluralVerb("is") returns "are"
//	PluralVerb("was") returns "were"
//	PluralVerb("has") returns "have"
//	PluralVerb("doesn't") returns "don't"
//	PluralVerb("runs") returns "run"
//	PluralVerb("is", 1) returns "is"
//	PluralVerb("is", 2) returns "are"
func PluralVerb(word string, count ...int) string {
	if word == "" {
		return ""
	}

	// Preserve leading/trailing whitespace
	prefix, trimmed, suffix := extractWhitespace(word)
	if trimmed == "" {
		return word
	}

	// Handle count parameter - if singular count, return singular form
	if len(count) > 0 && (count[0] == 1 || count[0] == -1) {
		// Return appropriate singular form
		lower := strings.ToLower(trimmed)
		if singular, ok := verbPluralToSingular[lower]; ok {
			return prefix + matchCase(trimmed, singular) + suffix
		}
		return word
	}

	lower := strings.ToLower(trimmed)

	// Check for unchanged modal verbs
	if verbUnchanged[lower] {
		return word
	}

	// Check for known irregular verb mappings
	if plural, ok := verbSingularToPlural[lower]; ok {
		return prefix + matchCase(trimmed, plural) + suffix
	}

	// Check custom verb definitions
	if plural, ok := customVerbs[lower]; ok {
		return prefix + matchCase(trimmed, plural) + suffix
	}

	// For regular verbs in third person singular (ends in -s/-es),
	// convert to base form (which is the plural form)

	// Handle -ies -> -y first (tries -> try, flies -> fly)
	if strings.HasSuffix(lower, "ies") && len(lower) > 3 {
		return prefix + trimmed[:len(trimmed)-3] + matchSuffix(trimmed, "y") + suffix
	}

	// Handle -es after sibilants: -sses, -shes, -ches, -xes, -zes
	if strings.HasSuffix(lower, "es") && len(lower) > 2 {
		base := lower[:len(lower)-2]
		if strings.HasSuffix(base, "ss") ||
			strings.HasSuffix(base, "sh") ||
			strings.HasSuffix(base, "ch") ||
			strings.HasSuffix(base, "x") ||
			strings.HasSuffix(base, "z") {
			return prefix + trimmed[:len(trimmed)-2] + suffix
		}
		// -oes -> -o (goes -> go) - but goes is in the irregular list
		if strings.HasSuffix(base, "o") {
			return prefix + trimmed[:len(trimmed)-2] + suffix
		}
		// For other -es endings like "sees", just remove -s (not -es)
		// "sees" -> "see", "flees" -> "flee"
	}

	// Regular -s ending: just remove the -s
	if strings.HasSuffix(lower, "s") && len(lower) > 1 {
		// Don't remove -s from words ending in -ss
		if !strings.HasSuffix(lower, "ss") {
			// Regular third person singular: runs -> run, sees -> see
			return prefix + trimmed[:len(trimmed)-1] + suffix
		}
	}

	// Return unchanged if no rule matches (already plural or base form)
	return word
}

// PluralAdj returns the plural form of an English adjective.
//
// This function handles:
//   - Demonstrative adjectives: "this" -> "these", "that" -> "those"
//   - Indefinite articles: "a" -> "some", "an" -> "some"
//   - Possessive adjectives: "my" -> "our", "his"/"her"/"its" -> "their"
//
// If count is provided and equals 1 or -1, returns the singular form.
// If count is not 1, returns the plural form.
// If no count is provided, returns the plural form.
//
// Examples:
//
//	PluralAdj("this") returns "these"
//	PluralAdj("that") returns "those"
//	PluralAdj("a") returns "some"
//	PluralAdj("an") returns "some"
//	PluralAdj("my") returns "our"
//	PluralAdj("this", 1) returns "this"
//	PluralAdj("this", 2) returns "these"
func PluralAdj(word string, count ...int) string {
	if word == "" {
		return ""
	}

	// Preserve leading/trailing whitespace
	prefix, trimmed, suffix := extractWhitespace(word)
	if trimmed == "" {
		return word
	}

	// Handle count parameter - if singular count, return singular form
	if len(count) > 0 && (count[0] == 1 || count[0] == -1) {
		// Return appropriate singular form
		lower := strings.ToLower(trimmed)
		if singular, ok := adjPluralToSingular[lower]; ok {
			return prefix + matchCase(trimmed, singular) + suffix
		}
		if genderMap, ok := adjPluralToSingularByGender[lower]; ok {
			if singular, ok := genderMap[gender]; ok {
				return prefix + matchCase(trimmed, singular) + suffix
			}
		}
		return word
	}

	lower := strings.ToLower(trimmed)

	// Check for known adjective mappings
	if plural, ok := adjSingularToPlural[lower]; ok {
		return prefix + matchCase(trimmed, plural) + suffix
	}

	// Check custom adjective definitions
	if plural, ok := customAdjs[lower]; ok {
		return prefix + matchCase(trimmed, plural) + suffix
	}

	// Most adjectives don't change between singular and plural in English
	return word
}

// SingularNoun returns the singular form of an English noun or pronoun.
//
// This function handles:
//   - Pronouns in nominative case: "we" -> "I", "they" -> he/she/it/they (depends on gender)
//   - Pronouns in accusative case: "us" -> "me", "them" -> him/her/it/them (depends on gender)
//   - Possessive pronouns: "our" -> "my", "ours" -> "mine", "their" -> his/her/its/their
//   - Reflexive pronouns: "ourselves" -> "myself", "themselves" -> himself/herself/itself/themself
//   - Regular nouns: defers to Singular()
//
// Third-person singular pronouns use the gender set by Gender():
//   - Gender("m"): masculine - "they" -> "he"
//   - Gender("f"): feminine - "they" -> "she"
//   - Gender("n"): neuter - "they" -> "it"
//   - Gender("t"): they (singular they) - "they" -> "they"
//
// If count is provided and equals 1 or -1, returns the singular form.
// If count is not 1, returns the plural form.
// If no count is provided, returns the singular form.
//
// Examples:
//
//	SingularNoun("we") returns "I"
//	SingularNoun("us") returns "me"
//	SingularNoun("our") returns "my"
//	SingularNoun("they") returns "it" (or he/she/they based on gender)
//	SingularNoun("cats") returns "cat"
//	SingularNoun("cats", 1) returns "cat"
//	SingularNoun("cats", 2) returns "cats"
func SingularNoun(word string, count ...int) string {
	if word == "" {
		return ""
	}

	// Preserve leading/trailing whitespace
	prefix, trimmed, suffix := extractWhitespace(word)
	if trimmed == "" {
		return word
	}

	// Handle count parameter - if plural count, return plural form
	if len(count) > 0 && count[0] != 1 && count[0] != -1 {
		return word // Return plural form as-is
	}

	lower := strings.ToLower(trimmed)

	// Check for nominative pronouns
	if genderMap, ok := pronounNominativeSingularByGender[lower]; ok {
		if singular, ok := genderMap[gender]; ok {
			return prefix + matchCase(trimmed, singular) + suffix
		}
	}

	// Check for accusative pronouns
	if genderMap, ok := pronounAccusativeSingularByGender[lower]; ok {
		if singular, ok := genderMap[gender]; ok {
			return prefix + matchCase(trimmed, singular) + suffix
		}
	}

	// Check for possessive pronouns
	if genderMap, ok := pronounPossessiveSingularByGender[lower]; ok {
		if singular, ok := genderMap[gender]; ok {
			return prefix + matchCase(trimmed, singular) + suffix
		}
	}

	// Check for reflexive pronouns
	if genderMap, ok := pronounReflexiveSingularByGender[lower]; ok {
		if singular, ok := genderMap[gender]; ok {
			return prefix + matchCase(trimmed, singular) + suffix
		}
	}

	// Fall back to regular Singular() for nouns
	return prefix + Singular(trimmed) + suffix
}

// processAdverb handles adverb('adj').
func processAdverb(args []string, original string) string {
	if len(args) == 0 {
		return original
	}
	return Adverb(args[0])
}

// processCapitalize handles capitalize('word').
func processCapitalize(args []string, original string) string {
	if len(args) == 0 {
		return original
	}
	return Capitalize(args[0])
}

// processTitleize handles titleize('text').
func processTitleize(args []string, original string) string {
	if len(args) == 0 {
		return original
	}
	return Titleize(args[0])
}

// processWordToOrdinal handles word_to_ordinal('word').
func processWordToOrdinal(args []string, original string) string {
	if len(args) == 0 {
		return original
	}
	return WordToOrdinal(args[0])
}

// processOrdinalToCardinal handles ordinal_to_cardinal('word').
func processOrdinalToCardinal(args []string, original string) string {
	if len(args) == 0 {
		return original
	}
	return OrdinalToCardinal(args[0])
}

// processFraction handles fraction(num, denom).
func processFraction(args []string, original string) string {
	if len(args) < 2 {
		return original
	}
	num, err := strconv.Atoi(args[0])
	if err != nil {
		return original
	}
	denom, err := strconv.Atoi(args[1])
	if err != nil {
		return original
	}
	return FractionToWords(num, denom)
}

// processFormatNumber handles format_number(n).
func processFormatNumber(args []string, original string) string {
	if len(args) == 0 {
		return original
	}
	n, err := strconv.Atoi(args[0])
	if err != nil {
		return original
	}
	return FormatNumber(n)
}

// processSnakeCase handles snake_case('text').
func processSnakeCase(args []string, original string) string {
	if len(args) == 0 {
		return original
	}
	return SnakeCase(args[0])
}

// processCamelCase handles camel_case('text').
func processCamelCase(args []string, original string) string {
	if len(args) == 0 {
		return original
	}
	return CamelCase(args[0])
}

// processPascalCase handles pascal_case('text').
func processPascalCase(args []string, original string) string {
	if len(args) == 0 {
		return original
	}
	return PascalCase(args[0])
}

// processKebabCase handles kebab_case('text').
func processKebabCase(args []string, original string) string {
	if len(args) == 0 {
		return original
	}
	return KebabCase(args[0])
}

// processHumanize handles humanize('text').
func processHumanize(args []string, original string) string {
	if len(args) == 0 {
		return original
	}
	return Humanize(args[0])
}

// processTableize handles tableize('word').
func processTableize(args []string, original string) string {
	if len(args) == 0 {
		return original
	}
	return Tableize(args[0])
}

// processForeignKey handles foreign_key('word').
func processForeignKey(args []string, original string) string {
	if len(args) == 0 {
		return original
	}
	return ForeignKey(args[0])
}

// processTypeify handles typeify('word').
func processTypeify(args []string, original string) string {
	if len(args) == 0 {
		return original
	}
	return Typeify(args[0])
}

// processParameterize handles parameterize('word').
func processParameterize(args []string, original string) string {
	if len(args) == 0 {
		return original
	}
	return Parameterize(args[0])
}

// processAsciify handles asciify('word').
func processAsciify(args []string, original string) string {
	if len(args) == 0 {
		return original
	}
	return Asciify(args[0])
}

// processNumberToWordsWithAnd handles number_to_words_with_and(n).
func processNumberToWordsWithAnd(args []string, original string) string {
	if len(args) == 0 {
		return original
	}
	n, err := strconv.Atoi(args[0])
	if err != nil {
		return original
	}
	return NumberToWordsWithAnd(n)
}

// processNumberToWordsThreshold handles number_to_words_threshold(n, threshold).
func processNumberToWordsThreshold(args []string, original string) string {
	if len(args) < 2 {
		return original
	}
	n, err := strconv.Atoi(args[0])
	if err != nil {
		return original
	}
	threshold, err := strconv.Atoi(args[1])
	if err != nil {
		return original
	}
	return NumberToWordsThreshold(n, threshold)
}

// processCurrencyToWords handles currency_to_words(amount, currency).
func processCurrencyToWords(args []string, original string) string {
	if len(args) < 2 {
		return original
	}
	amount, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return original
	}
	currency := args[1]
	return CurrencyToWords(amount, currency)
}

// processCompare handles compare('word1', 'word2').
func processCompare(args []string, original string) string {
	if len(args) < 2 {
		return original
	}
	return Compare(args[0], args[1])
}

// processCompareNouns handles compare_nouns('noun1', 'noun2').
func processCompareNouns(args []string, original string) string {
	if len(args) < 2 {
		return original
	}
	return CompareNouns(args[0], args[1])
}

// processCompareVerbs handles compare_verbs('verb1', 'verb2').
func processCompareVerbs(args []string, original string) string {
	if len(args) < 2 {
		return original
	}
	return CompareVerbs(args[0], args[1])
}

// processCompareAdjs handles compare_adjs('adj1', 'adj2').
func processCompareAdjs(args []string, original string) string {
	if len(args) < 2 {
		return original
	}
	return CompareAdjs(args[0], args[1])
}

// processWordCount handles word_count('text').
func processWordCount(args []string, original string) string {
	if len(args) == 0 {
		return original
	}
	return strconv.Itoa(WordCount(args[0]))
}

// processJoin handles join('a', 'b', 'c') -> "a, b, and c".
func processJoin(args []string, original string) string {
	if len(args) == 0 {
		return original
	}
	return Join(args)
}

// processJoinWith handles join_with('or', 'a', 'b', 'c') -> "a, b, or c".
func processJoinWith(args []string, original string) string {
	if len(args) < 2 {
		return original
	}
	return JoinWithConj(args[1:], args[0])
}
