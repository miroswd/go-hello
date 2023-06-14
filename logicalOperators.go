package main

import "fmt"


func main() {
	x := 4
	
	if x == 2 || x == 3 {
		fmt.Println("Hello World")
	}


	if x % 2 == 0 && x != 2 {
		fmt.Println("Hi!")
	}


	if y := 4; !(y + x % 2 == 0) {
		fmt.Println("Hey there")
	}

}