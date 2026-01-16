package main

import (
	"fmt"
	"math"
)

// Time:	O(log n)
// Space:	O(1)
func reverse(x int) int {
	positive := x > 0
	if !positive {
		x *= -1
	}
	reversed := 0
	for x != 0 {
		if reversed > math.MaxInt32/10 {
			return 0
		}
		reversed = reversed*10 + x%10
		x /= 10
	}
	if !positive {
		return -reversed
	}
	return reversed
}

func main() {
	fmt.Println(reverse(1534236469))
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
}
