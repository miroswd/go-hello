package main

import ("fmt")

const x = 10
// var y = 10


// Declare multiple const

const (
	a = 1
	b = "Hello"
	cc = true
)

var c int
// var d float64

func main() {
 	c = x

	fmt.Println(c)
	fmt.Println(a, b, cc)
}

// The type of a constant is only defined at the time of use, as for a var it is at the time of assignment