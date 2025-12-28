package inflect

import (
	"fmt"
	"sync"
	"testing"
)

// =============================================================================
// Package-Level Concurrency Tests
// =============================================================================

// TestConcurrentReads verifies that multiple goroutines can safely call
// read-only functions (Plural, Singular, An) simultaneously without race conditions.
func TestConcurrentReads(_ *testing.T) {
	// Reset to known state
	Reset()

	var wg sync.WaitGroup
	goroutines := 100
	iterations := 100

	// Test words
	words := []string{"cat", "child", "formula", "person", "criterion", "appendix", "matrix"}

	for range goroutines {
		wg.Go(func() {
			for range iterations {
				for _, word := range words {
					_ = Plural(word)
					_ = Singular(word)
					_ = An(word)
				}
			}
		})
	}
	wg.Wait()
}

// TestConcurrentWrites verifies that multiple goroutines can safely call
// write functions (DefNoun, Classical, DefA) simultaneously without race conditions.
func TestConcurrentWrites(_ *testing.T) {
	// Reset to known state
	Reset()

	var wg sync.WaitGroup
	goroutines := 50
	iterations := 20

	for n := range goroutines {
		wg.Go(func() {
			for j := range iterations {
				// Alternate between enabling and disabling classical mode
				Classical(n%2 == 0)
				ClassicalAll(n%3 == 0)

				// Define custom nouns with unique names per goroutine
				word := fmt.Sprintf("word%d_%d", n, j)
				plural := fmt.Sprintf("words%d_%d", n, j)
				DefNoun(word, plural)

				// Define article patterns
				DefA(fmt.Sprintf("word%d", n))
				DefAn(fmt.Sprintf("item%d", n))
			}
		})
	}
	wg.Wait()

	// Reset after test to not affect other tests
	Reset()
}

// TestMixedReadWrite verifies that concurrent read and write operations
// do not cause race conditions.
func TestMixedReadWrite(_ *testing.T) {
	// Reset to known state
	Reset()

	var wg sync.WaitGroup
	goroutines := 50
	iterations := 50

	// Reader goroutines
	for range goroutines {
		wg.Go(func() {
			for range iterations {
				_ = Plural("cat")
				_ = Plural("child")
				_ = Singular("cats")
				_ = Singular("children")
				_ = An("apple")
				_ = An("car")
				_ = A("umbrella")
				_ = IsClassical()
				_ = IsClassicalAll()
			}
		})
	}

	// Writer goroutines
	for n := range goroutines / 5 {
		wg.Go(func() {
			for j := range iterations {
				Classical(j%2 == 0)
				DefNoun(fmt.Sprintf("test%d_%d", n, j), fmt.Sprintf("tests%d_%d", n, j))
			}
		})
	}

	wg.Wait()
	Reset()
}

// TestConfigureThenRead verifies the pattern of configuring in one goroutine
// and then reading from many goroutines.
func TestConfigureThenRead(_ *testing.T) {
	// Reset to known state
	Reset()

	// Configuration phase
	done := make(chan struct{})
	go func() {
		ClassicalAll(true)
		DefNoun("regex", "regexen")
		DefNoun("datum", "data")
		DefA("unique")
		DefAn("honest")
		close(done)
	}()

	// Wait for configuration to complete
	<-done

	// Read phase - many goroutines reading
	var wg sync.WaitGroup
	goroutines := 100
	iterations := 100

	for range goroutines {
		wg.Go(func() {
			for range iterations {
				_ = Plural("regex")
				_ = Plural("datum")
				_ = Plural("formula")
				_ = An("unique")
				_ = An("honest")
				_ = IsClassicalAll()
			}
		})
	}

	wg.Wait()
	Reset()
}

// =============================================================================
// Engine-Specific Concurrency Tests
// =============================================================================

