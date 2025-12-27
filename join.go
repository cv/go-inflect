package inflect

import "strings"

// Join combines a slice of strings into a grammatically correct English list.
//
// The function uses the Oxford comma (serial comma) for lists of three or more items.
// It uses "and" as the conjunction. For custom conjunctions, use JoinWithConj.
//
// Examples:
//   - Join([]string{}) returns ""
//   - Join([]string{"a"}) returns "a"
//   - Join([]string{"a", "b"}) returns "a and b"
//   - Join([]string{"a", "b", "c"}) returns "a, b, and c"
func Join(words []string) string {
	return JoinWithConj(words, "and")
}

// JoinWithConj combines a slice of strings into a grammatically correct English list
// with a custom conjunction.
//
// The function uses the Oxford comma (serial comma) for lists of three or more items.
//
// Examples:
//   - JoinWithConj([]string{"a", "b"}, "or") returns "a or b"
//   - JoinWithConj([]string{"a", "b", "c"}, "or") returns "a, b, or c"
//   - JoinWithConj([]string{"a", "b", "c"}, "and/or") returns "a, b, and/or c"
func JoinWithConj(words []string, conj string) string {
	return JoinWithSep(words, conj, ", ")
}

// JoinWithAutoSep combines a slice of strings into a grammatically correct English list
// with a custom conjunction, automatically choosing the separator based on content.
//
// If any item contains a comma, semicolons are used as separators ("; ").
// Otherwise, commas are used (", ").
//
// This is useful when you don't know in advance whether items contain commas.
//
// Examples:
//   - JoinWithAutoSep([]string{"a", "b", "c"}, "and") returns "a, b, and c"
//   - JoinWithAutoSep([]string{"Jan 1, 2020", "Feb 2, 2021"}, "and") returns "Jan 1, 2020; and Feb 2, 2021"
func JoinWithAutoSep(words []string, conj string) string {
	// Check if any item contains a comma
	hasComma := false
	for _, w := range words {
		if strings.Contains(w, ",") {
			hasComma = true
			break
		}
	}

	if hasComma {
		// For 2 items with commas, use semicolon before conjunction
		// (JoinWithSep doesn't add separator for 2-item case)
		if len(words) == 2 {
			return words[0] + "; " + conj + " " + words[1]
		}
		return JoinWithSep(words, conj, "; ")
	}
	return JoinWithSep(words, conj, ", ")
}

// JoinWithSep combines a slice of strings into a grammatically correct English list
// with a custom conjunction and separator.
//
// This is useful when list items themselves contain commas.
//
// Examples:
//   - JoinWithSep([]string{"a", "b", "c"}, "and", "; ") returns "a; b; and c"
//   - JoinWithSep([]string{"Jan 1, 2020", "Feb 2, 2021"}, "and", "; ") returns "Jan 1, 2020; and Feb 2, 2021"
func JoinWithSep(words []string, conj, sep string) string {
	switch len(words) {
	case 0:
		return ""
	case 1:
		return words[0]
	case 2:
		return words[0] + " " + conj + " " + words[1]
	default:
		return strings.Join(words[:len(words)-1], sep) + sep + conj + " " + words[len(words)-1]
	}
}
