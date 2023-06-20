package main;


import ("fmt")

func main(){
	slice := []int{1,2,3,6}

	fmt.Println(sum(slice...))
	fmt.Println(sliceSum(slice))
}


func sum(x...int) int {

	total := 0
	for _, n := range x {
		total += n
	}

	return total;
}

func sliceSum(slice []int) int {
	total := 0

	for _, n := range slice {
		total += n
	}

	return total;
}