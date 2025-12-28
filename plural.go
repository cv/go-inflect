package inflect

import "strings"

// changeToVesWords contains words ending in -f/-fe that change to -ves.
var changeToVesWords = map[string]bool{
	"calf": true, "elf": true, "half": true, "knife": true, "leaf": true,
	"life": true, "loaf": true, "self": true, "sheaf": true, "shelf": true,
	"thief": true, "wife": true, "wolf": true,
}

// oExceptionWords contains words ending in -o that just take -s (not -es).
var oExceptionWords = map[string]bool{
	"alto": true, "auto": true, "basso": true, "canto": true, "casino": true,
	"combo": true, "contralto": true, "disco": true, "dynamo": true,
	"embryo": true, "espresso": true, "euro": true, "fiasco": true,
	"ghetto": true, "inferno": true, "kilo": true, "limo": true,
	"maestro": true, "memo": true, "metro": true, "piano": true,
	"photo": true, "pimento": true, "polo": true, "poncho": true,
	"pro": true, "ratio": true, "rhino": true, "silo": true, "solo": true,
	"soprano": true, "stiletto": true, "studio": true, "taco": true,
	"tattoo": true, "tempo": true, "tornado": true, "torso": true,
	"tuxedo": true, "video": true, "virtuoso": true, "zero": true,
	"albino": true, "archipelago": true, "armadillo": true, "commando": true,
	"dodo": true, "flamingo": true, "grotto": true, "magneto": true,
	"manifesto": true, "mosquito": true, "motto": true, "otto": true,
	"placebo": true, "portfolio": true, "quarto": true, "stucco": true,
	"tobacco": true, "volcano": true,
}

// unchangedPlurals contains words that don't change in plural form.
// Note: Some animals like bison, buffalo are in herdAnimals instead,
// since they have both unchanged (classical) and -s (modern) forms.
var unchangedPlurals = map[string]bool{
	"aircraft": true, "cod": true, "deer": true, "fish": true,
	"moose": true, "offspring": true, "pike": true, "salmon": true,
	"series": true, "sheep": true, "shrimp": true, "species": true,
	"squid": true, "swine": true, "trout": true, "tuna": true,
}

// herdAnimals contains animals that have both unchanged (classical) and
// regular -s (modern) plural forms. When classicalHerd is enabled,
// these remain unchanged; otherwise they take -s.
// Examples: bison -> bison (classical) vs bisons (modern)
//
//	wildebeest -> wildebeest (classical) vs wildebeests (modern)
var herdAnimals = map[string]bool{
	"bison":      true,
	"buffalo":    true,
	"caribou":    true,
	"elk":        true,
	"grouse":     true,
	"antelope":   true,
	"wildebeest": true,
}

// classicalLatinPlurals contains words with classical Latin/Greek plural forms.
// These are used when classicalAncient is enabled.
// Key is singular, value is classical plural.
var classicalLatinPlurals = map[string]string{
	// -a -> -ae (Latin feminine)
	"formula":   "formulae",
	"antenna":   "antennae",
	"vertebra":  "vertebrae",
	"alumna":    "alumnae",
	"larva":     "larvae",
	"nebula":    "nebulae",
	"aurora":    "aurorae",
	"alga":      "algae",
	"amoeba":    "amoebae",
	"minutia":   "minutiae",
	"lacuna":    "lacunae",
	"persona":   "personae",
	"vita":      "vitae",
	"cornea":    "corneae",
	"retina":    "retinae",
	"hernia":    "herniae",
	"nausea":    "nauseae",
	"arena":     "arenae",
	"zona":      "zonae",
	"lamina":    "laminae",
	"nova":      "novae",
	"supernova": "supernovae",

	// -us -> -i (Latin masculine, second declension)
	// Note: some of these are already in irregularPlurals

	// -um -> -a (Latin neuter)
	// Note: some of these are already in irregularPlurals

	// -ex/-ix -> -ices (Latin)
	// Note: already in irregularPlurals

	// -is -> -es (Greek/Latin)
	// Note: already in irregularPlurals

	// Other classical forms
	"octopus":     "octopodes", // Greek: -pous -> -podes
	"platypus":    "platypodes",
	"hippopotami": "hippopotami", // Already in irregular (but classical is hippopotami)
	"opus":        "opera",
	"corpus":      "corpora",
	"genus":       "genera",
	"viscus":      "viscera",
}

// defaultIrregularPlurals contains the built-in irregular plural mappings.
// This is used to reset irregularPlurals to its original state.
var defaultIrregularPlurals = map[string]string{
	"child":       "children",
	"foot":        "feet",
	"goose":       "geese",
	"louse":       "lice",
	"man":         "men",
	"mouse":       "mice",
	"ox":          "oxen",
	"person":      "people",
	"tooth":       "teeth",
	"woman":       "women",
	"die":         "dice",
	"criterion":   "criteria",
	"phenomenon":  "phenomena",
	"analysis":    "analyses",
	"basis":       "bases",
	"crisis":      "crises",
	"diagnosis":   "diagnoses",
	"hypothesis":  "hypotheses",
	"oasis":       "oases",
	"parenthesis": "parentheses",
	"synopsis":    "synopses",
	"thesis":      "theses",
	"alumnus":     "alumni",
	"cactus":      "cacti",
	"focus":       "foci",
	"fungus":      "fungi",
	"nucleus":     "nuclei",
	"radius":      "radii",
	"stimulus":    "stimuli",
	"syllabus":    "syllabi",
	"bacterium":   "bacteria",
	"curriculum":  "curricula",
	"datum":       "data",
	"medium":      "media",
	"memorandum":  "memoranda",
	"millennium":  "millennia",
	"stadium":     "stadia",
	"stratum":     "strata",
	"appendix":    "appendices",
	"index":       "indices",
	"matrix":      "matrices",
	"vertex":      "vertices",
	"apex":        "apices",
}

