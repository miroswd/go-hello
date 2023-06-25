package main

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)

func main() {
		channel1 := make(chan int)
		channel2 := make(chan int)


		go send(10, channel1)
		go other(channel1, channel2)

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


func other(channel1, channel2 chan int) {
	var wg sync.WaitGroup
	for v := range channel1 {
		wg.Add(1)
		go func(x int) {
			channel2 <- work(x)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(channel2)
}


func work(n int) int {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
	return n 
}