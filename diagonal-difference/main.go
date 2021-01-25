package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(diagonalDifference([][]int32{{1, 2, 3}, {4, 5, 6}, {9, 8, 9}}))
	fmt.Println(diagonalDifference([][]int32{{11, 2, 4}, {4, 5, 6}, {10, 8, -12}}))
	fmt.Println(diagonalDifference([][]int32{{-1, 1, -7, -8}, {-10, -8, -5, -2}, {0, 9, 7, -1}, {4, 4, -2, 1}}))
	fmt.Println(diagonalDifference([][]int32{{-1, 1, -7, -8, 3}, {-10, -8, -5, -2, 3}, {0, 9, 7, -1, 3}, {4, 4, -2, 1, 3}, {2, 3, 4, 5, 6}}))
}

func diagonalDifference(arr [][]int32) int32 {
	var leftSum int32 = 0
	var rightSum int32 = 0

	size := len(arr)

	for i := 0; i < size; i++ {
		left := 0 + i
		right := (size - 1) - i

		leftSum += arr[i][left]
		rightSum += arr[i][right]
	}

	return int32(math.Abs(float64(leftSum) - float64(rightSum)))
}
