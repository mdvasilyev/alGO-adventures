package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

const N = 40

type Input struct {
	Values []int64
}

func getInput() Input {
	in := bufio.NewReader(os.Stdin)

	values := make([]int64, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &values[i])
	}

	return Input{Values: values}
}

func subsetSum(values []int64, mask uint64) int64 {
	var sum int64

	for i := 0; i < N; i++ {
		if mask&(uint64(1)<<i) != 0 {
			sum += values[i]
		}
	}

	return sum
}

func buildSet(mask uint64) []int {
	result := make([]int, 0)

	for i := 0; i < N; i++ {
		if mask&(uint64(1)<<i) != 0 {
			result = append(result, i+1)
		}
	}

	return result
}

func findEqualSubsets(values []int64) ([]int, []int) {
	const maskLimit uint64 = (uint64(1) << N) - 1

	seen := make(map[int64]uint64)
	rng := rand.New(rand.NewSource(1))

	for {
		mask := uint64(rng.Int63()) & maskLimit

		if mask == 0 {
			continue
		}

		sum := subsetSum(values, mask)

		if oldMask, ok := seen[sum]; ok && oldMask != mask {
			first := oldMask &^ mask
			second := mask &^ oldMask

			if first != 0 && second != 0 {
				return buildSet(first), buildSet(second)
			}
		}

		seen[sum] = mask
	}
}

func printSet(out *bufio.Writer, set []int) {
	fmt.Fprintln(out, len(set))

	for _, idx := range set {
		fmt.Fprint(out, idx, " ")
	}

	fmt.Fprintln(out)
}

func main() {
	input := getInput()

	first, second := findEqualSubsets(input.Values)

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	printSet(out, first)
	printSet(out, second)
}
