package main

import (
	"bufio"
	"fmt"
	"os"
)

func isOpeningBracket(r rune) bool {
	return r == '(' || r == '[' || r == '{'
}

func closingBracket(r rune) rune {
	if r == '(' {
		return ')'
	} else if r == '[' {
		return ']'
	} else {
		return '}'
	}
}

func checkBrackets(s string) bool {
	stack := make([]rune, 0)

	for _, val := range s {
		if isOpeningBracket(val) {
			stack = append(stack, val)
		} else {
			if len(stack) == 0 || val != closingBracket(stack[len(stack)-1]) {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	s := scanner.Text()

	if checkBrackets(s) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
