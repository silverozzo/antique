package deck

import "antique/handy"

// ConditionCard Тип карты состояний сыщиков
type ConditionCard struct {
	handy.CardInterface
}

// ConditionDeck Колода карт состояний в запасе
type ConditionDeck struct {
	handy.DeckInterface
}

func NewConditionDeck() *ConditionDeck {
	deck := ConditionDeck{
		handy.NewDeck([]handy.CardInterface{}),
	}

	return &deck
}
