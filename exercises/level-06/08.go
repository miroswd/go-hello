package main

import ("fmt")


func returnAFunc() func(x string) {
	return func(name string) {
		fmt.Println("Hey Hey Hey", name)
	}
}


func main(){

	x := returnAFunc();
	x("Miro")

}