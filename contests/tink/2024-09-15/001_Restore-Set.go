package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var str string

	_, _ = fmt.Scan(&str)

	sepStr := strings.Split(str, ",")

	for _, val := range sepStr {
		if strings.Contains(val, "-") {
			lNr := strings.Split(val, "-")
			l, _ := strconv.Atoi(lNr[0])
			r, _ := strconv.Atoi(lNr[1])
			for i := l; i <= r; i++ {
				fmt.Print(i, " ")
			}
			continue
		}
		fmt.Print(val, " ")
	}
	fmt.Print("\n")
}
