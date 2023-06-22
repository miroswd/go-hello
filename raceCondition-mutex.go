package main

import "fmt"
import "sync"
import "runtime"

func main(){

	count := 0
	totalOfGoRoutines := 10

	var wg sync.WaitGroup
	var mutex sync.Mutex

	wg.Add(totalOfGoRoutines)

	for i := 0; i < totalOfGoRoutines; i++ {
		go func() {
			mutex.Lock()
			v := count

			// yield -> length of time the process switch occurs on the processor
			runtime.Gosched()

			v++
			count = v
			mutex.Unlock()
			wg.Done()
		}()
	}

	wg.Wait() // wait all go routines inside the loop
	
	fmt.Println(count)
}