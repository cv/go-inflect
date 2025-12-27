// Package inflect provides English language inflection utilities.
//
// It offers functions for pluralization, singularization, indefinite article
// selection (a/an), number-to-words conversion, ordinals, and more.
//
// This is a Go port of the Python inflect library.
package inflect

import (
	"fmt"
	"strings"
	"unicode"
)

// defaultNum stores the default count for number-related functions.
// A value of 0 indicates no default is set.
var defaultNum int

// classicalMode controls whether to prefer classical Latin/Greek plurals.
// When true, words like "formula" become "formulae" instead of "formulas".
// Default is false (modern English plurals).
// Deprecated: Use classicalAll and individual classical options instead.
var classicalMode bool

// Classical pluralization option flags.
// These control various aspects of classical (Latin/Greek) pluralization.
var (
	// classicalAll is the master switch for all classical options.
	// When true, all classical pluralization rules are enabled.
	classicalAll bool

	// classicalZero controls pluralization of zero count.
	// When true: "0 cat" (singular); when false: "0 cats" (plural).
	classicalZero bool

	// classicalHerd controls pluralization of herd animals.
	// When true: "wildebeest" (unchanged); when false: "wildebeests".
	classicalHerd bool

	// classicalNames controls pluralization of proper names.
	// When true: classical proper name pluralization is used.
	classicalNames bool

	// classicalAncient controls Latin/Greek plural forms.
	// When true: "formula" -> "formulae"; when false: "formula" -> "formulas".
	classicalAncient bool

	// classicalPersons controls person/people pluralization.
	// When true: "person" -> "persons"; when false: "person" -> "people".
	classicalPersons bool
)

// customAWords stores words that should be forced to use "a" instead of "an".
// The keys are lowercase versions of the words/patterns.
var customAWords = make(map[string]bool)

// customAnWords stores words that should be forced to use "an" instead of "a".
// The keys are lowercase versions of the words/patterns.
var customAnWords = make(map[string]bool)

// silentHWords contains words starting with silent 'h' that take "an".
var silentHWords = map[string]bool{
	"honest": true, "heir": true, "heiress": true, "heirloom": true,
	"honor": true, "honour": true, "hour": true, "hourly": true,
}

// lowercaseAbbrevs contains lowercase abbreviations pronounced letter-by-letter.
var lowercaseAbbrevs = map[string]bool{
	"mpeg": true, "jpeg": true, "gif": true, "sql": true, "html": true,
	"xml": true, "fbi": true, "cia": true, "nsa": true,
}

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

// feWordBases contains base words whose singular ends in -fe (plural is -ves).
var feWordBases = map[string]bool{
	"kni": true, "wi": true, "li": true,
}

// doubleConsonantWords contains multi-syllable words that double the final consonant.
var doubleConsonantWords = map[string]bool{
	"admit": true, "begin": true, "commit": true, "compel": true,
	"confer": true, "control": true, "defer": true, "deter": true,
	"equip": true, "excel": true, "expel": true, "forget": true,
	"incur": true, "occur": true, "omit": true, "patrol": true,
	"permit": true, "prefer": true, "propel": true, "rebel": true,
	"recur": true, "refer": true, "regret": true, "repel": true,
	"submit": true, "transfer": true, "transmit": true, "upset": true,
}

// isAllUpper checks if all letters in a word are uppercase.
func isAllUpper(word string) bool {
	for _, r := range word {
		if unicode.IsLetter(r) && !unicode.IsUpper(r) {
			return false
		}
	}
	return true
}

// isProperName checks if a word is a proper name.
// A proper name is detected by having a capitalized first letter and not being
// all uppercase. Examples: "Jones", "Mary", "Smith".
func isProperName(word string) bool {
	if len(word) < 2 {
		return false
	}

	// Check if the first letter is uppercase (proper name indicator)
	firstRune := []rune(word)[0]
	if !unicode.IsUpper(firstRune) {
		return false
	}

	// If all uppercase, it's likely an acronym, not a proper name
	if isAllUpper(word) {
		return false
	}

	return true
}

// isProperNameEndingInS checks if a word is a proper name ending in 's'.
// A proper name is detected by having a capitalized first letter and not being
// all uppercase. Examples: "Jones", "Williams", "Hastings".
func isProperNameEndingInS(word string) bool {
	if !isProperName(word) {
		return false
	}

	// Check if the word ends in 's' or 'S'
	lastRune := rune(word[len(word)-1])
	return unicode.ToLower(lastRune) == 's'
}

// An returns the word prefixed with the appropriate indefinite article ("a" or "an").
//
// The selection follows standard English rules:
//   - Use "an" before vowel sounds: "an apple", "an hour"
//   - Use "a" before consonant sounds: "a cat", "a university"
//
// Special cases handled:
//   - Silent 'h': "an honest person"
//   - Vowels with consonant sounds: "a Ukrainian", "a unanimous decision"
//   - Abbreviations: "a YAML file", "a JSON object"
//
// Custom patterns defined via DefA() and DefAn() take precedence over default rules.
func An(word string) string {
	if word == "" {
		return ""
	}

	// Get the first word for pattern matching
	firstWord := strings.Fields(word)[0]
	lowerFirst := strings.ToLower(firstWord)

	// Check custom "a" patterns first (highest priority)
	if customAWords[lowerFirst] {
		return "a " + word
	}

	// Check custom "an" patterns second
	if customAnWords[lowerFirst] {
		return "an " + word
	}

	// Fall back to default rules
	if needsAn(word) {
		return "an " + word
	}
	return "a " + word
}

// needsAn determines if a word/phrase should be preceded by "an" (vs "a").
func needsAn(text string) bool {
	// Get the first word to analyze
	firstWord := strings.Fields(text)[0]
	lower := strings.ToLower(firstWord)

	// Check for silent 'h' words that take "an"
	for h := range silentHWords {
		if strings.HasPrefix(lower, h) {
			return true
		}
	}

	// Check for abbreviations/acronyms (all uppercase or known patterns)
	if isAbbreviation(firstWord) {
		return abbreviationNeedsAn(firstWord)
	}

	// Check for known lowercase abbreviations pronounced letter-by-letter
	if lowercaseAbbrevs[lower] {
		return abbreviationNeedsAn(strings.ToUpper(lower))
	}

	// Check for special vowel patterns that sound like consonants (take "a")
	// "uni-" sounds like "yoo-", "eu-" sounds like "yoo-", etc.
	consonantVowelPatterns := []string{
		"uni", "upon", "use", "used", "user", "using", "usual",
		"usu", "uran", "uret", "euro", "ewe", "onc", "one",
		"onet", // onetime
	}
	for _, pat := range consonantVowelPatterns {
		if strings.HasPrefix(lower, pat) {
			return false
		}
	}

	// Check for "U" followed by consonant then vowel pattern (sounds like "you")
	// e.g., Ugandan, Ukrainian, Unabomber, unanimous
	if len(lower) >= 2 && lower[0] == 'u' {
		if isConsonantYSound(lower) {
			return false
		}
	}

	// Special case: single letters
	if len(firstWord) == 1 {
		return isVowelSound(rune(lower[0]))
	}

	// Default: check if first letter is a vowel
	first := rune(lower[0])
	return isVowelSound(first)
}

