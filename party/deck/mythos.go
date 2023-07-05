package deck

import "antique/handy"

// Тип карты мифов
type MythCard struct {
	handy.CardInterface
}

// Колода карт мифов, подготовленная к партии
type MythDeck struct {
	handy.DeckInterface
}

func newGreenMythCard() *MythCard {
	return &MythCard{}
}

func newYellowMythCard() *MythCard {
	return &MythCard{}
}

func newBlueMythCard() *MythCard {
	return &MythCard{}
}

func NewMythDeck() *MythDeck {
	cards := []handy.CardInterface{
		newBlueMythCard(),
		newYellowMythCard(),
		newGreenMythCard(),
	}

	deck := MythDeck{
		handy.NewDeck(cards),
	}

	return &deck
}
