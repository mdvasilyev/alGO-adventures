package main

import (
	"fmt"
	"maps"
)

// Time:	O(len(s))
// Space:	O(len(s))
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sMap := make(map[byte]int)
	tMap := make(map[byte]int)
	for i := range s {
		sMap[s[i]]++
		tMap[t[i]]++
	}
	return maps.Equal(sMap, tMap)
}

func main() {
	s1, t1 := "anagram", "nagaram"
	s2, t2 := "rat", "car"

	fmt.Println(isAnagram(s1, t1))
	fmt.Println(isAnagram(s2, t2))
}
