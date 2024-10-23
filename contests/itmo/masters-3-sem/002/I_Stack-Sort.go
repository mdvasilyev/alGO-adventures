package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Command string

func containsValue(values []int, value int) (int, bool) {
	for i, val := range values {
		if val == value {
			return i, true
		}
	}
	return -1, false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()

	n, _ := strconv.Atoi(scanner.Text())

	firstStack := make([]int, n)
	secondStack := make([]int, 0, n)
	resultingStack := make([]int, 0, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		firstStack[i], _ = strconv.Atoi(scanner.Text())
	}

	insertValue := 1
	commands := make([]Command, 0, 2*n)

	for len(resultingStack) != n {
		if idx, ok := containsValue(secondStack, insertValue); ok && idx != len(secondStack)-1 {
			fmt.Println("impossible")
			return
		} else if ok && idx == len(secondStack)-1 {
			resultingStack = append(resultingStack, secondStack[len(secondStack)-1])
			insertValue++
			secondStack = secondStack[:len(secondStack)-1]
			commands = append(commands, "pop")
		} else {
			secondStack = append(secondStack, firstStack[0])
			firstStack = firstStack[1:]
			commands = append(commands, "push")
		}
	}

	for _, cmd := range commands {
		fmt.Println(cmd)
	}
}
