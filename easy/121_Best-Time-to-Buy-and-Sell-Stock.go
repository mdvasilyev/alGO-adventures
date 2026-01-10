package main

import (
	"fmt"
)

// Time:	O(n)
// Space:	O(1)
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	initial := prices[0]
	best := 0
	for i := range prices {
		initial = min(initial, prices[i])
		best = max(best, prices[i]-initial)
	}
	return best
}

func main() {
	prices1 := []int{7, 1, 5, 3, 6, 4}
	prices2 := []int{7, 6, 4, 3, 1}

	fmt.Println(maxProfit(prices1))
	fmt.Println(maxProfit(prices2))
}
