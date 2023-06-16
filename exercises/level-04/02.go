package main 


import "fmt"

func main() {
	slice := []int{}

	for i := 1; i <= 10; i++ {
		slice = append(slice, i)
	}

	for _, e := range slice {
		fmt.Println(e)
	}


	fmt.Printf("%T", slice)
}