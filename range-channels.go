package main

import "fmt"


func main() {
	c := make(chan int)

	go myLoop(10, c)
	prints(c)
}


func myLoop(t int, s chan<- int) {
	for i := 0; i < t; i++ {
		s <- i;
	}

	close(s) // finish channel listening
}

func prints(r <-chan int) {
	for v := range r {
		fmt.Println("Received:", v)
	}
}