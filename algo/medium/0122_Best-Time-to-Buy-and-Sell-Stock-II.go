package main

func maxProfit1(prices []int) int {
	noShare := 0
	withShare := -prices[0]
	for i := 1; i < len(prices); i++ {
		noShare = max(noShare, withShare+prices[i])
		withShare = max(withShare, noShare-prices[i])
	}
	return noShare
}
