package inflect_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	inflect "github.com/cv/go-inflect"
)

func TestFutureTense(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		// Empty string
		{name: "empty", input: "", want: ""},

		// Regular verbs
		{name: "walk", input: "walk", want: "will walk"},
		{name: "talk", input: "talk", want: "will talk"},
		{name: "run", input: "run", want: "will run"},
		{name: "go", input: "go", want: "will go"},
		{name: "be", input: "be", want: "will be"},
		{name: "have", input: "have", want: "will have"},
		{name: "do", input: "do", want: "will do"},
		{name: "make", input: "make", want: "will make"},
		{name: "say", input: "say", want: "will say"},
		{name: "get", input: "get", want: "will get"},
		{name: "think", input: "think", want: "will think"},
		{name: "come", input: "come", want: "will come"},
		{name: "see", input: "see", want: "will see"},
		{name: "know", input: "know", want: "will know"},
		{name: "take", input: "take", want: "will take"},
		{name: "find", input: "find", want: "will find"},
		{name: "give", input: "give", want: "will give"},
		{name: "tell", input: "tell", want: "will tell"},
		{name: "work", input: "work", want: "will work"},
		{name: "play", input: "play", want: "will play"},
		{name: "try", input: "try", want: "will try"},
		{name: "stop", input: "stop", want: "will stop"},

		// Case preservation
		{name: "WALK uppercase", input: "WALK", want: "WILL WALK"},
		{name: "Walk titlecase", input: "Walk", want: "Will Walk"},
		{name: "GO uppercase", input: "GO", want: "WILL GO"},
		{name: "Go titlecase", input: "Go", want: "Will Go"},
		{name: "RUN uppercase", input: "RUN", want: "WILL RUN"},
		{name: "Run titlecase", input: "Run", want: "Will Run"},
		{name: "BE uppercase", input: "BE", want: "WILL BE"},
		{name: "Be titlecase", input: "Be", want: "Will Be"},

		// Mixed case (first letter determines case)
		{name: "wAlK mixed", input: "wAlK", want: "will wAlK"},
		{name: "PLAY uppercase", input: "PLAY", want: "WILL PLAY"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.FutureTense(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func BenchmarkFutureTense(b *testing.B) {
	benchmarks := []struct {
		name  string
		input string
	}{
		{"regular", "walk"},
		{"irregular", "go"},
		{"uppercase", "WALK"},
		{"titlecase", "Walk"},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for range b.N {
				inflect.FutureTense(bm.input)
			}
		})
	}
}
