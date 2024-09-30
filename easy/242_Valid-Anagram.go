package main

import "fmt"

// Time:	O(len(s))
// Space:	O(len(s))
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m1 := make(map[rune]int)
	m2 := make(map[rune]int)
	for _, val := range s {
		m1[val]++
	}
	for _, val := range t {
		m2[val]++
	}
	for key := range m1 {
		if m1[key] != m2[key] {
			return false
		}
	}
	return true
}

func main() {
	s1, t1 := "anagram", "nagaram"
	s2, t2 := "rat", "car"

	fmt.Println(isAnagram(s1, t1))
	fmt.Println(isAnagram(s2, t2))
}
