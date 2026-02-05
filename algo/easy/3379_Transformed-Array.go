package main

func constructTransformedArray(nums []int) []int {
	length := len(nums)
	res := make([]int, length)
	for i := range nums {
		res[i] = nums[((i+nums[i])%length+length)%length]
	}
	return res
}
