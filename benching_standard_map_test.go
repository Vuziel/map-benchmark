package bench

import (
	"fmt"
	"testing"
)

func BenchmarkInsertStandardMap(b *testing.B) {
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
		bench.mP = make(map[int]int, bench.elems)
		b.ResetTimer()

		b.Run(fmt.Sprintf("input map with amount of elements %d", bench.elems), func(t *testing.B) {
			for i := 0; i < b.N; i++ {
				insertStandardMap(bench.elems, bench.mP)
			}
		})
	}
}

func insertStandardMap(elems int, testMap map[int]int) {
	for i := 0; i < elems; i++ {
		testMap[i] = i
	}
}

func BenchmarkSelectStandardMap(b *testing.B) {
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
		bench.mP = make(map[int]int, bench.elems)
		for i := 0; i < bench.elems; i++ {
			bench.mP[i] = i
		}
		b.ResetTimer()

		b.Run(fmt.Sprintf("map with amount of elements %d", bench.elems), func(t *testing.B) {
			for i := 0; i < b.N; i++ {
				selectStandardMap(bench.elems, bench.mP)
			}
		})
	}
}

func selectStandardMap(elems int, testMap map[int]int) {
	for i := 0; i < elems; i++ {
		j := testMap[i]

		if j != 0 {
		}
	}
}

func BenchmarkDeleteStandardMap(b *testing.B) {
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
		bench.mP = make(map[int]int, bench.elems)
		for i := 0; i < bench.elems; i++ {
			bench.mP[i] = i
		}
		b.ResetTimer()

		b.Run(fmt.Sprintf("map with amount of elements %d", bench.elems), func(t *testing.B) {
			for i := 0; i < b.N; i++ {
				deleteStandardMap(bench.elems, bench.mP)
			}
		})
	}
}

func deleteStandardMap(x int, testMap map[int]int) {
	for i := 0; i < x; i++ {
		delete(testMap, i)
	}
}
