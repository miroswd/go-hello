package main

import ("fmt")

func main() {
	x := 10

	func (a int) {
		fmt.Println(a * 2)		
	}(x) // disposable function

}
