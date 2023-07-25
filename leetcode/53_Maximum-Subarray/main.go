package main

import "fmt"

func main() {
	fmt.Println(maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}))
}

func maxSubArray(nums []int) int {
	max := nums[0]
	sum := 0

	for start := 0; start < len(nums); start++ {
		// check if sum with next index value is greater than switching starting subarr index
		if sum + nums[start] < nums[start] {
			sum = nums[start]
			fmt.Println("masuk kondisi, sum =", sum)
			if sum > max {
				max = sum
			}
			continue
		}
		// traverse to next index
		sum += nums[start]
		if sum > max {
			max = sum
		}
		fmt.Println("sum =", sum)
	}

	return max
}
