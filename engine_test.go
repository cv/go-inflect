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
