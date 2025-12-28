package inflect

import "maps"

// pronounNominativePlural maps singular nominative pronouns to plural forms.
var pronounNominativePlural = map[string]string{
	"i":   "we",
	"he":  "they",
	"she": "they",
	"it":  "they",
}

// pronounAccusativePlural maps singular accusative pronouns to plural forms.
var pronounAccusativePlural = map[string]string{
	"me":  "us",
	"him": "them",
	"her": "them",
}

// pronounPossessivePlural maps singular possessive pronouns to plural forms.
var pronounPossessivePlural = map[string]string{
	"my":    "our",
	"mine":  "ours",
	"his":   "their",
	"hers":  "theirs",
	"its":   "their",
	"one's": "one's", // unchanged
}

// pronounReflexivePlural maps singular reflexive pronouns to plural forms.
var pronounReflexivePlural = map[string]string{
	"myself":   "ourselves",
	"yourself": "yourselves",
	"himself":  "themselves",
	"herself":  "themselves",
	"itself":   "themselves",
	"oneself":  "oneselves",
}

// allPronounsToPlural combines all pronoun plural mappings.
var allPronounsToPlural = buildAllPronounsToPlural()

func buildAllPronounsToPlural() map[string]string {
	m := make(map[string]string)
	maps.Copy(m, pronounNominativePlural)
	maps.Copy(m, pronounAccusativePlural)
	maps.Copy(m, pronounPossessivePlural)
	maps.Copy(m, pronounReflexivePlural)
	return m
}

// pronounNominativeSingularByGender maps plural nominative pronouns to singular forms.
// The actual singular form depends on the current gender setting.
var pronounNominativeSingularByGender = map[string]map[string]string{
	"we": {
		"m": "I",
		"f": "I",
		"n": "I",
		"t": "I",
	},
	"they": {
		"m": "he",
		"f": "she",
		"n": "it",
		"t": "they",
	},
}

// pronounAccusativeSingularByGender maps plural accusative pronouns to singular.
var pronounAccusativeSingularByGender = map[string]map[string]string{
	"us": {
		"m": "me",
		"f": "me",
		"n": "me",
		"t": "me",
	},
	"them": {
		"m": "him",
		"f": "her",
		"n": "it",
		"t": "them",
	},
}

// pronounPossessiveSingularByGender maps plural possessive pronouns to singular.
var pronounPossessiveSingularByGender = map[string]map[string]string{
	"our": {
		"m": "my",
		"f": "my",
		"n": "my",
		"t": "my",
	},
	"ours": {
		"m": "mine",
		"f": "mine",
		"n": "mine",
		"t": "mine",
	},
	"their": {
		"m": "his",
		"f": "her",
		"n": "its",
		"t": "their",
	},
	"theirs": {
		"m": "his",
		"f": "hers",
		"n": "its",
		"t": "theirs",
	},
}

// pronounReflexiveSingularByGender maps plural reflexive pronouns to singular.
var pronounReflexiveSingularByGender = map[string]map[string]string{
	"ourselves": {
		"m": "myself",
		"f": "myself",
		"n": "myself",
		"t": "myself",
	},
	"yourselves": {
		"m": "yourself",
		"f": "yourself",
		"n": "yourself",
		"t": "yourself",
	},
	"themselves": {
		"m": "himself",
		"f": "herself",
		"n": "itself",
		"t": "themself",
	},
}
