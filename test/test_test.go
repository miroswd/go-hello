package main

import "testing"

func TestSum(t *testing.T) {
	expected := 6
	result := sum(3,2,1)

	if expected != result {
		t.Error("Expected:", expected, "Got:", result)
	}
}


func TestSumFailed(t *testing.T) {
	expected := 0
	result := sum(3,2,1)

	if expected != result {
		t.Error("Expected:", expected, "Got:", result)
	}
}

func TestFailedMultiply(t *testing.T) {
	expected := 100
	result := multiply(10, 10)

	if expected != result {
		t.Error("Expected:", expected, "Got:", result)
	}
}