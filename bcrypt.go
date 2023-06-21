package main


import "fmt"
import "golang.org/x/crypto/bcrypt"


func main(){
	
	pass := "june03"
	wrongPass := "jone03"

	
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(pass), 8) 

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(hashedPasswordBytes))


	fmt.Println("success:", bcrypt.CompareHashAndPassword(hashedPasswordBytes, []byte(pass)))
	fmt.Println("error:", bcrypt.CompareHashAndPassword(hashedPasswordBytes, []byte(wrongPass)))

	
}