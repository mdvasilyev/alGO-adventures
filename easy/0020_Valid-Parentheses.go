package main

import "fmt"

// Time:	O(len(s))
// Space:	O(len(s))
func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}
	m := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	stack := make([]rune, 0)

	for _, val := range s {
		if _, ok := m[val]; ok {
			stack = append(stack, val)
		} else if len(stack) == 0 || val != m[stack[len(stack)-1]] {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func main() {
	s1 := "()"
	s2 := "()[]{}"
	s3 := "(]"
	s4 := "([])"
	s5 := "]"

	fmt.Println(isValid(s1))
	fmt.Println(isValid(s2))
	fmt.Println(isValid(s3))
	fmt.Println(isValid(s4))
	fmt.Println(isValid(s5))
}
