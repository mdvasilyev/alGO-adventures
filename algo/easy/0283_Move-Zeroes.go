package main

import "fmt"

func moveZeroes(nums []int)  {
    left, right := 0, len(nums)-1
    for left < right {
        for right >= 0 && nums[right] == 0 {
            right--
        }
        if right >= 0 {
            nums[left], nums[right] = nums[right], nums[left]
        }
        left++
        right--
    }
    firstZeroPos := len(nums)
    for i := range nums {
        if nums[i] == 0 {
            firstZeroPos = i
            break
        }
    }
    for i, j := 0, firstZeroPos-1; i < j; i, j = i+1, j-1 {
        nums[i], nums[j] = nums[j], nums[i]
    }
}

func main() {
	nums := []int{0,1,0,3,12}
	moveZeroes(nums)
	fmt.Println(nums)
}