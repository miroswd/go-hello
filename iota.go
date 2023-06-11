package main



import ("fmt")


const (
	_ = iota
	x = iota
	y = iota
	a = iota
	b = iota + a
	c = iota + iota
)

// the iota serves for use a random sequencial value

func main() {
	fmt.Println(x, y, a, b, c)
}