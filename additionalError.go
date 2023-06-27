package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrNorgateMath = errors.New("norgate math: square root of negative number")

type norgateMathErr struct {
	lat string
	long string
	err error
}


func main(){
	fmt.Printf("%T\n", ErrNorgateMath)

	_, err := sqrt(-10)

	if err != nil {
		log.Fatalln(err)
	}
}

func (n norgateMathErr) Error() string{
	return fmt.Sprintf("a norgate math error occured: %v %v %v", n.lat, n.long, n.err)
}


func sqrt(f float64) (float64, error) {
	if f < 0 {
		// return 0, ErrNorgateMath
		// return 0, fmt.Errorf("Square root of negative number: %v", f)

		norgateErr := fmt.Errorf("norgate math redux: square root of negative number: %v", f)

		return 0, norgateMathErr{
			"50000",
			"99999",
			norgateErr,
		}
	} 

	return 42, nil
}