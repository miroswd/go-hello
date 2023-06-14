package main


import "fmt"


func main() {
	slice := []int{1,2,3,4}
	anotherSlice := []int{9, 10, 11, 12}


	slice = append(slice, 5, 6, 7, 8)


	fmt.Printf("%v  - %T", slice, slice)
	fmt.Println()

	/*
	slice = append(slice, anotherSlice)
	// cannot use anotherSlice (type []int) as type int in append
	*/


	slice = append(slice, anotherSlice...) // enumeration || unfurl || rest
	fmt.Printf("%v", slice)


}