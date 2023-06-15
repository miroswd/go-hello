package main


import "fmt"


func main(){

	anything := map[int]string{
		123: "hi",
		321: "hello",
		111: "hey there!",
		000: "to delete",
	}


	fmt.Printf("%v %T\n", anything, anything)


	total := 0

	for k, v := range anything {
		fmt.Println(k, v)
		total += k
	}

	fmt.Println(total)


	delete(anything, 000)

	fmt.Printf("%v", anything)


}