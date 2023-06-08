package main

import "fmt"

func main() {
	numberOfBytes, error := fmt.Println("Hello World", "miro", 3)
	fmt.Println(numberOfBytes, error)

	// type
	x := 3
	y := "string"
	z := true

	fmt.Println(x, y, z)
}
