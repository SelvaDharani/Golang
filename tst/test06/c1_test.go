package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	assert.Equal(t, 5, result, "Add function result should be 5")
}

func TestSubtract(t *testing.T) {
	result := Subtract(5, 3)
	assert.NotEqual(t, 1, result, "Subtract function result should not be 1")
}

func TestDivide(t *testing.T) {
	result, err := Divide(6, 2)
	assert.Nil(t, err, "Divide function should not return an error")
	assert.Equal(t, 3, result, "Divide function result should be 3")

	_, err = Divide(6, 0)
	assert.NotNil(t, err, "Divide by zero should return an error")
	assert.True(t, ErrDivideByZero == err, "Error should be divide by zero error")
}
