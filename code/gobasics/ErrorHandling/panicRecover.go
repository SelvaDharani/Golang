package main

import "fmt"

func main() {
	// Declare constants for demonstration
	const one = 1
	const two = 1

	// Defer the saveExit function for panic recovery (first instance)
	defer saveExit()

	// Check if one is not equal to 1 (will not panic)
	if one != 1 {
		panic("one is not equal to 1")
	}

	// Defer the saveExit function again (second instance)
	defer saveExit()

	// Check if two is not equal to 2 (also will not panic)
	// will cause panic as variable two is equal to 1 in line no. 8
	if two != 2 {
		panic("two is not equal to 2")
	}
}

// Function to handle panics
func saveExit() {
	// Use recover() to regain control after a panic
	if r := recover(); r != nil {
		fmt.Println("Panic Is Recovered")
	}
}
