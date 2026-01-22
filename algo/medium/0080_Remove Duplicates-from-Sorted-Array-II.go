package main

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length <= 2 {
		return length
	}

	write := 2

	for read := 2; read < length; read++ {
		if nums[read] != nums[write-2] {
			nums[write] = nums[read]
			write++
		}
	}

	return write
}

//func removeDuplicates(nums []int) int {
//	unique := 1
//	length := len(nums)
//	if length == 1 {
//		return unique
//	}
//	left, right := 1, 1
//	curLen := 1
//	for right < length {
//		if nums[left-1] == nums[right] {
//			if curLen > 1 {
//				right++
//			} else {
//				nums[left], nums[right] = nums[right], nums[left]
//				curLen++
//				left++
//				right++
//				unique++
//			}
//		} else {
//			nums[left], nums[right] = nums[right], nums[left]
//			curLen = 1
//			left++
//			right++
//			unique++
//		}
//	}
//	return unique
//}
