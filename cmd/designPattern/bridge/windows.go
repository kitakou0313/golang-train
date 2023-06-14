package bridge

import "fmt"

type Windows struct {
	priner Printer
}

func (w *Windows) Print() {
	fmt.Println("Print req for Windows")
}

func (w *Windows) SetPrinter(p Printer) {
	w.priner = p
}
