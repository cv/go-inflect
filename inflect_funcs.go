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
//   - num(n) - returns the number as a string
//
// Examples:
//   - Inflect("The plural of cat is plural('cat')") -> "The plural of cat is cats"
//   - Inflect("I saw an('apple')") -> "I saw an apple"
//   - Inflect("There are num(3) plural('error', 3)") -> "There are 3 errors"
//   - Inflect("This is the ordinal(1) item") -> "This is the 1st item"
//   - Inflect("plural_noun('I') saw it") -> "We saw it"
//   - Inflect("The cat plural_verb('is') happy") -> "The cat are happy"
func Inflect(text string) string {
	return inflectFuncPattern.ReplaceAllStringFunc(text, processInflectCall)
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

	// Parse arguments
	args := parseInflectArgs(argsStr)

	switch funcName {
	case "plural":
		return processPlural(args, match)
	case "plural_noun":
		return processPluralNoun(args, match)
	case "plural_verb":
		return processPluralVerb(args, match)
	case "plural_adj":
		return processPluralAdj(args, match)
	case "singular", "singular_noun":
		return processSingularNoun(args, match)
	case "an", "a":
		return processArticle(args, match)
	case "ordinal":
		return processOrdinal(args, match)
	case "num":
		return processNum(args, match)
	default:
		return match
	}
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
