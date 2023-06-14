package main

import "fmt"

func main() {
	esporteFavorito := "futebol"

	switch esporteFavorito {
	case "futebol":
		fmt.Println("Come to fut, come!")
	case "basquete":
		fmt.Println("Who put the basketball on the electric wire?")
	}
}