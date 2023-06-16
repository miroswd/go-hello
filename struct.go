package main

import "fmt"

type customer struct {
	name string
	lastName string
	smoker bool
}


func main() {
	customer1 := customer{
		name: "john",
		lastName: "doe",
		smoker: false,
	}

	customer2 := customer{
		"xuxa",
		"da silva",
		true,
	}


	fmt.Println(customer1)
	fmt.Println(customer2)
}