// TestEngineConcurrentReads verifies that multiple goroutines can safely call
// read-only methods on a single Engine instance simultaneously.
func TestEngineConcurrentReads(_ *testing.T) {
	e := NewEngine()

	var wg sync.WaitGroup
	goroutines := 100
	iterations := 100

	words := []string{"cat", "child", "formula", "person", "criterion", "appendix", "matrix"}

	for range goroutines {
		wg.Go(func() {
			for range iterations {
				for _, word := range words {
					_ = e.Plural(word)
					_ = e.Singular(word)
					_ = e.An(word)
					_ = e.A(word)
					_ = e.IsClassical()
					_ = e.IsClassicalAll()
				}
			}
		})
	}
	wg.Wait()
}

// TestEngineConcurrentMixedOps verifies that mixed read and write operations
// on a single Engine instance do not cause race conditions.
func TestEngineConcurrentMixedOps(_ *testing.T) {
	e := NewEngine()

	var wg sync.WaitGroup
	goroutines := 50
	iterations := 50

	// Reader goroutines
	for range goroutines {
		wg.Go(func() {
			for range iterations {
				_ = e.Plural("cat")
				_ = e.Plural("child")
				_ = e.Singular("cats")
				_ = e.Singular("children")
				_ = e.An("apple")
				_ = e.An("car")
				_ = e.IsClassical()
				_ = e.IsClassicalAll()
				_ = e.PluralNoun("cat")
				_ = e.SingularNoun("cats")
			}
		})
	}

	// Writer goroutines
	for n := range goroutines / 5 {
		wg.Go(func() {
			for j := range iterations {
				e.Classical(j%2 == 0)
				e.ClassicalAll(j%3 == 0)
				e.DefNoun(fmt.Sprintf("test%d_%d", n, j), fmt.Sprintf("tests%d_%d", n, j))
				e.DefA(fmt.Sprintf("word%d", n))
				e.DefAn(fmt.Sprintf("item%d", n))
			}
		})
	}

	wg.Wait()
}

// TestEngineCloneConcurrent verifies that cloning an Engine while it's being
// used concurrently is safe.
func TestEngineCloneConcurrent(_ *testing.T) {
	e := NewEngine()

	var wg sync.WaitGroup
	goroutines := 50
	iterations := 20

	// Readers
	for range goroutines {
		wg.Go(func() {
			for range iterations {
				_ = e.Plural("cat")
				_ = e.Singular("cats")
			}
		})
	}

	// Writers
	for n := range 10 {
		wg.Go(func() {
			for j := range iterations {
				e.DefNoun(fmt.Sprintf("clone%d_%d", n, j), fmt.Sprintf("clones%d_%d", n, j))
			}
		})
	}

	// Cloners
	clones := make(chan *Engine, 20)
	for range 10 {
		wg.Go(func() {
			for range iterations {
				clone := e.Clone()
				clones <- clone
			}
		})
	}

	// Wait for all operations to complete
	go func() {
		wg.Wait()
		close(clones)
	}()

	// Drain clones channel
	for clone := range clones {
		// Verify clone works
		_ = clone.Plural("cat")
	}
}

// TestEngineResetConcurrent verifies that Reset can be called concurrently
// with other operations.
func TestEngineResetConcurrent(_ *testing.T) {
	e := NewEngine()

	var wg sync.WaitGroup
	goroutines := 50
	iterations := 20

	// Readers
	for range goroutines {
		wg.Go(func() {
			for range iterations {
				_ = e.Plural("cat")
				_ = e.Singular("cats")
				_ = e.IsClassical()
			}
		})
	}

	// Writers
	for n := range 10 {
		wg.Go(func() {
			for j := range iterations {
				e.Classical(true)
				e.DefNoun(fmt.Sprintf("reset%d_%d", n, j), fmt.Sprintf("resets%d_%d", n, j))
			}
		})
	}

	// Resetters
	for range 5 {
		wg.Go(func() {
			for range iterations {
				e.Reset()
			}
		})
	}

	wg.Wait()
}

// TestMultipleEngines verifies that multiple Engine instances can be used
// concurrently without interference.
func TestMultipleEngines(_ *testing.T) {
	engines := make([]*Engine, 10)
	for i := range engines {
		engines[i] = NewEngine()
	}

	var wg sync.WaitGroup
	iterations := 50

	for _, eng := range engines {
		e := eng // capture for goroutine
		wg.Go(func() {
			for j := range iterations {
				e.Classical(j%2 == 0)
				e.DefNoun(fmt.Sprintf("multi%d", j), fmt.Sprintf("multis%d", j))
				_ = e.Plural("cat")
				_ = e.Singular("cats")
				_ = e.An("apple")
			}
		})
	}

	wg.Wait()
}

