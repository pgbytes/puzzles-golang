package mergesort

//  This is the modified merge sort using go concurrency
// if A of length 2x is split into X & Y of size x each,
// sorting of these can take place concurrently.

func SortConcurrent(arr []int) []int {
	// make channel to receive sorted values
	ch := make(chan int, len(arr))
	sortConcurrentIntoChannel(arr, ch)
	var sorted []int
	for v := range ch {
		sorted = append(sorted, v)
	}
	return sorted
}

func sortConcurrentIntoChannel(arr []int, ch chan int) {
	// recursion exit
	defer close(ch)
	if len(arr) <= 1 {
		if len(arr) == 1 {
			ch <- arr[0]
		}
		return
	}

	// channels to receive split up arrays
	mid := len(arr) / 2
	s1 := make(chan int, mid)
	s2 := make(chan int, len(arr)-mid)

	// split and sort concurrently
	go sortConcurrentIntoChannel(arr[:mid], s1)
	go sortConcurrentIntoChannel(arr[mid:], s2)
	mergeConcurrent(s1, s2, ch)
}

func mergeConcurrent(s1, s2, ch chan int) {
	// v, ok = <- s; receiving value from channel, ok returns false when there is no more elements in channel s
	v1, ok1 := <-s1
	v2, ok2 := <-s2

	for ok1 && ok2 {
		if v1 < v2 {
			updateConcurrent(s1, ch, &v1, &ok1)
		} else {
			updateConcurrent(s2, ch, &v2, &ok2)
		}
	}

	for ok1 {
		updateConcurrent(s1, ch, &v1, &ok1)
	}

	for ok2 {
		updateConcurrent(s2, ch, &v2, &ok2)
	}
}

// puts the value v into the final receiving channel
// extracts the next element from array channel s
func updateConcurrent(s chan int, ch chan int, v *int, ok *bool) {
	ch <- *v
	*v, *ok = <-s
}
