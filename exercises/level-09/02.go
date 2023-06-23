package main

import "fmt"

type pessoa struct {
	name string
	age int
}

type humano interface {
	falar()
}

func (x *pessoa) falar() {
	fmt.Print("Hi")
}

func dizerAlgumaCoisa(h humano) {
  h.falar()
}

func main() {
	p := pessoa{
		"Miro",
		23,
	}

	dizerAlgumaCoisa(&p)
}