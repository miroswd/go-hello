package main


import (
	"fmt"
)


func main() {
	channel := make(chan int)

	num(10, channel)
	
	for x := 0; x < 100; x++ {
		fmt.Println(x, <- channel)
	}
}


func num(n int, channel chan <- int) {
	for i:= 0; i < n; i++ {
		go func(){
			for j:= 0; j < n; j++ {
				channel <- j
			}
			}()
	}

}

