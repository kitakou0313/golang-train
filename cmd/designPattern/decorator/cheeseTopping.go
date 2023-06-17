package main

type CheezeTopping struct {
	pizza IPizza
}

func (c *CheezeTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()

	return pizzaPrice * 10
}
