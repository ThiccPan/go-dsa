package main

import (
	"fmt"
	// "strings"
)

func main() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	combinations := []string{}
	res := ""
	genCombo(&combinations, 0, 0, res, n)
	fmt.Println(combinations)
	return combinations
}

func genCombo(res *[]string, left int, right int, s string, n int) {
	// base case
	fmt.Println(s)

	if len(s) >= n*2 {
		fmt.Println("return s:", s)
		*res = append(*res, s)
		return
	}

	if left < n {
		genCombo(res, left+1, right, s+"(", n)
	}

	if right < left {
		genCombo(res, left, right+1, s+")", n)
	}
}
