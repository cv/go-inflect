package inflect

// verbSingularToPlural maps singular verb forms to plural forms.
var verbSingularToPlural = map[string]string{
	"is":      "are",
	"was":     "were",
	"has":     "have",
	"does":    "do",
	"goes":    "go",
	"isn't":   "aren't",
	"wasn't":  "weren't",
	"hasn't":  "haven't",
	"doesn't": "don't",
}

// verbPluralToSingular maps plural verb forms to singular forms.
var verbPluralToSingular = map[string]string{
	"are":     "is",
	"were":    "was",
	"have":    "has",
	"do":      "does",
	"go":      "goes",
	"aren't":  "isn't",
	"weren't": "wasn't",
	"haven't": "hasn't",
	"don't":   "doesn't",
}

// verbUnchanged contains verbs that don't change between singular and plural.
var verbUnchanged = map[string]bool{
	"can":     true,
	"could":   true,
	"may":     true,
	"might":   true,
	"must":    true,
	"shall":   true,
	"should":  true,
	"will":    true,
	"would":   true,
	"can't":   true,
	"won't":   true,
	"shan't":  true,
	"mustn't": true,
}

// adjSingularToPlural maps singular adjectives to plural forms.
var adjSingularToPlural = map[string]string{
	"this": "these",
	"that": "those",
	"a":    "some",
	"an":   "some",
	"my":   "our",
	"your": "your", // unchanged
	"her":  "their",
	"his":  "their",
	"its":  "their",
}

// adjPluralToSingular maps plural adjectives to singular forms.
// Note: Singular possessives depend on gender.
var adjPluralToSingular = map[string]string{
	"these": "this",
	"those": "that",
	"some":  "a", // or "an" depending on next word
	"our":   "my",
}

// adjPluralToSingularByGender maps possessive adjectives to singular by gender.
var adjPluralToSingularByGender = map[string]map[string]string{
	"their": {
		"m": "his",
		"f": "her",
		"n": "its",
		"t": "their",
	},
}
