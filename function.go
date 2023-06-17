package main



import "fmt"


func main() {
	withArgument("after")
	withArgument("morning")
	value := sum(10, 23)

	fmt.Println(value)

	total, length := multipleArgs(1,2,4)
	fmt.Println(total, length)


}

/*
func (receiver) name(parameters) (return1, return2){
	

	func t(parameter){

	}
->
	t(argument)


}
*/


// func (receiver) name (thingShouldCount) int {

// }

func basic() {
	fmt.Println("Hi! good morning")
}

func withArgument(s string) {
	if s == "morning" {
		 basic()
	} else {
		fmt.Println("Hello")
	}
}

				// x int, y int -> as it is the same type can be x, y int
func sum(x, y int) int{
	return x + y
}


func multipleArgs(x ...int) (int, int) {
	sum := 0

	for _, v := range x {
		sum += v;
	}

	return sum, len(x)
}