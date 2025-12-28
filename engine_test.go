package inflect

import (
	"regexp"
	"testing"
)

func TestNewEngine(t *testing.T) {
	e := NewEngine()

	// Verify classical flags are false by default
	if e.classicalMode {
		t.Error("classicalMode should be false by default")
	}
	if e.classicalAll {
		t.Error("classicalAll should be false by default")
	}
	if e.classicalZero {
		t.Error("classicalZero should be false by default")
	}
	if e.classicalHerd {
		t.Error("classicalHerd should be false by default")
	}
	if e.classicalNames {
		t.Error("classicalNames should be false by default")
	}
	if e.classicalAncient {
		t.Error("classicalAncient should be false by default")
	}
	if e.classicalPersons {
		t.Error("classicalPersons should be false by default")
	}

	// Verify irregularPlurals is initialized from defaultIrregularPlurals
	if len(e.irregularPlurals) != len(defaultIrregularPlurals) {
		t.Errorf("irregularPlurals length = %d, want %d", len(e.irregularPlurals), len(defaultIrregularPlurals))
	}
	if e.irregularPlurals["child"] != "children" {
		t.Error("irregularPlurals should contain 'child' -> 'children'")
	}

	// Verify singularIrregulars is built correctly
	if e.singularIrregulars["children"] != "child" {
		t.Error("singularIrregulars should contain 'children' -> 'child'")
	}

	// Verify custom maps are empty
	if len(e.customVerbs) != 0 {
		t.Error("customVerbs should be empty by default")
	}
	if len(e.customVerbsReverse) != 0 {
		t.Error("customVerbsReverse should be empty by default")
	}
	if len(e.customAdjs) != 0 {
		t.Error("customAdjs should be empty by default")
	}
	if len(e.customAdjsReverse) != 0 {
		t.Error("customAdjsReverse should be empty by default")
	}
	if len(e.customAWords) != 0 {
		t.Error("customAWords should be empty by default")
	}
	if len(e.customAnWords) != 0 {
		t.Error("customAnWords should be empty by default")
	}
	if e.customAPatterns != nil {
		t.Error("customAPatterns should be nil by default")
	}
	if e.customAnPatterns != nil {
		t.Error("customAnPatterns should be nil by default")
	}

	// Verify gender is "t" (singular they) by default
	if e.gender != "t" {
		t.Errorf("gender = %q, want \"t\"", e.gender)
	}

	// Verify possessiveStyle is PossessiveModern by default
	if e.possessiveStyle != PossessiveModern {
		t.Errorf("possessiveStyle = %d, want %d", e.possessiveStyle, PossessiveModern)
	}

	// Verify defaultNum is 0 by default
	if e.defaultNum != 0 {
		t.Errorf("defaultNum = %d, want 0", e.defaultNum)
	}
}

