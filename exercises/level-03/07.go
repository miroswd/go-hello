package main 


import "fmt"


func main() {
	name := "John"

	if name == "Miro" {
		fmt.Print("Hello, ", name)
	} else if name == "John" {
		fmt.Print("Hi, little ", name)
	} else {
		fmt.Printf("Welcome %v", name)
	}
}