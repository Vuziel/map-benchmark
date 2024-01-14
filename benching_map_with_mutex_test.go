package bench

import (
	"sync"
	"testing"
)

type StandardMapWithMutex struct {
	mu      sync.Mutex
	testMap map[int]int
}

func (s *StandardMapWithMutex) Insert(i int) {
	s.mu.Lock()
	s.testMap[i] = i
	s.mu.Unlock()
}

func (s *StandardMapWithMutex) Select(key int) int {
	s.mu.Lock()
	val := s.testMap[key]

	s.mu.Unlock()

	return val
}

func (s *StandardMapWithMutex) Delete(key int) {
	s.mu.Lock()
	delete(s.testMap, key)
	s.mu.Unlock()
}

func insertXStandardMapWithMutex(x int, b *testing.B) {
	testMap := StandardMapWithMutex{testMap: make(map[int]int, x)}
	b.ResetTimer()
	for i := 0; i < x; i++ {
		testMap.Insert(i)
	}
}

func BenchmarkInsertStandardMapWithMutex_1_000_000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXStandardMapWithMutex(1_000_000, b)
	}
}

func BenchmarkInsertStandardMapWithMutex_100_000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXStandardMapWithMutex(100_000, b)
	}
}

func BenchmarkInsertStandardMapWithMutex_10_000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXStandardMapWithMutex(10_000, b)
	}
}

func BenchmarkInsertStandardMapWithMutex_1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXStandardMapWithMutex(1000, b)
	}
}

func BenchmarkInsertStandardMapWithMutex_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXStandardMapWithMutex(100, b)
	}
}

func selectXStandardMapWithMutex(x int, b *testing.B) {
	testMap := StandardMapWithMutex{testMap: make(map[int]int, x)}
	for i := 0; i < x; i++ {
		testMap.Insert(i)
	}

	var holder int
	b.ResetTimer()

	for i := 0; i < x; i++ {
		holder = testMap.Select(i)
	}

	if holder != 0 {
	}
}

func BenchmarkSelectStandardMapWithMutex_1_000_000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		selectXStandardMapWithMutex(1_000_000, b)
	}
}

func BenchmarkSelectStandardMapWithMutex_100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		selectXStandardMapWithMutex(100_000, b)
	}
}

func BenchmarkSelectStandardMapWithMutex_10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		selectXStandardMapWithMutex(10_000, b)
	}
}

func BenchmarkSelectStandardMapWithMutex_1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		selectXStandardMapWithMutex(1000, b)
	}
}

func BenchmarkSelectStandardMapWithMutex_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		selectXStandardMapWithMutex(100, b)
	}
}

func deleteXStandardMapWithMutex(x int, b *testing.B) {
	testMap := StandardMapWithMutex{testMap: make(map[int]int, x)}
	for i := 0; i < x; i++ {
		testMap.Insert(i)
	}

	var holder int
	b.ResetTimer()

	for i := 0; i < x; i++ {
		testMap.Delete(i)
	}

	if holder != 0 {
	}
}

func BenchmarkDeleteStandardMapWithMutex_1_000_000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deleteXStandardMapWithMutex(1_000_000, b)
	}
}

func BenchmarkDeleteStandardMapWithMutex_100_000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deleteXStandardMapWithMutex(100_000, b)
	}
}

func BenchmarkDeleteStandardMapWithMutex_10_000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deleteXStandardMapWithMutex(10_000, b)
	}
}

func BenchmarkDeleteStandardMapWithMutex_1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deleteXStandardMapWithMutex(1000, b)
	}
}

func BenchmarkDeleteStandardMapWithMutex_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deleteXStandardMapWithMutex(100, b)
	}
}
