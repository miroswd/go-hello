package main

import "fmt"


func main() {
	bornYear := 2000
	currentYear := 2023

	differenceBetweenYears := currentYear - bornYear;

	for differenceBetweenYears >= 0 {
		fmt.Println(currentYear - differenceBetweenYears);
		differenceBetweenYears--
	}
}
