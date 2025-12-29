package inflect

import "strings"

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

// irregularPastParticiples maps base verbs to their irregular past participle forms.
var irregularPastParticiples = map[string]string{
	"be": "been", "bear": "borne", "beat": "beaten", "become": "become", "begin": "begun",
	"bend": "bent", "bet": "bet", "bind": "bound", "bite": "bitten",
	"bleed": "bled", "blow": "blown", "break": "broken", "breed": "bred",
	"bring": "brought", "build": "built", "burn": "burnt", "burst": "burst",
	"buy": "bought", "catch": "caught", "choose": "chosen", "cling": "clung",
	"come": "come", "cost": "cost", "creep": "crept", "cut": "cut",
	"deal": "dealt", "dig": "dug", "do": "done", "draw": "drawn",
	"dream": "dreamt", "drink": "drunk", "drive": "driven", "eat": "eaten",
	"fall": "fallen", "feed": "fed", "feel": "felt", "fight": "fought",
	"find": "found", "flee": "fled", "fling": "flung", "fly": "flown",
	"forbid": "forbidden", "forget": "forgotten", "forgive": "forgiven",
	"freeze": "frozen", "get": "got", "give": "given", "go": "gone",
	"grind": "ground", "grow": "grown", "hang": "hung", "have": "had",
	"hear": "heard", "hide": "hidden", "hit": "hit", "hold": "held",
	"hurt": "hurt", "keep": "kept", "kneel": "knelt", "know": "known",
	"lay": "laid", "lead": "led", "lean": "leant", "leap": "leapt",
	"learn": "learnt", "leave": "left", "lend": "lent", "let": "let",
	"lie": "lain", "light": "lit", "lose": "lost", "make": "made",
	"mean": "meant", "meet": "met", "pay": "paid", "prove": "proven",
	"put": "put", "quit": "quit", "read": "read", "ride": "ridden",
	"ring": "rung", "rise": "risen", "run": "run", "say": "said",
	"see": "seen", "seek": "sought", "sell": "sold", "send": "sent",
	"set": "set", "shake": "shaken", "shed": "shed", "shine": "shone",
	"shoot": "shot", "show": "shown", "shrink": "shrunk", "shut": "shut",
	"sing": "sung", "sink": "sunk", "sit": "sat", "sleep": "slept",
	"slide": "slid", "sling": "slung", "slit": "slit", "smell": "smelt",
	"speak": "spoken", "speed": "sped", "spell": "spelt", "spend": "spent",
	"spill": "spilt", "spin": "spun", "spit": "spat", "split": "split",
	"spoil": "spoilt", "spread": "spread", "spring": "sprung", "stand": "stood",
	"steal": "stolen", "stick": "stuck", "sting": "stung", "stink": "stunk",
	"stride": "stridden", "strike": "struck", "string": "strung", "strive": "striven",
	"swear": "sworn", "sweep": "swept", "swim": "swum", "swing": "swung",
	"take": "taken", "teach": "taught", "tear": "torn", "tell": "told",
	"think": "thought", "throw": "thrown", "thrust": "thrust", "tread": "trodden",
	"understand": "understood", "wake": "woken", "wear": "worn", "weave": "woven",
	"weep": "wept", "win": "won", "wind": "wound", "withdraw": "withdrawn",
	"wring": "wrung", "write": "written",
	// Additional irregular past participles
	"beseech": "besought", "beget": "begotten", "dwell": "dwelt",
	"forsake": "forsaken", "inlay": "inlaid", "slay": "slain",
	"awake": "awoken", "arise": "arisen",
	// Compound verbs with irregular bases
	"foresee": "foreseen", "outdo": "outdone", "outgrow": "outgrown",
	"overdo": "overdone", "overhear": "overheard", "override": "overridden",
	"oversee": "overseen", "oversleep": "overslept", "overthrow": "overthrown",
	"partake": "partaken", "rebuild": "rebuilt", "redo": "redone",
	"remake": "remade", "repay": "repaid", "retell": "retold",
	"rewind": "rewound", "rewrite": "rewritten", "unbind": "unbound",
	"undo": "undone", "unwind": "unwound", "uphold": "upheld",
	"withstand": "withstood", "withhold": "withheld", "overcome": "overcome",
	"undergo": "undergone", "undertake": "undertaken", "mistake": "mistaken",
	"overtake": "overtaken",
}

