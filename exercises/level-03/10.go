package main

import "fmt"

func main() {
	a := true  // true && true
	b := false // true && false
	c := true  // true || true
	d := true  // true || false
	e := false // !(true)

	fmt.Println(true && true == a)
	fmt.Println(true && false == b)
	fmt.Println(true || true == c)
	fmt.Println(true || false == d)
	fmt.Println(!true == e)	
}