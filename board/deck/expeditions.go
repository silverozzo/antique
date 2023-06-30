package deck

import "antique/handy"

func NewExpeditionDeck() *handy.Deck {
	one := handy.Card{}
	deck := handy.NewDeck([]handy.Card{one})

	return deck
}
