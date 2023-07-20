package main

import "fmt"

func productExceptSelf(nums []int) []int {
    res := make([]int, len(nums))
	res2 := make([]int, len(nums))
    
    for i := range nums {
        res[i] = 1
		res2[i] = 1
    }
    
    prefix := 1
    for i := 0; i < len(nums); i++ {
        res[i] = prefix
		fmt.Println("pre", prefix)
        prefix = prefix * nums[i]
    }
	fmt.Println(res)
    
    postfix := 1
    for i := len(nums) - 1; i >= 0; i-- {
		res2[i] = res2[i] * postfix
        res[i] = res[i] * postfix
		fmt.Println("post", postfix)
        postfix = postfix * nums[i]
    }

	fmt.Println(res2)
    return res
}

func main() {
	fmt.Println(productExceptSelf([]int{1,2,3,4}))
}