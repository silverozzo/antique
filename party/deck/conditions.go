package deck

import "antique/handy"

type ConditionCard struct {
	handy.CardInterface
}

type ConditionDeck struct {
	handy.DeckInterface
}

func NewConsitionDeck() *ConditionDeck {
	deck := ConditionDeck{
		handy.NewDeck([]handy.CardInterface{}),
	}

	return &deck
}
