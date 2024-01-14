package bench

import (
	"sync"
	"testing"
)

func storeSyncMap(x int, b *testing.B) {
	var testMap sync.Map
	b.ResetTimer()
	for i := 0; i < x; i++ {
		testMap.Store(i, i)
	}
}

func loadSyncMap(x int, b *testing.B) {
	var testMap sync.Map
	for i := 0; i < x; i++ {
		testMap.Store(i, i)
	}

	var holder any
	b.ResetTimer()

	for i := 0; i < x; i++ {
		holder, _ = testMap.Load(i)
	}

	if holder != 0 {
	}
}

func deleteSyncMap(x int, b *testing.B) {
	var testMap sync.Map
	for i := 0; i < x; i++ {
		testMap.Store(i, i)
	}

	b.ResetTimer()

	for i := 0; i < x; i++ {
		testMap.Delete(i)
	}
}

func BenchmarkInsertSyncMap_1_000_000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		storeSyncMap(1_000_000, b)
	}
}

func BenchmarkInsertSyncMap_100_000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		storeSyncMap(100_000, b)
	}
}

func BenchmarkInsertSyncMap_10_000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		storeSyncMap(10_000, b)
	}
}

func BenchmarkInsertSyncMap_1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		storeSyncMap(1000, b)
	}
}

func BenchmarkInsertSyncMap_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		storeSyncMap(100, b)
	}
}

func BenchmarkSelectSyncMap_1_000_000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		loadSyncMap(1_000_000, b)
	}
}

func BenchmarkSelectSyncMap_100_000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		loadSyncMap(100_000, b)
	}
}

func BenchmarkSelectSyncMap_10_000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		loadSyncMap(10_000, b)
	}
}

func BenchmarkSelectSyncMap_1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		loadSyncMap(1000, b)
	}
}

func BenchmarkSelectSyncMap_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		loadSyncMap(100, b)
	}
}

func BenchmarkDeleteSyncMap_1_000_000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deleteSyncMap(1_000_000, b)
	}
}

func BenchmarkDeleteSyncMap_100_000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deleteSyncMap(100_000, b)
	}
}

func BenchmarkDeleteSyncMap_10_000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deleteSyncMap(10_000, b)
	}
}

func BenchmarkDeleteSyncMap_1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deleteSyncMap(1000, b)
	}
}

func BenchmarkDeleteSyncMap_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deleteSyncMap(100, b)
	}
}
