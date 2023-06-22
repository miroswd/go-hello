package main

import "fmt"
import "sync"
import "sync/atomic"
import "runtime"

func main(){

	var count int64
	count = 0
	totalOfGoRoutines := 10

	var wg sync.WaitGroup

	wg.Add(totalOfGoRoutines)

	for i := 0; i < totalOfGoRoutines; i++ {
		go func() {
			atomic.AddInt64(&count, 1)
			
			// yield -> length of time the process switch occurs on the processor
			runtime.Gosched()
			
			fmt.Println("Count: ", atomic.LoadInt64(&count))
			wg.Done()
		}()
	}

	wg.Wait() // wait all go routines inside the loop
	
	fmt.Println(count)
}