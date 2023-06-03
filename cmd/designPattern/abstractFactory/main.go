package main

type IShoe interface {
	getMark() string
}

type IShirt interface {
	getMark() string
}

// 複数のproductの生成を行うクラスのinterfaceを定義し，生成するvariation毎にfactoryサブクラスを定義する
type ISportsFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
}

type AMakerShoe struct {
	logo string
}

func (as *AMakerShoe) getMark() string {
	return as.logo
}

type AMakerShirt struct {
	logo string
}

func (asi *AMakerShirt) getMark() string {
	return asi.logo
}

type AMakerGoodsFactory struct {
	logo string
}

func (amf *AMakerGoodsFactory) makeShoe() IShoe {
	return &AMakerShoe{
		logo: amf.logo,
	}
}

func (amf *AMakerGoodsFactory) makeShirt() IShirt {
	return &AMakerShirt{
		logo: amf.logo,
	}
}

type BMakerShoe struct {
	logo string
}

func (bs *BMakerShoe) getMark() string {
	return bs.logo
}

type BMakerShirt struct {
	logo string
}

func (bsi *BMakerShirt) getMark() string {
	return bsi.logo
}

type BMakerGoodsFactory struct {
	logo string
}

func (bmf *BMakerGoodsFactory) makeShoe() IShoe {
	return &BMakerShoe{
		logo: bmf.logo,
	}
}

func (bmf *BMakerGoodsFactory) makeShirt() IShirt {
	return &BMakerShirt{
		logo: bmf.logo,
	}
}

func main() {

	arg := "B"

	var factory ISportsFactory

	if arg == "A" {
		factory = &AMakerGoodsFactory{
			logo: "A",
		}
	} else {
		factory = &BMakerGoodsFactory{
			logo: "B",
		}
	}

	shoe := factory.makeShoe()
	shirt := factory.makeShirt()
}
