package main

func partitionLabels(s string) []int {
	res := make([]int, 0)
	m := make(map[rune]int)

	for i, b := range s {
		m[b] = i
	}

	left, right := 0, 0
	for i, b := range s {
		right = max(right, m[b])
		if i == right {
			res = append(res, right-left+1)
			left = right + 1
		}
	}

	return res
}