func TestEngineClone(t *testing.T) {
	// Create and modify an engine
	e1 := NewEngine()
	e1.classicalAll = true
	e1.classicalAncient = true
	e1.irregularPlurals["test"] = "tests"
	e1.singularIrregulars["tests"] = "test"
	e1.customVerbs["run"] = "runs"
	e1.customAWords["euro"] = true
	e1.customAPatterns = []*regexp.Regexp{regexp.MustCompile("^euro.*$")}
	e1.gender = "m"
	e1.possessiveStyle = PossessiveTraditional
	e1.defaultNum = 42

	// Clone it
	e2 := e1.Clone()

	// Verify all values are copied correctly
	if e2.classicalAll != true {
		t.Error("Clone: classicalAll should be true")
	}
	if e2.classicalAncient != true {
		t.Error("Clone: classicalAncient should be true")
	}
	if e2.irregularPlurals["test"] != "tests" {
		t.Error("Clone: irregularPlurals should contain 'test' -> 'tests'")
	}
	if e2.singularIrregulars["tests"] != "test" {
		t.Error("Clone: singularIrregulars should contain 'tests' -> 'test'")
	}
	if e2.customVerbs["run"] != "runs" {
		t.Error("Clone: customVerbs should contain 'run' -> 'runs'")
	}
	if !e2.customAWords["euro"] {
		t.Error("Clone: customAWords should contain 'euro'")
	}
	if len(e2.customAPatterns) != 1 {
		t.Error("Clone: customAPatterns should have 1 pattern")
	}
	if e2.gender != "m" {
		t.Errorf("Clone: gender = %q, want \"m\"", e2.gender)
	}
	if e2.possessiveStyle != PossessiveTraditional {
		t.Error("Clone: possessiveStyle should be PossessiveTraditional")
	}
	if e2.defaultNum != 42 {
		t.Errorf("Clone: defaultNum = %d, want 42", e2.defaultNum)
	}

	// Verify independence: modify e2 and check e1 is unchanged
	e2.irregularPlurals["foo"] = "foos"
	e2.customVerbs["walk"] = "walks"
	e2.customAWords["unique"] = true

	if _, ok := e1.irregularPlurals["foo"]; ok {
		t.Error("Clone: modifying clone's irregularPlurals should not affect original")
	}
	if _, ok := e1.customVerbs["walk"]; ok {
		t.Error("Clone: modifying clone's customVerbs should not affect original")
	}
	if e1.customAWords["unique"] {
		t.Error("Clone: modifying clone's customAWords should not affect original")
	}

	// Verify independence: modify e1 and check e2 is unchanged
	e1.irregularPlurals["bar"] = "bars"

	if _, ok := e2.irregularPlurals["bar"]; ok {
		t.Error("Clone: modifying original's irregularPlurals should not affect clone")
	}
}

func TestEngineCloneNilSlices(t *testing.T) {
	// Test cloning with nil slices
	e1 := NewEngine()
	e2 := e1.Clone()

	if e2.customAPatterns != nil {
		t.Error("Clone: nil customAPatterns should remain nil")
	}
	if e2.customAnPatterns != nil {
		t.Error("Clone: nil customAnPatterns should remain nil")
	}
}

func TestEngineCloneWithPatterns(t *testing.T) {
	e1 := NewEngine()
	e1.customAPatterns = []*regexp.Regexp{
		regexp.MustCompile("^euro.*$"),
		regexp.MustCompile("^uni.*$"),
	}
	e1.customAnPatterns = []*regexp.Regexp{
		regexp.MustCompile("^honor.*$"),
	}

	e2 := e1.Clone()

	// Verify patterns are copied
	if len(e2.customAPatterns) != 2 {
		t.Errorf("Clone: customAPatterns length = %d, want 2", len(e2.customAPatterns))
	}
	if len(e2.customAnPatterns) != 1 {
		t.Errorf("Clone: customAnPatterns length = %d, want 1", len(e2.customAnPatterns))
	}

	// Verify independence: modifying slice doesn't affect original
	e2.customAPatterns = append(e2.customAPatterns, regexp.MustCompile("^new.*$"))
	if len(e1.customAPatterns) != 2 {
		t.Error("Clone: modifying clone's customAPatterns slice should not affect original")
	}
}

