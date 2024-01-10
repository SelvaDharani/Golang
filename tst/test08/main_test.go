// main_test.go
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddToEachElement(t *testing.T) {
	originalSlice := []int{1, 2, 3}
	increment := 2
	processedSlice := AddToEachElement(originalSlice, increment)

	// Using assert.Contains to check if a specific element is in the slice
	assert.Contains(t, processedSlice, 3, "Processed slice should contain 3")

	// Using assert.NotContains to check if a specific element is not in the slice
	assert.NotContains(t, processedSlice, 1, "Processed slice should not contain 1")

	// Using assert.Len to check the length of the slice
	assert.Len(t, processedSlice, 3, "Processed slice should have 3 elements")
}
