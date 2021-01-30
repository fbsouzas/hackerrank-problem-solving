package main

import "fmt"

func main() {
	fmt.Println(birthdayCakeCandles([]int32{2, 3, 1, 3}))
	fmt.Println(birthdayCakeCandles([]int32{3, 2, 1, 3}))
	fmt.Println(birthdayCakeCandles([]int32{4, 4, 1, 3}))
	fmt.Println(birthdayCakeCandles([]int32{10000, 1, 1, 1, 5000000, 10000}))
	fmt.Println(birthdayCakeCandles([]int32{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}))
	fmt.Println(birthdayCakeCandles([]int32{10000000, 10000000, 10000000, 50, 50, 50, 10000000}))
	fmt.Println(birthdayCakeCandles([]int32{0, 0, 0, 1, 1}))
	fmt.Println(birthdayCakeCandles([]int32{0, 0, 0, 0, 1}))
	fmt.Println(birthdayCakeCandles([]int32{1, 1, 1, 1, 2}))
}

func birthdayCakeCandles(candles []int32) int32 {
	tallest := []int32{}
	tallest = append(tallest, 0)

	for _, candle := range candles {
		if candle == tallest[0] {
			tallest = append(tallest, candle)
		}

		if candle > tallest[0] {
			tallest = []int32{}
			tallest = append(tallest, candle)
		}
	}

	return int32(len(tallest))
}
