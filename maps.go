package main


import "fmt"


func main() {

		friends := map[string]int{
			"alfredo": 5551234,
			"joana": 9996674,
		}

		fmt.Println(friends)
		fmt.Println(friends["joana"])


		friends["gopher"] = 44444

		fmt.Println(friends)


		fmt.Println(friends["miro"]) // this return 0, because does not exists miro in the friends map

		friend := "miro"


		if _, ok := friends[friend]; ok {
			fmt.Println(friends[friend], ok)
		} else {
			fmt.Println("The friend does not exists")
		}

}	