package inflect_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	inflect "github.com/cv/go-inflect"
)

func TestDasherize(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		// Basic conversions
		{name: "PascalCase", input: "HelloWorld", want: "hello-world"},
		{name: "camelCase", input: "helloWorld", want: "hello-world"},
		{name: "snake_case", input: "hello_world", want: "hello-world"},
		{name: "already kebab-case", input: "hello-world", want: "hello-world"},
		{name: "space separated", input: "hello world", want: "hello-world"},

		// Consecutive uppercase (acronyms)
		{name: "acronym at start", input: "HTTPServer", want: "http-server"},
		{name: "acronym in middle", input: "getHTTPResponse", want: "get-http-response"},
		{name: "acronym at end", input: "parseJSON", want: "parse-json"},
		{name: "multiple acronyms", input: "XMLHTTPRequest", want: "xmlhttp-request"}, // consecutive uppercase is one word
		{name: "all uppercase", input: "HTTP", want: "http"},
		{name: "uppercase with numbers", input: "HTTP2Server", want: "http-2-server"},

		// Numbers
		{name: "number at end", input: "version2", want: "version-2"},
		{name: "number in middle", input: "v2release", want: "v-2-release"},
		{name: "multiple numbers", input: "test123value", want: "test-123-value"},
		{name: "number at start", input: "123test", want: "123-test"},

		// Edge cases
		{name: "empty string", input: "", want: ""},
		{name: "single word", input: "hello", want: "hello"},
		{name: "single uppercase", input: "H", want: "h"},
		{name: "single lowercase", input: "h", want: "h"},
		{name: "leading separator", input: "_hello", want: "hello"},
		{name: "trailing separator", input: "hello_", want: "hello"},
		{name: "multiple separators", input: "hello__world", want: "hello-world"},
		{name: "mixed separators", input: "hello_-world", want: "hello-world"},
		{name: "leading and trailing", input: "_hello_world_", want: "hello-world"},

		// Complex cases
		{name: "mixed format", input: "XMLParser_v2", want: "xml-parser-v-2"},
		{name: "complex mixed", input: "getHTTPResponseCode", want: "get-http-response-code"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Dasherize(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestKebabCase(t *testing.T) {
	// KebabCase should be an alias for Dasherize
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "PascalCase", input: "HelloWorld", want: "hello-world"},
		{name: "snake_case", input: "hello_world", want: "hello-world"},
		{name: "HTTPServer", input: "HTTPServer", want: "http-server"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.KebabCase(tt.input)
			assert.Equal(t, tt.want, got)
			// Verify it matches Dasherize
			assert.Equal(t, inflect.Dasherize(tt.input), got)
		})
	}
}

func TestUnderscore(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		// Basic conversions
		{name: "PascalCase", input: "HelloWorld", want: "hello_world"},
		{name: "camelCase", input: "helloWorld", want: "hello_world"},
		{name: "kebab-case", input: "hello-world", want: "hello_world"},
		{name: "already snake_case", input: "hello_world", want: "hello_world"},
		{name: "space separated", input: "hello world", want: "hello_world"},

		// Consecutive uppercase (acronyms)
		{name: "acronym at start", input: "HTTPServer", want: "http_server"},
		{name: "acronym in middle", input: "getHTTPResponse", want: "get_http_response"},
		{name: "acronym at end", input: "parseJSON", want: "parse_json"},
		{name: "multiple acronyms", input: "XMLHTTPRequest", want: "xmlhttp_request"}, // consecutive uppercase is one word
		{name: "all uppercase", input: "HTTP", want: "http"},
		{name: "uppercase with numbers", input: "HTTP2Server", want: "http_2_server"},

		// Numbers
		{name: "number at end", input: "version2", want: "version_2"},
		{name: "number in middle", input: "v2release", want: "v_2_release"},
		{name: "multiple numbers", input: "test123value", want: "test_123_value"},
		{name: "number at start", input: "123test", want: "123_test"},

		// Edge cases
		{name: "empty string", input: "", want: ""},
		{name: "single word", input: "hello", want: "hello"},
		{name: "single uppercase", input: "H", want: "h"},
		{name: "single lowercase", input: "h", want: "h"},
		{name: "leading separator", input: "-hello", want: "hello"},
		{name: "trailing separator", input: "hello-", want: "hello"},
		{name: "multiple separators", input: "hello--world", want: "hello_world"},
		{name: "mixed separators", input: "hello-_world", want: "hello_world"},
		{name: "leading and trailing", input: "-hello-world-", want: "hello_world"},

		// Complex cases
		{name: "mixed format", input: "XMLParser-v2", want: "xml_parser_v_2"},
		{name: "complex mixed", input: "getHTTPResponseCode", want: "get_http_response_code"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.Underscore(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSnakeCase(t *testing.T) {
	// SnakeCase should be an alias for Underscore
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "PascalCase", input: "HelloWorld", want: "hello_world"},
		{name: "kebab-case", input: "hello-world", want: "hello_world"},
		{name: "HTTPServer", input: "HTTPServer", want: "http_server"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.SnakeCase(tt.input)
			assert.Equal(t, tt.want, got)
			// Verify it matches Underscore
			assert.Equal(t, inflect.Underscore(tt.input), got)
		})
	}
}

