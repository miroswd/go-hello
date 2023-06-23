package main


import ("fmt")


func main() {
	channel := make(chan int)

	go func(){
		channel <- 0
		close(channel)
	}()

	v, ok := <- channel

	fmt.Println(v, ok) // 0, true

	v, ok = <- channel

	fmt.Println(v, ok) // 0, false
}