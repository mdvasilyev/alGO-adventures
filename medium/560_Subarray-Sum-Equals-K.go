package main

func subarraySum(nums []int, k int) int {
	res := 0
	m := make(map[int]int)
	m[0] = 1
	cur := 0

	for i := 0; i < len(nums); i++ {
		cur += nums[i]
		if n, ok := m[cur-k]; ok {
			res += n
		}
		m[cur]++
	}

	return res
}
