package main

import "fmt"

// Time:	O(log(n*m))
// Space:	O(n*m)
func searchMatrix(matrix [][]int, target int) bool {
	n := len(matrix)
	m := len(matrix[0])
	lo := 0
	hi := n*m - 1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		val := matrix[mid/m][mid%m]
		if val < target {
			lo = mid + 1
		} else if val > target {
			hi = mid - 1
		} else {
			return true
		}
	}
	return false
}

func main() {
	matrix1, target1 := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3
	matrix2, target2 := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13
	matrix3, target3 := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}, 11
	matrix4, target4 := [][]int{{1}, {2}, {3}}, 3
	matrix5, target5 := [][]int{{1}}, 1

	fmt.Println(searchMatrix(matrix1, target1))
	fmt.Println(searchMatrix(matrix2, target2))
	fmt.Println(searchMatrix(matrix3, target3))
	fmt.Println(searchMatrix(matrix4, target4))
	fmt.Println(searchMatrix(matrix5, target5))
}
