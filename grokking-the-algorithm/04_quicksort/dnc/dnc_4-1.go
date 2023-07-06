package dnc

// start from the end of array until it reach index 0
func RecSum(arr []int, arrLen int) int {
	sum := 0
	sum += arr[arrLen]
	// recursive case: 
	if arrLen > 0 {
		sum += RecSum(arr[:arrLen], arrLen-1)
	}
	// base case: arrLen is zero
	return sum
}