package inflect

import (
	"maps"
	"regexp"
	"sync"
)

// Engine holds all mutable state for inflection operations.
// Use NewEngine() to create an instance with default settings.
// The Engine is safe for concurrent use; all methods are protected by a mutex.
//
// # Thread Safety Architecture
//
// All mutable state is encapsulated within the Engine struct and protected
// by a sync.RWMutex. Package-level functions delegate to a defaultEngine
// instance, providing backward-compatible API while ensuring thread safety.
//
// # Mutable State (in Engine)
//
// The following state is mutable and protected by Engine.mu:
//   - Classical mode flags: classicalMode, classicalAll, classicalZero,
//     classicalHerd, classicalNames, classicalAncient, classicalPersons
//   - Custom noun mappings: irregularPlurals, singularIrregulars
//   - Custom verb mappings: customVerbs, customVerbsReverse
//   - Custom adjective mappings: customAdjs, customAdjsReverse
//   - Custom article patterns: customAWords, customAnWords, customAPatterns, customAnPatterns
//   - Gender setting: gender (for third-person pronoun singularization)
//   - Possessive style: possessiveStyle (modern vs traditional)
//   - Default number: defaultNum (for Num/GetNum)
//
// # Immutable State (package-level variables)
//
// The following package-level variables are IMMUTABLE after initialization
// and are safe for concurrent read access without locking:
//
// Lookup tables (never modified after init):
//   - adjective.go: irregularComparatives, irregularSuperlatives, twoSyllableWithSuffix
//   - adverb.go: irregularAdverbs, unchangedAdverbs
//   - article.go: silentHWords, lowercaseAbbrevs
//   - currency.go: currencies
//   - number.go: onesCardinal, onesOrdinal, tensCardinal, tensOrdinal
//   - ordinal.go: cardinalToOrdinal, ordinalToCardinalMap, ordinalWords
//   - participle.go: doubleConsonantWords, irregularPastParticiples, knownParticiples
//   - past_tense.go: irregularPastTense
//   - plural.go: changeToVesWords, oExceptionWords, unchangedPlurals, herdAnimals,
//     classicalLatinPlurals, defaultIrregularPlurals
//   - possessive.go: irregularPluralNoS, singularEndsInS, commonNouns, truncatedNames, validShortA
//   - pronouns.go: pronounNominativePlural, pronounAccusativePlural, pronounPossessivePlural,
//     pronounReflexivePlural, allPronounsToPlural, pronoun*SingularByGender maps
//   - singular.go: feWordBases
//   - verbs.go: verbSingularToPlural, verbPluralToSingular, verbUnchanged,
//     adjSingularToPlural, adjPluralToSingular, adjPluralToSingularByGender
//
// Compiled regular expressions (immutable after compilation):
//   - inflect_funcs.go: inflectFuncPattern
//   - rails.go: notURLSafe, multiSep
//
// Function lookup tables (immutable after init):
//   - inflect_funcs.go: inflectFuncs
//
// # Default Engine
//
// The package-level defaultEngine (in classical.go) is created at package
// initialization and used by all package-level functions. It is safe for
// concurrent use but modifications affect all callers globally.
// For isolated configurations, use NewEngine() to create separate instances.
type Engine struct {
	mu sync.RWMutex

	// Classical mode settings
	classicalMode    bool
	classicalAll     bool
	classicalZero    bool
	classicalHerd    bool
	classicalNames   bool
	classicalAncient bool
	classicalPersons bool

	// Noun mappings
	irregularPlurals   map[string]string
	singularIrregulars map[string]string

	// Custom verb definitions
	customVerbs        map[string]string
	customVerbsReverse map[string]string

	// Custom adjective definitions
	customAdjs        map[string]string
	customAdjsReverse map[string]string

	// Article patterns
	customAWords     map[string]bool
	customAnWords    map[string]bool
	customAPatterns  []*regexp.Regexp
	customAnPatterns []*regexp.Regexp

	// Gender for singular third-person pronouns
	// Valid values: "m" (masculine), "f" (feminine), "n" (neuter), "t" (they/singular they)
	gender string

	// Possessive style: PossessiveModern or PossessiveTraditional
	possessiveStyle PossessiveStyleType

	// Default number for Num/GetNum
	defaultNum int
}

// NewEngine creates a new Engine instance with default settings.
//
// Default settings:
//   - All classical flags are false
//   - irregularPlurals is initialized from defaultIrregularPlurals
//   - singularIrregulars is built as the reverse of irregularPlurals
//   - All custom maps are empty
//   - Gender is "t" (singular they)
//   - Possessive style is PossessiveModern
//   - Default number is 0
//
// Example:
//
//	e := NewEngine()
//	e.Plural("cat") // returns "cats"
func NewEngine() *Engine {
	irregulars := copyMap(defaultIrregularPlurals)

	// Build singularIrregulars as reverse of irregularPlurals
	singulars := make(map[string]string, len(irregulars))
	for singular, plural := range irregulars {
		singulars[plural] = singular
	}

	return &Engine{
		// Classical mode settings - all false by default
		classicalMode:    false,
		classicalAll:     false,
		classicalZero:    false,
		classicalHerd:    false,
		classicalNames:   false,
		classicalAncient: false,
		classicalPersons: false,

		// Noun mappings
		irregularPlurals:   irregulars,
		singularIrregulars: singulars,

		// Custom verb definitions - empty by default
		customVerbs:        make(map[string]string),
		customVerbsReverse: make(map[string]string),

		// Custom adjective definitions - empty by default
		customAdjs:        make(map[string]string),
		customAdjsReverse: make(map[string]string),

		// Article patterns - empty by default
		customAWords:     make(map[string]bool),
		customAnWords:    make(map[string]bool),
		customAPatterns:  nil,
		customAnPatterns: nil,

		// Gender - default to singular they
		gender: "t",

		// Possessive style - default to modern
		possessiveStyle: PossessiveModern,

		// Default number - 0 means not set
		defaultNum: 0,
	}
}

