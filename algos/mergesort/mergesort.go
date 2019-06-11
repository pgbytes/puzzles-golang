package mergesort

// This is the normal merge sort with sequential flow of steps.

func Sort(arr []int) []int {
	// recursion exit
	if len(arr) <= 1 {
		return arr
	}
	// split the arrays into half
	mid := len(arr) / 2
	// recursively split the halves into further halves, till there is only 1 element in the array.
	// putting them on the stack while splitting
	s1 := Sort(arr[:mid])
	s2 := Sort(arr[mid:])
	// recursively merge all the split arrays
	return merge(s1, s2)
}

func merge(arr1 []int, arr2 []int) []int {
	size1 := len(arr1)
	size2 := len(arr2)

	finalArr := make([]int, size1+size2)
	i := 0
	j := 0
	index := 0

	for i < size1 && j < size2 {
		if arr1[i] < arr2[j] {
			update(finalArr, arr1, &index, &i)
		} else {
			update(finalArr, arr2, &index, &j)
		}
	}
	for i < size1 {
		update(finalArr, arr1, &index, &i)
	}
	for j < size2 {
		update(finalArr, arr2, &index, &j)
	}
	return finalArr
}

func update(finalArr []int, arr []int, index *int, incrementIndex *int) {
	finalArr[*index] = arr[*incrementIndex]
	*incrementIndex++
	*index++
}
