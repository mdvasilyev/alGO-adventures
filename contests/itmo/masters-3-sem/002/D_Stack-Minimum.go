package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	stack := make([]int, 0, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		s := scanner.Text()

		cmd := strings.Split(s, " ")

		if cmd[0] == "1" {
			val, _ := strconv.Atoi(cmd[1])
			if len(stack) == 0 {
				stack = append(stack, val)
			} else {
				stack = append(stack, min(val, stack[len(stack)-1]))
			}
		} else if cmd[0] == "2" {
			stack = stack[:len(stack)-1]
		} else {
			fmt.Println(stack[len(stack)-1])
		}
	}
}
