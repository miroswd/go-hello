package main

import (
	"fmt"
	"math/rand"
	"time"
)


func main(){
	c := converges(work("x"), work("y"))

	for x := 0; x < 10; x++ {
		fmt.Println(<- c)
	}
}


func work(s string) chan string{
	channel := make(chan string)

	go func(s string, c chan string){
		for i := 0; ; i++ {
			c <- fmt.Sprintf("Function %v says: %v", s, i)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
		}
	}(s, channel)

	return channel
}


func converges(x, y chan string) chan string {
	newChannel := make(chan string)

	go func(){
		for {
			newChannel <- <- x
		}
	}()

	go func(){
		for {
			newChannel <- <- y
		}
	}()
	
	return newChannel
}