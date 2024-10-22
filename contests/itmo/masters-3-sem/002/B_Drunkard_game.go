package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()

	nMax, _ := strconv.Atoi(scanner.Text())

	firstDeck := make([]int, nMax/2, nMax)
	secondDeck := make([]int, nMax/2, nMax)

	for i := 0; i < nMax/2; i++ {
		scanner.Scan()
		firstDeck[i], _ = strconv.Atoi(scanner.Text())
	}
	for i := 0; i < nMax/2; i++ {
		scanner.Scan()
		secondDeck[i], _ = strconv.Atoi(scanner.Text())
	}

	for counter, maxCounter := 0, int(2e5); counter < maxCounter; counter++ {
		if len(firstDeck) == 0 {
			fmt.Printf("second %d\n", counter)
			return
		}
		if len(secondDeck) == 0 {
			fmt.Printf("first %d\n", counter)
			return
		}

		if int(math.Abs(float64(firstDeck[0]-secondDeck[0]))) == nMax-1 {
			if firstDeck[0] < secondDeck[0] {
				firstDeck = append(firstDeck, firstDeck[0], secondDeck[0])
			} else {
				secondDeck = append(secondDeck, firstDeck[0], secondDeck[0])
			}
		} else if firstDeck[0] > secondDeck[0] {
			firstDeck = append(firstDeck, firstDeck[0], secondDeck[0])
		} else {
			secondDeck = append(secondDeck, firstDeck[0], secondDeck[0])
		}
		firstDeck = firstDeck[1:]
		secondDeck = secondDeck[1:]
	}

	fmt.Println("draw")
}
