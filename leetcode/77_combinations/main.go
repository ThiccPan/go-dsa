package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(combine(4,2))
	// a := []int{1, 2}
	// val := pop(a)
	// fmt.Println(val)
	// fmt.Println(a)
}

func pop(arr []int) (int) {
	val, newArr := arr[len(arr)-1], arr[:len(arr)-1]
	arr = newArr
	arr[0] = 3
	return val
}

func combine(n int, k int) [][]int {
	solution := [][]int{}

	var combining func([]int, int)
	combining = func(comb []int, appended int) {
		if len(comb) == k {
			combCpy := make([]int, len(comb))
			copy(combCpy, comb)
			solution = append(solution, combCpy)
			return
		}

		i := appended

		for ; i <= n; i++ {
			comb = append(comb, i)
			fmt.Println("cur comb: ", comb)
			fmt.Println("n+i:", appended)
			combining(comb, i+1)
			comb = comb[:len(comb)-1]
			
		}
	}
	combining([]int{}, 1)
	return solution
}

func wait1s() {
	time.Sleep(time.Second)
}
