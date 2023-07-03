package main

import "fmt"

type ItemRequestedState struct {
	vendingMachine *VendingMachine
}

func (i *ItemRequestedState) requestItem() error {
	return fmt.Errorf("Item already requested")
}
func (i *ItemRequestedState) addItem(count int) error {
	return fmt.Errorf("Item dispense in progress")
}
func (i *ItemRequestedState) insertMoney(monry int) error {
	if monry < i.vendingMachine.itemPrice {
		return fmt.Errorf("Inserted money is less. Please insert %d", i.vendingMachine.itemPrice)
	}
	fmt.Println("Monry enterd is ok")
	i.vendingMachine.setState(i.vendingMachine.hasItem)
	return nil
}
func (i *ItemRequestedState) dispenseItem() error {
	return fmt.Errorf("Please insert money first")
}
