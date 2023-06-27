package main

import (
	"fmt"
	"log"
)


type especialError struct {
	message string
}


func (e especialError) Error() string {
	return "Failed!"
}

func f(e error) {
	fmt.Println("Error:", e.(especialError).message)
	log.Fatalln(e)
}

func main() {
	er := especialError{
		"Failed",
	}

 f(er)
}


