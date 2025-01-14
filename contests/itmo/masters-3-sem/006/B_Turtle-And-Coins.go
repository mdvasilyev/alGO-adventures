package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findOptimalPath(grid [][]int, n, m int) (int, string) {
	dp := make([][]int, n)
	prev := make([][]rune, n)

	for i := range dp {
		dp[i] = make([]int, m)
		prev[i] = make([]rune, m)
	}

	dp[0][0] = grid[0][0]

	for j := 1; j < m; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
		prev[0][j] = 'R'
	}

	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
		prev[i][0] = 'D'
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			left := dp[i][j-1]
			above := dp[i-1][j]

			if left > above {
				dp[i][j] = left + grid[i][j]
				prev[i][j] = 'R'
			} else {
				dp[i][j] = above + grid[i][j]
				prev[i][j] = 'D'
			}
		}
	}

	path := make([]rune, 0, n+m-2)
	curI, curJ := n-1, m-1

	for curI > 0 || curJ > 0 {
		dir := prev[curI][curJ]
		path = append([]rune{dir}, path...)

		if dir == 'R' {
			curJ--
		} else {
			curI--
		}
	}

	return dp[n-1][m-1], string(path)
}

func splitInts(s string) []int {
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanWords)

	var nums []int
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
	}
	return nums
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	dims := splitInts(scanner.Text())
	n, m := dims[0], dims[1]

	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		row := splitInts(scanner.Text())
		grid[i] = make([]int, m)
		copy(grid[i], row)
	}

	coins, path := findOptimalPath(grid, n, m)

	fmt.Println(coins)
	fmt.Println(path)
}
