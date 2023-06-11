package main

import ("fmt")

func main() {
	s := "Hello World"

	p := `
		Hello 
						world
				!
	
	`

	sliceOfBytes := []byte(s)

	fmt.Printf("%v\n%T", s, s)
	fmt.Printf("%v\n%T", sliceOfBytes, sliceOfBytes)

	fmt.Printf("\n%v", p);


	for _, v := range sliceOfBytes {
		fmt.Printf("%v - %T - %#U - %#x\n", v,v,v,v)
	}

	fmt.Println()

	for i := 0; i < len(s); i++ {
		fmt.Printf("%v - %T - %#U - %#x\n", s[i], s[i],s[i],s[i])
	}


}