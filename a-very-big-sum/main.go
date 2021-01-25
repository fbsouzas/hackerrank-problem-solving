package main

import "fmt"

func main() {
	result := aVeryBigSum([]int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005})

	fmt.Println(result)
}

// Complete the aVeryBigSum function below.
func aVeryBigSum(ar []int64) int64 {
	var summed int64

	summed = 0

	for _, v := range ar {
		summed += v
	}

	return summed
}
