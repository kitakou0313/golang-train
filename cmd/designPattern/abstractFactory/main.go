package main

type IShoe interface {
	getMark() string
}

type IShirt interface {
	getMark() string
}

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

type AMakerGoodsFactory struct {
	logo string
}

func (amf *AMakerGoodsFactory) makeShoe() IShoe {
	return &AMakerShoe{
		logo: amf.logo,
	}
}

type BMakerShoe struct {
	logo string
}

func (bs *BMakerShoe) getMark() string {
	return bs.logo
}

type BMakerGoodsFactory struct {
	logo string
}

func (bmf *BMakerGoodsFactory) makeShoe() IShoe {
	return &BMakerShoe{
		logo: bmf.logo,
	}
}

func main() {

}
