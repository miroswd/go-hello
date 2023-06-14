package main


import "fmt"


func main() {
	x := 2

	switch {
	case x == 1: 
		fmt.Println("X equal 1")
	case x == 2: 
		fmt.Println("X equal 2")
	default:
		fmt.Printf("%v is invalid", x)
	}
}