// isAbbreviation checks if a word appears to be an abbreviation/acronym.
func isAbbreviation(word string) bool {
	if len(word) < 2 {
		return false
	}
	return isAllUpper(word)
}

// abbreviationNeedsAn checks if an abbreviation should take "an".
// This depends on how the first letter is pronounced.
func abbreviationNeedsAn(abbrev string) bool {
	if len(abbrev) == 0 {
		return false
	}

	// Letters whose names start with vowel sounds:
	// A (ay), E (ee), F (eff), H (aitch), I (eye), L (ell), M (em),
	// N (en), O (oh), R (ar), S (ess), X (ex)
	vowelSoundLetters := "AEFHILMNORSX"
	first := unicode.ToUpper(rune(abbrev[0]))
	return strings.ContainsRune(vowelSoundLetters, first)
}

// isConsonantYSound checks if a word starting with 'u' has a consonant "y" sound.
// e.g., "unanimous" starts with "yoo-" sound.
func isConsonantYSound(lower string) bool {
	if len(lower) < 2 {
		return false
	}

	// Explicit patterns that have the "you" sound
	youPatterns := []string{
		"uga", "ukr", "ula", "ule", "uli", "ulo", "ulu",
		"una", "uni", "uno", "unu", "ura", "ure", "uri", "uro", "uru",
		"usa", "use", "usi", "uso", "usu", "uta", "ute", "uti", "uto", "utu",
	}

	if len(lower) >= 3 {
		prefix := lower[:3]
		for _, pat := range youPatterns {
			if prefix == pat {
				return true
			}
		}
	}

	return false
}

// isVowelSound checks if a letter represents a vowel sound.
func isVowelSound(r rune) bool {
	return strings.ContainsRune("aeiou", unicode.ToLower(r))
}

// A is an alias for An - returns word prefixed with appropriate indefinite article.
func A(word string) string {
	return An(word)
}

// DefA defines a custom pattern that forces "a" instead of "an" for a word.
//
// The pattern is matched against the first word of the input (case-insensitive).
// Custom "a" patterns take precedence over custom "an" patterns and default rules.
//
// Examples:
//
//	DefA("ape")
//	An("ape") // returns "a ape" instead of "an ape"
//	An("Ape") // returns "a Ape" (case-insensitive matching)
func DefA(word string) {
	lower := strings.ToLower(word)
	customAWords[lower] = true
	// Remove from customAnWords if present to avoid conflicts
	delete(customAnWords, lower)
}

// DefAn defines a custom pattern that forces "an" instead of "a" for a word.
//
// The pattern is matched against the first word of the input (case-insensitive).
// Custom "an" patterns take precedence over default rules but not custom "a" patterns.
//
// Examples:
//
//	DefAn("hero")
//	An("hero") // returns "an hero" instead of "a hero"
//	An("Hero") // returns "an Hero" (case-insensitive matching)
func DefAn(word string) {
	lower := strings.ToLower(word)
	customAnWords[lower] = true
	// Remove from customAWords if present to avoid conflicts
	delete(customAWords, lower)
}

// UndefA removes a custom "a" pattern.
//
// Returns true if the pattern was removed, false if it didn't exist.
//
// Examples:
//
//	DefA("ape")
//	An("ape") // returns "a ape"
//	UndefA("ape")
//	An("ape") // returns "an ape" (default rule)
func UndefA(word string) bool {
	lower := strings.ToLower(word)
	if customAWords[lower] {
		delete(customAWords, lower)
		return true
	}
	return false
}

// UndefAn removes a custom "an" pattern.
//
// Returns true if the pattern was removed, false if it didn't exist.
//
// Examples:
//
//	DefAn("hero")
//	An("hero") // returns "an hero"
//	UndefAn("hero")
//	An("hero") // returns "a hero" (default rule)
func UndefAn(word string) bool {
	lower := strings.ToLower(word)
	if customAnWords[lower] {
		delete(customAnWords, lower)
		return true
	}
	return false
}

// DefAReset resets all custom a/an patterns to defaults (empty).
//
// This removes all custom patterns added via DefA() and DefAn().
//
// Example:
//
//	DefA("ape")
//	DefAn("hero")
//	DefAReset()
//	An("ape")  // returns "an ape" (default rule)
//	An("hero") // returns "a hero" (default rule)
func DefAReset() {
	customAWords = make(map[string]bool)
	customAnWords = make(map[string]bool)
}

// ClassicalAll enables or disables all classical pluralization options at once.
//
// This is a master switch that sets all classical options:
//   - classicalAll: master switch
//   - classicalZero: 0 cat vs 0 cats
//   - classicalHerd: wildebeest vs wildebeests
//   - classicalNames: proper name pluralization
//   - classicalAncient: Latin/Greek forms (formulae vs formulas)
//   - classicalPersons: persons vs people
//
// When enabled (true), Plural() prefers classical Latin/Greek plural forms:
//   - formula -> formulae (instead of formulas)
//   - antenna -> antennae (instead of antennas)
//   - vertebra -> vertebrae (instead of vertebras)
//   - alumna -> alumnae (instead of alumnas)
//
// When disabled (false, the default), modern English plurals are used.
//
// Examples:
//
//	ClassicalAll(true)
//	Plural("formula") // returns "formulae"
//	ClassicalAll(false)
//	Plural("formula") // returns "formulas"
func ClassicalAll(enabled bool) {
	classicalAll = enabled
	classicalZero = enabled
	classicalHerd = enabled
	classicalNames = enabled
	classicalAncient = enabled
	classicalPersons = enabled
	// Also set the legacy classicalMode for backward compatibility
	classicalMode = enabled
}

