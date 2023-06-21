package main

import (
	"fmt"
	"sort"
)

func main(){
		ss := []string{"is","boy","of","dad","is"}
		ii := []int{4,1,2,6,5}
		
		
		sort.Strings(ss)
		sort.Ints(ii)

		fmt.Print(ss)
		fmt.Print(ii)
}