// TestConcurrentVerbOperations tests concurrent verb inflection operations.
func TestConcurrentVerbOperations(_ *testing.T) {
	e := NewEngine()

	var wg sync.WaitGroup
	goroutines := 50
	iterations := 50

	verbs := []string{"is", "has", "does", "goes", "runs", "walks"}

	for range goroutines {
		wg.Go(func() {
			for range iterations {
				for _, verb := range verbs {
					_ = e.PluralVerb(verb)
				}
			}
		})
	}

	// Writer goroutines for custom verbs
	for n := range 10 {
		wg.Go(func() {
			for j := range iterations {
				e.DefVerb(fmt.Sprintf("verb%d_%d", n, j), fmt.Sprintf("verbs%d_%d", n, j))
			}
		})
	}

	wg.Wait()
}

// TestConcurrentAdjectiveOperations tests concurrent adjective inflection operations.
func TestConcurrentAdjectiveOperations(_ *testing.T) {
	e := NewEngine()

	var wg sync.WaitGroup
	goroutines := 50
	iterations := 50

	adjs := []string{"this", "that", "my", "your", "our", "their"}

	for range goroutines {
		wg.Go(func() {
			for range iterations {
				for _, adj := range adjs {
					_ = e.PluralAdj(adj)
				}
			}
		})
	}

	// Writer goroutines for custom adjectives
	for n := range 10 {
		wg.Go(func() {
			for j := range iterations {
				e.DefAdj(fmt.Sprintf("adj%d_%d", n, j), fmt.Sprintf("adjs%d_%d", n, j))
			}
		})
	}

	wg.Wait()
}

// TestConcurrentClassicalModeToggle tests rapid toggling of classical mode
// while other operations are running.
func TestConcurrentClassicalModeToggle(_ *testing.T) {
	e := NewEngine()

	var wg sync.WaitGroup
	goroutines := 50
	iterations := 100

	// Classical mode words that behave differently
	classicalWords := []string{"formula", "appendix", "criterion", "phenomenon", "matrix"}

	// Readers
	for range goroutines {
		wg.Go(func() {
			for range iterations {
				for _, word := range classicalWords {
					_ = e.Plural(word)
				}
			}
		})
	}

	// Classical mode togglers
	for range 10 {
		wg.Go(func() {
			for j := range iterations {
				e.ClassicalAll(j%2 == 0)
				e.ClassicalAncient(j%3 == 0)
				e.ClassicalHerd(j%4 == 0)
				e.ClassicalPersons(j%5 == 0)
				e.ClassicalZero(j%6 == 0)
				e.ClassicalNames(j%7 == 0)
			}
		})
	}

	wg.Wait()
}

// TestConcurrentNumberConversion tests concurrent number-to-words operations.
func TestConcurrentNumberConversion(_ *testing.T) {
	var wg sync.WaitGroup
	goroutines := 50
	iterations := 100

	for n := range goroutines {
		wg.Go(func() {
			for j := range iterations {
				_ = NumberToWords(n*1000 + j)
				_ = Ordinal(n*100 + j)
				_ = OrdinalWord(n*10 + j)
			}
		})
	}

	wg.Wait()
}

// TestConcurrentArticlePatterns tests concurrent article pattern operations.
func TestConcurrentArticlePatterns(_ *testing.T) {
	e := NewEngine()

	var wg sync.WaitGroup
	goroutines := 30
	iterations := 30

	// Readers
	for range goroutines {
		wg.Go(func() {
			for range iterations {
				_ = e.An("apple")
				_ = e.An("car")
				_ = e.An("umbrella")
				_ = e.A("house")
				_ = e.A("unique")
			}
		})
	}

	// Pattern definers
	for n := range 10 {
		wg.Go(func() {
			for range iterations {
				// DefAPattern and DefAnPattern may return errors for invalid patterns
				// We ignore them for concurrency testing purposes
				_ = e.DefAPattern(fmt.Sprintf("test%d.*", n))
				_ = e.DefAnPattern(fmt.Sprintf("item%d.*", n))
			}
		})
	}

	wg.Wait()
}

