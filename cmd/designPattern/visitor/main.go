package main

import "fmt"

func main() {

	square := &Square{
		side: 2,
	}
	circle := &Circle{
		radius: 3,
	}
	rectangle := &Rectangle{
		l: 2, b: 3,
	}

	areaCalculator := &AreaCalculator{}

	square.accept(areaCalculator)
	circle.accept(areaCalculator)
	rectangle.accept(areaCalculator)

	fmt.Println()
	middleCordinates := &MiddleCoordinates{}

	square.accept(middleCordinates)
	circle.accept(middleCordinates)
	rectangle.accept(middleCordinates)
}
