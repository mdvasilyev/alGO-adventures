package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func mergeAndCount(left, right []int) ([]int, int) {
	inversions := 0
	leftSize, rightSize := len(left), len(right)
	res := make([]int, 0, leftSize+rightSize)
	leftIdx, rightIdx := 0, 0
	for leftIdx < leftSize && rightIdx < rightSize {
		if left[leftIdx] <= right[rightIdx] {
			res = append(res, left[leftIdx])
			leftIdx++
		} else {
			inversions += leftSize - leftIdx
			res = append(res, right[rightIdx])
			rightIdx++
		}
	}
	for ; leftIdx < leftSize; leftIdx++ {
		res = append(res, left[leftIdx])
	}
	for ; rightIdx < rightSize; rightIdx++ {
		res = append(res, right[rightIdx])
	}
	return res, inversions
}

func countInversion(s []int) ([]int, int) {
	size := len(s)
	if size < 2 {
		return s, 0
	}
	left, leftInversions := countInversion(s[:size/2])
	right, rightInversions := countInversion(s[size/2:])
	res, resInversions := mergeAndCount(left, right)
	return res, leftInversions + resInversions + rightInversions
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()

	n, _ := strconv.Atoi(scanner.Text())
	s := make([]int, 0, n)

	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		s = append(s, value)
	}

	_, inversions := countInversion(s)

	fmt.Println(inversions)
}
