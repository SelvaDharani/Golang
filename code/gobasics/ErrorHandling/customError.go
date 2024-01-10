package main

import (
	"fmt"
	"strconv"
)

// The main function where program execution begins
func main() {
	// Call a function that might return a custom error
	result, err := returnCustomError(false)
	// Check for errors
	if err != nil {
		// Print the error message if an error occurred
		fmt.Println(err)
	} else {
		// Print the result if no error
		fmt.Println(result)
	}

	// Call the same function again, forcing an error this time
	result, err = returnCustomError(true)
	// Handle the potential error
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

// Define a custom error type with additional fields
type customError struct {
	errorMessage string
	errorCode    int
}

// Implement the Error() method to provide a custom error message
func (c customError) Error() string {
	// Combine the error message and code into a single string
	return c.errorMessage + " : " + strconv.Itoa(c.errorCode)
}

// Function that returns a string and potentially a custom error
func returnCustomError(returnError bool) (string, error) {
	if returnError {
		// Create a custom error and return it
		return "", customError{errorMessage: "Custom Error Message", errorCode: 1234}
	}
	// Return a string and a nil error (no error)
	return "Some Random String", nil
}
