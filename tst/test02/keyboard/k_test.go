package main

import "testing"

func TestKeyboard(t *testing.T) {
	expected := "Keys are working"
	actual := keyboard()

	if expected != actual {
		t.Errorf("Expected %s do not match actual %s", expected, actual)
	}

}
