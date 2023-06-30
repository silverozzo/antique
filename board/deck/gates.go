package deck

import "antique/handy"

// Возвращает колоду врат
func NewGateStack() *handy.Deck {
	one := handy.Card{}
	deck := handy.NewDeck([]handy.Card{one})

	return deck
}
