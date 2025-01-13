package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func lowerBound(s []int, val int) int {
	left, right := 0, len(s)

	for left < right {
		mid := left + (right-left)/2
		if s[mid] < val {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

func upperBound(s []int, val int) int {
	left, right := 0, len(s)

	for left < right {
		mid := left + (right-left)/2
		if s[mid] <= val {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()

	n, _ := strconv.Atoi(scanner.Text())
	s := make([]int, 0, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		value, _ := strconv.Atoi(scanner.Text())
		s = append(s, value)
	}

	scanner.Scan()

	k, _ := strconv.Atoi(scanner.Text())

	sort.Ints(s)

	for i := 0; i < k; i++ {
		scanner.Scan()
		fst, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		snd, _ := strconv.Atoi(scanner.Text())

		leftIdx := lowerBound(s, fst)
		rightIdx := upperBound(s, snd)

		fmt.Printf("%v ", rightIdx-leftIdx)
	}
}
