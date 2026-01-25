package main

import "math"

func minSubArrayLen(target int, nums []int) int {
	res := math.MaxInt
	curSum := 0
	for left, right := 0, 0; right < len(nums); right++ {
		curSum += nums[right]
		for curSum >= target {
			res = min(res, right-left+1)
			curSum -= nums[left]
			left++
		}
	}
	if res == math.MaxInt {
		return 0
	}
	return res
}

func main() {
	target := 7
	nums := []int{2, 3, 1, 2, 4, 3}
	println(minSubArrayLen(target, nums))
}
