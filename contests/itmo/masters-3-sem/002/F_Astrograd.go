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

	globalCounter := 0
	ticketsBought := 0
	stack := make([]int, 0, n)
	m := make([]int, 10e5)

	for i := 0; i < n; i++ {
		scanner.Scan()
		s := scanner.Text()
		cmd := strings.Split(s, " ")

		if cmd[0] == "1" {
			id, _ := strconv.Atoi(cmd[1])
			m[id] = globalCounter
			globalCounter++
			stack = append(stack, id)
		} else if cmd[0] == "2" {
			ticketsBought++
			m[stack[0]] = 0
			stack = stack[1:]
			if len(stack) == 0 {
				ticketsBought = 0
				globalCounter = 0
			}
		} else if cmd[0] == "3" {
			m[stack[len(stack)-1]] = 0
			globalCounter--
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				ticketsBought = 0
				globalCounter = 0
			}
		} else if cmd[0] == "4" {
			id, _ := strconv.Atoi(cmd[1])
			fmt.Printf("%d\n", m[id]-ticketsBought)
		} else {
			fmt.Println(stack[0])
		}
	}
}
