package inflect

// This file contains irregular verb data shared between past_tense.go and participle.go.
// Many irregular verbs have identical past tense and past participle forms (e.g., "bought",
// "brought", "caught"). These are stored in irregularVerbsSame to avoid duplication.
// Verbs with different forms are stored separately in irregularPastTenseOnly and
// irregularPastParticipleOnly.

// irregularVerbsSame contains verbs where past tense == past participle.
// Example: buy -> bought (past), bought (participle).
//
//nolint:dupl // Similar structure to other verb maps but different data.
var irregularVerbsSame = map[string]string{
	// Common verbs with same past/participle
	"bend": "bent", "bet": "bet", "bid": "bid", "bind": "bound",
	"bleed": "bled", "breed": "bred", "bring": "brought", "build": "built",
	"burst": "burst", "buy": "bought", "catch": "caught", "cling": "clung",
	"cost": "cost", "creep": "crept", "cut": "cut", "deal": "dealt",
	"dig": "dug", "feed": "fed", "feel": "felt", "fight": "fought",
	"find": "found", "flee": "fled", "fling": "flung", "get": "got",
	"grind": "ground", "hang": "hung", "have": "had", "hear": "heard",
	"hit": "hit", "hold": "held", "hurt": "hurt", "keep": "kept",
	"kneel": "knelt", "lay": "laid", "lead": "led", "leave": "left",
	"lend": "lent", "let": "let", "light": "lit", "lose": "lost",
	"make": "made", "mean": "meant", "meet": "met", "pay": "paid",
	"put": "put", "quit": "quit", "read": "read", "say": "said",
	"seek": "sought", "sell": "sold", "send": "sent", "set": "set",
	"shed": "shed", "shine": "shone", "shut": "shut", "sit": "sat",
	"sleep": "slept", "slide": "slid", "sling": "slung", "slit": "slit",
	"speed": "sped", "spend": "spent", "spin": "spun", "spit": "spat",
	"split": "split", "spread": "spread", "stand": "stood", "stick": "stuck",
	"sting": "stung", "strike": "struck", "sweep": "swept", "swing": "swung",
	"teach": "taught", "tell": "told", "think": "thought", "thrust": "thrust",
	"weep": "wept", "win": "won", "wind": "wound", "wring": "wrung",

	// Compound verbs with same past/participle
	"inlay": "inlaid", "overhear": "overheard", "oversleep": "overslept",
	"rebuild": "rebuilt", "remake": "remade", "repay": "repaid",
	"retell": "retold", "rewind": "rewound", "unbind": "unbound",
	"understand": "understood", "unwind": "unwound", "uphold": "upheld",
	"withhold": "withheld", "withstand": "withstood",

	// Additional verbs with same past/participle
	"abide": "abode", "behold": "beheld", "beseech": "besought",
	"beset": "beset", "dwell": "dwelt", "wet": "wet",
}

