package main


import ("fmt")


func main() {
	fmt.Println(returnInt())
	fmt.Println(returnMultipleTypes())
}


func returnInt() int {
	return 1
}

func returnMultipleTypes() (int, string){
	return 2, "two"
}