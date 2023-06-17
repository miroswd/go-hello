package main


import "fmt"


type person struct {
	name string
	age int
}

/*
func (receiver) identifier(parameters) (returns) {
	// code
}
*/

// receiver -> can only receive these types and the function is attached to this type
// parameter -> receive the values

func (p person) goodMorning() {
	fmt.Println(p.name, "says good morning")
}


func (p person) getAge() {
	fmt.Println(p.age)
}

func (p person) generateUserName() string {
	return fmt.Sprintf("%s%d", p.name, p.age)
}


func main(){
	mario := person{"MarioBros", 25}

	mario.goodMorning()	
	mario.getAge()

	username := mario.generateUserName()

	fmt.Println(username)
}


