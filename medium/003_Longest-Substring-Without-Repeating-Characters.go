package main

import "fmt"

// Time:	O(n)
// Space:	O(n)
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	m := make(map[rune]bool)
	arr := []rune(s)
	left := 0
	right := 1
	maximum := 1
	current := 1
	m[arr[left]] = true
	for left < len(s) && right < len(s) {
		in, ok := m[arr[right]]
		if in && ok {
			m[arr[left]] = false
			current--
			left++
		} else {
			m[arr[right]] = true
			current++
			right++
		}
		if current > maximum {
			maximum = current
		}
	}
	return maximum
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
