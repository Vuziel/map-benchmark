package bench

import (
	"fmt"
	"testing"

	"github.com/cornelk/hashmap"
)

func BenchmarkInsertCornelkHashmap(b *testing.B) {
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
		testMap := hashmap.NewSized[int, int](uintptr(bench.elems))
		b.ResetTimer()

		b.Run(fmt.Sprintf("input map with amount of elements %d", bench.elems), func(t *testing.B) {
			for i := 0; i < b.N; i++ {
				insertCornelkHashmap(bench.elems, testMap)
			}
		})
	}
}

func insertCornelkHashmap(elems int, testMap *hashmap.Map[int, int]) {
	for i := 0; i < elems; i++ {
		testMap.Insert(i, i)
	}
}

func BenchmarkSelectCornelkHashmap(b *testing.B) {
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
		testMap := hashmap.NewSized[int, int](uintptr(bench.elems))
		for i := 0; i < bench.elems; i++ {
			testMap.Insert(i, i)
		}
		b.ResetTimer()

		b.Run(fmt.Sprintf("input map with amount of elements %d", bench.elems), func(t *testing.B) {
			for i := 0; i < b.N; i++ {
				selectCornelkHashmap(bench.elems, testMap)
			}
		})
	}
}

func selectCornelkHashmap(elems int, testMap *hashmap.Map[int, int]) {
	for i := 0; i < elems; i++ {
		_, ok := testMap.Get(i)
		if ok != true {

		}
	}
}

func BenchmarkDeleteCornelkHashmap(b *testing.B) {
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
		testMap := hashmap.NewSized[int, int](uintptr(bench.elems))
		for i := 0; i < bench.elems; i++ {
			testMap.Insert(i, i)
		}
		b.ResetTimer()

		b.Run(fmt.Sprintf("input map with amount of elements %d", bench.elems), func(t *testing.B) {
			for i := 0; i < b.N; i++ {
				deleteCornelkHashmap(bench.elems, testMap)
			}
		})
	}
}

func deleteCornelkHashmap(elems int, testMap *hashmap.Map[int, int]) {
	for i := 0; i < elems; i++ {
		ok := testMap.Del(i)
		if ok != true {

		}
	}
}
