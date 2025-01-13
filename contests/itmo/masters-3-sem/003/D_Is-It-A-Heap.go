package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isHeap(s []int) bool {
	for i := range s {
		leftChildIdx, rightChildIdx := 2*i+1, 2*i+2
		if leftChildIdx < len(s) {
			if s[i] > s[leftChildIdx] {
				return false
			}
		}
		if rightChildIdx < len(s) {
			if s[i] > s[rightChildIdx] {
				return false
			}
		}
	}
	return true
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

	if isHeap(s) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
