package main

import "strings"

func reverseWords(s string) string {
	res := &strings.Builder{}
	cur := &strings.Builder{}
	for i := len(s) - 1; i >= 0; i-- {
		ch := s[i]
		if ch == ' ' && cur.Len() > 0 {
			writeReversedWord(cur, res)
		}
		if ch != ' ' {
			cur.WriteByte(ch)
		}
	}
	if cur.Len() > 0 {
		writeReversedWord(cur, res)
	}
	return res.String()
}

func writeReversedWord(cur, res *strings.Builder) {
	word := reverseWord(cur.String())
	if res.Len() > 0 {
		res.WriteString(" ")
	}
	res.WriteString(word)
	cur.Reset()
}

func reverseWord(s string) string {
	sList := []byte(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		sList[i], sList[j] = sList[j], sList[i]
	}
	return string(sList)
}
