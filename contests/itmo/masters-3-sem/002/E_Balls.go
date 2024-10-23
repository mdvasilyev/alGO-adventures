package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type BallCounter struct {
	Color   int
	Counter int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	explodedBalls := 0
	stack := make([]BallCounter, 0)

	for scanner.Scan() {
		val, _ := strconv.Atoi(scanner.Text())

		if len(stack) == 0 {
			stack = append(stack, BallCounter{Color: val, Counter: 1})
		} else if stack[len(stack)-1].Color == val {
			stack = append(stack, BallCounter{Color: val, Counter: stack[len(stack)-1].Counter + 1})
		} else {
			if stack[len(stack)-1].Counter >= 3 {
				explodedBalls += stack[len(stack)-1].Counter
				stack = stack[:len(stack)-stack[len(stack)-1].Counter]
			}
			if len(stack) != 0 && stack[len(stack)-1].Color == val {
				stack = append(stack, BallCounter{Color: val, Counter: stack[len(stack)-1].Counter + 1})
			} else {
				stack = append(stack, BallCounter{Color: val, Counter: 1})
			}
		}
	}

	if len(stack) != 0 {
		if stack[len(stack)-1].Counter >= 3 {
			explodedBalls += stack[len(stack)-1].Counter
		}
	}

	fmt.Println(explodedBalls)
}
