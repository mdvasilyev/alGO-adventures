package main

import "fmt"

func maxProfit(prices []int, fee int) int {
	length := len(prices)
	if length < 2 {
		return 0
	}

	withShare := -prices[0]
	noShare := 0
	for i := 1; i < length; i++ {
		noShare = max(noShare, withShare+prices[i]-fee)
		withShare = max(withShare, noShare-prices[i])
	}

	return noShare
}

func main() {
	prices := []int{1, 3, 7, 5, 10, 3}
	fee := 3
	fmt.Println(maxProfit(prices, fee))
}
