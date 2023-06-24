package main

import "fmt"

type Medical struct {
	next Department
}

func (m *Medical) handle(p *Patient) {
	if p.medichineDone {
		fmt.Println("Medical step is already done")
		m.next.handle(p)
	}

	fmt.Println("Medical is giving medicine to patient")
	p.medichineDone = true
	m.next.handle(p)
}

func (m *Medical) setNext(next Department) {
	m.next = next
}
