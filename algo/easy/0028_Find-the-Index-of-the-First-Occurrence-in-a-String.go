package main

func strStr(haystack string, needle string) int {
	res := -1
	lenH, lenN := len(haystack), len(needle)
	if lenN > lenH {
		return res
	}
	for i := 0; i <= lenH-lenN; i++ {
		if haystack[i:i+lenN] == needle {
			return i
		}
	}
	return res
}
