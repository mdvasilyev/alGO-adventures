package main

func hIndex(citations []int) int {
	length := len(citations)
	hIndexes := make([]int, length+1)
	for i := range citations {
		cit := citations[i]
		if cit >= length {
			hIndexes[length]++
		} else {
			hIndexes[cit]++
		}
	}
	cnt := 0
	for i := length; i >= 0; i-- {
		cnt += hIndexes[i]
		if cnt >= i {
			return i
		}
	}
	return 0
}
