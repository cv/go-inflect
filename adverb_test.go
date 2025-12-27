package inflect_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	inflect "github.com/cv/go-inflect"
)

func TestAdverb(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		// Empty string
		{name: "empty", input: "", want: ""},

		// Irregular forms
		{name: "good", input: "good", want: "well"},
		{name: "fast", input: "fast", want: "fast"},
		{name: "hard", input: "hard", want: "hard"},
		{name: "late", input: "late", want: "late"},
		{name: "early", input: "early", want: "early"},
		{name: "straight", input: "straight", want: "straight"},

		// Most adjectives: add -ly
		{name: "quick", input: "quick", want: "quickly"},
		{name: "slow", input: "slow", want: "slowly"},
		{name: "soft", input: "soft", want: "softly"},
		{name: "loud", input: "loud", want: "loudly"},
		{name: "quiet", input: "quiet", want: "quietly"},
		{name: "bright", input: "bright", want: "brightly"},
		{name: "calm", input: "calm", want: "calmly"},
		{name: "clear", input: "clear", want: "clearly"},
		{name: "clever", input: "clever", want: "cleverly"},
		{name: "close", input: "close", want: "closely"},
		{name: "cold", input: "cold", want: "coldly"},
		{name: "cool", input: "cool", want: "coolly"},
		{name: "correct", input: "correct", want: "correctly"},
		{name: "deep", input: "deep", want: "deeply"},
		{name: "direct", input: "direct", want: "directly"},
		{name: "exact", input: "exact", want: "exactly"},
		{name: "fair", input: "fair", want: "fairly"},
		{name: "firm", input: "firm", want: "firmly"},
		{name: "free", input: "free", want: "freely"},
		{name: "great", input: "great", want: "greatly"},
		{name: "kind", input: "kind", want: "kindly"},
		{name: "neat", input: "neat", want: "neatly"},
		{name: "nice", input: "nice", want: "nicely"},
		{name: "open", input: "open", want: "openly"},
		{name: "perfect", input: "perfect", want: "perfectly"},
		{name: "plain", input: "plain", want: "plainly"},
		{name: "poor", input: "poor", want: "poorly"},
		{name: "proud", input: "proud", want: "proudly"},
		{name: "pure", input: "pure", want: "purely"},
		{name: "rare", input: "rare", want: "rarely"},
		{name: "real", input: "real", want: "really"},
		{name: "rough", input: "rough", want: "roughly"},
		{name: "safe", input: "safe", want: "safely"},
		{name: "secret", input: "secret", want: "secretly"},
		{name: "sharp", input: "sharp", want: "sharply"},
		{name: "short", input: "short", want: "shortly"},
		{name: "silent", input: "silent", want: "silently"},
		{name: "sincere", input: "sincere", want: "sincerely"},
		{name: "smooth", input: "smooth", want: "smoothly"},
		{name: "stern", input: "stern", want: "sternly"},
		{name: "strict", input: "strict", want: "strictly"},
		{name: "strong", input: "strong", want: "strongly"},
		{name: "sudden", input: "sudden", want: "suddenly"},
		{name: "sweet", input: "sweet", want: "sweetly"},
		{name: "tight", input: "tight", want: "tightly"},
		{name: "usual", input: "usual", want: "usually"},
		{name: "warm", input: "warm", want: "warmly"},
		{name: "weak", input: "weak", want: "weakly"},
		{name: "wide", input: "wide", want: "widely"},
		{name: "wrong", input: "wrong", want: "wrongly"},

		// Adjectives ending in -y: change y to -ily
		{name: "happy", input: "happy", want: "happily"},
		{name: "easy", input: "easy", want: "easily"},
		{name: "angry", input: "angry", want: "angrily"},
		{name: "busy", input: "busy", want: "busily"},
		{name: "crazy", input: "crazy", want: "crazily"},
		{name: "dirty", input: "dirty", want: "dirtily"},
		{name: "hungry", input: "hungry", want: "hungrily"},
		{name: "lazy", input: "lazy", want: "lazily"},
		{name: "lucky", input: "lucky", want: "luckily"},
		{name: "merry", input: "merry", want: "merrily"},
		{name: "messy", input: "messy", want: "messily"},
		{name: "noisy", input: "noisy", want: "noisily"},
		{name: "pretty", input: "pretty", want: "prettily"},
		{name: "ready", input: "ready", want: "readily"},
		{name: "steady", input: "steady", want: "steadily"},
		{name: "thirsty", input: "thirsty", want: "thirstily"},
		{name: "weary", input: "weary", want: "wearily"},

		// Adjectives ending in -le: change -le to -ly
		{name: "gentle", input: "gentle", want: "gently"},
		{name: "simple", input: "simple", want: "simply"},
		{name: "humble", input: "humble", want: "humbly"},
		{name: "possible", input: "possible", want: "possibly"},
		{name: "probable", input: "probable", want: "probably"},
		{name: "terrible", input: "terrible", want: "terribly"},
		{name: "horrible", input: "horrible", want: "horribly"},
		{name: "incredible", input: "incredible", want: "incredibly"},
		{name: "comfortable", input: "comfortable", want: "comfortably"},
		{name: "reasonable", input: "reasonable", want: "reasonably"},
		{name: "responsible", input: "responsible", want: "responsibly"},
		{name: "visible", input: "visible", want: "visibly"},
		{name: "noble", input: "noble", want: "nobly"},
		{name: "able", input: "able", want: "ably"},
		{name: "subtle", input: "subtle", want: "subtly"},
		{name: "feeble", input: "feeble", want: "feebly"},

		// Adjectives ending in -ue: drop e, add -ly
		{name: "true", input: "true", want: "truly"},
		{name: "due", input: "due", want: "duly"},

		// Adjectives ending in -ic: add -ally
		{name: "basic", input: "basic", want: "basically"},
		{name: "automatic", input: "automatic", want: "automatically"},
		{name: "dramatic", input: "dramatic", want: "dramatically"},
		{name: "enthusiastic", input: "enthusiastic", want: "enthusiastically"},
		{name: "fantastic", input: "fantastic", want: "fantastically"},
		{name: "frantic", input: "frantic", want: "frantically"},
		{name: "historic", input: "historic", want: "historically"},
		{name: "magic", input: "magic", want: "magically"},
		{name: "romantic", input: "romantic", want: "romantically"},
		{name: "scientific", input: "scientific", want: "scientifically"},
		{name: "specific", input: "specific", want: "specifically"},
		{name: "systematic", input: "systematic", want: "systematically"},
		{name: "tragic", input: "tragic", want: "tragically"},

		// Exception: public -> publicly (not publically)
		{name: "public", input: "public", want: "publicly"},

		// Adjectives ending in -ll: add -y (not -ly)
		{name: "full", input: "full", want: "fully"},
		{name: "dull", input: "dull", want: "dully"},

		// Case preservation
		{name: "QUICK uppercase", input: "QUICK", want: "QUICKLY"},
		{name: "Quick titlecase", input: "Quick", want: "Quickly"},
		{name: "GOOD uppercase", input: "GOOD", want: "WELL"},
		{name: "Good titlecase", input: "Good", want: "Well"},
		{name: "HAPPY uppercase", input: "HAPPY", want: "HAPPILY"},
		{name: "Happy titlecase", input: "Happy", want: "Happily"},
		{name: "GENTLE uppercase", input: "GENTLE", want: "GENTLY"},
		{name: "Gentle titlecase", input: "Gentle", want: "Gently"},
		{name: "BASIC uppercase", input: "BASIC", want: "BASICALLY"},
		{name: "Basic titlecase", input: "Basic", want: "Basically"},
		{name: "TRUE uppercase", input: "TRUE", want: "TRULY"},
		{name: "True titlecase", input: "True", want: "Truly"},
		{name: "FAST uppercase", input: "FAST", want: "FAST"},
		{name: "Fast titlecase", input: "Fast", want: "Fast"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Adverb(tt.input)
			assert.Equal(t, tt.want, got, "Adverb(%q)", tt.input)
		})
	}
}

func BenchmarkAdverb(b *testing.B) {
	benchmarks := []struct {
		name  string
		input string
	}{
		{"irregular_good", "good"},
		{"irregular_unchanged", "fast"},
		{"regular_ly", "quick"},
		{"ending_y", "happy"},
		{"ending_le", "gentle"},
		{"ending_ue", "true"},
		{"ending_ic", "basic"},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for range b.N {
				inflect.Adverb(bm.input)
			}
		})
	}
}
