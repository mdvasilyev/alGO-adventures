package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func countingSort(s []int, maxValue int) []int {
	counts := make([]int, maxValue+1)
	res := make([]int, 0, len(s))

	for _, val := range s {
		counts[val]++
	}

	for pos, num := range counts {
		for i := 0; i < num; i++ {
			res = append(res, pos)
		}
	}

	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()

	n, _ := strconv.Atoi(scanner.Text())
	s := make([]int, 0, n)
	maxValue := 0

	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		maxValue = max(maxValue, value)
		s = append(s, value)
	}

	sorted := countingSort(s, maxValue)

	for _, val := range sorted {
		fmt.Printf("%v ", val)
	}
}
