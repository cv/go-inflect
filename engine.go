package inflect

import (
	"maps"
	"regexp"
	"sync"
)

// Engine holds all mutable state for inflection operations.
// Use NewEngine() to create an instance with default settings.
// The Engine is safe for concurrent use; all methods are protected by a mutex.
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
