package main

import "fmt"

type Reception struct {
	next Department
}

func (r *Reception) handle(p *Patient) {
	if p.registarationDone {
		fmt.Println("Patient registration already Done")
		r.next.handle(p)
		return
	}

	fmt.Println("Reception registering patient")
	p.registarationDone = true
	r.next.handle(p)
}

func (r *Reception) setNext(next Department) {
	r.next = next
}
