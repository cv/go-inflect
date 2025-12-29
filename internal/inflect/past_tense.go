package inflect

import "strings"

// irregularPastTense maps verbs to their irregular past tense forms.
var irregularPastTense = map[string]string{
	// Be, have, do
	"be":   "was",
	"am":   "was",
	"is":   "was",
	"are":  "were",
	"have": "had",
	"has":  "had",
	"do":   "did",
	"does": "did",

	// Common irregular verbs
	"go":         "went",
	"say":        "said",
	"make":       "made",
	"get":        "got",
	"see":        "saw",
	"come":       "came",
	"take":       "took",
	"know":       "knew",
	"think":      "thought",
	"find":       "found",
	"give":       "gave",
	"tell":       "told",
	"become":     "became",
	"leave":      "left",
	"put":        "put",
	"keep":       "kept",
	"let":        "let",
	"begin":      "began",
	"run":        "ran",
	"write":      "wrote",
	"read":       "read",
	"bring":      "brought",
	"buy":        "bought",
	"catch":      "caught",
	"teach":      "taught",
	"fight":      "fought",
	"build":      "built",
	"send":       "sent",
	"spend":      "spent",
	"lose":       "lost",
	"feel":       "felt",
	"meet":       "met",
	"sit":        "sat",
	"stand":      "stood",
	"hear":       "heard",
	"hold":       "held",
	"speak":      "spoke",
	"break":      "broke",
	"choose":     "chose",
	"grow":       "grew",
	"throw":      "threw",
	"blow":       "blew",
	"fly":        "flew",
	"draw":       "drew",
	"drive":      "drove",
	"ride":       "rode",
	"rise":       "rose",
	"hide":       "hid",
	"eat":        "ate",
	"fall":       "fell",
	"swim":       "swam",
	"sing":       "sang",
	"ring":       "rang",
	"drink":      "drank",
	"sink":       "sank",
	"win":        "won",
	"hit":        "hit",
	"cut":        "cut",
	"shut":       "shut",
	"set":        "set",
	"hurt":       "hurt",
	"cost":       "cost",
	"sleep":      "slept",
	"wake":       "woke",
	"wear":       "wore",
	"tear":       "tore",
	"bear":       "bore",
	"swear":      "swore",
	"steal":      "stole",
	"freeze":     "froze",
	"forget":     "forgot",
	"forgive":    "forgave",
	"bite":       "bit",
	"shake":      "shook",
	"shine":      "shone",
	"lie":        "lay",
	"lay":        "laid",
	"pay":        "paid",
	"mean":       "meant",
	"seek":       "sought",
	"sell":       "sold",
	"lend":       "lent",
	"bend":       "bent",
	"dig":        "dug",
	"stick":      "stuck",
	"strike":     "struck",
	"hang":       "hung",
	"swing":      "swung",
	"cling":      "clung",
	"spin":       "spun",
	"sting":      "stung",
	"sling":      "slung",
	"spring":     "sprang",
	"shrink":     "shrank",
	"stink":      "stank",
	"feed":       "fed",
	"bleed":      "bled",
	"breed":      "bred",
	"lead":       "led",
	"speed":      "sped",
	"flee":       "fled",
	"creep":      "crept",
	"weep":       "wept",
	"sweep":      "swept",
	"split":      "split",
	"quit":       "quit",
	"spit":       "spat",
	"spread":     "spread",
	"shed":       "shed",
	"bid":        "bid",
	"rid":        "rid",
	"slide":      "slid",
	"grind":      "ground",
	"bind":       "bound",
	"wind":       "wound",
	"withdraw":   "withdrew",
	"withstand":  "withstood",
	"withhold":   "withheld",
	"overcome":   "overcame",
	"undergo":    "underwent",
	"understand": "understood",
	"undertake":  "undertook",
	"mistake":    "mistook",
	"overtake":   "overtook",
	"awake":      "awoke",
	"arise":      "arose",

	// Unchanged past tense forms
	"bet":       "bet",
	"burst":     "burst",
	"cast":      "cast",
	"forecast":  "forecast",
	"preset":    "preset",
	"reset":     "reset",
	"proofread": "proofread",
	"reread":    "reread",
	"sublet":    "sublet",
	"slit":      "slit",
	"thrust":    "thrust",
	"wed":       "wed",
	"fit":       "fit",
	"knit":      "knit",
	"upset":     "upset",
	"bust":      "bust",
	"shit":      "shit",

	// Additional irregular verbs
	"beseech": "besought",
	"beget":   "begot",
	"deal":    "dealt",
	"dwell":   "dwelt",
	"fling":   "flung",
	"forsake": "forsook",
	"inlay":   "inlaid",
	"kneel":   "knelt",
	"light":   "lit",
	"slay":    "slew",
	"stride":  "strode",
	"strive":  "strove",
	"tread":   "trod",
	"weave":   "wove",
	"wring":   "wrung",

	// Compound verbs with irregular bases
	"foresee":   "foresaw",
	"outdo":     "outdid",
	"outgrow":   "outgrew",
	"overdo":    "overdid",
	"overhear":  "overheard",
	"override":  "overrode",
	"oversee":   "oversaw",
	"oversleep": "overslept",
	"overthrow": "overthrew",
	"partake":   "partook",
	"rebuild":   "rebuilt",
	"redo":      "redid",
	"remake":    "remade",
	"repay":     "repaid",
	"retell":    "retold",
	"rewind":    "rewound",
	"rewrite":   "rewrote",
	"unbind":    "unbound",
	"undo":      "undid",
	"unwind":    "unwound",
	"uphold":    "upheld",

	// Can/could, may/might, etc. are modal verbs
	"can":   "could",
	"may":   "might",
	"shall": "should",
	"will":  "would",
}

