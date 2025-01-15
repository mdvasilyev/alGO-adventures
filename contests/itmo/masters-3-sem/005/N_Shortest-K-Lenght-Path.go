package main

import (
	"fmt"
	"math"
)

type Edge struct {
	from   int
	to     int
	weight int64
}

func findKPaths(n, m, k, start int, edges []Edge) []int64 {
	dp := make([][]int64, k+1)
	for i := range dp {
		dp[i] = make([]int64, n)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt64
		}
	}

	dp[0][start] = 0

	for i := 0; i < k; i++ {
		for _, edge := range edges {
			if dp[i][edge.from] < math.MaxInt64 {
				newDist := dp[i][edge.from] + edge.weight
				if newDist < dp[i+1][edge.to] {
					dp[i+1][edge.to] = newDist
				}
			}
		}
	}

	res := make([]int64, n)
	for i := range res {
		if dp[k][i] == math.MaxInt64 {
			res[i] = -1
		} else {
			res[i] = dp[k][i]
		}
	}

	return res
}

func main() {
	var n, m, k, start int
	fmt.Scan(&n, &m, &k, &start)
	start--

	edges := make([]Edge, m)
	for i := 0; i < m; i++ {
		var from, to int
		var weight int64
		fmt.Scan(&from, &to, &weight)
		edges[i] = Edge{
			from:   from - 1,
			to:     to - 1,
			weight: weight,
		}
	}

	res := findKPaths(n, m, k, start, edges)
	for _, dist := range res {
		fmt.Println(dist)
	}
}
