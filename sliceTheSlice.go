package main

import "fmt"


func main() {
	flavors := []string{ "pepperoni", "muçarela", "abacaxi", "marguerita", "calabresa"  }

	slice := flavors[2:len(flavors)]

	fmt.Println(slice)

	// delete items

	flavors = append(flavors[:2], flavors[3:]...)
	fmt.Println(flavors)

}