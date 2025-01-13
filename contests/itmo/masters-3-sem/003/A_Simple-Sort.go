package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func merge(left, right []int) []int {
	leftSize, rightSize := len(left), len(right)
	res := make([]int, 0, leftSize+rightSize)
	leftIdx, rightIdx := 0, 0
	for leftIdx < leftSize && rightIdx < rightSize {
		if left[leftIdx] <= right[rightIdx] {
			res = append(res, left[leftIdx])
			leftIdx++
		} else {
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
	return res
}

func mergeSort(s []int) []int {
	size := len(s)
	if size < 2 {
		return s
	}
	left := mergeSort(s[:size/2])
	right := mergeSort(s[size/2:])
	return merge(left, right)
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

	sorted := mergeSort(s)

	for _, val := range sorted {
		fmt.Printf("%v ", val)
	}
}
