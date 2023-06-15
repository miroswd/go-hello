package main

import "fmt"

func main() {
	slice := []int{1,2,3,4,5,6}

	fmt.Println(slice)

	secondSlice := append(slice[:2], slice[4:]...)

	fmt.Println(secondSlice)
	fmt.Println(slice)
}