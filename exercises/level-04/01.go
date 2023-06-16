package main

import "fmt"

func main() {
	arr := [5]int{1,2,3,4,5}

	for _, e := range arr {
		fmt.Println(e)
	}

	fmt.Printf("%T", arr)
}
