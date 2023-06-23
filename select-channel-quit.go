package main

import (
	"fmt"
)

func main(){
	channel := make(chan int)
	quit := make(chan int)

	go receiveQuit(channel, quit)
	sendToChannel(channel, quit)
}


func receiveQuit(channel, quit chan int) {
	for i := 0; i < 50; i++ {
		fmt.Println("Received:", <-channel)

		if (i == 30) {
			quit <- 0
		}
	}

}

func sendToChannel(channel, quit chan int) {
	anything := 0

	for {
		select {
		case channel <- anything:
				anything++
		case <- quit:
			return
		}
	}
}