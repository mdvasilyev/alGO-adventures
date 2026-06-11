package main

import (
	"bufio"
	"fmt"
	"os"
)

const INF int64 = 1 << 60

type Input struct {
	N int
	W []int64
}

func getInput() *Input {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	w := make([]int64, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &w[i])
	}

	return &Input{
		N: n,
		W: w,
	}
}

func rangeSum(prefix []int64, l, r int) int64 {
	if l > r {
		return 0
	}

	return prefix[r] - prefix[l-1]
}

func solve(input *Input) (int, [][]int) {
	n := input.N
	w := input.W

	prefix := make([]int64, n+1)
	for i := 1; i <= n; i++ {
		prefix[i] = prefix[i-1] + w[i]
	}

	dp := make([][]int64, n+2)
	cnt := make([][]int, n+2)
	optMin := make([][]int, n+2)
	optMax := make([][]int, n+2)
	roots := make([][][]int, n+2)

	for i := 0; i <= n+1; i++ {
		dp[i] = make([]int64, n+2)
		cnt[i] = make([]int, n+2)
		optMin[i] = make([]int, n+2)
		optMax[i] = make([]int, n+2)
		roots[i] = make([][]int, n+2)
	}

	for i := 1; i <= n; i++ {
		dp[i][i] = w[i]
		cnt[i][i] = 1
		optMin[i][i] = i
		optMax[i][i] = i
		roots[i][i] = []int{i}
	}

	for length := 2; length <= n; length++ {
		for l := 1; l+length-1 <= n; l++ {
			r := l + length - 1

			leftBound := optMin[l][r-1]
			rightBound := optMax[l+1][r]

			best := INF
			sum := rangeSum(prefix, l, r)

			for root := leftBound; root <= rightBound; root++ {
				cur := dp[l][root-1] + dp[root+1][r] + sum

				if cur < best {
					best = cur
					roots[l][r] = roots[l][r][:0]
					roots[l][r] = append(roots[l][r], root)
				} else if cur == best {
					roots[l][r] = append(roots[l][r], root)
				}
			}

			dp[l][r] = best
			optMin[l][r] = roots[l][r][0]
			optMax[l][r] = roots[l][r][len(roots[l][r])-1]

			totalCount := 0

			for _, root := range roots[l][r] {
				leftCount := 1
				rightCount := 1

				if l <= root-1 {
					leftCount = cnt[l][root-1]
				}
				if root+1 <= r {
					rightCount = cnt[root+1][r]
				}

				totalCount += leftCount * rightCount
			}

			cnt[l][r] = totalCount
		}
	}

	var build func(l, r, limit int) [][]int

	build = func(l, r, limit int) [][]int {
		if l > r {
			return [][]int{{}}
		}

		res := make([][]int, 0)

		for _, root := range roots[l][r] {
			if len(res) >= limit {
				break
			}

			leftLimit := limit - len(res)
			leftVariants := build(l, root-1, leftLimit)

			for _, left := range leftVariants {
				if len(res) >= limit {
					break
				}

				rightLimit := limit - len(res)
				rightVariants := build(root+1, r, rightLimit)

				for _, right := range rightVariants {
					if len(res) >= limit {
						break
					}

					seq := make([]int, 0, r-l+1)
					seq = append(seq, root)
					seq = append(seq, left...)
					seq = append(seq, right...)

					res = append(res, seq)
				}
			}
		}

		return res
	}

	ans := build(1, n, cnt[1][n])

	return cnt[1][n], ans
}

func main() {
	input := getInput()

	k, algos := solve(input)

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fprintln(out, k)

	for _, algo := range algos {
		for i, x := range algo {
			if i > 0 {
				fmt.Fprint(out, " ")
			}
			fmt.Fprint(out, x)
		}
		fmt.Fprintln(out)
	}
}
