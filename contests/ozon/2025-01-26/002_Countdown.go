package main

import (
	"bufio"
	"fmt"
	"os"
)

var answers map[int]int

func allDigitsSame(n int) bool {
	digit := n % 10
	for n != 0 {
		if digit != n%10 {
			return false
		}
		n /= 10
	}
	return true
}

func checkDecades(n int) bool {
	return n == 100 || n == 1000 || n == 10000 || n == 100000 || n == 1000000 || n == 10000000 || n == 100000000 || n == 1000000000
}

func getPlatesNumber(n int) int {
	if val, ok := answers[n]; ok {
		return val
	} else if allDigitsSame(n) || checkDecades(n) {
		answers[n] = getPlatesNumber(n-1) + 1
	} else {
		answers[n] = getPlatesNumber(n - 1)
	}
	return answers[n]
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	res := make([]int, 0)
	answers = make(map[int]int)
	answers[0] = 1
	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(in, &n)
		res = append(res, getPlatesNumber(n))
	}
	for i := range res {
		fmt.Fprintln(out, res[i])
	}
}
