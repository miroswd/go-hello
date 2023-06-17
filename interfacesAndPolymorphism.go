package main


import ("fmt")


type person struct {
	name     string
	lastName string
	age      int
}


type dentist struct {
	person
	pulledTeeth int
	salary  	  float64
}


type architect struct {
	person
	typeOfBuild string
	salary 			float64

}

func (x dentist) hi(){
	fmt.Println("Hi, I'm", x.name, "good morning! I pulled", x.pulledTeeth, "teeth")
}

func (x architect) hi(){
	fmt.Println("Hi, I'm", x.name, "good morning! I build", x.typeOfBuild )
}

type professional interface {
	hi()	
}

func greetings(p professional) {
	p.hi()

	switch p.(type) {
		case dentist:
			fmt.Println("I earn", p.(dentist).salary)
		case architect:
			fmt.Println(p.(architect).typeOfBuild)
}
}


func main(){
	marioDentist := dentist{
		person: person{
			"Mario",
			"Bros",
			24,
		},
		pulledTeeth: 50,
		salary: 1000.00,
	}

	luigiArchitect := architect{
		person: person{
			"Mario",
			"Bros",
			24,
		},
		typeOfBuild: "house",
		salary: 2000.00,
	}

	marioDentist.hi()
	
	fmt.Println("\n------------ or -------------")

	greetings(luigiArchitect)
	greetings(marioDentist)
}