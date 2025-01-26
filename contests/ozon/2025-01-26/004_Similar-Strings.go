package main

import (
	"bufio"
	"fmt"
	"os"
)

func evenCompare(s1, s2 string) bool {
	s1Len, s2Len := len(s1), len(s2)
	minLen, maxLen := min(s1Len, s2Len), max(s1Len, s2Len)
	if maxLen-minLen > 1 || minLen%2 == 0 && maxLen%2 == 1 {
		return false
	}
	for i := 0; i < minLen; i += 2 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func oddCompare(s1, s2 string) bool {
	s1Len, s2Len := len(s1), len(s2)
	minLen, maxLen := min(s1Len, s2Len), max(s1Len, s2Len)
	if s1Len == 1 && s2Len == 1 {
		return false
	}
	if maxLen-minLen > 1 || minLen%2 == 1 && maxLen%2 == 0 {
		return false
	}
	for i := 1; i < minLen; i += 2 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func countSimilarStrings(n int, ss []string) int {
	totalSimilarStrings := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if evenCompare(ss[i], ss[j]) || oddCompare(ss[i], ss[j]) {
				totalSimilarStrings++
			}
		}
	}
	return totalSimilarStrings
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	res := make([]int, t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(in, &n)
		ss := make([]string, n)
		for j := 0; j < n; j++ {
			var s string
			fmt.Fscan(in, &s)
			ss[j] = s
		}
		res[i] = countSimilarStrings(n, ss)
	}
	for i := range res {
		fmt.Fprintln(out, res[i])
	}
}
