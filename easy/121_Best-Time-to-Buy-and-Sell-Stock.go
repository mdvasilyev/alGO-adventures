package main

import (
	"fmt"
)

// Time:	O(n)
// Space:	O(n)
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	maximum := 0
	initialAmount := prices[0]
	for _, price := range prices {
		if price < initialAmount {
			initialAmount = price
		} else {
			profit := price - initialAmount
			if profit > maximum {
				maximum = profit
			}
		}
	}
	return maximum
}

func main() {
	prices1 := []int{7, 1, 5, 3, 6, 4}
	prices2 := []int{7, 6, 4, 3, 1}

	fmt.Println(maxProfit(prices1))
	fmt.Println(maxProfit(prices2))
}
