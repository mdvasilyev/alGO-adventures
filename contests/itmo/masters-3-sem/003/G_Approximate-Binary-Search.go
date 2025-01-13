package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func closestValue(s []int, val int) int {
	left, right := 0, len(s)-1
	res := s[0]

	for left <= right {
		mid := left + (right-left)/2

		if math.Abs(float64(s[mid]-val)) < math.Abs(float64(res-val)) {
			res = s[mid]
		} else if math.Abs(float64(s[mid]-val)) == math.Abs(float64(res-val)) {
			res = min(res, s[mid])
		}

		if s[mid] == val {
			return s[mid]
		} else if s[mid] < val {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	s := make([]int, 0, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		val, _ := strconv.Atoi(scanner.Text())
		s = append(s, val)
	}

	for i := 0; i < k; i++ {
		scanner.Scan()
		val, _ := strconv.Atoi(scanner.Text())

		fmt.Println(closestValue(s, val))
	}
}
