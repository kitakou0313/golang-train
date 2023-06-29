package main

import "fmt"

/*
Memento pattern
- objectのsnapsthopの保存，復元を可能にするパターン
- redo, undoオペレーションの実装に便利
*/
func main() {

	caretaker := &Caretaker{
		mementoArray: make([]*Memento, 0),
	}

	originator := &Originator{
		state: "A",
	}

	caretaker.addMemento(originator.createMemento())

	originator.setState("B")
	caretaker.addMemento(originator.createMemento())

	originator.setState("C")
	caretaker.addMemento(originator.createMemento())

	originator.restoreMement(caretaker.getMemento(1))
	fmt.Println(originator.getState())

	originator.restoreMement(caretaker.getMemento(0))
	fmt.Println(originator.getState())
}
