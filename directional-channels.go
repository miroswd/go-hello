package main

import "fmt"


func main(){
	canal := make(chan int) // send or receive

	go send(canal) 
	receive(canal)
}


func send(s chan<- int){ // put values in the channel
	s <- 42
}

func receive(r <-chan int){ // get channel values
	fmt.Println("The channel value received:", <- r)
}