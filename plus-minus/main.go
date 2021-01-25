package main

import "fmt"

func main() {
	plusMinus([]int32{-4, 3, -9, 0, 4, 1})
	fmt.Println("")
	plusMinus([]int32{1, 2, 3, -1, -2, -3, 0, 0})
}

func plusMinus(arr []int32) {
	size := len(arr)
	var positive float32 = 0
	var negative float32 = 0
	var zero float32 = 0

	for _, v := range arr {
		if v > 0 {
			positive++
		}

		if v < 0 {
			negative++
		}

		if v == 0 {
			zero++
		}
	}

	fmt.Printf("%.6f\n", positive/float32(size))
	fmt.Printf("%.6f\n", negative/float32(size))
	fmt.Printf("%.6f\n", zero/float32(size))
}
