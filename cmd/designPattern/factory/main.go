package main

type Product interface {
	doStuff()
}

type ProductA struct {
}

func (pa ProductA) doStuff() {
	println("Prodoct A")
}

type ProductB struct {
}

func (pb ProductB) doStuff() {
	println("Prodoct B")
}

type Creater interface {
	createProduct() Product
}

// creatorの責務はインスタンスの生成では無くビジネスロジックの管理

type CreaterA struct {
	cache *ProductA
}

func (cA CreaterA) createProduct() Product {
	if cA.cache != nil {
		return cA.cache
	}

	cA.cache = &ProductA{}

	return cA.cache
}

type CreaterB struct {
	cache *ProductB
}

func (cB CreaterB) createProduct() Product {
	if cB.cache != nil {
		return cB.cache
	}

	cB.cache = &ProductB{}
	return cB.cache
}

func main() {
	/*
		Factory pattern
		creational patternの一種

		インスタンスを生成するFactoryのスーパークラスについてinterfaceを定義し，
		実際のインスタンスの生成はサブクラスに行わせて生成するインスタンスの型を変化させるデザインパターン
		本質はcreatorの方

		同じインターフェースで取り扱いたいけど，内部のビジネスロジックは異なる　みたいなパターンで有効そう

		Factoryクラスが生成したインスタンス（product）をclientが使う（内部の振る舞いについては関知しない）イメージっぽいわね
	*/
	createrB := CreaterB{}
	productB := createrB.createProduct()
	createrA := CreaterA{}
	productA := createrA.createProduct()

	productA.doStuff()
	productB.doStuff()

}
