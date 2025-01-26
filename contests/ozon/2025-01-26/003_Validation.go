package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func validate(n int, pairs map[string]string, s string) bool {
	ss := strings.Split(s, ",")
	if len(ss) > len(pairs) {
		return false
	}
	usedPrices := make(map[string]bool)
	for k := range pairs {
		usedPrices[pairs[k]] = false
	}
	for i := range ss {
		if strings.ContainsAny(ss[i], " ,.;`/~!@#$%^&*()_+?><") || !strings.Contains(ss[i], ":") {
			return false
		}
		parts := strings.Split(ss[i], ":")
		if len(parts) != 2 {
			return false
		}
		if len(parts[0]) == 0 || len(parts[1]) == 0 {
			return false
		}
		if parts[1][0] == '0' || usedPrices[parts[1]] {
			return false
		}
		usedPrices[parts[1]] = true
		if price, ok := pairs[parts[0]]; !ok || price != parts[1] {
			return false
		}
	}
	for k := range usedPrices {
		if !usedPrices[k] {
			return false
		}
	}
	return true
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	res := make([]bool, 0)
	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(in, &n)
		pairs := make(map[string]string, n)
		for j := 0; j < n; j++ {
			var name, price string
			fmt.Fscan(in, &name, &price)
			pairs[name] = price
		}
		var s string
		fmt.Fscan(in, &s)
		res = append(res, validate(n, pairs, s))
	}
	for i := range res {
		if res[i] {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
