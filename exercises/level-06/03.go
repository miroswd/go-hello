package main

import (
	"fmt"
)

func main(){
	defer fmt.Println("Defer is called")
	fmt.Println("Hi")
}
