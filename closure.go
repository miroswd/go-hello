package main


import ("fmt")


func main(){
	// x in the scope from A function
	a := i()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	// 1, 2, 3

	// x in the scope from B function
	b := i()
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	// 1, 2, 3
}


func i() func() int {
	x := 0

	return func() int {
		x++
		return x
	}
}