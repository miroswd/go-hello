package main


import (
	"fmt"
)


func main() {
	f()	
	
}


func f() {
	defer func(){
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	fmt.Println("calling g")
	g(0)
}

func g(i int){
	if i > 3 {
		fmt.Println("Panicking")
		panic(fmt.Sprintf("%v", i))
	}

	
	defer fmt.Println("Defer g")
	fmt.Println("Printing in g", i)

	g(i+1)
}