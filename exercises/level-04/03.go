package main 


import "fmt"

func main() {
	slice := []int{}

	for i := 1; i <= 10; i++ {
		slice = append(slice, i)
	}

	fmt.Println(slice[:3])
	fmt.Println(slice[4:])
	fmt.Println(slice[1:7])
	fmt.Println(slice[2:len(slice) - 1])

}