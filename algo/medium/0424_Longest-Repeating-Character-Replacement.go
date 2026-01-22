package main

func characterReplacement(s string, k int) int {
	freqMap := make(map[byte]int)
	res, left, maxFreq := 0, 0, 0

	for right := 0; right < len(s); right++ {
		freqMap[s[right]]++
		maxFreq = max(maxFreq, freqMap[s[right]])

		for (right-left+1)-maxFreq > k {
			freqMap[s[left]]--
			left++
		}

		if right-left+1 > res {
			res = right - left + 1
		}
	}

	return res
}
