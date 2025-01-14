package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func findMaxIndex(s []int) int {
	maxIndex := 0
	for i := 1; i < len(s); i++ {
		if s[i] > s[maxIndex] {
			maxIndex = i
		}
	}
	return maxIndex
}

func findLIS(s []int, n int) (int, []int) {
	dp := make([]int, n)
	prev := make([]int, n)

	for i := 0; i < n; i++ {
		dp[i] = 0
		prev[i] = i
	}

	for i := 0; i < n; i++ {
		maxLen := 0
		lastIndex := i

		for j := 0; j < n; j++ {
			if s[i] > s[j] && dp[j] > maxLen {
				maxLen = dp[j]
				lastIndex = j
			}
		}

		dp[i] = maxLen + 1
		prev[i] = lastIndex
	}

	maxIndex := findMaxIndex(dp)
	result := make([]int, 0)
	result = append(result, s[maxIndex])

	for maxIndex != prev[maxIndex] {
		maxIndex = prev[maxIndex]
		result = append(result, s[maxIndex])
	}

	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-1-i] = result[len(result)-1-i], result[i]
	}

	return dp[findMaxIndex(dp)], result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		arr[i], _ = strconv.Atoi(scanner.Text())
	}

	length, subsequence := findLIS(arr, n)

	fmt.Println(length)
	for _, val := range subsequence {
		fmt.Printf("%d ", val)
	}
}
