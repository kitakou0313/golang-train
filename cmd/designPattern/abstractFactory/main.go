package main

type IShoe interface {
	setMark(logo string)
}

type IShirt interface {
	setMark(logo string)
}

type ISportsFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
}

func main() {

}
