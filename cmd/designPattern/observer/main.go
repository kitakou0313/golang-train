package main

func main() {
	shirtItem := newItem("Nike Shirt")

	observerFirst := &Customer{
		id: "abc",
	}
	observerSecond := &Customer{
		id: "def",
	}

	shirtItem.register(observerFirst)
	shirtItem.register(observerSecond)

	shirtItem.updateAvailablity()
}
