package main


import "fmt"

// array de array

func main() {
	ss := [][]int{
					[]int{1, 2, 3},
					[]int{4, 5, 6},
	}


	sss := [][][]int{
		[][]int{
			[]int{1,2,3},
		},
	}

	fmt.Println(ss) // [[1 2 3] [4 5 6]]

	fmt.Println(ss[1][0])


	fmt.Println()
	fmt.Println(sss)

}