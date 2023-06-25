package main

import (
	"fmt"
)


func main() {
	channel := make(chan int)

	go num(100, channel)
	print(channel)
}


func num(qt int, channel chan<- int) {
	for i := 0; i < qt; i++ {
		channel <- i
	}
	close(channel)
}

func print(c <- chan int) {
	for v := range c {
		fmt.Println(v)
	}
}