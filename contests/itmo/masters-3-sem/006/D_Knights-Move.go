package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const mod = 1000000000

func numbersNumber(n int) int {
	s := make([][10]int, n)

	for i := 0; i < 10; i++ {
		if i != 0 && i != 8 {
			s[0][i] = 1
		}
	}

	for i := 1; i < n; i++ {
		for j := 0; j < 10; j++ {
			switch j {
			case 0:
				s[i][j] = s[i-1][4] + s[i-1][6]
			case 1:
				s[i][j] = s[i-1][8] + s[i-1][6]
			case 2:
				s[i][j] = s[i-1][7] + s[i-1][9]
			case 3:
				s[i][j] = s[i-1][4] + s[i-1][8]
			case 4:
				s[i][j] = s[i-1][3] + s[i-1][9]
				s[i][j] += s[i-1][0]
			case 5:
				s[i][j] = 0
			case 6:
				s[i][j] = s[i-1][1] + s[i-1][7]
				s[i][j] += s[i-1][0]
			case 7:
				s[i][j] = s[i-1][2] + s[i-1][6]
			case 8:
				s[i][j] = s[i-1][1] + s[i-1][3]
			case 9:
				s[i][j] = s[i-1][2] + s[i-1][4]
			}
			if s[i][j] >= mod {
				s[i][j] %= mod
			}
		}

	}

	sum := 0

	for i := 0; i < 10; i++ {
		sum += s[n-1][i]
		if sum >= mod {
			sum %= mod
		}
	}

	return sum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	fmt.Println(numbersNumber(n))
}
