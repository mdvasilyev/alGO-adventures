package main

func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	counterMap := make(map[[26]int][]int)
	for i := range strs {
		counter := count(strs[i])
		counterMap[counter] = append(counterMap[counter], i)
	}
	for _, val := range counterMap {
		cur := make([]string, len(val))
		for i := range val {
			cur[i] = strs[val[i]]
		}
		res = append(res, cur)
	}
	return res
}

func count(s string) [26]int {
	res := [26]int{}
	for i := range s {
		res[s[i]-'a']++
	}
	return res
}
