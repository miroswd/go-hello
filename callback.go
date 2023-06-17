package main


import "fmt"


func main() {
	resultOfSum, even := evenNumbers(sum, []int{1, 2, 3, 4, 5, 6}...)
	resultOfMultiplication, odd := oddNumbers(multiplicationFunc, []int{1,2,3,4,5,6,7}...)
	fmt.Print(resultOfSum, even)
	fmt.Print("\n ----------- \n")
	fmt.Print(resultOfMultiplication, odd)
}

func sum(x ...int) int {
	total := 0

	for _, n := range x {
		total += n;
	}

	return total;
}

func evenNumbers(sumFunc func(x ...int) int, y ...int) (int, []int) {

	var slice []int;


	for _, v := range y {
		if v % 2 == 0 {
			slice = append(slice, v)
		}
	}


	total := sumFunc(slice...);

	return total, slice	
}


func multiplicationFunc(x ...int) int {
	total := 1;
	
	for _, v := range x {
		total *= v
	}

	return total;
}


func oddNumbers(multiplicationFunc func(x ...int) int, y ...int ) (int, []int) {
	
	slice := []int{};

	for _, n := range y {
		if n % 2 != 0 {
			slice = append(slice, n)
		}
	}

	return multiplicationFunc(slice...), slice
}