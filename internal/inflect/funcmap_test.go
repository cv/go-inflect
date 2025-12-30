package inflect_test

import (
	"bytes"
	htmltemplate "html/template"
	"testing"
	texttemplate "text/template"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cv/go-inflect/internal/inflect"
)

func TestFuncMap(t *testing.T) {
	fm := inflect.FuncMap()

	// Verify all expected functions are present
	expectedFuncs := []string{
		"plural", "singular", "pluralNoun", "pluralVerb", "pluralAdj", "singularNoun",
		"an", "a",
		"ordinal", "ordinalWord", "numberToWords",
		"pastTense", "presentParticiple",
		"comparative", "superlative",
		"possessive",
		"join", "joinWith",
		"camelCase", "snakeCase", "kebabCase", "pascalCase",
	}

	for _, name := range expectedFuncs {
		assert.Contains(t, fm, name, "FuncMap should contain %q", name)
	}
}

func TestEngineFuncMap(t *testing.T) {
	e := inflect.NewEngine()
	fm := e.FuncMap()

	// Should have same functions as package-level FuncMap
	assert.Contains(t, fm, "plural")
	assert.Contains(t, fm, "an")
	assert.Contains(t, fm, "ordinal")
}

func TestFuncMapPlural(t *testing.T) {
	tests := []struct {
		name     string
		template string
		data     any
		want     string
	}{
		{
			name:     "basic plural",
			template: `{{plural "cat"}}`,
			data:     nil,
			want:     "cats",
		},
		{
			name:     "plural with count 2",
			template: `{{plural "cat" .Count}}`,
			data:     map[string]int{"Count": 2},
			want:     "cats",
		},
		{
			name:     "plural with count 1 returns singular",
			template: `{{plural "cat" .Count}}`,
			data:     map[string]int{"Count": 1},
			want:     "cat",
		},
		{
			name:     "plural with count 0",
			template: `{{plural "cat" .Count}}`,
			data:     map[string]int{"Count": 0},
			want:     "cats",
		},
		{
			name:     "irregular plural",
			template: `{{plural "child"}}`,
			data:     nil,
			want:     "children",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpl, err := texttemplate.New("test").Funcs(inflect.FuncMap()).Parse(tt.template)
			require.NoError(t, err)

			var buf bytes.Buffer
			err = tmpl.Execute(&buf, tt.data)
			require.NoError(t, err)
			assert.Equal(t, tt.want, buf.String())
		})
	}
}

func TestFuncMapSingular(t *testing.T) {
	tests := []struct {
		name     string
		template string
		want     string
	}{
		{name: "basic singular", template: `{{singular "cats"}}`, want: "cat"},
		{name: "irregular singular", template: `{{singular "children"}}`, want: "child"},
		{name: "already singular", template: `{{singular "cat"}}`, want: "cat"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpl, err := texttemplate.New("test").Funcs(inflect.FuncMap()).Parse(tt.template)
			require.NoError(t, err)

			var buf bytes.Buffer
			err = tmpl.Execute(&buf, nil)
			require.NoError(t, err)
			assert.Equal(t, tt.want, buf.String())
		})
	}
}

func TestFuncMapArticles(t *testing.T) {
	tests := []struct {
		name     string
		template string
		want     string
	}{
		{name: "an consonant", template: `{{an "cat"}}`, want: "a cat"},
		{name: "an vowel", template: `{{an "apple"}}`, want: "an apple"},
		{name: "a consonant", template: `{{a "cat"}}`, want: "a cat"},
		{name: "a vowel", template: `{{a "apple"}}`, want: "an apple"},
		{name: "an hour", template: `{{an "hour"}}`, want: "an hour"},
		{name: "an university", template: `{{an "university"}}`, want: "a university"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpl, err := texttemplate.New("test").Funcs(inflect.FuncMap()).Parse(tt.template)
			require.NoError(t, err)

			var buf bytes.Buffer
			err = tmpl.Execute(&buf, nil)
			require.NoError(t, err)
			assert.Equal(t, tt.want, buf.String())
		})
	}
}

func TestFuncMapOrdinals(t *testing.T) {
	tests := []struct {
		name     string
		template string
		want     string
	}{
		{name: "ordinal 1", template: `{{ordinal 1}}`, want: "1st"},
		{name: "ordinal 2", template: `{{ordinal 2}}`, want: "2nd"},
		{name: "ordinal 3", template: `{{ordinal 3}}`, want: "3rd"},
		{name: "ordinal 11", template: `{{ordinal 11}}`, want: "11th"},
		{name: "ordinal 21", template: `{{ordinal 21}}`, want: "21st"},
		{name: "ordinalWord 1", template: `{{ordinalWord 1}}`, want: "first"},
		{name: "ordinalWord 21", template: `{{ordinalWord 21}}`, want: "twenty-first"},
		{name: "numberToWords 42", template: `{{numberToWords 42}}`, want: "forty-two"},
		{name: "numberToWords 100", template: `{{numberToWords 100}}`, want: "one hundred"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpl, err := texttemplate.New("test").Funcs(inflect.FuncMap()).Parse(tt.template)
			require.NoError(t, err)

			var buf bytes.Buffer
			err = tmpl.Execute(&buf, nil)
			require.NoError(t, err)
			assert.Equal(t, tt.want, buf.String())
		})
	}
}

