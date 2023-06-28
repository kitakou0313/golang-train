package main

// mediactor（中間者） pattern
// Component間のやり取りをmediatorが仲介し，直接やりとりしない

// 例ではTrain, Stationが登場するが，互いにはやり取りせずMediator interfaceを経由する

func main() {
	stationManeger := newStationManeger()

	passengerTrain := &PassengerTrain{
		mediator: stationManeger,
	}
	freightTrain := &FreightTrain{
		mediator: stationManeger,
	}

	passengerTrain.arrive()
	freightTrain.arrive()
	passengerTrain.depart()
}
