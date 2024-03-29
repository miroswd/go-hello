package main

import "fmt"

type person struct{
	name string
	age int
}

type professional struct {
	person
	title string
	salary float64
}

func main(){

	person1 := person{
		name: "Mario",
		age: 25,
	}

	person2 := professional{
		person: person{
			name: "Luigi",
			age: 24,
		},
		title: "Plumber",
		salary: 1000.00,
	}


	person3 := professional{
		person: person1,
		title: "Plumber Sr",
		salary: 3000.00,
	}

	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println(person3)

	fmt.Println(person1.name)
	fmt.Println(person2.person.name)

	person4 := person{"Xuxa", 9999}

	fmt.Println(person4)

	person5 := professional{person{"Miro", 23}, "Dev", 1.00}

	fmt.Println(person5)

}
