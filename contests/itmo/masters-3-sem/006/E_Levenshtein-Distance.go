package main

import (
	"bufio"
	"fmt"
	"os"
)

func levenshteinDistance(s1, s2 string) int {
	s1Length := len(s1)
	s2Length := len(s2)
	dp := make([][]int, s1Length+1)

	for i := 0; i < s1Length+1; i++ {
		dp[i] = make([]int, s2Length+1)
		dp[i][0] = i
	}

	for i := 0; i < s2Length+1; i++ {
		dp[0][i] = i
	}

	for i := 1; i < s1Length+1; i++ {
		for j := 1; j < s2Length+1; j++ {
			diff := 0
			if s1[i-1] != s2[j-1] {
				diff = 1
			}
			minimum := min(dp[i-1][j]+1, dp[i][j-1]+1)
			dp[i][j] = min(minimum, dp[i-1][j-1]+diff)
		}
	}
	return dp[s1Length][s2Length]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	s1 := scanner.Text()

	scanner.Scan()
	s2 := scanner.Text()

	fmt.Println(levenshteinDistance(s1, s2))
}
