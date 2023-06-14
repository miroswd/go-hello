package main

import "fmt"

func main() {
	array := [5]int{1,2,3,4,5}
	fmt.Println(array)

	slice := []int{1,2,3,4,5}
	fmt.Println(slice)


	// array2 = append(array, 6) -> is not allowed add item to array 5 int
	
	slice2 := append(slice, 6)
	slice2[3] = 11
	fmt.Println(slice2)

}