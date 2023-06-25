package main


import (
	"fmt"
	"context"
	"time"
)


func main() {
	// kill the go routines

	ctx, cancel := context.WithCancel(context.Background())


	go func(){
		n := 0;

		for {
			select {
			case <- ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("Working", n)
			}
		}
	}()

	time.Sleep(time.Second* 2)
	fmt.Println("about to cancel context")

	cancel()

	fmt.Print("cancelled context")

}
