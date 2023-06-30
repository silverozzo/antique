package deck

import "antique/handy"

func NewClueStack() *handy.Deck {
	one := handy.Card{}
	deck := handy.NewDeck([]handy.Card{one})

	return deck
}
