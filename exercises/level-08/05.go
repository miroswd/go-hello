package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}


type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type orderUsersByAge []user
type orderUsersByLastName []user

// orderUsersByAge
func (u orderUsersByAge) Len() int {
	return len(u)
}

func (u orderUsersByAge) Less(i, j int) bool {
	return u[i].Age < u[j].Age
}

func (u orderUsersByAge) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

// orderUsersByLastName
func (u orderUsersByLastName) Len() int {
	return len(u)
}

func (u orderUsersByLastName) Less(i, j int) bool {
	return u[i].Last < u[j].Last
}

func (u orderUsersByLastName) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
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

	for _, v := range users {
		sort.Strings(v.Sayings)
	}

	sort.Sort(orderUsersByAge(users))

	sort.Sort(orderUsersByLastName(users))

	for _, v := range users {
		fmt.Println()

		fmt.Printf("%v %v (%v)\n\n", v.First, v.Last, v.Age)
		fmt.Println("Sayings:")
		
		for _, s := range v.Sayings {
			fmt.Printf("\t - %v\n", s)
		}
	}




}
