package bench

import (
	"fmt"
	"strconv"
	"testing"

	cmap "github.com/orcaman/concurrent-map"
)

func BenchmarkInsertConcurrentMap(b *testing.B) {
	benchmarks := []struct {
		elems int
		mP    map[int]int
	}{
		{elems: 100},
		{elems: 1_000},
		{elems: 10_000},
		{elems: 100_000},
		{elems: 1_000_000},
	}

	for _, bench := range benchmarks {
		testMap := cmap.New()
		b.ResetTimer()

		b.Run(fmt.Sprintf("input map with amount of elements %d", bench.elems), func(t *testing.B) {
			for i := 0; i < b.N; i++ {
				insertConcurrentMap(bench.elems, &testMap)
			}
		})
	}
}

func insertConcurrentMap(elems int, testMap *cmap.ConcurrentMap) {
	for i := 0; i < elems; i++ {
		testMap.Set(strconv.Itoa(i), i)
	}
}

func BenchmarkSelectConcurrentMap(b *testing.B) {
	benchmarks := []struct {
		elems int
		mP    map[int]int
	}{
		{elems: 100},
		{elems: 1_000},
		{elems: 10_000},
		{elems: 100_000},
		{elems: 1_000_000},
	}

	for _, bench := range benchmarks {
		testMap := cmap.New()
		for i := 0; i < bench.elems; i++ {
			testMap.Set(strconv.Itoa(i), i)
		}
		b.ResetTimer()

		b.Run(fmt.Sprintf("input map with amount of elements %d", bench.elems), func(t *testing.B) {
			for i := 0; i < b.N; i++ {
				selectConcurrentMap(bench.elems, &testMap)
			}
		})
	}
}

func selectConcurrentMap(elems int, testMap *cmap.ConcurrentMap) {
	for i := 0; i < elems; i++ {
		_, ok := testMap.Get(strconv.Itoa(i))
		if ok != true {

		}
	}
}

func BenchmarkDeleteConcurrentMap(b *testing.B) {
	benchmarks := []struct {
		elems int
		mP    map[int]int
	}{
		{elems: 100},
		{elems: 1_000},
		{elems: 10_000},
		{elems: 100_000},
		{elems: 1_000_000},
	}

	for _, bench := range benchmarks {
		testMap := cmap.New()
		for i := 0; i < bench.elems; i++ {
			testMap.Set(strconv.Itoa(i), i)
		}
		b.ResetTimer()

		b.Run(fmt.Sprintf("input map with amount of elements %d", bench.elems), func(t *testing.B) {
			for i := 0; i < b.N; i++ {
				deleteConcurrentMap(bench.elems, &testMap)
			}
		})
	}
}

func deleteConcurrentMap(elems int, testMap *cmap.ConcurrentMap) {
	for i := 0; i < elems; i++ {
		testMap.Remove(strconv.Itoa(i))
	}
}
