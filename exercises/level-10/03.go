package main


import (
	"fmt"
)


func main(){
	c := gen()	

	receive(c)
}


func gen() <- chan int{
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func receive(x <- chan int) {
	for v := range x {
		fmt.Println(v)
	}
}