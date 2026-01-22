package main

func removeElement(nums []int, val int) int {
	counter := 0
	length := len(nums)
	for i := range nums {
		if nums[i] == val {
			counter++
		}
	}

	for left, right := 0, length-1; left < length-counter; left++ {
		for nums[right] == val {
			right--
		}
		if nums[left] == val {
			nums[left], nums[right] = nums[right], nums[left]
		}
	}
	return length - counter
}
