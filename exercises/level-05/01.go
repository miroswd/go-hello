package main

import "fmt"

type pessoa struct {
	nome string
	sobrenome string
	sorvetesFavoritos []string
}

func main(){

	pessoa1 := pessoa{
		"Miro",
		"Leto",
		[]string{"Chocolate", "Morango",},
	}

	pessoa2 := pessoa{
		"John",
		"Doe",
		[]string{"Limão", "Caqui",},
	}

	pessoas := []pessoa{
		pessoa1,
		pessoa2,
	}


	for _, p := range pessoas{
		fmt.Println(p.nome, p.sobrenome)
		for _, v := range p.sorvetesFavoritos {
			fmt.Println("\t", v)
		}
	}
}