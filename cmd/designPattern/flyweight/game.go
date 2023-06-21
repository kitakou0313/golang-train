package main

type game struct {
	terrorits        []*Player
	counterTerrorist []*Player
}

func newGame() *game {
	return &game{
		terrorits:        make([]*Player, 1),
		counterTerrorist: make([]*Player, 1),
	}
}

func (c *game) addTerrorist(dressType string) {
	player := newPlayer("T", dressType)
	c.terrorits = append(c.terrorits, player)
	return
}

func (c *game) addCounterTerrorist(dressType string) {
	player := newPlayer("CT", dressType)
	c.counterTerrorist = append(c.counterTerrorist, player)
	return
}
