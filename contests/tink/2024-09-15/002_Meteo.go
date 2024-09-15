package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int
	strs := make([]string, 0, n)
	res := make([]int, 0, n)

	_, _ = fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		var s string
		_, _ = fmt.Scan(&s)
		strs = append(strs, s)
	}

	prev := 0
	for _, val := range strs {
		if strings.Contains(val, "-") {
			res = append(res, 1)
			prev++
		} else if cur, _ := strconv.Atoi(val); cur > prev {
			res = append(res, cur-prev)
			prev = cur
		} else {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
	for _, v := range res {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
