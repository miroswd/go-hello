package main



import "fmt"


func main() {

	users := [][]string{
						 []string{
							"name",
							"lastname",
							"hobby",
						 },
						 []string{
							"miro",
							"leto",
							"run",
						 },
						 []string{
							"john",
							"doe",
							"play guitar",
						 },
	}


	for _, e := range users {
		fmt.Println(e)
	}


}