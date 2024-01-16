package bench

import (
	"fmt"
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

func BenchmarkInsertStandardMapWithMutexMap(b *testing.B) {
	benchmarks := []struct {
		elems int
	}{
		{elems: 100},
		{elems: 1_000},
		{elems: 10_000},
		{elems: 100_000},
		{elems: 1_000_000},
	}

	for _, bench := range benchmarks {
		testMap := StandardMapWithMutex{testMap: make(map[int]int, bench.elems)}
		b.ResetTimer()

		b.Run(fmt.Sprintf("standard map mutex with amount of elements %d", bench.elems), func(t *testing.B) {
			for i := 0; i < b.N; i++ {
				insertStandardMapWithMutex(bench.elems, &testMap)
			}
		})
	}
}

func insertStandardMapWithMutex(x int, testMap *StandardMapWithMutex) {
	for i := 0; i < x; i++ {
		testMap.Insert(i)
	}
}

func BenchmarkSelectStandardMapWithMutexMap(b *testing.B) {
	benchmarks := []struct {
		elems int
	}{
		{elems: 100},
		{elems: 1_000},
		{elems: 10_000},
		{elems: 100_000},
		{elems: 1_000_000},
	}

	for _, bench := range benchmarks {
		testMap := StandardMapWithMutex{testMap: make(map[int]int, bench.elems)}
		for i := 0; i < bench.elems; i++ {
			testMap.Insert(i)
		}
		b.ResetTimer()

		b.Run(fmt.Sprintf("standard map mutex with amount of elements %d", bench.elems), func(t *testing.B) {
			for i := 0; i < b.N; i++ {
				selectStandardMapWithMutex(bench.elems, &testMap)
			}
		})
	}
}

func selectStandardMapWithMutex(x int, testMap *StandardMapWithMutex) {
	for i := 0; i < x; i++ {
		j := testMap.Select(i)
		if j != 0 {

		}
	}
}

func BenchmarkDeleteStandardMapWithMutexMap(b *testing.B) {
	benchmarks := []struct {
		elems int
	}{
		{elems: 100},
		{elems: 1_000},
		{elems: 10_000},
		{elems: 100_000},
		{elems: 1_000_000},
	}

	for _, bench := range benchmarks {
		testMap := StandardMapWithMutex{testMap: make(map[int]int, bench.elems)}
		for i := 0; i < bench.elems; i++ {
			testMap.Insert(i)
		}
		b.ResetTimer()

		b.Run(fmt.Sprintf("standard map mutex with amount of elements %d", bench.elems), func(t *testing.B) {
			for i := 0; i < b.N; i++ {
				deleteStandardMapWithMutex(bench.elems, &testMap)
			}
		})
	}
}

func deleteStandardMapWithMutex(x int, testMap *StandardMapWithMutex) {
	for i := 0; i < x; i++ {
		testMap.Delete(i)
	}
}
