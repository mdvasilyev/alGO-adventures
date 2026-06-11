package main

import (
	"bufio"
	"fmt"
	"os"
)

type Input struct {
	H     int
	W     int
	Sand  []int
	Stone []int
}

func getInput() *Input {
	in := bufio.NewReader(os.Stdin)

	var h, w int
	fmt.Fscan(in, &h)
	fmt.Fscan(in, &w)

	sand := make([]int, w)
	stone := make([]int, w)

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			var cell string
			fmt.Fscan(in, &cell)

			if cell == "0" {
				sand[j]++
			} else if cell == "x" {
				stone[j]++
			}
		}
	}

	return &Input{
		H:     h,
		W:     w,
		Sand:  sand,
		Stone: stone,
	}
}

func moveSand(source, target int, sand, total []int) {
	diff := total[source] - total[target]

	if diff <= 1 || sand[source] == 0 {
		return
	}

	cnt := diff / 2

	if cnt > sand[source] {
		cnt = sand[source]
	}

	sand[source] -= cnt
	total[source] -= cnt

	sand[target] += cnt
	total[target] += cnt
}

func getRes(input *Input) []int {
	w := input.W

	sand := input.Sand
	total := make([]int, w)

	for col := 0; col < w; col++ {
		total[col] = input.Sand[col] + input.Stone[col]
	}

	for col := 0; col < w; col++ {
		total[col] = sand[col]

		if col+1 < w {
			moveSand(col+1, col, sand, total)
		}

		if col-1 >= 0 {
			moveSand(col-1, col, sand, total)
		}
	}

	return sand
}

func printRes(h, w int, sand []int) {
	out := bufio.NewWriter(os.Stdout)

	for row := 0; row < h; row++ {
		for col := 0; col < w; col++ {
			if col > 0 {
				fmt.Fprint(out, " ")
			}

			if row >= h-sand[col] {
				fmt.Fprint(out, "0")
			} else {
				fmt.Fprint(out, "-")
			}
		}
		fmt.Fprintln(out)
	}
}

func main() {
	input := getInput()
	res := getRes(input)
	printRes(input.H, input.W, res)
}
