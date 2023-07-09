package quicksort

func Quicksort(arr []int) []int {
	// base case 1: partition contains 0/1 element
	if len(arr) <= 1 {
		return arr
	}
	// base case 2: partition contais 2 element
	if len(arr) == 2 {
		if arr[0] > arr[1] {
			arr[0], arr[1] = arr[1], arr[0]
		}
		return arr
	}

	// recursive case
	if len(arr) > 2 {
		// pick pivot
		pivot := arr[0]

		less := []int{}
		greater := []int{}

		for _, v := range arr[1:] {
			if v > pivot {
				greater = append(greater, v)
			}
			if v <= pivot {
				less = append(less, v)
			}
		}
		less = append(Quicksort(less), pivot)
		less = append(less, Quicksort(greater)...)
		return less
	}

	return arr
}
