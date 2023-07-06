package main

import "fmt"

func main() {
	arr := []int{2,3,4,5,1}
	fmt.Println(recSum(arr, len(arr)-1))
}

// start from the end of array until it reach index 0
func recSum(arr []int, arrLen int) int {
	sum := 0
	sum += arr[arrLen]
	// recursive case: 
	if arrLen > 0 {
		sum += recSum(arr[:arrLen], arrLen-1)
	}
	// base case: arrLen is zero
	return sum
}