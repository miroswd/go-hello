package main

import ("fmt")

func main() {
	
	x := returnFunction()

	y := 3;

	fmt.Println(x(y))

	// execute straight
	fmt.Println(returnFunction()(5))
}

// returns a function that return int
func returnFunction() func(int) int {
	return func(a int) int {
		return a * 10;
	}
}