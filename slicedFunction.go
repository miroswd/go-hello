package main


import "fmt"



func main() {
	slice := []int{1,2,3,4,5}

	total := sum(slice...)

	fmt.Println(total)
}


func sum(x ...int) int{
	total := 0;

	for _, e := range x {
		total += e;
	}


	return total;
}
