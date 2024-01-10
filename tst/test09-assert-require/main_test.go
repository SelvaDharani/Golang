// // main_test.go
package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIntToString(t *testing.T) {
	// Test with a positive number using assert
	fmt.Println("Testing with a positive number...")
	result, err := IntToString(10)
	assert.NoError(t, err, "Expecting no error for positive number")
	assert.Equal(t, "10", result, "Expecting '10' for input 10")
	fmt.Println("Asserts completed for positive number")

	fmt.Println("-------------")
	// Test with a positive number using require
	fmt.Println("Testing with a positive number...")
	result2, err := IntToString(15)
	require.NoError(t, err, "Expecting no error for positive number")
	require.Equal(t, "15", result2, "Expecting '10' for input 10")
	fmt.Println("Require assertions completed for positive number")

	fmt.Println("----------------")
	// Test with a negative number using require
	fmt.Println("Testing with a negative number...")
	_, err = IntToString(-11)
	require.Error(t, err, "Expecting an error for negative number")
	fmt.Println("Require completed for negative number")

	// This line will only execute if require.Error above does not fail
	fmt.Println("This line is executed only if require.Error passes")
}
