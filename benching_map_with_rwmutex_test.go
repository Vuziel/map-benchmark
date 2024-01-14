package bench

import (
	"sync"
	"testing"
)

type StandardMapRWWithMutex struct {
	rwMutex sync.RWMutex
	testMap map[int]int
}

func (s *StandardMapRWWithMutex) Insert(i int) {
	s.rwMutex.Lock()
	s.testMap[i] = i
	s.rwMutex.Unlock()
}

func (s *StandardMapRWWithMutex) Select(key int) int {
	s.rwMutex.RLock()
	val := s.testMap[key]
	s.rwMutex.RUnlock()

	return val
}

func (s *StandardMapRWWithMutex) Delete(key int) {
	s.rwMutex.Lock()
	delete(s.testMap, key)
	s.rwMutex.Unlock()
}

func insertXStandardMapRWWithMutex(x int, b *testing.B) {
	testMap := StandardMapRWWithMutex{testMap: make(map[int]int, x)}
	b.ResetTimer()
	for i := 0; i < x; i++ {
		testMap.Insert(i)
	}
}

func BenchmarkInsertStandardMapRWWithMutex_1_000_000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXStandardMapRWWithMutex(1_000_000, b)
	}
}

func BenchmarkInsertStandardMapRWWithMutex_100_000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXStandardMapRWWithMutex(100_000, b)
	}
}

func BenchmarkInsertStandardMapRWWithMutex_10_000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXStandardMapRWWithMutex(10_000, b)
	}
}

func BenchmarkInsertStandardMapRWWithMutex_1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXStandardMapRWWithMutex(1000, b)
	}
}

func BenchmarkInsertStandardMapRWWithMutex_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXStandardMapRWWithMutex(100, b)
	}
}

func selectXStandardMapRWWithMutex(x int, b *testing.B) {
	testMap := StandardMapRWWithMutex{testMap: make(map[int]int, x)}
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

func BenchmarkSelectStandardMapRWWithMutex_1_000_000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		selectXStandardMapRWWithMutex(1_000_000, b)
	}
}

func BenchmarkSelectStandardMapRWWithMutex_100_000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		selectXStandardMapRWWithMutex(100_000, b)
	}
}

func BenchmarkSelectStandardMapRWWithMutex_10_000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		selectXStandardMapRWWithMutex(10_000, b)
	}
}

func BenchmarkSelectStandardMapRWWithMutex_1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		selectXStandardMapRWWithMutex(1000, b)
	}
}

func BenchmarkSelectStandardMapRWWithMutex_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		selectXStandardMapRWWithMutex(100, b)
	}
}

func deleteXStandardMapRWWithMutex(x int, b *testing.B) {
	testMap := StandardMapRWWithMutex{testMap: make(map[int]int, x)}
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

func BenchmarkDeleteStandardMapRWWithMutex_1_000_0000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deleteXStandardMapRWWithMutex(1_000_000, b)
	}
}

func BenchmarkDeleteStandardMapRWWithMutex_100_000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deleteXStandardMapRWWithMutex(100_000, b)
	}
}

func BenchmarkDeleteStandardMapRWWithMutex_10_000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deleteXStandardMapRWWithMutex(10_000, b)
	}
}

func BenchmarkDeleteStandardMapRWWithMutex_1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deleteXStandardMapRWWithMutex(1000, b)
	}
}

func BenchmarkDeleteStandardMapRWWithMutex_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deleteXStandardMapRWWithMutex(100, b)
	}
}
