package main

import (
	"fmt"
	"testing"
)

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




// Table test


type test struct {
	data []int
	answer int
}


func TestSumTable(t *testing.T) {
	tests := []test{
		test{data: []int{1,2,3}, answer:6},
		test{[]int{10,11,12},33},
		test{[]int{-5,0,5,10},10},
	}

	for _, v := range tests {
		if x := sum(v.data...); v.answer != x  {
			t.Error("Expected:", v.answer, "Got:", x)

		}
	}

}


// Test as an example

func ExampleSum(){
	fmt.Println(sum(3,2,1))
	fmt.Println(sum(2,2,1))
	fmt.Println(sum(4,2,1))
	// Output:
	// 6
	// 5
	// 7

	fmt.Println(sum(3,2,1))
	// Output: 10

}