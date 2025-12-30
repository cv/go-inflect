package inflect

// FutureTense returns the future tense form of an English verb using "will".
//
// Examples:
//   - FutureTense("walk") returns "will walk"
//   - FutureTense("go") returns "will go"
//   - FutureTense("be") returns "will be"
//   - FutureTense("WALK") returns "WILL WALK"
//   - FutureTense("Walk") returns "Will Walk"
func FutureTense(verb string) string {
	if verb == "" {
		return ""
	}

	will := matchCase(verb, "will")
	return will + " " + verb
}
