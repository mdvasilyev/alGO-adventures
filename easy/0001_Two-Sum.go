package main

import (
	"fmt"
	"sort"
)

// Time: 	O(n log n)
// Space:	O(n)
func twoSum(nums []int, target int) []int {
	sorted := make([]int, 0, len(nums))
	for _, value := range nums {
		sorted = append(sorted, value)
	}
	sort.Ints(sorted)

	left := 0
	right := len(sorted) - 1
	current := sorted[left] + sorted[right]
	for current != target {
		if current < target {
			left++
		} else {
			right--
		}
		current = sorted[left] + sorted[right]
	}

	var result []int
	for i := 0; i < len(nums) && len(result) != 2; i++ {
		if nums[i] == sorted[left] || nums[i] == sorted[right] {
			result = append(result, i)
		}
	}

	return result
}

// Time:	O(n)
// Space:	O(n)
func twoSum2(nums []int, target int) []int {
	m := make(map[int]int)
	for i, val := range nums {
		other, ok := m[target-val]
		if ok && i != other {
			return []int{i, other}
		}
		m[val] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
