package inflect_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	inflect "github.com/cv/go-inflect/v2"
)

func TestPastTense(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		// Empty string
		{name: "empty", input: "", want: ""},

		// Regular verbs: add -ed
		{name: "walk", input: "walk", want: "walked"},
		{name: "talk", input: "talk", want: "talked"},
		{name: "work", input: "work", want: "worked"},
		{name: "play", input: "play", want: "played"}, // vowel + y
		{name: "stay", input: "stay", want: "stayed"},
		{name: "enjoy", input: "enjoy", want: "enjoyed"},
		{name: "destroy", input: "destroy", want: "destroyed"},
		{name: "help", input: "help", want: "helped"},
		{name: "start", input: "start", want: "started"},
		{name: "finish", input: "finish", want: "finished"},
		{name: "watch", input: "watch", want: "watched"},
		{name: "wash", input: "wash", want: "washed"},
		{name: "push", input: "push", want: "pushed"},
		{name: "pull", input: "pull", want: "pulled"},
		{name: "open", input: "open", want: "opened"},
		{name: "close", input: "close", want: "closed"}, // ends in e
		{name: "need", input: "need", want: "needed"},
		{name: "want", input: "want", want: "wanted"},
		{name: "ask", input: "ask", want: "asked"},
		{name: "answer", input: "answer", want: "answered"},
		{name: "clean", input: "clean", want: "cleaned"},
		{name: "cook", input: "cook", want: "cooked"},
		{name: "look", input: "look", want: "looked"},

		// Verbs ending in -e: add -d
		{name: "love", input: "love", want: "loved"},
		{name: "like", input: "like", want: "liked"},
		{name: "live", input: "live", want: "lived"},
		{name: "move", input: "move", want: "moved"},
		{name: "change", input: "change", want: "changed"},
		{name: "create", input: "create", want: "created"},
		{name: "use", input: "use", want: "used"},
		{name: "hope", input: "hope", want: "hoped"},
		{name: "smile", input: "smile", want: "smiled"},
		{name: "dance", input: "dance", want: "danced"},
		{name: "arrive", input: "arrive", want: "arrived"},
		{name: "decide", input: "decide", want: "decided"},
		{name: "believe", input: "believe", want: "believed"},
		{name: "receive", input: "receive", want: "received"},

		// Consonant + y: change y to -ied
		{name: "try", input: "try", want: "tried"},
		{name: "cry", input: "cry", want: "cried"},
		{name: "carry", input: "carry", want: "carried"},
		{name: "study", input: "study", want: "studied"},
		{name: "hurry", input: "hurry", want: "hurried"},
		{name: "worry", input: "worry", want: "worried"},
		{name: "marry", input: "marry", want: "married"},
		{name: "copy", input: "copy", want: "copied"},
		{name: "apply", input: "apply", want: "applied"},
		{name: "reply", input: "reply", want: "replied"},
		{name: "supply", input: "supply", want: "supplied"},
		{name: "occupy", input: "occupy", want: "occupied"},
		{name: "deny", input: "deny", want: "denied"},
		{name: "rely", input: "rely", want: "relied"},

		// CVC pattern: double final consonant
		{name: "stop", input: "stop", want: "stopped"},
		{name: "drop", input: "drop", want: "dropped"},
		{name: "shop", input: "shop", want: "shopped"},
		{name: "plan", input: "plan", want: "planned"},
		{name: "rob", input: "rob", want: "robbed"},
		{name: "rub", input: "rub", want: "rubbed"},
		{name: "hug", input: "hug", want: "hugged"},
		{name: "jog", input: "jog", want: "jogged"},
		{name: "grab", input: "grab", want: "grabbed"},
		{name: "trip", input: "trip", want: "tripped"},
		{name: "slip", input: "slip", want: "slipped"},
		{name: "step", input: "step", want: "stepped"},
		{name: "beg", input: "beg", want: "begged"},
		{name: "nod", input: "nod", want: "nodded"},
		{name: "chat", input: "chat", want: "chatted"},

		// Don't double w, x, y
		{name: "show", input: "show", want: "showed"},
		{name: "fix", input: "fix", want: "fixed"},
		{name: "box", input: "box", want: "boxed"},
		{name: "mix", input: "mix", want: "mixed"},

		// Irregular verbs
		{name: "go", input: "go", want: "went"},
		{name: "be", input: "be", want: "was"},
		{name: "have", input: "have", want: "had"},
		{name: "do", input: "do", want: "did"},
		{name: "say", input: "say", want: "said"},
		{name: "make", input: "make", want: "made"},
		{name: "get", input: "get", want: "got"},
		{name: "see", input: "see", want: "saw"},
		{name: "come", input: "come", want: "came"},
		{name: "take", input: "take", want: "took"},
		{name: "know", input: "know", want: "knew"},
		{name: "think", input: "think", want: "thought"},
		{name: "find", input: "find", want: "found"},
		{name: "give", input: "give", want: "gave"},
		{name: "tell", input: "tell", want: "told"},
		{name: "become", input: "become", want: "became"},
		{name: "leave", input: "leave", want: "left"},
		{name: "put", input: "put", want: "put"},
		{name: "keep", input: "keep", want: "kept"},
		{name: "let", input: "let", want: "let"},
		{name: "begin", input: "begin", want: "began"},
		{name: "run", input: "run", want: "ran"},
		{name: "write", input: "write", want: "wrote"},
		{name: "read", input: "read", want: "read"},
		{name: "bring", input: "bring", want: "brought"},
		{name: "buy", input: "buy", want: "bought"},
		{name: "catch", input: "catch", want: "caught"},
		{name: "teach", input: "teach", want: "taught"},
		{name: "fight", input: "fight", want: "fought"},
		{name: "build", input: "build", want: "built"},
		{name: "send", input: "send", want: "sent"},
		{name: "spend", input: "spend", want: "spent"},
		{name: "lose", input: "lose", want: "lost"},
		{name: "feel", input: "feel", want: "felt"},
		{name: "meet", input: "meet", want: "met"},
		{name: "sit", input: "sit", want: "sat"},
		{name: "stand", input: "stand", want: "stood"},
		{name: "hear", input: "hear", want: "heard"},
		{name: "hold", input: "hold", want: "held"},
		{name: "speak", input: "speak", want: "spoke"},
		{name: "break", input: "break", want: "broke"},
		{name: "choose", input: "choose", want: "chose"},
		{name: "grow", input: "grow", want: "grew"},
		{name: "throw", input: "throw", want: "threw"},
		{name: "blow", input: "blow", want: "blew"},
		{name: "fly", input: "fly", want: "flew"},
		{name: "draw", input: "draw", want: "drew"},
		{name: "drive", input: "drive", want: "drove"},
		{name: "ride", input: "ride", want: "rode"},
		{name: "rise", input: "rise", want: "rose"},
		{name: "write", input: "write", want: "wrote"},
		{name: "hide", input: "hide", want: "hid"},
		{name: "eat", input: "eat", want: "ate"},
		{name: "fall", input: "fall", want: "fell"},
		{name: "swim", input: "swim", want: "swam"},
		{name: "sing", input: "sing", want: "sang"},
		{name: "ring", input: "ring", want: "rang"},
		{name: "drink", input: "drink", want: "drank"},
		{name: "sink", input: "sink", want: "sank"},
		{name: "win", input: "win", want: "won"},
		{name: "hit", input: "hit", want: "hit"},
		{name: "cut", input: "cut", want: "cut"},
		{name: "shut", input: "shut", want: "shut"},
		{name: "set", input: "set", want: "set"},
		{name: "hurt", input: "hurt", want: "hurt"},
		{name: "cost", input: "cost", want: "cost"},
		{name: "sleep", input: "sleep", want: "slept"},
		{name: "wake", input: "wake", want: "woke"},
		{name: "wear", input: "wear", want: "wore"},
		{name: "tear", input: "tear", want: "tore"},
		{name: "bear", input: "bear", want: "bore"},
		{name: "swear", input: "swear", want: "swore"},
		{name: "steal", input: "steal", want: "stole"},
		{name: "freeze", input: "freeze", want: "froze"},
		{name: "forget", input: "forget", want: "forgot"},
		{name: "forgive", input: "forgive", want: "forgave"},
		{name: "bite", input: "bite", want: "bit"},
		{name: "shake", input: "shake", want: "shook"},
		{name: "mistake", input: "mistake", want: "mistook"},
		{name: "undertake", input: "undertake", want: "undertook"},
		{name: "shine", input: "shine", want: "shone"},
		{name: "lie", input: "lie", want: "lay"},
		{name: "lay", input: "lay", want: "laid"},
		{name: "pay", input: "pay", want: "paid"},
		{name: "mean", input: "mean", want: "meant"},
		{name: "lean", input: "lean", want: "leaned"},    // regular
		{name: "learn", input: "learn", want: "learned"}, // regular (American)
		{name: "burn", input: "burn", want: "burned"},    // regular (American)
		{name: "dream", input: "dream", want: "dreamed"}, // regular (American)
		{name: "leap", input: "leap", want: "leaped"},    // regular (American)
		{name: "spell", input: "spell", want: "spelled"}, // regular (American)
		{name: "smell", input: "smell", want: "smelled"}, // regular (American)
		{name: "spill", input: "spill", want: "spilled"}, // regular (American)

		// New irregular forms (unchanged)
		{name: "bet", input: "bet", want: "bet"},
		{name: "burst", input: "burst", want: "burst"},
		{name: "cast", input: "cast", want: "cast"},
		{name: "forecast", input: "forecast", want: "forecast"},
		{name: "fit", input: "fit", want: "fit"},
		{name: "upset", input: "upset", want: "upset"},
		{name: "thrust", input: "thrust", want: "thrust"},

		// New irregular forms
		{name: "deal", input: "deal", want: "dealt"},
		{name: "dwell", input: "dwell", want: "dwelt"},
		{name: "kneel", input: "kneel", want: "knelt"},
		{name: "light", input: "light", want: "lit"},
		{name: "slay", input: "slay", want: "slew"},
		{name: "stride", input: "stride", want: "strode"},
		{name: "strive", input: "strive", want: "strove"},
		{name: "tread", input: "tread", want: "trod"},
		{name: "weave", input: "weave", want: "wove"},
		{name: "fling", input: "fling", want: "flung"},
		{name: "wring", input: "wring", want: "wrung"},

		// Compound verbs
		{name: "foresee", input: "foresee", want: "foresaw"},
		{name: "outdo", input: "outdo", want: "outdid"},
		{name: "overdo", input: "overdo", want: "overdid"},
		{name: "redo", input: "redo", want: "redid"},
		{name: "undo", input: "undo", want: "undid"},
		{name: "rebuild", input: "rebuild", want: "rebuilt"},
		{name: "uphold", input: "uphold", want: "upheld"},

		// Case preservation
		{name: "WALK uppercase", input: "WALK", want: "WALKED"},
		{name: "Walk titlecase", input: "Walk", want: "Walked"},
		{name: "GO uppercase", input: "GO", want: "WENT"},
		{name: "Go titlecase", input: "Go", want: "Went"},
		{name: "TRY uppercase", input: "TRY", want: "TRIED"},
		{name: "Try titlecase", input: "Try", want: "Tried"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.PastTense(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func BenchmarkPastTense(b *testing.B) {
	benchmarks := []struct {
		name  string
		input string
	}{
		{"regular_ed", "walk"},
		{"ending_e", "love"},
		{"consonant_y", "try"},
		{"cvc_double", "stop"},
		{"irregular", "go"},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for range b.N {
				inflect.PastTense(bm.input)
			}
		})
	}
}
