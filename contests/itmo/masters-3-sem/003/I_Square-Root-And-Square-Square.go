package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func solveEquation(c float64) float64 {
	left, right := float64(0), c

	for i := 0; i < 100; i++ {
		mid := left + (right-left)/2
		val := mid*mid + math.Sqrt(mid)
		if val > c {
			right = mid
		} else {
			left = mid
		}
	}

	return left
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	c, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Println(solveEquation(c))
}
