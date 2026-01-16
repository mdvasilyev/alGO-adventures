package main

import "fmt"

// Time:	O(n)
// Space:	O(n)
func sortColors(nums []int) {
	counter := map[int]int{
		0: 0,
		1: 0,
		2: 0,
	}
	for _, val := range nums {
		counter[val]++
	}
	for pos, val := 0, 0; val < 3; val++ {
		for i := 0; i < counter[val]; i++ {
			nums[pos] = val
			pos++
		}
	}
}

func main() {
	nums1 := []int{2, 0, 2, 1, 1, 0}
	nums2 := []int{2, 0, 1}

	sortColors(nums1)
	sortColors(nums2)

	fmt.Println(nums1)
	fmt.Println(nums2)
}
