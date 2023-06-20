package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string
	LastName string
	Age int
	Job string
	BankAccount float64
}

func main(){
	jamesBond := person{
		Name: "James",
		LastName: "Bond",
		Age: 40,
		Job: "Spy",
		BankAccount: 1000000.50,
	}

	darthVader := person{
		"Anakin",
		"Skywalker",
		50,
		"Vilan",
		50000000.11,
	}

	j, jErr := json.Marshal(jamesBond)

	if jErr != nil {
		fmt.Println(jErr)
	}

	d, dErr := json.Marshal(darthVader)

	if dErr != nil {
		fmt.Println(dErr)
	}

	fmt.Println(string(j))
	fmt.Println(string(d))
}