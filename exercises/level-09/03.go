package main

import (
	"fmt"
	"sync"
	"runtime"
)


var count = 0;
var wg sync.WaitGroup;


func increment(numberOfGoRoutines int) {

	
	wg.Add(numberOfGoRoutines)
	
	for gr := 0; gr < numberOfGoRoutines; gr++ {
		go func(){

		currentPosition := count
		fmt.Println("count:", currentPosition, "gr:", gr)

		runtime.Gosched()

		currentPosition++

		count = currentPosition;

		wg.Done()
		}()

	}

}

func main() {
	 increment(50)
	wg.Wait()
}