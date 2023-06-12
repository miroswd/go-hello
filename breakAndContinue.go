package main

import (
	"fmt"
)


func main() {
	for i := 0; i < 20; i++ {

		if i % 2 != 0 {
			continue // skip i
		}

		fmt.Println(i)

		if (i == 10) {
			break // break the loop
		}

	}
}