func TestEngineClassicalMethods(t *testing.T) {
	t.Run("ClassicalAll", func(t *testing.T) {
		e := NewEngine()

		// Default should be false
		if e.IsClassicalAll() {
			t.Error("IsClassicalAll should be false by default")
		}

		// Enable all
		e.ClassicalAll(true)
		if !e.IsClassicalAll() {
			t.Error("IsClassicalAll should be true after ClassicalAll(true)")
		}
		if !e.IsClassicalAncient() {
			t.Error("IsClassicalAncient should be true after ClassicalAll(true)")
		}
		if !e.IsClassicalZero() {
			t.Error("IsClassicalZero should be true after ClassicalAll(true)")
		}
		if !e.IsClassicalHerd() {
			t.Error("IsClassicalHerd should be true after ClassicalAll(true)")
		}
		if !e.IsClassicalNames() {
			t.Error("IsClassicalNames should be true after ClassicalAll(true)")
		}
		if !e.IsClassicalPersons() {
			t.Error("IsClassicalPersons should be true after ClassicalAll(true)")
		}

		// Disable all
		e.ClassicalAll(false)
		if e.IsClassicalAll() {
			t.Error("IsClassicalAll should be false after ClassicalAll(false)")
		}
	})

	t.Run("Classical", func(t *testing.T) {
		e := NewEngine()

		e.Classical(true)
		if !e.IsClassical() {
			t.Error("IsClassical should be true after Classical(true)")
		}
		if !e.IsClassicalAll() {
			t.Error("IsClassicalAll should be true after Classical(true)")
		}

		e.Classical(false)
		if e.IsClassical() {
			t.Error("IsClassical should be false after Classical(false)")
		}
	})

	t.Run("ClassicalAncient", func(t *testing.T) {
		e := NewEngine()

		if e.IsClassicalAncient() {
			t.Error("IsClassicalAncient should be false by default")
		}

		e.ClassicalAncient(true)
		if !e.IsClassicalAncient() {
			t.Error("IsClassicalAncient should be true after ClassicalAncient(true)")
		}
		if !e.IsClassical() {
			t.Error("IsClassical should be true when ClassicalAncient is enabled")
		}

		e.ClassicalAncient(false)
		if e.IsClassicalAncient() {
			t.Error("IsClassicalAncient should be false after ClassicalAncient(false)")
		}
	})

	t.Run("ClassicalZero", func(t *testing.T) {
		e := NewEngine()

		if e.IsClassicalZero() {
			t.Error("IsClassicalZero should be false by default")
		}

		e.ClassicalZero(true)
		if !e.IsClassicalZero() {
			t.Error("IsClassicalZero should be true after ClassicalZero(true)")
		}

		e.ClassicalZero(false)
		if e.IsClassicalZero() {
			t.Error("IsClassicalZero should be false after ClassicalZero(false)")
		}
	})

	t.Run("ClassicalHerd", func(t *testing.T) {
		e := NewEngine()

		if e.IsClassicalHerd() {
			t.Error("IsClassicalHerd should be false by default")
		}

		e.ClassicalHerd(true)
		if !e.IsClassicalHerd() {
			t.Error("IsClassicalHerd should be true after ClassicalHerd(true)")
		}

		e.ClassicalHerd(false)
		if e.IsClassicalHerd() {
			t.Error("IsClassicalHerd should be false after ClassicalHerd(false)")
		}
	})

	t.Run("ClassicalNames", func(t *testing.T) {
		e := NewEngine()

		if e.IsClassicalNames() {
			t.Error("IsClassicalNames should be false by default")
		}

		e.ClassicalNames(true)
		if !e.IsClassicalNames() {
			t.Error("IsClassicalNames should be true after ClassicalNames(true)")
		}

		e.ClassicalNames(false)
		if e.IsClassicalNames() {
			t.Error("IsClassicalNames should be false after ClassicalNames(false)")
		}
	})

	t.Run("ClassicalPersons", func(t *testing.T) {
		e := NewEngine()

		if e.IsClassicalPersons() {
			t.Error("IsClassicalPersons should be false by default")
		}

		e.ClassicalPersons(true)
		if !e.IsClassicalPersons() {
			t.Error("IsClassicalPersons should be true after ClassicalPersons(true)")
		}

		e.ClassicalPersons(false)
		if e.IsClassicalPersons() {
			t.Error("IsClassicalPersons should be false after ClassicalPersons(false)")
		}
	})

	t.Run("Independence", func(t *testing.T) {
		e := NewEngine()

		// Enable individual flags and verify they don't affect others
		e.ClassicalAncient(true)
		if e.IsClassicalPersons() {
			t.Error("ClassicalAncient should not affect ClassicalPersons")
		}
		if e.IsClassicalHerd() {
			t.Error("ClassicalAncient should not affect ClassicalHerd")
		}

		e.ClassicalAncient(false)
		e.ClassicalPersons(true)
		if e.IsClassicalAncient() {
			t.Error("ClassicalPersons should not affect ClassicalAncient")
		}

		// After enabling all, disabling one should not affect others
		e.ClassicalAll(true)
		e.ClassicalAncient(false)
		if !e.IsClassicalPersons() {
			t.Error("Disabling ClassicalAncient should not affect ClassicalPersons")
		}
		if !e.IsClassicalHerd() {
			t.Error("Disabling ClassicalAncient should not affect ClassicalHerd")
		}
		if e.IsClassicalAll() {
			t.Error("IsClassicalAll should be false after disabling one flag")
		}
	})
}

