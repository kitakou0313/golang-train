package main

import "fmt"

type PassengerTrain struct {
	mediator Mediator
}

func (g *PassengerTrain) arrive() {
	if !g.mediator.canArrive(g) {
		fmt.Println("Passanger Tarin: Arrival blocked, wating")
		return
	}

	fmt.Println("Passengers Train: Arrived")

}

func (g *PassengerTrain) depart() {
	fmt.Println("PassengerTrain: Leaving")
	g.mediator.notifyAboutDeparture()
}

func (g *PassengerTrain) permitArrival() {
	fmt.Println("PassengersTrain: Arraival permitted, arriving")
	g.arrive()
}
