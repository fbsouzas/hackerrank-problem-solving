package main

import "fmt"

func main() {
	miniMaxSum([]int32{1, 2, 3, 4, 5})
	miniMaxSum([]int32{1, 3, 5, 7, 9})
	miniMaxSum([]int32{1, 2, 3})
	miniMaxSum([]int32{100, 101})
	miniMaxSum([]int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15})
	miniMaxSum([]int32{999999990, 999999991, 999999992, 999999993, 999999994, 999999995, 999999996})
}

func miniMaxSum(arr []int32) {
	min := int32(1)
	max := int32(1)
	minSum := int64(0)
	maxSum := int64(0)
	size := int32(len(arr))

	for i := int32(0); i < size; i++ {
		if arr[i] < arr[min] || arr[min] == 0 {
			min = i
		}

		if arr[i] > arr[max] {
			max = i
		}
	}

	for i, v := range arr {
		if int32(i) != min {
			minSum += int64(v)
		}

		if int32(i) != max {
			maxSum += int64(v)
		}
	}

	if minSum > maxSum {
		fmt.Println(maxSum, minSum)
	} else {
		fmt.Println(minSum, maxSum)
	}
}
