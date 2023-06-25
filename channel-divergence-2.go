package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
		channel1 := make(chan int)
		channel2 := make(chan int)

		functions := 5 // every second it executes 5 goroutines, until completing the "send function" loop item

		go send(20, channel1)
		go other(functions, channel1, channel2)

		for v := range channel2 {
			fmt.Println(v)
		}
}


func send(n int, channel chan int) {
	for i := 0; i < n; i++ {
		channel <- i
	}

	close(channel)
}


func other(functions int, channel1, channel2 chan int) {
	var wg sync.WaitGroup


	for i := 0; i < functions; i++ {
		wg.Add(1)
	
		go func(){
			for v := range channel1{
				channel2 <- work(v)
			}
			wg.Done()
		}()

	}
	wg.Wait()
	close(channel2)
}


func work(n int) int {
	time.Sleep(time.Millisecond * 1000)
	return n 
}