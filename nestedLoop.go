package main

import ("fmt")


// Loop inside loop
func main() {

	for hours := 0; hours <= 12; hours++ {
		fmt.Print("Hour: ", hours)

		for minutes := 0; minutes < 60; minutes++ {
			fmt.Print(" ", minutes)
		}

		fmt.Println("\n")
	}

}
