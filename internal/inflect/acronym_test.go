package inflect

import (
	"slices"
	"testing"
)

func TestSplitPascalCase(t *testing.T) {
	tests := []struct {
		input string
		want  []string
	}{
		{"", nil},
		{"Hello", []string{"Hello"}},
		{"HelloWorld", []string{"Hello", "World"}},
		{"GPUConfig", []string{"GPU", "Config"}},
		{"DNSRecord", []string{"DNS", "Record"}},
		{"LDAPUser", []string{"LDAP", "User"}},
		{"CoreweaveGPUStatus", []string{"Coreweave", "GPU", "Status"}},
		{"myXMLParser", []string{"my", "XML", "Parser"}},
		{"OAuth2Client", []string{"O", "Auth", "2", "Client"}}, // Known limitation: "OA" treated as acronym
		{"HTTPSConnection", []string{"HTTPS", "Connection"}},
		{"getHTTPResponse", []string{"get", "HTTP", "Response"}},
		{"ID", []string{"ID"}},
		{"UUID", []string{"UUID"}},
		{"SimpleWord", []string{"Simple", "Word"}},
		{"ABC", []string{"ABC"}},
		{"ABCDef", []string{"ABC", "Def"}},
		{"DefABC", []string{"Def", "ABC"}},
		{"already_snake", []string{"already", "snake"}},
		{"kebab-case", []string{"kebab", "case"}},
		{"Mixed_CASE-input", []string{"Mixed", "CASE", "input"}},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := SplitPascalCase(tt.input)
			if !slices.Equal(got, tt.want) {
				t.Errorf("SplitPascalCase(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestAddAcronym(t *testing.T) {
	e := NewEngine()
	e.ClearAcronyms()

	if e.IsAcronym("GPU") {
		t.Error("GPU should not be an acronym after ClearAcronyms")
	}

	e.AddAcronym("GPU")
	if !e.IsAcronym("GPU") {
		t.Error("GPU should be an acronym after AddAcronym")
	}
	if !e.IsAcronym("gpu") {
		t.Error("IsAcronym should be case-insensitive")
	}
	if !e.IsAcronym("Gpu") {
		t.Error("IsAcronym should be case-insensitive")
	}
}

func TestRemoveAcronym(t *testing.T) {
	e := NewEngine()
	e.ClearAcronyms()
	e.AddAcronym("GPU")

	if !e.RemoveAcronym("GPU") {
		t.Error("RemoveAcronym should return true for registered acronym")
	}
	if e.IsAcronym("GPU") {
		t.Error("GPU should not be an acronym after RemoveAcronym")
	}
	if e.RemoveAcronym("GPU") {
		t.Error("RemoveAcronym should return false for non-registered acronym")
	}
}

func TestClearAcronyms(t *testing.T) {
	e := NewEngine()
	// Should have defaults
	if !e.IsAcronym("GPU") {
		t.Error("GPU should be a default acronym")
	}

	e.ClearAcronyms()
	if e.IsAcronym("GPU") {
		t.Error("GPU should not be an acronym after ClearAcronyms")
	}
}

func TestResetAcronyms(t *testing.T) {
	e := NewEngine()
	e.ClearAcronyms()
	if e.IsAcronym("GPU") {
		t.Error("GPU should not be an acronym after ClearAcronyms")
	}

	e.ResetAcronyms()
	if !e.IsAcronym("GPU") {
		t.Error("GPU should be an acronym after ResetAcronyms")
	}
}

func TestGetAcronyms(t *testing.T) {
	e := NewEngine()
	e.ClearAcronyms()
	e.AddAcronym("GPU")
	e.AddAcronym("CPU")
	e.AddAcronym("API")

	acronyms := e.GetAcronyms()
	if len(acronyms) != 3 {
		t.Errorf("Expected 3 acronyms, got %d", len(acronyms))
	}
	// Should be sorted
	if acronyms[0] != "API" || acronyms[1] != "CPU" || acronyms[2] != "GPU" {
		t.Errorf("Acronyms should be sorted, got %v", acronyms)
	}
}

func TestDefaultAcronyms(t *testing.T) {
	defaults := DefaultAcronyms()
	if len(defaults) == 0 {
		t.Error("DefaultAcronyms should not be empty")
	}
	// Check some expected defaults
	found := make(map[string]bool)
	for _, a := range defaults {
		found[a] = true
	}
	expectedDefaults := []string{"GPU", "CPU", "API", "DNS", "HTTP", "JSON", "URL", "ID"}
	for _, expected := range expectedDefaults {
		if !found[expected] {
			t.Errorf("DefaultAcronyms should include %q", expected)
		}
	}
}

func TestIsAcronymWithDefaults(t *testing.T) {
	// Package-level function should use defaults
	e := NewEngine()
	// Test that defaults are active even without explicit initialization
	if !e.IsAcronym("GPU") {
		t.Error("GPU should be a default acronym")
	}
	if !e.IsAcronym("API") {
		t.Error("API should be a default acronym")
	}
	if e.IsAcronym("Cat") {
		t.Error("Cat should not be an acronym")
	}
}

func TestHumanizeWithAcronyms(t *testing.T) {
	e := NewEngine()
	e.ResetAcronyms()

	tests := []struct {
		input string
		want  string
	}{
		{"GPUConfig", "GPU config"},
		{"DNSRecord", "DNS record"},
		{"myAPIServer", "My API server"},
		{"CoreweaveGPUStatus", "Coreweave GPU status"},
		{"HTTPSConnection", "HTTPS connection"},
		{"employee_salary", "Employee salary"},
		{"author_id", "Author"},
		{"hello-world", "Hello world"},
		{"SimpleWord", "Simple word"},
		{"ID", "ID"},
		{"UserID", "User"},
		{"user_api_key", "User API key"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := e.Humanize(tt.input)
			if got != tt.want {
				t.Errorf("Humanize(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestHumanizeWithoutAcronyms(t *testing.T) {
	e := NewEngine()
	e.ClearAcronyms()

	tests := []struct {
		input string
		want  string
	}{
		{"GPUConfig", "Gpu config"},
		{"DNSRecord", "Dns record"},
		{"myAPIServer", "My api server"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := e.Humanize(tt.input)
			if got != tt.want {
				t.Errorf("Humanize(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestPluralUppercase(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"GPU", "GPUs"},
		{"VM", "VMs"},
		{"ACL", "ACLs"},
		{"CPU", "CPUs"},
		{"API", "APIs"},
		{"URL", "URLs"},
		// Regular words should use normal rules
		{"cat", "cats"},
		{"box", "boxes"},
		// Single uppercase letter - uses existing matchSuffix behavior
		{"A", "AS"},
		// Mixed case uses normal rules
		{"Gpu", "Gpus"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := Plural(tt.input)
			if got != tt.want {
				t.Errorf("Plural(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestEngineCloneWithAcronyms(t *testing.T) {
	e1 := NewEngine()
	e1.ClearAcronyms()
	e1.AddAcronym("FOO")

	e2 := e1.Clone()
	e2.AddAcronym("BAR")

	if !e1.IsAcronym("FOO") {
		t.Error("e1 should have FOO")
	}
	if e1.IsAcronym("BAR") {
		t.Error("e1 should not have BAR")
	}
	if !e2.IsAcronym("FOO") {
		t.Error("e2 should have FOO")
	}
	if !e2.IsAcronym("BAR") {
		t.Error("e2 should have BAR")
	}
}

func TestEngineResetWithAcronyms(t *testing.T) {
	e := NewEngine()
	e.ClearAcronyms()
	e.AddAcronym("FOO")

	e.Reset()

	// After reset, should have defaults again
	if !e.IsAcronym("GPU") {
		t.Error("GPU should be present after Reset")
	}
	if e.IsAcronym("FOO") {
		t.Error("FOO should not be present after Reset")
	}
}

func TestIsAllUppercase(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"GPU", true},
		{"ABC", true},
		{"Gpu", false},
		{"gpu", false},
		{"GPu", false},
		{"", false},
		{"A", true},
		{"123", false},
		{"GPU123", false},
		{"ABC-DEF", false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := isAllUppercase(tt.input)
			if got != tt.want {
				t.Errorf("isAllUppercase(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func BenchmarkSplitPascalCase(b *testing.B) {
	inputs := []string{
		"HelloWorld",
		"GPUConfig",
		"CoreweaveGPUStatus",
		"myXMLParser",
		"OAuth2Client",
	}
	for b.Loop() {
		for _, input := range inputs {
			SplitPascalCase(input)
		}
	}
}

func BenchmarkHumanizeWithAcronyms(b *testing.B) {
	e := NewEngine()
	inputs := []string{
		"GPUConfig",
		"DNSRecord",
		"CoreweaveGPUStatus",
		"employee_salary",
	}
	for b.Loop() {
		for _, input := range inputs {
			e.Humanize(input)
		}
	}
}
