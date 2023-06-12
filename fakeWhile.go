package main


import ("fmt");



func main() {
	
	x := 0
	// for accept only the condition

	for x < 10 {
		
		fmt.Println(x)
		
		x++	
	}

	for {
		if x < 20 {
			fmt.Printf("%v less than 20\n", x)
			x++ 
		} else {
			fmt.Println("BREAK LOOP")
			break
		}
	}

	// infinite loop
	for {
		fmt.Print("Epa")
	}
	
}