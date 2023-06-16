package main


import "fmt"


func main(){
	x := struct{
		example map[string]string
		slice []int
	}{
		map[string]string{
			"foo": "bar",
		},
		[]int{1,2,3,},
	}

	fmt.Println(x)
}