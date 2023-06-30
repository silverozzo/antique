package deck

import "antique/handy"

// Возвращает колоду врат
func NewGateStack() *handy.Deck {

	// TODO пока хрен знает как сделать... болваночно
	one := handy.Card{}
	deck := handy.NewDeck([]handy.Card{one})

	return deck
}
