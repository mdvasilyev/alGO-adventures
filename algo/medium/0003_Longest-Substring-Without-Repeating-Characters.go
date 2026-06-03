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

// func lengthOfLongestSubstring(s string) int {
//     length := len(s)
//     if length <= 1 {
//         return length
//     }
//     res := 0
//     m := make(map[byte]int)
//     left, right := 0, 1
//     m[s[left]]++
//     for right < length {
//         val := s[right]
//         if cnt, ok := m[val]; ok && cnt > 0 {
//             m[s[left]]--
//             left++
//         } else {
//             m[val]++
//             res = max(res, right-left+1)
//             right++
//         }
//     }
//     return res
// }

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
