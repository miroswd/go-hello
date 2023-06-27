package main

import "testing"

func sum(a, b int) int {
	return a + b
}

func testSum(t *testing.T) {
	v := sum(1,2)
	
	if v != 3 {
		t.Error("Expected 3, got", v)
	}

}
