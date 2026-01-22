package main

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}

	res := make([]int, 0)
	m, cur := [26]int{}, [26]int{}

	for i := 0; i < len(p); i++ {
		m[p[i]-'a']++
		cur[s[i]-'a']++
	}

	for i := 0; i < len(s)-len(p)+1; i++ {
		if m == cur {
			res = append(res, i)
		}
		if i+len(p) < len(s) {
			cur[s[i]-'a']--
			cur[s[i+len(p)]-'a']++
		}
	}

	return res
}
