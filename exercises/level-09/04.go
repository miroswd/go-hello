package main

import (
	"fmt"
	"sync"
	"runtime"
)


var count = 0;
var wg sync.WaitGroup;
var mutex sync.Mutex


func increment(numberOfGoRoutines int) {

	
	wg.Add(numberOfGoRoutines)
	
	for gr := 0; gr < numberOfGoRoutines; gr++ {
		go func(){
			mutex.Lock()
		currentPosition := count
		fmt.Println("count:", currentPosition)

		runtime.Gosched()

		currentPosition++

		count = currentPosition;

		mutex.Unlock()
		wg.Done()
		}()

	}

}

func main() {
	 increment(50)
	wg.Wait()
}