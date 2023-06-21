// run at the same time and not one by one


package main


import (
	"fmt"
	"sync"
	"time"
	"runtime"
)

var wg sync.WaitGroup;

func main(){
	// wait groups


	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())

	wg.Add(2)
	// add (func total)
	go func1()

	go func2()

	fmt.Println(runtime.NumGoroutine())

	// wait all funcs
	wg.Wait()
}

func func1() {
	for i := 0; i < 100; i++ {
		fmt.Println("func1:", i)
		time.Sleep(20)
	}

	// done, executed func
	wg.Done()
}


func func2() {
	for i := 0; i < 100; i++ {
		fmt.Println("func2:", i)
		time.Sleep(20)

	}

	wg.Done()

}
