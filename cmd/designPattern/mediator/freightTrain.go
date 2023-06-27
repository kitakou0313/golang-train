package main

import "fmt"

type FreightTrain struct {
	mediator Mediator
}

func (g *FreightTrain) arrive() {
	if !g.mediator.canArrive(g) {
		fmt.Println("Passanger Tarin: Arrival blocked, wating")
		return
	}

	fmt.Println("Passengers Train: Arrived")

}

func (g *FreightTrain) depart() {
	fmt.Println("PassengerTrain: Leaving")
	g.mediator.notifyAboutDeparture()
}

func (g *FreightTrain) permitArrival() {
	fmt.Println("PassengersTrain: Arraival permitted, arriving")
	g.arrive()
}
