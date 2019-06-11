package mergesort

import (
	"testing"
)

func BenchmarkConcurrentMergeSort_20(b *testing.B) {
	sample := generateRandomArrayInt(20, 100)
	benchmarkConcurrentSort(sample, b)
}

func BenchmarkConcurrentMergeSort_50(b *testing.B) {
	sample := generateRandomArrayInt(50, 100)
	benchmarkConcurrentSort(sample, b)
}

func BenchmarkConcurrentMergeSort_80(b *testing.B) {
	sample := generateRandomArrayInt(80, 300)
	benchmarkConcurrentSort(sample, b)
}

func BenchmarkConcurrentMergeSort_200(b *testing.B) {
	sample := generateRandomArrayInt(200, 500)
	benchmarkConcurrentSort(sample, b)
}

func BenchmarkConcurrentMergeSort_500(b *testing.B) {
	sample := generateRandomArrayInt(500, 1000)
	benchmarkConcurrentSort(sample, b)
}

func BenchmarkConcurrentMergeSort_5000(b *testing.B) {
	sample := generateRandomArrayInt(5000, 10000)
	benchmarkConcurrentSort(sample, b)
}

func BenchmarkConcurrentMergeSort_100000(b *testing.B) {
	sample := generateRandomArrayInt(100000, 400000)
	benchmarkConcurrentSort(sample, b)
}

func benchmarkConcurrentSort(sample []int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		SortConcurrent(sample)
	}
}
