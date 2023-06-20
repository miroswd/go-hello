package main

import "fmt"


type person struct {
	name string
	age int
}


func changeMe(p *person) {
	(*p).name = "Miro"
	(*p).age = 23
	
}

func main() {

	user := person{
		"John",
		20,
	}

	fmt.Println(user)

	changeMe(&user)

	fmt.Println(user)
}