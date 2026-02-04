package main

func rotateMatrix(matrix [][]int) {
	length := len(matrix)
	lDiff := length - 1
	for i := 0; i < length/2; i++ {
		for j := i; j < lDiff-i; j++ {
			matrix[i][j], matrix[j][lDiff-i], matrix[lDiff-i][lDiff-j], matrix[lDiff-j][i] = matrix[lDiff-j][i], matrix[i][j], matrix[j][lDiff-i], matrix[lDiff-i][lDiff-j]
		}
	}
}

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotateMatrix(matrix)
	print(matrix)
}
