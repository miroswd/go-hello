package main

import "fmt"

var x int = 10
var y bool // we cannot assign the y var outside the scope
// y = false // it will not work
func main () {
	//x = 10.5 // failed to run this, because GO has a static typing
	x = 15

	y = false

	fmt.Println(x, y)
}