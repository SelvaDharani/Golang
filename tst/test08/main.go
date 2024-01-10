// main.go
package main

import "fmt"

func AddToEachElement(slice []int, increment int) []int {
	result := make([]int, len(slice))
	for i, v := range slice {
		result[i] = v + increment
	}
	return result
}

func main() {

	originalSlice := []int{1, 2, 3}
	increment := 2
	processedSlice := AddToEachElement(originalSlice, increment)
	fmt.Println(processedSlice)
}
