package main

import (
	"bufio"
	"fmt"
	"os"
)

func getPositions(n, m int) string {
	if m == 1 {
		return "1\n1 1 D\n"
	}
	if n == 1 {
		return "1\n1 1 R\n"
	}
	if n >= m {
		return fmt.Sprintf("2\n1 1 D\n%v %v U\n", n, m)
	} else {
		return fmt.Sprintf("2\n1 1 R\n%v %v L\n", n, m)
	}
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	var s string
	for i := 0; i < t; i++ {
		var n, m int
		fmt.Fscan(in, &n, &m)
		tmpString := getPositions(n, m)
		s += tmpString
	}
	fmt.Fprintln(out, s)
}
