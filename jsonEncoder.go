package main


import (
	"encoding/json"
	"os"
)

type person struct {
	Name 				string 	
	LastName		string 	
	Age 				int    	
	Job 				string 	
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


	encoder := json.NewEncoder(os.Stdout)
	
	encoder.Encode(jamesBond)

}