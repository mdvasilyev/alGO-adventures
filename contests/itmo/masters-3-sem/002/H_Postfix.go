package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isOperator(s string) bool {
	return s == "+" || s == "-" || s == "*"
}

func operatorToFunc(s string) func(int, int) int {
	if s == "+" {
		return func(a, b int) int {
			return a + b
		}
	} else if s == "-" {
		return func(a, b int) int {
			return a - b
		}
	} else {
		return func(a, b int) int {
			return a * b
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	stack := make([]int, 0)

	for scanner.Scan() {
		var val int
		token := scanner.Text()
		if isOperator(token) {
			val = operatorToFunc(token)(stack[len(stack)-2], stack[len(stack)-1])
			stack = stack[:len(stack)-2]
		} else {
			val, _ = strconv.Atoi(token)
		}
		stack = append(stack, val)
	}

	fmt.Println(stack[0])
}
