package main


import (
	"fmt"
)


func main(){
	x := 1
	y := x >> 1 // shift bit to the right
	z := x << 1 // shift bit to the left

	fmt.Printf("%b\n %v", x, x)
	fmt.Printf("%b\n", y)
	fmt.Printf("%b\n %v", z, z)
}