package main


import "fmt"


func main() {
	slice := []string{"Hey", "Hi"}

	for i, v := range slice {
		fmt.Println("Index:", i, "Value:", v)	
	}

	fmt.Println()

	slice = append(slice, "Hello")

	for i, v := range slice {
		fmt.Println("Index:", i, "Value:", v)	
	}

	fmt.Println()

	for i := range [3]int{1,2,3} {
		fmt.Printf("%v\n", slice[i])
	}
}


