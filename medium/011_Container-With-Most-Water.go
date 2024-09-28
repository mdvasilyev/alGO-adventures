package main

import (
	"fmt"
)

// Time:	O(n)
// Space:	O(n)
func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	left := 0
	right := len(height) - 1
	maximum := 0
	for left < right {
		cur := (right - left) * min(height[left], height[right])
		if cur > maximum {
			maximum = cur
		}
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}
	return maximum
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	res := maxArea(height)
	fmt.Println(res)
}
