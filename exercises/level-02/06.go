package main

import ("fmt")

const (
	_ = iota + 2023
	add1Year
	add2Year
	add3Year
	add4Year
)

func main() {
	fmt.Println(
		add1Year,
		add2Year,
		add3Year,
		add4Year,
	)
}
