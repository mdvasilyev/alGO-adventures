package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()

	n, _ := strconv.Atoi(scanner.Text())

	var cur, prev int

	scanner.Scan()
	cur, _ = strconv.Atoi(scanner.Text())

	counter := 0

	for i := 0; i < n-1; i++ {
		scanner.Scan()
		prev = cur
		cur, _ = strconv.Atoi(scanner.Text())
		if cur < prev {
			fmt.Println("-1")
			return
		}
		counter += cur - prev
	}

	fmt.Println(counter)
}
