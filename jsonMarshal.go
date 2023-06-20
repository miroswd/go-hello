package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type color struct {
	Id int
	Name string
} // first letter needs to be capitalized for marshal to interpret


func main() {

	red := color{
		1,
		"red",
	}

	j, err := json.Marshal(red)

	if err != nil {
		fmt.Println("[error]", err)
	}

	os.Stdout.Write(j)

}