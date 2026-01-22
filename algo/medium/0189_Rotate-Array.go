package main

func rotate(nums []int, k int) {
	length := len(nums)
	k %= length
	rev(nums, 0, length-1)
	rev(nums, 0, k-1)
	rev(nums, k, length-1)
}

func rev(nums []int, left, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
