package main

import (
	"fmt" 
	"runtime"
)


func main() {
	a := "e"
	b := "é"

	fmt.Printf("%v, %v\n", a, b)

	c := []byte(a) // convert to byte
	d := []byte(b)

	fmt.Printf("%v, %v", c, d)



	x := 10
	y := 10.1

	fmt.Printf("%v, %T\n", x, x)
	fmt.Printf("%v, %T\n", y, y)


	fmt.Println(runtime.GOOS) // Get the operating system running GO
	fmt.Println(runtime.GOARCH) // Represents the architecture of the processor (amd64: 64 bits)

	// add a more byte to limited type

	var i uint16 //  unsigned integers - can only hold non-negative values (zero and positive integers
	// int - can hold both positive and negative values
	// i = 65536 // constant 65536 overflows uint16
	i = 65535
	
	fmt.Println(i)
	i++

	// When we add the maximum byte of the type, the value is zeroed. Similar to a 3-digit count, when we reach 999 and add 1, the counter goes to zero
	fmt.Println(i)

	i++
	fmt.Println(i)



}