package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Triple struct {
	val   int
	left  int
	right int
}

func isBST(s []Triple) bool {
	for _, t := range s {
		if t.left >= 0 {
			if s[t.left].val >= t.val {
				return false
			}
		}
		if t.right >= 0 {
			if s[t.right].val <= t.val {
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
	s := make([]Triple, 0, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		value, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		left, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		right, _ := strconv.Atoi(scanner.Text())

		s = append(s, Triple{val: value, left: left - 1, right: right - 1})
	}

	scanner.Scan()

	if isBST(s) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
