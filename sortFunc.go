package main


import (
	"fmt"
	"sort"
)

type car struct {
	name string
	power int
	consume int
}

type Interface interface{
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type orderByPower 								[]car
type orderByConsume 							[]car
type orderByGasStationProfitOwner []car

// by power

func (x orderByPower) Len() int {
	return len(x)
}

func (x orderByPower) Less(i, j int) bool {
	return x[i].power < x[j].power;
}

func (x orderByPower) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

// by consume

func (x orderByConsume) Len() int {
	return len(x)
}

func (x orderByConsume) Less(i, j int) bool {
	return x[i].consume > x[j].consume;
}

func (x orderByConsume) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

// by profit 

func (x orderByGasStationProfitOwner) Len() int {
	return len(x)
}

func (x orderByGasStationProfitOwner) Less(i, j int) bool {
	return x[i].consume < x[j].consume;
}

func (x orderByGasStationProfitOwner) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func main(){

	cars := []car{
		car{
			name: "civic",
			power: 160,
			consume: 16,
		},
		car{
			name: "lancer",
			power: 140,
			consume: 8,
		},
		car{
			name: "fit",
			power: 116,
			consume: 20,
		},
	}


	fmt.Println(cars)

	sort.Sort(orderByPower(cars))

	fmt.Println("order by power:", cars)

	sort.Sort(orderByConsume(cars))

	fmt.Println("order by consume:", cars)

	sort.Sort(orderByGasStationProfitOwner(cars))

	fmt.Println("order by profit:", cars)


}