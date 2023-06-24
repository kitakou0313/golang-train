package main

import "fmt"

type Doctor struct {
	next Department
}

func (d *Doctor) handle(p *Patient) {
	if p.dockerCheckUpDone {
		fmt.Println("Docker step is already done.")
		d.next.handle(p)
	}

	fmt.Println("Doctor checking patient")
	p.dockerCheckUpDone = true
	d.next.handle(p)
}

func (d *Doctor) setNext(next Department) {
	d.next = next
}
