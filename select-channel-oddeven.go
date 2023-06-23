package main


import (
	"fmt"
)


func main() {
	quit        := make(chan bool)
	oddChannel  := make(chan int)
	evenChannel := make(chan int)

	go sendNumbers(oddChannel, evenChannel, quit)
	receive(oddChannel, evenChannel, quit)
}


func sendNumbers(oddChannel, evenChannel chan int, quit chan bool) {
	total :=10

	for i := 0; i < total; i++ {
		if i % 2 == 0 {
			evenChannel <- i
		} else {
			oddChannel <- i
		}
	}

	quit <- true
	
	close(evenChannel)
	close(oddChannel)
	
	/*
	The number: 46 is even
	The number: 47 is odd
	The number: 48 is even
	The number: 49 is odd
	The number: 0 is even 
	The number: 0 is odd

	the zeros represent the closed channel
	*/


}



func receive (oddChannel, evenChannel chan int, quit chan bool) {
	for {
		select {
		case v := <- oddChannel:
			fmt.Println("The number:", v, "is odd")
		case v := <- evenChannel:
			fmt.Println("The number:", v, "is even")
		case v, ok := <- quit:
			if !ok {
				fmt.Println("Failed:", v)
			} else {
				fmt.Println("Finishing...", v)
			}
			return
		}
	}

}