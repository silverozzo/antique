package deck

import "antique/handy"

func NewMonsterStack() *handy.Deck {
	one := handy.Card{}
	deck := handy.NewDeck([]handy.Card{one})

	return deck
}
