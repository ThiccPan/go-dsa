package dnc

import "log"

// param and return refer to the index
func RecursiveBSearch(arr []int, search, low, high, point int) int {
	// recursive case: arr[point] lower than search
	if arr[point] < search {
		log.Println(arr[point+1:])
		return RecursiveBSearch(arr, search, point+1, high, ((point+1) + high)/2)
	}
	// recursive case: arr[point] greater than search
	if arr[point] > search {
		log.Println(arr[:point-1])
		return RecursiveBSearch(arr, search, low, point-1, (low + (point-1))/2)
	}
	// base case: high equals to low
	if low >= high {
		return point
	}
	// base case: point equals to search
	return point
}