// Plural returns the plural form of an English noun.
//
// Examples:
//   - Plural("cat") returns "cats"
//   - Plural("box") returns "boxes"
//   - Plural("child") returns "children"
//   - Plural("sheep") returns "sheep"
func Plural(word string) string {
	return defaultEngine.Plural(word)
}

// Plural returns the plural form of an English noun.
//
// Examples:
//   - e.Plural("cat") returns "cats"
//   - e.Plural("box") returns "boxes"
//   - e.Plural("child") returns "children"
//   - e.Plural("sheep") returns "sheep"
func (e *Engine) Plural(word string) string {
	if word == "" {
		return ""
	}

	lower := strings.ToLower(word)

	// Check for classical proper name handling when classicalNames is enabled.
	// Proper names (capitalized words) ending in 's' remain unchanged.
	// Examples: Jones -> Jones, Williams -> Williams
	if e.IsClassicalNames() && isProperNameEndingInS(word) {
		return word
	}

	// Check for classical Latin/Greek plurals when classicalAncient is enabled
	if e.IsClassical() {
		if plural, ok := classicalLatinPlurals[lower]; ok {
			return matchCase(word, plural)
		}
	}

	// Handle classicalPersons: person -> persons (instead of people)
	if e.IsClassicalPersons() && lower == "person" {
		return matchCase(word, "persons")
	}

	// Check for irregular plurals first
	e.mu.RLock()
	plural, ok := e.irregularPlurals[lower]
	e.mu.RUnlock()
	if ok {
		return matchCase(word, plural)
	}

	// Check for uncountable/unchanged words
	if unchangedPlurals[lower] {
		return word
	}

	// Check for herd animals (affected by classicalHerd flag)
	if herdAnimals[lower] {
		if e.IsClassicalHerd() {
			return word // unchanged in classical mode
		}
		// Modern mode: apply standard suffix rules (adds -s or -es)
		return applySuffixRules(word, lower)
	}

	// Check for words ending in -ese, -ois (nationalities that don't change)
	if strings.HasSuffix(lower, "ese") || strings.HasSuffix(lower, "ois") {
		return word
	}

	// Apply suffix rules
	return applySuffixRules(word, lower)
}

// applySuffixRules applies standard English pluralization suffix rules.
func applySuffixRules(word, lower string) string {
	// Words ending in -man -> -men
	if strings.HasSuffix(lower, "man") && !strings.HasSuffix(lower, "human") {
		return word[:len(word)-3] + matchCase(word[len(word)-3:], "men")
	}

	// Words ending in -s, -ss, -sh, -ch, -x, -z -> add -es
	if strings.HasSuffix(lower, "s") || strings.HasSuffix(lower, "ss") ||
		strings.HasSuffix(lower, "sh") || strings.HasSuffix(lower, "ch") ||
		strings.HasSuffix(lower, "x") || strings.HasSuffix(lower, "z") {
		return word + matchSuffix(word, "es")
	}

	// Words ending in consonant + y -> -ies
	// Exception: proper names (capitalized words like "Mary") just add -s
	if strings.HasSuffix(lower, "y") && len(lower) > 1 {
		beforeY := lower[len(lower)-2]
		if !isVowel(rune(beforeY)) {
			// Proper names just add -s: Mary -> Marys, not Maries
			if isProperName(word) {
				return word + matchSuffix(word, "s")
			}
			return word[:len(word)-1] + matchSuffix(word, "ies")
		}
	}

	// Words ending in -f or -fe -> -ves (with exceptions)
	if strings.HasSuffix(lower, "fe") {
		if shouldChangeF(lower) {
			return word[:len(word)-2] + matchSuffix(word, "ves")
		}
	} else if strings.HasSuffix(lower, "f") && !strings.HasSuffix(lower, "ff") {
		if shouldChangeF(lower) {
			return word[:len(word)-1] + matchSuffix(word, "ves")
		}
	}

	// Words ending in -o -> -oes (with exceptions)
	if strings.HasSuffix(lower, "o") && len(lower) > 1 {
		beforeO := lower[len(lower)-2]
		// Vowel + o -> just add s (radio, studio, zoo)
		if isVowel(rune(beforeO)) {
			return word + matchSuffix(word, "s")
		}
		// Check if it's an exception that just takes -s
		if oExceptionTakesS(lower) {
			return word + matchSuffix(word, "s")
		}
		return word + matchSuffix(word, "es")
	}

	// Default: add -s
	return word + matchSuffix(word, "s")
}

// shouldChangeF determines if a word ending in -f/-fe should change to -ves.
func shouldChangeF(lower string) bool {
	return changeToVesWords[lower]
}

// oExceptionTakesS returns true if a word ending in -o just takes -s.
func oExceptionTakesS(lower string) bool {
	return oExceptionWords[lower]
}
