package main

import (
	"fmt"
	"sync"
)


func main() {

	even := make(chan int)
	odd  := make(chan int)
	convergence := make(chan int)

	go send(even, odd)
	go receive(even, odd, convergence)

	for v := range convergence {
		fmt.Println("Received value:", v)
	}

}



func send(e, o chan int) {
	x := 100

	for n := 0; n < x; n++{
		if n % 2 == 0 {
			e <- n
		} else {
			o <- n
		}		
	}

	close(e)
	close(o)
}


func receive(e, o, c chan int) {
	var wg sync.WaitGroup
	
	wg.Add(2)
	go func(){
			for v := range e {
				c <- v
			}
			wg.Done()
		}();

		go func(){
			for v := range o {
				c <- v
			}
			wg.Done()
		}()

		wg.Wait()
		close(c)
}