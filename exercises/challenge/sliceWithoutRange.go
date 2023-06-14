package main

import "fmt"


func main(){
	// access all items from slice without use the range

	sliceName := []string{"John", "Doe", "Da", "Silva"}

	// fullName := sliceName[0:len(sliceName)]
	fullName := sliceName[:]

	fmt.Println(fullName)
}