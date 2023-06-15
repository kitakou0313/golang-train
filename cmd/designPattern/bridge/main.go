package bridge

import "fmt"

func main() {
	hpPrinter := &Hp{}
	epsonPrinter := &Epson{}

	macCom := &Mac{}

	macCom.SetPrinter(hpPrinter)
	macCom.Print()
	fmt.Println()

	macCom.SetPrinter(epsonPrinter)
	macCom.Print()
	fmt.Println()

	winCom := &Windows{}

	winCom.SetPrinter(hpPrinter)
	winCom.Print()
	fmt.Println()

	winCom.SetPrinter(epsonPrinter)
	winCom.Print()
	fmt.Println()
}
