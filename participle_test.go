package inflect_test

import (
	"testing"

	inflect "github.com/cv/go-inflect"
)

func TestPresentParticiple(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		// Empty string
		{name: "empty", input: "", want: ""},

		// Single letter verbs
		{name: "single letter", input: "a", want: "aing"},

		// Already ending in -ing
		{name: "already -ing", input: "running", want: "running"},
		{name: "already -ing sing", input: "sing", want: "singing"},

		// Double consonant (CVC pattern)
		{name: "run", input: "run", want: "running"},
		{name: "sit", input: "sit", want: "sitting"},
		{name: "hit", input: "hit", want: "hitting"},
		{name: "cut", input: "cut", want: "cutting"},
		{name: "stop", input: "stop", want: "stopping"},
		{name: "drop", input: "drop", want: "dropping"},
		{name: "plan", input: "plan", want: "planning"},
		{name: "skip", input: "skip", want: "skipping"},
		{name: "begin", input: "begin", want: "beginning"},
		{name: "occur", input: "occur", want: "occurring"},
		{name: "prefer", input: "prefer", want: "preferring"},
		{name: "admit", input: "admit", want: "admitting"},
		{name: "commit", input: "commit", want: "committing"},
		{name: "regret", input: "regret", want: "regretting"},

		// Drop silent e
		{name: "make", input: "make", want: "making"},
		{name: "take", input: "take", want: "taking"},
		{name: "come", input: "come", want: "coming"},
		{name: "give", input: "give", want: "giving"},
		{name: "have", input: "have", want: "having"},
		{name: "write", input: "write", want: "writing"},
		{name: "live", input: "live", want: "living"},
		{name: "move", input: "move", want: "moving"},
		{name: "hope", input: "hope", want: "hoping"},
		{name: "dance", input: "dance", want: "dancing"},

		// Just add -ing (no changes needed)
		{name: "play", input: "play", want: "playing"},
		{name: "stay", input: "stay", want: "staying"},
		{name: "enjoy", input: "enjoy", want: "enjoying"},
		{name: "show", input: "show", want: "showing"},
		{name: "follow", input: "follow", want: "following"},
		{name: "fix", input: "fix", want: "fixing"},
		{name: "mix", input: "mix", want: "mixing"},
		{name: "go", input: "go", want: "going"},
		{name: "do", input: "do", want: "doing"},
		{name: "eat", input: "eat", want: "eating"},
		{name: "read", input: "read", want: "reading"},
		{name: "think", input: "think", want: "thinking"},
		{name: "walk", input: "walk", want: "walking"},
		{name: "talk", input: "talk", want: "talking"},
		{name: "open", input: "open", want: "opening"},
		{name: "listen", input: "listen", want: "listening"},
		{name: "visit", input: "visit", want: "visiting"},

		// ie -> ying
		{name: "die", input: "die", want: "dying"},
		{name: "lie", input: "lie", want: "lying"},
		{name: "tie", input: "tie", want: "tying"},

		// ee -> eeing
		{name: "see", input: "see", want: "seeing"},
		{name: "flee", input: "flee", want: "fleeing"},
		{name: "agree", input: "agree", want: "agreeing"},
		{name: "free", input: "free", want: "freeing"},

		// be -> being (special vowel + e case)
		{name: "be", input: "be", want: "being"},

		// Words ending in -c (add k)
		{name: "panic", input: "panic", want: "panicking"},
		{name: "picnic", input: "picnic", want: "picnicking"},
		{name: "traffic", input: "traffic", want: "trafficking"},
		{name: "mimic", input: "mimic", want: "mimicking"},
		{name: "frolic", input: "frolic", want: "frolicking"},

		// Words ending in -ye, -oe (keep e)
		{name: "dye", input: "dye", want: "dyeing"},
		{name: "hoe", input: "hoe", want: "hoeing"},
		{name: "toe", input: "toe", want: "toeing"},

		// Words ending in -nge/-inge (keep e)
		{name: "singe", input: "singe", want: "singeing"},

		// Case preservation
		{name: "RUN uppercase", input: "RUN", want: "RUNNING"},
		{name: "Run titlecase", input: "Run", want: "Running"},
		{name: "MAKE uppercase", input: "MAKE", want: "MAKING"},
		{name: "Make titlecase", input: "Make", want: "Making"},
		{name: "DIE uppercase", input: "DIE", want: "DYING"},
		{name: "PANIC uppercase", input: "PANIC", want: "PANICKING"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.PresentParticiple(tt.input)
			if got != tt.want {
				t.Errorf("PresentParticiple(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func BenchmarkPresentParticiple(b *testing.B) {
	// Test with verbs covering different transformation rules
	benchmarks := []struct {
		name  string
		input string
	}{
		{"add_ing", "play"},         // playing (just add -ing)
		{"double_consonant", "run"}, // running (double final consonant)
		{"drop_e", "make"},          // making (drop silent e)
		{"ie_to_y", "die"},          // dying (ie -> y)
		{"ee_keep", "see"},          // seeing (keep ee)
		{"add_k", "panic"},          // panicking (add k before -ing)
		{"already_ing", "sing"},     // singing
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for range b.N {
				inflect.PresentParticiple(bm.input)
			}
		})
	}
}

func TestPastParticiple(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		// Empty string
		{name: "empty", input: "", want: ""},

		// Regular verbs (-ed)
		{name: "walk", input: "walk", want: "walked"},
		{name: "talk", input: "talk", want: "talked"},
		{name: "play", input: "play", want: "played"},
		{name: "stay", input: "stay", want: "stayed"},
		{name: "work", input: "work", want: "worked"},
		{name: "help", input: "help", want: "helped"},
		{name: "ask", input: "ask", want: "asked"},
		{name: "call", input: "call", want: "called"},
		{name: "open", input: "open", want: "opened"},
		{name: "listen", input: "listen", want: "listened"},

		// Verbs ending in -e (just add -d)
		{name: "like", input: "like", want: "liked"},
		{name: "love", input: "love", want: "loved"},
		{name: "dance", input: "dance", want: "danced"},
		{name: "hope", input: "hope", want: "hoped"},
		{name: "use", input: "use", want: "used"},
		{name: "close", input: "close", want: "closed"},

		// Verbs ending in consonant + y (y -> ied)
		{name: "try", input: "try", want: "tried"},
		{name: "cry", input: "cry", want: "cried"},
		{name: "study", input: "study", want: "studied"},
		{name: "carry", input: "carry", want: "carried"},
		{name: "worry", input: "worry", want: "worried"},

		// Verbs ending in vowel + y (just add -ed)
		{name: "play already", input: "play", want: "played"},
		{name: "enjoy", input: "enjoy", want: "enjoyed"},
		{name: "stay already", input: "stay", want: "stayed"},
		{name: "delay", input: "delay", want: "delayed"},

		// CVC pattern - double consonant
		{name: "stop", input: "stop", want: "stopped"},
		{name: "drop", input: "drop", want: "dropped"},
		{name: "plan", input: "plan", want: "planned"},
		{name: "skip", input: "skip", want: "skipped"},
		{name: "admit", input: "admit", want: "admitted"},
		{name: "occur", input: "occur", want: "occurred"},
		{name: "prefer", input: "prefer", want: "preferred"},
		{name: "regret", input: "regret", want: "regretted"},

		// Don't double w, x, y
		{name: "fix", input: "fix", want: "fixed"},
		{name: "mix", input: "mix", want: "mixed"},
		{name: "show", input: "show", want: "shown"},

		// Verbs ending in -c (add k)
		{name: "panic", input: "panic", want: "panicked"},
		{name: "picnic", input: "picnic", want: "picnicked"},
		{name: "traffic", input: "traffic", want: "trafficked"},

		// Irregular verbs (common ones)
		{name: "go", input: "go", want: "gone"},
		{name: "be", input: "be", want: "been"},
		{name: "have", input: "have", want: "had"},
		{name: "do", input: "do", want: "done"},
		{name: "say", input: "say", want: "said"},
		{name: "get", input: "get", want: "got"},
		{name: "make", input: "make", want: "made"},
		{name: "know", input: "know", want: "known"},
		{name: "think", input: "think", want: "thought"},
		{name: "take", input: "take", want: "taken"},
		{name: "see", input: "see", want: "seen"},
		{name: "come", input: "come", want: "come"},
		{name: "give", input: "give", want: "given"},
		{name: "find", input: "find", want: "found"},
		{name: "tell", input: "tell", want: "told"},
		{name: "write", input: "write", want: "written"},
		{name: "run", input: "run", want: "run"},
		{name: "eat", input: "eat", want: "eaten"},
		{name: "drink", input: "drink", want: "drunk"},
		{name: "sing", input: "sing", want: "sung"},
		{name: "swim", input: "swim", want: "swum"},
		{name: "begin", input: "begin", want: "begun"},
		{name: "break", input: "break", want: "broken"},
		{name: "choose", input: "choose", want: "chosen"},
		{name: "speak", input: "speak", want: "spoken"},
		{name: "steal", input: "steal", want: "stolen"},
		{name: "forget", input: "forget", want: "forgotten"},
		{name: "drive", input: "drive", want: "driven"},
		{name: "ride", input: "ride", want: "ridden"},
		{name: "hide", input: "hide", want: "hidden"},
		{name: "bite", input: "bite", want: "bitten"},
		{name: "fly", input: "fly", want: "flown"},
		{name: "grow", input: "grow", want: "grown"},
		{name: "throw", input: "throw", want: "thrown"},
		{name: "draw", input: "draw", want: "drawn"},
		{name: "fall", input: "fall", want: "fallen"},
		{name: "buy", input: "buy", want: "bought"},
		{name: "bring", input: "bring", want: "brought"},
		{name: "catch", input: "catch", want: "caught"},
		{name: "teach", input: "teach", want: "taught"},
		{name: "fight", input: "fight", want: "fought"},
		{name: "seek", input: "seek", want: "sought"},
		{name: "feel", input: "feel", want: "felt"},
		{name: "keep", input: "keep", want: "kept"},
		{name: "sleep", input: "sleep", want: "slept"},
		{name: "leave", input: "leave", want: "left"},
		{name: "meet", input: "meet", want: "met"},
		{name: "read", input: "read", want: "read"},
		{name: "lead", input: "lead", want: "led"},
		{name: "sit", input: "sit", want: "sat"},
		{name: "stand", input: "stand", want: "stood"},
		{name: "lose", input: "lose", want: "lost"},
		{name: "win", input: "win", want: "won"},
		{name: "put", input: "put", want: "put"},
		{name: "cut", input: "cut", want: "cut"},
		{name: "hit", input: "hit", want: "hit"},
		{name: "let", input: "let", want: "let"},
		{name: "set", input: "set", want: "set"},
		{name: "shut", input: "shut", want: "shut"},
		{name: "hurt", input: "hurt", want: "hurt"},
		{name: "cost", input: "cost", want: "cost"},
		{name: "build", input: "build", want: "built"},
		{name: "send", input: "send", want: "sent"},
		{name: "spend", input: "spend", want: "spent"},
		{name: "lend", input: "lend", want: "lent"},
		{name: "bend", input: "bend", want: "bent"},

		// Case preservation
		{name: "WALK uppercase", input: "WALK", want: "WALKED"},
		{name: "Walk titlecase", input: "Walk", want: "Walked"},
		{name: "GO uppercase", input: "GO", want: "GONE"},
		{name: "Go titlecase", input: "Go", want: "Gone"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.PastParticiple(tt.input)
			if got != tt.want {
				t.Errorf("PastParticiple(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsParticiple(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		// Present participles
		{name: "running", input: "running", want: true},
		{name: "walking", input: "walking", want: true},
		{name: "playing", input: "playing", want: true},
		{name: "making", input: "making", want: true},
		{name: "dying", input: "dying", want: true},
		{name: "seeing", input: "seeing", want: true},

		// Past participles (regular)
		{name: "walked", input: "walked", want: true},
		{name: "played", input: "played", want: true},
		{name: "stopped", input: "stopped", want: true},
		{name: "tried", input: "tried", want: true},

		// Past participles (irregular -en)
		{name: "taken", input: "taken", want: true},
		{name: "eaten", input: "eaten", want: true},
		{name: "written", input: "written", want: true},
		{name: "given", input: "given", want: true},
		{name: "driven", input: "driven", want: true},
		{name: "chosen", input: "chosen", want: true},
		{name: "spoken", input: "spoken", want: true},
		{name: "broken", input: "broken", want: true},
		{name: "forgotten", input: "forgotten", want: true},

		// Past participles (irregular -t)
		{name: "kept", input: "kept", want: true},
		{name: "slept", input: "slept", want: true},
		{name: "felt", input: "felt", want: true},
		{name: "left", input: "left", want: true},
		{name: "built", input: "built", want: true},
		{name: "sent", input: "sent", want: true},
		{name: "spent", input: "spent", want: true},
		{name: "thought", input: "thought", want: true},
		{name: "bought", input: "bought", want: true},
		{name: "brought", input: "brought", want: true},
		{name: "caught", input: "caught", want: true},
		{name: "taught", input: "taught", want: true},
		{name: "fought", input: "fought", want: true},
		{name: "sought", input: "sought", want: true},

		// Other irregular participles
		{name: "gone", input: "gone", want: true},
		{name: "done", input: "done", want: true},
		{name: "seen", input: "seen", want: true},
		{name: "been", input: "been", want: true},
		{name: "had", input: "had", want: true},
		{name: "made", input: "made", want: true},
		{name: "said", input: "said", want: true},

		// Not participles (base verbs that aren't also participles)
		{name: "walk", input: "walk", want: false},
		// Note: "run" is both base and participle (I run / I have run)
		{name: "play", input: "play", want: false},
		{name: "sing", input: "sing", want: false},
		{name: "bring", input: "bring", want: false},
		{name: "thing", input: "thing", want: false},
		{name: "ring", input: "ring", want: false},

		// Words that are both base and participle (unchanged verbs)
		{name: "run (also participle)", input: "run", want: true},
		{name: "cut (also participle)", input: "cut", want: true},
		{name: "put (also participle)", input: "put", want: true},
		{name: "hit (also participle)", input: "hit", want: true},

		// Not participles (other words)
		{name: "cat", input: "cat", want: false},
		{name: "dog", input: "dog", want: false},
		{name: "red", input: "red", want: false},
		{name: "bed", input: "bed", want: false},
		{name: "kid", input: "kid", want: false},
		{name: "open", input: "open", want: false},
		{name: "often", input: "often", want: false},
		{name: "garden", input: "garden", want: false},
		{name: "kitchen", input: "kitchen", want: false},

		// Edge cases
		{name: "empty", input: "", want: false},
		{name: "a", input: "a", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.IsParticiple(tt.input)
			if got != tt.want {
				t.Errorf("IsParticiple(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func BenchmarkPastParticiple(b *testing.B) {
	inputs := []string{"walk", "stop", "try", "go", "take", "run"}
	for i := range b.N {
		inflect.PastParticiple(inputs[i%len(inputs)])
	}
}

func BenchmarkIsParticiple(b *testing.B) {
	inputs := []string{"running", "walked", "taken", "walk", "cat"}
	for i := range b.N {
		inflect.IsParticiple(inputs[i%len(inputs)])
	}
}
