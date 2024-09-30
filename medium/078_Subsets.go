package main

import "fmt"

// Time:	O(2^n)
// Space:	O(n)
func subsets(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{{}, {nums[0]}}
	}
	return addValueToSubsets(subsets(nums[:len(nums)-1]), nums[len(nums)-1])
}

func addValueToSubsets(subsets [][]int, value int) [][]int {
	newSubsets := make([][]int, len(subsets), 2*len(subsets))
	copy(newSubsets, subsets)
	for _, sub := range subsets {
		newSub := make([]int, len(sub), len(sub)+1)
		copy(newSub, sub)
		newSubsets = append(newSubsets, append(newSub, value))
	}
	return newSubsets
}

func main() {
	nums1 := []int{1}
	nums2 := []int{1, 2, 3}
	nums3 := []int{1, 2, 3, 4}
	nums4 := []int{1, 2, 3, 4, 5}
	nums5 := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(subsets(nums1))
	fmt.Println(subsets(nums2))
	fmt.Println(subsets(nums3))
	fmt.Println(subsets(nums4))
	fmt.Println(subsets(nums5))
}
