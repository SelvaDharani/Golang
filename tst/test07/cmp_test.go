// main_test.go
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHigherTemperature(t *testing.T) {
	higher := HigherTemperature(30.0, 25.0)
	assert.Greater(t, higher, 25.0, "The higher temperature should be greater than 25.0")
	assert.Equal(t, 30.0, higher, "The higher temperature should be equal to 30.0")
}

func TestLowerTemperature(t *testing.T) {
	lower := LowerTemperature(18.0, 22.0)
	assert.Less(t, lower, 22.0, "The lower temperature should be less than 22.0")
	assert.Equal(t, 18.0, lower, "The lower temperature should be equal to 18.0")
}
