package main

func lengthOfLastWord(s string) int {
	res := 0
	for i := len(s) - 1; i >= 0; i-- {
		ch := s[i]
		if ch == ' ' && res != 0 {
			break
		}
		if ch != ' ' {
			res++
		}
	}
	return res
}
