package main

import (
	"fmt"
)


var a interface{}


func main() {

	x := 1

	switch { // without expression, switch true {
		case x < 5:
			fmt.Println("Less than 5")
		case x == 5:
			fmt.Println("Equal to 5")
		case x > 5:
			fmt.Println("More than 5")
	}


	whoIAm := "miro"
 
	switch whoIAm {
		case "miro":
			fmt.Println(`Hi miro`);
		case "john":
			fmt.Println("Hello john");
		default:
			fmt.Println("Hey there")
	}


	switch x {
		case 1, 2, 3:
			fmt.Println("Forbidden number")
		case 4, 5:
			fmt.Println("Ok")
		default: 
			fmt.Println("More than expected")
	}


	a = "true";

	switch a.(type) {
		case int: 
			fmt.Println("Is int")
		case bool: 
			fmt.Println("Is bool")
		case string: 
			fmt.Println("Is string")
		case float64:
			fmt.Println("Is float")

	}


	switch b := 0; {
		case b == 1:
			fmt.Println("On")
		case b == 0:
			fmt.Println("Off")
	}

}