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
