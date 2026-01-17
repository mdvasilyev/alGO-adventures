package main

func jump(nums []int) int {
	length := len(nums)
	minJumps := make([]int, length)
	for i := 0; i < length-1; i++ {
		for j := 1; j <= nums[i]; j++ {
			idx := i + j
			if idx < length {
				if minJumps[idx] == 0 {
					minJumps[idx] = minJumps[i] + 1
				} else {
					minJumps[idx] = min(minJumps[idx], minJumps[i]+1)
				}
			}
		}
	}
	return minJumps[length-1]
}
