package main

import "fmt"

type test int
var x test
var y int

func main() {
	fmt.Printf("%v -> %T\n", x, x)
	x = 42
	fmt.Println(x)

	y = int(x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%T\n", y)
}