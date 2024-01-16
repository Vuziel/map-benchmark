package bench

import (
	"fmt"
	"sync"
	"testing"
)

func storeSyncMap(x int, testMap *sync.Map) {
	for i := 0; i < x; i++ {
		testMap.Store(i, i)
	}
}

func loadSyncMap(x int, testMap *sync.Map) {
	for i := 0; i < x; i++ {
		j, _ := testMap.Load(i)

		if j != 0 {
		}
	}
}

func deleteSyncMap(x int, testMap *sync.Map) {
	for i := 0; i < x; i++ {
		testMap.Delete(i)
	}
}

func BenchmarkInsertSyncMap(b *testing.B) {
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
		var testMap sync.Map
		b.ResetTimer()

		b.Run(fmt.Sprintf("sync map with amount of elements %d", bench.elems), func(t *testing.B) {
			for i := 0; i < b.N; i++ {
				storeSyncMap(bench.elems, &testMap)
			}
		})
	}
}

func BenchmarkSelectSyncMap(b *testing.B) {
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
		var testMap sync.Map
		for i := 0; i < bench.elems; i++ {
			testMap.Store(i, i)
		}
		b.ResetTimer()

		b.Run(fmt.Sprintf("sync map with amount of elements %d", bench.elems), func(t *testing.B) {
			for i := 0; i < b.N; i++ {
				loadSyncMap(bench.elems, &testMap)
			}
		})
	}
}

func BenchmarkDeleteSyncMap(b *testing.B) {
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
		var testMap sync.Map
		for i := 0; i < bench.elems; i++ {
			testMap.Store(i, i)
		}
		b.ResetTimer()

		b.Run(fmt.Sprintf("sync map with amount of elements %d", bench.elems), func(t *testing.B) {
			for i := 0; i < b.N; i++ {
				deleteSyncMap(bench.elems, &testMap)
			}
		})
	}
}
