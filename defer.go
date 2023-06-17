package main

import ("fmt")


func main() {
	defer fmt.Println("With defer 1")
	fmt.Println("Without defer")
	fmt.Println("Without defer")
	fmt.Println("Without defer")
	fmt.Println("Without defer")
	fmt.Println("Without defer")
	defer fmt.Println("With defer 2")
	fmt.Println("Without defer")
	fmt.Println("Without defer")
	fmt.Println("Without defer")
	fmt.Println("Without defer")
	fmt.Println("Without defer")
	fmt.Println("Without defer")
	fmt.Println("Without defer")
	fmt.Println("Without defer")
	fmt.Println("Without defer")
	fmt.Println("Without defer")
}