// PastTense returns the simple past tense form of an English verb.
//
// Examples:
//   - PastTense("walk") returns "walked"
//   - PastTense("go") returns "went"
//   - PastTense("try") returns "tried"
//   - PastTense("stop") returns "stopped"
func PastTense(verb string) string {
	if verb == "" {
		return ""
	}

	lower := strings.ToLower(verb)

	// Check irregular verbs first
	if past, ok := irregularPastTense[lower]; ok {
		return matchCase(verb, past)
	}

	// Apply regular rules
	return applyPastTenseRules(verb, lower)
}

// applyPastTenseRules applies regular past tense formation rules.
func applyPastTenseRules(verb, lower string) string {
	// Verbs ending in -e: add -d
	if strings.HasSuffix(lower, "e") {
		return verb + matchSuffix(verb, "d")
	}

	// Verbs ending in consonant + y: change y to -ied
	if strings.HasSuffix(lower, "y") && len(lower) > 1 {
		beforeY := lower[len(lower)-2]
		if !isVowel(rune(beforeY)) {
			return verb[:len(verb)-1] + matchSuffix(verb, "ied")
		}
	}

	// CVC pattern: double the final consonant and add -ed
	if shouldDoubleFinalConsonantForPast(lower) {
		lastChar := string(lower[len(lower)-1])
		return verb + matchSuffix(verb, lastChar+"ed")
	}

	// Default: add -ed
	return verb + matchSuffix(verb, "ed")
}

// shouldDoubleFinalConsonantForPast checks if the final consonant should be doubled
// when forming the past tense. This applies to short verbs with a CVC pattern.
func shouldDoubleFinalConsonantForPast(lower string) bool {
	if len(lower) < 2 {
		return false
	}

	// Only for short (one-syllable) words
	if countSyllables(lower) != 1 {
		return false
	}

	lastChar := rune(lower[len(lower)-1])
	secondLastChar := rune(lower[len(lower)-2])

	// Last char must be a consonant (not w, x, or y)
	if isVowel(lastChar) || lastChar == 'w' || lastChar == 'x' || lastChar == 'y' {
		return false
	}

	// Second-to-last must be a single vowel
	if !isVowel(secondLastChar) {
		return false
	}

	// Check that there's a consonant before the vowel (CVC pattern)
	if len(lower) >= 3 {
		thirdLastChar := rune(lower[len(lower)-3])
		if isVowel(thirdLastChar) {
			return false // VVC pattern, don't double
		}
	}

	return true
}
