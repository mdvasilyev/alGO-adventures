package main

import "fmt"

// Time:	O(n)
// Space:	O(n)
func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	res := make([]int, n)
	res[0] = 1
	res[1] = 2
	for i := 2; i < n; i++ {
		res[i] = res[i-1] + res[i-2]
	}
	return res[len(res)-1]
}

func main() {
	n1 := 1
	n2 := 2
	n3 := 3

	fmt.Println(climbStairs(n1))
	fmt.Println(climbStairs(n2))
	fmt.Println(climbStairs(n3))
}
