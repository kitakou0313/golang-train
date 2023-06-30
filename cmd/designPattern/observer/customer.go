package main

import "fmt"

type Customer struct {
	id string
}

func (c *Customer) update(itemName string) {
	fmt.Printf(
		"Sending email to %s for item : %s\n", c.id, itemName,
	)
}

func (c *Customer) getID() string {
	return c.id
}