func TestPascalCase(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		// Basic conversions
		{name: "snake_case", input: "hello_world", want: "HelloWorld"},
		{name: "kebab-case", input: "hello-world", want: "HelloWorld"},
		{name: "space separated", input: "hello world", want: "HelloWorld"},
		{name: "already PascalCase", input: "HelloWorld", want: "HelloWorld"},
		{name: "camelCase", input: "helloWorld", want: "HelloWorld"},

		// Uppercase handling
		{name: "all uppercase", input: "HTTP_SERVER", want: "HttpServer"},
		{name: "uppercase word", input: "HTTP", want: "Http"},
		{name: "acronym at start", input: "http_server", want: "HttpServer"},

		// Numbers
		{name: "number at end", input: "version_2", want: "Version2"},
		{name: "number in middle", input: "v2_release", want: "V2Release"},
		{name: "number only word", input: "test_123_value", want: "Test123Value"},

		// Edge cases
		{name: "empty string", input: "", want: ""},
		{name: "single word", input: "hello", want: "Hello"},
		{name: "single uppercase", input: "H", want: "H"},
		{name: "single lowercase", input: "h", want: "H"},
		{name: "leading separator", input: "_hello", want: "Hello"},
		{name: "trailing separator", input: "hello_", want: "Hello"},
		{name: "multiple separators", input: "hello__world", want: "HelloWorld"},

		// Complex cases
		{name: "three words", input: "one_two_three", want: "OneTwoThree"},
		{name: "mixed separators", input: "one-two_three", want: "OneTwoThree"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.PascalCase(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestTitleCase(t *testing.T) {
	// TitleCase should be an alias for PascalCase
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "snake_case", input: "hello_world", want: "HelloWorld"},
		{name: "kebab-case", input: "hello-world", want: "HelloWorld"},
		{name: "HTTP_SERVER", input: "HTTP_SERVER", want: "HttpServer"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.TitleCase(tt.input)
			assert.Equal(t, tt.want, got)
			// Verify it matches PascalCase
			assert.Equal(t, inflect.PascalCase(tt.input), got)
		})
	}
}

func TestCamelCase(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		// Basic conversions
		{name: "snake_case", input: "hello_world", want: "helloWorld"},
		{name: "kebab-case", input: "hello-world", want: "helloWorld"},
		{name: "space separated", input: "hello world", want: "helloWorld"},
		{name: "PascalCase", input: "HelloWorld", want: "helloWorld"},
		{name: "already camelCase", input: "helloWorld", want: "helloWorld"},

		// Uppercase handling
		{name: "all uppercase", input: "HTTP_SERVER", want: "httpServer"},
		{name: "uppercase word", input: "HTTP", want: "http"},
		{name: "acronym at start", input: "http_server", want: "httpServer"},

		// Numbers
		{name: "number at end", input: "version_2", want: "version2"},
		{name: "number in middle", input: "v2_release", want: "v2Release"},
		{name: "number only word", input: "test_123_value", want: "test123Value"},

		// Edge cases
		{name: "empty string", input: "", want: ""},
		{name: "single word", input: "hello", want: "hello"},
		{name: "single uppercase", input: "H", want: "h"},
		{name: "single lowercase", input: "h", want: "h"},
		{name: "leading separator", input: "_hello", want: "hello"},
		{name: "trailing separator", input: "hello_", want: "hello"},
		{name: "multiple separators", input: "hello__world", want: "helloWorld"},

		// Complex cases
		{name: "three words", input: "one_two_three", want: "oneTwoThree"},
		{name: "mixed separators", input: "one-two_three", want: "oneTwoThree"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.CamelCase(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

// TestCaseRoundTrip tests round-trip conversions.
func TestCaseRoundTrip(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{name: "simple", input: "hello_world"},
		{name: "three words", input: "one_two_three"},
		{name: "with number", input: "version_2"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// snake_case -> PascalCase -> snake_case
			pascal := inflect.PascalCase(tt.input)
			backToSnake := inflect.SnakeCase(pascal)
			assert.Equal(t, tt.input, backToSnake)

			// snake_case -> kebab-case -> snake_case
			kebab := inflect.Dasherize(tt.input)
			backToSnake2 := inflect.Underscore(kebab)
			assert.Equal(t, tt.input, backToSnake2)
		})
	}
}

// Benchmarks for case conversion functions.
func BenchmarkDasherize(b *testing.B) {
	inputs := []string{"HelloWorld", "getHTTPResponse", "hello_world", "already-kebab"}
	for i := range b.N {
		inflect.Dasherize(inputs[i%len(inputs)])
	}
}

func BenchmarkUnderscore(b *testing.B) {
	inputs := []string{"HelloWorld", "getHTTPResponse", "hello-world", "already_snake"}
	for i := range b.N {
		inflect.Underscore(inputs[i%len(inputs)])
	}
}

func BenchmarkPascalCase(b *testing.B) {
	inputs := []string{"hello_world", "hello-world", "already_pascal_case"}
	for i := range b.N {
		inflect.PascalCase(inputs[i%len(inputs)])
	}
}

func BenchmarkCamelCase(b *testing.B) {
	inputs := []string{"hello_world", "hello-world", "already_camel_case"}
	for i := range b.N {
		inflect.CamelCase(inputs[i%len(inputs)])
	}
}

func BenchmarkSnakeCase(b *testing.B) {
	benchmarks := []struct {
		name  string
		input string
	}{
		{"camel", "camelCase"},
		{"pascal", "PascalCase"},
		{"kebab", "kebab-case"},
		{"spaces", "hello world"},
		{"mixed", "XMLHttpRequest"},
		{"already_snake", "already_snake_case"},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for range b.N {
				inflect.SnakeCase(bm.input)
			}
		})
	}
}

func BenchmarkKebabCase(b *testing.B) {
	benchmarks := []struct {
		name  string
		input string
	}{
		{"camel", "camelCase"},
		{"pascal", "PascalCase"},
		{"snake", "snake_case"},
		{"spaces", "hello world"},
		{"mixed", "XMLHttpRequest"},
		{"already_kebab", "already-kebab-case"},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for range b.N {
				inflect.KebabCase(bm.input)
			}
		})
	}
}
