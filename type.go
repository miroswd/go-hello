package main

import "fmt"


type hotdog int 
var b hotdog = 10 // the base type behind b is int

var x int = 10
var y bool // we cannot assign the y var outside the scope
// y = false // it will not work
func main () {
	//x = 10.5 // failed to run this, because GO has a static typing
	x = 15

	y = false

	fmt.Println(x, y)
	fmt.Printf("%T: %v", b, b)

	// b = x //cannot use x (type int) as type hotdog in assignment
}