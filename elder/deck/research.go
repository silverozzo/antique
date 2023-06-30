package deck

import "antique/handy"

func NewResearchDeck() *handy.Deck {
	one := handy.Card{}
	deck := handy.NewDeck([]handy.Card{one})

	return deck
}
