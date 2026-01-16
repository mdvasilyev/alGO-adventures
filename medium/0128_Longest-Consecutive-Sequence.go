package main

import "fmt"

// Time:	O(n)
// Space:	O(n)
func longestConsecutive(nums []int) int {
	m := make(map[int]int)
	length := 0
	for _, num := range nums {
		if _, exists := m[num]; !exists {
			prevNum := m[num-1]
			nextNum := m[num+1]
			curLength := prevNum + 1 + nextNum
			m[num] = curLength
			if prevNum > 0 {
				m[num-prevNum] = curLength
			}
			if nextNum > 0 {
				m[num+nextNum] = curLength
			}
			if curLength > length {
				length = curLength
			}
		}
	}
	return length
}

func main() {
	nums1 := []int{100, 4, 200, 1, 3, 2}
	nums2 := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}

	fmt.Println(longestConsecutive(nums1))
	fmt.Println(longestConsecutive(nums2))
}
