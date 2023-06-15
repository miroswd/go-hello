package main

import "fmt"

func main() {

	slice := make([]int, 5, 10);

	slice[0], slice[1], slice[2] = 1, 2, 3

	slice = append(slice, 10)
	slice = append(slice, 10)
	slice = append(slice, 10)
	slice = append(slice, 10)
	slice = append(slice, 10)

	// exceeded the capacity

	// make creates a new array, based on the current one, increases the capacity and discards the previous one	

	slice = append(slice, 10)
	slice = append(slice, 10)

	fmt.Println(slice, len(slice), cap(slice))

}