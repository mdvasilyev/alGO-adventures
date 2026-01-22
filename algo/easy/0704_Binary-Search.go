package main

import "fmt"

// Time:	O(log(n))
// Space:	O(n)
func search(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		m := lo + (hi-lo)/2
		if nums[m] < target {
			lo = m + 1
		} else if nums[m] > target {
			hi = m - 1
		} else {
			return m
		}
	}
	return -1
}

func main() {
	nums1, target1 := []int{-1, 0, 3, 5, 9, 12}, 9
	nums2, target2 := []int{-1, 0, 3, 5, 9, 12}, 2

	fmt.Println(search(nums1, target1))
	fmt.Println(search(nums2, target2))
}
