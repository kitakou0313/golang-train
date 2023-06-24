package main

import "fmt"

type Cashier struct {
	next Department
}

func (c *Cashier) handle(p *Patient) {
	if p.paymentDone {
		fmt.Println("Cashier step is already done")

	}
	fmt.Println("Cashier getting money from patient")
}

func (c *Cashier) setNext(next Department) {
	c.next = next
}
