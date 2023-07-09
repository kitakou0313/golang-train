package main

import "fmt"

type MiddleCoordinates struct {
	x int
	y int
}

func (a *MiddleCoordinates) visitForSquare(s *Square) {
	fmt.Println(
		"Calculating middle point corrdinates for square and assign to a.x, a.y",
	)
}

func (a *MiddleCoordinates) visitForCircle(c *Circle) {
	fmt.Println("Calculating middle point corrdinates for Circle and assign to a.x, a.y")
}

func (a *MiddleCoordinates) visitForRectangle(t *Rectangle) {
	fmt.Println(
		"Calculating middle point corrdinates for reutagnle for a.x, a.y",
	)
}
