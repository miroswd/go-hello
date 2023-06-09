package main


import ("fmt")


func main() {
	x := "Bom dia! \n \ttab \"entre aspas\""
	fmt.Println(x)
	fmt.Print("Line without break row")

	y := "New line";


	// concatStrings := fmt.Sprint(x, " => ", y)

	fmt.Println(fmt.Sprint(x, " => ", y))
}