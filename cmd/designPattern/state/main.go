package main

import (
	"fmt"
	"log"
)

func main() {
	vendingMachine := newVendingMachine(
		1, 10,
	)
	fmt.Printf("After newVendingMachine:%T\n", vendingMachine.currentState)

	err := vendingMachine.requestItem()
	fmt.Printf("After requestItem:%T\n", vendingMachine.currentState)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.insertMoney(10)
	fmt.Printf("After insertMoney:%T\n", vendingMachine.currentState)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.dispenseItem()
	fmt.Printf("After dispenseItem:%T\n", vendingMachine.currentState)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	err = vendingMachine.addItem(2)
	fmt.Printf("After addItem:%T\n", vendingMachine.currentState)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	err = vendingMachine.requestItem()
	fmt.Printf("After requestItem:%T\n", vendingMachine.currentState)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.insertMoney(10)
	fmt.Printf("After insertMoney:%T\n", vendingMachine.currentState)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.dispenseItem()
	fmt.Printf("After dispenseItem:%T\n", vendingMachine.currentState)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
