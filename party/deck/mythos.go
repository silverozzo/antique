package deck

import "antique/handy"

func NewMythDeck() *handy.Deck {
	one := handy.Card{}
	deck := handy.NewDeck([]handy.Card{one})

	return deck
}
