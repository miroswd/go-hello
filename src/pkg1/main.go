package main

import "fmt"
import "github.com/miroswd/go-hello/src/pkg2"

// go run src/pkg1/*.go

func main() {
	fmt.Println("Main")
	pkg2.Pk2()
	Pk1()
}