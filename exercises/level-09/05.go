package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	// "runtime"
)


var count int64;
var wg sync.WaitGroup;


func increment(numberOfGoRoutines int) {

	
	wg.Add(numberOfGoRoutines)
	
	for gr := 0; gr < numberOfGoRoutines; gr++ {
		go func(){

		atomic.AddInt64(&count, 2)
		currentPosition := count

		// runtime.Gosched()

		currentPosition++

		count = currentPosition;
		fmt.Println("Count: ", atomic.LoadInt64(&count))

		wg.Done()
		}()

	}

}

func main() {
	 increment(50)
	wg.Wait()
}