package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func maxPower(s string) int {
	terms := strings.Split(s, "+")
	pow := 0
	for _, term := range terms {
		if strings.Contains(term, "^") {
			curPow, _ := strconv.Atoi(strings.Trim(strings.SplitAfterN(term, "^", 2)[1], " d"))
			pow = max(curPow, pow)
		} else if strings.Contains(s, "x") {
			pow = max(1, pow)
		}
	}
	return pow
}

func main() {
	functions := make([]string, 0, 2)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		functions = append(functions, scanner.Text())
	}

	n, m := maxPower(functions[0]), maxPower(functions[1])
	if n < m {
		fmt.Println("YES")
		fmt.Println("NO")
		fmt.Println("NO")
	} else if n == m {
		fmt.Println("YES")
		fmt.Println("YES")
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
		fmt.Println("YES")
		fmt.Println("NO")
	}
}
