package bench

import (
	"fmt"
	"sync"
	"testing"
)

type StandardMapWithRWMutex struct {
	rwMutex sync.RWMutex
	testMap map[int]int
}

func (s *StandardMapWithRWMutex) Insert(i int) {
	s.rwMutex.Lock()
	s.testMap[i] = i
	s.rwMutex.Unlock()
}

func (s *StandardMapWithRWMutex) Select(key int) int {
	s.rwMutex.RLock()
	val := s.testMap[key]
	s.rwMutex.RUnlock()

	return val
}

func (s *StandardMapWithRWMutex) Delete(key int) {
	s.rwMutex.Lock()
	delete(s.testMap, key)
	s.rwMutex.Unlock()
}

func BenchmarkInsertStandardMapWithRWMutexMap(b *testing.B) {
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
		testMap := StandardMapWithRWMutex{testMap: make(map[int]int, bench.elems)}
		b.ResetTimer()

		b.Run(fmt.Sprintf("standard map RWMutex with amount of elements %d", bench.elems), func(t *testing.B) {
			for i := 0; i < b.N; i++ {
				insertStandardMapWithRWMutex(bench.elems, &testMap)
			}
		})
	}
}

func insertStandardMapWithRWMutex(x int, testMap *StandardMapWithRWMutex) {
	for i := 0; i < x; i++ {
		testMap.Insert(i)
	}
}

func BenchmarkSelectStandardMapWithRWMutexMap(b *testing.B) {
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
		testMap := StandardMapWithRWMutex{testMap: make(map[int]int, bench.elems)}
		for i := 0; i < bench.elems; i++ {
			testMap.Insert(i)
		}
		b.ResetTimer()

		b.Run(fmt.Sprintf("standard map RWMutex with amount of elements %d", bench.elems), func(t *testing.B) {
			for i := 0; i < b.N; i++ {
				selectStandardMapWithRWMutex(bench.elems, &testMap)
			}
		})
	}
}

func selectStandardMapWithRWMutex(x int, testMap *StandardMapWithRWMutex) {
	for i := 0; i < x; i++ {
		j := testMap.Select(i)
		if j != 0 {

		}
	}
}

func BenchmarkDeleteStandardMapWithRWMutexMap(b *testing.B) {
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
		testMap := StandardMapWithRWMutex{testMap: make(map[int]int, bench.elems)}
		for i := 0; i < bench.elems; i++ {
			testMap.Insert(i)
		}
		b.ResetTimer()

		b.Run(fmt.Sprintf("standard map RWMutex with amount of elements %d", bench.elems), func(t *testing.B) {
			for i := 0; i < b.N; i++ {
				deleteStandardMapWithRWMutex(bench.elems, &testMap)
			}
		})
	}
}

func deleteStandardMapWithRWMutex(x int, testMap *StandardMapWithRWMutex) {
	for i := 0; i < x; i++ {
		testMap.Delete(i)
	}
}
