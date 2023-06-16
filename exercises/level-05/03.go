package main

import "fmt" 

type vehicle struct {
	doors int
	color string
}

type pickupTruck struct {
	vehicle
	fourWheels bool
}

type sedan struct {
	vehicle
	luxModel bool
}



func main() {

	pickup := pickupTruck{
		vehicle{2, "silver",}, true,
	}

	sedanLux := sedan{
		vehicle{4, "black",}, false,
	}


	fmt.Println(pickup, sedanLux)
	fmt.Println(pickup.color, sedanLux.doors)

}