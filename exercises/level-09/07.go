package main

import "fmt"

// interface -> descrever os comportamentos de um tipo

type dog struct {
	name string
}

type cat struct {
	name string
}

type animal interface {
	speak()
}

type bicho interface {
	speaking() string
}


func (d dog) speak() {
	fmt.Println(d.name, "mandou um salve")
}


func (c cat) speak() {
	fmt.Println(c.name, "miau")
}

func (d dog) speaking() string {
	return "AUUUUUUUUUUUUUUUUUUUU"
}

func (c cat) speaking() string {
	return "salve meu patrão"
}


func speakAnything(b bicho) {
	fmt.Println(b.speaking())
}

func main(){
	d := dog{
		"Adilsu",
	}

	c := cat{
		"januário",
	}
	
	speakAnything(c)
	speakAnything(d)
}