func TestFuncMapVerbTenses(t *testing.T) {
	tests := []struct {
		name     string
		template string
		want     string
	}{
		{name: "pastTense walk", template: `{{pastTense "walk"}}`, want: "walked"},
		{name: "pastTense go", template: `{{pastTense "go"}}`, want: "went"},
		{name: "pastTense try", template: `{{pastTense "try"}}`, want: "tried"},
		{name: "presentParticiple run", template: `{{presentParticiple "run"}}`, want: "running"},
		{name: "presentParticiple make", template: `{{presentParticiple "make"}}`, want: "making"},
		{name: "presentParticiple play", template: `{{presentParticiple "play"}}`, want: "playing"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpl, err := texttemplate.New("test").Funcs(inflect.FuncMap()).Parse(tt.template)
			require.NoError(t, err)

			var buf bytes.Buffer
			err = tmpl.Execute(&buf, nil)
			require.NoError(t, err)
			assert.Equal(t, tt.want, buf.String())
		})
	}
}

func TestFuncMapComparison(t *testing.T) {
	tests := []struct {
		name     string
		template string
		want     string
	}{
		{name: "comparative big", template: `{{comparative "big"}}`, want: "bigger"},
		{name: "comparative good", template: `{{comparative "good"}}`, want: "better"},
		{name: "comparative beautiful", template: `{{comparative "beautiful"}}`, want: "more beautiful"},
		{name: "superlative big", template: `{{superlative "big"}}`, want: "biggest"},
		{name: "superlative good", template: `{{superlative "good"}}`, want: "best"},
		{name: "superlative beautiful", template: `{{superlative "beautiful"}}`, want: "most beautiful"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpl, err := texttemplate.New("test").Funcs(inflect.FuncMap()).Parse(tt.template)
			require.NoError(t, err)

			var buf bytes.Buffer
			err = tmpl.Execute(&buf, nil)
			require.NoError(t, err)
			assert.Equal(t, tt.want, buf.String())
		})
	}
}

func TestFuncMapPossessive(t *testing.T) {
	tests := []struct {
		name     string
		template string
		want     string
	}{
		{name: "possessive cat", template: `{{possessive "cat"}}`, want: "cat's"},
		{name: "possessive James", template: `{{possessive "James"}}`, want: "James's"},
		{name: "possessive children", template: `{{possessive "children"}}`, want: "children's"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpl, err := texttemplate.New("test").Funcs(inflect.FuncMap()).Parse(tt.template)
			require.NoError(t, err)

			var buf bytes.Buffer
			err = tmpl.Execute(&buf, nil)
			require.NoError(t, err)
			assert.Equal(t, tt.want, buf.String())
		})
	}
}

func TestFuncMapJoin(t *testing.T) {
	tests := []struct {
		name     string
		template string
		data     any
		want     string
	}{
		{
			name:     "join empty",
			template: `{{join .Items}}`,
			data:     map[string][]string{"Items": {}},
			want:     "",
		},
		{
			name:     "join single",
			template: `{{join .Items}}`,
			data:     map[string][]string{"Items": {"apple"}},
			want:     "apple",
		},
		{
			name:     "join two",
			template: `{{join .Items}}`,
			data:     map[string][]string{"Items": {"apple", "banana"}},
			want:     "apple and banana",
		},
		{
			name:     "join three with Oxford comma",
			template: `{{join .Items}}`,
			data:     map[string][]string{"Items": {"apple", "banana", "cherry"}},
			want:     "apple, banana, and cherry",
		},
		{
			name:     "joinWith or",
			template: `{{joinWith .Items "or"}}`,
			data:     map[string][]string{"Items": {"red", "blue", "green"}},
			want:     "red, blue, or green",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpl, err := texttemplate.New("test").Funcs(inflect.FuncMap()).Parse(tt.template)
			require.NoError(t, err)

			var buf bytes.Buffer
			err = tmpl.Execute(&buf, tt.data)
			require.NoError(t, err)
			assert.Equal(t, tt.want, buf.String())
		})
	}
}

func TestFuncMapCaseConversion(t *testing.T) {
	tests := []struct {
		name     string
		template string
		want     string
	}{
		{name: "camelCase", template: `{{camelCase "hello_world"}}`, want: "helloWorld"},
		{name: "snakeCase", template: `{{snakeCase "HelloWorld"}}`, want: "hello_world"},
		{name: "kebabCase", template: `{{kebabCase "HelloWorld"}}`, want: "hello-world"},
		{name: "pascalCase", template: `{{pascalCase "hello_world"}}`, want: "HelloWorld"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpl, err := texttemplate.New("test").Funcs(inflect.FuncMap()).Parse(tt.template)
			require.NoError(t, err)

			var buf bytes.Buffer
			err = tmpl.Execute(&buf, nil)
			require.NoError(t, err)
			assert.Equal(t, tt.want, buf.String())
		})
	}
}

