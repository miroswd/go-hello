package main

import ("fmt")

var a int = 200

func main() {

	b := a << 1

	fmt.Printf("%d - %#b - %#x\n", a, a, a)
	fmt.Printf("%d - %#b - %#x\n", b, b, b)
	
}
