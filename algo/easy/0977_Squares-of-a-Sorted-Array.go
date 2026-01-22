package main

func sortedSquares(nums []int) []int {
	res := make([]int, 0)
	left, right := 0, len(nums)-1
	for left <= right {
		leftSquare, rightSquare := nums[left]*nums[left], nums[right]*nums[right]
		if leftSquare >= rightSquare {
			res = append(res, leftSquare)
			left++
		} else {
			res = append(res, rightSquare)
			right--
		}
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
