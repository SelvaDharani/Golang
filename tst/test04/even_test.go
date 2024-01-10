// main_test.go
package main

import "testing"

func TestIsEven(t *testing.T) {
	// Correct case
	//t.Errorf reports an error with formatting and continues the test.
	if !IsEven(4) {
		t.Errorf("Expected true for 4, got false")
	}

	// Incorrect case, demonstrate t.Error
	//t.Error reports an error without formatting and continues the test.
	if IsEven(5) {
		t.Error("Expected false for 5, got true")
	}

	// Incorrect case, demonstrate t.Fail
	if !IsEven(6) {
		t.Log("6 should be even, but test failed")
		t.Fail()
		t.Log("After t.Fail, this will still execute")
	}

	// Incorrect case, demonstrate t.FailNow
	if IsEven(7) {
		t.Log("Before t.FailNow for 7")
		t.FailNow()
		t.Log("After t.FailNow, this will not execute")
	}

	// This line will not be reached if t.FailNow() is called for 7
	t.Log("End of TestIsEven")
}
