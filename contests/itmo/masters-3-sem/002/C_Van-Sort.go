package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Pair struct {
	Path   int
	Number int
}

func contains(values []int, value int) (int, bool) {
	for i, val := range values {
		if val == value {
			return i, true
		}
	}
	return -1, false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()

	n, _ := strconv.Atoi(scanner.Text())

	deadEnd := make([]int, 0, n)
	firstPath := make([]int, n)
	secondPath := make([]int, 0, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		firstPath[i], _ = strconv.Atoi(scanner.Text())
	}

	insertValue := 1
	pairs := make([]Pair, 0, 2*n)

	for len(secondPath) != n {
		if idx, ok := contains(deadEnd, insertValue); ok && idx != len(deadEnd)-1 {
			fmt.Println("0")
			return
		} else if ok && idx == len(deadEnd)-1 {
			secondPath = append(secondPath, deadEnd[len(deadEnd)-1])
			insertValue++
			deadEnd = deadEnd[:len(deadEnd)-1]
			pairs = append(pairs, Pair{Path: 2, Number: 1})
		} else {
			deadEnd = append(deadEnd, firstPath[0])
			firstPath = firstPath[1:]
			pairs = append(pairs, Pair{Path: 1, Number: 1})
		}
	}

	for _, p := range pairs {
		fmt.Println(p.Path, p.Number)
	}
}
