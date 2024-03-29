package main


import ("fmt")

// create a func that receive stuff from different channels

func main(){

	a := make(chan int)
	b := make(chan int)

	x := 10


	go func(x int){
		for i := 0; i < x; i++ {
			a <- i
		}
	}(x/2)

	go func(x int){
		for i := 0; i < x; i++ {
			b <- i
		}
	}(x/2)


	for i := 0; i < x; i++ {
		select {
		case  v:= <- a:
			fmt.Println("Channel A:", v)
		case  v:= <- b:
			fmt.Println("Channel B:", v)
			
		}
	}

}
