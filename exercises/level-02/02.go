package main


import ("fmt")


func main() {
	a := 1
	b := 2
	
	c := (a == b)
	d := (a != b)
	e := (a <= b)
	f := (a < b)
	g := (a >= b)
	h := (a > b)

	fmt.Println(c, d, e, f, g, h)
}