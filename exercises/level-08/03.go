package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}


func (u user) printName() user {
	fmt.Printf("%v %v\n", u.First, u.Last)
	return u
}

func (u user) printAge() user {
	fmt.Printf("%v\n", u.Age)
	return u
}

func (u user) printFirstSayings() {
	fmt.Printf("%v\n", u.Sayings[0])
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)
	
	encode := json.NewEncoder(os.Stdout)
	
	encode.Encode(users)
	

	for _, v := range users {

		fmt.Println("\n----------------\n")
		v.printAge().printName().printFirstSayings()

	}

}
