package main

import "fmt"

func main() {

	pizza := &VeggeMania{}

	pizzaWithCheeze := &CheezeTopping{
		pizza: pizza,
	}

	pizzaWithCheeseAndTomato := &TomatoTopping{
		pizza: pizza,
	}

	fmt.Println(pizzaWithCheeze.getPrice(), pizzaWithCheeseAndTomato.getPrice())
}
