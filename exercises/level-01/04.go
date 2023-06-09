package main

import "fmt"

type test int
var x test

func main() {
	fmt.Printf("%v -> %T\n", x, x)
	x = 42
	fmt.Println(x)
}