package inflect

// classicalMode controls whether to prefer classical Latin/Greek plurals.
// When true, words like "formula" become "formulae" instead of "formulas".
// Default is false (modern English plurals).
//
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