// Clone creates a deep copy of the Engine.
// The returned Engine is independent of the original - modifications to one
// will not affect the other.
//
// Example:
//
//	e1 := NewEngine()
//	e1.DefNoun("foo", "foos")
//	e2 := e1.Clone()
//	e2.DefNoun("bar", "bars")
//	// e1 has "foo" -> "foos" but not "bar" -> "bars"
//	// e2 has both mappings
func (e *Engine) Clone() *Engine {
	e.mu.RLock()
	defer e.mu.RUnlock()

	// Deep copy all maps using maps.Copy
	irregulars := make(map[string]string, len(e.irregularPlurals))
	maps.Copy(irregulars, e.irregularPlurals)

	singulars := make(map[string]string, len(e.singularIrregulars))
	maps.Copy(singulars, e.singularIrregulars)

	verbs := make(map[string]string, len(e.customVerbs))
	maps.Copy(verbs, e.customVerbs)

	verbsReverse := make(map[string]string, len(e.customVerbsReverse))
	maps.Copy(verbsReverse, e.customVerbsReverse)

	adjs := make(map[string]string, len(e.customAdjs))
	maps.Copy(adjs, e.customAdjs)

	adjsReverse := make(map[string]string, len(e.customAdjsReverse))
	maps.Copy(adjsReverse, e.customAdjsReverse)

	aWords := make(map[string]bool, len(e.customAWords))
	maps.Copy(aWords, e.customAWords)

	anWords := make(map[string]bool, len(e.customAnWords))
	maps.Copy(anWords, e.customAnWords)

	// Copy regex slices (regexes are immutable, so shallow copy is safe)
	var aPatterns []*regexp.Regexp
	if e.customAPatterns != nil {
		aPatterns = make([]*regexp.Regexp, len(e.customAPatterns))
		copy(aPatterns, e.customAPatterns)
	}

	var anPatterns []*regexp.Regexp
	if e.customAnPatterns != nil {
		anPatterns = make([]*regexp.Regexp, len(e.customAnPatterns))
		copy(anPatterns, e.customAnPatterns)
	}

	return &Engine{
		classicalMode:      e.classicalMode,
		classicalAll:       e.classicalAll,
		classicalZero:      e.classicalZero,
		classicalHerd:      e.classicalHerd,
		classicalNames:     e.classicalNames,
		classicalAncient:   e.classicalAncient,
		classicalPersons:   e.classicalPersons,
		irregularPlurals:   irregulars,
		singularIrregulars: singulars,
		customVerbs:        verbs,
		customVerbsReverse: verbsReverse,
		customAdjs:         adjs,
		customAdjsReverse:  adjsReverse,
		customAWords:       aWords,
		customAnWords:      anWords,
		customAPatterns:    aPatterns,
		customAnPatterns:   anPatterns,
		gender:             e.gender,
		possessiveStyle:    e.possessiveStyle,
		defaultNum:         e.defaultNum,
	}
}

// Reset restores the Engine to its default state.
// This clears all custom definitions and resets all options to defaults.
//
// All state is reset to match the values from NewEngine():
//   - All classical flags are set to false
//   - irregularPlurals is restored from defaultIrregularPlurals
//   - singularIrregulars is rebuilt as the reverse of irregularPlurals
//   - All custom maps (verbs, adjectives, article patterns) are cleared
//   - Gender is reset to "t" (singular they)
//   - Possessive style is reset to PossessiveModern
//   - Default number is reset to 0
//
// Example:
//
//	e := NewEngine()
//	e.Classical(true)
//	e.DefNoun("foo", "foos")
//	e.Reset()
//	e.IsClassical() // returns false
//	e.Plural("foo") // returns "foos" (standard rule, custom removed)
func (e *Engine) Reset() {
	e.mu.Lock()
	defer e.mu.Unlock()

	// Reset classical options
	e.classicalMode = false
	e.classicalAll = false
	e.classicalZero = false
	e.classicalHerd = false
	e.classicalNames = false
	e.classicalAncient = false
	e.classicalPersons = false

	// Reset noun mappings
	e.irregularPlurals = copyMap(defaultIrregularPlurals)
	// Build singularIrregulars as reverse of irregularPlurals
	e.singularIrregulars = make(map[string]string, len(e.irregularPlurals))
	for singular, plural := range e.irregularPlurals {
		e.singularIrregulars[plural] = singular
	}

	// Reset custom definitions
	e.customVerbs = make(map[string]string)
	e.customVerbsReverse = make(map[string]string)
	e.customAdjs = make(map[string]string)
	e.customAdjsReverse = make(map[string]string)

	// Reset article patterns
	e.customAWords = make(map[string]bool)
	e.customAnWords = make(map[string]bool)
	e.customAPatterns = nil
	e.customAnPatterns = nil

	// Reset gender
	e.gender = "t"

	// Reset other state
	e.defaultNum = 0
	e.possessiveStyle = PossessiveModern
}
