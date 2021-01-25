package main

import "fmt"

func main() {
	staircase(6)
}

// Complete the staircase function below.
func staircase(n int32) {
	for i := int32(0); i < n; i++ {
		for j := int32(1); j < n-i; j++ {
			fmt.Print(" ")
		}

		for j := n - i; j <= n; j++ {
			fmt.Print("#")
		}

		fmt.Println("")
	}
}
