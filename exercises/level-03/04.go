package main

import "fmt"

func main() {
	bornYear := 2000
	currentYear := 2023

	for {
		fmt.Printf("%v\n", bornYear)
		
		if bornYear == currentYear {
			break
		}
		
		bornYear++
	}
}