// irregularPastTenseOnly contains verbs where past tense differs from participle.
// These are looked up only for past tense; participles are in irregularPastParticipleOnly.
//
//nolint:dupl // Similar structure to other verb maps but different data.
var irregularPastTenseOnly = map[string]string{
	// Be, have, do variants
	"be": "was", "am": "was", "is": "was", "are": "were",
	"has": "had", "do": "did", "does": "did",

	// Common verbs with different past/participle
	"arise": "arose", "awake": "awoke", "bear": "bore", "beat": "beat",
	"become": "became", "begin": "began", "bite": "bit", "blow": "blew",
	"break": "broke", "choose": "chose", "come": "came", "draw": "drew",
	"drink": "drank", "drive": "drove", "eat": "ate", "fall": "fell",
	"fly": "flew", "forget": "forgot", "forgive": "forgave", "freeze": "froze",
	"give": "gave", "go": "went", "grow": "grew", "hide": "hid",
	"know": "knew", "lie": "lay", "ride": "rode", "ring": "rang",
	"rise": "rose", "run": "ran", "see": "saw", "shake": "shook",
	"shrink": "shrank", "sing": "sang", "sink": "sank", "speak": "spoke",
	"spring": "sprang", "steal": "stole", "stink": "stank", "stride": "strode",
	"swear": "swore", "swim": "swam", "take": "took", "tear": "tore",
	"throw": "threw", "tread": "trod", "wake": "woke", "wear": "wore",
	"weave": "wove", "write": "wrote",

	// Compound verbs with different past/participle
	"beget": "begot", "cleave": "clove", "forego": "forewent",
	"foresee": "foresaw", "forsake": "forsook", "mistake": "mistook",
	"outdo": "outdid", "outgrow": "outgrew", "overcome": "overcame",
	"overdo": "overdid", "override": "overrode", "oversee": "oversaw",
	"overtake": "overtook", "overthrow": "overthrew", "partake": "partook",
	"redo": "redid", "rewrite": "rewrote", "slay": "slew", "smite": "smote",
	"strive": "strove", "undergo": "underwent", "undertake": "undertook",
	"undo": "undid", "withdraw": "withdrew",

	// Unchanged past tense forms (different from participle or no participle)
	"bet": "bet", "bust": "bust", "cast": "cast", "fit": "fit",
	"forecast": "forecast", "knit": "knit", "preset": "preset",
	"proofread": "proofread", "reread": "reread", "reset": "reset",
	"rid": "rid", "shit": "shit", "sublet": "sublet", "upset": "upset",
	"wed": "wed",

	// Modal verbs
	"can": "could", "may": "might", "shall": "should", "will": "would",
}

// irregularPastParticipleOnly contains verbs where participle differs from past tense.
// These are looked up only for past participle; past tense is in irregularPastTenseOnly.
var irregularPastParticipleOnly = map[string]string{
	// Be, do
	"be": "been", "do": "done",

	// Common verbs with different past/participle
	"arise": "arisen", "awake": "awoken", "bear": "borne", "beat": "beaten",
	"become": "become", "begin": "begun", "bite": "bitten", "blow": "blown",
	"break": "broken", "choose": "chosen", "come": "come", "draw": "drawn",
	"drink": "drunk", "drive": "driven", "eat": "eaten", "fall": "fallen",
	"fly": "flown", "forbid": "forbidden", "forget": "forgotten",
	"forgive": "forgiven", "freeze": "frozen", "give": "given", "go": "gone",
	"grow": "grown", "hide": "hidden", "know": "known", "lie": "lain",
	"ride": "ridden", "ring": "rung", "rise": "risen", "run": "run",
	"see": "seen", "shake": "shaken", "shrink": "shrunk", "sing": "sung",
	"sink": "sunk", "speak": "spoken", "spring": "sprung", "steal": "stolen",
	"stink": "stunk", "stride": "stridden", "swear": "sworn", "swim": "swum",
	"take": "taken", "tear": "torn", "throw": "thrown", "tread": "trodden",
	"wake": "woken", "wear": "worn", "weave": "woven", "write": "written",

	// Compound verbs with different past/participle
	"beget": "begotten", "cleave": "cloven", "forego": "foregone",
	"foresee": "foreseen", "forsake": "forsaken", "mistake": "mistaken",
	"outdo": "outdone", "outgrow": "outgrown", "overcome": "overcome",
	"overdo": "overdone", "override": "overridden", "oversee": "overseen",
	"overtake": "overtaken", "overthrow": "overthrown", "partake": "partaken",
	"redo": "redone", "rewrite": "rewritten", "slay": "slain", "smite": "smitten",
	"strive": "striven", "undergo": "undergone", "undertake": "undertaken",
	"undo": "undone", "withdraw": "withdrawn",

	// Participles ending in -n (from regular past tense verbs)
	"hew": "hewn", "lade": "laden", "mow": "mown", "prove": "proven",
	"saw": "sawn", "sew": "sewn", "shear": "shorn", "show": "shown",
	"sow": "sown", "strew": "strewn", "string": "strung",

	// Participles ending in -t (British spellings)
	"burn": "burnt", "dream": "dreamt", "lean": "leant", "leap": "leapt",
	"learn": "learnt", "smell": "smelt", "spell": "spelt", "spill": "spilt",
	"spoil": "spoilt",

	// Other participle-only forms
	"shoot": "shot",
}
