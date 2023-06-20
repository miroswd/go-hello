package main


import ("fmt")


type person struct {
	name string
	lastName string
	age int
}

func (p person) showData() {
	fmt.Printf("%v %v, %v\n", p.name, p.lastName, p.age)
} 

func main(){
	mario := person{
		name: "Mario",
		lastName: "Bros",
		age: 25,
	}

	mario.showData()
}