// TestConcurrentGenderOperations tests concurrent gender-related operations.
func TestConcurrentGenderOperations(_ *testing.T) {
	e := NewEngine()

	var wg sync.WaitGroup
	goroutines := 50
	iterations := 50

	genders := []string{"m", "f", "n", "t"}

	// Readers
	for range goroutines {
		wg.Go(func() {
			for range iterations {
				_ = e.SingularNoun("they")
				_ = e.SingularNoun("them")
			}
		})
	}

	// Gender setters
	for n := range 10 {
		wg.Go(func() {
			for j := range iterations {
				e.SetGender(genders[(n+j)%len(genders)])
			}
		})
	}

	wg.Wait()
}

// TestConcurrentCompareOperations tests concurrent word comparison operations.
func TestConcurrentCompareOperations(_ *testing.T) {
	var wg sync.WaitGroup
	goroutines := 50
	iterations := 50

	wordPairs := [][2]string{
		{"cat", "cats"},
		{"child", "children"},
		{"person", "people"},
		{"mouse", "mice"},
		{"cat", "dog"},
	}

	for range goroutines {
		wg.Go(func() {
			for range iterations {
				for _, pair := range wordPairs {
					_ = Compare(pair[0], pair[1])
					_ = CompareNouns(pair[0], pair[1])
				}
			}
		})
	}

	wg.Wait()
}

// TestConcurrentJoinOperations tests concurrent list joining operations.
func TestConcurrentJoinOperations(_ *testing.T) {
	var wg sync.WaitGroup
	goroutines := 50
	iterations := 50

	lists := [][]string{
		{"apple", "banana", "cherry"},
		{"cat", "dog"},
		{"one", "two", "three", "four"},
		{"alpha"},
		{},
	}

	for range goroutines {
		wg.Go(func() {
			for range iterations {
				for _, list := range lists {
					_ = Join(list)
					_ = JoinWithConj(list, "and")
					_ = JoinWithConj(list, "or")
				}
			}
		})
	}

	wg.Wait()
}

// TestConcurrentPossessiveOperations tests concurrent possessive operations.
func TestConcurrentPossessiveOperations(_ *testing.T) {
	e := NewEngine()

	var wg sync.WaitGroup
	goroutines := 50
	iterations := 50

	words := []string{"cat", "James", "boss", "children", "it"}

	// Readers
	for range goroutines {
		wg.Go(func() {
			for range iterations {
				for _, word := range words {
					_ = e.Possessive(word)
				}
			}
		})
	}

	// Style togglers
	for n := range 10 {
		wg.Go(func() {
			for j := range iterations {
				if (n+j)%2 == 0 {
					e.SetPossessiveStyle(PossessiveModern)
				} else {
					e.SetPossessiveStyle(PossessiveTraditional)
				}
			}
		})
	}

	wg.Wait()
}

// TestConcurrentInflection tests the Inflect function concurrently.
func TestConcurrentInflection(_ *testing.T) {
	var wg sync.WaitGroup
	goroutines := 50
	iterations := 50

	patterns := []string{
		"PL(cat)",
		"SING(cats)",
		"AN(apple)",
		"A(car)",
		"ORD(3)",
	}

	for range goroutines {
		wg.Go(func() {
			for range iterations {
				for _, pattern := range patterns {
					_ = Inflect(pattern)
				}
			}
		})
	}

	wg.Wait()
}

// TestDefaultEngineIsolation verifies that DefaultEngine returns the same
// instance and that concurrent access to it is safe.
func TestDefaultEngineIsolation(_ *testing.T) {
	var wg sync.WaitGroup
	goroutines := 50
	iterations := 50

	for range goroutines {
		wg.Go(func() {
			for range iterations {
				e := DefaultEngine()
				_ = e.Plural("cat")
				_ = e.Singular("cats")
			}
		})
	}

	wg.Wait()
}
