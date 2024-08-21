package main

import (
	"fmt"
	"strconv"
)

// Time:	O(n)
// Space:	O(n)
func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// Time:	O(n)
// Space:	O(n)
func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	}
	original := x
	reversed := 0
	for x > 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}
	return original == reversed
}

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
}
