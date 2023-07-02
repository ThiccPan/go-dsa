package main

import "fmt"

func main() {
	arr := []int{4, 2, 3, -7}
	selectionSort(arr)
	fmt.Println(arr)
}

func selectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		minIdx := i
		for j := i; j < len(arr); j++ {
			// search for minimum element
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		// swap with first element
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
		// iterate from the next element
	}
}