// knownParticiples is a set of known irregular past participles for IsParticiple.
var knownParticiples = map[string]bool{
	// Irregular past participles ending in -en
	"been": true, "beaten": true, "bitten": true, "blown": true, "broken": true,
	"chosen": true, "driven": true, "eaten": true, "fallen": true, "forbidden": true,
	"forgotten": true, "forgiven": true, "frozen": true, "given": true, "gone": true,
	"grown": true, "hidden": true, "known": true, "lain": true, "proven": true,
	"ridden": true, "risen": true, "seen": true, "shaken": true, "shown": true,
	"spoken": true, "stolen": true, "stridden": true, "striven": true, "sworn": true,
	"taken": true, "torn": true, "trodden": true, "woken": true, "worn": true,
	"woven": true, "written": true,
	// Irregular past participles ending in -t
	"bent": true, "built": true, "burnt": true, "crept": true, "dealt": true,
	"dreamt": true, "dwelt": true, "felt": true, "kept": true, "knelt": true,
	"leant": true, "leapt": true, "learnt": true, "left": true, "lent": true,
	"lit": true, "lost": true, "meant": true, "met": true, "slept": true,
	"smelt": true, "spelt": true, "spent": true, "spilt": true, "spoilt": true,
	"swept": true, "wept": true,
	// Irregular -ght endings
	"bought": true, "brought": true, "caught": true, "fought": true,
	"sought": true, "taught": true, "thought": true,
	// Irregular -nt endings
	"sent": true, "rent": true,
	// Other irregular forms
	"bound": true, "bled": true, "bred": true, "clung": true, "done": true,
	"dug": true, "fed": true, "fled": true, "flung": true, "found": true,
	"ground": true, "had": true, "heard": true, "held": true, "hung": true,
	"laid": true, "led": true, "made": true, "paid": true, "rung": true,
	"said": true, "sat": true, "shed": true, "shone": true, "shot": true,
	"shrunk": true, "slid": true, "slit": true, "slung": true, "sold": true,
	"sped": true, "spun": true, "spat": true, "sprung": true, "stood": true,
	"stuck": true, "stung": true, "stunk": true, "struck": true, "strung": true,
	"sung": true, "sunk": true, "swum": true, "swung": true, "told": true,
	"understood": true, "withdrawn": true, "won": true, "wound": true, "wrung": true,
	// Unchanged forms (also participles)
	"bet": true, "burst": true, "come": true, "cost": true, "cut": true,
	"hit": true, "hurt": true, "let": true, "put": true, "quit": true,
	"read": true, "run": true, "set": true, "shut": true, "split": true,
	"spread": true, "thrust": true,
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

// PastParticiple converts a verb to its past participle form.
//
// Examples:
//   - PastParticiple("walk") returns "walked" (regular -ed)
//   - PastParticiple("stop") returns "stopped" (double consonant)
//   - PastParticiple("try") returns "tried" (y -> ied)
//   - PastParticiple("go") returns "gone" (irregular)
//   - PastParticiple("take") returns "taken" (irregular)
//   - PastParticiple("run") returns "run" (unchanged irregular)
func PastParticiple(verb string) string {
	if verb == "" {
		return ""
	}

	lower := strings.ToLower(verb)

	// Check for irregular verbs first
	if participle, ok := irregularPastParticiples[lower]; ok {
		return matchCase(verb, participle)
	}

	n := len(lower)

	// Single letter - just add -ed
	if n == 1 {
		return verb + matchSuffix(verb, "ed")
	}

	// Words ending in -e: just add -d
	if strings.HasSuffix(lower, "e") {
		return verb + matchSuffix(verb, "d")
	}

	// Words ending in consonant + y: change y to ied
	if strings.HasSuffix(lower, "y") && n >= 2 && !isVowel(rune(lower[n-2])) {
		return verb[:len(verb)-1] + matchSuffix(verb, "ied")
	}

	// Words ending in -c: add k before -ed
	if strings.HasSuffix(lower, "c") {
		return verb + matchSuffix(verb, "ked")
	}

	// Check for CVC pattern that requires doubling the final consonant
	if shouldDoubleConsonant(lower) {
		lastChar := verb[len(verb)-1:]
		return verb + matchSuffix(verb, strings.ToLower(lastChar)+"ed")
	}

	// Default: just add -ed
	return verb + matchSuffix(verb, "ed")
}

// IsParticiple checks if a word is a participle (present or past).
//
// Examples:
//   - IsParticiple("running") returns true (present participle)
//   - IsParticiple("walked") returns true (past participle)
//   - IsParticiple("taken") returns true (irregular past participle)
//   - IsParticiple("walk") returns false (base verb)
//   - IsParticiple("cat") returns false (not a verb form)
func IsParticiple(word string) bool {
	if word == "" {
		return false
	}

	lower := strings.ToLower(word)
	n := len(lower)

	// Check known irregular participles first
	if knownParticiples[lower] {
		return true
	}

	// Check for present participle (-ing)
	if strings.HasSuffix(lower, "ing") && n >= 4 {
		// Exclude base verbs like "sing", "ring", "bring", "thing"
		base := lower[:n-3]
		// If base is a known verb, it's likely a participle
		// For now, assume -ing words longer than 4 chars are participles
		// unless they're known non-participles
		nonParticiples := map[string]bool{
			"sing": true, "ring": true, "bring": true, "thing": true,
			"king": true, "wing": true, "spring": true, "string": true,
			"swing": true, "sting": true, "sling": true, "cling": true,
		}
		if nonParticiples[lower] {
			return false
		}
		// Check if it could be a base verb ending in -ing
		if len(base) >= 1 {
			return true
		}
	}

	// Check for regular past participle (-ed)
	if strings.HasSuffix(lower, "ed") && n >= 3 {
		// Exclude words that just happen to end in -ed but aren't participles
		nonParticiples := map[string]bool{
			"bed": true, "red": true, "shed": true, "wed": true,
			"sled": true, "ted": true, "ned": true, "fed": true,
		}
		// Most -ed words are participles
		if nonParticiples[lower] {
			return false
		}
		return true
	}

	// Check for -ied ending (cried, tried)
	if strings.HasSuffix(lower, "ied") && n >= 4 {
		return true
	}

	return false
}
