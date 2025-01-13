package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func minTime(n, x, y int) int {
	n--
	minimal, maximal := 0, n*max(x, y)

	for minimal < maximal {
		mid := minimal + (maximal-minimal)/2
		samples := mid/x + mid/y
		if samples >= n {
			maximal = mid
		} else {
			minimal = mid + 1
		}
	}

	return maximal + min(x, y)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	x, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	y, _ := strconv.Atoi(scanner.Text())

	fmt.Printf("%v.6f\n", minTime(n, x, y))
}
