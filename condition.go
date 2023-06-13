package main


import (
	"fmt"
)


func main() {
	x := 10

	if !(x < 100) {
		fmt.Println("Hello World")
	}

	if y := 100; y == 100 {
		fmt.Println("Hi")
	} 

	z := 18

	if z > 18 {
		fmt.Println("More than 18")
	} else if z < 18 {
		fmt.Println("Less than 18")
	} else {
		fmt.Println("Equal to 18")
	}
}