func TestFuncMapPOS(t *testing.T) {
	tests := []struct {
		name     string
		template string
		data     any
		want     string
	}{
		{
			name:     "pluralNoun",
			template: `{{pluralNoun "I"}}`,
			data:     nil,
			want:     "We", // preserves case of input "I"
		},
		{
			name:     "pluralNoun with count 1",
			template: `{{pluralNoun "cat" .Count}}`,
			data:     map[string]int{"Count": 1},
			want:     "cat",
		},
		{
			name:     "pluralVerb",
			template: `{{pluralVerb "is"}}`,
			data:     nil,
			want:     "are",
		},
		{
			name:     "pluralVerb with count 1",
			template: `{{pluralVerb "is" .Count}}`,
			data:     map[string]int{"Count": 1},
			want:     "is",
		},
		{
			name:     "pluralAdj",
			template: `{{pluralAdj "this"}}`,
			data:     nil,
			want:     "these",
		},
		{
			name:     "singularNoun cats",
			template: `{{singularNoun "cats"}}`,
			data:     nil,
			want:     "cat",
		},
		{
			name:     "singularNoun with count 2 returns as-is",
			template: `{{singularNoun "cats" .Count}}`,
			data:     map[string]int{"Count": 2},
			want:     "cats", // count != 1 returns input as-is (plural form)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpl, err := texttemplate.New("test").Funcs(inflect.FuncMap()).Parse(tt.template)
			require.NoError(t, err)

			var buf bytes.Buffer
			err = tmpl.Execute(&buf, tt.data)
			require.NoError(t, err)
			assert.Equal(t, tt.want, buf.String())
		})
	}
}

func TestEngineFuncMapWithCustomNoun(t *testing.T) {
	e := inflect.NewEngine()
	e.DefNoun("gremlin", "gremloz")

	tmpl, err := texttemplate.New("test").Funcs(e.FuncMap()).Parse(`{{plural "gremlin"}}`)
	require.NoError(t, err)

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, nil)
	require.NoError(t, err)
	assert.Equal(t, "gremloz", buf.String())
}

func TestEngineFuncMapWithClassicalMode(t *testing.T) {
	e := inflect.NewEngine()
	e.ClassicalAncient(true)

	tmpl, err := texttemplate.New("test").Funcs(e.FuncMap()).Parse(`{{plural "formula"}}`)
	require.NoError(t, err)

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, nil)
	require.NoError(t, err)
	assert.Equal(t, "formulae", buf.String())
}

func TestFuncMapWithHTMLTemplate(t *testing.T) {
	// Verify FuncMap works with html/template as well
	// html/template.FuncMap is the same underlying type as text/template.FuncMap
	tmpl, err := htmltemplate.New("test").Funcs(inflect.FuncMap()).Parse(`<p>I have {{plural "cat" .Count}}</p>`)
	require.NoError(t, err)

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, map[string]int{"Count": 3})
	require.NoError(t, err)
	assert.Equal(t, "<p>I have cats</p>", buf.String())
}

func TestFuncMapComplexTemplate(t *testing.T) {
	tmplStr := `I have {{.Count}} {{plural "cat" .Count}} and {{an "apple"}}. 
The {{ordinalWord .Position}} cat is the {{superlative "big"}}.`

	tmpl, err := texttemplate.New("test").Funcs(inflect.FuncMap()).Parse(tmplStr)
	require.NoError(t, err)

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, map[string]int{"Count": 3, "Position": 1})
	require.NoError(t, err)

	want := `I have 3 cats and an apple. 
The first cat is the biggest.`
	assert.Equal(t, want, buf.String())
}

func TestFuncMapPluralWithNegativeOne(t *testing.T) {
	// -1 should also return singular (special case for "no items" vs "-1 item")
	tmpl, err := texttemplate.New("test").Funcs(inflect.FuncMap()).Parse(`{{plural "cat" .Count}}`)
	require.NoError(t, err)

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, map[string]int{"Count": -1})
	require.NoError(t, err)
	assert.Equal(t, "cat", buf.String())
}

func BenchmarkFuncMapPlural(b *testing.B) {
	fm := inflect.FuncMap()
	tmpl, _ := texttemplate.New("test").Funcs(fm).Parse(`{{plural "cat"}}`)
	var buf bytes.Buffer

	for b.Loop() {
		buf.Reset()
		_ = tmpl.Execute(&buf, nil)
	}
}

func BenchmarkFuncMapComplex(b *testing.B) {
	fm := inflect.FuncMap()
	tmpl, _ := texttemplate.New("test").Funcs(fm).Parse(`{{plural "cat" .Count}} and {{an "apple"}}`)
	data := map[string]int{"Count": 3}
	var buf bytes.Buffer

	for b.Loop() {
		buf.Reset()
		_ = tmpl.Execute(&buf, data)
	}
}
