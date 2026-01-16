package main

func removeDuplicates(nums []int) int {
	unique := 1
	left, right := 0, 1
	for right < len(nums) {
		for right < len(nums) && nums[left] == nums[right] {
			right++
		}
		if right < len(nums) {
			unique++
			left++
			nums[left], nums[right] = nums[right], nums[left]
			right++
		}
	}
	return unique
}
