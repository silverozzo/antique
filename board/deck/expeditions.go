package deck

import "antique/handy"

type ExpeditionCard struct {
	handy.CardInterface
}

type ExpeditionDeck struct {
	handy.DeckInterface
}

func newExpeditionCard(locName string) *ExpeditionCard {
	return &ExpeditionCard{
		handy.NewCard(handy.Side{}, handy.Side{LocationTag: locName}),
	}
}

func NewExpeditionDeck() *ExpeditionDeck {
	one := newExpeditionCard("Глушь")
	deck := ExpeditionDeck{
		handy.NewDeck([]handy.CardInterface{one}),
	}

	return &deck
}
