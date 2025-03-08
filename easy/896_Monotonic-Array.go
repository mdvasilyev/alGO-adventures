package main

func isMonotonic(nums []int) bool {
	if len(nums) < 2 {
		return true
	}

	isUp, isDown := false, false

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			isDown = true
		} else if nums[i] < nums[i+1] {
			isUp = true
		}

		if isUp && isDown {
			return false
		}
	}

	return true
}
