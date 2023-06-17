package main


import ("fmt")


func main(){

	x := 10

	y := func(a int){
		fmt.Println(a * 2)
	}

	z := func(a int) int {
		return a + 2;
	}

	y(x)
	fmt.Println(z(x))
	
}