// Classical enables or disables classical pluralization mode.
//
// This is an alias for ClassicalAll() for backward compatibility.
// It sets all classical pluralization options at once.
//
// When enabled (true), Plural() prefers classical Latin/Greek plural forms:
//   - formula -> formulae (instead of formulas)
//   - antenna -> antennae (instead of antennas)
//   - vertebra -> vertebrae (instead of vertebras)
//   - alumna -> alumnae (instead of alumnas)
//
// When disabled (false, the default), modern English plurals are used.
//
// Examples:
//
//	Classical(true)
//	Plural("formula") // returns "formulae"
//	Classical(false)
//	Plural("formula") // returns "formulas"
func Classical(enabled bool) {
	ClassicalAll(enabled)
}

// IsClassicalAll returns whether all classical pluralization options are enabled.
//
// Returns true only if all classical options are enabled, false otherwise.
//
// Examples:
//
//	IsClassicalAll() // returns false (default)
//	ClassicalAll(true)
//	IsClassicalAll() // returns true
//	ClassicalAll(false)
//	IsClassicalAll() // returns false
func IsClassicalAll() bool {
	return classicalAll && classicalZero && classicalHerd &&
		classicalNames && classicalAncient && classicalPersons
}

// IsClassical returns whether classical pluralization mode is enabled.
//
// Returns true if Classical(true) or ClassicalAll(true) was called, false otherwise.
// This checks the classicalAncient flag which controls Latin/Greek plurals.
//
// Examples:
//
//	IsClassical() // returns false (default)
//	Classical(true)
//	IsClassical() // returns true
//	Classical(false)
//	IsClassical() // returns false
func IsClassical() bool {
	return classicalAncient || classicalMode
}

// ClassicalAncient enables or disables classical Latin/Greek plural forms.
//
// This controls the classicalAncient flag independently of other classical
// options like classicalZero, classicalHerd, classicalNames, and classicalPersons.
//
// When enabled (true), Plural() prefers classical Latin/Greek plural forms:
//   - formula -> formulae (instead of formulas)
//   - antenna -> antennae (instead of antennas)
//   - vertebra -> vertebrae (instead of vertebras)
//   - alumna -> alumnae (instead of alumnas)
//
// When disabled (false, the default), modern English plurals are used.
//
// Note: This also controls the legacy classicalMode flag, since both affect
// Latin/Greek plural forms.
//
// Examples:
//
//	ClassicalAncient(true)
//	Plural("formula") // returns "formulae"
//	ClassicalAncient(false)
//	Plural("formula") // returns "formulas"
func ClassicalAncient(enabled bool) {
	classicalAncient = enabled
	// Also set the legacy classicalMode to keep them in sync
	classicalMode = enabled
}

// IsClassicalAncient returns whether classical Latin/Greek plural forms are enabled.
//
// Returns true if ClassicalAncient(true), Classical(true), or ClassicalAll(true)
// was called, false otherwise.
//
// Examples:
//
//	IsClassicalAncient() // returns false (default)
//	ClassicalAncient(true)
//	IsClassicalAncient() // returns true
//	ClassicalAncient(false)
//	IsClassicalAncient() // returns false
func IsClassicalAncient() bool {
	return classicalAncient
}

// ClassicalZero enables or disables classical zero count pluralization.
//
// This controls the classicalZero flag independently of other classical
// options like classicalHerd, classicalNames, classicalAncient, and classicalPersons.
//
// When enabled (true), No() uses singular form for zero count:
//   - No("cat", 0) -> "no cat"
//
// When disabled (false, the default), No() uses plural form for zero count:
//   - No("cat", 0) -> "no cats"
//
// Examples:
//
//	ClassicalZero(true)
//	No("cat", 0) // returns "no cat"
//	ClassicalZero(false)
//	No("cat", 0) // returns "no cats"
func ClassicalZero(enabled bool) {
	classicalZero = enabled
}

// IsClassicalZero returns whether classical zero count pluralization is enabled.
//
// Returns true if ClassicalZero(true) or ClassicalAll(true) was called,
// false otherwise.
//
// Examples:
//
//	IsClassicalZero() // returns false (default)
//	ClassicalZero(true)
//	IsClassicalZero() // returns true
//	ClassicalZero(false)
//	IsClassicalZero() // returns false
func IsClassicalZero() bool {
	return classicalZero
}

// ClassicalPersons enables or disables classical person/persons pluralization.
//
// This controls the classicalPersons flag independently of other classical
// options like classicalZero, classicalHerd, classicalNames, and classicalAncient.
//
// When enabled (true), Plural() uses "persons" as the plural of "person":
//   - person -> persons
//
// When disabled (false, the default), the irregular plural "people" is used:
//   - person -> people
//
// Examples:
//
//	ClassicalPersons(true)
//	Plural("person") // returns "persons"
//	ClassicalPersons(false)
//	Plural("person") // returns "people"
func ClassicalPersons(enabled bool) {
	classicalPersons = enabled
}

// IsClassicalPersons returns whether classical person/persons pluralization is enabled.
//
// Returns true if ClassicalPersons(true) or ClassicalAll(true) was called,
// false otherwise.
//
// Examples:
//
//	IsClassicalPersons() // returns false (default)
//	ClassicalPersons(true)
//	IsClassicalPersons() // returns true
//	ClassicalPersons(false)
//	IsClassicalPersons() // returns false
func IsClassicalPersons() bool {
	return classicalPersons
}

// ClassicalHerd enables or disables classical herd animal pluralization.
//
// This controls the classicalHerd flag independently of other classical
// options like classicalZero, classicalNames, classicalAncient, and classicalPersons.
//
// When enabled (true), Plural() uses unchanged forms for herd animals:
//   - bison -> bison
//   - buffalo -> buffalo
//   - wildebeest -> wildebeest
//
// When disabled (false, the default), modern English plurals are used:
//   - bison -> bisons
//   - buffalo -> buffaloes
//   - wildebeest -> wildebeests
//
// Note: Some animals like deer, sheep, moose always remain unchanged
// regardless of this setting.
//
// Examples:
//
//	ClassicalHerd(true)
//	Plural("wildebeest") // returns "wildebeest"
//	ClassicalHerd(false)
//	Plural("wildebeest") // returns "wildebeests"
func ClassicalHerd(enabled bool) {
	classicalHerd = enabled
}

