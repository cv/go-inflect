package inflect_test

import (
	"testing"

	inflect "github.com/cv/go-inflect"
)

func TestComparative(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		// Empty string
		{name: "empty", input: "", want: ""},

		// Irregular forms
		{name: "good", input: "good", want: "better"},
		{name: "bad", input: "bad", want: "worse"},
		{name: "far", input: "far", want: "farther"},
		{name: "little", input: "little", want: "less"},
		{name: "much", input: "much", want: "more"},
		{name: "many", input: "many", want: "more"},
		{name: "old irregular", input: "old", want: "older"}, // also has "elder" but older is more common

		// One-syllable adjectives: add -er
		{name: "tall", input: "tall", want: "taller"},
		{name: "short", input: "short", want: "shorter"},
		{name: "fast", input: "fast", want: "faster"},
		{name: "slow", input: "slow", want: "slower"},
		{name: "young", input: "young", want: "younger"},
		{name: "long", input: "long", want: "longer"},
		{name: "strong", input: "strong", want: "stronger"},
		{name: "weak", input: "weak", want: "weaker"},
		{name: "cheap", input: "cheap", want: "cheaper"},
		{name: "deep", input: "deep", want: "deeper"},
		{name: "high", input: "high", want: "higher"},
		{name: "low", input: "low", want: "lower"},
		{name: "new", input: "new", want: "newer"},
		{name: "poor", input: "poor", want: "poorer"},
		{name: "rich", input: "rich", want: "richer"},
		{name: "warm", input: "warm", want: "warmer"},
		{name: "cold", input: "cold", want: "colder"},
		{name: "dark", input: "dark", want: "darker"},
		{name: "light", input: "light", want: "lighter"},
		{name: "hard", input: "hard", want: "harder"},
		{name: "soft", input: "soft", want: "softer"},
		{name: "clean", input: "clean", want: "cleaner"},
		{name: "loud", input: "loud", want: "louder"},

		// One-syllable ending in -e: add -r
		{name: "large", input: "large", want: "larger"},
		{name: "wide", input: "wide", want: "wider"},
		{name: "close", input: "close", want: "closer"},
		{name: "late", input: "late", want: "later"},
		{name: "nice", input: "nice", want: "nicer"},
		{name: "safe", input: "safe", want: "safer"},
		{name: "wise", input: "wise", want: "wiser"},
		{name: "rude", input: "rude", want: "ruder"},
		{name: "rare", input: "rare", want: "rarer"},
		{name: "pale", input: "pale", want: "paler"},
		{name: "fine", input: "fine", want: "finer"},
		{name: "cute", input: "cute", want: "cuter"},
		{name: "pure", input: "pure", want: "purer"},

		// CVC pattern (consonant-vowel-consonant): double final consonant
		{name: "big", input: "big", want: "bigger"},
		{name: "hot", input: "hot", want: "hotter"},
		{name: "thin", input: "thin", want: "thinner"},
		{name: "fat", input: "fat", want: "fatter"},
		{name: "wet", input: "wet", want: "wetter"},
		{name: "sad", input: "sad", want: "sadder"},
		{name: "red", input: "red", want: "redder"},
		{name: "dim", input: "dim", want: "dimmer"},
		{name: "fit", input: "fit", want: "fitter"},

		// Consonant + y: change y to -ier
		{name: "happy", input: "happy", want: "happier"},
		{name: "easy", input: "easy", want: "easier"},
		{name: "busy", input: "busy", want: "busier"},
		{name: "funny", input: "funny", want: "funnier"},
		{name: "pretty", input: "pretty", want: "prettier"},
		{name: "heavy", input: "heavy", want: "heavier"},
		{name: "dirty", input: "dirty", want: "dirtier"},
		{name: "angry", input: "angry", want: "angrier"},
		{name: "crazy", input: "crazy", want: "crazier"},
		{name: "lazy", input: "lazy", want: "lazier"},
		{name: "tiny", input: "tiny", want: "tinier"},
		{name: "ugly", input: "ugly", want: "uglier"},
		{name: "early", input: "early", want: "earlier"},
		{name: "noisy", input: "noisy", want: "noisier"},

		// Two-syllable adjectives that take -er (common ones)
		{name: "simple", input: "simple", want: "simpler"},
		{name: "gentle", input: "gentle", want: "gentler"},
		{name: "narrow", input: "narrow", want: "narrower"},
		{name: "shallow", input: "shallow", want: "shallower"},
		{name: "quiet", input: "quiet", want: "quieter"},
		{name: "clever", input: "clever", want: "cleverer"},

		// Long adjectives: use "more"
		{name: "beautiful", input: "beautiful", want: "more beautiful"},
		{name: "dangerous", input: "dangerous", want: "more dangerous"},
		{name: "expensive", input: "expensive", want: "more expensive"},
		{name: "important", input: "important", want: "more important"},
		{name: "interesting", input: "interesting", want: "more interesting"},
		{name: "comfortable", input: "comfortable", want: "more comfortable"},
		{name: "difficult", input: "difficult", want: "more difficult"},
		{name: "intelligent", input: "intelligent", want: "more intelligent"},
		{name: "wonderful", input: "wonderful", want: "more wonderful"},
		{name: "terrible", input: "terrible", want: "more terrible"},
		{name: "horrible", input: "horrible", want: "more horrible"},
		{name: "incredible", input: "incredible", want: "more incredible"},
		{name: "successful", input: "successful", want: "more successful"},
		{name: "popular", input: "popular", want: "more popular"},
		{name: "famous", input: "famous", want: "more famous"},
		{name: "nervous", input: "nervous", want: "more nervous"},

		// Case preservation
		{name: "BIG uppercase", input: "BIG", want: "BIGGER"},
		{name: "Big titlecase", input: "Big", want: "Bigger"},
		{name: "GOOD uppercase", input: "GOOD", want: "BETTER"},
		{name: "Good titlecase", input: "Good", want: "Better"},
		{name: "BEAUTIFUL uppercase", input: "BEAUTIFUL", want: "MORE BEAUTIFUL"},
		{name: "Beautiful titlecase", input: "Beautiful", want: "More Beautiful"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Comparative(tt.input)
			if got != tt.want {
				t.Errorf("Comparative(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestSuperlative(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		// Empty string
		{name: "empty", input: "", want: ""},

		// Irregular forms
		{name: "good", input: "good", want: "best"},
		{name: "bad", input: "bad", want: "worst"},
		{name: "far", input: "far", want: "farthest"},
		{name: "little", input: "little", want: "least"},
		{name: "much", input: "much", want: "most"},
		{name: "many", input: "many", want: "most"},
		{name: "old irregular", input: "old", want: "oldest"}, // also has "eldest"

		// One-syllable adjectives: add -est
		{name: "tall", input: "tall", want: "tallest"},
		{name: "short", input: "short", want: "shortest"},
		{name: "fast", input: "fast", want: "fastest"},
		{name: "slow", input: "slow", want: "slowest"},
		{name: "young", input: "young", want: "youngest"},
		{name: "long", input: "long", want: "longest"},
		{name: "strong", input: "strong", want: "strongest"},
		{name: "weak", input: "weak", want: "weakest"},
		{name: "cheap", input: "cheap", want: "cheapest"},
		{name: "deep", input: "deep", want: "deepest"},
		{name: "high", input: "high", want: "highest"},
		{name: "low", input: "low", want: "lowest"},
		{name: "new", input: "new", want: "newest"},
		{name: "poor", input: "poor", want: "poorest"},
		{name: "rich", input: "rich", want: "richest"},
		{name: "warm", input: "warm", want: "warmest"},
		{name: "cold", input: "cold", want: "coldest"},
		{name: "dark", input: "dark", want: "darkest"},
		{name: "light", input: "light", want: "lightest"},
		{name: "hard", input: "hard", want: "hardest"},
		{name: "soft", input: "soft", want: "softest"},
		{name: "clean", input: "clean", want: "cleanest"},
		{name: "loud", input: "loud", want: "loudest"},

		// One-syllable ending in -e: add -st
		{name: "large", input: "large", want: "largest"},
		{name: "wide", input: "wide", want: "widest"},
		{name: "close", input: "close", want: "closest"},
		{name: "late", input: "late", want: "latest"},
		{name: "nice", input: "nice", want: "nicest"},
		{name: "safe", input: "safe", want: "safest"},
		{name: "wise", input: "wise", want: "wisest"},
		{name: "rude", input: "rude", want: "rudest"},
		{name: "rare", input: "rare", want: "rarest"},
		{name: "pale", input: "pale", want: "palest"},
		{name: "fine", input: "fine", want: "finest"},
		{name: "cute", input: "cute", want: "cutest"},
		{name: "pure", input: "pure", want: "purest"},

		// CVC pattern: double final consonant
		{name: "big", input: "big", want: "biggest"},
		{name: "hot", input: "hot", want: "hottest"},
		{name: "thin", input: "thin", want: "thinnest"},
		{name: "fat", input: "fat", want: "fattest"},
		{name: "wet", input: "wet", want: "wettest"},
		{name: "sad", input: "sad", want: "saddest"},
		{name: "red", input: "red", want: "reddest"},
		{name: "dim", input: "dim", want: "dimmest"},
		{name: "fit", input: "fit", want: "fittest"},

		// Consonant + y: change y to -iest
		{name: "happy", input: "happy", want: "happiest"},
		{name: "easy", input: "easy", want: "easiest"},
		{name: "busy", input: "busy", want: "busiest"},
		{name: "funny", input: "funny", want: "funniest"},
		{name: "pretty", input: "pretty", want: "prettiest"},
		{name: "heavy", input: "heavy", want: "heaviest"},
		{name: "dirty", input: "dirty", want: "dirtiest"},
		{name: "angry", input: "angry", want: "angriest"},
		{name: "crazy", input: "crazy", want: "craziest"},
		{name: "lazy", input: "lazy", want: "laziest"},
		{name: "tiny", input: "tiny", want: "tiniest"},
		{name: "ugly", input: "ugly", want: "ugliest"},
		{name: "early", input: "early", want: "earliest"},
		{name: "noisy", input: "noisy", want: "noisiest"},

		// Two-syllable adjectives that take -est
		{name: "simple", input: "simple", want: "simplest"},
		{name: "gentle", input: "gentle", want: "gentlest"},
		{name: "narrow", input: "narrow", want: "narrowest"},
		{name: "shallow", input: "shallow", want: "shallowest"},
		{name: "quiet", input: "quiet", want: "quietest"},
		{name: "clever", input: "clever", want: "cleverest"},

		// Long adjectives: use "most"
		{name: "beautiful", input: "beautiful", want: "most beautiful"},
		{name: "dangerous", input: "dangerous", want: "most dangerous"},
		{name: "expensive", input: "expensive", want: "most expensive"},
		{name: "important", input: "important", want: "most important"},
		{name: "interesting", input: "interesting", want: "most interesting"},
		{name: "comfortable", input: "comfortable", want: "most comfortable"},
		{name: "difficult", input: "difficult", want: "most difficult"},
		{name: "intelligent", input: "intelligent", want: "most intelligent"},
		{name: "wonderful", input: "wonderful", want: "most wonderful"},
		{name: "terrible", input: "terrible", want: "most terrible"},
		{name: "horrible", input: "horrible", want: "most horrible"},
		{name: "incredible", input: "incredible", want: "most incredible"},
		{name: "successful", input: "successful", want: "most successful"},
		{name: "popular", input: "popular", want: "most popular"},
		{name: "famous", input: "famous", want: "most famous"},
		{name: "nervous", input: "nervous", want: "most nervous"},

		// Case preservation
		{name: "BIG uppercase", input: "BIG", want: "BIGGEST"},
		{name: "Big titlecase", input: "Big", want: "Biggest"},
		{name: "GOOD uppercase", input: "GOOD", want: "BEST"},
		{name: "Good titlecase", input: "Good", want: "Best"},
		{name: "BEAUTIFUL uppercase", input: "BEAUTIFUL", want: "MOST BEAUTIFUL"},
		{name: "Beautiful titlecase", input: "Beautiful", want: "Most Beautiful"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Superlative(tt.input)
			if got != tt.want {
				t.Errorf("Superlative(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func BenchmarkComparative(b *testing.B) {
	benchmarks := []struct {
		name  string
		input string
	}{
		{"irregular", "good"},
		{"short", "tall"},
		{"ending_e", "large"},
		{"cvc_double", "big"},
		{"y_to_ier", "happy"},
		{"long_more", "beautiful"},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for range b.N {
				inflect.Comparative(bm.input)
			}
		})
	}
}

func BenchmarkSuperlative(b *testing.B) {
	benchmarks := []struct {
		name  string
		input string
	}{
		{"irregular", "good"},
		{"short", "tall"},
		{"ending_e", "large"},
		{"cvc_double", "big"},
		{"y_to_iest", "happy"},
		{"long_most", "beautiful"},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for range b.N {
				inflect.Superlative(bm.input)
			}
		})
	}
}
