package main

import "fmt"

// Time:	O(n)
// Space:	O(1)
func checkInclusion(s1 string, s2 string) bool {
	s1Len, s2Len := len(s1), len(s2)
	if s1Len > s2Len {
		return false
	}
	s1Map, s2Map := [26]int{}, [26]int{}

	for i := 0; i < s1Len; i++ {
		s1Map[s1[i]-'a']++
		s2Map[s2[i]-'a']++
	}

	for i := 0; i < s2Len-s1Len; i++ {
		if s1Map == s2Map {
			return true
		}
		s2Map[s2[i]-'a']--
		s2Map[s2[i+s1Len]-'a']++
	}

	return s1Map == s2Map
}

func main() {
	s1 := "adc"
	s2 := "dcda"
	fmt.Println(checkInclusion(s1, s2))
}
