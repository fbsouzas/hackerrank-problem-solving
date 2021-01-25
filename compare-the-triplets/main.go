package main

import "fmt"

func main() {
	result := compareTriplets([]int32{1, 2, 3}, []int32{3, 2, 1})

	fmt.Println(result)
}

func compareTriplets(a []int32, b []int32) []int32 {
	var alice int32
	var bob int32

	alice = 0
	bob = 0

	for i := range a {
		if a[i] > b[i] {
			alice++
		}

		if b[i] > a[i] {
			bob++
		}
	}

	return []int32{alice, bob}
}
