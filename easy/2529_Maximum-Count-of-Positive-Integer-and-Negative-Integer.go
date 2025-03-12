package main

import "sort"

func maximumCount(nums []int) int {
	posStart := sort.Search(len(nums), func(i int) bool { return nums[i] > 0 })

	negEnd := sort.Search(len(nums), func(i int) bool { return nums[i] >= 0 })

	negCount := negEnd
	posCount := len(nums) - posStart

	return max(negCount, posCount)
}
