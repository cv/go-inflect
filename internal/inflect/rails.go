package inflect

import (
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// Humanize converts an underscored or dasherized string into a human-readable
// form. It capitalizes the first letter, replaces underscores and dashes with
// spaces, and strips trailing "_id" suffixes.
//
// This function is provided for compatibility with github.com/go-openapi/inflect
// and Rails ActiveSupport.
//
// Examples:
//
//	Humanize("employee_salary")  // "Employee salary"
//	Humanize("author_id")        // "Author"
//	Humanize("hello-world")      // "Hello world"
//	Humanize("XMLParser")        // "Xml parser"
func Humanize(word string) string {
	// Strip _id suffix
	word = strings.TrimSuffix(word, "_id")
	word = strings.TrimSuffix(word, "_ID")
	word = strings.TrimSuffix(word, "ID")

	// Convert to separated words (handles camelCase, snake_case, kebab-case)
	word = separatedWords(word, " ")

	// Lowercase and capitalize first letter
	word = strings.ToLower(word)
	if word != "" {
		word = strings.ToUpper(word[:1]) + word[1:]
	}

	return word
}

// ForeignKey creates an underscored foreign key name from a type name.
//
// This function is provided for compatibility with github.com/go-openapi/inflect
// and Rails ActiveSupport.
//
// Examples:
//
//	ForeignKey("Person")     // "person_id"
//	ForeignKey("Message")    // "message_id"
//	ForeignKey("AdminUser")  // "admin_user_id"
func ForeignKey(word string) string {
	return SnakeCase(word) + "_id"
}

// ForeignKeyCondensed creates a foreign key name without an underscore before "id".
//
// This function is provided for compatibility with github.com/go-openapi/inflect
// and Rails ActiveSupport.
//
// Examples:
//
//	ForeignKeyCondensed("Person")     // "personid"
//	ForeignKeyCondensed("Message")    // "messageid"
//	ForeignKeyCondensed("AdminUser")  // "admin_userid"
func ForeignKeyCondensed(word string) string {
	return SnakeCase(word) + "id"
}

// Tableize creates a table name from a type name. It underscores and pluralizes
// the word.
//
// This function is provided for compatibility with github.com/go-openapi/inflect
// and Rails ActiveSupport.
//
// Examples:
//
//	Tableize("Person")         // "people"
//	Tableize("RawScaledScorer") // "raw_scaled_scorers"
//	Tableize("MouseTrap")      // "mouse_traps"
func Tableize(word string) string {
	return Plural(SnakeCase(word))
}

// notURLSafe matches characters that are not safe for URLs.
// Matches anything that's not alphanumeric, dash, underscore, or space.
var notURLSafe = regexp.MustCompile(`[^a-zA-Z0-9\-_ ]`)

// multiSep matches multiple consecutive separators.
var multiSep = regexp.MustCompile(`[-_\s]+`)

// Parameterize converts a string to a URL-safe slug using dashes as separators.
//
// This function is provided for compatibility with github.com/go-openapi/inflect
// and Rails ActiveSupport.
//
// Examples:
//
//	Parameterize("Hello World!")     // "hello-world"
//	Parameterize("Hello, World!")    // "hello-world"
//	Parameterize("  Multiple   Spaces  ") // "multiple-spaces"
func Parameterize(word string) string {
	return ParameterizeJoin(word, "-")
}

// ParameterizeJoin converts a string to a URL-safe slug using a custom separator.
//
// This function is provided for compatibility with github.com/go-openapi/inflect
// and Rails ActiveSupport.
//
// Examples:
//
//	ParameterizeJoin("Hello World!", "_") // "hello_world"
//	ParameterizeJoin("Hello World!", "-") // "hello-world"
func ParameterizeJoin(word, sep string) string {
	word = strings.ToLower(word)
	word = Asciify(word)
	word = notURLSafe.ReplaceAllString(word, "")
	word = strings.TrimSpace(word)
	word = multiSep.ReplaceAllString(word, sep)
	word = strings.Trim(word, sep)
	return word
}

// Typeify converts a table name or plural word to a type name (singular, PascalCase).
//
// This function is provided for compatibility with github.com/go-openapi/inflect
// and Rails ActiveSupport.
//
// Examples:
//
//	Typeify("users")           // "User"
//	Typeify("raw_scaled_scorers") // "RawScaledScorer"
//	Typeify("people")          // "Person"
func Typeify(word string) string {
	return PascalCase(Singular(word))
}

// Asciify removes or transliterates non-ASCII characters from a string.
// Accented characters are converted to their ASCII equivalents where possible.
//
// This function is provided for compatibility with github.com/go-openapi/inflect
// and Rails ActiveSupport.
//
// Examples:
//
//	Asciify("café")    // "cafe"
//	Asciify("naïve")   // "naive"
//	Asciify("日本語")  // "" (non-Latin characters removed)
func Asciify(word string) string {
	// Use unicode normalization to decompose characters
	// NFD decomposes é into e + combining acute accent
	// Then we remove the combining marks
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, _ := transform.String(t, word)

	// Remove any remaining non-ASCII characters
	var sb strings.Builder
	for _, r := range result {
		if r < 128 {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}

// separatedWords converts a string to space-separated words, handling
// camelCase, snake_case, and kebab-case.
func separatedWords(word, sep string) string {
	// Replace underscores and dashes with separator
	word = strings.ReplaceAll(word, "_", sep)
	word = strings.ReplaceAll(word, "-", sep)

	// Handle camelCase by inserting separator before uppercase letters
	// Also handle acronyms like "XMLParser" -> "XML Parser"
	wordRunes := []rune(word)
	var result strings.Builder
	for i, r := range wordRunes {
		if i > 0 && unicode.IsUpper(r) {
			prev := wordRunes[i-1]
			// Insert separator if:
			// 1. Previous char is lowercase or digit (e.g., "myXML" -> "my XML")
			// 2. Previous char is uppercase but next char is lowercase (e.g., "XMLParser" -> "XML Parser")
			if unicode.IsLower(prev) || unicode.IsDigit(prev) {
				result.WriteString(sep)
			} else if unicode.IsUpper(prev) && i+1 < len(wordRunes) && unicode.IsLower(wordRunes[i+1]) {
				result.WriteString(sep)
			}
		}
		result.WriteRune(r)
	}

	// Clean up multiple separators
	output := result.String()
	if sep != "" {
		multiSepPattern := regexp.MustCompile(regexp.QuoteMeta(sep) + `+`)
		output = multiSepPattern.ReplaceAllString(output, sep)
	}

	return strings.TrimSpace(output)
}
