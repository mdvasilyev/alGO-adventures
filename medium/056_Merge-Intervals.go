package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := make([][]int, 0)

	for i := 0; i < len(intervals); i++ {
		interval := intervals[i]
		for i < len(intervals)-1 && interval[1] >= intervals[i+1][0] {
			interval[1] = max(interval[1], intervals[i+1][1])
			i++
		}
		res = append(res, interval)
	}

	return res
}

func main() {
	fmt.Println(merge([][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}))
}
