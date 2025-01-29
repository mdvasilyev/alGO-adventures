package main

import (
	"bufio"
	"fmt"
	"os"
)

func isStringCorrect(s string) bool {
	rSeenBefore := false
	for i := range s {
		if s[i] == 'R' {
			rSeenBefore = true
		} else if s[i] == 'M' && !rSeenBefore {
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

	var s string
	fmt.Fscan(in, &s)

	if isStringCorrect(s) {
		fmt.Fprintln(out, "YES")
	} else {
		fmt.Fprintln(out, "NO")
	}
}