// IsClassicalHerd returns whether classical herd animal pluralization is enabled.
//
// Returns true if ClassicalHerd(true) or ClassicalAll(true) was called,
// false otherwise.
//
// Examples:
//
//	IsClassicalHerd() // returns false (default)
//	ClassicalHerd(true)
//	IsClassicalHerd() // returns true
//	ClassicalHerd(false)
//	IsClassicalHerd() // returns false
func IsClassicalHerd() bool {
	return classicalHerd
}

// ClassicalNames enables or disables classical proper name pluralization.
//
// This controls the classicalNames flag independently of other classical
// options like classicalZero, classicalHerd, classicalAncient, and classicalPersons.
//
// When enabled (true), Plural() leaves proper names ending in 's' unchanged:
//   - Jones -> Jones (unchanged)
//   - Williams -> Williams (unchanged)
//   - Hastings -> Hastings (unchanged)
//
// When disabled (false, the default), regular pluralization rules apply:
//   - Jones -> Joneses
//   - Williams -> Williamses
//   - Hastings -> Hastingses
//
// Note: This only affects capitalized words ending in 's'. Other proper names
// like "Mary" still pluralize normally (Mary -> Marys).
//
// Examples:
//
//	ClassicalNames(true)
//	Plural("Jones") // returns "Jones"
//	ClassicalNames(false)
//	Plural("Jones") // returns "Joneses"
func ClassicalNames(enabled bool) {
	classicalNames = enabled
}

// IsClassicalNames returns whether classical proper name pluralization is enabled.
//
// Returns true if ClassicalNames(true) or ClassicalAll(true) was called,
// false otherwise.
//
// Examples:
//
//	IsClassicalNames() // returns false (default)
//	ClassicalNames(true)
//	IsClassicalNames() // returns true
//	ClassicalNames(false)
//	IsClassicalNames() // returns false
func IsClassicalNames() bool {
	return classicalNames
}

