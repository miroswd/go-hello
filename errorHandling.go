package main


import (
	"fmt"
	"os"
	"log"
)


func main(){
	_, err := os.Open("README.md1")

	if err != nil {
		fmt.Println("error:", err)
		
		log.Println("Log error:", err) // timestamp
		// log.Fatalln(err)
		log.Panicln(err)

	}


	f, errToCreateLogFile := os.Create("log.txt")
		
	if errToCreateLogFile != nil  {
		log.Println("Create file err:", errToCreateLogFile) // timestamp
	}

	defer f.Close()

	log.SetOutput(f)
} 