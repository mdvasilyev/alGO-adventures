package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func collectCoins(n, k int, coins []int) (int, int, []int) {
	initialArray := make([]int, n)
	for i := 1; i < n-1; i++ {
		initialArray[i] = coins[i-1]
	}

	d := make([]int, n)
	for i := range d {
		d[i] = math.MinInt32
	}
	d[0] = 0

	p := make([]int, n)

	for i := 1; i < n; i++ {
		bestPrevPos := i - 1
		for j := max(0, i-k); j < i; j++ {
			if d[j] > d[bestPrevPos] {
				bestPrevPos = j
			}
		}
		d[i] = d[bestPrevPos] + initialArray[i]
		p[i] = bestPrevPos
	}

	path := make([]int, 0)
	currentPos := n - 1
	path = append(path, n) // Add final position

	for currentPos != 0 {
		currentPos = p[currentPos]
		path = append(path, currentPos+1)
	}

	for i := 0; i < len(path)/2; i++ {
		path[i], path[len(path)-1-i] = path[len(path)-1-i], path[i]
	}

	return d[n-1], len(path) - 1, path
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	s := make([]int, n-2)
	for i := 0; i < n-2; i++ {
		scanner.Scan()
		val, _ := strconv.Atoi(scanner.Text())
		s[i] = val
	}

	maxCoins, jumps, positions := collectCoins(n, k, s)

	fmt.Println(maxCoins)
	fmt.Println(jumps)
	for _, pos := range positions {
		fmt.Printf("%d ", pos)
	}
}
