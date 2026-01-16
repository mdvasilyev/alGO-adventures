package main

import (
	"fmt"
	"strings"
)

// Time:	O(len(s))
// Space:	O(len(s))
func isPalindrome3(s string) bool {
	s = strip(s)
	for lo, hi := 0, len(s)-1; lo < hi; lo, hi = lo+1, hi-1 {
		if s[lo] != s[hi] {
			return false
		}
	}
	return true
}

func strip(s string) string {
	var result strings.Builder
	result.Grow(len(s))
	for i := 0; i < len(s); i++ {
		b := s[i]
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') ||
			('0' <= b && b <= '9') ||
			b == ' ' {
			result.WriteByte(b)
		}
	}
	return strings.ToLower(strings.ReplaceAll(result.String(), " ", ""))
}

func main() {
	s1 := "A man, a plan, a canal: Panama"
	s2 := "race a car"
	s3 := " "

	fmt.Println(isPalindrome3(s1))
	fmt.Println(isPalindrome3(s2))
	fmt.Println(isPalindrome3(s3))
}
