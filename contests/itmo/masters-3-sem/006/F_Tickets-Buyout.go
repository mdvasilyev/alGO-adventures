package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	A := make([]int, n+1)
	B := make([]int, n+1)
	C := make([]int, n+1)

	for i := 1; i <= n; i++ {
		fmt.Scan(&A[i], &B[i], &C[i])
	}

	dp := make([]int, n+1)
	dp[0] = 0

	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + A[i]
		if i > 1 {
			dp[i] = min(dp[i], dp[i-2]+B[i-1])
		}
		if i > 2 {
			dp[i] = min(dp[i], dp[i-3]+C[i-2])
		}
	}

	fmt.Println(dp[n])
}
