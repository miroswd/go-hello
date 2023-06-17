package main

import "fmt"

func main(){
	fmt.Println(fatorial(6))
}

func fatorial(x int) int {
	// !5  -> 5 * 4 * 3 * 2 * 1
	total := x

	if x > 1 {
		total *= fatorial(x-1)
	}

	return total
}
