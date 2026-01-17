package main

import "fmt"

func canJump(nums []int) bool {
	possibleJumpTo := make([]bool, len(nums))
	possibleJumpTo[0] = true
	for i := 0; i < len(nums)-1; i++ {
		if possibleJumpTo[i] == false {
			continue
		}
		maxJump := nums[i]
		for j := 1; j < maxJump+1; j++ {
			if i+j >= len(nums)-1 {
				return true
			}
			possibleJumpTo[i+j] = true
		}
	}
	return possibleJumpTo[len(nums)-1]
}

func main() {
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
}
