// go run variables.go

package main

import "fmt"

var name = "Miro" // visible to entire package | package level block
var zzz = 56

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

	isEvenNumber(z)
}



func isEvenNumber (x int){
	// getting the z variable outside the code block, nobody sees the variable
	fmt.Println(zzz); // using the package level block


	isEvenNumber := x % 2 == 0;
	fmt.Printf("\n%v is even number: %v", x, isEvenNumber);
}
