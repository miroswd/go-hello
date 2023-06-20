package main


import "fmt" 


func closure() func() {
	x := 10;

	return func() {
		fmt.Println(x*2)
	}
}


func main() {
	closure()()
}