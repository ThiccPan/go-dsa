package main

import "fmt"

func binarySearch(target int, nums []int) (idx int) {
	min := 0
	max := len(nums) - 1
	
	for min < max {
		pointing := (min + max)/2
		if nums[pointing] == target {
			return pointing
		}
		if nums[pointing] > target {
			max = pointing - 1
			fmt.Println("max: ", max)
			fmt.Println("pointing: ", pointing)
			continue
		}

		if nums[pointing] < target{
			min = pointing + 1
			fmt.Println("min: ", min)
			fmt.Println("pointing: ", pointing)
			continue
		}
	}
	return max
}

func main() {
	fmt.Println(binarySearch(10, []int{1, 2, 4, 5, 6, 8, 10}))
}
