package inflect

// Gender sets the gender for singular third-person pronouns.
//
// The gender affects pronoun selection in SingularNoun():
//   - Gender("m") - masculine: they -> he, them -> him, their -> his
//   - Gender("f") - feminine: they -> she, them -> her, their -> hers
//   - Gender("n") - neuter: they -> it, them -> it, their -> its
//   - Gender("t") - they (singular they): they -> they, them -> them, their -> their
//
// The default gender is "t" (singular they).
//
// Invalid gender values are ignored; the gender remains unchanged.
//
// Examples:
//
//	Gender("m")
//	GetGender() // returns "m"
//	Gender("f")
//	GetGender() // returns "f"
//	Gender("invalid")
//	GetGender() // returns "f" (unchanged)
func Gender(g string) {
	defaultEngine.SetGender(g)
}

// SetGender sets the gender for singular third-person pronouns.
//
// The gender affects pronoun selection in SingularNoun():
//   - SetGender("m") - masculine: they -> he, them -> him, their -> his
//   - SetGender("f") - feminine: they -> she, them -> her, their -> hers
//   - SetGender("n") - neuter: they -> it, them -> it, their -> its
//   - SetGender("t") - they (singular they): they -> they, them -> them, their -> their
//
// The default gender is "t" (singular they).
//
// Invalid gender values are ignored; the gender remains unchanged.
//
// Examples:
//
//	e.SetGender("m")
//	e.GetGender() // returns "m"
//	e.SetGender("f")
//	e.GetGender() // returns "f"
//	e.SetGender("invalid")
//	e.GetGender() // returns "f" (unchanged)
func (e *Engine) SetGender(g string) {
	switch g {
	case "m", "f", "n", "t":
		e.mu.Lock()
		e.gender = g
		e.mu.Unlock()
	}
}

// GetGender returns the current gender setting for singular third-person pronouns.
//
// Returns one of:
//   - "m" - masculine
//   - "f" - feminine
//   - "n" - neuter
//   - "t" - they (singular they, the default)
//
// Examples:
//
//	GetGender() // returns "t" (default)
//	Gender("m")
//	GetGender() // returns "m"
func GetGender() string {
	return defaultEngine.GetGender()
}

// GetGender returns the current gender setting for singular third-person pronouns.
//
// Returns one of:
//   - "m" - masculine
//   - "f" - feminine
//   - "n" - neuter
//   - "t" - they (singular they, the default)
//
// Examples:
//
//	e.GetGender() // returns "t" (default)
//	e.SetGender("m")
//	e.GetGender() // returns "m"
func (e *Engine) GetGender() string {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return e.gender
}
