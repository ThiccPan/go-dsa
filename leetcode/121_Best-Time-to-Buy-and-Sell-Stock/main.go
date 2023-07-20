package main

import "fmt"

func maxProfit(prices []int) int {
	buyAt := 0
	sellAt := 1
	profit := 0

	for i := 0; i < len(prices); i++ {
		fmt.Println(i)
		if sellAt >= len(prices) {
			fmt.Println(sellAt)
			return profit
		}

		if sum := prices[sellAt] - prices[buyAt]; sum > profit {
			profit = sum
			fmt.Println(profit)
		}

		if prices[sellAt] < prices[buyAt] {
			buyAt = sellAt
			fmt.Println("swap", buyAt, sellAt)
		}
		sellAt++
	}

	return profit
}

func main() {
	arr := []int{2, 1, 4}
	fmt.Println(maxProfit(arr))
}
