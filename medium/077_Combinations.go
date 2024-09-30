package main

import "fmt"

// Time:	O(2^n)
// Space:	O(C_n^k)
func combine(n int, k int) [][]int {
	if k > n {
		return nil
	}
	var res [][]int
	if k == 1 {
		res = make([][]int, n)
		for i := 1; i <= n; i++ {
			res[i-1] = []int{i}
		}
		return res
	}
	if n == k {
		res = make([][]int, 1)
		valRange := make([]int, n)
		for i := 1; i <= n; i++ {
			valRange[i-1] = i
		}
		res[0] = valRange
		return res
	}
	combinations1 := combine(n-1, k-1)
	res = make([][]int, 0, len(combinations1))
	for _, c := range combinations1 {
		newC := make([]int, len(c), len(c)+1)
		copy(newC, c)
		newC = append(newC, n)
		res = append(res, newC)
	}
	combinations2 := combine(n-1, k)
	for _, c := range combinations2 {
		res = append(res, c)
	}
	return res
}

func main() {
	n1, k1 := 4, 2
	n2, k2 := 1, 1

	fmt.Println(combine(n1, k1))
	fmt.Println(combine(n2, k2))
}
