package main


import "fmt"


func main() {
	x := 10
	y := &x

	// get address -> &x
	// get content from address -> *&x

	fmt.Println(*&x, *y)
	fmt.Printf("%T", y)

	receiveValue(x) 
	/* return
		Receive value: 11
		>> 10 // created a copy for the var x, change the copy
	*/

	receivePointer(&x)
	/* return
		Receive value: 11
		>> 11 // changed the reference of the var, change the original
	*/


	fmt.Println(">>", x)

}


func receiveValue(x int){
	x++
	fmt.Println("Receive value:", x)
}

func receivePointer(x *int) {
	*x++
	fmt.Println("Receive pointer:", *x)
}