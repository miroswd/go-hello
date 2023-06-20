package main

import "fmt"


func closeFile() {
	defer fmt.Println("Closed file")
	
	for i := 0; i < 10; i++ {
		fmt.Printf("processing %d%%...\n", (i+1)*10)
	}
}

func readFile(cf func()) {
	defer cf()

	fmt.Println("Read line")
	fmt.Println("Read line")
	fmt.Println("Read line")
	fmt.Println("Read line")
}


func main(){
	readFile(closeFile)
}