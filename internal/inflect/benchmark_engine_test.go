package inflect

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

// =============================================================================
// Serial Benchmarks - Measure baseline single-goroutine performance
// =============================================================================

// BenchmarkPluralSerial measures package-level Plural performance in serial.
func BenchmarkPluralSerial(b *testing.B) {
	for b.Loop() {
		Plural("cat")
	}
}

// BenchmarkEnginePluralSerial measures Engine.Plural performance in serial
// with a dedicated Engine instance (no contention from other callers).
func BenchmarkEnginePluralSerial(b *testing.B) {
	e := NewEngine()
	for b.Loop() {
		e.Plural("cat")
	}
}

// BenchmarkSingularSerial measures package-level Singular performance in serial.
func BenchmarkSingularSerial(b *testing.B) {
	for b.Loop() {
		Singular("cats")
	}
}

// BenchmarkAnSerial measures package-level An performance in serial.
func BenchmarkAnSerial(b *testing.B) {
	for b.Loop() {
		An("apple")
	}
}

// =============================================================================
// Parallel Benchmarks - Measure read performance under contention
// =============================================================================

// BenchmarkPluralParallel measures package-level Plural with parallel access.
func BenchmarkPluralParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Plural("cat")
		}
	})
}

// BenchmarkEnginePluralParallel measures Engine.Plural with parallel access
// to a shared Engine instance.
func BenchmarkEnginePluralParallel(b *testing.B) {
	e := NewEngine()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			e.Plural("cat")
		}
	})
}

// BenchmarkSingularParallel measures package-level Singular with parallel access.
func BenchmarkSingularParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Singular("cats")
		}
	})
}

// BenchmarkAnParallel measures package-level An with parallel access.
func BenchmarkAnParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			An("apple")
		}
	})
}

// =============================================================================
// Write Operation Benchmarks - Measure write lock overhead
// =============================================================================

// BenchmarkClassicalToggle measures the cost of toggling classical mode.
func BenchmarkClassicalToggle(b *testing.B) {
	e := NewEngine()
	for i := range b.N {
		e.Classical(i%2 == 0)
	}
}

// BenchmarkDefNoun measures the cost of defining custom nouns.
func BenchmarkDefNoun(b *testing.B) {
	e := NewEngine()
	for i := range b.N {
		e.DefNoun(fmt.Sprintf("word%d", i), fmt.Sprintf("words%d", i))
	}
}

// =============================================================================
// Contention Scenarios - Measure lock overhead under realistic workloads
// =============================================================================

// BenchmarkHighContentionReads measures read performance with many parallel readers.
// This tests the RWMutex scalability for read-heavy workloads.
func BenchmarkHighContentionReads(b *testing.B) {
	e := NewEngine()
	words := []string{"cat", "dog", "child", "person", "formula", "matrix", "appendix"}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			word := words[i%len(words)]
			_ = e.Plural(word)
			_ = e.Singular(word + "s")
			_ = e.An(word)
			i++
		}
	})
}

// BenchmarkMixedContentionReadWrite measures performance with occasional writes
// among many reads. This simulates realistic usage where reads dominate but
// occasional configuration changes occur.
//
// Ratio: approximately 100 reads per 1 write.
func BenchmarkMixedContentionReadWrite(b *testing.B) {
	e := NewEngine()
	words := []string{"cat", "dog", "child", "person", "formula", "matrix", "appendix"}
	var writeCounter atomic.Int64

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			// Every 100th operation is a write
			if i%100 == 0 {
				n := writeCounter.Add(1)
				e.DefNoun(fmt.Sprintf("bench%d", n), fmt.Sprintf("benchs%d", n))
			} else {
				word := words[i%len(words)]
				_ = e.Plural(word)
			}
			i++
		}
	})
}

// =============================================================================
// Isolated vs Shared Engine Comparison
// =============================================================================

// BenchmarkIsolatedEngines measures performance when each goroutine has its
// own Engine instance (no lock contention).
func BenchmarkIsolatedEngines(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		e := NewEngine()
		for pb.Next() {
			e.Plural("cat")
		}
	})
}

// BenchmarkSharedEngine measures performance when all goroutines share a
// single Engine instance (maximum lock contention for reads).
func BenchmarkSharedEngine(b *testing.B) {
	e := NewEngine()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			e.Plural("cat")
		}
	})
}

// =============================================================================
// Read-Heavy vs Write-Heavy Workloads
// =============================================================================

// BenchmarkReadHeavyWorkload simulates a read-heavy workload (99% reads).
func BenchmarkReadHeavyWorkload(b *testing.B) {
	e := NewEngine()
	var writeCounter atomic.Int64

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			if i%100 == 0 {
				n := writeCounter.Add(1)
				e.Classical(n%2 == 0)
			} else {
				_ = e.Plural("cat")
			}
			i++
		}
	})
}

// BenchmarkWriteHeavyWorkload simulates a write-heavy workload (50% writes).
// This is an unrealistic but useful stress test for lock contention.
func BenchmarkWriteHeavyWorkload(b *testing.B) {
	e := NewEngine()
	var writeCounter atomic.Int64

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			if i%2 == 0 {
				n := writeCounter.Add(1)
				e.DefNoun(fmt.Sprintf("heavy%d", n), fmt.Sprintf("heavys%d", n))
			} else {
				_ = e.Plural("cat")
			}
			i++
		}
	})
}

// =============================================================================
// Clone Overhead Benchmark
// =============================================================================

// BenchmarkClone measures the cost of cloning an Engine.
func BenchmarkClone(b *testing.B) {
	e := NewEngine()
	// Add some custom definitions to make the clone more realistic
	for i := range 100 {
		e.DefNoun(fmt.Sprintf("custom%d", i), fmt.Sprintf("customs%d", i))
	}

	for b.Loop() {
		_ = e.Clone()
	}
}

// BenchmarkCloneParallel measures clone performance under contention.
func BenchmarkCloneParallel(b *testing.B) {
	e := NewEngine()
	for i := range 100 {
		e.DefNoun(fmt.Sprintf("custom%d", i), fmt.Sprintf("customs%d", i))
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = e.Clone()
		}
	})
}

// =============================================================================
// Real-World Pattern Benchmarks
// =============================================================================

// BenchmarkTypicalWorkload simulates a typical mixed workload pattern:
// - Multiple read goroutines doing continuous pluralization.
// - Occasional configuration changes.
// - Periodic cloning for isolated operations.
func BenchmarkTypicalWorkload(b *testing.B) {
	e := NewEngine()
	words := []string{"cat", "dog", "child", "person", "formula", "matrix", "appendix",
		"mouse", "house", "box", "bus", "class", "hero", "potato"}
	var counter atomic.Int64

	b.ResetTimer()

	var wg sync.WaitGroup
	goroutines := 8
	opsPerGoroutine := max(b.N/goroutines, 1)

	for range goroutines {
		wg.Go(func() {
			for range opsPerGoroutine {
				n := counter.Add(1)
				switch n % 1000 {
				case 0:
					// Rare: toggle classical mode
					e.Classical(n%2000 == 0)
				case 500:
					// Rare: define a custom noun
					e.DefNoun(fmt.Sprintf("typical%d", n), fmt.Sprintf("typicals%d", n))
				default:
					// Common: read operations
					word := words[int(n)%len(words)]
					_ = e.Plural(word)
				}
			}
		})
	}

	wg.Wait()
}
