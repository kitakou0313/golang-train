package main

import "fmt"

type NoItemState struct {
	vendingMachine *VendingMachine
}

func (i *NoItemState) addItem(count int) error {
	i.vendingMachine.incrementItemCount(count)
	i.vendingMachine.setState(i.vendingMachine.hasItem)
	return nil
}

func (i *NoItemState) requestItem() error {
	return fmt.Errorf("Item out of stock")
}

func (i *NoItemState) insertMoney(monry int) error {
	return fmt.Errorf("Item out of stock")
}

func (i *NoItemState) dispenseItem() error {
	return fmt.Errorf("Item out of stock")
}
