package recursive_algorithms_in_go

func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	first := 0
	last := len(a) - 1

	left := first + 1
	right := last

	pivot := len(a) / 2
	a[first], a[pivot] = a[pivot], a[first]

	for left <= right {
		for ; left <= last && a[left] < a[first]; left++ {}
		for ; first <= right && a[first] < a[right]; right -- {}

		if left <= right {
			a[left], a[right] = a[right], a[left]
			left++
			right--
		}
	}

	a[first], a[right] = a[right], a[first]

	quicksort(a[:right])
	quicksort(a[left:])
	return a
}