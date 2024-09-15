package main

import (
	"fmt"
	"math/big"
)

func divisorNumber(n int64) int64 {
	var number int64
	for i := int64(1); i <= n; i++ {
		if n%i == 0 {
			number++
		}
	}
	return number
}

func main() {
	var l, r int64
	var res int

	_, _ = fmt.Scan(&l, &r)
	big.NewInt(l).ProbablyPrime(0)

	for l <= r {
		if !big.NewInt(l).ProbablyPrime(0) && big.NewInt(divisorNumber(l)).ProbablyPrime(0) {
			res++
		}
		l++
	}
	fmt.Println(res)
}
