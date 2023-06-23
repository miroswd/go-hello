package main


import (
	"fmt"
	"sync"
	"time"
)


func p1() {
	for x := 0; x < 100; x++ {
		
		fmt.Println("p1", x)
		time.Sleep(200)

}

	wg.Done()
}

func p2() {

	for x := 0; x < 100; x++ {
		
			fmt.Println("p2", x)

			time.Sleep(200)
	}

	wg.Done()
	
}


var wg sync.WaitGroup

func main() {

	wg.Add(2)

	go p1();
	go p2()

	wg.Wait()

}