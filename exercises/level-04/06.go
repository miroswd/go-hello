package main


import "fmt"

func main() {
	statesSlice := []string{"Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins"}
	
	states := make([]string, len(statesSlice), 27);


  for i, e := range statesSlice {
		states[i] = e
	}


	fmt.Println(len(states), cap(states))
	
	for i := 0; i < len(states); i++ {
		fmt.Println(states[i])
	}
	

}