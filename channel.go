package main


import (
	"fmt"
)


func main(){
	canal := make(chan int) // create a channel 

	go func(){
		canal <- 42
	}()

	fmt.Println(<- canal)
}
