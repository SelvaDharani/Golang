package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// In your Go test file
func TestMyFunction(t *testing.T) {
	// Test case 1: Positive inputs
	sum, err := MyFunction(2, 3)
	assert.Equal(t, 5, sum, "Sum of 2 and 3 is 5")
	assert.NoError(t, err, "There are no errors")

}

// In your Go test file
func TestMyFunction2(t *testing.T) {
	// Test case 2: One negative input
	sum, err := MyFunction(-2, 3)
	assert.Equal(t, 0, sum, "Sum of -2 and 3 is 0")
	assert.Error(t, err, "Error is not nil for negative input")

}
