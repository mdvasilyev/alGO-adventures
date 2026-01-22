package main

import "fmt"

// Time:	O(len(s))
// Space:	O(len(s))
func romanToInt(s string) int {
	if len(s) == 0 {
		return 0
	}
	res := 0
	symbols := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var prev rune
	for i, val := range s {
		if i == 0 {
			prev = val
			continue
		}
		if symbols[val] <= symbols[prev] {
			res += symbols[prev]
		} else {
			res -= symbols[prev]
		}
		prev = val
	}
	res += symbols[prev]
	return res
}

func main() {
	s1 := "III"
	s2 := "LVIII"
	s3 := "MCMXCIV"

	fmt.Println(romanToInt(s1))
	fmt.Println(romanToInt(s2))
	fmt.Println(romanToInt(s3))
}
