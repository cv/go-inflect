package inflect

// gender stores the current gender for singular third-person pronouns.
// Valid values: 'm' (masculine), 'f' (feminine), 'n' (neuter), 't' (they/singular they).
// Default is 't' (singular they).
var gender = "t"

// Gender sets the gender for singular third-person pronouns.
//
// The gender affects pronoun selection in Singular():
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
	switch g {
	case "m", "f", "n", "t":
		gender = g
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
	return gender
}
