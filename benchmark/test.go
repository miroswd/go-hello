package main


import (
	"fmt"
)

func main() {
	x := sum(1, 2, 3)
	y := multiply(1, 2, 3)
	fmt.Println(x, y)
}


func sum(x ...int) int{

	total := 0
	for _, v := range x {
		total += v
	}

	return total
}

func multiply(x ...int) int{

	total := 0
	for _, v := range x {
		total *= v
	}

	return total
}