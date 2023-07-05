package deck

import "antique/handy"

// Тип карты состояний сыщиков
type ConditionCard struct {
	handy.CardInterface
}

// Колода карт состояний в запасе
type ConditionDeck struct {
	handy.DeckInterface
}

func NewConsitionDeck() *ConditionDeck {
	deck := ConditionDeck{
		handy.NewDeck([]handy.CardInterface{}),
	}

	return &deck
}
