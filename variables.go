// go run variables.go

package main

import "fmt"

var name = "Miro" // visible to entire package | package level block

func main() {
	// scope main
	x := 10.0
	y := "bom dia"

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n\n", y, y)

	x, z := 20.3, 31 // allowed to use := because z is not yet defined

	fmt.Println(x, z)
	

	fmt.Printf("Hi, %v\n", name) // use the package level block var

	sum := 10 + x


	fmt.Printf("result of 10 with x: %v\n", sum);


	isEvenNumber := z % 2 == 0

	fmt.Printf("%v is even number: %v", z, isEvenNumber);
	
	
}
