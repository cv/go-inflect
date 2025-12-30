package inflect_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	inflect "github.com/cv/go-inflect/v2"
)

// Benchmarks for Rails-style helper functions

func BenchmarkHumanize(b *testing.B) {
	benchmarks := []struct {
		name  string
		input string
	}{
		{"snake_case", "employee_salary"},
		{"with_id", "author_id"},
		{"camelCase", "XMLParser"},
		{"long", "hello_world_foo_bar"},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for range b.N {
				inflect.Humanize(bm.input)
			}
		})
	}
}

func BenchmarkForeignKey(b *testing.B) {
	benchmarks := []struct {
		name  string
		input string
	}{
		{"simple", "Person"},
		{"camelCase", "AdminUser"},
		{"acronym", "XMLParser"},
		{"lowercase", "user"},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for range b.N {
				inflect.ForeignKey(bm.input)
			}
		})
	}
}

func BenchmarkForeignKeyCondensed(b *testing.B) {
	benchmarks := []struct {
		name  string
		input string
	}{
		{"simple", "Person"},
		{"camelCase", "AdminUser"},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for range b.N {
				inflect.ForeignKeyCondensed(bm.input)
			}
		})
	}
}

func BenchmarkTableize(b *testing.B) {
	benchmarks := []struct {
		name  string
		input string
	}{
		{"simple", "Person"},
		{"camelCase", "RawScaledScorer"},
		{"compound", "MouseTrap"},
		{"irregular", "Child"},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for range b.N {
				inflect.Tableize(bm.input)
			}
		})
	}
}

func BenchmarkParameterize(b *testing.B) {
	benchmarks := []struct {
		name  string
		input string
	}{
		{"simple", "Hello World"},
		{"special_chars", "Special!@#$%Characters"},
		{"unicode", "café au lait"},
		{"spaces", "Multiple   Spaces"},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for range b.N {
				inflect.Parameterize(bm.input)
			}
		})
	}
}

func BenchmarkParameterizeJoin(b *testing.B) {
	for b.Loop() {
		inflect.ParameterizeJoin("Hello World!", "_")
	}
}

func BenchmarkTypeify(b *testing.B) {
	benchmarks := []struct {
		name  string
		input string
	}{
		{"simple", "users"},
		{"snake_case", "raw_scaled_scorers"},
		{"irregular", "people"},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for range b.N {
				inflect.Typeify(bm.input)
			}
		})
	}
}

func BenchmarkAsciify(b *testing.B) {
	benchmarks := []struct {
		name  string
		input string
	}{
		{"accented", "café"},
		{"multiple", "Crème brûlée"},
		{"ascii_only", "hello"},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for range b.N {
				inflect.Asciify(bm.input)
			}
		})
	}
}

// Benchmarks for compatibility aliases

func BenchmarkPluralizeFn(b *testing.B) {
	benchmarks := []struct {
		name  string
		input string
	}{
		{"regular", "cat"},
		{"irregular", "child"},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for range b.N {
				inflect.Pluralize(bm.input)
			}
		})
	}
}

func BenchmarkSingularizeFn(b *testing.B) {
	benchmarks := []struct {
		name  string
		input string
	}{
		{"regular", "cats"},
		{"irregular", "children"},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for range b.N {
				inflect.Singularize(bm.input)
			}
		})
	}
}

func BenchmarkCamelize(b *testing.B) {
	benchmarks := []struct {
		name  string
		input string
	}{
		{"snake_case", "hello_world"},
		{"kebab-case", "foo-bar"},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for range b.N {
				inflect.Camelize(bm.input)
			}
		})
	}
}

func BenchmarkCamelizeDownFirst(b *testing.B) {
	benchmarks := []struct {
		name  string
		input string
	}{
		{"snake_case", "hello_world"},
		{"kebab-case", "foo-bar"},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for range b.N {
				inflect.CamelizeDownFirst(bm.input)
			}
		})
	}
}

func TestHumanize(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"employee_salary", "Employee salary"},
		{"author_id", "Author"},
		{"author_ID", "Author"},
		{"authorID", "Author"},
		{"hello-world", "Hello world"},
		{"XMLParser", "Xml parser"},
		{"user_name", "User name"},
		{"firstName", "First name"},
		{"", ""},
		{"a", "A"},
		{"already humanized", "Already humanized"},
		{"multiple__underscores", "Multiple underscores"},
		{"trailing_id", "Trailing"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := inflect.Humanize(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestForeignKey(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"Person", "person_id"},
		{"Message", "message_id"},
		{"AdminUser", "admin_user_id"},
		{"user", "user_id"},
		{"XMLParser", "xml_parser_id"},
		{"", "_id"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := inflect.ForeignKey(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestForeignKeyCondensed(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"Person", "personid"},
		{"Message", "messageid"},
		{"AdminUser", "admin_userid"},
		{"user", "userid"},
		{"", "id"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := inflect.ForeignKeyCondensed(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestTableize(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"Person", "people"},
		{"RawScaledScorer", "raw_scaled_scorers"},
		{"MouseTrap", "mouse_traps"},
		{"User", "users"},
		{"admin_user", "admin_users"},
		{"Child", "children"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := inflect.Tableize(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestParameterize(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"Hello World!", "hello-world"},
		{"Hello, World!", "hello-world"},
		{"  Multiple   Spaces  ", "multiple-spaces"},
		{"Special!@#$%Characters", "specialcharacters"},
		{"Already-Dashed", "already-dashed"},
		{"under_scored", "under-scored"},
		{"MixedCase", "mixedcase"},
		{"", ""},
		{"café au lait", "cafe-au-lait"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := inflect.Parameterize(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestParameterizeJoin(t *testing.T) {
	tests := []struct {
		input string
		sep   string
		want  string
	}{
		{"Hello World!", "_", "hello_world"},
		{"Hello World!", "-", "hello-world"},
		{"Hello World!", "", "helloworld"},
		{"Multiple   Spaces", "_", "multiple_spaces"},
	}

	for _, tt := range tests {
		t.Run(tt.input+"_"+tt.sep, func(t *testing.T) {
			got := inflect.ParameterizeJoin(tt.input, tt.sep)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestTypeify(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"users", "User"},
		{"raw_scaled_scorers", "RawScaledScorer"},
		{"people", "Person"},
		{"mice", "Mouse"},
		{"admin_users", "AdminUser"},
		{"categories", "Category"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := inflect.Typeify(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestAsciify(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"café", "cafe"},
		{"naïve", "naive"},
		{"résumé", "resume"},
		{"日本語", ""},
		{"hello", "hello"},
		{"über", "uber"},
		{"Ångström", "Angstrom"},
		{"", ""},
		{"Crème brûlée", "Creme brulee"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := inflect.Asciify(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
