package main

import (
	"testing"
)


func TestSum(t *testing.T) {
	expected := 6
	result := sum(3,2,1)

	if expected != result {
		t.Error("Expected:", expected, "Got:", result)
	}
}


func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum(1,1)
	}
}

func BenchmarkMultiply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		multiply(1,1)
	}
}