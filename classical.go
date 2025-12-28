package inflect

// defaultEngine is the package-level Engine instance used by all package-level functions.
// This provides backward compatibility while allowing thread-safe concurrent access.
// All package-level functions are thin wrappers that delegate to this Engine.
var defaultEngine = NewEngine()

// DefaultEngine returns the default package-level Engine.
// This can be used to access Engine methods directly or to configure
// the default behavior.
//
// The returned Engine is safe for concurrent use; all methods are protected
// by an internal mutex. However, callers should be aware that modifications
// to the default Engine affect all package-level function calls globally.
//
// For isolated configurations, use NewEngine() to create a separate instance.
//
// Examples:
//
//	// Access the default engine directly
//	e := inflect.DefaultEngine()
//	e.Classical(true)
//	e.Plural("formula") // returns "formulae"
//
//	// All package-level calls use the same engine
//	inflect.Plural("formula") // also returns "formulae"
func DefaultEngine() *Engine {
	return defaultEngine
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
	defaultEngine.ClassicalAll(enabled)
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
//	e := NewEngine()
//	e.ClassicalAll(true)
//	e.Plural("formula") // returns "formulae"
//	e.ClassicalAll(false)
//	e.Plural("formula") // returns "formulas"
func (e *Engine) ClassicalAll(enabled bool) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.classicalAll = enabled
	e.classicalZero = enabled
	e.classicalHerd = enabled
	e.classicalNames = enabled
	e.classicalAncient = enabled
	e.classicalPersons = enabled
	// Also set the legacy classicalMode for backward compatibility
	e.classicalMode = enabled
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
	defaultEngine.Classical(enabled)
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
//	e := NewEngine()
//	e.Classical(true)
//	e.Plural("formula") // returns "formulae"
//	e.Classical(false)
//	e.Plural("formula") // returns "formulas"
func (e *Engine) Classical(enabled bool) {
	e.ClassicalAll(enabled)
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
	return defaultEngine.IsClassicalAll()
}

// IsClassicalAll returns whether all classical pluralization options are enabled.
//
// Returns true only if all classical options are enabled, false otherwise.
//
// Examples:
//
//	e := NewEngine()
//	e.IsClassicalAll() // returns false (default)
//	e.ClassicalAll(true)
//	e.IsClassicalAll() // returns true
//	e.ClassicalAll(false)
//	e.IsClassicalAll() // returns false
func (e *Engine) IsClassicalAll() bool {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return e.classicalAll && e.classicalZero && e.classicalHerd &&
		e.classicalNames && e.classicalAncient && e.classicalPersons
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
	return defaultEngine.IsClassical()
}

// IsClassical returns whether classical pluralization mode is enabled.
//
// Returns true if Classical(true) or ClassicalAll(true) was called, false otherwise.
// This checks the classicalAncient flag which controls Latin/Greek plurals.
//
// Examples:
//
//	e := NewEngine()
//	e.IsClassical() // returns false (default)
//	e.Classical(true)
//	e.IsClassical() // returns true
//	e.Classical(false)
//	e.IsClassical() // returns false
func (e *Engine) IsClassical() bool {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return e.classicalAncient || e.classicalMode
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
	defaultEngine.ClassicalAncient(enabled)
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
//	e := NewEngine()
//	e.ClassicalAncient(true)
//	e.Plural("formula") // returns "formulae"
//	e.ClassicalAncient(false)
//	e.Plural("formula") // returns "formulas"
func (e *Engine) ClassicalAncient(enabled bool) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.classicalAncient = enabled
	// Also set the legacy classicalMode to keep them in sync
	e.classicalMode = enabled
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
	return defaultEngine.IsClassicalAncient()
}

// IsClassicalAncient returns whether classical Latin/Greek plural forms are enabled.
//
// Returns true if ClassicalAncient(true), Classical(true), or ClassicalAll(true)
// was called, false otherwise.
//
// Examples:
//
//	e := NewEngine()
//	e.IsClassicalAncient() // returns false (default)
//	e.ClassicalAncient(true)
//	e.IsClassicalAncient() // returns true
//	e.ClassicalAncient(false)
//	e.IsClassicalAncient() // returns false
func (e *Engine) IsClassicalAncient() bool {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return e.classicalAncient
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
	defaultEngine.ClassicalZero(enabled)
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
//	e := NewEngine()
//	e.ClassicalZero(true)
//	e.No("cat", 0) // returns "no cat"
//	e.ClassicalZero(false)
//	e.No("cat", 0) // returns "no cats"
func (e *Engine) ClassicalZero(enabled bool) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.classicalZero = enabled
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
	return defaultEngine.IsClassicalZero()
}

// IsClassicalZero returns whether classical zero count pluralization is enabled.
//
// Returns true if ClassicalZero(true) or ClassicalAll(true) was called,
// false otherwise.
//
// Examples:
//
//	e := NewEngine()
//	e.IsClassicalZero() // returns false (default)
//	e.ClassicalZero(true)
//	e.IsClassicalZero() // returns true
//	e.ClassicalZero(false)
//	e.IsClassicalZero() // returns false
func (e *Engine) IsClassicalZero() bool {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return e.classicalZero
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
	defaultEngine.ClassicalPersons(enabled)
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
//	e := NewEngine()
//	e.ClassicalPersons(true)
//	e.Plural("person") // returns "persons"
//	e.ClassicalPersons(false)
//	e.Plural("person") // returns "people"
func (e *Engine) ClassicalPersons(enabled bool) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.classicalPersons = enabled
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
	return defaultEngine.IsClassicalPersons()
}

// IsClassicalPersons returns whether classical person/persons pluralization is enabled.
//
// Returns true if ClassicalPersons(true) or ClassicalAll(true) was called,
// false otherwise.
//
// Examples:
//
//	e := NewEngine()
//	e.IsClassicalPersons() // returns false (default)
//	e.ClassicalPersons(true)
//	e.IsClassicalPersons() // returns true
//	e.ClassicalPersons(false)
//	e.IsClassicalPersons() // returns false
func (e *Engine) IsClassicalPersons() bool {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return e.classicalPersons
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
	defaultEngine.ClassicalHerd(enabled)
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
//	e := NewEngine()
//	e.ClassicalHerd(true)
//	e.Plural("wildebeest") // returns "wildebeest"
//	e.ClassicalHerd(false)
//	e.Plural("wildebeest") // returns "wildebeests"
func (e *Engine) ClassicalHerd(enabled bool) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.classicalHerd = enabled
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
	return defaultEngine.IsClassicalHerd()
}

// IsClassicalHerd returns whether classical herd animal pluralization is enabled.
//
// Returns true if ClassicalHerd(true) or ClassicalAll(true) was called,
// false otherwise.
//
// Examples:
//
//	e := NewEngine()
//	e.IsClassicalHerd() // returns false (default)
//	e.ClassicalHerd(true)
//	e.IsClassicalHerd() // returns true
//	e.ClassicalHerd(false)
//	e.IsClassicalHerd() // returns false
func (e *Engine) IsClassicalHerd() bool {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return e.classicalHerd
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
	defaultEngine.ClassicalNames(enabled)
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
//	e := NewEngine()
//	e.ClassicalNames(true)
//	e.Plural("Jones") // returns "Jones"
//	e.ClassicalNames(false)
//	e.Plural("Jones") // returns "Joneses"
func (e *Engine) ClassicalNames(enabled bool) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.classicalNames = enabled
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
	return defaultEngine.IsClassicalNames()
}

// IsClassicalNames returns whether classical proper name pluralization is enabled.
//
// Returns true if ClassicalNames(true) or ClassicalAll(true) was called,
// false otherwise.
//
// Examples:
//
//	e := NewEngine()
//	e.IsClassicalNames() // returns false (default)
//	e.ClassicalNames(true)
//	e.IsClassicalNames() // returns true
//	e.ClassicalNames(false)
//	e.IsClassicalNames() // returns false
func (e *Engine) IsClassicalNames() bool {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return e.classicalNames
}
