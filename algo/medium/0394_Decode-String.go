package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func decodeString(s string) string {
	res, _ := decodeSubstring(s, 0)
	return res
}

func decodeSubstring(s string, start int) (string, int) {
	res := strings.Builder{}
	numStr := strings.Builder{}
	num := 0
	i := 0
	for i = start; i < len(s); i++ {
		ch := rune(s[i])
		if unicode.IsDigit(ch) {
			numStr.WriteRune(ch)
		} else if ch == '[' {
			num, _ = strconv.Atoi(numStr.String())
			numStr.Reset()
			subS, newI := decodeSubstring(s, i+1)
			res.WriteString(strings.Repeat(subS, num))
			i = newI
		} else if ch == ']' {
			break
		} else {
			res.WriteRune(ch)
		}
	}
	return res.String(), i
}

func main() {
	s0 := "2[a]"
	s1 := "3[a]2[bc]"     // aaabcbc
	s2 := "3[a2[c]]"      // accaccacc
	s3 := "2[abc]3[cd]ef" // abcabccdcdcdef

	fmt.Println(decodeString(s0))
	fmt.Println(decodeString(s1))
	fmt.Println(decodeString(s2))
	fmt.Println(decodeString(s3))
}
