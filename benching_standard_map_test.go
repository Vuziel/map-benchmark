package bench

import (
	"testing"
)

// todo -> table tests

// todo -> https://github.com/orcaman/concurrent-map

// todo - hash table исполь., метод цепочек

func insertXStandardMap(x int, b *testing.B) {
	testMap := make(map[int]int, x)
	b.ResetTimer()
	for i := 0; i < x; i++ {
		testMap[i] = i
	}
}

func BenchmarkInsertStandardMap_1_000_000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXStandardMap(1_000_000, b)
	}
}

func BenchmarkInsertStandardMap_100_000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXStandardMap(100_000, b)
	}
}

func BenchmarkInsertStandardMap_10_000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXStandardMap(10_000, b)
	}
}

func BenchmarkInsertStandardMap_1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXStandardMap(1000, b)
	}
}

func BenchmarkInsertStandardMap_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXStandardMap(100, b)
	}
}

func selectXStandardMap(x int, b *testing.B) {
	testMap := make(map[int]int, x)
	for i := 0; i < x; i++ {
		testMap[i] = i
	}
	var holder int
	b.ResetTimer()

	for i := 0; i < x; i++ {
		holder = testMap[i]
	}

	if holder != 0 {
	}
}

func BenchmarkSelectStandardMap_1_000_000(b *testing.B) {
	// todo map create, reset time
	for i := 0; i < b.N; i++ {
		selectXStandardMap(1_000_000, b)
	}
}

func BenchmarkSelectStandardMap_100_000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		selectXStandardMap(100_000, b)
	}
}

func BenchmarkSelectStandardMap_10_000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		selectXStandardMap(10_000, b)
	}
}

func BenchmarkSelectStandardMap_1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		selectXStandardMap(1000, b)
	}
}

func BenchmarkSelectStandardMap_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		selectXStandardMap(100, b)
	}
}

func deleteXStandardMap(x int, b *testing.B) {
	testMap := make(map[int]int, x)
	for i := 0; i < x; i++ {
		testMap[i] = i
	}
	b.ResetTimer()

	for i := 0; i < x; i++ {
		delete(testMap, i)
	}
}

func BenchmarkDeleteStandardMap_1_000_000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deleteXStandardMap(1_000_000, b)
	}
}

func BenchmarkDeleteStandardMap_100_000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		selectXStandardMap(100_000, b)
	}
}

func BenchmarkDeleteStandardMap_10_000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		selectXStandardMap(10_000, b)
	}
}

func BenchmarkDeleteStandardMap_1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		selectXStandardMap(1000, b)
	}
}

func BenchmarkDeleteStandardMap_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		selectXStandardMap(100, b)
	}
}
