package main


import "fmt"


func main() {

	x := struct{
		name string
		age int
	}{"Miro", 23,} // implicit declaration

	fmt.Print(x)
}