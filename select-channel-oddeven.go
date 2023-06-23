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
	total := 50

	for i := 0; i < total; i++ {
		if i % 2 == 0 {
			evenChannel <- i
		} else {
			oddChannel <- i
		}
	}

	quit <- true

	close(oddChannel)
	close(evenChannel)

}



func receive (oddChannel, evenChannel chan int, quit chan bool) {
	for {
		select {
		case v := <- oddChannel:
			fmt.Println("The number:", v, "is odd")
		case v := <- evenChannel:
			fmt.Println("The number:", v, "is even")
		case <- quit:
			close(quit)
			return
		}
	}

}