func TestEngineIsolation(t *testing.T) {
	// Verify that different Engine instances are isolated
	e1 := NewEngine()
	e2 := NewEngine()

	e1.ClassicalAll(true)

	if e2.IsClassicalAll() {
		t.Error("Modifying e1 should not affect e2")
	}
	if e2.IsClassicalAncient() {
		t.Error("Modifying e1 should not affect e2")
	}
}

func TestEnginePluralNoun(t *testing.T) {
	e := NewEngine()

	// Test basic functionality
	tests := []struct {
		name     string
		word     string
		count    []int
		expected string
	}{
		{name: "I -> We", word: "I", expected: "We"},
		{name: "cat -> cats", word: "cat", expected: "cats"},
		{name: "cat count=1 singular", word: "cat", count: []int{1}, expected: "cat"},
		{name: "cat count=2 plural", word: "cat", count: []int{2}, expected: "cats"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got string
			if len(tt.count) > 0 {
				got = e.PluralNoun(tt.word, tt.count[0])
			} else {
				got = e.PluralNoun(tt.word)
			}
			if got != tt.expected {
				t.Errorf("PluralNoun(%q, %v) = %q, want %q", tt.word, tt.count, got, tt.expected)
			}
		})
	}
}

func TestEnginePluralVerb(t *testing.T) {
	e := NewEngine()

	// Test basic functionality
	tests := []struct {
		name     string
		word     string
		count    []int
		expected string
	}{
		{name: "is -> are", word: "is", expected: "are"},
		{name: "was -> were", word: "was", expected: "were"},
		{name: "is count=1 singular", word: "is", count: []int{1}, expected: "is"},
		{name: "is count=2 plural", word: "is", count: []int{2}, expected: "are"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got string
			if len(tt.count) > 0 {
				got = e.PluralVerb(tt.word, tt.count[0])
			} else {
				got = e.PluralVerb(tt.word)
			}
			if got != tt.expected {
				t.Errorf("PluralVerb(%q, %v) = %q, want %q", tt.word, tt.count, got, tt.expected)
			}
		})
	}
}

func TestEnginePluralAdj(t *testing.T) {
	e := NewEngine()

	// Test basic functionality
	tests := []struct {
		name     string
		word     string
		count    []int
		expected string
	}{
		{name: "this -> these", word: "this", expected: "these"},
		{name: "that -> those", word: "that", expected: "those"},
		{name: "this count=1 singular", word: "this", count: []int{1}, expected: "this"},
		{name: "this count=2 plural", word: "this", count: []int{2}, expected: "these"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got string
			if len(tt.count) > 0 {
				got = e.PluralAdj(tt.word, tt.count[0])
			} else {
				got = e.PluralAdj(tt.word)
			}
			if got != tt.expected {
				t.Errorf("PluralAdj(%q, %v) = %q, want %q", tt.word, tt.count, got, tt.expected)
			}
		})
	}
}

