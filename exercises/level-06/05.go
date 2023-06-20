package main

import ("fmt")

type square struct {
	width, height float64
}

type circle struct {
	radius float64
}

type figure interface {
	area() float64
}


func (s square) area() float64 {
	return s.height * s.width;
}

func (c circle) area() float64 {
	return 2 * 3.14 * c.radius
}


func info(f figure) {
	fmt.Println(f.area());
}


func main() {
	squareValue := square{
		width: 15,
		height: 15,
	}

	circleValue := circle{
		radius: 30,
	}

	info(squareValue)
	info(circleValue)
}