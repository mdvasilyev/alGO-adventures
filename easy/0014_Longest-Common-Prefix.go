package main

func longestCommonPrefix(strs []string) string {
	for i := range strs[0] {
		for j := range strs {
			if i == len(strs[j]) || strs[0][i] != strs[j][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
