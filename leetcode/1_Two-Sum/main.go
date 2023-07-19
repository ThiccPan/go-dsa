package main

import "fmt"

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for firstIdx, num := range nums {
		secondIdx, ok := numMap[target-num]
		if !ok {
			numMap[num] = firstIdx
			continue
		}
		return []int{firstIdx, secondIdx}
	}
	
	return []int{}
}

func main() {
	a := twoSum([]int{3,2,4}, 6)
	fmt.Println(a)
}
