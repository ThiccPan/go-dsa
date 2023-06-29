package main

import "fmt"

func TopKFrequent(nums []int, k int) []int {
	angkaMap := map[int]int{}
	angkaMuncul := make([][]int, len(nums) + 1)
	for _, v := range nums {
		angkaMap[v]++
	}

	// fmt.Println(angkaMap)

	for angka, muncul := range angkaMap {
		angkaMuncul[muncul] = append(angkaMuncul[muncul], angka)
	}

	// fmt.Println(angkaMuncul)

	answer := new([]int)
	for i := len(angkaMuncul) -1; i >= 0; i-- {
		fmt.Println(angkaMuncul[i])
		*answer = append(*answer, angkaMuncul[i]...)
		if len(*answer) == k {
			break
		}
	}
	return *answer
}
