package main

type Product interface {
	doStuff()
}

type ProductA struct {
}

func (pa *ProductA) doStuff() {
	println("Prodoct A")
}

type ProductB struct {
}

func (pb *ProductB) doStuff() {
	println("Prodoct B")
}

type Creater interface {
	createProduct() Product
}

type CreaterA struct {
}

func (cA *CreaterA) createProduct() Product {
	return &ProductA{}
}

type CreaterB struct {
}

func (cB *CreaterB) createProduct() Product {
	res := ProductB{}
	return &res
}

func main() {
	createrB := CreaterB{}
	productB := createrB.createProduct()

	createrA := CreaterA{}
	productA := createrA.createProduct()

	productA.doStuff()
	productB.doStuff()

}
