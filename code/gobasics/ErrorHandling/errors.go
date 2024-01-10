package main

import (
	"errors"
	"fmt"
)

func main() {
	// Call a function that returns a single error value
	err := returnErrorSingleValue()
	// Check if an error occurred
	if err != nil {
		// Print the error message
		fmt.Println(err)
	}

	// Call a function that returns multiple values, including an error
	result, err := returnErrorMultipleValue(false)
	// Check for errors
	if err != nil {
		fmt.Println(err)
	} else {
		// If no error, print the result
		fmt.Println(result)
	}

	// Call the same function again, forcing an error this time
	result, err = returnErrorMultipleValue(true)
	// Handle potential errors
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

// Function that always returns an error
func returnErrorSingleValue() error {
	// Create a new error with the message "Error Here"
	return errors.New("Error Here")
}

// Function that returns a string and an error, depending on input
func returnErrorMultipleValue(returnError bool) (string, error) {
	if returnError {
		// Return an empty string and an error if returnError is true
		return "", errors.New("Error Here")
	}
	// Otherwise, return a string and a nil error
	return "Some Random String", nil
}
