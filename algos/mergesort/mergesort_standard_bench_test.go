package mergesort

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkStdMergeSort_20(b *testing.B) {
	sample := generateRandomArrayInt(20, 100)
	//b.Logf("benchmarking standard merge sort with size: %d", len(sample))
	benchmarkSort(sample, b)
}

func BenchmarkStdMergeSort_50(b *testing.B) {
	sample := generateRandomArrayInt(50, 100)
	//b.Logf("benchmarking standard merge sort with size: %d", len(sample))
	benchmarkSort(sample, b)
}

func BenchmarkStdMergeSort_80(b *testing.B) {
	sample := generateRandomArrayInt(80, 300)
	//b.Logf("benchmarking standard merge sort with size: %d", len(sample))
	benchmarkSort(sample, b)
}

func BenchmarkStdMergeSort_200(b *testing.B) {
	sample := generateRandomArrayInt(200, 500)
	//b.Logf("benchmarking standard merge sort with size: %d", len(sample))
	benchmarkSort(sample, b)
}

func BenchmarkStdMergeSort_500(b *testing.B) {
	sample := generateRandomArrayInt(500, 1000)
	//b.Logf("benchmarking standard merge sort with size: %d", len(sample))
	benchmarkSort(sample, b)
}

func BenchmarkStdMergeSort_5000(b *testing.B) {
	sample := generateRandomArrayInt(5000, 10000)
	//b.Logf("benchmarking standard merge sort with size: %d", len(sample))
	benchmarkSort(sample, b)
}

func BenchmarkStdMergeSort_100000(b *testing.B) {
	sample := generateRandomArrayInt(100000, 400000)
	//b.Logf("benchmarking standard merge sort with size: %d", len(sample))
	benchmarkSort(sample, b)
}

func generateRandomArrayInt(size int, limit int) []int {
	var sample []int
	seed := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(seed)
	for i := 0; i < size; i++ {
		sample = append(sample, rnd.Intn(limit))
	}
	return sample
}

func benchmarkSort(sample []int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Sort(sample)
	}
}