func TestEngineSingularNoun(t *testing.T) {
	e := NewEngine()

	// Test basic functionality with default gender (t)
	tests := []struct {
		name     string
		word     string
		count    []int
		expected string
	}{
		{name: "we -> I", word: "we", expected: "I"},
		{name: "cats -> cat", word: "cats", expected: "cat"},
		{name: "cats count=1 singular", word: "cats", count: []int{1}, expected: "cat"},
		{name: "cats count=2 plural", word: "cats", count: []int{2}, expected: "cats"},
		{name: "they -> they (default gender t)", word: "they", expected: "they"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got string
			if len(tt.count) > 0 {
				got = e.SingularNoun(tt.word, tt.count[0])
			} else {
				got = e.SingularNoun(tt.word)
			}
			if got != tt.expected {
				t.Errorf("SingularNoun(%q, %v) = %q, want %q", tt.word, tt.count, got, tt.expected)
			}
		})
	}
}

func TestEngineSingularNounWithGender(t *testing.T) {
	// Test gender independence between engines
	e1 := NewEngine()
	e2 := NewEngine()

	e1.SetGender("m")
	e2.SetGender("f")

	// e1 should use masculine gender
	if got := e1.SingularNoun("they"); got != "he" {
		t.Errorf("e1.SingularNoun(\"they\") with masculine = %q, want \"he\"", got)
	}

	// e2 should use feminine gender
	if got := e2.SingularNoun("they"); got != "she" {
		t.Errorf("e2.SingularNoun(\"they\") with feminine = %q, want \"she\"", got)
	}
}

func TestEngineNo(t *testing.T) {
	e := NewEngine()

	tests := []struct {
		name     string
		word     string
		count    int
		expected string
	}{
		{name: "no errors", word: "error", count: 0, expected: "no errors"},
		{name: "1 error", word: "error", count: 1, expected: "1 error"},
		{name: "2 errors", word: "error", count: 2, expected: "2 errors"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := e.No(tt.word, tt.count)
			if got != tt.expected {
				t.Errorf("No(%q, %d) = %q, want %q", tt.word, tt.count, got, tt.expected)
			}
		})
	}
}

func TestEngineNoWithClassicalZero(t *testing.T) {
	e := NewEngine()
	e.ClassicalZero(true)

	if got := e.No("error", 0); got != "no error" {
		t.Errorf("No(\"error\", 0) with ClassicalZero = %q, want \"no error\"", got)
	}
}

func TestEngineNum(t *testing.T) {
	e := NewEngine()

	// Test setting and getting
	if got := e.Num(5); got != 5 {
		t.Errorf("Num(5) = %d, want 5", got)
	}
	if got := e.GetNum(); got != 5 {
		t.Errorf("GetNum() = %d, want 5", got)
	}

	// Test clearing
	if got := e.Num(0); got != 0 {
		t.Errorf("Num(0) = %d, want 0", got)
	}
	if got := e.GetNum(); got != 0 {
		t.Errorf("GetNum() after Num(0) = %d, want 0", got)
	}

	// Test independence between engines
	e1 := NewEngine()
	e2 := NewEngine()

	e1.Num(10)
	e2.Num(20)

	if got := e1.GetNum(); got != 10 {
		t.Errorf("e1.GetNum() = %d, want 10", got)
	}
	if got := e2.GetNum(); got != 20 {
		t.Errorf("e2.GetNum() = %d, want 20", got)
	}
}

func TestEngineGetGender(t *testing.T) {
	e := NewEngine()

	// Test default gender
	if got := e.GetGender(); got != "t" {
		t.Errorf("GetGender() = %q, want \"t\"", got)
	}

	// Test setting gender
	e.SetGender("m")
	if got := e.GetGender(); got != "m" {
		t.Errorf("GetGender() after SetGender(\"m\") = %q, want \"m\"", got)
	}

	// Test invalid gender is ignored
	e.SetGender("invalid")
	if got := e.GetGender(); got != "m" {
		t.Errorf("GetGender() after SetGender(\"invalid\") = %q, want \"m\"", got)
	}
}
