package main


import "fmt"


func main(){

	users := map[string][]string{
		"leto_miro": []string{
			"play guitar", "fly", "run",
		},
		"doe_john": []string{
			"read", "eat", "sleep",
		},
	}


	for k, v := range users {
		fmt.Println(k)

		for i, e := range v {
			fmt.Println("	", i, e)
		}
	}
}