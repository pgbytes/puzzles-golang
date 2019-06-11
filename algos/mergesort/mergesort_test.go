package mergesort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSort(t *testing.T) {
	t.Log("testing standard merge sort")
	sample := []int{4, 5, 1, 9, 6, 12, 44, 11, 3, 5, 21}
	sampleSortedExpected := []int{1, 3, 4, 5, 5, 6, 9, 11, 12, 21, 44}
	sorted := Sort(sample)
	assert.ElementsMatch(t, sorted, sampleSortedExpected)
}

func TestSortConcurrent(t *testing.T) {
	t.Log("testing concurrent merge sort")
	sample := []int{4, 5, 1, 9, 6, 12, 44, 11, 3, 5, 21}
	sampleSortedExpected := []int{1, 3, 4, 5, 5, 6, 9, 11, 12, 21, 44}
	sorted := SortConcurrent(sample)
	assert.ElementsMatch(t, sorted, sampleSortedExpected)
}
