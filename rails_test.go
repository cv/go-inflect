package inflect_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	inflect "github.com/cv/go-inflect"
)

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
