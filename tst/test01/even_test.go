package main

import "testing"

func TestIsEven(t *testing.T) {
	if !IsEven(4) {
		t.Errorf("Expected true for 4, got false")
	}

	if IsEven(5) {
		t.Error("Expected false for 5, got true")
	}

	if !IsEven(6) {
		t.Log("6 should be even, but test failed")
		t.Fail()
		t.Log("After t.Fail, this will still execute")
	}

	if IsEven(7) {
		t.Log("Before t.FailNow for 7")
		t.FailNow()
		t.Log("After t.FailNow, this will not execute")
	}

}
