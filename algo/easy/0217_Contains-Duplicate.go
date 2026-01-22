package main

import "fmt"

// Time:	O(n)
// Space:	O(n)
func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for _, val := range nums {
		if _, exists := m[val]; exists {
			return true
		} else {
			m[val]++
		}
	}
	return false
}

func main() {
	nums1 := []int{1, 2, 3, 1}
	nums2 := []int{1, 2, 3, 4}
	nums3 := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}

	fmt.Println(containsDuplicate((nums1)))
	fmt.Println(containsDuplicate((nums2)))
	fmt.Println(containsDuplicate((nums3)))
}
