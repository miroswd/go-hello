package main


import (
	"encoding/json"
	"fmt"
)

type info struct {
	Name 				string 	`json: "Name"`
	LastName		string 	`json: "LastName"`
	Age 				int    	`json: "Age"`
	Job 				string 	`json: "Profession"`
	BankAccount float64 `json: "Balance"`
}


func main(){
	sb := []byte(`{"Name":"James","LastName":"Bond","Age":40,"Job":"Spy","BankAccount":1000000.5}`)


	var j info
	err := json.Unmarshal(sb, &j)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(j)
	fmt.Println(j.Name)

}