// Plural returns the plural form of an English noun.
//
// Examples:
//   - Plural("cat") returns "cats"
//   - Plural("box") returns "boxes"
//   - Plural("child") returns "children"
//   - Plural("sheep") returns "sheep"
func Plural(word string) string {
	if word == "" {
		return ""
	}

	lower := strings.ToLower(word)

	// Check for classical proper name handling when classicalNames is enabled.
	// Proper names (capitalized words) ending in 's' remain unchanged.
	// Examples: Jones -> Jones, Williams -> Williams
	if classicalNames && isProperNameEndingInS(word) {
		return word
	}

	// Check for classical Latin/Greek plurals when classicalAncient is enabled
	if classicalAncient || classicalMode {
		if plural, ok := classicalLatinPlurals[lower]; ok {
			return matchCase(word, plural)
		}
	}

	// Handle classicalPersons: person -> persons (instead of people)
	if classicalPersons && lower == "person" {
		return matchCase(word, "persons")
	}

	// Check for irregular plurals first
	if plural, ok := irregularPlurals[lower]; ok {
		return matchCase(word, plural)
	}

	// Check for uncountable/unchanged words
	if unchangedPlurals[lower] {
		return word
	}

	// Check for herd animals (affected by classicalHerd flag)
	if herdAnimals[lower] {
		if classicalHerd {
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

// matchSuffix returns the suffix in uppercase if the word is all uppercase.
func matchSuffix(word, suffix string) string {
	if isAllUpper(word) {
		return strings.ToUpper(suffix)
	}
	return suffix
}

// isVowel checks if a rune is a vowel.
func isVowel(r rune) bool {
	return strings.ContainsRune("aeiouAEIOU", r)
}

// shouldChangeF determines if a word ending in -f/-fe should change to -ves.
func shouldChangeF(lower string) bool {
	return changeToVesWords[lower]
}

// oExceptionTakesS returns true if a word ending in -o just takes -s.
func oExceptionTakesS(lower string) bool {
	return oExceptionWords[lower]
}

// matchCase adjusts the replacement to match the case pattern of the original.
func matchCase(original, replacement string) string {
	if len(original) == 0 || len(replacement) == 0 {
		return replacement
	}

	// Check if original is all uppercase
	if isAllUpper(original) {
		return strings.ToUpper(replacement)
	}

	// Check if original starts with uppercase
	firstRune := []rune(original)[0]
	if unicode.IsUpper(firstRune) {
		runes := []rune(replacement)
		runes[0] = unicode.ToUpper(runes[0])
		return string(runes)
	}

	return replacement
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

// irregularPlurals maps singular forms to their irregular plural forms.
var irregularPlurals = copyMap(defaultIrregularPlurals)

// singularIrregulars maps plural forms to their singular forms (reverse of irregularPlurals).
var singularIrregulars = buildSingularIrregulars()

// copyMap creates a shallow copy of a map.
func copyMap(src map[string]string) map[string]string {
	dst := make(map[string]string, len(src))
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

// buildSingularIrregulars builds the reverse mapping from irregularPlurals.
func buildSingularIrregulars() map[string]string {
	m := make(map[string]string, len(irregularPlurals))
	for singular, plural := range irregularPlurals {
		m[plural] = singular
	}
	return m
}

// DefNoun defines a custom noun pluralization rule.
//
// The singular and plural forms are stored in lowercase, and subsequent calls
// to Plural() and Singular() will use this custom rule with case preservation.
//
// Examples:
//
//	DefNoun("foo", "foos")
//	Plural("foo") // returns "foos"
//	Plural("Foo") // returns "Foos"
//	Singular("foos") // returns "foo"
func DefNoun(singular, plural string) {
	lower := strings.ToLower(singular)
	lowerPlural := strings.ToLower(plural)
	irregularPlurals[lower] = lowerPlural
	singularIrregulars[lowerPlural] = lower
}

// UndefNoun removes a custom noun pluralization rule.
//
// This removes only user-defined rules; it cannot remove built-in irregular
// plurals. To restore a built-in rule that was overwritten, use DefNounReset().
//
// Returns true if the rule was removed, false if it didn't exist or was a
// built-in rule.
//
// Examples:
//
//	DefNoun("foo", "foos")
//	Plural("foo") // returns "foos"
//	UndefNoun("foo")
//	Plural("foo") // returns "foos" (standard rule)
func UndefNoun(singular string) bool {
	lower := strings.ToLower(singular)

	// Check if this is a built-in rule
	if _, isBuiltIn := defaultIrregularPlurals[lower]; isBuiltIn {
		return false
	}

	// Check if the rule exists
	plural, exists := irregularPlurals[lower]
	if !exists {
		return false
	}

	// Remove from both maps
	delete(irregularPlurals, lower)
	delete(singularIrregulars, plural)
	return true
}

// DefNounReset resets all noun pluralization rules to their defaults.
//
// This removes all custom rules added via DefNoun() and restores any
// built-in rules that may have been overwritten.
//
// Example:
//
//	DefNoun("child", "childs")  // override built-in
//	DefNoun("foo", "foos")      // add custom
//	DefNounReset()
//	Plural("child") // returns "children" (restored)
//	Plural("foo")   // returns "foos" (standard rule, custom removed)
func DefNounReset() {
	irregularPlurals = copyMap(defaultIrregularPlurals)
	singularIrregulars = buildSingularIrregulars()
}

// customVerbs stores custom verb conjugation rules (singular -> plural).
// This is a placeholder map for future verb conjugation support.
var customVerbs = make(map[string]string)

// customVerbsReverse stores reverse verb conjugation rules (plural -> singular).
var customVerbsReverse = make(map[string]string)

// DefVerb defines a custom verb conjugation rule.
//
// NOTE: This is a placeholder stub for future implementation.
// Full verb conjugation is not yet implemented; this function only stores
// the singular/plural pair in internal maps for future use.
//
// The singular and plural forms are stored in lowercase.
//
// Examples:
//
//	DefVerb("run", "runs")
//	DefVerb("be", "are")
func DefVerb(singular, plural string) {
	lower := strings.ToLower(singular)
	lowerPlural := strings.ToLower(plural)
	customVerbs[lower] = lowerPlural
	customVerbsReverse[lowerPlural] = lower
}

// UndefVerb removes a custom verb conjugation rule.
//
// NOTE: This is a placeholder stub for future implementation.
//
// Returns true if the rule was removed, false if it didn't exist.
//
// Examples:
//
//	DefVerb("run", "runs")
//	UndefVerb("run") // returns true
//	UndefVerb("walk") // returns false (not defined)
func UndefVerb(singular string) bool {
	lower := strings.ToLower(singular)
	plural, exists := customVerbs[lower]
	if !exists {
		return false
	}
	delete(customVerbs, lower)
	delete(customVerbsReverse, plural)
	return true
}

// DefVerbReset resets all custom verb conjugation rules.
//
// NOTE: This is a placeholder stub for future implementation.
//
// This removes all custom rules added via DefVerb().
func DefVerbReset() {
	customVerbs = make(map[string]string)
	customVerbsReverse = make(map[string]string)
}

// customAdjs stores custom adjective pluralization rules (singular -> plural).
// This is a placeholder map for future adjective pluralization support.
var customAdjs = make(map[string]string)

// customAdjsReverse stores reverse adjective rules (plural -> singular).
var customAdjsReverse = make(map[string]string)

// DefAdj defines a custom adjective pluralization rule.
//
// NOTE: This is a placeholder stub for future implementation.
// Full adjective pluralization is not yet implemented; this function only stores
// the singular/plural pair in internal maps for future use.
//
// The singular and plural forms are stored in lowercase.
//
// Examples:
//
//	DefAdj("big", "bigs")
//	DefAdj("happy", "happies")
func DefAdj(singular, plural string) {
	lower := strings.ToLower(singular)
	lowerPlural := strings.ToLower(plural)
	customAdjs[lower] = lowerPlural
	customAdjsReverse[lowerPlural] = lower
}

// UndefAdj removes a custom adjective pluralization rule.
//
// NOTE: This is a placeholder stub for future implementation.
//
// Returns true if the rule was removed, false if it didn't exist.
//
// Examples:
//
//	DefAdj("big", "bigs")
//	UndefAdj("big") // returns true
//	UndefAdj("small") // returns false (not defined)
func UndefAdj(singular string) bool {
	lower := strings.ToLower(singular)
	plural, exists := customAdjs[lower]
	if !exists {
		return false
	}
	delete(customAdjs, lower)
	delete(customAdjsReverse, plural)
	return true
}

// DefAdjReset resets all custom adjective pluralization rules.
//
// NOTE: This is a placeholder stub for future implementation.
//
// This removes all custom rules added via DefAdj().
func DefAdjReset() {
	customAdjs = make(map[string]string)
	customAdjsReverse = make(map[string]string)
}

// Singular returns the singular form of an English noun.
//
// Examples:
//   - Singular("cats") returns "cat"
//   - Singular("boxes") returns "box"
//   - Singular("children") returns "child"
//   - Singular("sheep") returns "sheep"
func Singular(word string) string {
	if word == "" {
		return ""
	}

	lower := strings.ToLower(word)

	// Check for irregular plurals first
	if singular, ok := singularIrregulars[lower]; ok {
		return matchCase(word, singular)
	}

	// Check for uncountable/unchanged words
	if unchangedPlurals[lower] {
		return word
	}

	// Check for words ending in -ese, -ois (nationalities that don't change)
	if strings.HasSuffix(lower, "ese") || strings.HasSuffix(lower, "ois") {
		return word
	}

	// Apply suffix rules to singularize
	return applySingularSuffixRules(word, lower)
}

// applySingularSuffixRules applies standard English singularization suffix rules.
func applySingularSuffixRules(word, lower string) string {
	n := len(lower)

	// Words ending in -men -> -man (but not "women" which is irregular)
	if strings.HasSuffix(lower, "men") && n > 3 {
		return word[:len(word)-3] + matchCase(word[len(word)-3:], "man")
	}

	// Words ending in -ves -> -f or -fe
	if strings.HasSuffix(lower, "ves") && n > 3 {
		base := lower[:n-3]
		// Check if original was -fe (knives -> knife, wives -> wife)
		if singularEndsInFe(base) {
			return word[:len(word)-3] + matchSuffix(word, "fe")
		}
		// Otherwise was -f (wolves -> wolf, leaves -> leaf)
		return word[:len(word)-3] + matchSuffix(word, "f")
	}

	// Words ending in -ies (consonant + ies) -> -y
	if strings.HasSuffix(lower, "ies") && n > 3 {
		beforeIes := lower[n-4]
		if !isVowel(rune(beforeIes)) {
			return word[:len(word)-3] + matchSuffix(word, "y")
		}
	}

	// Words ending in -es after sibilants (s, ss, sh, ch, x, z)
	if strings.HasSuffix(lower, "es") && n > 2 {
		base := lower[:n-2]
		// -sses -> -ss (classes -> class)
		if strings.HasSuffix(base, "ss") {
			return word[:len(word)-2]
		}
		// -shes -> -sh (bushes -> bush)
		if strings.HasSuffix(base, "sh") {
			return word[:len(word)-2]
		}
		// -ches -> -ch (churches -> church)
		if strings.HasSuffix(base, "ch") {
			return word[:len(word)-2]
		}
		// -xes -> -x (boxes -> box)
		if strings.HasSuffix(base, "x") {
			return word[:len(word)-2]
		}
		// -zes -> -z (buzzes -> buzz)
		if strings.HasSuffix(base, "zz") {
			return word[:len(word)-2]
		}
		// -oes -> -o (heroes -> hero, potatoes -> potato)
		// But not words like "shoes" -> "shoe"
		if strings.HasSuffix(base, "o") && !oExceptionTakesS(base) {
			// Check if this looks like a word that would have taken -oes
			beforeO := base[len(base)-2]
			if !isVowel(rune(beforeO)) {
				return word[:len(word)-2]
			}
		}
		// Single -s ending with -es: buses -> bus
		if strings.HasSuffix(base, "s") && !strings.HasSuffix(base, "ss") {
			return word[:len(word)-2]
		}
	}

	// Default: remove trailing -s
	if strings.HasSuffix(lower, "s") && n > 1 {
		// Don't remove -s from words ending in -ss
		if strings.HasSuffix(lower, "ss") {
			return word
		}
		return word[:len(word)-1]
	}

	// Word doesn't appear to be plural
	return word
}

// singularEndsInFe checks if a base word's singular form ends in -fe.
func singularEndsInFe(base string) bool {
	return feWordBases[base]
}

// Ordinal converts an integer to its ordinal string representation.
//
// Examples:
//   - Ordinal(1) returns "1st"
//   - Ordinal(2) returns "2nd"
//   - Ordinal(3) returns "3rd"
//   - Ordinal(11) returns "11th"
//   - Ordinal(21) returns "21st"
//   - Ordinal(-1) returns "-1st"
func Ordinal(n int) string {
	suffix := ordinalSuffix(n)
	return fmt.Sprintf("%d%s", n, suffix)
}

// ordinalSuffix returns the ordinal suffix for a number.
func ordinalSuffix(n int) string {
	// Handle negative numbers by using absolute value
	if n < 0 {
		n = -n
	}

	// Special case: 11, 12, 13 always use "th"
	// Check the last two digits to handle 111, 112, 113, etc.
	lastTwo := n % 100
	if lastTwo >= 11 && lastTwo <= 13 {
		return "th"
	}

	// Otherwise, check the last digit
	switch n % 10 {
	case 1:
		return "st"
	case 2:
		return "nd"
	case 3:
		return "rd"
	default:
		return "th"
	}
}

// OrdinalWord converts an integer to its ordinal word representation.
//
// Examples:
//   - OrdinalWord(1) returns "first"
//   - OrdinalWord(2) returns "second"
//   - OrdinalWord(11) returns "eleventh"
//   - OrdinalWord(21) returns "twenty-first"
//   - OrdinalWord(100) returns "one hundredth"
//   - OrdinalWord(101) returns "one hundred first"
//   - OrdinalWord(-1) returns "negative first"
func OrdinalWord(n int) string {
	if n == 0 {
		return "zeroth"
	}

	if n < 0 {
		return "negative " + OrdinalWord(-n)
	}

	return convertToOrdinalWord(n)
}

// convertToOrdinalWord converts a positive integer to its ordinal word form.
func convertToOrdinalWord(n int) string {
	// Handle numbers 1-19 with direct lookup
	if n <= 19 {
		return onesOrdinal[n]
	}

	// Handle exact tens (20, 30, 40, ...)
	if n < 100 && n%10 == 0 {
		return tensOrdinal[n/10]
	}

	// Handle 20-99 with compound form (twenty-first, etc.)
	if n < 100 {
		return tensCardinal[n/10] + "-" + onesOrdinal[n%10]
	}

	// Handle exact hundreds (100, 200, ...)
	if n < 1000 && n%100 == 0 {
		return onesCardinal[n/100] + " hundredth"
	}

	// Handle 100-999
	if n < 1000 {
		return onesCardinal[n/100] + " hundred " + convertToOrdinalWord(n%100)
	}

	// Handle exact thousands (1000, 2000, ...)
	if n < 1000000 && n%1000 == 0 {
		return cardinalWord(n/1000) + " thousandth"
	}

	// Handle 1000-999999
	if n < 1000000 {
		return cardinalWord(n/1000) + " thousand " + convertToOrdinalWord(n%1000)
	}

	// Handle exact millions (1000000, 2000000, ...)
	if n < 1000000000 && n%1000000 == 0 {
		return cardinalWord(n/1000000) + " millionth"
	}

	// Handle 1000000-999999999
	if n < 1000000000 {
		return cardinalWord(n/1000000) + " million " + convertToOrdinalWord(n%1000000)
	}

	// Handle exact billions
	if n%1000000000 == 0 {
		return cardinalWord(n/1000000000) + " billionth"
	}

	// Handle billions and above
	return cardinalWord(n/1000000000) + " billion " + convertToOrdinalWord(n%1000000000)
}

// NumberToWords converts an integer to its English word representation.
//
// Examples:
//   - NumberToWords(0) returns "zero"
//   - NumberToWords(1) returns "one"
//   - NumberToWords(42) returns "forty-two"
//   - NumberToWords(100) returns "one hundred"
//   - NumberToWords(1000) returns "one thousand"
//   - NumberToWords(-5) returns "negative five"
func NumberToWords(n int) string {
	if n < 0 {
		return "negative " + cardinalWord(-n)
	}
	return cardinalWord(n)
}

// cardinalWord converts a positive integer to its cardinal word form.
func cardinalWord(n int) string {
	if n == 0 {
		return "zero"
	}

	if n <= 19 {
		return onesCardinal[n]
	}

	if n < 100 && n%10 == 0 {
		return tensCardinal[n/10]
	}

	if n < 100 {
		return tensCardinal[n/10] + "-" + onesCardinal[n%10]
	}

	if n < 1000 && n%100 == 0 {
		return onesCardinal[n/100] + " hundred"
	}

	if n < 1000 {
		return onesCardinal[n/100] + " hundred " + cardinalWord(n%100)
	}

	if n < 1000000 && n%1000 == 0 {
		return cardinalWord(n/1000) + " thousand"
	}

	if n < 1000000 {
		return cardinalWord(n/1000) + " thousand " + cardinalWord(n%1000)
	}

	if n < 1000000000 && n%1000000 == 0 {
		return cardinalWord(n/1000000) + " million"
	}

	if n < 1000000000 {
		return cardinalWord(n/1000000) + " million " + cardinalWord(n%1000000)
	}

	if n%1000000000 == 0 {
		return cardinalWord(n/1000000000) + " billion"
	}

	return cardinalWord(n/1000000000) + " billion " + cardinalWord(n%1000000000)
}

// onesCardinal maps 1-19 to their cardinal word forms.
var onesCardinal = []string{
	"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen",
	"seventeen", "eighteen", "nineteen",
}

// onesOrdinal maps 1-19 to their ordinal word forms.
var onesOrdinal = []string{
	"", "first", "second", "third", "fourth", "fifth", "sixth", "seventh",
	"eighth", "ninth", "tenth", "eleventh", "twelfth", "thirteenth",
	"fourteenth", "fifteenth", "sixteenth", "seventeenth", "eighteenth",
	"nineteenth",
}

// tensCardinal maps tens (2-9 representing 20-90) to their cardinal word forms.
var tensCardinal = []string{
	"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety",
}

// tensOrdinal maps tens (2-9 representing 20-90) to their ordinal word forms.
var tensOrdinal = []string{
	"", "", "twentieth", "thirtieth", "fortieth", "fiftieth", "sixtieth",
	"seventieth", "eightieth", "ninetieth",
}

// PresentParticiple converts a verb to its present participle (-ing) form.
//
// Examples:
//   - PresentParticiple("run") returns "running" (double consonant)
//   - PresentParticiple("make") returns "making" (drop silent e)
//   - PresentParticiple("play") returns "playing" (just add -ing)
//   - PresentParticiple("die") returns "dying" (ie -> ying)
//   - PresentParticiple("see") returns "seeing" (ee -> eeing)
//   - PresentParticiple("panic") returns "panicking" (c -> ck)
func PresentParticiple(verb string) string {
	if verb == "" {
		return ""
	}

	lower := strings.ToLower(verb)
	n := len(lower)

	// Already a present participle (ends in doubled consonant + ing, like "running")
	if isAlreadyParticiple(lower) {
		return verb
	}

	// Single letter verbs - just add -ing
	if n == 1 {
		return verb + matchSuffix(verb, "ing")
	}

	// Words ending in -ie: change to -ying (die -> dying, lie -> lying)
	if strings.HasSuffix(lower, "ie") {
		return verb[:len(verb)-2] + matchSuffix(verb, "ying")
	}

	// Words ending in -ee: just add -ing (see -> seeing, flee -> fleeing)
	if strings.HasSuffix(lower, "ee") {
		return verb + matchSuffix(verb, "ing")
	}

	// Words ending in -ye, -oe: just add -ing (dye -> dyeing, hoe -> hoeing)
	if strings.HasSuffix(lower, "ye") || strings.HasSuffix(lower, "oe") {
		return verb + matchSuffix(verb, "ing")
	}

	// Words ending in -nge/-inge: keep the e (singe -> singeing)
	if strings.HasSuffix(lower, "nge") || strings.HasSuffix(lower, "inge") {
		return verb + matchSuffix(verb, "ing")
	}

	// Words ending in -c: add k before -ing (panic -> panicking)
	if strings.HasSuffix(lower, "c") {
		return verb + matchSuffix(verb, "king")
	}

	// Words ending in consonant + e (silent e): drop e, add -ing
	// But keep e if it's the only vowel in the word (be -> being)
	if strings.HasSuffix(lower, "e") && n >= 2 {
		beforeE := rune(lower[n-2])
		if !isVowel(beforeE) {
			// Check if 'e' is the only vowel (not a silent e)
			vowelCount := countVowels(lower[:n-1]) // count vowels excluding final 'e'
			if vowelCount == 0 {
				// 'e' is the only vowel, keep it (be -> being)
				return verb + matchSuffix(verb, "ing")
			}
			return verb[:len(verb)-1] + matchSuffix(verb, "ing")
		}
		// vowel + e: just add -ing
		return verb + matchSuffix(verb, "ing")
	}

	// Check for CVC pattern that requires doubling the final consonant
	if shouldDoubleConsonant(lower) {
		lastChar := verb[len(verb)-1:]
		return verb + matchSuffix(verb, strings.ToLower(lastChar)+"ing")
	}

	// Default: just add -ing
	return verb + matchSuffix(verb, "ing")
}

// isAlreadyParticiple checks if a word is already a present participle.
// This catches words like "running", "sitting" but not base verbs like "sing".
func isAlreadyParticiple(lower string) bool {
	if !strings.HasSuffix(lower, "ing") {
		return false
	}
	n := len(lower)
	if n < 5 {
		return false
	}

	// Check for doubled consonant before -ing (running, sitting, hitting)
	beforeIng := lower[n-4]
	beforeThat := lower[n-5]
	if beforeIng == beforeThat && !isVowel(rune(beforeIng)) {
		return true
	}

	// Check for common participle patterns
	// Words ending in -ting, -ning, -ping, etc. after a consonant
	// But not "sing", "ring", "bring" which are base verbs

	return false
}

// shouldDoubleConsonant checks if the final consonant should be doubled.
// This applies to CVC (consonant-vowel-consonant) patterns in stressed syllables.
func shouldDoubleConsonant(lower string) bool {
	n := len(lower)
	if n < 3 {
		return false
	}

	lastChar := rune(lower[n-1])

	// Don't double w, x, y
	if lastChar == 'w' || lastChar == 'x' || lastChar == 'y' {
		return false
	}

	// Must end in a consonant
	if isVowel(lastChar) {
		return false
	}

	// Check for single vowel before the final consonant
	beforeLast := rune(lower[n-2])
	if !isVowel(beforeLast) {
		return false
	}

	// Don't double if there's a vowel digraph (two vowels in a row before consonant)
	// Examples: eat, read, beat, lead - these have "ea" before the final consonant
	if n >= 3 && isVowel(rune(lower[n-3])) {
		return false
	}

	// At this point we know the word ends in consonant + single vowel + consonant

	// For short words (3 letters): double if CVC pattern
	// Examples: run, sit, hit, cut
	if n == 3 {
		return true
	}

	// For 4-letter words: double only if there's a single vowel cluster
	// "stop", "drop", "skip", "plan" -> double (single vowel)
	// "open" -> don't double (two separate vowels = multi-syllable)
	if n == 4 {
		// Count distinct vowel clusters
		if countVowels(lower) == 1 {
			return true
		}
		return false
	}

	// For multi-syllable words, check for common patterns that double
	// Words ending in stressed syllables typically double
	return doubleConsonantWords[lower]
}

// countVowels counts the number of vowels in a string.
func countVowels(s string) int {
	count := 0
	for _, r := range s {
		if isVowel(r) {
			count++
		}
	}
	return count
}

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

// JoinWithSep combines a slice of strings into a grammatically correct English list
// with a custom conjunction and separator.
//
// This is useful when list items themselves contain commas.
//
// Examples:
//   - JoinWithSep([]string{"a", "b", "c"}, "and", "; ") returns "a; b; and c"
//   - JoinWithSep([]string{"Jan 1, 2020", "Feb 2, 2021"}, "and", "; ") returns "Jan 1, 2020; and Feb 2, 2021"
func JoinWithSep(words []string, conj string, sep string) string {
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

// Compare compares two words for singular/plural equality.
//
// It returns:
//   - "eq" if the words are equal (case-insensitive)
//   - "s:p" if word1 is singular and word2 is its plural form
//   - "p:s" if word1 is plural and word2 is its singular form
//   - "p:p" if both words are different plural forms of the same word
//   - "" if the words are not related
//
// Examples:
//   - Compare("cat", "cat") returns "eq"
//   - Compare("cat", "cats") returns "s:p"
//   - Compare("cats", "cat") returns "p:s"
//   - Compare("indexes", "indices") returns "p:p"
//   - Compare("cat", "dog") returns ""
func Compare(word1, word2 string) string {
	// Handle empty strings
	if word1 == "" || word2 == "" {
		if word1 == "" && word2 == "" {
			return "eq"
		}
		return ""
	}

	// Normalize for comparison
	lower1 := strings.ToLower(word1)
	lower2 := strings.ToLower(word2)

	// Same word
	if lower1 == lower2 {
		return "eq"
	}

	// Check if word1 is singular and word2 is its plural
	// Use Plural() for verification since Singular() has edge cases
	if strings.ToLower(Plural(word1)) == lower2 {
		return "s:p"
	}

	// Check if word2 is singular and word1 is its plural
	if strings.ToLower(Plural(word2)) == lower1 {
		return "p:s"
	}

	// Check if both are different plural forms of the same singular word
	singular1 := strings.ToLower(Singular(word1))
	singular2 := strings.ToLower(Singular(word2))

	if singular1 == singular2 {
		// Verify both are actually plurals (different from their singular form)
		// by checking that pluralizing the singular gives us something related
		pluralOfSingular := strings.ToLower(Plural(singular1))
		// If both words singularize to the same thing, and that singular
		// can be pluralized, they're both plural forms
		if lower1 != singular1 && lower2 != singular2 {
			// Additional verification: ensure the singular is valid
			if pluralOfSingular == lower1 || pluralOfSingular == lower2 {
				return "p:p"
			}
		}
	}

	return ""
}

// CompareNouns compares two nouns for singular/plural equality.
//
// This is an alias for Compare that makes the intent explicit when working
// specifically with nouns.
//
// It returns:
//   - "eq" if the nouns are equal (case-insensitive)
//   - "s:p" if noun1 is singular and noun2 is its plural form
//   - "p:s" if noun1 is plural and noun2 is its singular form
//   - "p:p" if both nouns are different plural forms of the same word
//   - "" if the nouns are not related
//
// Examples:
//   - CompareNouns("cat", "cats") returns "s:p"
//   - CompareNouns("mice", "mouse") returns "p:s"
func CompareNouns(noun1, noun2 string) string {
	return Compare(noun1, noun2)
}

// CompareVerbs compares two verbs for conjugation equality.
//
// NOTE: This is a placeholder stub for future implementation.
// Verb conjugation comparison is not yet implemented.
//
// Currently always returns an empty string.
//
// Future implementation will return:
//   - "eq" if the verbs are equal (case-insensitive)
//   - Comparison codes for different conjugation relationships
//   - "" if the verbs are not related
func CompareVerbs(verb1, verb2 string) string {
	// TODO: Implement verb conjugation comparison
	return ""
}

// CompareAdjs compares two adjectives for comparative/superlative equality.
//
// NOTE: This is a placeholder stub for future implementation.
// Adjective comparison is not yet implemented.
//
// Currently always returns an empty string.
//
// Future implementation will return:
//   - "eq" if the adjectives are equal (case-insensitive)
//   - Comparison codes for different adjective relationships (e.g., base/comparative/superlative)
//   - "" if the adjectives are not related
func CompareAdjs(adj1, adj2 string) string {
	// TODO: Implement adjective comparison
	return ""
}

// Num stores and retrieves a default count for number-related operations.
//
// When called with a positive integer, it stores that value as the default
// count and returns it. When called with 0 or no arguments, it clears the
// default count and returns 0.
//
// Examples:
//   - Num(5) stores 5 as default count, returns 5
//   - Num(0) clears the default count, returns 0
//   - Num() clears the default count, returns 0
func Num(n ...int) int {
	if len(n) == 0 || n[0] == 0 {
		defaultNum = 0
		return 0
	}
	defaultNum = n[0]
	return defaultNum
}

// GetNum retrieves the current default count.
//
// Returns 0 if no default has been set or if it was cleared.
//
// Examples:
//   - After Num(5): GetNum() returns 5
//   - After Num(0) or Num(): GetNum() returns 0
//   - Before any Num() call: GetNum() returns 0
func GetNum() int {
	return defaultNum
}

// No returns a count and noun phrase in English, using "no" for zero counts.
//
// The function handles pluralization automatically:
//   - For count 0 with classicalZero=false (default): returns "no" + plural form
//   - For count 0 with classicalZero=true: returns "no" + singular form
//   - For count 1: returns "1" + singular form
//   - For count > 1: returns count + plural form
//
// Examples:
//   - No("error", 0) returns "no errors" (default)
//   - No("error", 1) returns "1 error"
//   - No("error", 2) returns "2 errors"
//   - No("child", 0) returns "no children" (default)
//   - No("child", 1) returns "1 child"
//   - No("child", 3) returns "3 children"
//
// With ClassicalZero(true):
//   - No("error", 0) returns "no error"
//   - No("child", 0) returns "no child"
func No(word string, count int) string {
	if count == 0 {
		if classicalZero {
			return "no " + word
		}
		return "no " + Plural(word)
	}
	if count == 1 || count == -1 {
		return fmt.Sprintf("%d %s", count, word)
	}
	return fmt.Sprintf("%d %s", count, Plural(word))
}
