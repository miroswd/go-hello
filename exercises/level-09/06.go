package main

import (
	"fmt"
	"runtime"
)



func main() {
	fmt.Println("Os:", runtime.GOOS)
	fmt.Println("Arch:", runtime.GOARCH)
}