package main

import (
	"fmt"
)

// MyFunction returns the sum of two integers.
func MyFunction(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, fmt.Errorf("both arguments must be non-negative")
	}
	return a + b, nil
}

// func main() {
// 	result, err := MyFunction(2, 3)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Println("Result:", result)
// }
